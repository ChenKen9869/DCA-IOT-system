package dao

import (
	"go-backend/api/common/db"
	"go-backend/api/server/entity"
)

func CreateBiology(biology entity.Biology) uint {
	db.GetDB().Create(&biology)
	return biology.ID
}

func DeleteBiology(biologyId uint) entity.Biology {
	var biology entity.Biology
	db.GetDB().Where("id = ?", biologyId).First(&biology)
	db.GetDB().Delete(&biology)
	return biology
}

func CreateBiologyType(biologyType entity.BiologyType) {
	db.GetDB().Create(&biologyType)
}

func DeleteBiologyType(biologyTypeId string) entity.BiologyType {
	var biologyType entity.BiologyType
	db.GetDB().Where("biology_type_id = ?", biologyType).First(&biologyType)
	db.GetDB().Delete(&biologyType)
	return biologyType
}

func ExistBiologyType(biologyTypeId string) bool {
	var biologyType entity.BiologyType
	db.GetDB().Table("biology_types").Where("biology_type_id = ?", biologyTypeId).First(&biologyType)
	return len(biologyType.BiologyTypeID) != 0
}

func GetBiologyInfoById(biologyId uint) entity.Biology {
	var biology entity.Biology
	db.GetDB().Table("biologies").Where("id = ?", biologyId).First(&biology)
	return biology
}

func GetBiologyListByFarmhouse(farmhouseId uint) []entity.Biology {
	var biologyList []entity.Biology
	db.GetDB().Table("biologies").Where("farmhouse_id = ?", farmhouseId).Find(&biologyList)
	return biologyList
}

func UpdateBiologyFarmhouse(biologyId uint, farmhouseId uint) {
	db.GetDB().Model(&entity.Biology{}).Where("id = ?", biologyId).Update("farmhouse_id", farmhouseId)
}

func UpdateBiologyPicturePath(biologyId uint, picturePath string) {
	db.GetDB().Model(&entity.Biology{}).Where("id = ?", biologyId).Update("picture_path", picturePath)
}

func CreateEpidemicPreventionRecord(record entity.EpidemicPrevention) {
	db.GetDB().Create(&record)
}

func GetEpidemicPreventionRecordListByBiology(biologyId uint) []entity.EpidemicPrevention {
	var recordList []entity.EpidemicPrevention
	db.GetDB().Table("epidemic_preventions").Where("biology_id = ?", biologyId).Find(&recordList)
	return recordList
}

func CreateOperationRecord(record entity.OperationHistory) {
	db.GetDB().Create(&record)
}

func GetOperationRecordListByBiology(biologyId uint) []entity.OperationHistory {
	var recordList []entity.OperationHistory
	db.GetDB().Table("operation_histories").Where("biology_id = ?", biologyId).Find(&recordList)
	return recordList
}

func CreateMedicalRecord(record entity.MedicalHistory) {
	db.GetDB().Create(&record)
}

func GetMedicalRecordListByBiology(biologyId uint) []entity.MedicalHistory {
	var recordList []entity.MedicalHistory
	db.GetDB().Table("medical_histories").Where("biology_id = ?", biologyId).Find(&recordList)
	return recordList
}

func CreateBiologyChangeRecord(record entity.BiologyChange) {
	db.GetDB().Create(&record)
}

func GetOwnBiologyList(userId uint) []entity.Biology {
	var biologyList []entity.Biology
	db.GetDB().Table("biologies").Where("owner = ?", userId).Find(&biologyList)
	return biologyList
}
