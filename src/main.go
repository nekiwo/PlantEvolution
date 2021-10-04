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

// Global vars:
var AllGens [][]plant.Plant // All simulated blobs of all time
var AllPlants []plant.Plant // Plants to be simulated

func main() {
	//test
	err := os.RemoveAll("out")
	helpers.ErrCheck(err)
	fmt.Println("start")

	// Prepare the simulation
	evo.GenPlants(config.TotalPlants)
	AllGens = append(AllGens, []plant.Plant{})

	// Run every generation
	for i := 0; i < config.TotalGenerations; i++ {
		// Simulate blobs
		for ID, blob := range AllPlants {
			sim.SimulatePlant(blob, ID)
		}

		// Sort plants results (worst to best)
		SortedPlants := plant.SortByScore(AllGens[len(AllGens) - 1])

		// Render highlights
		graphics.RenderImage(SortedPlants[0], "worst" + strconv.Itoa(i), "out/highlights/") // Worst performing blob
		graphics.RenderImage(SortedPlants[int(math.Ceil(float64(len(SortedPlants)) / 2) - 1)], "median" + strconv.Itoa(i), "out/highlights/") // Median performing blob
		graphics.RenderImage(SortedPlants[len(SortedPlants) - 1], "best" + strconv.Itoa(i), "out/highlights/") // Best performing blob

		// Delete every other plant render
		if config.HighlightsOnly {
			err = os.RemoveAll("out/all/")
			helpers.ErrCheck(err)
		}

		// Render graphs
		if config.GenerateGraphs {

		}

		// Copy last generation to a new empty generation
		AllGens = append(AllGens, make([]BlobData, 0))
		NextGen := len(AllGens) - 1
		AllGens[NextGen] = make([]BlobData, len(SortedBlobs))
		copy(AllGens[NextGen], SortedBlobs)

		// Kill worst performing blobs
		// Copy the remaining
		AllGens[NextGen] = KillWorstBlobs(AllGens[NextGen])

		// Generate new blobs by randomizing the last generation
		AllGens[NextGen] = RandomizeBlobs(AllGens[NextGen])

		// Copy plants over to AllPlants
		AllPlants = AllGens[NextGen]

		fmt.Println("gen #" + strconv.Itoa(len(AllGens) - 1) + " done")
	}
}