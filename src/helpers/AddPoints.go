package helpers

func AddPoints(a [2]int, b [2]int) [2]int {
	return [2]int{
		a[0] + b[0],
		a[1] + b[1],
	}
}