package preprosess

import (
	"go-backend/api/common/common"
	"go-backend/api/rule/accepter"
	"go-backend/api/rule/ruleparser"
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
		var rule entity.Rule
		common.GetDB().Table(accepter.DeviceDBMap[typeS]).Where("id = ?", id).First(&rule)
		if rule.ParentId != companyId {
			return false
		}
	}
	return true
}
