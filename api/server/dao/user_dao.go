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

