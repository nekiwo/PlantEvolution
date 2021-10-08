package plant

import (
	"github.com/nekiwo/PlantEvolution/src/config"
)

func GenPlants() []Plant {
	plants := make([]Plant, 0)

	for i := 0; i < config.TotalPlants; i++ {
		// Generate a random blob
		plants = append(plants, GenPlant(i))
	}

	return plants
}
