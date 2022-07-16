package entity

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Biology struct {
	gorm.Model
	Name string
	FarmhouseID uint
	BiologyTypeID string
	Owner uint
	PicturePath string
}

type BiologyType struct {
	BiologyTypeID string `gorm:"primary_key"`
}

type EpidemicPrevention struct {
	gorm.Model
	BiologyId uint
	VaccineDescription string
	InoculationTime time.Time
}

type MedicalHistory struct {
	gorm.Model
	BiologyId uint
	Disease string
	IllnessTime time.Time
	TreatmentPlan string
}

type OperationHistory struct {
	gorm.Model
	BiologyId uint
	Doctor string
	OperationTime time.Time
	ProcessDescription string
	Result string
}