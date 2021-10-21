package helpers

import (
	"github.com/fogleman/gg"
	"math"
)

func DegreeToPoint(degree float64, radius float64) [2]int {
	return [2]int{
		int(math.Round(math.Cos(gg.Radians(degree)) * radius)),
		int(math.Round(math.Sin(gg.Radians(degree)) * radius)),
	}
}
