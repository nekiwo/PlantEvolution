package graphics

import (
	"encoding/hex"
	"github.com/fogleman/gg"
	"github.com/nekiwo/PlantEvolution/src/config"
	"github.com/nekiwo/PlantEvolution/src/helpers"
	"github.com/nekiwo/PlantEvolution/src/plant"
	"math/rand"
	"os"
)

func RenderImage(data plant.Plant, rays [][][2]int, directory string, FileName string) {
	ImageDimen := [2]float64{float64(config.ImageBorders[1][1][0]), float64(config.ImageBorders[1][1][1])}
	draw := gg.NewContext(int(ImageDimen[0]), int(ImageDimen[1]))

	// Draw background
	draw.SetHexColor("#99daff")
	draw.DrawRectangle(0, 0, ImageDimen[0], ImageDimen[1])
	draw.Fill()

	ConvertedRotations := plant.ConvertToPoints(data.Segments)

	for _, segment := range ConvertedRotations {
		DrawSegment(draw, segment, "#228f3c")
	}
	for _, segment := range config.SimBox.Segments {
		DrawSegment(draw, segment, "#ba842d")
	}
	for _, segment := range config.ImageBorders {
		DrawSegment(draw, segment, "#4261ff")
	}

	for _, polygon := range rays {
		GenColor, _ := randomHex(3)
		for i := 0; i < len(polygon) - 1; i++ {
			DrawSegment(draw, [2][2]int{
				polygon[i],
				polygon[i + 1],
			}, "#" + GenColor)
		}
	}

	// Create directory
	err := os.MkdirAll("out/" + directory, os.ModePerm)
	helpers.ErrCheck(err)

	err = draw.SavePNG("out/" + directory + FileName + ".png")
	helpers.ErrCheck(err)
}

func randomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}