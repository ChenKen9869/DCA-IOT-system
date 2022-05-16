package model

import (

	"github.com/jinzhu/gorm"
)

type Dept struct {
	gorm.Model
	Name	string `gorm:"type:varchar(30);not null"`
	ParentID	uint
	Ancestors	string `gorm:"type:varchar(50);not null"`
	OrderNum	uint
	Leader	string `gorm:"type:varchar(20);not null"`
}

/*

type Company struct {
	gorm.Model
	Name string `gorm:"type:varchar(20);not null"`
}

type Ranch struct {
	gorm.Model
	Name string `gorm:"type:varchar(20);not null"`
	CompanyID uint
}

type Instance struct {
	gorm.Model
	Name string `gorm:"type:varchar(20);not null"`
	RanchID uint
}

type TypeInstance struct {
	TypeInstanceID uint `gorm:"primary_key"`
}

*/