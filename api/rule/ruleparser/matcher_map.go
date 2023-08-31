package ruleparser

// 内符号表
type InnerTable = map[string]float64

type ConditionType = string

var MatcherMap map[ConditionType]func([]Token, InnerTable) bool = make(map[string]func([]Token, map[string]float64) bool)
