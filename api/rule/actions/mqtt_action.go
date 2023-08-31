package actions

import (
	"fmt"
	"strconv"
	"strings"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var MqttActionChannel chan (string) = make(chan string)

// 发布一条消息到 mqtt
func ExecMqttAction(params string) {
	// params: address, port, username, password, topic, msg
	address, port, username, password, topic, msg := parseMqttParams(params)

	opts := mqtt.NewClientOptions().AddBroker(fmt.Sprintf("tcp://%s:%d", address, port)).SetUsername(username).SetPassword(password)

	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	// 断开连接
	defer c.Disconnect(200)

	// 发布消息
	token := c.Publish(topic, 0, false, msg)
	token.Wait()
}

func parseMqttParams(params string) (address string, port int, username string, password string, topic string, msg string) {
	params = strings.Replace(params, " ", "", -1)
	paramList := strings.Split(params, ";")
	port, err := strconv.Atoi(paramList[1])
	if err != nil {
		panic(err.Error())
	}
	return paramList[0], port, paramList[2], paramList[3], paramList[4], paramList[5]
}
