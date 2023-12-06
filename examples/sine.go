package main

import (
	"github.com/stefan-muehlebach/gg"
	"image/color"
	"math"
)

const (
	outFile = "sine.png"
)

func main() {
	const W = 1200
	const H = 60
	dc := gg.NewContext(W, H)
	// dc.SetHexColor("#FFFFFF")
	// dc.Clear()
	dc.ScaleAbout(0.95, 0.75, W/2, H/2)
	for i := 0; i < W; i++ {
		a := float64(i) * 2 * math.Pi / W * 8
		x := float64(i)
		y := (math.Sin(a) + 1) / 2 * H
		dc.LineTo(x, y)
	}
	dc.ClosePath()
	dc.SetFillColor(color.RGBA{0x3e, 0x60, 0x6f, 0xff})
	dc.SetStrokeColor(color.RGBA{0x19, 0x34, 0x41, 0x80})
	dc.SetStrokeWidth(8)
	dc.FillStroke()
	dc.SavePNG(outFile)
    
}
