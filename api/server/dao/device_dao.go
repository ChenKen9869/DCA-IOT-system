package dao

import (
	"go-backend/api/common/db"
	"go-backend/api/server/entity"
)

func CreateFixedDevice(fixedDevice entity.FixedDevice) uint {
	db.GetDB().Create(&fixedDevice)
	return fixedDevice.ID
}
func CreatePortableDevice(portableDevice entity.PortableDevice) uint {
	db.GetDB().Create(&portableDevice)
	return portableDevice.ID
}

func CreateFixedDeviceType(fixedDeviceType entity.FixedDeviceType) {
	db.GetDB().Create(fixedDeviceType)
}

func CreatePortableDeviceType(portableDeviceType entity.PortableDeviceType) {
	db.GetDB().Create(&portableDeviceType)
}

func DeletePortableDevice(portableDeviceId uint) entity.PortableDevice {
	var portableDevice entity.PortableDevice
	db.GetDB().Where("id = ?", portableDeviceId).First(&portableDevice)
	db.GetDB().Delete(&portableDevice)
	return portableDevice
}

func DeletePortableDeviceType(portableDeviceTypeId string) entity.PortableDeviceType {
	var portableDeviceType entity.PortableDeviceType
	db.GetDB().Where("portable_device_type_id = ?", portableDeviceTypeId).First(&portableDeviceType)
	db.GetDB().Delete(&portableDeviceType)
	return portableDeviceType
}

func DeleteFixedDevice(fixedDeviceId uint) entity.FixedDevice {
	var fixedDevice entity.FixedDevice
	db.GetDB().Where("id = ?", fixedDeviceId).First(&fixedDevice)
	db.GetDB().Delete(&fixedDevice)
	return fixedDevice
}

func DeleteFixedDeviceType(fixedDeviceTypeId string) entity.FixedDeviceType {
	var fixedDeviceType entity.FixedDeviceType
	db.GetDB().Where("fixed_device_type_id = ?", fixedDeviceTypeId).First(&fixedDeviceType)
	db.GetDB().Delete(&fixedDeviceType)
	return fixedDeviceType
}

func ExistFixedDeviceType(fixedDeviceTypeId string) bool {
	var fixedType entity.FixedDeviceType
	db.GetDB().Table("fixed_device_types").Where("fixed_device_type_id = ?", fixedDeviceTypeId).First(&fixedType)
	return len(fixedType.FixedDeviceTypeID) != 0
}

func ExistPortableDeviceType(portableDeviceTypeId string) bool {
	var portableType entity.PortableDeviceType
	db.GetDB().Table("portable_device_types").Where("portable_device_type_id = ?", portableDeviceTypeId).First(&portableType)
	return len(portableType.PortableDeviceTypeID) != 0
}

func GetPortableDeviceListByBiology(biologyId uint) []entity.PortableDevice {
	var portableDeviceList []entity.PortableDevice
	db.GetDB().Where("biology_id = ?", biologyId).Find(&portableDeviceList)
	return portableDeviceList
}

func GetFixedDeviceListByFarmhouse(farmhouseId uint) []entity.FixedDevice {
	var fixedDeviceList []entity.FixedDevice
	db.GetDB().Where("farmhouse_id = ?", farmhouseId).Find(&fixedDeviceList)
	return fixedDeviceList
}

func GetFixedDeviceInfoById(fixedDeviceId uint) entity.FixedDevice {
	var fixedDevice entity.FixedDevice
	db.GetDB().Where("id = ?", fixedDeviceId).First(&fixedDevice)
	return fixedDevice
}

func GetPortableDeviceInfoById(portableDeviceId uint) entity.PortableDevice {
	var portableDevice entity.PortableDevice
	db.GetDB().Where("id = ?", portableDeviceId).First(&portableDevice)
	return portableDevice
}

func GetOwnFixedDeviceList(userId uint) []entity.FixedDevice {
	var fixedDeviceList []entity.FixedDevice
	db.GetDB().Table("fixed_devices").Where("owner = ?", userId).Find(&fixedDeviceList)
	return fixedDeviceList
}

func GetOwnPortableDeviceList(userId uint) []entity.PortableDevice {
	var portableDeviceList []entity.PortableDevice
	db.GetDB().Table("portable_devices").Where("owner = ?", userId).Find(&portableDeviceList)
	return portableDeviceList
}

func GetFixedDeviceInfoByMessagePayload(deviceId string, deviceType string) entity.FixedDevice {
	var deviceInfo entity.FixedDevice
	db.GetDB().Table("fixed_devices").Where("fixed_device_type_id = ?", deviceType).Where("device_id = ?", deviceId).First(&deviceInfo)
	return deviceInfo
}

func GetPortableDeviceInfoByMessagePayload(deviceId string, deviceType string) entity.PortableDevice {
	var deviceInfo entity.PortableDevice
	db.GetDB().Table("portable_devices").Where("portable_device_type_id = ?", deviceType).Where("device_id = ?", deviceId).First(&deviceInfo)
	return deviceInfo
}
