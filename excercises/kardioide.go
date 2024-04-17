package main

import (
	"math"

	"github.com/stefan-muehlebach/gg"
	"github.com/stefan-muehlebach/gg/color"
	"github.com/stefan-muehlebach/gg/colornames"
	"github.com/stefan-muehlebach/gg/geom"
)

const (
	// Breite und Höhe des Bildes in Pixel
	Width, Height = 512.0, 512.0
	// Freizulassender Rand
	MarginSize = 20.0

	NumPoints = 360

	OutFileName = "kardioide.png"
)

var (
	BackColor = color.RGBAF{0.851, 0.811, 0.733, 1.0}
	LineColor = colornames.DarkSlateBlue.Dark(0.4)
    LineWidth = 1.0
)

func CircularPointList(midPoint geom.Point, radius, startAngle, endAngle float64, n int) []geom.Point {
	divider := n - 1
	if (startAngle == 0) && (endAngle == 2*math.Pi) {
		divider = n
	}
	pl := make([]geom.Point, n)
	for i := 0; i < n; i++ {
		t := float64(i) / float64(divider)
		angle := (1-t)*startAngle + t*endAngle
		pl[i] = midPoint.Add(geom.Point{math.Cos(angle), math.Sin(angle)}.Mul(radius))
	}
	return pl
}

func Cardioide(gc *gg.Context, pl []geom.Point, f int) {
	gc.SetStrokeWidth(LineWidth)
	gc.SetStrokeColor(LineColor)
    for i, p1 := range pl[1:] {
	// for i := 1; i < len(pl); i++ {
		// p1 := pl[i]
		j := (f * (i+1)) % (len(pl))
		p2 := pl[j]
		gc.DrawLine(p1.X, p1.Y, p2.X, p2.Y)
	}
    	gc.Stroke()
}

func main() {
	var pointList []geom.Point

	pointList = make([]geom.Point, 0)
	pointList = append(pointList, CircularPointList(geom.Point{Width/2, Height/2}, Width/2-MarginSize, 0, 2*math.Pi, NumPoints)...)

	gc := gg.NewContext(Width, Height)
	gc.SetFillColor(BackColor)
	gc.Clear()

	Cardioide(gc, pointList, 2)

	gc.SavePNG(OutFileName)
}
