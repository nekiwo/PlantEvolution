package helpers

import "math"

func MedianIndex(length int) int {
	return int(math.Ceil(float64(length) / 2) - 1)
}
