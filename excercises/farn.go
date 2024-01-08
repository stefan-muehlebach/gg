package main

import (
	"github.com/stefan-muehlebach/gg/colornames"
	"github.com/stefan-muehlebach/gg"
	"github.com/stefan-muehlebach/gg/color"
	"github.com/stefan-muehlebach/gg/geom"
	"math/rand"
)

const (
	Width, Height = 512.0, 512.0
	MarginSize    = 24.0
	GraphSize     = Width - 2*MarginSize
	OutFileName   = "farn.png"

	xMin, xMax = -5.5,  5.5
	yMin, yMax =  0.0, 11.0

	numPoints = 100_000
)

var (
	BackColor = color.RGBAF{0.851, 0.811, 0.733, 1.0}
	LineColor = color.RGBAF{0.153, 0.157, 0.133, 1.0}

	p = []float64{0.01, 0.86, 0.93, 1.0}

	m = []geom.Matrix{
		{0.0, 0.0, 0.0, 0.0, 0.16, 0.0},
		{0.85, 0.04, 0.0, -0.04, 0.85, 1.6},
		{0.2, -0.26, 0.0, 0.23, 0.22, 1.6},
		{-0.15, 0.28, 0.0, 0.26, 0.24, 0.44},
	}
)

func colToX(col int) float64 {
	return xMin + (float64(col)/float64(GraphSize))*(xMax-xMin)
}
func rowToY(row int) float64 {
	return yMax - (float64(row)/float64(GraphSize))*(yMax-yMin)
}
func xToCol(x float64) float64 {
	return float64(GraphSize) * (x - xMin) / (xMax - xMin)
}
func yToRow(y float64) float64 {
	return float64(GraphSize) * (yMax - y) / (yMax - yMin)
}

func main() {
	gc := gg.NewContext(Width, Height)
	gc.SetFillColor(BackColor)
	gc.Clear()
	gc.SetFillColor(colornames.DarkGreen)
	pt := geom.Point{0.0, 0.0}
	for i := 0; i < numPoints; i++ {
		rnd := rand.Float64()
		pos := 0
		for rnd >= p[pos] {
			pos += 1
		}
		pt = m[pos].Transform(pt)
		gc.DrawPoint(xToCol(pt.X), yToRow(pt.Y), 0.5)
		gc.Fill()

	}
	gc.SavePNG(OutFileName)
}
