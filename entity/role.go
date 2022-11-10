package entity

import "github.com/jinzhu/gorm"

// type AuthMapRole struct {
// 	gorm.Model
// 	RoleID  uint
// 	AuthID uint
// }

// type Role struct {
// 	RoleID uint `gorm:"primary_key"`
// 	RoleName string `gorm:"varchar(30);not null"`
// }

// type Auth struct {
// 	AuthID uint `gorm:"primary_key"`
// }

// type RoleMapUser struct {
// 	gorm.Model
// 	UserID uint
// 	CompanyID uint
// 	RanchID int
// 	InstanceID int
// 	RoleId uint
// }

type Visitor struct {
	gorm.Model
	CompanyId uint
	UserId uint
}
