package geocontainer

import "go-backend/api/server/tools/util"

type Point struct {
	Longitude float64
	Latitude  float64
}

type Line struct {
	Start Point
	End   Point
}

type Polygon struct {
	PointSet []Point
}

func (polygon Polygon) Size() int {
	return len(polygon.PointSet)
}

func (polygon *Polygon) InitFromString(position string) {
	pointList := util.String2Float2D(position)
	polygon.PointSet = []Point{}
	for _, point := range pointList {
		polygon.PointSet = append(polygon.PointSet, Point{Longitude: point[0], Latitude: point[1]})
	}
}