package main

import (
	"github.com/stefan-muehlebach/gg"
	"image/color"
)

const (
	outFile = "quadratic.png"
)

func main() {
	const S = 1000
	dc := gg.NewContext(S, S)
	dc.SetFillColor(color.White)
	dc.Clear()
	dc.Translate(S/2, S/2)
	dc.Scale(40, 40)

	var x0, y0, x1, y1, x2, y2, x3, y3, x4, y4 float64
	x0, y0 = -10, 0
	x1, y1 = -5, -10
	x2, y2 = 0, 0
	x3, y3 = 5, 10
	x4, y4 = 10, 0

	dc.MoveTo(x0, y0)
	dc.LineTo(x1, y1)
	dc.LineTo(x2, y2)
	dc.LineTo(x3, y3)
	dc.LineTo(x4, y4)
	dc.SetStrokeColor(color.RGBA{0xFF, 0x2D, 0x00, 0xff})
	dc.SetLineWidth(8)
	dc.Stroke()

	dc.MoveTo(x0, y0)
	dc.QuadraticTo(x1, y1, x2, y2)
	dc.QuadraticTo(x3, y3, x4, y4)
	dc.SetFillColor(color.RGBA{0x3E, 0x60, 0x6F, 0xff})
	dc.SetLineWidth(16)
	dc.SetStrokeColor(color.Black)
	dc.FillStroke()

	dc.DrawCircle(x0, y0, 0.5)
	dc.DrawCircle(x1, y1, 0.5)
	dc.DrawCircle(x2, y2, 0.5)
	dc.DrawCircle(x3, y3, 0.5)
	dc.DrawCircle(x4, y4, 0.5)
	dc.SetFillColor(color.White)
	dc.SetStrokeColor(color.Black)
	dc.SetLineWidth(4)
	dc.FillStroke()

	dc.LoadFontFace("Ubuntu-Regular.ttf", 20)
	dc.DrawStringAnchored("g", -5, 5, 0.5, 0.5)
	dc.DrawStringAnchored("G", 5, -5, 0.5, 0.5)

	dc.SavePNG(outFile)
}
