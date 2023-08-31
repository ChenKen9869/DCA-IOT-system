package ruleparser

import (
	"go-backend/api/rule/accepter"
	"go-backend/api/rule/actions"
)

func MatcherGenerator(symbolTable SymbolTable, conditionType string, tokenList []Token, actionList []Action) func() {
	return func() {
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
			for _, ac := range actionList {
				actions.ActionChannels[ac.ActionType] <- ac.ActionParams
			}
		}
	}
}
