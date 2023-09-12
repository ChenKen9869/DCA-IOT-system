package ruleparser

func InitRuleparser() {
	MatcherMap = make(map[string]func([]Token, map[string]float64) bool)
}
