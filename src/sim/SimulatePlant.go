package sim

import (
	"github.com/nekiwo/PlantEvolution/src/config"
	"github.com/nekiwo/PlantEvolution/src/graphics"
	"github.com/nekiwo/PlantEvolution/src/helpers"
	"github.com/nekiwo/PlantEvolution/src/plant"
)

func SimulatePlant(data *plant.Plant, directory, ID string) {
	rays := make([][][2]int, 0)

	for i := 0; i < 360; i += 360 / config.TotalRays {
		rays = append(rays, CastRay([2][2]int{
			{config.SimBox.Sun[0], config.SimBox.Sun[1]}, // origin
			helpers.DegreeToPoint(float64(i), 2000), // direction
		}, data))
	}

	graphics.RenderImage(*data, rays, directory, ID)
}