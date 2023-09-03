package ruleparser

import (
	"fmt"
	"go-backend/api/rule/accepter"
	"go-backend/api/rule/actions"
	"strconv"
)

func MatcherGenerator(ruleIdStr string, symbolTable SymbolTable, conditionType string, tokenList []Token, actionList []Action) func() {
	return func() {
		// 构建内符号表
		currData := make(InnerTable)
		for symbol, symbolData := range symbolTable {
			accepter.DMLock.Lock()
			v, exist := accepter.DatasourceManagement[accepter.DeviceIndex{
				Id:         symbolData.DeviceId,
				DeviceType: symbolData.DeviceType,
			}]
			if !exist {
				panic("[Rule Matcher: " + ruleIdStr + "] Rule Device information does not exist in datasource management!")
			}
			data, e := v[symbolData.Attr]
			if !e {
				panic("[Rule Matcher: " + ruleIdStr + "] Attribute information does not exist in device info of datasource!")
			}
			currData[symbol] = data.Value
			accepter.DMLock.Unlock()
		}
		// 查找并调用匹配算法
		if MatcherMap[conditionType](tokenList, currData) {
			fmt.Println("[Rule Matcher: " + ruleIdStr + "] Rule matched! ")
			for _, ac := range actionList {
				// 使用内符号表替换 actionparams
				params := ac.ActionParams
				params = replaceSymbolInParams(params, currData)
				actionChannel, exist := actions.ActionChannels[ac.ActionType]
				if exist {
					actionChannel <- params
				} else {
					panic("[Rule Matcher: " + ruleIdStr + "]Syntax error: action type " + ac.ActionType + "does not exist! ")
				}
			}
			return
		}
		fmt.Println("[Rule Matcher: " + ruleIdStr + "] Rule not matched! ")
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
