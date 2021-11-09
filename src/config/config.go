package config

import (
	"github.com/nekiwo/PlantEvolution/src/box"
)

// Evolution Config:
const TotalPlants int = 10 // Total amount of plants each generation
const TotalGenerations int = 10 // Total number of generation cycles to be simulated

// Plant Config:
const SegmentLength float64 = 100 // Length of each plant segment (pixels)
const MaxSegments int = 20 // Max segments that can be earned
const PointsPerSegment float64 = 5 // Points needed to gain/lose a segment
const PointsThreshold float64 = 5 // Point threshold needed to be achieved to get at least 1 segment

// Simulation/Physics Config:
const TotalRays int = 45 // Amount of rays the sun shoots out
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