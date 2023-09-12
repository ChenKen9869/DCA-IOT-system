package accepter

import "github.com/bits-and-blooms/bloom"

const (
	PortableDeviceType string = "Portable"
	FixedDeviceType    string = "Fixed"
)

func InitAccepter() {
	BloomFilter = bloom.NewWithEstimates(10000000, 0.01)
	DatasourceManagement = make(map[DeviceIndex]KeyAttr)
	DeviceDBMap = make(map[string]DBTable)
	DeviceDBMap[PortableDeviceType] = DBTable{
		TableName:  "portable_devices",
		ColumnName: "portable_device_type_id",
	}
	DeviceDBMap[FixedDeviceType] = DBTable{
		TableName:  "fixed_devices",
		ColumnName: "fixed_device_type_id",
	}

	go StartExampleAccepter()
}