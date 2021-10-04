package evo

import (
	"github.com/nekiwo/PlantEvolution/src/config"
	"github.com/nekiwo/PlantEvolution/src/plant"
)

func GenPlants() []plant.Plant {
	plants := make([]plant.Plant, 0)

	for i := 0; i < config.TotalPlants; i++ {
		// Generate a random blob
		plants = append(plants, GenPlant())
	}

	return plants
}
