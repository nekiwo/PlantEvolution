package plant

import (
	"github.com/nekiwo/PlantEvolution/src/config"
	"github.com/nekiwo/PlantEvolution/src/helpers"
	"math/rand"
)

func GenSegments(segments []float64, amount int) []float64 {
	NewSegments := segments

	for i := 0; i < amount; i++ {
		// Pre-calculated
		ConvertedSegments := ConvertToPoints(NewSegments)
		PlantTopPoint := ConvertedSegments[len(ConvertedSegments) - 1][1]

		// Check if a new segment is possible by checking intersections of 3 positions:
		// 0, 45, 90 degrees
		// Not a complete method but it is works well enough
		AnglesToCheck := [3]float64{0, 45, 90}
		IsIntersecting := false

		for _, angle := range AnglesToCheck {
			SegmentCheck := [2][2]int{
				PlantTopPoint,
				helpers.DegreeToPoint(angle - 135, config.SegmentLength),
			}

			for _, BoxSegment := range config.SimBox.Segments {
				Intersection := helpers.GetHitPoint(SegmentCheck, BoxSegment)
				if Intersection != helpers.AddPoints(SegmentCheck[0], SegmentCheck[1]) {
					IsIntersecting = true
				}
			}
		}

		if !IsIntersecting {
			// Repeat search for a non-intersecting angle until it is found
			result := 0.0
			for result == 0.0 {
				RandRot := rand.Float64() * 90 - 135
				NewSegment := [2][2]int{
					PlantTopPoint,
					helpers.DegreeToPoint(RandRot, config.SegmentLength),
				}

				IsIntersectingNew := false

				for _, BoxSegment := range config.SimBox.Segments {
					Intersection := helpers.GetHitPoint(NewSegment, BoxSegment)
					if Intersection != helpers.AddPoints(NewSegment[0], NewSegment[1]) {
						// Intersecting a segment
						IsIntersectingNew = true
					}
				}

				if !IsIntersectingNew {
					result = RandRot
				}
			}

			NewSegments = append(NewSegments, result)
		}
	}

	return NewSegments
}