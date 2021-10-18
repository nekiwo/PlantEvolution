package sim

import (
	"github.com/nekiwo/PlantEvolution/src/config"
	"github.com/nekiwo/PlantEvolution/src/helpers"
	"github.com/nekiwo/PlantEvolution/src/plant"
)

func CastRay(ray [2][2]int) [][2]int {
	ImageBorders := [][2][2]int{
		{{0, 0}, {1080, 0}},
		{{1080, 0}, {1080, 1920}},
		{{1080, 1920}, {0, 1920}},
		{{0, 1920}, {0, 0}},
	}
	ConvertedRotations := plant.ConvertToPoints(config.AllPlants[len(config.AllPlants) - 1].Segments)
	LatestPlant := config.AllPlants[len(config.AllPlants) - 1]

	RayPolygon := make([][2]int, 0)
	RayPolygon = append(RayPolygon, ray[0])

	PointsToBeGiven := 1.00
	record := 9999.00
	RecordType := "none"
	RecordSegment := [2][2]int{{0, 0}, {0, 0}}

	EndLoop := false
	for !EndLoop {
		if PointsToBeGiven == 0 {
			EndLoop = true
		}

		for _, segment := range ConvertedRotations {
			OverwriteRecord(&ray, &segment, &record, &RecordType, "plant")
		}
		for _, segment := range config.SimBox.Segments {
			OverwriteRecord(&ray, &segment, &record, &RecordType, "box")
		}
		for _, segment := range ImageBorders {
			OverwriteRecord(&ray, &segment, &record, &RecordType, "border")
		}

		switch RecordType {
			case "plant":
				LatestPlant.Points += PointsToBeGiven
				EndLoop = true
				break
			case "box":
				SegmentAngle := helpers.SlopeToDegree(helpers.Slope(RecordSegment))

				ray[0] = RayPolygon[len(RayPolygon) - 1]
				ray[1] = helpers.DegreeToPoint(-helpers.PointToDegree(ray[1]) + 360 + SegmentAngle) // Reflect off the surface
				PointsToBeGiven -= 0.25
				break
			case "border":

		}
	}
}