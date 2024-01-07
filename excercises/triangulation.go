package main

import (
	"github.com/stefan-muehlebach/gg"
	"github.com/stefan-muehlebach/gg/color"
	"github.com/stefan-muehlebach/gg/colornames"
	"github.com/stefan-muehlebach/gg/fonts"
	"github.com/stefan-muehlebach/gg/geom"

	// "image/color"
	"math/rand"
)

const (
	Width, Height = 512.0, 512.0
	Padding       = 24
	OutFileName   = "triangulation.png"
	NumPoints     = 30
	TitleFontSize = 12.0
	TextFontSize  = 10.0
)

var (
	BackColor = color.RGBAF{0.851, 0.811, 0.733, 1.0}
	LineColor = color.RGBAF{0.153, 0.157, 0.133, 1.0}
	TitleFont = fonts.GoBold
	TextFont  = fonts.GoRegular
)

type PointList struct {
    points []geom.Point
    sortOnX, sortOnY []int
}

type Edge struct {
    from, to int
}

func CreatePoints(n int) []geom.Point {
	points := make([]geom.Point, n)
	for i := 0; i < n; i++ {
		x := 2.0*rand.Float64() - 1.0
		y := 2.0*rand.Float64() - 1.0
		points[i] = geom.Point{x, y}
	}
	return points
}

func CreateEdges(points []geom.Point) []Edge {
    edges := make([]Edge, 0)
    for i := range points {
        from := i
        to := (i+1) % len(points)
         e := Edge{from, to}
        edges = append(edges, e)
    }
    return edges
}

func main() {
	dc := gg.NewContext(Width, Height)
	dc.SetFillColor(BackColor)
	dc.Clear()

	dc.Translate(Padding, Height-Padding)
	dc.Scale((Width-2*Padding)/2.0, -(Height-2*Padding)/2.0)
	dc.Translate(1, 1)

	points := CreatePoints(NumPoints)
    edges := CreateEdges(points)
    
    dc.SetStrokeColor(colornames.Silver)
    dc.SetStrokeWidth(1.0)
    dc.DrawLine(-1, 1, 1, -1)
    dc.DrawLine(-1, -1, 1, 1)
    dc.DrawRectangle(-1, -1, 2, 2)
    dc.DrawCircle(0, 0, 1)
    dc.Stroke()

	// draw points
	dc.SetStrokeColor(colornames.Crimson)
	dc.SetFillColor(colornames.Crimson)
	for _, p := range points {
		dc.DrawPoint(p.X, p.Y, 1.5)
		dc.FillStroke()
	}
    
    dc.SetStrokeColor(color.Black.Alpha(0.3))
    dc.SetStrokeWidth(2.0)
    for _, e := range edges {
        x0, y0 := points[e.from].AsCoord()
        x1, y1 := points[e.to].AsCoord()
        dc.DrawLine(x0, y0, x1, y1)
        dc.Stroke()
    }

	dc.SavePNG(OutFileName)
}
