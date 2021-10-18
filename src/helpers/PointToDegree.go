package helpers

import "math"

func PointToDegree(point [2]int) int {
	return int(math.Atan2(float64(point[1]), float64(point[0])))
}
