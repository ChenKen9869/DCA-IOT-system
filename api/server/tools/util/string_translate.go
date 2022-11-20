package util

import (
	"strconv"
	"strings"
)

func String2Float2D(str string) [][]float64 {
	if len(str) == 0 {
		panic("str empty")
	}
	result := [][]float64{}
	s := strings.Trim(str, "[]")
	sList := strings.Split(s, ",")
	for i := 0; i < len(sList); i += 2 {
		longitudeString := strings.Trim(sList[i], "[")
		longitude, errLong := strconv.ParseFloat(longitudeString, 32)
		if errLong != nil {
			panic(errLong.Error())
		}
		latitudeString := strings.Trim(sList[i+1], "]")
		latitude, errLat := strconv.ParseFloat(latitudeString, 32)
		if errLat != nil {
			panic(errLat.Error())
		}
		point := []float64{longitude, latitude}
		result = append(result, point)
	}
	return result
}

func String2Point(str string) (float64, float64) {
	if len(str) == 0 {
		panic("str empty")
	}
	s := strings.Trim(str, "[]")
	sl := strings.Split(s, ",")
	if len(sl) != 2 {
		panic("format error")
	}
	longitude, errLong := strconv.ParseFloat(sl[0], 32)
	if errLong != nil {
		panic(errLong.Error())
	}
	latitude, errLat := strconv.ParseFloat(sl[1], 32)
	if errLat != nil {
		panic(errLat.Error())
	}
	return longitude, latitude
}

func String2ListUint(str string) []uint {
	strList := strings.Split(str, ",")
	result := []uint{}
	for _, s := range strList {
		num, _ := strconv.Atoi(s)
		result = append(result, uint(num))
	}
	return result
}

func ListUint2String(l []uint) string {
	str := ""
	for index, ui := range l {
		str += strconv.Itoa(int(ui))
		if index != (len(l) - 1) {
			str += ","
		}
	}
	return str
}