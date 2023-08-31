package accepter

import (
	"go-backend/api/common/common"
	"go-backend/api/server/entity"
	"time"
)

// (可选)存入 mongoDB 数据库

func startExampleAccepter() {
	for {
		deviceId, msgDeviceType, attribute, value := messageArrive()
		deviceType := getDeviceTypeInMysql(msgDeviceType)
		// 在 mysql 中查找对应设备的 主键id
		var id int
		if deviceType == PortableDeviceType {
			var deviceInfo entity.PortableDevice
			common.GetDB().Table(DeviceDBMap[deviceType]).Where("device_id = ?", deviceId).Where("device_type = ?", msgDeviceType).First(&deviceInfo)
			id = int(deviceInfo.ID)
		} else {
			var deviceInfo entity.FixedDevice
			common.GetDB().Table(DeviceDBMap[deviceType]).Where("device_id = ?", deviceId).Where("device_type = ?", msgDeviceType).First(&deviceInfo)
			id = int(deviceInfo.ID)
		}
		updateDatasourceManagement(id, deviceType, attribute, value)

		time.Sleep(5 * time.Minute)
	}
}

// 查找并更新数据源管理器的数据
func updateDatasourceManagement(id int, deviceType string, attr string, value float64) {
	index := DeviceIndex{
		Id:         id,
		DeviceType: deviceType,
	}

	DMLock.Lock()
	v := DatasourceManagement[index][attr]
	v.Value = value
	DatasourceManagement[index][attr] = v
	DMLock.Unlock()
}

// 模拟数据到达
func messageArrive() (deviceId string, deviceType string, attribute string, value float64) {
	return "0000001", "collar", "temperature", float64(25.6)
}

func getDeviceTypeInMysql(msgDeviceType string) string {
	if msgDeviceType == "collar" || msgDeviceType == "position-collar" {
		return PortableDeviceType
	} else {
		return FixedDeviceType
	}
}
