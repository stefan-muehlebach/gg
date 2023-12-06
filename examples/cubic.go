package main

import (
	"github.com/stefan-muehlebach/gg"
	"github.com/stefan-muehlebach/gg/color"
)

const (
	outFile = "cubic.png"
)

func main() {
	const S = 512
	dc := gg.NewContext(S, S)
	dc.SetFillColor(color.White)
	dc.Clear()
	dc.Translate(S/2, S/2)
	dc.Scale(20, 20)

	var x0, y0, x1, y1, x2, y2, x3, y3 float64
	x0, y0 = -10,  0
	x1, y1 =  -8, -8
	x2, y2 =   8,  8
	x3, y3 =  10,  0

	dc.MoveTo(x0, y0)
	dc.CubicTo(x1, y1, x2, y2, x3, y3)
	dc.SetStrokeWidth(8)
	dc.SetDash(16, 24)
	dc.SetFillColor(color.RGBAF{0, 0, 0, 0.2})
	dc.SetStrokeColor(color.Black)
	dc.FillStroke()

	dc.MoveTo(x0, y0)
	dc.LineTo(x1, y1)
	dc.LineTo(x2, y2)
	dc.LineTo(x3, y3)
	dc.SetStrokeWidth(2)
	dc.SetDash(4, 8, 1, 8)
	dc.SetStrokeColor(color.RGBAF{1, 0, 0, 0.4})
	dc.Stroke()

	dc.SavePNG(outFile)
}
