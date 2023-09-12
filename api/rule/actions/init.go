package actions

import (
	"github.com/bytedance/gopkg/util/gopool"
)

const (
	WebsocketActionType string = "WebSocket"
	MqttActionType      string = "Mqtt"
)

func startActionExecutor() {
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

func InitAction() {
	ActionChannels = make(map[string]chan string)

	WsActionChannel = make(chan string)
	MqttActionChannel = make(chan string)
	ActionChannels[WebsocketActionType] = WsActionChannel
	ActionChannels[MqttActionType] = MqttActionChannel
	go startActionExecutor()
}
