package entity

import "github.com/jinzhu/gorm"

type Visitor struct {
	gorm.Model
	CompanyId uint
	UserId uint
}
