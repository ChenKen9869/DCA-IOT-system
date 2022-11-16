package service

import (
	"errors"
	"go-backend/api/server/dao"
	"go-backend/api/server/entity"
	"go-backend/api/server/tools/server"
	"go-backend/api/server/vo"
	"strconv"
	"strings"
)

// 公司与用户的权限验证
func AuthCompanyUser(userId uint, companyId uint) bool {
	ancestorList, errAtoi := GetAncestorsList((dao.GetCompanyInfoByID(uint(companyId))).Ancestors)
	if errAtoi != nil {
		panic(errAtoi.Error())
	}
	userCompanyList := dao.GetCompanyListByUserID(userId)
	for _, userCompany := range userCompanyList {
		if userCompany.CompanyID == uint(companyId) {
			return true
		}
		for _, ancestorId := range ancestorList {
			if userCompany.CompanyID == ancestorId {
				return true
			}
		}
	}
	return false
}

func AuthVisitor(userId uint, companyId uint) bool {
	ancestorList, errAtoi := GetAncestorsList((dao.GetCompanyInfoByID(uint(companyId))).Ancestors)
	if errAtoi != nil {
		panic(errAtoi.Error())
	}
	visitorList := dao.GetVisitorListByUserID(userId)
	for _, visitor := range visitorList {
		if visitor.CompanyId == uint(companyId) {
			return true
		}
		for _, ancestorId := range ancestorList {
			if visitor.CompanyId == ancestorId {
				return true
			}
		}
	}
	return false
}

// @Summary API of golang gin backend
// @Tags Company
// @description delete company : 删除一个公司 参数列表：[公司ID] 访问携带token
// @version 1.0
// @accept application/json
// @param CompanyId query int true "company id"
// @param Authorization header string true "token"
// @Success 200 {object} server.SuccessResponse200 "成功"
// @router /company/delete [delete]
func DeleteCompanyService(companyId uint, operator entity.User) error {
	companyNode := dao.GetCompanyInfoByID(companyId)
	if companyNode.ID == 0 {
		err := errors.New(server.CompanyNotExist)
		return err
	}
	childNode := dao.GetCompanyListByParent(companyId)
	if len(childNode) != 0 {
		err := errors.New(server.NodeHasSubcompany)
		return err
	}
	biologyList := dao.GetBiologyListByFarmhouse(companyId)
	for _, biology := range biologyList {
		DeleteBiologyService(operator.Name, operator.Telephone, "farmhouse been deleted", biology.ID)
	}
	fixedDeviceList := dao.GetFixedDeviceListByFarmhouse(companyId)
	for _, fixedDevice := range fixedDeviceList {
		dao.DeleteFixedDevice(fixedDevice.ID)
	}
	managerList := dao.GetUserListByCompanyId(companyId)
	visitorList := dao.GetVisitorListByCompanyId(companyId)
	for _, manager := range managerList {
		dao.DeleteCompanyUser(manager.ID)
	}
	for _, visitor := range visitorList {
		dao.DeleteVisitorById(visitor.ID)
	}
	dao.DeleteCompanyByID(companyId)
	return nil
}

// 用 id 获得 tree
func makeTreeByCompanyId(currentId uint) vo.CompanyTreeNode {
	root := vo.CompanyTreeNode{}
	currentNode := dao.GetCompanyInfoByID(currentId)
	ancestorList, err := GetAncestorsList(currentNode.Ancestors)
	if err != nil {
		panic("error occurs!")
	}
	point := &root
	if len(ancestorList) > 1 {
		// len(ancestors) != 1, 则当前节点不是机构树的根节点
		// 用 ancestor_list 的第二个元素初始化根节点（第一个元素是标识）
		rootNode := dao.GetCompanyInfoByID(ancestorList[1])
		root.Id = ancestorList[1]
		root.Name = rootNode.Name
		root.Ancestors = rootNode.Ancestors
		root.Children = []vo.CompanyTreeNode{}
		// 挂载其余的 ancestor
		for i := 2; i < len(ancestorList); i++ {
			ancestorNode := dao.GetCompanyInfoByID(ancestorList[i])
			node := vo.CompanyTreeNode{
				Id:        ancestorNode.ID,
				Name:      ancestorNode.Name,
				Ancestors: ancestorNode.Ancestors,
				Children:  []vo.CompanyTreeNode{},
			}
			(*point).Children = append((*point).Children, node)
			point = &(((*point).Children)[0])
		}
		// 初始化并挂载当前节点
		current := vo.CompanyTreeNode{
			Id:        currentNode.ID,
			Name:      currentNode.Name,
			Ancestors: currentNode.Ancestors,
			Children:  nil,
		}
		(*point).Children = append((*point).Children, current)
		// 让 point 指向当前节点
		point = &(((*point).Children)[0])
	} else {
		// len(ancestors) == 1, 则当前节点是机构树的根节点
		// 用当前节点初始化机构树根节点
		root.Id = currentId
		root.Name = currentNode.Name
		root.Ancestors = currentNode.Ancestors
		root.Children = nil
	}
	// 从当前节点开始构造孩子树，此时 point 指向当前节点
	(*point).Children = makeChildrenTreeListRecursive((*point).Id)
	return root
}

