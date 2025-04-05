package main

import (
	"github.com/stefan-muehlebach/gg"
	"github.com/stefan-muehlebach/gg/colors"
	"github.com/stefan-muehlebach/gg/fonts"
)

const (
	width, height = 512, 512
	outFile       = "gofont.png"
)

func main() {
	face := fonts.NewFace(fonts.GoBold, 48)

	dc := gg.NewContext(width, height)
	dc.SetFontFace(face)
	dc.SetFillColor(colors.White)
	dc.Clear()
	dc.SetStrokeColor(colors.Black)
	dc.DrawStringAnchored("Hello, world!", width/2, height/2, 0.5, 0.5)
	dc.SavePNG(outFile)
}
