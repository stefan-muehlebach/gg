package main

import (
	"image/color"

	"github.com/stefan-muehlebach/gg"
)

const (
	W = 1024
	H = 512
    outFile = "gradient-text.png"
)

func main() {
	dc := gg.NewContext(W, H)

	// draw text
	dc.SetStrokeColor(color.Black)
	dc.LoadFontFace("Ubuntu-Bold.ttf", 128)
	dc.DrawStringAnchored("Gradient Text", W/2, H/2, 0.5, 0.5)

	// get the context as an alpha mask
	mask := dc.AsMask()

	// clear the context
	dc.SetFillColor(color.White)
	dc.Clear()

	// set a gradient
	g := gg.NewLinearGradient(0, 0, W, H)
	g.AddColorStop(0, color.RGBA{255, 0, 0, 255})
	g.AddColorStop(1, color.RGBA{0, 0, 255, 255})
	dc.SetFillStyle(g)

	// using the mask, fill the context with the gradient
	dc.SetMask(mask)
	dc.DrawRectangle(0, 0, W, H)
	dc.Fill()

	dc.SavePNG(outFile)
}
