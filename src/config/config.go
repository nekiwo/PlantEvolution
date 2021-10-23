package config

import (
	"github.com/nekiwo/PlantEvolution/src/box"
	"github.com/nekiwo/PlantEvolution/src/plant"
)

// Evolution Config:
const TotalPlants int = 10 // Total amount of plants each generation
const TotalGenerations int = 10 // Total number of generation cycles to be simulated

// Plant Config:
const SegmentLength float64 = 200 // Length of each plant segment (pixels)

// Simulation/Physics Config:
const TotalRays int = 45 // Amount of rays the sun shoots out
const PointsGoal float64 = 5 // Points needed to gain/lose a segment
var SimBox box.Box = box.LabyrinthEasy // Chosen box for the simulation
var ImageBorders [][2][2]int = [][2][2]int{
	{{0, 0}, {1080, 0}},
	{{1080, 0}, {1080, 1920}},
	{{1080, 1920}, {0, 1920}},
	{{0, 1920}, {0, 0}},
} // Image borders segments

// Output Config:
const HighlightsOnly bool = true // Keep images of only median, most, and least successful plants (every generation)
const GenerateGraphs bool = false // Keep image graphs for the results
const ShowRays bool = true // Show light rays for debugging



// Global vars:
var AllGens [][]plant.Plant // All simulated plants of all generations
var AllPlants []plant.Plant // Plants to be simulated