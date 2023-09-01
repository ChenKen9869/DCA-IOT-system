package accepter

type DeviceType = string
type DBTable struct {
	TableName string
	ColumnName string
}

var DeviceDBMap map[DeviceType]DBTable 
