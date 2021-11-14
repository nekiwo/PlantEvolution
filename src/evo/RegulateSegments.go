package evo

import (
	"github.com/nekiwo/PlantEvolution/src/config"
	"github.com/nekiwo/PlantEvolution/src/plant"
	"math"
)

func RegulateSegments(data []plant.Plant) []plant.Plant {
	NewData := data

	for i, plnt := range NewData {
		RequiredPoints := config.PointsToSustain * float64(len(plnt.Segments)) // Points needed to sustain every segment
		if plnt.Points > RequiredPoints {
			// Count new segments from points
			NewSegments := math.Floor((plnt.Points - RequiredPoints) / config.PointsToGrow)
			AdjustedNum := int(math.Min(NewSegments, float64(config.MaxSegments)))

			// Add new segments to plant
			NewData[i].Segments = plant.GenSegments(plnt.Segments, AdjustedNum)
		} else {
			// Remove segments from the plant
			SustainedSegments := int(math.Max(math.Floor(plnt.Points / config.PointsToSustain), 1))
			NewData[i].Segments = plnt.Segments[:SustainedSegments]
		}
	}

	return NewData
}