package entity

import (
	"github.com/jinzhu/gorm"
)

type PortableDevice struct {
	gorm.Model
	DeviceID string `gorm:"varchar(50);not null"`
	BiologyID uint
	PortableDeviceTypeID string
	Owner uint
}

type FixedDevice struct {
	gorm.Model
	DeviceID string `gorm:"varchar(50);not null"`
	FarmhouseID uint
	FixedDeviceTypeID string
	Owner uint
}

type FixedDeviceType struct {
	FixedDeviceTypeID string `gorm:"primary_key"`
}

type PortableDeviceType struct {
	PortableDeviceTypeID string `gorm:"primary_key"`
}
