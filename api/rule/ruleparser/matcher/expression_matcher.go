package matcher

import (
	"fmt"
	"go-backend/api/rule/ruleparser"
	"go-backend/api/server/tools/util"
)

// 基于 stack 的匹配算法
func MatchExpressionCondition(tokenList []ruleparser.Token, innerTable ruleparser.InnerTable) bool {
	fmt.Println("entered match function")
	var st util.Stack
	for _, token := range tokenList {

		fmt.Println(token.TokenType, token.TokenValue, token.RealNum)

		if token.TokenType == ruleparser.NumTokenType {
			// 数值直接进栈
			st.Push(token)
		} else if token.TokenType == ruleparser.ValTokenType {
			// 变量用内表替换成数值后，入栈
			v := token.TokenValue
			for symbol, value := range innerTable {
				fmt.Println(symbol, value)
				if v == symbol {
					token.RealNum = value
					token.TokenType = ruleparser.NumTokenType
					st.Push(token)
				}
			}

			fmt.Println(token.RealNum)

		} else if token.TokenType == ruleparser.OptTokenType {
			// 运算符，则取出栈里的两个token，做运算，运算结果压入栈中
			optA := st.Pop().(ruleparser.Token)
			optB := st.Pop().(ruleparser.Token)
			st.Push(caculateToken(optB, optA, token))
		}
	}
	if st.Top().(ruleparser.Token).TokenType != ruleparser.BoolTokenType {
		panic("syntax error!")
	}

	fmt.Println(st.Top().(ruleparser.Token).RealNum)

	return util.IsFloat64Equal(st.Top().(ruleparser.Token).RealNum, float64(1))
}

func caculateToken(operandA, operandB, operator ruleparser.Token) ruleparser.Token {
	operatorStr := operator.TokenValue
	AType := operandA.TokenType
	BType := operandB.TokenType
	switch operatorStr {
	case "*":
		if AType != ruleparser.NumTokenType || BType != ruleparser.NumTokenType {
			panic("syntax error!")
		}
		return ruleparser.Token{
			TokenType:  ruleparser.NumTokenType,
			TokenValue: "",
			RealNum:    operandA.RealNum * operandB.RealNum,
		}
	case "/":
		if AType != ruleparser.NumTokenType || BType != ruleparser.NumTokenType {
			panic("syntax error!")
		}
		return ruleparser.Token{
			TokenType:  ruleparser.NumTokenType,
			TokenValue: "",
			RealNum:    operandA.RealNum / operandB.RealNum,
		}
	case "+":
		if AType != ruleparser.NumTokenType || BType != ruleparser.NumTokenType {
			panic("syntax error!")
		}
		return ruleparser.Token{
			TokenType:  ruleparser.NumTokenType,
			TokenValue: "",
			RealNum:    operandA.RealNum + operandB.RealNum,
		}
	case "-":
		if AType != ruleparser.NumTokenType || BType != ruleparser.NumTokenType {
			panic("syntax error!")
		}
		return ruleparser.Token{
			TokenType:  ruleparser.NumTokenType,
			TokenValue: "",
			RealNum:    operandA.RealNum - operandB.RealNum,
		}
	case ">":
		if AType != ruleparser.NumTokenType || BType != ruleparser.NumTokenType {
			panic("syntax error!")
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
			panic("syntax error!")
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
			panic("syntax error!")
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
			panic("syntax error!")
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
			panic("syntax error!")
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
