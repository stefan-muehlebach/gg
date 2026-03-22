package main

import (
	"github.com/stefan-muehlebach/gg"
	"github.com/stefan-muehlebach/gg/colors"
	"github.com/stefan-muehlebach/gg/fonts"
)

const (
    W = 800
    H = 400
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
	dc := gg.NewContext(W, H)
    face, _ := fonts.NewFace(fonts.GoRegular, 18)

	dc.SetFillColor(colors.White)
	dc.Clear()
	dc.SetStrokeColor(colors.Black)
	dc.SetFontFace(face)
	const h = 24
	for i, line := range lines {
		y := H/2 - h*len(lines)/2 + i*h
		dc.DrawStringAnchored(line, 400, float64(y), 0.5, 0.5)
	}
	dc.SavePNG(outFile)
}
