package matcher

import "go-backend/api/rule/ruleparser"

func InitMatcher() {
	ruleparser.MatcherMap[ruleparser.Expression] = MatchExpressionCondition
}