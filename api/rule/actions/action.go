package actions

type ActionSignal struct {
	ActionType string
	ParamList  string
}

// 还是要用枚举，不容易出错
// const {
// 	MqttSignal
// }

var ActionChannels map[string](chan string)
