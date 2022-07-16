package dao

import (
	"go-backend/common"
	"go-backend/entity"
)


func CreateCompany(company entity.Company) uint {
	common.GetDB().Create(&company)
	return company.ID
}

func CreateCompanyUser(companyUser entity.CompanyUser) {
	common.GetDB().Create(&companyUser)
}

func DeleteCompanyByID(companyId uint) entity.Company {
	var company entity.Company
	common.GetDB().Where("id = ?", companyId).First(&company)
	common.GetDB().Delete(&company)
	return company
}

func DeleteCompanyUser(Id uint)  {
	common.GetDB().Delete(&entity.CompanyUser{}, Id)
}

func GetCompanyListByParent(parentId uint) []entity.Company {
	var companyList []entity.Company
	common.GetDB().Where("parent_id = ?", parentId).Find(&companyList)
	return companyList
}

func GetCompanyInfoByID(id uint) entity.Company {
	var company entity.Company
	common.GetDB().Where("id = ?", id).First(&company)
	return company
}

func GetCompanyUserInfoExists(companyId uint, userId uint) bool {
	var companyUser entity.CompanyUser
	common.GetDB().Table("company_users").Where("company_id = ? and user_id = ?", companyId, userId).First(&companyUser)
	return companyUser.ID != 0
}

func GetCompanyListByUserID(userId uint) []entity.CompanyUser {
	var companyList []entity.CompanyUser
	common.GetDB().Table("company_users").Where("user_id = ?", userId).Find(&companyList)
	return companyList
}

func GetCompanyUser(companyId uint, userId uint) entity.CompanyUser {
	var companyUser entity.CompanyUser
	common.GetDB().Table("company_users").Where("companY_id = ? and user_id = ?", companyId, userId).First(&companyUser)
	return companyUser
}

func GetUserListByCompanyId(companyId uint) []entity.CompanyUser {
	var userList []entity.CompanyUser
	common.GetDB().Table("company_users").Where("company_id = ?", companyId).Find(&userList)
	return userList
}