package main

import "github.com/stefan-muehlebach/gg"

const (
	outFile = "pattern-fill.png"
)

func main() {
	im, err := gg.LoadPNG("baboon.png")
	if err != nil {
		panic(err)
	}
	pattern := gg.NewSurfacePattern(im, gg.RepeatBoth)
	dc := gg.NewContext(600, 600)
	dc.MoveTo(20, 20)
	dc.LineTo(590, 20)
	dc.LineTo(590, 590)
	dc.LineTo(20, 590)
	dc.ClosePath()
	dc.SetFillStyle(pattern)
	dc.Fill()
	dc.SavePNG(outFile)
}
