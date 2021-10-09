package main

import (
	"fmt"
	"github.com/nekiwo/PlantEvolution/src/config"
	"github.com/nekiwo/PlantEvolution/src/evo"
	"github.com/nekiwo/PlantEvolution/src/graphics"
	"github.com/nekiwo/PlantEvolution/src/helpers"
	"github.com/nekiwo/PlantEvolution/src/plant"
	"github.com/nekiwo/PlantEvolution/src/sim"
	"math"
	"os"
	"strconv"
)

func main() {
	//test
	err := os.RemoveAll("out")
	helpers.ErrCheck(err)
	fmt.Println("start")

	// Prepare the simulation
	config.AllPlants = plant.GenPlants() // Gen new blobs
	config.AllGens = append(config.AllGens, []plant.Plant{}) // Add new generation

	// Run every generation
	for i := 0; i < config.TotalGenerations; i++ {
		// Simulate plants
		for ID, blob := range config.AllPlants {
			sim.SimulatePlant(blob, ID)
		}

		// Sort plants results (worst to best)
		SortedPlants := plant.SortByScore(config.AllGens[len(config.AllGens) - 1])

		// Render highlights
		graphics.RenderImage(SortedPlants[0], "worst" + strconv.Itoa(i), "out/highlights/") // Worst performing plant
		graphics.RenderImage(SortedPlants[int(math.Ceil(float64(len(SortedPlants)) / 2) - 1)], "median" + strconv.Itoa(i), "out/highlights/") // Median performing plant
		graphics.RenderImage(SortedPlants[len(SortedPlants) - 1], "best" + strconv.Itoa(i), "out/highlights/") // Best performing plant

		// Delete every other plant render
		if config.HighlightsOnly {
			err = os.RemoveAll("out/all/")
			helpers.ErrCheck(err)
		}

		// Render graphs
		if config.GenerateGraphs {

		}

		// Copy last generation to a new empty generation
		config.AllGens = append(config.AllGens, make([]plant.Plant, 0))
		NextGen := len(config.AllGens) - 1
		config.AllGens[NextGen] = make([]plant.Plant, len(SortedPlants))
		copy(config.AllGens[NextGen], SortedPlants)

		// Kill worst performing plants
		// Copy the remaining
		config.AllGens[NextGen] = evo.KillWorstPlants(config.AllGens[NextGen])

		// Generate new plants by randomizing the last generation
		config.AllGens[NextGen] = evo.RandomizePlants(config.AllGens[NextGen])

		// Copy plants over to AllPlants
		config.AllPlants = config.AllGens[NextGen]

		fmt.Println("gen #" + strconv.Itoa(len(config.AllGens) - 1) + " done")
	}
}