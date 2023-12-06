package main

import (
	"github.com/stefan-muehlebach/gg"
	"github.com/stefan-muehlebach/gg/color"
	"github.com/stefan-muehlebach/gg/fonts"
)

const (
    width, height = 512, 512
	outFile = "gofont.png"
)

func main() {
	face := fonts.NewFace(fonts.GoBold, 48)

	dc := gg.NewContext(width, height)
	dc.SetFontFace(face)
	dc.SetFillColor(color.White)
	dc.Clear()
	dc.SetStrokeColor(color.Black)
	dc.DrawStringAnchored("Hello, world!", width/2, height/2, 0.5, 0.5)
	dc.SavePNG(outFile)
}
