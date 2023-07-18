package rule

import (
	"go-backend/api/rule/actions"
	"go-backend/api/server/tools/util"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/robfig/cron/v3"
)

// 订阅 mqtt 指定 topic : 用 deviceId 查表，查出 topic
func subscribeDatasource(Datasource) *(chan float64)

type Datasource struct {
	DeviceId  uint
	Attribute string
}

type SymbolTable map[string]*(chan float64)

var datasourceManager map[Datasource]*(chan float64)

func getDatasource(datasource Datasource) *(chan float64) {
	if value, exist := datasourceManager[datasource]; exist {
		return value
	}
	dataChan := subscribeDatasource(datasource)
	datasourceManager[datasource] = dataChan
	// 数据源管理器中存放消息队列
	// 但是这里返回管道引用，因为getDatasource将创建一个管道，管道订阅消息队列，然后将管道引用返回
	// 如果直接把消息队列返回可以吗？不可以，如果直接用消息队列的引用
	// 会导致不同规则之间相互影响
	return dataChan
}

func parseTime(start string) time.Time {
	return util.ParseDate(start)
}

func parseInterval(interval string) string {
	// 判断语法是否合法
	return interval
}

// 解析数据源字段，返回符号表
func parseDatasource(datasource string) SymbolTable {
	symbolTable := make(SymbolTable)
	// 创建内表

	// 解析出datasource
	// 去除所有空格
	sources := strings.Replace(datasource, " ", "", -1)
	// 先按 & 分割
	sourceCollection := strings.Split(sources, "&")
	// for 遍历所有datasource字符串
	for _, source := range sourceCollection {
		sourceList := strings.Split(source, ",")
		name := sourceList[0]
		deviceId, err := strconv.Atoi(sourceList[1])
		if err != nil {
			panic(err.Error())
		}
		attribute := sourceList[2]
		symbolTable[name] = getDatasource(Datasource{
			DeviceId:  uint(deviceId),
			Attribute: attribute,
		})
	}

	go func() {
		// 无限循环
		// 让外表的数据传输到内表中
	}()

	// 返回内表引用
	return symbolTable
}

// 解析 condition ，返回 后序 token
func parseCondition(condition string) []Token {
	// 去除空格
	// 调用 lexer
	tokenList := LexerCondition(condition)
	return tokenList
}

func parseAction(actionStr string) []actions.ActionSignal {
	var actionList []actions.ActionSignal
	// 以 & 作为分隔符 split
	actionCollection := strings.Split(actionStr, "&")
	// 解析 action
	for _, action := range actionCollection {
		// split 到第一个逗号，然后 split 出 acntionType 和 paramList
		typeAndParam := strings.SplitN(action, ",", 1)
		actionType := strings.Replace(typeAndParam[0], " ", "", -1)
		paramList := typeAndParam[1]
		actionList = append(actionList, actions.ActionSignal{
			ActionType: actionType,
			ParamList:  paramList,
		})
	}
	return actionList
}

func createMatcher(symbolTable SymbolTable, tokenList []Token, actionList []actions.ActionSignal) func() {
	return func() {
		// 应该保证尽量同时，最近的消息比较合理,如果是多个datasource的话
		currDatas := make(map[string]float64)
		var wg sync.WaitGroup
		wg.Add(len(symbolTable))
		// 获取所有最近数据

		// 深拷贝 内表
		for name, channel := range symbolTable {
			go func(str string, ch chan float64) {
				currDatas[str] = <-ch
				wg.Done()
			}(name, *channel)
		}
		wg.Wait()
		if Match(tokenList, currDatas) {
			// 发送出去
			for _, action := range actionList {
				actions.ActionChannels[action.ActionType] <- action.ParamList
			}
		}
	}
}

// 解析规则，传入规则所在文件
// 规则文件传到后端后，先统一保存在某个文件目录下，然后再解析
func ParseRule(filePathOfRule string) {
	start := "start"
	end := "end"
	interval := "interval"
	datasource := "datasource"
	condition := "condition"
	action := "action"

	symbolTable := parseDatasource(datasource)

	matcher := createMatcher(symbolTable, parseCondition(condition), parseAction(action))

	startTime := time.NewTimer(time.Until(parseTime(start)))

	endTime := time.NewTimer(time.Until(parseTime(end)))

	c := cron.New()
	c.AddFunc(parseInterval(interval), matcher)

	<-startTime.C
	c.Start()

	<-endTime.C
	c.Stop()

	// 同时清理所有资源：清理内表更新协程，清理内外表，清理token list，清理 cron，检查数据源管理器，清理数据源管道
}
