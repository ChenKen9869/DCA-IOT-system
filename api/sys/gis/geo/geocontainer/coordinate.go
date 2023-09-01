package geocontainer

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

var Coordinates map[string]bool

const (
	GCJ_02 = "GCJ-02"
	WGS84 = "WGS84"
	BD_09 = "BD-09"
)

func InitContainer() {
	coordinateList := viper.GetString("geo.coordinates")
	coordinates := strings.Split(coordinateList, ",")
	InitCoordinates()
	for _, coordinate := range coordinates {
		if _, exists := Coordinates[coordinate]; exists {
			Coordinates[coordinate] = true
		}
	}
	fmt.Println("[INITIAL SUCCESS] The gis module is initialized successfully!")
}

func InitCoordinates() {
	Coordinates = make(map[string]bool)
	Coordinates[GCJ_02] = false
	Coordinates[WGS84] = false
	Coordinates[BD_09] = false
}