package ruleparser

type InnerTable = map[string]float64

type ConditionType = string

var MatcherMap map[ConditionType]func([]Token, InnerTable) bool
