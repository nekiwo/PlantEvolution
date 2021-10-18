package sim

import "github.com/nekiwo/PlantEvolution/src/helpers"

func OverwriteRecord(ray *[2][2]int, segment *[2][2]int, record *float64, RecordType *string, CurrentType string) {
	HitDist := helpers.CalcDist(ray[0], GetHitPoint(*ray, *segment))
	if HitDist > *record {
		*record = HitDist
		*RecordType = CurrentType
	}
}