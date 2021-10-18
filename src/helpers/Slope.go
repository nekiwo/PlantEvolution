package helpers

func Slope(segment [2][2]int) float64 {
	return float64(segment[1][1] - segment[0][1]) / float64(segment[1][0] - segment[0][0])
}
