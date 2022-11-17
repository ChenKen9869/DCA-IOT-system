package vo

import "time"

type OwnBiology struct {
	Id          uint      `json:"id"`
	Name        string    `json:"name"`
	Type        string    `json:"type"`
	Gender      string    `json:"gender"`
	Birthday    time.Time `json:"birthday"`
	CreateTime  time.Time `json:"create_time"`
	FarmhouseId uint      `json:"farmhouse_id"`
}

type OwnFixedDevice struct {
	Id          uint      `json:"id"`
	Type        string    `json:"type"`
	DeviceId	string	  `json:"device_id"`
	InstallTime	time.Time `json:"install_time"`
	CreateTime  time.Time `json:"create_time"`
	FarmhouseId uint      `json:"farmhouse_id"`
	BoughtTime	time.Time `json:"bought_time"`
}

type OwnPortableDevice struct {
	Id          uint      `json:"id"`
	Type        string    `json:"type"`
	DeviceId	string	  `json:"device_id"`
	InstallTime	time.Time `json:"install_time"`
	CreateTime  time.Time `json:"create_time"`
	BiologyId	uint	  `json:"biology_id"`
	BoughtTime	time.Time `json:"bought_time"` 
}