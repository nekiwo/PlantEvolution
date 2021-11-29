package sim

import (
	"fmt"
	"github.com/nekiwo/PlantEvolution/src/helpers"
)

func OverwriteRecord(ray [2][2]int, segment [2][2]int, record *float64, RecordType *string, RecordSegment *[2][2]int, RayPolygon *[][2]int, CurrentType string, test1 *[][2]int, test2 *[]string) {
	HitPoint := helpers.GetHitPoint(ray, segment)
	HitDist := helpers.CalcDist(ray[0], HitPoint)
	if HitDist < *record {
		*record = HitDist
		*RecordType = CurrentType
		*RecordSegment = segment
		(*RayPolygon)[len(*RayPolygon) - 1] = HitPoint
		*test1 = append(*test1, HitPoint)
		switch CurrentType {
			case "plant":
				*test2 = append(*test2, "#ff0000")
				break
			case "box":
				*test2 = append(*test2, "#0000ff")
				break
			case "border":
				*test2 = append(*test2, "#00ff00")
				break
			default:
				fmt.Println(CurrentType)
		}
	}
}