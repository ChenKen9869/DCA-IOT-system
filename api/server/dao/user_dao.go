package dao

import (
	"go-backend/api/common/common"
	"go-backend/api/server/entity"
)

func CreateUser(user entity.User) uint {
	common.GetDB().Create(&user)
	return user.ID
}

func GetUserInfoByName(userName string) entity.User {
	var userInfo entity.User
	common.GetDB().Where("name = ?", userName).First(&userInfo)
	return userInfo
}

func GetUserInfoById(userId uint) entity.User {
	var userInfo entity.User
	common.GetDB().Where("id = ?", userId).First(&userInfo)
	return userInfo
}

func UpdateUserInfo(userId uint, name string, password string, telephone string, email string) {
	common.GetDB().Model(&entity.User{}).Where("id = ?", userId).Update("name", name).Update("password", password).Update("telephone", telephone).Update("email", email)
}