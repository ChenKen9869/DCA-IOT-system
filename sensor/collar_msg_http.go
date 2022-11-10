package sensor

type CollarTokenMessage struct {
	Code string
	Msg string
	Data CollarToken
}

type CollarToken struct {
	Token string
}

type CollarHttpMessage struct {

}

type CollarHttpCurrentMessage struct {

}