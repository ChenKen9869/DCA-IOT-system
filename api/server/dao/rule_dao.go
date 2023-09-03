package dao

import (
	"go-backend/api/common/db"
	"go-backend/api/server/entity"
)

func CreateRule(rule entity.Rule) uint {
	db.GetDB().Create(&rule)
	return rule.ID
}

func DeleteRule(ruleId uint) entity.Rule {
	var rule entity.Rule
	db.GetDB().Model(&entity.Rule{}).Where("id = ?", ruleId).First(&rule)
	db.GetDB().Delete(&rule)
	return rule
}

func UpdateRuleStat(ruleId uint, newStat string) {
	db.GetDB().Model(&entity.Rule{}).Where("id = ?", ruleId).Update("stat", newStat)
}

func GetRuleInfo(ruleId uint) entity.Rule {
	var rule entity.Rule
	db.GetDB().Table("rules").Where("id = ?", ruleId).Find(&rule)
	return rule
}

func GetRuleListByUser(userId uint) []entity.Rule {
	var ruleList []entity.Rule
	db.GetDB().Table("rules").Where("owner = ?", userId).Find(&ruleList)
	return ruleList
}

func GetRuleListByCompany(companyId uint) []entity.Rule {
	var ruleList []entity.Rule
	db.GetDB().Table("rules").Where("parent_id = ?", companyId).Find(&ruleList)
	return ruleList
}

func UpdateRuleDCA(ruleId uint, datasource string, condition string, action string) {
	db.GetDB().Model(&entity.Rule{}).Where("id = ?", ruleId).Update("datasource", datasource).Update("condition", condition).Update("action", action)
}