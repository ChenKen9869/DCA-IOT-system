package accepter

import (
	"bufio"
	"fmt"
	"go-backend/api/common/common"
	"go-backend/api/server/entity"
	"net"
	"strconv"
	"strings"
)

func parseExampleMessage(msg string) (deviceId string, deviceType string, attribute string, value float64) {
	// 0000001, collar, temperature, 25.6
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

		// deviceId, msgDeviceType, attribute, value := messageArrive()
		deviceId, msgDeviceType, attribute, value := parseExampleMessage(recvStr)
		fmt.Println("[Example Accepter] Message has parsed! ")
		deviceType := getDeviceTypeInMysql(msgDeviceType)
		// 在 mysql 中查找对应设备的 主键id
		var id int
		if deviceType == PortableDeviceType {
			var deviceInfo entity.PortableDevice
			common.GetDB().Table(DeviceDBMap[deviceType].TableName).Where(DeviceDBMap[deviceType].ColumnName+" = ?", msgDeviceType).Where("device_id = ?", deviceId).First(&deviceInfo)
			id = int(deviceInfo.ID)
		} else {
			var deviceInfo entity.FixedDevice
			common.GetDB().Table(DeviceDBMap[deviceType].TableName).Where(DeviceDBMap[deviceType].ColumnName+" = ?", msgDeviceType).Where("device_id = ?", deviceId).First(&deviceInfo)
			id = int(deviceInfo.ID)
		}
		updateDatasourceManagement(id, deviceType, attribute, value)
		// time.Sleep(30 * time.Second)
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

// 查找并更新数据源管理器的数据
func updateDatasourceManagement(id int, deviceType string, attr string, value float64) {

	index := DeviceIndex{
		Id:         id,
		DeviceType: deviceType,
	}

	DMLock.Lock()
	v1, exist1 := DatasourceManagement[index]
	if exist1 {
		v, exist := v1[attr]
		if exist {
			v.Value = value
			DatasourceManagement[index][attr] = v
		}
		fmt.Println("[Example Accepter] Datasource management update is complete!")
	} else {
		fmt.Println("[Example Accepter] Datasource management was not updated!")
	}
	DMLock.Unlock()
}

func getDeviceTypeInMysql(msgDeviceType string) string {
	if msgDeviceType == "collar" || msgDeviceType == "position-collar" {
		return PortableDeviceType
	} else {
		return FixedDeviceType
	}
}

// 模拟数据到达
// func messageArrive() (deviceId string, deviceType string, attribute string, value float64) {
// 	return "0000001", "collar", "temperature", float64(25.6)
// }
