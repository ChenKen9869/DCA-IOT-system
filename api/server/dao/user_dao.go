package dao

import (
	"go-backend/api/common/db"
	"go-backend/api/server/entity"
)

func CreateUser(user entity.User) uint {
	db.GetDB().Create(&user)
	return user.ID
}

func GetUserInfoByName(userName string) entity.User {
	var userInfo entity.User
	db.GetDB().Where("name = ?", userName).First(&userInfo)
	return userInfo
}

func GetUserInfoById(userId uint) entity.User {
	var userInfo entity.User
	db.GetDB().Where("id = ?", userId).First(&userInfo)
	return userInfo
}

func UpdateUserInfo(userId uint, name string, password string, telephone string, email string) {
	db.GetDB().Model(&entity.User{}).Where("id = ?", userId).Update("name", name).Update("password", password).Update("telephone", telephone).Update("email", email)
}

func UpdateUserDefaultCompany(userId uint, companyId uint) {
	db.GetDB().Model(&entity.User{}).Where("id = ?", userId).Update("default_company", companyId)
}