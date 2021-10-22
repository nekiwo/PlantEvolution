package sim

import (
	"github.com/nekiwo/PlantEvolution/src/config"
	"github.com/nekiwo/PlantEvolution/src/graphics"
	"github.com/nekiwo/PlantEvolution/src/helpers"
	"github.com/nekiwo/PlantEvolution/src/plant"
)

func SimulatePlant(data *plant.Plant, ID string) {
	rays := make([][][2]int, 0)

	for i := 0; i < 360; i += 360 / config.TotalRays {
		// shoot ray from the sun
		// if intersecting a wall/boundary/plant
		// 		if colliding with the boundaries
		// 			terminate the ray
		// 		elif colliding with the segments
		// 			90 degree bounce
		// 			keep a slice of hit coords
		// 		elif colliding with the plant
		// 			add points
		// 			terminate the ray
		// repeat the process, but starting from the hit point
		// send data for rendering

		rays = append(rays, CastRay([2][2]int{
			{config.SimBox.Sun[0], config.SimBox.Sun[1]}, // origin
			helpers.DegreeToPoint(float64(i), 1), // direction
		}, data))
	}

	graphics.RenderImage(*data, ID)
}