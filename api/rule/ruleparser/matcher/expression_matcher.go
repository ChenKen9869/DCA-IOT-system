package matcher

import (
	"go-backend/api/rule/ruleparser"
	"go-backend/api/server/tools/util"
)

func MatchExpressionCondition(tokenList []ruleparser.Token, innerTable ruleparser.InnerTable) bool {
	var st util.Stack
	for _, token := range tokenList {
		if token.TokenType == ruleparser.NumTokenType {
			st.Push(token)
		} else if token.TokenType == ruleparser.ValTokenType {
			v := token.TokenValue
			for symbol, value := range innerTable {
				if v == symbol {
					token.RealNum = value
					token.TokenType = ruleparser.NumTokenType
					st.Push(token)
				}
			}
		} else if token.TokenType == ruleparser.OptTokenType {
			optA := st.Pop().(ruleparser.Token)
			optB := st.Pop().(ruleparser.Token)
			st.Push(caculateToken(optB, optA, token))
		}
	}
	if st.Top().(ruleparser.Token).TokenType != ruleparser.BoolTokenType {
		panic("syntax error!")
	}
	return util.IsFloat64Equal(st.Top().(ruleparser.Token).RealNum, float64(1))
}

func caculateToken(operandA, operandB, operator ruleparser.Token) ruleparser.Token {
	operatorStr := operator.TokenValue
	AType := operandA.TokenType
	BType := operandB.TokenType
	switch operatorStr {
	case "*":
		if AType != ruleparser.NumTokenType || BType != ruleparser.NumTokenType {
			panic("Syntax Error!")
		}
		return ruleparser.Token{
			TokenType:  ruleparser.NumTokenType,
			TokenValue: "",
			RealNum:    operandA.RealNum * operandB.RealNum,
		}
	case "/":
		if AType != ruleparser.NumTokenType || BType != ruleparser.NumTokenType {
			panic("Syntax Error!")
		}
		if util.IsFloat64Equal(operandB.RealNum, float64(0)) {
			panic("Error: The divisor is 0!")
		}
		return ruleparser.Token{
			TokenType:  ruleparser.NumTokenType,
			TokenValue: "",
			RealNum:    operandA.RealNum / operandB.RealNum,
		}
	case "+":
		if AType != ruleparser.NumTokenType || BType != ruleparser.NumTokenType {
			panic("Syntax Error!")
		}
		return ruleparser.Token{
			TokenType:  ruleparser.NumTokenType,
			TokenValue: "",
			RealNum:    operandA.RealNum + operandB.RealNum,
		}
	case "-":
		if AType != ruleparser.NumTokenType || BType != ruleparser.NumTokenType {
			panic("Syntax Error!")
		}
		return ruleparser.Token{
			TokenType:  ruleparser.NumTokenType,
			TokenValue: "",
			RealNum:    operandA.RealNum - operandB.RealNum,
		}
	case ">":
		if AType != ruleparser.NumTokenType || BType != ruleparser.NumTokenType {
			panic("Syntax Error!")
		}
		var boolValue int
		if operandA.RealNum > operandB.RealNum {
			boolValue = 1
		} else {
			boolValue = 0
		}
		return ruleparser.Token{
			TokenType:  ruleparser.BoolTokenType,
			TokenValue: "",
			RealNum:    float64(boolValue),
		}
	case "<":
		if AType != ruleparser.NumTokenType || BType != ruleparser.NumTokenType {
			panic("Syntax Error!")
		}
		var boolValue int
		if operandA.RealNum < operandB.RealNum {
			boolValue = 1
		} else {
			boolValue = 0
		}
		return ruleparser.Token{
			TokenType:  ruleparser.BoolTokenType,
			TokenValue: "",
			RealNum:    float64(boolValue),
		}
	case "!=":
		if AType != ruleparser.NumTokenType || BType != ruleparser.NumTokenType {
			panic("Syntax Error!")
		}
		var boolValue int
		if !util.IsFloat64Equal(operandA.RealNum, operandB.RealNum) {
			boolValue = 1
		} else {
			boolValue = 0
		}
		return ruleparser.Token{
			TokenType:  ruleparser.BoolTokenType,
			TokenValue: "",
			RealNum:    float64(boolValue),
		}
	case "==":
		if AType != ruleparser.NumTokenType || BType != ruleparser.NumTokenType {
			panic("Syntax Error!")
		}
		var boolValue int
		if util.IsFloat64Equal(operandA.RealNum, operandB.RealNum) {
			boolValue = 1
		} else {
			boolValue = 0
		}
		return ruleparser.Token{
			TokenType:  ruleparser.BoolTokenType,
			TokenValue: "",
			RealNum:    float64(boolValue),
		}
	case "&":
		if AType != ruleparser.BoolTokenType || BType != ruleparser.BoolTokenType {
			panic("Syntax Error!")
		}
		var boolValue int
		if util.IsFloat64Equal(operandA.RealNum, float64(1)) && util.IsFloat64Equal(operandB.RealNum, float64(1)) {
			boolValue = 1
		} else {
			boolValue = 0
		}
		return ruleparser.Token{
			TokenType:  ruleparser.BoolTokenType,
			TokenValue: "",
			RealNum:    float64(boolValue),
		}
	}
	return ruleparser.Token{}
}
