package sim

import (
	"github.com/nekiwo/PlantEvolution/src/config"
	"github.com/nekiwo/PlantEvolution/src/plant"
)

func SimulatePlant(data plant.Plant, ID interface{}) {
	SimulatedPlant := data

	for i := 0; i < config.TotalRays; i++ {
		// shoot ray from the sun
		// check for collision
		// keep a slice of hit coords
		// if colliding with the boundaries, terminate the ray
		// if colliding with the segments, 90 degree bounce
		// repeat the process, but starting from the hit point
		// if colliding with the plant, add point
		// send data for rendering
	}
}