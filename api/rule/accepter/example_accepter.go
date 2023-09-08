package accepter

import (
	"bufio"
	"fmt"
	"go-backend/api/server/dao"
	"net"
	"strconv"
	"strings"
)

func parseExampleMessage(msg string) (deviceId string, deviceType string, attribute string, value float64) {
	msg = strings.Replace(msg, " ", "", -1)
	msgList := strings.Split(msg, ",")
	deviceId = msgList[0]
	deviceType = msgList[1]
	attribute = msgList[2]
	v, err := strconv.ParseFloat(msgList[3], 64)
	if err != nil {
		panic(err.Error())
	}
	value = v
	return deviceId, deviceType, attribute, value
}

func processExampleMsg(conn net.Conn) {
	defer conn.Close()
	for {
		reader := bufio.NewReader(conn)
		fmt.Println("[Example Accepter] Waiting for message from client ...")
		var buf [128]byte
		n, err := reader.Read(buf[:])
		if err != nil {
			fmt.Println("[Example Accepter] Error Occur: Read from client failed, " + err.Error())
			fmt.Println("[Example Accepter] Client closed the connection")
			return
		}
		recvStr := string(buf[:n])
		fmt.Println("[Example Accepter] Accept Message: " + recvStr)
		conn.Write([]byte("[Message from server] Message accepted!"))

		deviceId, msgDeviceType, attribute, value := parseExampleMessage(recvStr)
		fmt.Println("[Example Accepter] Message has parsed! ")
		deviceType := getDeviceTypeInMysql(msgDeviceType)
		var id int
		if deviceType == PortableDeviceType {
			deviceInfo := dao.GetPortableDeviceInfoByMessagePayload(deviceId, msgDeviceType)
			id = int(deviceInfo.ID)
		} else if deviceType == FixedDeviceType {
			deviceInfo := dao.GetFixedDeviceInfoByMessagePayload(deviceId, msgDeviceType)
			id = int(deviceInfo.ID)
		}
		updateDatasourceManagement(id, deviceType, attribute, value)
	}
}

func StartExampleAccepter() {
	listen, err := net.Listen("tcp", "localhost:9869")
	if err != nil {
		panic(err.Error())
	}
	for {
		fmt.Println("[Example Accepter] Waiting for connect ... ")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("[Example Accepter] Connection establish failed, " + err.Error())
			continue
		}
		fmt.Println("[Example Accepter] Connection established successfully!")
		processExampleMsg(conn)
	}
}
