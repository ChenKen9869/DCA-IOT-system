package util

import "math"

const MIN = 0.000001

func IsFloat64Equal(x, y float64) bool {
	return math.Abs(x-y) < MIN
}
