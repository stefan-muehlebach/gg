package main

import (
    "image/color"
    "github.com/stefan-muehlebach/gg"
)

const (
    outFile = "linewidth.png"
)

func main() {
	dc := gg.NewContext(1000, 1000)
	dc.SetFillColor(color.White)
	dc.Clear()
	dc.SetStrokeColor(color.Black)
	w := 0.1
	for i := 100; i <= 900; i += 20 {
		x := float64(i)
		dc.DrawLine(x+50, 0, x-50, 1000)
		dc.SetLineWidth(w)
		dc.Stroke()
		w += 0.1
	}
	dc.SavePNG(outFile)
}
