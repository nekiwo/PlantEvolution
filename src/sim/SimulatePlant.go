package sim

import (
	"github.com/nekiwo/PlantEvolution/src/config"
	"github.com/nekiwo/PlantEvolution/src/plant"
	"github.com/fogleman/gg"
)

func SimulatePlant(data plant.Plant, ID interface{}) {
	SimulatedPlant := data

	for i := 0; i < 360; i += 360 / config.TotalRays {

		CastRay(gg.Radians(float64(i)))
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
	}
}