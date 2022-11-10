package dao

import (
	"go-backend/common"
	"go-backend/entity"
)

// func CountBiologyByUserId(userId uint) int {
// 	var count int
// 	common.GetDB().Table("biologies").Where("owner = ?", userId).Count(&count)
// 	return count
// }

func CreateBiology(biology entity.Biology) uint {
	common.GetDB().Create(&biology)
	return biology.ID
}

func DeleteBiology(biologyId uint) entity.Biology {
	var biology entity.Biology
	common.GetDB().Where("id = ?", biologyId).First(&biology)
	common.GetDB().Delete(&biology)
	return biology
}

func CreateBiologyType(biologyType entity.BiologyType) {
	common.GetDB().Create(&biologyType)
}

func DeleteBiologyType(biologyTypeId string) entity.BiologyType {
	var biologyType entity.BiologyType
	common.GetDB().Where("biology_type_id = ?", biologyType).First(&biologyType)
	common.GetDB().Delete(&biologyType)
	return biologyType
}


func ExistBiologyType(biologyTypeId string) bool {
	var biologyType entity.BiologyType
	common.GetDB().Table("biology_types").Where("biology_type_id = ?", biologyTypeId).First(&biologyType)
	return len(biologyType.BiologyTypeID) != 0
}

func GetBiologyInfoById(biologyId uint) entity.Biology {
	var biology entity.Biology
	common.GetDB().Table("biologies").Where("id = ?", biologyId).First(&biology)
	return biology
}

func GetBiologyListByFarmhouse(farmhouseId uint) []entity.Biology {
	var biologyList []entity.Biology
	common.GetDB().Table("biologies").Where("farmhouse_id = ?", farmhouseId).Find(&biologyList)
	return biologyList
}

func UpdateBiologyFarmhouse(biologyId uint, farmhouseId uint) {
	common.GetDB().Model(&entity.Biology{}).Where("id = ?", biologyId).Update("farmhouse_id", farmhouseId)
}

func UpdateBiologyPicturePath(biologyId uint, picturePath string) {
	common.GetDB().Model(&entity.Biology{}).Where("id = ?", biologyId).Update("picture_path", picturePath)
}

func CreateEpidemicPreventionRecord(record entity.EpidemicPrevention) {
	common.GetDB().Create(&record)
}

func GetEpidemicPreventionRecordListByBiology(biologyId uint) []entity.EpidemicPrevention {
	var recordList []entity.EpidemicPrevention
	common.GetDB().Table("epidemic_preventions").Where("biology_id = ?", biologyId).Find(&recordList)
	return recordList
}

func CreateOperationRecord(record entity.OperationHistory) {
	common.GetDB().Create(&record)
}

func GetOperationRecordListByBiology(biologyId uint) []entity.OperationHistory {
	var recordList []entity.OperationHistory
	common.GetDB().Table("operation_histories").Where("biology_id = ?", biologyId).Find(&recordList)
	return recordList
}

func CreateMedicalRecord(record entity.MedicalHistory) {
	common.GetDB().Create(&record)
}

func GetMedicalRecordListByBiology(biologyId uint) []entity.MedicalHistory {
	var recordList []entity.MedicalHistory
	common.GetDB().Table("medical_histories").Where("biology_id = ?", biologyId).Find(&recordList)
	return recordList
}

func CreateBiologyChangeRecord(record entity.BiologyChange) {
	common.GetDB().Create(&record)
}