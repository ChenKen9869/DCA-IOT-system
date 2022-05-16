package model

import "github.com/jinzhu/gorm"

type Biology struct {
	gorm.Model
	Name string `gorm:"type:varchar(20);not null"`
	InstanceID uint
	TypeBiologyID uint
}

type TypeBiology struct {
	TypeBiologyID uint `gorm:"primary_key"`
}