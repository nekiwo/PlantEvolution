package helpers

import "math"

func CalcDist(a [2]int, b [2]int) float64 {
	return math.Sqrt(
		float64((a[0] - b[0]) * (a[0] - b[0])) + float64((a[1] - b[1]) * (a[1] - b[1])),
	)
}