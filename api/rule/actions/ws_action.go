package actions

import (
	"fmt"
	"go-backend/api/sys/iot/monitor"
	"strconv"
	"strings"
)

// websocket 推送

// params: userId, msg
func sendWebsocketMSG(userId uint, msg string) {
	ch, exist := monitor.MonitorCenter[userId]
	if exist {
		ch.MessageChan <- msg

		fmt.Println("msg has send")
	} else {
		panic("warning: Msg was not sent. The target user is not connected to Monitoring Center!")
	}

}

var WsActionChannel chan (string)

func ExecWsAction(params string) {
	fmt.Println("doing websocket action... ")
	// params: userId,msg
	// params = strings.Replace(params, " ", "", -1)
	var paramList []string
	for i, c := range params {
		if string(c) == "," {
			if i == len(params)-1 {
				paramList = append(paramList, params)
				paramList = append(paramList, "")
			} else {
				paramList = append(paramList, params[:i])
				paramList = append(paramList, params[i+1:])
			}
		}
	}
	paramList[0] = strings.Replace(paramList[0], " ", "", -1)
	userId, err := strconv.Atoi(paramList[0])
	if err != nil {
		panic(err.Error())
	}
	sendWebsocketMSG(uint(userId), paramList[1])
}
