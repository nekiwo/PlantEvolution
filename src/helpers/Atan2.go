package helpers

import (
	"math"
)

func Atan2(x, y float64) float64 {
	result := 3434.65

	if x >= 0 {
		result = math.Atan(y/x)
	} else if x * y <= 0 {
		result = math.Atan(y/x) + math.Pi
	} else if x < 0 {
		result = math.Atan(y/x) - math.Pi
	}

	if x == 0 {
		//result = -math.Atan(x/y)
	}

	return result
}