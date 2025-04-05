package main

import (
	"github.com/stefan-muehlebach/gg"
	"github.com/stefan-muehlebach/gg/colors"
)

const (
	width, height = 512, 512
	marginSize    = 24
	canvasSize    = width - 2*marginSize
	outFile       = "gradient-linear.png"
)

func main() {
	dc := gg.NewContext(512, 512)

	grad := gg.NewLinearGradient(32, 512-32, 512-32, 32)
	grad.AddColorStop(0, colors.RGBAF{0, 1.0, 0, 1.0})
	grad.AddColorStop(1, colors.RGBAF{0, 0, 1.0, 1.0})
	grad.AddColorStop(0.5, colors.RGBAF{1.0, 0, 0, 1.0})

	dc.SetStrokeColor(colors.White)
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
