package actions

type ActionSignal struct {
	ActionType ActionType
	ParamList  string
}

type ActionType = string

var ActionChannels map[ActionType](chan string)
