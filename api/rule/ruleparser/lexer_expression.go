package ruleparser

import (
	"strconv"
	"unicode"
)

func LexerCondition(exp string, symbolTable SymbolTable) []Token {
	var result []Token
	var temp *string
	skip := false
	for index, c := range exp {
		if skip {
			skip = false
			continue
		}
		if isOpt(string(c)) {
			generateToken(temp, &result, symbolTable)
			if string(c) == "!" || string(c) == "=" {
				if next := string(exp[index+1]); next == "=" {
					skip = true
					result = append(result, Token{
						TokenType:  OptTokenType,
						TokenValue: string(c) + next,
					})
				} else {
					panic("syntax error")
				}
			} else {
				if string(c) == "(" || string(c) == ")" {
					result = append(result, Token{
						TokenType:  PairTokenType,
						TokenValue: string(c),
					})
				} else {
					result = append(result, Token{
						TokenType:  OptTokenType,
						TokenValue: string(c),
					})
				}
			}
		} else if isNumOrCharacter(string(c)) {
			newTemp := *temp + string(c)
			temp = &newTemp
			if index == len(exp)-1 {
				generateToken(temp, &result, symbolTable)
			}
		}
	}

	// 返回中缀
	return result
}

// 是否是运算符？
func isOpt(str string) bool {
	return str == "*" || str == "/" ||
		str == "+" || str == "-" ||
		str == ">" || str == "<" ||
		str == "!" || str == "=" ||
		str == "&"
}

// 是否是0-9的数字，浮点，或者 _ ，或者大小写字母？
func isNumOrCharacter(str string) bool {
	_, err := strconv.Atoi(str)
	if err == nil {
		return true
	}
	return str == "." || str == "_" ||
		unicode.IsUpper(rune(str[0])) ||
		unicode.IsLower(rune(str[0]))
}

// 根据符号表生成 token，并加入到 tokenList 中
func generateToken(value *string, tokenList *[]Token, symbolTable SymbolTable) {
	intValue, errInt := strconv.Atoi(*value)
	if errInt == nil {
		*tokenList = append(*tokenList, Token{
			TokenType:  NumTokenType,
			TokenValue: *value,
			RealNum:    float64(intValue),
		})
	} else if floatValue, errFloat := strconv.ParseFloat(*value, 64); errFloat == nil {
		*tokenList = append(*tokenList, Token{
			TokenType:  NumTokenType,
			TokenValue: *value,
			RealNum:    floatValue,
		})
	} else {
		found := false
		for symbol := range symbolTable {
			if symbol == *value {
				*tokenList = append(*tokenList, Token{
					TokenType:  ValTokenType,
					TokenValue: *value,
				})
				found = true
			}
		}
		if !found {
			panic("syntax error: symbol not exist!")
		}
	}
	*value = ""
}
