package main

import (
	"github.com/stefan-muehlebach/gg"
	"image/color"
)

const (
	outFile = "text.png"
)

func main() {
	const S = 1024
	dc := gg.NewContext(S, S)
	dc.SetFillColor(color.White)
	dc.Clear()
	dc.SetStrokeColor(color.Black)
	if err := dc.LoadFontFace("Ubuntu-Regular.ttf", 96); err != nil {
		panic(err)
	}
	dc.DrawStringAnchored("Hello, world!", S/2, S/2, 0.5, 0.5)
	dc.SavePNG(outFile)
}
