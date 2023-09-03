package actions

import (
	"github.com/bytedance/gopkg/util/gopool"
)

const (
	WebsocketActionType string = "WebSocket"
	MqttActionType      string = "Mqtt"
)

func StartActionExecutor() {
	for {
		select {
		case params := <-WsActionChannel:
			gopool.Go(func() {
				ExecWsAction(params)
			})
		case params := <-MqttActionChannel:
			gopool.Go(func() {
				ExecMqttAction(params)
			})
		}
	}
}
