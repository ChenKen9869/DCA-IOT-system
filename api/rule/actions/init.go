package actions

import (
	"fmt"

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
			fmt.Println("[Action Executor] Params arrived at websocket executor!")
			gopool.Go(func() {
				ExecWsAction(params)
			})
			// go ExecWsAction(params)
		case params := <-MqttActionChannel:
			fmt.Println("[Action Executor] Params arrived at mqtt executor!")
			gopool.Go(func() {
				ExecMqttAction(params)
			})
			// go ExecMqttAction(params)
		}
	}
}
