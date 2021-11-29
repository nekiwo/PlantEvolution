package config

import (
	"github.com/nekiwo/PlantEvolution/src/box"
)

// Evolution Config:
const TotalPlants int = 10 // Total amount of plants each generation
const TotalGenerations int = 10 // Total number of generation cycles to be simulated

// Plant Config:
const SegmentLength float64 = 100 // Length of each plant segment (pixels)
const MaxSegments int = 30 // Max segments that can be earned
const PointsToGrow float64 = 1 // Points needed to gain a segment
const PointsToSustain float64 = 0.75 // Points needed to sustain an existing segment

// Simulation/Physics Config:
const TotalRays int = 200 // Amount of rays the sun shoots out
var SimBox box.Box = box.LabyrinthEasy // Chosen box for the simulation
var ImageBorders [][2][2]int = [][2][2]int{
	{{0, 0}, {1080, 0}},
	{{1080, 0}, {1080, 1920}},
	{{1080, 1920}, {0, 1920}},
	{{0, 1920}, {0, 0}},
} // Image borders segments

// Output Config:
const HighlightsOnly bool = false // Keep images of only median, most, and least successful plants (every generation)
const GenerateGraphs bool = false // Keep image graphs for the results
const ShowRays bool = true // Show light rays for debugging

var Test1 [][2]int = [][2]int{}
var Test2 []string = []string{}