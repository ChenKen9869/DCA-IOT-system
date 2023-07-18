package rule

func isNumToken(t Token) bool {
	return true
}

func isSymbolToken(t Token) bool {
	return true
}

func isOptToken(t Token) bool {
	return true
}

func getCurrData(t Token) Token

func isPairToken(t Token) bool {
	return true
}

func Convert(midExpression []Token) []Token {
	var postExpression []Token
	var stOpt Stack
	var stValue Stack
	for _, token := range midExpression {
		if isNumToken(token) {
			stValue.Push(token)
		} else if isSymbolToken(token) {
			stValue.Push(getCurrData(token))
		} else if isOptToken(token) {
			if isOptToken(Token(stOpt.Peak())) {

			}
		}
	}
	// 数值直接进栈
	// 左括号直接压入
	// 右括号就弹出直到匹配到左括号
	// 运算符需要判断优先级, 优先级高的在下面
	return postExpression
}

func Match(expression []Token, currDatas map[string]float64) bool {

	return true
}

// 优先级定义

// 字符数字运算
