package vo

import "time"

type FixedDeviceVo struct {
	Id   uint   `json:"id"`
	Type string `json:"type"`
}

type AuthFixedDevice struct {
	DeviceId    uint      `json:"device_id"`
	DeviceType  string    `json:"device_type"`
	FarmhouseId uint      `json:"farmhouse_id"`
	CreateDate  time.Time `json:"create_date"`
	BoughtDate  time.Time `json:"bought_date"`
	InstallDate	time.Time	`json:"install_date"`
	Stat	string	`json:"status"`
}