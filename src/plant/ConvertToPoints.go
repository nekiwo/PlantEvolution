package plant

import (
	"github.com/nekiwo/PlantEvolution/src/config"
	"github.com/nekiwo/PlantEvolution/src/helpers"
)

func ConvertToPoints(rotations []float64) [][2][2]int {
	points := make([][2][2]int, 0)
	LastPoint := [2]int{config.ImageBorders[1][1][0] / 2, config.ImageBorders[1][1][1]} // Bottom middle of the image

	for _, rot := range rotations {
		points = append(points, [2][2]int{
			LastPoint,
			helpers.AddPoints(LastPoint, helpers.DegreeToPoint(rot, config.SegmentLength)),
		})
		LastPoint = points[len(points) - 1][1]
	}

	if len(rotations) == 0 {
		points = append(points, [2][2]int{
			LastPoint,
			LastPoint,
		})
	}

	return points
}
