package plant

type Plant struct {
	GenerationNum int // Generation index
	Origin int // Specie ID (specie IDs are only generated in the 1st gen)
	Segments []float64 // Slice of segments rotations
	Points float64 // Points earned by sun ray hits
}