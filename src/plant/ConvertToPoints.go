package plant

import (
	"github.com/nekiwo/PlantEvolution/src/config"
	"github.com/nekiwo/PlantEvolution/src/helpers"
)

func ConvertToPoints(rotations []float64) [][2][2]int {
	points := make([][2][2]int, 0)
	LastPoint := [2]int{540, 0}

	for _, rot := range rotations {
		points = append(points, [2][2]int{
			LastPoint,
			helpers.AddPoints(LastPoint, helpers.DegreeToPoint(rot, config.SegmentLength)),
		})
		LastPoint = points[len(points) - 1][1]
	}

	return points
}
