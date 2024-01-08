package main

import (
	"fmt"
	"math"

	"github.com/stefan-muehlebach/gg"
	"github.com/stefan-muehlebach/gg/color"
	"github.com/stefan-muehlebach/gg/colornames"
	"github.com/stefan-muehlebach/gg/fonts"
	"github.com/stefan-muehlebach/gg/geom"

	// "image/color"
	"math/rand"
)

const (
	ImageSize     = 1024.0
	Padding       = 20
	NumPoints     = 40
    NumNeigbours  = 3
	TextFontSize  = 16.0
	OutFileName   = "triangulation.png"
)

var (
	BackColor = color.RGBAF{0.851, 0.811, 0.733, 1.0}
	LineColor = color.RGBAF{0.153, 0.157, 0.133, 1.0}
	TextFont  = fonts.GoRegular
    TextColor = colornames.DarkSlateGray
)

type PointList struct {
	points []geom.Point
}

func NewPointList() *PointList {
	pl := &PointList{}
	pl.points = make([]geom.Point, 0)
	return pl
}

func (pl *PointList) AddPoint(x, y float64) {
	pl.points = append(pl.points, geom.Point{x, y})
}

func (pl *PointList) FindNearest(idxA int, lowerDist float64) (int, float64) {
	var idxB int = -1
	var minDist float64 = math.MaxFloat64

	ptA := pl.points[idxA]
	for j, ptB := range pl.points {
		if j == idxA {
			continue
		}
		d := ptA.Distance(ptB)
		if d <= lowerDist {
			continue
		}
		if d < minDist {
			idxB = j
			minDist = d
		}
	}
	return idxB, minDist
}

func (pl *PointList) Draw(gc *gg.Context) {
    face := fonts.NewFace(TextFont, TextFontSize)
    matrix := gc.Matrix()

    gc.SetFontFace(face)
	gc.SetStrokeColor(colornames.Crimson)
	gc.SetFillColor(colornames.Crimson)
	for _, pt := range pl.points {
		gc.DrawPoint(pt.X, pt.Y, 1.5)
		gc.FillStroke()
    }
    gc.Push()
    gc.Identity()
    gc.SetStrokeColor(TextColor)
	for i, pt := range pl.points {
        ptNew := matrix.Transform(pt)
        // fmt.Printf("%v > %v\n", pt, ptNew)
        gc.DrawString(fmt.Sprintf("%d", i), ptNew.X+4, ptNew.Y-4)
	}
    gc.Pop()
}

type Edge struct {
	from, to int
}

type EdgeList struct {
	edges []Edge
	pl    *PointList
}

func NewEdgeList(pl *PointList) *EdgeList {
	el := &EdgeList{}
	el.edges = make([]Edge, 0)
	el.pl = pl
	return el
}

func (el *EdgeList) AddEdge(from, to int) {
	el.edges = append(el.edges, Edge{from, to})
}

func (el *EdgeList) Length(edge int) (float64) {
    p0 := el.pl.points[el.edges[edge].from]
    p1 := el.pl.points[el.edges[edge].to]
    return p0.Distance(p1)
}

func (el *EdgeList) Search(from, to int) bool {
	for _, edge := range el.edges {
		if edge.from == from && edge.to == to {
			return true
		}
	}
	return false
}

func (el *EdgeList) Draw(gc *gg.Context) {
	gc.SetStrokeColor(LineColor.Alpha(0.3))
	gc.SetStrokeWidth(2.0)
	for _, e := range el.edges {
		p0 := el.pl.points[e.from]
		p1 := el.pl.points[e.to]
		gc.DrawLine(p0.X, p0.Y, p1.X, p1.Y)
		gc.Stroke()
	}
}

func DrawOrientation(gc *gg.Context) {
	gc.SetStrokeColor(colornames.Silver)
	gc.SetStrokeWidth(1.0)
	gc.DrawLine(-1, 1, 1, -1)
	gc.DrawLine(-1, -1, 1, 1)
	gc.DrawRectangle(-1, -1, 2, 2)
	gc.DrawCircle(0, 0, 1)
	gc.Stroke()
}

func main() {
	rand.Seed(321_654_000)
	dc := gg.NewContext(ImageSize, ImageSize)
	dc.SetFillColor(BackColor)
	dc.Clear()

	dc.Translate(Padding, ImageSize-Padding)
	dc.Scale((ImageSize-2*Padding)/2.0, -(ImageSize-2*Padding)/2.0)
	dc.Translate(1, 1)

	pointList := NewPointList()
	for i := 0; i < NumPoints; i++ {
		x := 2.0*rand.Float64() - 1.0
		y := 2.0*rand.Float64() - 1.0
		pointList.AddPoint(x, y)
	}
	edgeList := NewEdgeList(pointList)

	for from := range pointList.points {
		dist := 0.0
		to := -1
		for i := 0; i < NumNeigbours; i++ {
            // fmt.Printf("%d: finding neighbour %d with more then %f distance\n",
                    // from, i, dist)
			if to, dist = pointList.FindNearest(from, dist); to < 0 {
                // fmt.Printf("  > break with %d, %f\n", to, dist)
				break
			}
            // fmt.Printf("  > %d found with a distance of %f\n", to, dist)
			if edgeList.Search(to, from) {
				continue
			}
			edgeList.AddEdge(from, to)
		}
	}

	DrawOrientation(dc)

	edgeList.Draw(dc)
	pointList.Draw(dc)

	// fmt.Printf("edges: %v\n", edgeList.edges)

	dc.SavePNG(OutFileName)
}
