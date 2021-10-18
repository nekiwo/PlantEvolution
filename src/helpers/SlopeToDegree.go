package helpers

import (
	"github.com/fogleman/gg"
	"math"
)

func SlopeToDegree(slope float64) int {
	return int(math.Round(gg.Degrees(math.Atan(slope))))
}
