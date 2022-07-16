package vo

type BiologyDevice struct {
	BiologyId		uint		`json:"biologyId"`
	BiologyName		string		`json:"biologyName"`
	BiologyType		string		`json:"biologyType"`
	DeviceId		uint		`json:"deviceId"`
	DeviceType		string		`json:"deviceType"`
}