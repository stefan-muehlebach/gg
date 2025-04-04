package main

import (
	"github.com/stefan-muehlebach/gg"
	"github.com/stefan-muehlebach/gg/colors"
	"github.com/stefan-muehlebach/gg/fonts"
)

const (
	outFile  = "rotated-text.png"
	size     = 512.0
	fontSize = 40.0
)

func main() {
	const S = 400
	dc := gg.NewContext(size, size)
	dc.SetFillColor(colors.White)
	dc.Clear()

	dc.SetStrokeColor(colors.Black)
	dc.SetFontFace(fonts.NewFace(fonts.GoRegular, fontSize))
	text := "Hello, world!"
	w, h := dc.MeasureString(text)
	dc.RotateAbout(gg.Radians(10), size/2, size/2)
	dc.DrawRectangle(size/2-w/2, size/2-h/2, w, h)
	dc.Stroke()
	dc.DrawStringAnchored(text, size/2, size/2, 0.5, 0.5)
	dc.SavePNG(outFile)
}