func makeChildrenTreeListRecursive(parentId uint) []vo.CompanyTreeNode {
	// 查数据库，找到所有第一层的孩子节点
	childrenList := dao.GetCompanyListByParent(parentId)
	// 递归出口：已经没有孩子了
	if len(childrenList) == 0 {
		return []vo.CompanyTreeNode{}
	}
	// 初始化孩子列表
	var children []vo.CompanyTreeNode
	// 遍历孩子节点列表
	for _, child := range childrenList {
		// 初始化孩子
		childNode := vo.CompanyTreeNode{
			Id:        child.ID,
			Name:      child.Name,
			Ancestors: child.Ancestors,
			Children:  makeChildrenTreeListRecursive(child.ID),
		}
		// 挂载孩子到孩子列表
		children = append(children, childNode)
	}
	// 返回孩子列表
	return children
}

// @Summary API of golang gin backend
// @Tags Company
// @description get user's company tree : 获取当前用户有权限的所有公司信息（以树形结构返回） 参数列表：[] 访问携带token
// @version 1.0
// @accept application/json
// @param Authorization header string true "token"
// @Success 200 {object} server.SuccessResponse200 "成功"
// @router /company/get/treelist [get]
func GetCompanyTreeListService(userId uint) ([]vo.CompanyTreeNode, []uint) {
	// 根据 user_id 获取 company list : dao
	companies := dao.GetCompanyListByUserID(userId)
	var companyList []uint
	for _, company := range companies {
		companyList = append(companyList, company.CompanyID)
	}
	// 构造 tree list
	var treeList []vo.CompanyTreeNode

	// 遍历 company list ，得到每颗树，并加入 tree list
	for _, companyId := range companyList {
		root := makeTreeByCompanyId(companyId)
		treeList = append(treeList, root)
	}
	// 返回 tree list

	return treeList, companyList
	// 1. 使用 user_id 获取 comnpany list : dao
	// 2. 使用 company_id 构造树

}

func GetAncestorsList(ancestors string) ([]uint, error) {
	var ancestorList []uint
	ancestorStringList := strings.Split(ancestors, ",")
	for _, ancestorString := range ancestorStringList {
		ancestor, err := strconv.Atoi(ancestorString)
		if err != nil {
			return ancestorList, err
		}
		ancestorList = append(ancestorList, uint(ancestor))
	}
	return ancestorList, nil
}

// @Summary API of golang gin backend
// @Tags Company
// @description create company : 创建一个公司 参数列表：[公司名称、该公司的父公司ID（root公司的父公司ID填写0）、该公司的地理位置信息描述（前端自己决定格式，具体看第三方天气定位等服务的接口要求，后端只负责保存地理信息，不做其他处理）] 访问携带token
// @version 1.0
// @accept mpfd
// @param Name formData string true "company_name"
// @param ParentId formData string true "parent id"
// @param Location formData string true "location"
// @param Authorization header string true "token"
// @Success 200 {object} server.SuccessResponse200 "成功"
// @router /company/create [post]
func CreateCompanyService(parentId uint, name string, owner uint, location string) (uint, error) {
	// 判断名字是否为空
	if len(name) == 0 {
		err := errors.New("name is null")
		return 0, err
	}
	// 如果 parent_id 是 0，说明是新机构
	if parentId == 0 {
		newCompany := entity.Company{
			Ancestors: "0",
			Name:      name,
			Owner:    owner,
			ParentID:  0,
			Location: location,
		}
		id := dao.CreateCompany(newCompany)
		return id, nil
	}
	// 根据 parentID 去查表：GetCompanyByID，得到parent
	parent := dao.GetCompanyInfoByID(parentId)
	newCompany := entity.Company{
		Ancestors: parent.Ancestors + "," + strconv.Itoa(int(parentId)),
		Name:      name,
		Owner:     parent.Owner,
		ParentID:  parentId,
		Location:  location,
	}
	id := dao.CreateCompany(newCompany)
	return id, nil
}

