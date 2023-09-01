package ruleparser

import (
	"fmt"
	"go-backend/api/rule/accepter"
	"go-backend/api/rule/actions"
)

func MatcherGenerator(symbolTable SymbolTable, conditionType string, tokenList []Token, actionList []Action) func() {
	return func() {
		fmt.Println("rule match start, waiting for result ... ")
		// 构建内符号表
		currData := make(InnerTable)
		for symbol, symbolData := range symbolTable {
			accepter.DMLock.Lock()
			currData[symbol] = accepter.DatasourceManagement[accepter.DeviceIndex{
				Id:         symbolData.DeviceId,
				DeviceType: symbolData.DeviceType,
			}][symbolData.Attr].Value
			accepter.DMLock.Unlock()
		}
		// 查找并调用匹配算法
		if MatcherMap[conditionType](tokenList, currData) {
			fmt.Println("rule matched! ")
			for _, ac := range actionList {
				// 使用内符号表替换 actionparams

				actions.ActionChannels[ac.ActionType] <- ac.ActionParams
			}
			return
		}
		fmt.Println("rule not matched! ")
	}
}
