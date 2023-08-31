package entity

import "github.com/jinzhu/gorm"

type Rule struct {
	gorm.Model
	Datasource string
	Condition  string
	Action     string
	Owner      uint
	ParentId   uint
	Stat       string
}

type RuleStat = string

const (
	RuleInactive  RuleStat = "Inactive"
	RuleActive    RuleStat = "Active"
	RuleScheduled RuleStat = "Scheduled"
)
