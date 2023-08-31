package actions

import (
	"go-backend/api/sys/iot/monitor"
	"strconv"
	"strings"
)

// websocket 推送

// params: userId, msg
func pushWebsocketMSG(userId uint, msg string) {
	ch, exist := monitor.MonitorCentor[userId]
	if exist {
		ch.MessageChan <- msg
	} else {
		panic("warning: Msg was not sent. The target user is not connected to Monitoring Center!")
	}

}

var WsActionChannel chan (string) = make(chan string)

func ExecWsAction(params string) {
	// params: userId, msg
	params = strings.Replace(params, " ", "", -1)
	paramList := strings.Split(params, ";")
	userId, err := strconv.Atoi(paramList[0])
	if err != nil {
		panic(err.Error())
	}
	pushWebsocketMSG(uint(userId), paramList[1])
}
