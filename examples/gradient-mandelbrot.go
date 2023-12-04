package main

import (
	"github.com/stefan-muehlebach/gg"
	"image/color"
)

const (
	outFile = "gradient-mandelbrot.png"
)

func main() {
	dc := gg.NewContext(1000, 450)

	grad := gg.NewLinearGradient(0, 0, 1000, 0)
	grad.AddColorStop(0.0, color.RGBA{255, 255, 255, 255})
	grad.AddColorStop(0.15, color.RGBA{255, 204, 0, 255})
	grad.AddColorStop(0.33, color.RGBA{135, 31, 19, 255})
	grad.AddColorStop(0.67, color.RGBA{0, 0, 153, 255})
	grad.AddColorStop(0.85, color.RGBA{0, 98, 255, 255})
	grad.AddColorStop(1.0, color.RGBA{255, 255, 255, 255})

	for row := 0; row < 4; row++ {
		for col := 0; col < 9; col++ {
			x := float64(10 + 110*col)
			y := float64(10 + 110*row)
			dc.DrawRectangle(x, y, 100, 100)
			dc.Fill()
		}
	}
	mask := dc.AsMask()
	dc.Clear()

	dc.SetMask(mask)
	dc.DrawRectangle(0, 0, 1000, 450)
	dc.SetFillStyle(grad)
	dc.SetStrokeWidth(5.0)
	dc.SetStrokeColor(color.Black)
	dc.FillStroke()

	dc.SavePNG(outFile)
}
