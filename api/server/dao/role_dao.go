package dao

import (
	"go-backend/api/common/db"
	"go-backend/api/server/entity"
)

func CreateVisitor(visitor entity.Visitor) uint {
	db.GetDB().Create(&visitor)
	return visitor.ID
}

func DeleteVisitorById(visitorId uint) entity.Visitor {
	var visitor entity.Visitor
	db.GetDB().Table("visitors").Where("id = ?", visitorId).First(&visitor)
	db.GetDB().Delete(&visitor)
	return visitor
}

func GetVisitorInfoExists(companyId uint, userId uint) bool {
	var visitor entity.Visitor
	db.GetDB().Table("visitors").Where("company_id = ? and user_id = ?", companyId, userId).First(&visitor)
	return visitor.ID != 0
}

func GetVisitorListByCompanyId(companyId uint) []entity.Visitor {
	var visitorList []entity.Visitor
	db.GetDB().Table("visitors").Where("company_id = ?", companyId).Find(&visitorList)
	return visitorList
}

func GetVisitor(companyId uint, userId uint) entity.Visitor {
	var visitor entity.Visitor
	db.GetDB().Table("visitors").Where("companY_id = ? and user_id = ?", companyId, userId).First(&visitor)
	return visitor
}

func GetVisitorListByUserID(userId uint) []entity.Visitor {
	var visitorList []entity.Visitor
	db.GetDB().Table("visitors").Where("user_id = ?", userId).Find(&visitorList)
	return visitorList
}