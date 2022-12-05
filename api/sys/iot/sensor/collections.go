package sensor

import (
	"github.com/spf13/viper"
)

var collarCollection string
var fioCollection string
var posCollarCollection string
func InitCollections() {
	collarCollection = viper.GetString("mongodb.collar")
	fioCollection = viper.GetString("mongodb.fio")
	posCollarCollection = viper.GetString("mongodb.position-collar")
}