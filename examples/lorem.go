package main

import (
	"github.com/stefan-muehlebach/gg"
	"image/color"
)

const (
	outFile = "lorem.png"
)

var lines = []string{
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod",
	"tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam,",
	"quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo",
	"consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse",
	"cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat",
	"non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
}

func main() {
	const W = 800
	const H = 400
	dc := gg.NewContext(W, H)
	dc.SetFillColor(color.White)
	dc.Clear()
	dc.SetStrokeColor(color.Black)
	dc.LoadFontFace("Ubuntu-Regular.ttf", 18)
	const h = 24
	for i, line := range lines {
		y := H/2 - h*len(lines)/2 + i*h
		dc.DrawStringAnchored(line, 400, float64(y), 0.5, 0.5)
	}
	dc.SavePNG(outFile)
}
