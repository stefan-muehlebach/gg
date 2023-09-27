package main

import (
    "image/color"
	"math/rand"
	"mju.net/gg"
)

const (
    outFile = "scatter.png"
)

func CreatePoints(n int) []gg.Point {
	points := make([]gg.Point, n)
	for i := 0; i < n; i++ {
		x := 0.5 + rand.NormFloat64()*0.1
		y := x + rand.NormFloat64()*0.1
		points[i] = gg.Point{x, y}
	}
	return points
}

func main() {
	const S = 1024
	const P = 64
	dc := gg.NewContext(S, S)
	dc.InvertY()
	dc.SetFillColor(color.White)
	dc.Clear()
	points := CreatePoints(1000)
	dc.Translate(P, P)
	dc.Scale(S-P*2, S-P*2)
	// draw minor grid
	for i := 1; i <= 10; i++ {
		x := float64(i) / 10
		dc.MoveTo(x, 0)
		dc.LineTo(x, 1)
		dc.MoveTo(0, x)
		dc.LineTo(1, x)
	}
	dc.SetStrokeColor(color.RGBA{0, 0, 0, 63})
	dc.SetLineWidth(1)
	dc.Stroke()
	// draw axes
	dc.MoveTo(0, 0)
	dc.LineTo(1, 0)
	dc.MoveTo(0, 0)
	dc.LineTo(0, 1)
	dc.SetStrokeColor(color.Black)
	dc.SetLineWidth(4)
	dc.Stroke()
	// draw points
	dc.SetFillColor(color.NRGBA{0, 0, 255, 127})
	for _, p := range points {
		dc.DrawCircle(p.X, p.Y, 3.0/S)
		dc.Fill()
	}
	// draw text
	dc.Identity()
	dc.SetStrokeColor(color.Black)
	if err := dc.LoadFontFace("Ubuntu-Bold.ttf", 24); err != nil {
		panic(err)
	}
	dc.DrawStringAnchored("Chart Title", S/2, P/2, 0.5, 0.5)
	if err := dc.LoadFontFace("Ubuntu-Regular.ttf", 18); err != nil {
		panic(err)
	}
	dc.DrawStringAnchored("X Axis Title", S/2, S-P/2, 0.5, 0.5)
	dc.SavePNG(outFile)
}
