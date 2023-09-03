package dao

import (
	"go-backend/api/common/db"
	"go-backend/api/server/entity"
)

func CreateCompany(company entity.Company) uint {
	db.GetDB().Create(&company)
	return company.ID
}

func CreateCompanyUser(companyUser entity.CompanyUser) {
	db.GetDB().Create(&companyUser)
}

func DeleteCompanyByID(companyId uint) entity.Company {
	var company entity.Company
	db.GetDB().Where("id = ?", companyId).First(&company)
	db.GetDB().Delete(&company)
	return company
}

func DeleteCompanyUser(Id uint) {
	db.GetDB().Delete(&entity.CompanyUser{}, Id)
}

func GetCompanyListByParent(parentId uint) []entity.Company {
	var companyList []entity.Company
	db.GetDB().Where("parent_id = ?", parentId).Find(&companyList)
	return companyList
}

func GetCompanyInfoByID(id uint) entity.Company {
	var company entity.Company
    db.GetDB().Where("id = ?", id).First(&company)
    return company
}

func GetCompanyUserInfoExists(companyId uint, userId uint) bool {
	var companyUser entity.CompanyUser
	db.GetDB().Table("company_users").Where("company_id = ? and user_id = ?", companyId, userId).First(&companyUser)
	return companyUser.ID != 0
}

func GetCompanyListByUserID(userId uint) []entity.CompanyUser {
	var companyList []entity.CompanyUser
	db.GetDB().Table("company_users").Where("user_id = ?", userId).Find(&companyList)
	return companyList
}

func GetCompanyUser(companyId uint, userId uint) entity.CompanyUser {
	var companyUser entity.CompanyUser
	db.GetDB().Table("company_users").Where("companY_id = ? and user_id = ?", companyId, userId).First(&companyUser)
	return companyUser
}

func GetUserListByCompanyId(companyId uint) []entity.CompanyUser {
	var userList []entity.CompanyUser
	db.GetDB().Table("company_users").Where("company_id = ?", companyId).Find(&userList)
	return userList
}

func GetOwnCompanyList(userId uint) []entity.Company {
	var companyList []entity.Company
	db.GetDB().Table("companies").Where("owner = ? and parent_id = 0", userId).Find(&companyList)
	return companyList
}

func UpdateCompanyInfo(companyId uint, companyName string, location string) {
	db.GetDB().Table("companies").Where("id = ?", companyId).Update("name", companyName).Update("location", location)
}
