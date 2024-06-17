package main

import (
	"github.com/stefan-muehlebach/gg"
	"github.com/stefan-muehlebach/gg/color"
)

const (
	width, height  = 512, 512
	marginSize     = 24
	canvasSize     = width - 2*marginSize
	colorStripSize = 63.0
	outFile        = "gradient-transformed.png"
)

func main() {
	dc := gg.NewContext(width, height)

	grad := gg.NewLinearGradient(0, 0, canvasSize, 0)
	grad.AddColorStop(0.0, color.RGBAF{1.0, 0, 0, 0})
	grad.AddColorStop(0.25, color.RGBAF{1.0, 0, 0, 1.0})
	grad.AddColorStop(0.5, color.RGBAF{0, 1.0, 0, 1.0})
	grad.AddColorStop(0.75, color.RGBAF{0, 0, 1.0, 1.0})
	grad.AddColorStop(1.0, color.RGBAF{0, 0, 1.0, 0})

	dc.SetFillColor(color.Black)
	dc.Clear()

	dc.SetFillStyle(grad)
	dc.DrawRectangle(width/2.0, height/2.0-colorStripSize/2.0, canvasSize, colorStripSize)
	dc.Fill()

	dc.SavePNG(outFile)
}
