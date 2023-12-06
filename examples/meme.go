package main

import (
	"github.com/stefan-muehlebach/gg"
	"github.com/stefan-muehlebach/gg/color"
	"github.com/stefan-muehlebach/gg/fonts"
)

const (
	outFile = "meme.png"
)

func main() {
	const S = 1024
	dc := gg.NewContext(S, S)
	dc.SetFillColor(color.White)
	dc.Clear()
	dc.SetFontFace(fonts.NewFace(fonts.GoBold, 72))
	dc.SetStrokeColor(color.Black)
	s := "ONE DOES NOT SIMPLY"
	n := 6 // "stroke" size
	for dy := -n; dy <= n; dy++ {
		for dx := -n; dx <= n; dx++ {
			if dx*dx+dy*dy >= n*n {
				// give it rounded corners
				continue
			}
			x := S/2 + float64(dx)
			y := S/2 + float64(dy)
			dc.DrawStringAnchored(s, x, y, 0.5, 0.5)
		}
	}
	dc.SetStrokeColor(color.White)
	dc.DrawStringAnchored(s, S/2, S/2, 0.5, 0.5)
	dc.SavePNG(outFile)
}
