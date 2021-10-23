package graphics

import (
	"github.com/fogleman/gg"
	"github.com/nekiwo/PlantEvolution/src/config"
	"github.com/nekiwo/PlantEvolution/src/helpers"
	"github.com/nekiwo/PlantEvolution/src/plant"
)

func RenderImage(data plant.Plant, rays [][][2]int, directory string) {
	ImageDimen := [2]float64{float64(config.ImageBorders[1][1][0]), float64(config.ImageBorders[1][1][1])}
	draw := gg.NewContext(int(ImageDimen[0]), int(ImageDimen[1]))

	// Draw background
	draw.SetHexColor("#ffffff")
	draw.DrawRectangle(0, 0, ImageDimen[0], ImageDimen[1])
	draw.Fill()

	ConvertedRotations := plant.ConvertToPoints(data.Segments)

	for _, segment := range ConvertedRotations {
		DrawSegment(draw, segment, "#42ff6e")
	}
	for _, segment := range config.SimBox.Segments {
		DrawSegment(draw, segment, "#ffb742")
	}
	for _, segment := range config.ImageBorders {
		DrawSegment(draw, segment, "#4261ff")
	}

	err := draw.SavePNG("out/" + directory + ".png")
	helpers.ErrCheck(err)
}