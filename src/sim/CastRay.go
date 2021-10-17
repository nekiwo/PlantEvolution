package sim

import (
	"github.com/nekiwo/PlantEvolution/src/config"
	"github.com/nekiwo/PlantEvolution/src/plant"
)

func CastRay(a [2][2]int) [][2]int {
	ImageBorders := [][2][2]int{
		{{0, 0}, {1080, 0}},
		{{1080, 0}, {1080, 1920}},
		{{1080, 1920}, {0, 1920}},
		{{0, 1920}, {0, 0}},
	}

	RayPolygon := make([][2]float64, 0)

	ConvertedRotations := plant.ConvertToPoints(config.AllPlants[len(config.AllPlants) - 1].Segments)
	for _, segment := range ConvertedRotations {

	}

	for _, segment := range config.SimBox.Segments {

	}

	for _, segment := range ImageBorders {

	}
}