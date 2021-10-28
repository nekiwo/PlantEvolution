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
	plant.AllPlants = plant.GenPlants() // Gen new plants

	// Run every generation
	for i := 0; i < config.TotalGenerations; i++ {
		// Simulate plants
		for ID, plnt := range plant.AllPlants {
			plnt := plnt // Scope magic to get pointer
			sim.SimulatePlant(&plnt, "all/", "g" + strconv.Itoa(len(plant.AllGens)) + "p" + strconv.Itoa(ID))
		}

		// Sort plants results in ascending order
		SortedPlants := plant.SortByScore(plant.AllPlants)
		plant.AllGens = append(plant.AllGens, SortedPlants) // Add new generation

		// Render highlights
		sim.SimulatePlant(&SortedPlants[0], "highlights/", "worst" + strconv.Itoa(i)) // Worst performing plant
		sim.SimulatePlant(&SortedPlants[helpers.MedianIndex(len(SortedPlants))], "highlights/", "median" + strconv.Itoa(i)) // Median performing plant
		sim.SimulatePlant(&SortedPlants[len(SortedPlants) - 1], "highlights/", "best" + strconv.Itoa(i)) // Best performing plant

		// Delete every other plant render
		if config.HighlightsOnly {
			err = os.RemoveAll("out/all/")
			helpers.ErrCheck(err)
		}

		// Render graphs
		if config.GenerateGraphs {

		}

		// Copy last generation to a new empty generation
		plant.AllGens = append(plant.AllGens, make([]plant.Plant, 0))
		NextGen := len(plant.AllGens) - 1
		plant.AllGens[NextGen] = make([]plant.Plant, len(SortedPlants))
		copy(plant.AllGens[NextGen], SortedPlants)

		// Kill worst performing plants
		// Copy the remaining
		plant.AllGens[NextGen] = evo.KillWorstPlants(plant.AllGens[NextGen])

		// Generate new plants by adding/subtracting segments from the last generation plants
		plant.AllGens[NextGen] = evo.RegulateSegments(plant.AllGens[NextGen])

		// Copy plants over to AllPlants
		plant.AllPlants = plant.AllGens[NextGen]

		fmt.Println("gen #" + strconv.Itoa(len(plant.AllGens) - 1) + " done")
	}
}