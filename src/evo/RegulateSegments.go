package evo

import (
	"github.com/nekiwo/PlantEvolution/src/config"
	"github.com/nekiwo/PlantEvolution/src/plant"
	"math"
)

func RegulateSegments(data []plant.Plant) []plant.Plant {
	NewData := data

	for _, plnt := range NewData {
		if plnt.Points > config.PointsThreshold {
			// Count new segments from points
			NewSegments := int((plnt.Points - math.Mod(plnt.Points, config.PointsPerSegment)) / config.PointsPerSegment)

			// Add new segments to plant
			if len(plnt.Segments) + NewSegments < config.MaxSegments {
				plnt.Segments = plant.GenSegments(plnt.Segments, NewSegments)
			}
		} else {
			// Remove a segment from the plant
			plnt.Segments = plnt.Segments[:len(plnt.Segments) - 1]
		}
	}

	return NewData
}