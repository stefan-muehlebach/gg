package main

import (
	"github.com/stefan-muehlebach/gg"
	"github.com/stefan-muehlebach/gg/colors"
)

const (
	outFile = "linewidth.png"
)

func main() {
	dc := gg.NewContext(1000, 1000)
	dc.SetFillColor(colors.White)
	dc.Clear()
	dc.SetStrokeColor(colors.Black)
	w := 0.1
	for i := 100; i <= 900; i += 20 {
		x := float64(i)
		dc.DrawLine(x+50, 0, x-50, 1000)
		dc.SetStrokeWidth(w)
		dc.Stroke()
		w += 0.1
	}
	dc.SavePNG(outFile)
}
