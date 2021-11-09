package sim

import (
	"fmt"
	"github.com/nekiwo/PlantEvolution/src/helpers"
)

func OverwriteRecord(ray [2][2]int, segment [2][2]int, record *float64, RecordType *string, RecordSegment *[2][2]int, RayPolygon *[][2]int, CurrentType string, test int) {
	HitPoint := GetHitPoint(ray, segment, &CurrentType)
	HitDist := helpers.CalcDist(ray[0], HitPoint)
	if HitDist < *record {
		*record = HitDist
		*RecordType = CurrentType
		*RecordSegment = segment
		(*RayPolygon)[len(*RayPolygon) - 1] = HitPoint

		if CurrentType == "plant" {
			fmt.Println(CurrentType)
			fmt.Println("===========")
		}
		fmt.Println(test)
	}
}