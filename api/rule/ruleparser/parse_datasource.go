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
	/*
		datasource=
		name{id,type,attr},name{id,type,attr}
	*/
	ds := strings.Split(datasource, ";")
	/*
		ds=
		elem-01: name{id,type,attr}
		elem-02: name{id,type,attr}
	*/
	for _, d := range ds {
		nameAndParams := strings.Split(d, "{")
		/*
			nameAndParams=
			elem-01: name
			elem-02: id,type,attr
		*/
		name := nameAndParams[0]
		paramString := nameAndParams[1]
		paramString = paramString[:len(paramString)-1]

		params := strings.Split(paramString, ",")
		/*
			params=
			elem-01: id
			elem-02: type
			elem-03: attr
		*/
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
