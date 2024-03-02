package main

import (
	"image/color"

	"github.com/stefan-muehlebach/gg"
)

const (
	outFile = "gradient-radial.png"
)

func main() {
	dc := gg.NewContext(400, 200)

	grad := gg.NewRadialGradient(100, 100, 100, 300, 100, 40)
	grad.AddColorStop(0.0, color.RGBA{255, 0, 0, 255})
	grad.AddColorStop(0.5, color.RGBA{0, 255, 0, 255})
	grad.AddColorStop(1.0, color.RGBA{0, 0, 255, 255})

	dc.SetFillStyle(grad)
	dc.DrawRectangle(0, 0, 400, 200)
	dc.Fill()

	dc.SetStrokeColor(color.White)
	dc.DrawCircle(100, 100, 100)
	dc.Stroke()
	dc.DrawCircle(300, 100, 40)
	dc.Stroke()

	dc.SavePNG(outFile)
}
