package model

import "github.com/jinzhu/gorm"

type PortableEquipment struct {
	gorm.Model
	BiologyID uint
	TypePortableEID uint
}

type FixedEquipment struct {
	gorm.Model
	InstanceID uint
	TypeFixedEID uint
}

type TypeFixed struct {
	TypeFixedEID uint `gorm:"primary_key"`
}

type TypePortable struct {
	TypePortableEID uint `gorm:"primary_key"`
}