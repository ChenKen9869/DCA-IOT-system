package vo

import "time"

type ActiveFence struct {
	Name      	 string 		`json:"name"`
	Id        	 uint 			`json:"id"`
	ExpireTime 	 time.Time		`json:"expiretime"` 
	Position	 string			`json:"position"`
}

type FenceRunningStatus struct {
	Coordinate		string		`json:"coordinate"`
	Position		string		`json:"position"`
	DeviceList 		string		`json:"deviceList"`
	AlarmTime		uint		`json:"alarmTime"`
	Remain			string		`json:"remain"`
}