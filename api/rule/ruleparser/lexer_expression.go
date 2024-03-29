package ruleparser

import (
	"strconv"
	"unicode"
)

func LexerCondition(exp string, symbolTable SymbolTable) []Token {
	var result []Token
	tempString := ""
	var temp *string
	temp = &tempString
	skip := false
	for index, c := range exp {
		if skip {
			skip = false
			continue
		}
		if isOpt(string(c)) {
			if *temp != "" {
				generateToken(temp, &result, symbolTable)
			}
			if string(c) == "!" || string(c) == "=" {
				if next := string(exp[index+1]); next == "=" {
					skip = true
					result = append(result, Token{
						TokenType:  OptTokenType,
						TokenValue: string(c) + next,
					})
				} else {
					panic("Syntax Error")
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
	return result
}

func isOpt(str string) bool {
	return str == "*" || str == "/" ||
		str == "+" || str == "-" ||
		str == ">" || str == "<" ||
		str == "!" || str == "=" ||
		str == "&" || str == "(" ||
		str == ")"
}

func isNumOrCharacter(str string) bool {
	_, err := strconv.Atoi(str)
	if err == nil {
		return true
	}
	return str == "." || str == "_" ||
		unicode.IsUpper(rune(str[0])) ||
		unicode.IsLower(rune(str[0]))
}

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
			panic("Syntax Error: Symbol " + *value + " not exist!")
		}
	}
	*value = ""
}
