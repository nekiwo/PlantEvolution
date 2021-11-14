package helpers

import (
	"math"
)

func RadianToPoint(radian float64, radius float64) [2]int {
	return [2]int{
		int(math.Round(math.Cos(radian) * radius)),
		int(math.Round(math.Sin(radian) * radius)),
	}
}