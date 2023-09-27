package main

import (
    "image/color"
	"math/rand"
	"mju.net/gg"
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
		r := uint8(rand.Intn(255))
		g := uint8(rand.Intn(255))
		b := uint8(rand.Intn(255))
		a := uint8(127 + rand.Intn(127))
		w := rand.Float64()*4 + 1
		dc.SetStrokeColor(color.NRGBA{r, g, b, a})
		dc.SetLineWidth(w)
		dc.DrawLine(x1, y1, x2, y2)
		dc.Stroke()
	}
	dc.SavePNG(outFile)
}
