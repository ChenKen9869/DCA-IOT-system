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
	/* Datasource =
	name{id, type, attr}, name{id, type, attr}
	*/
	// 解析出 id, type, attr
	datasourceList := ruleparser.ParseDatasource(datasource)
	for _, ds := range datasourceList {
		id := ds.DeviceId
		typeS := ds.DeviceType
		attrS := ds.Attribute
		// 将 datasource 加入到数据源管理器中
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
	// remove from datasource management
	datasourceList := ruleparser.ParseDatasource(datasource)
	for _, ds := range datasourceList {
		id := ds.DeviceId
		typeS := ds.DeviceType
		attrS := ds.Attribute
		// 将 datasource 加入到数据源管理器中
		deviceIndex := accepter.DeviceIndex{
			Id:         id,
			DeviceType: typeS,
		}
		accepter.DMLock.Lock()
		val := accepter.DatasourceManagement[deviceIndex]
		curr := val[attrS]
		// 引用计数减少 1，若减少为 0，则删除该数据源
		curr.RefNum -= 1
		if curr.RefNum == 0 {
			// remove this entry
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
	// 判断数据源是否全部属于该企业
	/* Datasource =
	name{id, type, attr}, name{id, type, attr}
	*/
	datasourceList := ruleparser.ParseDatasource(datasource)
	for _, ds := range datasourceList {
		id := ds.DeviceId
		typeS := ds.DeviceType
		// auth company and device
		if ds.DeviceType == accepter.PortableDeviceType {
			var device entity.PortableDevice
			fmt.Println(id, typeS+"    "+accepter.DeviceDBMap[typeS].TableName)
			fmt.Println("__________-------------_______")
			common.GetDB().Table(accepter.DeviceDBMap[typeS].TableName).Where("id = ?", id).First(&device)
			pid := dao.GetBiologyInfoById(device.BiologyID).FarmhouseID
			fmt.Println(pid, companyId)
			if pid != companyId {
				return false
			}
		} else if ds.DeviceType == accepter.FixedDeviceType {
			var device entity.FixedDevice
			fmt.Println(id, typeS+"    "+accepter.DeviceDBMap[typeS].TableName)
			fmt.Println("__________-------------_______")
			common.GetDB().Table(accepter.DeviceDBMap[typeS].TableName).Where("id = ?", id).First(&device)
			fmt.Println(device.FarmhouseID, companyId)
			if device.FarmhouseID != companyId {
				return false
			}
		} else {
			panic("device type Error")
		}
	}
	return true
}
