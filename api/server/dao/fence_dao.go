package dao

// import (
// 	"go-backend/api/common/common"
// 	"go-backend/api/server/entity"
// 	"time"
// 	"github.com/jinzhu/gorm"
// )

// func CreateFenceRecord(fenceRecord entity.FenceRecord) uint {
// 	common.GetDB().Create(&fenceRecord)
// 	return fenceRecord.ID
// }

// func AddAlarmTime(fenceId uint) {
// 	common.GetDB().Model(&entity.FenceRecord{}).Where("id = ? ", fenceId).Update("alarm_time", gorm.Expr("alarm_time + ?", 1))
// }

// func ModifyFenceStat(fenceId uint, stat int) {
// 	common.GetDB().Model(&entity.FenceRecord{}).Where("id = ?", fenceId).Update("stat", stat)
// }

// func UpdateFenceEndTime(fenceId uint, endTime time.Time) {
// 	common.GetDB().Model(&entity.FenceRecord{}).Where("id = ?", fenceId).Update("end_time", endTime)
// }


// func GetFenceRecordById(fenceId uint) entity.FenceRecord {
// 	var fenceRecord entity.FenceRecord
// 	common.GetDB().Where("id = ?", fenceId).First(&fenceRecord)
// 	return fenceRecord
// }

// func GetActiveFenceListByUserId(userId uint) []entity.FenceRecord {
// 	var fenceRecordList []entity.FenceRecord
// 	common.GetDB().Where("owner = ? and stat = ?", userId, entity.FenceRunning).Find(&fenceRecordList)
// 	return fenceRecordList
// }

// func GetActiveFenceListByCompanyId(companyId uint) []entity.FenceRecord {
// 	var fenceRecordList []entity.FenceRecord
// 	common.GetDB().Where("parent_id = ? and stat = ?", companyId, entity.FenceRunning).Find(&fenceRecordList)
// 	return fenceRecordList
// }

// func GetFenceRecordListByCompanyId(companyId uint) []entity.FenceRecord {
// 	var fenceRecordList []entity.FenceRecord
// 	common.GetDB().Where("parent_id = ? and stat = ?", companyId, entity.FenceFinished).Find(&fenceRecordList)
// 	return fenceRecordList
// }