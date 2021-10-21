package main

import (
	"fmt"
	"github.com/nekiwo/PlantEvolution/src/config"
	"github.com/nekiwo/PlantEvolution/src/evo"
	"github.com/nekiwo/PlantEvolution/src/helpers"
	"github.com/nekiwo/PlantEvolution/src/plant"
	"github.com/nekiwo/PlantEvolution/src/sim"
	"os"
	"strconv"
)

func main() {
	//test
	err := os.RemoveAll("out")
	helpers.ErrCheck(err)
	fmt.Println("start")

	// Prepare the simulation
	config.AllPlants = plant.GenPlants() // Gen new plants
	config.AllGens = append(config.AllGens, []plant.Plant{}) // Add new generation

	// Run every generation
	for i := 0; i < config.TotalGenerations; i++ {
		// Simulate plants
		for ID, plnt := range config.AllPlants {
			plnt := plnt // Scope magic to get pointer
			sim.SimulatePlant(&plnt, "g" + strconv.Itoa(len(config.AllGens)) + "p" + strconv.Itoa(ID))
		}

		// Sort plants results (worst to best)
		SortedPlants := plant.SortByScore(config.AllGens[len(config.AllGens) - 1])

		// Render highlights
		sim.SimulatePlant(&SortedPlants[0], "highlights/worst" + strconv.Itoa(i), ) // Worst performing plant
		sim.SimulatePlant(&SortedPlants[helpers.MedianIndex(len(SortedPlants))], "highlights/median" + strconv.Itoa(i)) // Median performing plant
		sim.SimulatePlant(&SortedPlants[len(SortedPlants) - 1], "highlights/best" + strconv.Itoa(i)) // Best performing plant

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

		// Generate new plants by adding/subtracting segments from the last generation plants
		config.AllGens[NextGen] = evo.RegulateSegments(config.AllGens[NextGen])

		// Copy plants over to AllPlants
		config.AllPlants = config.AllGens[NextGen]

		fmt.Println("gen #" + strconv.Itoa(len(config.AllGens) - 1) + " done")
	}
}