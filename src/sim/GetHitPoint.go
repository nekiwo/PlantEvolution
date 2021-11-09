package sim

import (
	"fmt"
	"math"
)

func GetHitPoint(ray [2][2]int, segment [2][2]int, test *string) [2]int {
	HitPoint := [2]int{ray[0][0] + ray[1][0], ray[0][1] + ray[1][1]}

	// Segment
	x1 := float64(segment[0][0])
	y1 := float64(segment[0][1])
	x2 := float64(segment[1][0])
	y2 := float64(segment[1][1])

	// Ray
	x3 := float64(ray[0][0])
	y3 := float64(ray[0][1])
	x4 := float64(ray[0][0] + ray[1][0])
	y4 := float64(ray[0][1] + ray[1][1])

	// Denominator
	d := (x1 - x2) * (y3 - y4) - (y1 - y2) * (x3 - x4)
	if d == 0 {
		// The line is parallel; no intersection
		return HitPoint
	}

	t := ((x1 - x3) * (y3 - y4) - (y1 - y3) * (x3 - x4)) / d
	u := -((x1 - x2) * (y1 - y3) - (y1 - y2) * (x1 - x3)) / d

	if 0 < t && t < 1 && 0 < u && u < 1 {
		// There is an intersection; find the hit point
		HitPoint = [2]int{
			int(math.Round(x1 + t * (x2 - x1))),
			int(math.Round(y1 + t * (y2 - y1))),
		}
		/*fmt.Println("========")
		fmt.Println(x1,y1,x2,y2)
		fmt.Println(x3,y3,x4,y4)
		fmt.Println(t,u,d)*/
		fmt.Println(*test)
	}

	return HitPoint
}
