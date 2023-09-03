package ruleparser

import (
	"strconv"
	"strings"
)

type Datasuorce struct {
	Name       string
	DeviceId   int
	DeviceType string
	Attribute  string
}

/*
	Datasource =
	name{id, type, attr}; name{id, type, attr}
*/
func ParseDatasource(datasource string) []Datasuorce {

	var result []Datasuorce

	datasource = strings.Replace(datasource, " ", "", -1)
	ds := strings.Split(datasource, ";")

	for _, d := range ds {
		nameAndParams := strings.Split(d, "{")
		name := nameAndParams[0]
		paramString := nameAndParams[1]
		paramString = paramString[:len(paramString)-1]
		params := strings.Split(paramString, ",")
		id, err := strconv.Atoi(params[0])
		if err != nil {
			panic(err.Error())
		}
		typeS := params[1]
		attrS := params[2]
		result = append(result, Datasuorce{
			Name:       name,
			DeviceId:   id,
			DeviceType: typeS,
			Attribute:  attrS,
		})
	}
	return result
}
