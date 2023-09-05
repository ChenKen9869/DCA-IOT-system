package matcher

import (
	"fmt"
	"go-backend/api/rule/ruleparser"
	"go-backend/api/sys/gis/geo/geoalgorithm"
	"go-backend/api/sys/gis/geo/geocontainer"
	"strconv"
)

func MatchPointSurfaceFunctionCondition(tokenList []ruleparser.Token, innerTable ruleparser.InnerTable) bool {
	var point geocontainer.Point
	var polygon geocontainer.Polygon
	for symbol, value := range innerTable {
		if symbol == tokenList[0].TokenValue {
			point.Longitude = value
			break
		}
	}
	for symbol, value := range innerTable {
		if symbol == tokenList[1].TokenValue {
			point.Latitude = value
			break
		}
	}
	var numList []float64
	for _, token := range tokenList[2:] {
		currNum, err := strconv.ParseFloat(token.TokenValue, 64)
		if err != nil {
			fmt.Println("[Point Surface Rule Matcher] Error Occur: " + err.Error())
		}
		numList = append(numList, currNum)
	}
	for index := range numList {
		if index%2 == 0 {
			polygon.PointSet = append(polygon.PointSet, geocontainer.Point{
				Longitude: numList[index],
				Latitude:  numList[index+1],
			})
		}
	}
	return geoalgorithm.PolygonContainsPoint(polygon, point)
}