// @Summary API of golang gin backend
// @Tags Company
// @description add company auth to user : 为指定用户分配指定公司的权限（接口访问者需要事先拥有该公司的权限） 参数列表：[公司ID、用户ID] 访问携带token
// @version 1.0
// @accept mpfd
// @param CompanyId formData string true "company id"
// @param UserId formData string true "user id"
// @param Authorization header string true "token"
// @Success 200 {object} server.SuccessResponse200 "成功"
// @router /company/company_user/create [post]
func CreateCompanyUserService(companyId uint, userId uint) error {
	// 查 companyId , 和 userId 是否存在
	companyInfo := dao.GetCompanyInfoByID(companyId)
	if companyInfo.ID == 0 {
		err := errors.New(server.CompanyNotExist)
		return err
	}
	userInfo := dao.GetUserInfoById(userId)
	if userInfo.ID == 0 {
		err := errors.New(server.UserNotExist)
		return err
	}

	// 如果要创建的这个权限，用户已经拥有，则报错
	// 查目标 company 是否在用户的权限列表中，以及目标company的祖先中是否有某个节点在用户的权限列表
	companyList := dao.GetCompanyListByUserID(userId)
	ancestorList, errAtoi := GetAncestorsList((dao.GetCompanyInfoByID(companyId)).Ancestors)
	if errAtoi != nil {
		panic("atoi error")
	}
	for _, company := range companyList {
		if company.CompanyID == companyId {
			err := errors.New("权限已经存在")
			return err
		}
		for _, ancestorId := range ancestorList {
			if company.CompanyID == ancestorId {
				err := errors.New("权限已经存在")
				return err
			}
		}
	}

	companyUserInfo := entity.CompanyUser{
		CompanyID: companyId,
		UserID: userId,
	}
	dao.CreateCompanyUser(companyUserInfo)
	return nil
}

// @Summary API of golang gin backend
// @Tags Company
// @description delete company auth of user : 从指定用户处收回指定公司的权限 参数列表：[公司ID、用户ID] 访问携带token
// @version 1.0
// @accept application/json
// @param CompanyId query int true "company id"
// @param UserId formData string true "user id"
// @param Authorization header string true "token"
// @Success 200 {object} server.SuccessResponse200 "成功"
// @router /company/company_user/delete [delete]
func DeleteCompanyUserService(companyId uint, userId uint) {
	companyUser := dao.GetCompanyUser(companyId, userId)
	dao.DeleteCompanyUser(companyUser.ID)
}

// @Summary API of golang gin backend
// @Tags Company
// @description get employee list of company : 获取公司的员工列表 参数列表：[公司ID] 访问携带token
// @version 1.0
// @accept application/json
// @param CompanyId query string true "company id"
// @param Authorization header string true "token"
// @Success 200 {object} server.SuccessResponse200 "成功"
// @router /company/get/employeelist [get]
func GetEmployeeListService(companyId uint) map[entity.User][]uint {
	employeeList := make(map[entity.User][]uint)
	// 找到公司的所有子公司
	// 遍历该公司，以及该公司的所有子公司，在公司-员工表中找到匹配的信息，将信息填入雇员表
	// 雇员表： key：雇员id，value：该雇员有操作权的公司id	
	GetEmployeeRecursive(companyId, employeeList)
	return employeeList
}

func GetEmployeeRecursive(companyId uint, employeeList map[entity.User][]uint) {
	userList := dao.GetUserListByCompanyId(companyId)
	for _, user := range userList {
		employee := dao.GetUserInfoById(user.UserID)
		employeeList[employee] = append(employeeList[employee], companyId)
	}
	childrenList := dao.GetCompanyListByParent(companyId)
	for _, subCompany := range childrenList {
		GetEmployeeRecursive(subCompany.ID, employeeList)
	}
}

// @Summary API of golang gin backend
// @Tags Company
// @description get company information : 获取公司的详细信息 参数列表：[公司ID] 访问携带token
// @version 1.0
// @accept application/json
// @param CompanyId query string true "company id"
// @param Authorization header string true "token"
// @Success 200 {object} server.SuccessResponse200 "成功"
// @router /company/get/info [get]
func GetCompanyInfoService(companyId uint) vo.CompanyInfo {
	company := dao.GetCompanyInfoByID(companyId)
	companyInfo := vo.CompanyInfo{
		Id: company.ID,
		Name: company.Name,
		Location: company.Location,
	}
	return companyInfo
}