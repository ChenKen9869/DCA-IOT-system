package vo

import "time"

type BiologyPortableDevice struct {
	Id          uint      `json:"id"`
	Type        string    `json:"type"`
	DeviceId    string    `json:"device_id"`
	InstallTime time.Time `json:"install_time"`
	CreateTime  time.Time `json:"create_time"`
	BoughtTime  time.Time `json:"bought_time"`
	Stat 		string	  `json:"stat"`
}