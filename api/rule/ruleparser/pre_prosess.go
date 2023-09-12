package ruleparser

import (
	"go-backend/api/rule/accepter"
	"go-backend/api/rule/rulelog"
	"go-backend/api/server/dao"
)

func AddDatasource(datasource string) {
	datasourceList := ParseDatasource(datasource)
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
				val[attrS] = keyAttr
			} else {
				val[attrS] = accepter.InitFloatDatasource()
			}
		}
		accepter.DMLock.Unlock()
	}
}

func RemoveDatasource(datasource string) {
	datasourceList := ParseDatasource(datasource)
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
	datasourceList := ParseDatasource(datasource)
	for _, ds := range datasourceList {
		id := ds.DeviceId
		typeS := ds.DeviceType

		if typeS == accepter.PortableDeviceType {
			deviceInfo := dao.GetPortableDeviceInfoById(uint(id))
			pid := dao.GetBiologyInfoById(deviceInfo.BiologyID).FarmhouseID
			if pid != companyId {
				return false
			}
		} else if typeS == accepter.FixedDeviceType {
			deviceInfo := dao.GetFixedDeviceInfoById(uint(id))

			if deviceInfo.FarmhouseID != companyId {
				rulelog.RuleLog.Println("[Auth Rule] Auth Error: Company does not have the access permission to such datasource!")
				return false
			}
		} else {
			panic("[Auth Rule] Syntax Error: Device type does not exist!")
		}
	}
	return true
}
