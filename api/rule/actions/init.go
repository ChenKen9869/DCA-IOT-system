package actions

const (
	WebsocketActionType string = "WebSocket"
	MqttActionType      string = "Mqtt"
)

func InitAction() {
	// 将 action 的 channel 注册
	ActionChannels[WebsocketActionType] = WsActionChannel
	ActionChannels[MqttActionType] = MqttActionChannel
	// 启动监听器协程
	go startActionExecutor()
}

func startActionExecutor() {
	for {
		select {
		case params := <-WsActionChannel:
			go ExecWsAction(params)
		case params := <-MqttActionChannel:
			go ExecMqttAction(params)
		}
	}
}
