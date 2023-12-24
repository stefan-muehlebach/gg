package main

import (
	"github.com/stefan-muehlebach/gg"
	"github.com/stefan-muehlebach/gg/color"
	"github.com/stefan-muehlebach/gg/geom"
	"math/rand"
)

const (
	Width, Height = 512.0, 512.0
	MarginSize    = 24.0
	NumLines      = 35
	NumPoints     = 15
	SpaceBetween  = (Height - 2*MarginSize) / float64(NumLines-1)
	LineWidth     = 1.5
	xMin, xMax    = 0.1, 0.9
	yRange        = 1.5
)

var (
	BackColor = color.RGBAF{0.851, 0.811, 0.733, 1.0}
	LineColor = color.RGBAF{0.153, 0.157, 0.133, 1.0}
)

func CalcRelief(numPts int, x0, x1, yrange float64) []geom.Point {
	ptList := make([]geom.Point, numPts)
	ptList[0] = geom.Point{x0, 0.0}
	ptList[len(ptList)-1] = geom.Point{x1, 0.0}
	x := x0
	for i := 1; i < numPts-1; i++ {
		dx := rand.NormFloat64()*0.01 + 0.053
		y := rand.Float64()*yrange - yrange/2.0
		x += dx
		if x > x1 {
			x = x1
		}
		p := geom.Point{x, y}
		ptList[i] = p
		if x == x1 {
			break
		}
	}
	return ptList
}

func main() {
	var p1, p2 geom.Point
	var ptList []geom.Point

	gc := gg.NewContext(Width, Height)
	gc.SetFillColor(BackColor)
	gc.Clear()

	gc.SetStrokeWidth(LineWidth)
	gc.SetStrokeColor(LineColor)
	for i := 0; i < NumLines; i++ {
		p1 = geom.NewPoint(MarginSize, MarginSize+float64(i)*SpaceBetween)
		p2 = p1.AddXY(Width-2*MarginSize, 0.0)
		ptList = CalcRelief(NumPoints, xMin, xMax, yRange)
		gc.MoveTo(p1.X, p1.Y)
		for _, p := range ptList {
			pn := p1.Interpolate(p2, p.X)
			pn = pn.AddXY(0.0, p.Y*SpaceBetween)
			gc.LineTo(pn.X, pn.Y)
		}
		gc.LineTo(p2.X, p2.Y)
		gc.Stroke()
	}
	gc.SavePNG("relief.png")
}
