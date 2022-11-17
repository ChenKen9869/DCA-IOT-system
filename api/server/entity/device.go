package entity

import (
	"time"
	"github.com/jinzhu/gorm"
)

type PortableDevice struct {
	gorm.Model
	DeviceID string `gorm:"varchar(50);not null"`
	BiologyID uint
	PortableDeviceTypeID string
	Owner uint
	BoughtTime time.Time
	InstallTime time.Time
	Stat string
}

type FixedDevice struct {
	gorm.Model
	DeviceID string `gorm:"varchar(50);not null"`
	FarmhouseID uint
	FixedDeviceTypeID string
	Owner uint
	BoughtTime time.Time
	InstallTime time.Time
	Stat string
}

type FixedDeviceType struct {
	FixedDeviceTypeID string `gorm:"primary_key"`
}

type PortableDeviceType struct {
	PortableDeviceTypeID string `gorm:"primary_key"`
}