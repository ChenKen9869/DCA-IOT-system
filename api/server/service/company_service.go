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
		DeleteCompanyUserService(manager.CompanyID, manager.UserID)
	}
	for _, visitor := range visitorList {
		dao.DeleteVisitorById(visitor.ID)
	}
	dao.DeleteCompanyByID(companyId)
	return nil
}

func makeChildrenTreeListRecursive(parentId uint) []vo.CompanyTreeNode {
	childrenList := dao.GetCompanyListByParent(parentId)
	if len(childrenList) == 0 {
		return []vo.CompanyTreeNode{}
	}
	var children []vo.CompanyTreeNode
	for _, child := range childrenList {
		childNode := vo.CompanyTreeNode{
			Id:        child.ID,
			Name:      child.Name,
			Location:  child.Location,
			Owner:     child.Owner,
			Ancestors: child.Ancestors,
			Children:  makeChildrenTreeListRecursive(child.ID),
		}
		children = append(children, childNode)
	}
	return children
}

func makeTreeByCompanyId(currentId uint) vo.CompanyTreeNode {
	root := vo.CompanyTreeNode{}
	currentNode := dao.GetCompanyInfoByID(currentId)
	ancestorList, err := GetAncestorsList(currentNode.Ancestors)
	if err != nil {
		panic("error occurs!")
	}
	point := &root
	if len(ancestorList) > 1 {
		rootNode := dao.GetCompanyInfoByID(ancestorList[1])
		root.Id = ancestorList[1]
		root.Name = rootNode.Name
		root.Ancestors = rootNode.Ancestors
		root.Children = []vo.CompanyTreeNode{}
		for i := 2; i < len(ancestorList); i++ {
			ancestorNode := dao.GetCompanyInfoByID(ancestorList[i])
			node := vo.CompanyTreeNode{
				Id:        ancestorNode.ID,
				Name:      ancestorNode.Name,
				Location:  ancestorNode.Location,
				Owner:     ancestorNode.Owner,
				Ancestors: ancestorNode.Ancestors,
				Children:  []vo.CompanyTreeNode{},
			}
			(*point).Children = append((*point).Children, node)
			point = &(((*point).Children)[0])
		}
		current := vo.CompanyTreeNode{
			Id:        currentNode.ID,
			Name:      currentNode.Name,
			Location:  currentNode.Location,
			Owner:     currentNode.Owner,
			Ancestors: currentNode.Ancestors,
			Children:  nil,
		}
		(*point).Children = append((*point).Children, current)
		point = &(((*point).Children)[0])
	} else {
		root.Id = currentId
		root.Name = currentNode.Name
		root.Location = currentNode.Location
		root.Ancestors = currentNode.Ancestors
		root.Owner = currentNode.Owner
		root.Children = nil
	}
	(*point).Children = makeChildrenTreeListRecursive((*point).Id)
	return root
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
	companies := dao.GetCompanyListByUserID(userId)
	var companyList []uint
	for _, company := range companies {
		companyList = append(companyList, company.CompanyID)
	}
	var treeList []vo.CompanyTreeNode
	for _, companyId := range companyList {
		root := makeTreeByCompanyId(companyId)
		treeList = append(treeList, root)
	}
	return treeList, companyList
}

// @Summary API of golang gin backend
// @Tags Company
// @description create company : 创建一个公司 参数列表：[公司名称、该公司的父公司ID（root公司的父公司ID填写0）、该公司的地理位置信息描述（前端自己决定格式，具体看第三方天气定位等服务的接口要求，后端只负责保存地理信息，不做其他处理）] 访问携带token
// @version 1.0
// @accept mpfd
// @param Name formData string true "company_name"
// @param ParentId formData int true "parent id"
// @param Location formData string true "location"
// @param Authorization header string true "token"
// @Success 200 {object} server.SuccessResponse200 "成功"
// @router /company/create [post]
func CreateCompanyService(parentId uint, name string, owner uint, location string) (uint, error) {
	if len(name) == 0 {
		err := errors.New("name is null")
		return 0, err
	}
	if parentId == 0 {
		newCompany := entity.Company{
			Ancestors: "0",
			Name:      name,
			Owner:     owner,
			ParentID:  0,
			Location:  location,
		}
		id := dao.CreateCompany(newCompany)
		return id, nil
	}
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
		UserID:    userId,
	}
	dao.CreateCompanyUser(companyUserInfo)
	// l := dao.GetOwnCompanyList(userId)
	// _, l := GetCompanyTreeListService(userId)
	defaultCompany := dao.GetUserInfoById(userId).DefaultCompany
	if defaultCompany == 0 {
		dao.UpdateUserDefaultCompany(userId, companyId)
	}
	return nil
}

// @Summary API of golang gin backend
// @Tags Company
// @description delete company auth of user : 从指定用户处收回指定公司的权限 参数列表：[公司ID、用户ID] 访问携带token
// @version 1.0
// @accept application/json
// @param CompanyId query int true "company id"
// @param UserId query string true "user id"
// @param Authorization header string true "token"
// @Success 200 {object} server.SuccessResponse200 "成功"
// @router /company/company_user/delete [delete]
func DeleteCompanyUserService(companyId uint, userId uint) {
	companyUser := dao.GetCompanyUser(companyId, userId)
	dao.DeleteCompanyUser(companyUser.ID)
	if companyId == dao.GetUserInfoById(userId).DefaultCompany {
		_, l := GetCompanyTreeListService(userId)
		if len(l) == 0 {
			UpdateUserDefaultCompanyService(userId, 0)
		} else {
			UpdateUserDefaultCompanyService(userId, l[0])
		}
	}
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
		Id:       company.ID,
		Name:     company.Name,
		Location: company.Location,
	}
	return companyInfo
}

// @Summary API of golang gin backend
// @Tags Company
// @description get own company list : 获取当前用户拥有的公司列表 参数列表：[] 访问携带token
// @version 1.0
// @accept application/json
// @param Authorization header string true "token"
// @Success 200 {object} server.SuccessResponse200 "成功"
// @router /company/get/own_list [get]
func GetOwnCompanyListService(userId uint) []vo.CompanyTreeNode {
	companyList := dao.GetOwnCompanyList(userId)
	var treeList []vo.CompanyTreeNode
	for _, company := range companyList {
		root := makeTreeByCompanyId(company.ID)
		treeList = append(treeList, root)
	}
	return treeList
}

// @Summary API of golang gin backend
// @Tags Company
// @description update company info : 更新公司信息 参数列表：[公司ID，新名字，新地址] 访问携带token
// @version 1.0
// @accept application/json
// @param Name query string true "company_name"
// @param Location query string true "location"
// @param CompanyId query int true "company id"
// @param Authorization header string true "token"
// @Success 200 {object} server.SuccessResponse200 "成功"
// @router /company/update [put]
func UpdateCompanyInfoService(companyId uint, name string, location string) {
	dao.UpdateCompanyInfo(companyId, name, location)
}
