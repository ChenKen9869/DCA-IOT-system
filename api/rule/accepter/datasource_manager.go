package accepter

import (
	"fmt"
	"sync"
)

var DMLock *sync.Mutex = &sync.Mutex{}

type KeyAttr map[string]Attribute

type Attribute struct {
	RefNum uint
	Value  float64
}

type DeviceIndex struct {
	Id         int
	DeviceType string
}

var DatasourceManagement map[DeviceIndex]KeyAttr

func InitFloatDatasource() Attribute {
	return Attribute{RefNum: 1, Value: float64(0)}
}

func updateDatasourceManagement(id int, deviceType string, attr string, value float64) {

	index := DeviceIndex{
		Id:         id,
		DeviceType: deviceType,
	}

	DMLock.Lock()
	v1, exist1 := DatasourceManagement[index]
	if exist1 {
		v, exist := v1[attr]
		if exist {
			v.Value = value
			DatasourceManagement[index][attr] = v
		}
		fmt.Println("[Example Accepter] Datasource management update is complete!")
	} else {
		fmt.Println("[Example Accepter] Datasource management was not updated!")
	}
	DMLock.Unlock()
}
