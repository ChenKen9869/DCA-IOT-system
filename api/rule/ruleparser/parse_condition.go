package ruleparser

import (
	"strings"
)

var Expression string = "expression"

var StrTokenType string = "STR"
var NumTokenType string = "NUM"
var OptTokenType string = "OPT"
var ConstTokenType string = "CONST"
var ValTokenType string = "VAL"
var PairTokenType string = "PAIR"
var BoolTokenType string = "BOOL"

type Condition struct {
	ConditionType string
	StrTokenList  []string
}

/*
	Condition = (val_01 + 10 > val_02)&(val_02 < 30);
	Condition = type, val, val, const, const, ...
	函数式有逗号，表达式无逗号
*/
func ParseCondition(condition string) Condition {

	result := Condition{}
	if !strings.Contains(condition, ",") {
		// 表达式
		result.ConditionType = Expression
		result.StrTokenList = []string{strings.Replace(condition, " ", "", -1)}
	} else {
		condition = strings.Replace(condition, " ", "", -1)
		/*
			condition=
			type,val,val,const,const
		*/
		conditionList := strings.Split(condition, ",")
		/*
			conditionList=
			elem-01: type
			elem-02: val
			elem-03: val
			elem-04: const
			elem-05: const
		*/
		result.ConditionType = conditionList[0]
		result.StrTokenList = append(result.StrTokenList, conditionList[1:]...)
	}
	return result
}

func parseExpressionCondition(exp string, symbolTable SymbolTable) []Token {
	// 1. 解析符号
	infixTokenList := LexerCondition(exp, symbolTable)
	// 2. 中缀转后缀
	return TransformCondition(infixTokenList)
}

func parseFunctionCondition(paramsStr []string, symbolTable SymbolTable) []Token {
	var result []Token
	// 直接解析符号
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
