package rule

import (
	"fmt"
	"go-backend/api/rule/accepter"
	"go-backend/api/rule/actions"
	"go-backend/api/rule/ruleparser"
	"go-backend/api/rule/ruleparser/matcher"
	"go-backend/api/rule/scheduler"
	"time"

	"github.com/robfig/cron/v3"
)

func InitRule() {
	accepter.DatasourceManagement = make(map[accepter.DeviceIndex]accepter.KeyAttr)
	accepter.DeviceDBMap = make(map[string]accepter.DBTable)
	go accepter.StartExampleAccepter()
	accepter.DeviceDBMap[accepter.PortableDeviceType] = accepter.DBTable{
		TableName:  "portable_devices",
		ColumnName: "portable_device_type_id",
	}
	accepter.DeviceDBMap[accepter.FixedDeviceType] = accepter.DBTable{
		TableName:  "fixed_devices",
		ColumnName: "fixed_device_type_id",
	}

	ruleparser.MatcherMap = make(map[string]func([]ruleparser.Token, map[string]float64) bool)
	ruleparser.MatcherMap[ruleparser.Expression] = matcher.MatchExpressionCondition
	ruleparser.MatcherMap[ruleparser.PointSurfaceFunction] = matcher.MatchPointSurfaceFunctionCondition

	scheduler.RuleCron = cron.New()
	scheduler.RuleMap = make(map[uint]cron.EntryID)
	scheduler.RuleCron.Start()
	scheduler.ScheduledMap = make(map[uint]*time.Timer)

	actions.ActionChannels = make(map[string]chan string)

	actions.WsActionChannel = make(chan string)
	actions.MqttActionChannel = make(chan string)
	actions.ActionChannels[actions.WebsocketActionType] = actions.WsActionChannel
	actions.ActionChannels[actions.MqttActionType] = actions.MqttActionChannel
	go actions.StartActionExecutor()

	fmt.Println("[INITIAL SUCCESS] The rule module is initialized successfully!")
}
