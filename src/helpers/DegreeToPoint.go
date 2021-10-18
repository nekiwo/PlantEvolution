package helpers

import (
	"github.com/fogleman/gg"
	"math"
)

func DegreeToPoint(degree int) [2]int {
	return [2]int{
		int(math.Round(math.Cos(gg.Radians(float64(degree))))),
		int(math.Round(math.Sin(gg.Radians(float64(degree))))),
	}
}
