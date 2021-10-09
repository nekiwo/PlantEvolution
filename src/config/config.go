package config

import "github.com/nekiwo/PlantEvolution/src/plant"

// Evolution Config:
const TotalPlants int = 10 // Total amount of plants each generation
const TotalGenerations int = 10 // Total number of generation cycles to be simulated

// Plant Config:
const SegmentLength float64 = 50 // Length of each plant segment (pixels)

// Simulation/Physics Config:
const TotalRays int = 45 // Amount of rays the sun shoots out

// Output Config:
const HighlightsOnly bool = true // Keep images of only median, most, and least successful plants (every generation)
const GenerateGraphs bool = false // Keep image graphs for the results

// Global vars:
var AllGens [][]plant.Plant // All simulated blobs of all time
var AllPlants []plant.Plant // Plants to be simulated