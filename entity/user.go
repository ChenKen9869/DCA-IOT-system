package entity

import (
	// "time"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name string `gorm:"type:varchar(20);not null"`
	Password string `gorm:"size:255;not null"`
	// Telephone string `gorm:"type:varchar(20)"`
	// Email string `gorm:"type:varchar(50)"`
	// Stat string `gorm:"type:varchar(20)"`
	// LoginDate time.Time
	// LoginIP string `gorm:"type:varchar(20)"`
}