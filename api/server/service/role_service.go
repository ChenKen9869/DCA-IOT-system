package service

import (
	"go-backend/api/server/dao"
	"go-backend/api/server/entity"
	"go-backend/api/server/tools/server"
)

// @Summary API of golang gin backend
// @Tags Role
// @description add company visitor auth to user
// @version 1.0
// @accept mpfd
// @param CompanyId formData string true "company id"
// @param UserId formData string true "user id"
// @param Authorization header string true "token"
// @Success 200 {object} server.SuccessResponse200 "success"
// @router /role/visitor/create [post]
func CreateVisitorService(companyId uint, userId uint) {
	companyInfo := dao.GetCompanyInfoByID(companyId)
	if companyInfo.ID == 0 {
		panic(server.CompanyNotExist)
	}
	userInfo := dao.GetUserInfoById(userId)
	if userInfo.ID == 0 {
		panic(server.UserNotExist)
	}
	companyList := dao.GetCompanyListByUserID(userId)
	ancestorList, errAtoi := GetAncestorsList((dao.GetCompanyInfoByID(companyId)).Ancestors)
	if errAtoi != nil {
		panic("atoi error")
	}
	for _, company := range companyList {
		if company.CompanyID == companyId {
			panic("permission already exists")
		}
		for _, ancestorId := range ancestorList {
			if company.CompanyID == ancestorId {
				panic("permission already exists")
			}
		}
	}
	visitorList := dao.GetVisitorListByCompanyId(userId)
	ancestorList, errAtoi = GetAncestorsList((dao.GetCompanyInfoByID(companyId)).Ancestors)
	if errAtoi != nil {
		panic(errAtoi.Error())
	}
	for _, visitor := range visitorList {
		if visitor.CompanyId == companyId {
			panic("permission already exists")
		}
		for _, ancestorId := range ancestorList {
			if visitor.CompanyId == ancestorId {
				panic("permission already exists")
			}
		}
	}
	dao.CreateVisitor(entity.Visitor{
		CompanyId: companyId,
		UserId:    userId,
	})
}

// @Summary API of golang gin backend
// @Tags Role
// @description delete company visitor auth of user
// @version 1.0
// @accept application/json
// @param CompanyId query int true "company id"
// @param UserId query string true "user id"
// @param Authorization header string true "token"
// @Success 200 {object} server.SuccessResponse200 "success"
// @router /role/visitor/delete [delete]
func DeleteVisitorService(companyId uint, userId uint) {
	visitor := dao.GetVisitor(companyId, userId)
	dao.DeleteVisitorById(visitor.ID)
}

// @Summary API of golang gin backend
// @Tags Role
// @description get visitor list of company
// @version 1.0
// @accept application/json
// @param CompanyId query string true "company id"
// @param Authorization header string true "token"
// @Success 200 {object} server.SuccessResponse200 "success"
// @router /role/visitor/get_list [get]
func GetVisitorListService(companyId uint) map[entity.User]([]uint) {
	visitorList := make(map[entity.User]([]uint))
	GetVisitorRecursive(companyId, visitorList)
	return visitorList
}

func GetVisitorRecursive(companyId uint, visitorList map[entity.User]([]uint)) {
	userList := dao.GetVisitorListByCompanyId(companyId)
	for _, user := range userList {
		visitor := dao.GetUserInfoById(user.UserId)
		visitorList[visitor] = append(visitorList[visitor], companyId)
	}
	childrenList := dao.GetCompanyListByParent(companyId)
	for _, subCompany := range childrenList {
		GetVisitorRecursive(subCompany.ID, visitorList)
	}
}

// @Summary API of golang gin backend
// @Tags Role
// @description get user's visitor company list
// @version 1.0
// @accept application/json
// @param Authorization header string true "token"
// @Success 200 {object} server.SuccessResponse200 "success"
// @router /role/visitor/get_company_list [get]
func GetVisitorCompanyListService(userId uint) []uint {
	companyIdList := []uint{}
	visitorList := dao.GetVisitorListByUserID(userId)
	for _, visitor := range visitorList {
		companyIdList = append(companyIdList, visitor.CompanyId)
	}
	return companyIdList
}