package plant

import (
	"math/rand"
)

func GenSegments(segments []float64, amount int) []float64 {
	NewSegments := segments

	for i := 0; i < amount; i++ {
		NewSegments = append(NewSegments, rand.Float64() * 90 + 45)
	}

	return NewSegments
}