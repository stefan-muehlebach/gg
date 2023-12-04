package main

import (
	"github.com/stefan-muehlebach/gg"
	"image/color"
)

const (
	outFile = "rotated-image.png"
)

func main() {
	const W = 400
	const H = 500
	im, err := gg.LoadPNG("gopher.png")
	if err != nil {
		panic(err)
	}
	iw, ih := im.Bounds().Dx(), im.Bounds().Dy()
	dc := gg.NewContext(W, H)
	// draw outline
	dc.SetStrokeColor(color.RGBA{0xff, 0x00, 0x00, 0xff})
	dc.SetStrokeWidth(1)
	dc.DrawRectangle(0, 0, float64(W), float64(H))
	dc.Stroke()
	// draw full image
	dc.SetStrokeColor(color.RGBA{0x00, 0x00, 0xff, 0xff})
	dc.SetStrokeWidth(2)
	dc.DrawRectangle(100, 210, float64(iw), float64(ih))
	dc.Stroke()
	dc.DrawImage(im, 100, 210)
	// draw image with current matrix applied
	dc.SetStrokeColor(color.RGBA{0x00, 0x00, 0xff, 0xff})
	dc.SetStrokeWidth(2)
	dc.Rotate(gg.Radians(10))
	dc.DrawRectangle(100, 0, float64(iw), float64(ih)/2+20.0)
	dc.StrokePreserve()
	dc.Clip()
	dc.DrawImageAnchored(im, 100, 0, 0.0, 0.0)
	dc.SavePNG(outFile)
}
