package sim

func GetHitPoint(ray *[2][2]int, segment *[2][2]int) [2]int {
	// if it intersects
	// 		find coordinate
	// 		return it
	// else
	// 		return [2]int{math.Inf(1), math.Inf(1)}

	// Ray
	x1 := *ray[0][0]
	y1 := *ray[0][1]
	x2 := *ray[1][0]
	y2 := *ray[1][1]

	// Segment
	x3 := *segment[0][0]
	y3 := *segment[0][1]
	x4 := *segment[1][0]
	y4 := *segment[1][1]

	t := (x1 - x3)(y1 - y2) - (y1 - y3)(x1 - x2) /
		 (x1 - x2)(y3 - y4) - (y1 - y2)(x3 - x4)

	u := (x1 - x3)(y1 - y2) - (y1 - y3)(x1 - x3) /
		 (x1 - x2)(y3 - y4) - (y1 - y2)(x3 - x4)

	p := [2]int{
		
	}
}
