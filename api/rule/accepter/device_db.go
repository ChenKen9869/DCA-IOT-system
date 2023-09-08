package accepter

type DeviceType = string
type DBTable struct {
	TableName  string
	ColumnName string
}

var DeviceDBMap map[DeviceType]DBTable

func getDeviceTypeInMysql(msgDeviceType string) string {
	if msgDeviceType == "collar" || msgDeviceType == "position-collar" {
		return PortableDeviceType
	} else {
		return FixedDeviceType
	}
}
