package main

import (
	"math"

	"github.com/stefan-muehlebach/gg"
	"github.com/stefan-muehlebach/gg/colors"
)

const (
	outFile = "star.png"
)

type Point struct {
	X, Y float64
}

func Polygon(n int, x, y, r float64) []Point {
	result := make([]Point, n)
	for i := 0; i < n; i++ {
		a := float64(i)*2*math.Pi/float64(n) - math.Pi/2
		result[i] = Point{x + r*math.Cos(a), y + r*math.Sin(a)}
	}
	return result
}

func main() {
	n := 5
	points := Polygon(n, 512, 512, 400)
	dc := gg.NewContext(1024, 1024)
	dc.SetFillColor(colors.White)
	dc.Clear()
	for i := 0; i < n+1; i++ {
		index := (i * 2) % n
		p := points[index]
		dc.LineTo(p.X, p.Y)
	}
	dc.SetFillColor(colors.RGBAF{0, 0.5, 0, 1.0})
	dc.SetFillRule(gg.FillRuleEvenOdd)
	dc.SetStrokeColor(colors.RGBAF{0, 1.0, 0, 0.5})
	dc.SetStrokeWidth(16)
	dc.FillStroke()
	dc.SavePNG(outFile)
}
