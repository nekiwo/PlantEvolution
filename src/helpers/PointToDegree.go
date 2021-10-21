package helpers

import "math"

func PointToDegree(point [2]int) float64 {
	return math.Atan2(float64(point[1]), float64(point[0]))
}