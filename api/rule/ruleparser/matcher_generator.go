package ruleparser

import (
	"fmt"
	"go-backend/api/rule/accepter"
	"go-backend/api/rule/actions"
	"strconv"
)

func MatcherGenerator(symbolTable SymbolTable, conditionType string, tokenList []Token, actionList []Action) func() {
	return func() {
		fmt.Println("rule match start, waiting for result ... ")
		// 构建内符号表
		currData := make(InnerTable)
		for symbol, symbolData := range symbolTable {
			fmt.Println(symbol, symbolData.Attr, symbolData.DeviceType, symbolData.DeviceId)
			accepter.DMLock.Lock()
			v, exist := accepter.DatasourceManagement[accepter.DeviceIndex{
				Id:         symbolData.DeviceId,
				DeviceType: symbolData.DeviceType,
			}]
			if exist {
				fmt.Println("device information exist in datasource management!")
			}
			data, e := v[symbolData.Attr]
			if e {
				fmt.Println("attribute exist in device info!")
			}
			currData[symbol] = data.Value
			// currData[symbol] = accepter.DatasourceManagement[accepter.DeviceIndex{
			// 	Id:         symbolData.DeviceId,
			// 	DeviceType: symbolData.DeviceType,
			// }][symbolData.Attr].Value
			fmt.Println(currData[symbol])
			accepter.DMLock.Unlock()
		}
		// 查找并调用匹配算法
		if MatcherMap[conditionType](tokenList, currData) {
			fmt.Println("rule matched! ")
			for _, ac := range actionList {
				// 使用内符号表替换 actionparams
				fmt.Println(ac.ActionParams)
				params := ac.ActionParams
				params = replaceSymbolInParams(params, currData)
				actionChannel, exist := actions.ActionChannels[ac.ActionType]
				if exist {
					actionChannel <- params
				} else {
					panic("syntax error: action type " + ac.ActionType + "does not exist! ")
				}
			}
			return
		}
		fmt.Println("rule not matched! ")
	}
}

// 替换 $ 符号以及其后的匹配符号，默认匹配最长
func replaceParamsRecursive(params string, innerTable InnerTable) string {
	for index, c := range params {
		if string(c) == "$" {
			valueString := ""
			maxLenth := 0
			for symbol, value := range innerTable {
				l := len(symbol)
				if l+index <= len(params) {
					if params[index+1:index+len(symbol)+1] == symbol {
						// matched
						// 判断是否比最大已匹配符号长，如果是则替换 valueString，更新 maxLenth
						if l > maxLenth {
							// 保留 6 位小数
							valueString = strconv.FormatFloat(value, 'f', 6, 64)
							maxLenth = l
						}
					}
				}
			}
			if maxLenth != 0 {
				// matched, should replace
				newParams := ""
				if index+maxLenth < len(params) {
					newParams = params[:index] + valueString + params[index+maxLenth+1:]
				} else {
					newParams = params[:index] + valueString
				}
				params = newParams
				return replaceParamsRecursive(params, innerTable)
			}
		}
	}
	return params
}

func replaceSymbolInParams(params string, innerTable InnerTable) (strOfSymbolCurrentData string) {
	params = replaceParamsRecursive(params, innerTable)
	return params
}
