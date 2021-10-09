package box

type Box struct {
	Segments [][2][2]int // Slice of coords for box segments
	Sun [2]int // Coords for the light source
}