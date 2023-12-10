package main

import (
	"github.com/stefan-muehlebach/gg"
	"github.com/stefan-muehlebach/gg/color"
	"math/rand"
)

const (
	outFile = "lines.png"
)

func main() {
	const W = 1024
	const H = 1024
	dc := gg.NewContext(W, H)
	dc.SetFillColor(color.Black)
	dc.Clear()
	for i := 0; i < 1000; i++ {
		x1 := rand.Float64() * W
		y1 := rand.Float64() * H
		x2 := rand.Float64() * W
		y2 := rand.Float64() * H
		r  := rand.Float64()
		g  := rand.Float64()
		b  := rand.Float64()
		a  := 0.5 + 0.5*rand.Float64()
		w  := rand.Float64()*4 + 1
		dc.SetStrokeColor(color.RGBAF{r, g, b, a})
		dc.SetStrokeWidth(w)
		dc.DrawLine(x1, y1, x2, y2)
		dc.Stroke()
	}
	dc.SavePNG(outFile)
}
