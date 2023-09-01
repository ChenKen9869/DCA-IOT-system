package actions

const (
	WebsocketActionType string = "WebSocket"
	MqttActionType      string = "Mqtt"
)

func StartActionExecutor() {
	for {
		select {
		case params := <-WsActionChannel:
			go ExecWsAction(params)
		case params := <-MqttActionChannel:
			go ExecMqttAction(params)
		}
	}
}
