package sim

import "math"

func GetHitPoint(ray *[2][2]int, segment *[2][2]int) [2]int {
	HitPoint := [2]int{math.MaxInt32, math.MaxInt32}

	// Ray
	x1 := ray[0][0]
	y1 := ray[0][1]
	x2 := ray[0][0] + ray[1][0]
	y2 := ray[0][1] + ray[1][1]

	// Segment
	x3 := segment[0][0]
	y3 := segment[0][1]
	x4 := segment[1][0]
	y4 := segment[1][1]

	// Denominator
	d := (x1 - x2) * (y3 - y4) - (y1 - y2) * (x3 - x4)
	if d == 0 {
		// The line is parallel; no intersection
		return HitPoint
	}

	t := ((x1 - x3) * (y3 - y4) - (y1 - y3) * (x3 - x4)) / d
	u := -((x1 - x2) * (y1 - y3) - (y1 - y2) * (x1 - x3)) / d

	if 0 < t && t < 1 && 0 < u {
		// There is an intersection; find the hit point
		HitPoint = [2]int{
			x1 + t * (x2 - x1),
			y1 + t * (y2 - y1),
		}
	}

	return HitPoint
}
