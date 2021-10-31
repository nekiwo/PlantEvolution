package sim

func GetHitPoint(ray [2][2]int, segment [2][2]int) [2]int {
	HitPoint := [2]int{ray[0][0] + ray[1][0], ray[0][1] + ray[1][1]}

	// Segment
	x1 := segment[0][0]
	y1 := segment[0][1]
	x2 := segment[1][0]
	y2 := segment[1][1]

	// Ray
	x3 := ray[0][0]
	y3 := ray[0][1]
	x4 := ray[0][0] + ray[1][0]
	y4 := ray[0][1] + ray[1][1]

	// Denominator
	d := (x1 - x2) * (y3 - y4) - (y1 - y2) * (x3 - x4)
	if d == 0 {
		// The line is parallel; no intersection
		return HitPoint
	}

	t := ((x1 - x3) * (y3 - y4) - (y1 - y3) * (x3 - x4)) / d
	u := -((x1 - x2) * (y1 - y3) - (y1 - y2) * (x1 - x3)) / d



	if 0 <= t && t <= 1 && 0 <= u && u <= 1 {
		// There is an intersection; find the hit point
		HitPoint = [2]int{
			x1 + t * (x2 - x1),
			y1 + t * (y2 - y1),
		}
	}

	return HitPoint
}
