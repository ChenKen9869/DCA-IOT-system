package model

import "github.com/jinzhu/gorm"

type AuthMapRole struct {
	gorm.Model
	RoleID  uint
	AuthID uint
}

type Role struct {
	RoleID uint `gorm:"primary_key"`
}

type Auth struct {
	AuthID uint `gorm:"primary_key"`
}

type RoleMapUser struct {
	gorm.Model
	UserID uint
	CompanyID uint
	RanchID int
	InstanceID int
	RoleId uint
}
