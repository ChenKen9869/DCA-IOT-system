package entity

import (
	"time"

	"github.com/jinzhu/gorm"
)

type FenceRecord struct {
	gorm.Model
	Name string
	Coordinate string
	Position string
	DeviceList string
	StartTime time.Time
	EndTime time.Time
	AlarmTime uint
	ParentId uint
	Owner uint
	Stat int
}

const (
	FenceRunning = 1
	FenceFinished = 2
	FenceAbort = 3
)



