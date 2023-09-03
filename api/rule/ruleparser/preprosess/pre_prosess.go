package preprosess

import (
	"fmt"
	"go-backend/api/common/common"
	"go-backend/api/rule/accepter"
	"go-backend/api/rule/ruleparser"
	"go-backend/api/server/dao"
	"go-backend/api/server/entity"
)

func AddDatasource(datasource string) {
	datasourceList := ruleparser.ParseDatasource(datasource)
	for _, ds := range datasourceList {
		id := ds.DeviceId
		typeS := ds.DeviceType
		attrS := ds.Attribute
		deviceIndex := accepter.DeviceIndex{
			Id:         id,
			DeviceType: typeS,
		}
		accepter.DMLock.Lock()
		val, exist := accepter.DatasourceManagement[deviceIndex]
		if !exist {
			accepter.DatasourceManagement[deviceIndex] = make(accepter.KeyAttr)
			accepter.DatasourceManagement[deviceIndex][attrS] = accepter.InitFloatDatasource()
		} else {
			keyAttr, existA := val[attrS]
			if existA {
				keyAttr.RefNum += 1
			} else {
				val[attrS] = accepter.InitFloatDatasource()
			}
		}
		accepter.DMLock.Unlock()
	}
}

func RemoveDatasource(datasource string) {
	datasourceList := ruleparser.ParseDatasource(datasource)
	for _, ds := range datasourceList {
		id := ds.DeviceId
		typeS := ds.DeviceType
		attrS := ds.Attribute
		deviceIndex := accepter.DeviceIndex{
			Id:         id,
			DeviceType: typeS,
		}
		accepter.DMLock.Lock()
		val := accepter.DatasourceManagement[deviceIndex]
		curr := val[attrS]
		curr.RefNum -= 1
		if curr.RefNum == 0 {
			delete(val, attrS)
			if len(val) == 0 {
				delete(accepter.DatasourceManagement, deviceIndex)
			}
		} else {
			val[attrS] = curr
		}
		accepter.DMLock.Unlock()
	}
}

func AuthDevices(datasource string, companyId uint) bool {
	datasourceList := ruleparser.ParseDatasource(datasource)
	for _, ds := range datasourceList {
		id := ds.DeviceId
		typeS := ds.DeviceType
		// auth company and device
		if ds.DeviceType == accepter.PortableDeviceType {
			var device entity.PortableDevice
			common.GetDB().Table(accepter.DeviceDBMap[typeS].TableName).Where("id = ?", id).First(&device)
			pid := dao.GetBiologyInfoById(device.BiologyID).FarmhouseID
			if pid != companyId {
				return false
			}
		} else if ds.DeviceType == accepter.FixedDeviceType {
			var device entity.FixedDevice
			common.GetDB().Table(accepter.DeviceDBMap[typeS].TableName).Where("id = ?", id).First(&device)
			if device.FarmhouseID != companyId {
				fmt.Println("[Auth Rule] Auth Error: Company does not have the access permission to such datasource!")
				return false
			}
		} else {
			panic("[Auth Rule] Syntax Error: Device type does not exist!")
		}
	}
	return true
}
