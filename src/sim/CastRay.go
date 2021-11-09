package sim

import (
	"fmt"
	"github.com/fogleman/gg"
	"github.com/nekiwo/PlantEvolution/src/config"
	"github.com/nekiwo/PlantEvolution/src/helpers"
	"github.com/nekiwo/PlantEvolution/src/plant"
	"math"
)

func CastRay(ray [2][2]int, data *plant.Plant) [][2]int {
	ConvertedRotations := plant.ConvertToPoints(data.Segments)

	RayPolygon := make([][2]int, 0)
	RayPolygon = append(RayPolygon, ray[0])
	RayPolygon = append(RayPolygon, ray[1])

	PointsToBeGiven := 1.00
	record := 2000.0
	RecordType := "none"
	RecordSegment := [2][2]int{{0, 0}, {0, 0}}

	EndLoop := false
	for !EndLoop {
		for b, segment := range ConvertedRotations {
			OverwriteRecord(ray, segment, &record, &RecordType, &RecordSegment, &RayPolygon, "plant", b)
		}
		for b, segment := range config.ImageBorders {
			OverwriteRecord(ray, segment, &record, &RecordType, &RecordSegment, &RayPolygon, "border", b)
		}
		for b, segment := range config.SimBox.Segments {
			OverwriteRecord(ray, segment, &record, &RecordType, &RecordSegment, &RayPolygon, "box", b)
		}

		switch RecordType {
			case "plant":
				data.Points += PointsToBeGiven
				EndLoop = true
				fmt.Println("plant")
				break
			case "box":
				// Slopes to Radian
				SegmentAngle := -helpers.Atan2(
					float64(RecordSegment[1][1] - RecordSegment[0][1]),
					float64(RecordSegment[1][0] - RecordSegment[0][0]),
				) - math.Pi / 2

				RayAngle := -helpers.Atan2(
					float64((RecordSegment[1][1] + ray[1][1]) - (RecordSegment[0][1] + ray[0][1])),
					float64((RecordSegment[1][0] + ray[1][0]) - (RecordSegment[0][0] + ray[0][0])),
				) + math.Pi / 2

				ray[0] = RayPolygon[len(RayPolygon) - 1]
				ray[1] = helpers.DegreeToPoint( gg.Degrees(-RayAngle + 2 * SegmentAngle) - 180, 2000) // Reflect off the surface

				RayPolygon = append(RayPolygon, ray[1])

				PointsToBeGiven -= 0.25

				if PointsToBeGiven < 0.25 {
					EndLoop = true
				}

				//EndLoop = true
				break
			case "border":
				// Absorb the wave and do nothing
				//RayPolygon = append(RayPolygon, ray[1])
				//fmt.Println(ray[1])
				EndLoop = true
				break
			default:
				fmt.Println("???")
		}
	}

	return RayPolygon
}