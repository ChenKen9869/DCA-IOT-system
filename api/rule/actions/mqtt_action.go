package actions

import (
	"fmt"
	"strconv"
	"strings"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var MqttActionChannel chan (string)

func ExecMqttAction(params string) {
	address, port, username, password, topic, msg, portStr := parseMqttParams(params)

	opts := mqtt.NewClientOptions().AddBroker(fmt.Sprintf("tcp://%s:%d", address, port)).SetUsername(username).SetPassword(password)

	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	defer c.Disconnect(200)

	token := c.Publish(topic, 0, false, msg)
	token.Wait()
	fmt.Println("[Mqtt Executor] Message: { " + msg + " } has send to mqtt topic: { " + topic + " } of { " + address + ":" + portStr + " }! ")
}

func parseMqttParams(params string) (string, int, string, string, string, string, string) {
	paramList := strings.Split(params, ",")
	address := strings.Replace(paramList[0], " ", "", -1)
	portStr := strings.Replace(paramList[1], " ", "", -1)
	port, err := strconv.Atoi(portStr)
	if err != nil {
		panic(err.Error())
	}
	username := strings.Replace(paramList[2], " ", "", -1)
	password := strings.Replace(paramList[3], " ", "", -1)
	topic := strings.Replace(paramList[4], " ", "", -1)
	msg := paramList[5]

	return address, port, username, password, topic, msg, portStr
}
