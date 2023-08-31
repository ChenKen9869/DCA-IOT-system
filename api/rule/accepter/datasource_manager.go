package accepter

import (
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

var DatasourceManagement map[DeviceIndex]KeyAttr = make(map[DeviceIndex]KeyAttr)

func InitFloatDatasource() Attribute {
	return Attribute{RefNum: 1, Value: float64(0)}
}
