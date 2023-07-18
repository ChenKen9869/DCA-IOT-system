package rule

import "strings"

type Token struct {
	Type  string
	Value string
}

func isNum(string) bool

func isOpt(string) bool

func isSymbol(string) bool

func CheckSymbolTable(string) bool

func CreateNumToken(string) Token

func CreateSymbolToken(string) Token

func CreateOptToken(string) Token

// 需要解决： 下标遍历字符串的类型问题
func LexerCondition(condition string) []Token {
	var binOpt strings.Builder
	var tokenList []Token
	var temp string
	for index, value := range condition {
		if value == rune('=') {
			binOpt.WriteString("=")
			tokenList = append(tokenList, CreateOptToken(binOpt.String()))
			binOpt.Reset()
			continue
		} else {
			tokenList = append(tokenList, CreateOptToken(binOpt.String()))
			binOpt.Reset()
		}
		if index == (len(condition) - 1) {
			if isNum(temp) {
				tokenList = append(tokenList, CreateNumToken(temp))
			} else {
				if CheckSymbolTable(temp) {
					tokenList = append(tokenList, CreateSymbolToken(temp))
				} else {
					panic("input rule wrong! Symbol is not exists")
				}
			}
		}
		if value == rune('=') || value == rune('!') {
			if isNum(temp) {
				tokenList = append(tokenList, CreateNumToken(temp))
			} else {
				if CheckSymbolTable(temp) {
					tokenList = append(tokenList, CreateSymbolToken(temp))
				} else {
					panic("wrong")
				}
			}
			binOpt.WriteString(string(value))
			continue
		}
	}

	return tokenList
}
