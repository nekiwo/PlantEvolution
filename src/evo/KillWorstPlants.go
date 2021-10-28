package evo

import (
	"github.com/nekiwo/PlantEvolution/src/config"
	"github.com/nekiwo/PlantEvolution/src/plant"
	"math"
)

func KillWorstPlants(data []plant.Plant) []plant.Plant {
	NewData := make([]plant.Plant, 0)

	// Clone top half of blobs
	HalfBlobNum := int(math.Floor(float64(config.TotalPlants) / 2))
	for i := 0; i < HalfBlobNum; i++ {
		NewData = append(NewData, data[len(data) - i - 1])
		NewData = append(NewData, data[len(data) - i - 1]) // Clone the other half to fill the gap
	}

	// Fix number of blobs for odd numbered slices
	if len(NewData) != config.TotalPlants {
		NewData = append(NewData, data[len(data) - 1])
	}

	return NewData
}