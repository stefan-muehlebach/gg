package main

import (
	"image/color"

	"github.com/stefan-muehlebach/gg"
)

const (
	outFile = "gradient-linear.png"
)

func main() {
	dc := gg.NewContext(512, 512)

	grad := gg.NewLinearGradient(32, 512-32, 512-32, 32)
	grad.AddColorStop(0, color.RGBA{0, 255, 0, 255})
	grad.AddColorStop(1, color.RGBA{0, 0, 255, 255})
	grad.AddColorStop(0.5, color.RGBA{255, 0, 0, 255})

	dc.SetStrokeColor(color.White)
	dc.DrawRectangle(512/8, 512/8, 384, 384)
	dc.Stroke()

	dc.SetStrokeStyle(grad)
	dc.SetStrokeWidth(10)
    dc.DrawRectangle(20, 20, 390, 90)
	dc.Stroke()

	dc.SetFillStyle(grad)
    dc.DrawRectangle(20, 200, 512-40, 200)
	dc.Fill()

	dc.SavePNG(outFile)
}
