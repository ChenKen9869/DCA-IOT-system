package dao

import (
	"go-backend/api/common/common"
	"go-backend/api/server/entity"
)

func CreateFixedDevice(fixedDevice entity.FixedDevice) uint {
	common.GetDB().Create(&fixedDevice)
	return fixedDevice.ID
}
func CreatePortableDevice(portableDevice entity.PortableDevice) uint {
	common.GetDB().Create(&portableDevice)
	return portableDevice.ID
}

func CreateFixedDeviceType(fixedDeviceType entity.FixedDeviceType) {
	common.GetDB().Create(fixedDeviceType)
}

func CreatePortableDeviceType(portableDeviceType entity.PortableDeviceType) {
	common.GetDB().Create(&portableDeviceType)
}

func DeletePortableDevice(portableDeviceId uint) entity.PortableDevice {
	var portableDevice entity.PortableDevice
	common.GetDB().Where("id = ?", portableDeviceId).First(&portableDevice)
	common.GetDB().Delete(&portableDevice)
	return portableDevice
}

func DeletePortableDeviceType(portableDeviceTypeId string) entity.PortableDeviceType {
	var portableDeviceType entity.PortableDeviceType
	common.GetDB().Where("portable_device_type_id = ?", portableDeviceTypeId).First(&portableDeviceType)
	common.GetDB().Delete(&portableDeviceType)
	return portableDeviceType
}

func DeleteFixedDevice(fixedDeviceId uint) entity.FixedDevice {
	var fixedDevice entity.FixedDevice
	common.GetDB().Where("id = ?", fixedDeviceId).First(&fixedDevice)
	common.GetDB().Delete(&fixedDevice)
	return fixedDevice
}

func DeleteFixedDeviceType(fixedDeviceTypeId string) entity.FixedDeviceType {
	var fixedDeviceType entity.FixedDeviceType
	common.GetDB().Where("fixed_device_type_id = ?", fixedDeviceTypeId).First(&fixedDeviceType)
	common.GetDB().Delete(&fixedDeviceType)
	return fixedDeviceType
}

func ExistFixedDeviceType(fixedDeviceTypeId string) bool {
	var fixedType entity.FixedDeviceType
	common.GetDB().Table("fixed_device_types").Where("fixed_device_type_id = ?", fixedDeviceTypeId).First(&fixedType)
	return len(fixedType.FixedDeviceTypeID) != 0
}

func ExistPortableDeviceType(portableDeviceTypeId string) bool {
	var portableType entity.PortableDeviceType
	common.GetDB().Table("portable_device_types").Where("portable_device_type_id = ?", portableDeviceTypeId).First(&portableType)
	return len(portableType.PortableDeviceTypeID) != 0
}

func GetPortableDeviceListByBiology(biologyId uint) []entity.PortableDevice {
	var portableDeviceList []entity.PortableDevice
	common.GetDB().Where("biology_id = ?", biologyId).Find(&portableDeviceList)
	return portableDeviceList
}

func GetFixedDeviceListByFarmhouse(farmhouseId uint) []entity.FixedDevice {
	var fixedDeviceList []entity.FixedDevice
	common.GetDB().Where("farmhouse_id = ?", farmhouseId).Find(&fixedDeviceList)
	return fixedDeviceList
}

func GetFixedDeviceInfoById(fixedDeviceId uint) entity.FixedDevice {
	var fixedDevice entity.FixedDevice
	common.GetDB().Where("id = ?", fixedDeviceId).First(&fixedDevice)
	return fixedDevice
}

func GetPortableDeviceInfoById(portableDeviceId uint) entity.PortableDevice {
	var portableDevice entity.PortableDevice
	common.GetDB().Where("id = ?", portableDeviceId).First(&portableDevice)
	return portableDevice
}

func GetOwnFixedDeviceList(userId uint) []entity.FixedDevice {
	var fixedDeviceList []entity.FixedDevice
	common.GetDB().Table("fixed_devices").Where("owner = ?", userId).Find(&fixedDeviceList)
	return fixedDeviceList
}

func GetOwnPortableDeviceList(userId uint) []entity.PortableDevice {
	var portableDeviceList []entity.PortableDevice
	common.GetDB().Table("portable_devices").Where("owner = ?", userId).Find(&portableDeviceList)
	return portableDeviceList
}
