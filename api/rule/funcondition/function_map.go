package funcondition

import "go-backend/api/rule/ruleparser"

type FuncionConditionType = string
type FunctionConditionMatcher = func(ruleparser.SymbolTable) bool

var FunctionCondition map[FuncionConditionType]FunctionConditionMatcher = map[string]func(map[string]ruleparser.SymbolElem) bool{}
