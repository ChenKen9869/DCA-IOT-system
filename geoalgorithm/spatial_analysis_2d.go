package geoalgorithm

import "go-backend/geocontainer"

// 点面分析函数 : 引射线法
func PolygonContainsPoint(polygon geocontainer.Polygon, point geocontainer.Point) bool {
	longitude := point.Longitude
	latitude := point.Latitude
	crossTime := 0
	for i := 0; i < polygon.Size() - 1; i++ {
		start := polygon.PointSet[i]
		end := polygon.PointSet[i + 1]
		if (latitude > start.Latitude && latitude < end.Latitude) || (start.Latitude > latitude && latitude > end.Latitude) {
			duT := (latitude - start.Latitude) / (end.Latitude - start.Latitude)
			duXT := start.Longitude + duT * (end.Longitude - start.Longitude)
			if duXT > longitude {
				crossTime++
			}		
		}
	}
	startPoint := polygon.PointSet[polygon.Size() - 1]
	endPoint := polygon.PointSet[0]
	if (latitude > startPoint.Latitude && latitude < endPoint.Latitude) || (startPoint.Latitude > latitude && latitude > endPoint.Latitude) {
		duT := (latitude - startPoint.Latitude) / (endPoint.Latitude - startPoint.Latitude)
		duXT := startPoint.Longitude + duT * (endPoint.Longitude - startPoint.Longitude)
		if duXT > longitude {
			crossTime++
		}		
	}
	return (crossTime % 2) != 0
}
