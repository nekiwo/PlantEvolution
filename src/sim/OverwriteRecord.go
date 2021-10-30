package sim

import (
	"github.com/nekiwo/PlantEvolution/src/helpers"
)

func OverwriteRecord(ray [2][2]int, segment [2][2]int, record *float64, RecordType *string, RayPolygon *[][2]int, CurrentType string) {
	HitPoint := GetHitPoint(ray, segment)
	HitDist := helpers.CalcDist(ray[0], HitPoint)
	if HitDist < *record {
		*record = HitDist
		*RecordType = CurrentType
		*RayPolygon = append(*RayPolygon, HitPoint)
	}
}