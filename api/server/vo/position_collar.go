package vo

import "time"

type PosCollarData struct {
	Id        uint `json:"id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Altitude  float64 `json:"altitude"`
	Time      time.Time `json:"time"`
}