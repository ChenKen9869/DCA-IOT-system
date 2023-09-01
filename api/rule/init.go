package rule

import (
	"fmt"
	"go-backend/api/rule/accepter"
	"go-backend/api/rule/actions"
	"go-backend/api/rule/funcondition"
	"go-backend/api/rule/ruleparser"
	"go-backend/api/rule/ruleparser/matcher"
	"go-backend/api/rule/scheduler"

	"github.com/robfig/cron/v3"
)

func InitRule() {
	// accepter init
	accepter.DatasourceManagement = make(map[accepter.DeviceIndex]accepter.KeyAttr)
	accepter.DeviceDBMap = make(map[string]accepter.DBTable)
	// 启动各接收器协程
	go accepter.StartExampleAccepter()
	// 注册 device db map
	accepter.DeviceDBMap[accepter.PortableDeviceType] = accepter.DBTable{
		TableName:  "portable_devices",
		ColumnName: "portable_device_type_id",
	}
	accepter.DeviceDBMap[accepter.FixedDeviceType] = accepter.DBTable{
		TableName:  "fixed_devices",
		ColumnName: "fixed_device_type_id",
	}

	// rule parser matcher init
	ruleparser.MatcherMap = make(map[string]func([]ruleparser.Token, map[string]float64) bool)
	ruleparser.MatcherMap[ruleparser.Expression] = matcher.MatchExpressionCondition

	// scheduler init
	scheduler.RuleCron = cron.New()
	scheduler.RuleMap = make(map[uint]cron.EntryID)
	scheduler.RuleCron.Start()

	// function condition init
	funcondition.FunctionCondition = make(map[string]func(map[string]ruleparser.SymbolElem) bool)

	// action init
	actions.ActionChannels = make(map[string]chan string)

	// 将 action 的 channel 注册
	actions.ActionChannels[actions.WebsocketActionType] = actions.WsActionChannel
	actions.ActionChannels[actions.MqttActionType] = actions.MqttActionChannel
	// 启动监听器协程
	go actions.StartActionExecutor()

	fmt.Println("[INITIAL SUCCESS] The rule module is initialized successfully!")
}
