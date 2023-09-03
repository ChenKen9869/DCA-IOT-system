package ruleparser

import "fmt"

type Token struct {
	TokenType  string
	TokenValue string
	RealNum    float64
}

type SymbolElem struct {
	DeviceId   int
	DeviceType string
	Attr       string
}

type Name = string
type SymbolTable = map[Name]SymbolElem

func ParseRule(ruleIdStr string, datasource string, condition string, action string) func() {
	datasourceList := ParseDatasource(datasource)
	conditionWithType := ParseCondition(condition)
	actionList := ParseAction(action)
	symbolTable := make(SymbolTable)

	for _, ds := range datasourceList {
		symbolTable[ds.Name] = SymbolElem{
			DeviceId:   ds.DeviceId,
			DeviceType: ds.DeviceType,
			Attr:       ds.Attribute,
		}
	}

	var tokenList []Token
	if conditionWithType.ConditionType == Expression {
		tokenList = parseExpressionCondition(conditionWithType.StrTokenList[0], symbolTable)
	} else {
		tokenList = parseFunctionCondition(conditionWithType.StrTokenList, symbolTable)
	}

	fmt.Println("[Rule Parser: " + ruleIdStr + "] Rule has parsed. Condition type: " + conditionWithType.ConditionType)
	return MatcherGenerator(ruleIdStr, symbolTable, conditionWithType.ConditionType, tokenList, actionList)
}
