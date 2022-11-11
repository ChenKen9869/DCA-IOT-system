package entity

import (
	"github.com/jinzhu/gorm"
)

type Company struct {
	gorm.Model
	Name	string `gorm:"type:varchar(30);not null"`
	ParentID	uint
	Ancestors	string `gorm:"type:varchar(50);not null"`
	Owner	uint 
	Location string
}

type CompanyUser struct {
	ID        uint `gorm:"primary_key"`
	CompanyID uint 
	UserID uint	
}
