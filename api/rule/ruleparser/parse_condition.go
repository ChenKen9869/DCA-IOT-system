package ruleparser

import (
	"strings"
)

var Expression string = "expression"
var PointSurfaceFunction string = "PointSurface"

var StrTokenType string = "STR"
var NumTokenType string = "NUM"
var OptTokenType string = "OPT"
var ValTokenType string = "VAL"
var PairTokenType string = "PAIR"
var BoolTokenType string = "BOOL"

type Condition struct {
	ConditionType string
	StrTokenList  []string
}

/*
	Condition = (val_01 + 10 > val_02)&(val_02 < 30);
	Condition = type: val, val, const, const, ...
*/
func ParseCondition(condition string) Condition {

	result := Condition{}
	if !strings.Contains(condition, ":") {
		result.ConditionType = Expression
		result.StrTokenList = []string{strings.Replace(condition, " ", "", -1)}
	} else {
		condition = strings.Replace(condition, " ", "", -1)
		conditionList := strings.Split(condition, ":")
		result.ConditionType = conditionList[0]
		paramList := strings.Split(conditionList[1], ",")
		result.StrTokenList = append(result.StrTokenList, paramList...)
	}
	return result
}

func parseExpressionCondition(exp string, symbolTable SymbolTable) []Token {
	infixTokenList := LexerCondition(exp, symbolTable)
	return TransformCondition(infixTokenList)
}

func parseFunctionCondition(paramsStr []string, symbolTable SymbolTable) []Token {
	var result []Token
	for _, param := range paramsStr {
		flag := false
		for symbol := range symbolTable {
			if param == symbol {
				result = append(result, Token{
					TokenType:  ValTokenType,
					TokenValue: param,
				})
				flag = true
				break
			}
		}
		if !flag {
			result = append(result, Token{
				TokenType:  StrTokenType,
				TokenValue: param,
			})
		}
	}
	return result
}
