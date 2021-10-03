package plant

type Plant struct {
	GenerationNum int // Generation index
	Specie int // Specie ID (specie IDs are only generated in the 1st gen)
	TotalSegments int // Number of tail segments
	Segments []Segment // Slice of segments
	Points int // Points earned by sun ray hits
}