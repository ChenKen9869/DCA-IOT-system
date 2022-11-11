package vo

import "time"

type FioData struct {
	Id          uint		`json:"id"`
	Humidity    float32		`json:"humidity"`
	Temperature float32		`json:"temperature"`
	Methane     float32		`json:"methane"`
	Ammonia     float32		`json:"ammonia"`
	Hydrogen    float32		`json:"hydrogen"`
	Time        time.Time 	`json:"time"`
}