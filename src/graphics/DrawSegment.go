package graphics

import (
	"github.com/fogleman/gg"
)

func DrawSegment(draw *gg.Context, segment [2][2]int, color string, width float64) {
	draw.SetHexColor(color)
	draw.SetLineWidth(width)
	draw.DrawLine(float64(segment[0][0]), float64(segment[0][1]), float64(segment[1][0]), float64(segment[1][1]))
	draw.Stroke()
}