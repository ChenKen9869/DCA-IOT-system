package dao

import (
	"go-backend/api/common/common"
	"go-backend/api/server/entity"
)

// UserWriteInDatabase 将用户写入数据库
func CreateUser(user entity.User) uint {
	common.GetDB().Create(&user)
	return user.ID
}


// GetUserInfoByName 根据用户名获取用户
func GetUserInfoByName(userName string) entity.User {
	var userInfo entity.User
	common.GetDB().Where("name = ?", userName).First(&userInfo)
	return userInfo
}

// GetUserInfoById 根据用户 id 获取用户
func GetUserInfoById(userId uint) entity.User {
	var userInfo entity.User
	common.GetDB().Where("id = ?", userId).First(&userInfo)
	return userInfo
}

