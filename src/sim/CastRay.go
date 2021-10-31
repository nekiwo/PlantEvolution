package sim

import (
	"fmt"
	"github.com/nekiwo/PlantEvolution/src/config"
	"github.com/nekiwo/PlantEvolution/src/helpers"
	"github.com/nekiwo/PlantEvolution/src/plant"
)

func CastRay(ray [2][2]int, data *plant.Plant) [][2]int {
	ConvertedRotations := plant.ConvertToPoints(data.Segments)

	RayPolygon := make([][2]int, 0)
	RayPolygon = append(RayPolygon, ray[0])

	PointsToBeGiven := 1.00
	record := 2000.0
	RecordType := "none"
	RecordSegment := [2][2]int{{0, 0}, {0, 0}}

	EndLoop := false
	for !EndLoop {
		if PointsToBeGiven == 0 {
			EndLoop = true
		}

		for _, segment := range config.ImageBorders {
			OverwriteRecord(ray, segment, &record, &RecordType, &RayPolygon, "border")
		}
		for _, segment := range ConvertedRotations {
			OverwriteRecord(ray, segment, &record, &RecordType, &RayPolygon, "plant")
		}
		for _, segment := range config.SimBox.Segments {
			OverwriteRecord(ray, segment, &record, &RecordType, &RayPolygon, "box")
		}

		switch RecordType {
			case "plant":
				fmt.Println("plant")
				data.Points += PointsToBeGiven
				EndLoop = true
				break
			case "box":
				//fmt.Println("box")
				SegmentAngle := helpers.SlopeToDegree(helpers.Slope(RecordSegment))
				ray[0] = RayPolygon[len(RayPolygon) - 1]
				ray[1] = helpers.DegreeToPoint(-helpers.PointToDegree(ray[1]) + 360 + float64(SegmentAngle), 2000) // Reflect off the surface
				PointsToBeGiven -= 0.25
				break
			case "border":
				//fmt.Println("border")
				// Absorb the wave and do nothing
				EndLoop = true
				break
			default:
				fmt.Println("???")
		}
	}

	return RayPolygon
}