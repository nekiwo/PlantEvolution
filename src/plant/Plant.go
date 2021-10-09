package plant

type Plant struct {
	GenerationNum int // Generation index
	Origin int // Specie ID (specie IDs are only generated in the 1st gen)
	TotalSegments int // Number of tail segments
	Segments []float64 // Slice of segments rotations
	Points int // Points earned by sun ray hits
}