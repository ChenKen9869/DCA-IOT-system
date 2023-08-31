package ruleparser

type Token struct {
	TokenType  	string
	TokenValue 	string
	RealNum		float64
}

type SymbolElem struct {
	DeviceId   int
	DeviceType string
	Attr       string
}

type Name = string
type SymbolTable = map[Name]SymbolElem

// // 解析规则，传入规则所在文件
// // 规则文件传到后端后，先统一保存在某个文件目录下，然后再解析
func ParseRule(datasource string, condition string, action string) func() {
	datasourceList := ParseDatasource(datasource)
	conditionWithType := ParseCondition(condition)
	actionList := ParseAction(action)
	// 1. 建立外符号表
	symbolTable := make(SymbolTable)

	for _, ds := range datasourceList {
		symbolTable[ds.Name] = SymbolElem{
			DeviceId:   ds.DeviceId,
			DeviceType: ds.DeviceType,
			Attr:       ds.Attribute,
		}
	}

	// 2. 获取 TokenList
	var tokenList []Token
	if conditionWithType.ConditionType == Expression {
		tokenList = parseExpressionCondition(conditionWithType.StrTokenList[0], symbolTable)
	} else {
		// 函数式
		tokenList = parseFunctionCondition(conditionWithType.StrTokenList, symbolTable)
	}

	// 3. 生成并返回模式匹配函数，将外符号表，类型标识符，TokenList，actionList一起传进去
	return MatcherGenerator(symbolTable, conditionWithType.ConditionType, tokenList, actionList)
}

