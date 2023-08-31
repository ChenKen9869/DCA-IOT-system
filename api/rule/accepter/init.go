package accepter

const (
	PortableDeviceType string = "Portable"
	FixedDeviceType    string = "Fixed"
)

func InitAccepter() {
	// 启动各接收器协程
	go startExampleAccepter()
	// 注册 device db map
	DeviceDBMap[PortableDeviceType] = "portable_devices"
	DeviceDBMap[FixedDeviceType] = "fixed_devices"
}
