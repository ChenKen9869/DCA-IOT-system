package sensor

import (
	"github.com/spf13/viper"

)

var collarCollection string
var fioCollection string

func InitCollections() {
	collarCollection = viper.GetString("mongodb.collar")
	fioCollection = viper.GetString("mongodb.fio")
}