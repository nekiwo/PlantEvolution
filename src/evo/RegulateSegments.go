package evo

import (
	"github.com/nekiwo/PlantEvolution/src/config"
	"github.com/nekiwo/PlantEvolution/src/plant"
	"math"
)

func RegulateSegments(data []plant.Plant) []plant.Plant {
	NewData := data

	for i, plnt := range NewData {
		if plnt.Points > config.PointsThreshold {
			// Count new segments from points
			NewSegments := int((plnt.Points - math.Mod(plnt.Points, config.PointsPerSegment)) / config.PointsPerSegment)

			// Add new segments to plant
			if len(plnt.Segments) + NewSegments < config.MaxSegments {
				NewData[i].Segments = plant.GenSegments(plnt.Segments, NewSegments)
			}
		} else {
			// Remove a segment from the plant
			NewData[i].Segments = plnt.Segments[:int(math.Max(float64(len(plnt.Segments)), 2)) - 1]
		}
	}

	/*for _, p := range NewData {
		fmt.Println(len(p.Segments))
	}*/

	return NewData
}