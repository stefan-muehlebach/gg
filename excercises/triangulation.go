package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"

	"github.com/stefan-muehlebach/gg"
	"github.com/stefan-muehlebach/gg/color"
	"github.com/stefan-muehlebach/gg/fonts"
	"github.com/stefan-muehlebach/gg/geom"
	// "image/color"
)

const (
	ImageSize    = 1024.0
	Margin       = 20
	NumPoints    = 60
	MinPointDist = 0.2
	MaxPointDist = 1.0
	RandomSeed   = 2_345_678_900
	DrawLabels   = true
	OutFileName  = "triangulation.png"
)

var (
	BackColor = color.RGBAF{0.851, 0.811, 0.733, 1.0}

	AxesColor = color.Silver

	PointSize         = 3.5
	PointColor        = color.FireBrick
	PointTextFont     = fonts.GoBold
	PointTextColor    = PointColor.Dark(0.5)
	PointTextFontSize = 16.0

	EdgeSize         = 1.0
	EdgeColor        = color.SlateGray
	EdgeTextFont     = fonts.GoRegular
	EdgeTextColor    = EdgeColor.Dark(0.5)
	EdgeTextFontSize = 0.8 * PointTextFontSize
)

// Neue Superstruktur fuer Punkte, Kanten und Dreiecke.
type Grid struct {
	points    []geom.Point
	edges     []Edge
	triangles []Triangle
}

func NewGrid() *Grid {
	g := &Grid{}
	g.points = make([]geom.Point, 0)
	g.edges = make([]Edge, 0)
	g.triangles = make([]Triangle, 0)
	return g
}

func (g *Grid) AddPoint(p geom.Point) (idx int) {
	idx = len(g.points)
	g.points = append(g.points, p)
	return
}

func (g *Grid) MinDist(p geom.Point) float64 {
	var minDist float64 = math.MaxFloat64

	for _, q := range g.points {
		d := p.Distance(q)
		if d < minDist {
			minDist = d
		}
	}
	return minDist
}

func (g *Grid) AddEdge(pi0, pi1 int, length float64) (idx int) {
	idx = len(g.edges)
	if length < 0.0 {
		length = g.points[pi0].Distance(g.points[pi1])
	}
	g.edges = append(g.edges, Edge{pi0, pi1, length})
	return idx
}

func (g *Grid) GetEdges(pi0 int) []Edge {
	edges := make([]Edge, 0)
	for _, edge := range g.edges {
		if edge.pi0 == pi0 {
			edges = append(edges, edge)
		}
	}
	return edges
}

func (g *Grid) FindEdge(pi0, pi1 int) (bool, int) {
	for i, edge := range g.edges {
		if (edge.pi0 == pi0 && edge.pi1 == pi1) ||
			(edge.pi0 == pi1 && edge.pi1 == pi0) {
			return true, i
		}
	}
	return false, -1
}

func (g *Grid) AllEdgesByLength() []Edge {
	allEdges := make([]Edge, 0)
	for pi0 := 0; pi0 < len(g.points); pi0++ {
		for pi1 := pi0 + 1; pi1 < len(g.points); pi1++ {
			l := g.points[pi0].Distance(g.points[pi1])
			allEdges = append(allEdges, Edge{pi0, pi1, l})
		}
	}
	sort.SliceStable(allEdges, func(i, j int) bool {
		return allEdges[i].length < allEdges[j].length
	})
	return allEdges
}

func (g *Grid) Angle(edge Edge) float64 {
	v := g.points[edge.pi1].Sub(g.points[edge.pi0])
	if v.Y < 0.0 {
		return -math.Acos(v.X / edge.length)
	} else {
		return math.Acos(v.X / edge.length)
	}
}

func (g *Grid) Intersect(piA, piB int) (bool, int) {
	for i, edge := range g.edges {
		if g.isBetween(edge.pi0, edge.pi1, piA, piB) && g.isBetween(piA, piB, edge.pi0, edge.pi1) {
			return true, i
		}
	}
	return false, -1
}

func (g *Grid) isBetween(pi0, pi1, piA, piB int) bool {
	p0 := g.points[pi0]
	p1 := g.points[pi1]
	pA := g.points[piA]
	pB := g.points[piB]
	vE := p1.Sub(p0)
	vA := pA.Sub(p0)
	vB := pB.Sub(p0)
	detA := vE.X*vA.Y - vA.X*vE.Y
	detB := vE.X*vB.Y - vB.X*vE.Y
	if detA*detB < 0 {
		return true
	} else {
		return false
	}
}

func (g *Grid) Draw(gc *gg.Context, withLabels bool) {
	gc.SetStrokeColor(EdgeColor)
	gc.SetStrokeWidth(EdgeSize)
	for _, edge := range g.edges {
		p0 := g.points[edge.pi0]
		p1 := g.points[edge.pi1]
		gc.DrawLine(p0.X, p0.Y, p1.X, p1.Y)
		gc.Stroke()
	}

	gc.SetStrokeColor(PointColor)
	gc.SetFillColor(PointColor)
	for _, pt := range g.points {
		gc.DrawPoint(pt.X, pt.Y, PointSize)
		gc.FillStroke()
	}

	if !withLabels {
		return
	}
	matrix := gc.Matrix()
	gc.Push()
	gc.Identity()

	face := fonts.NewFace(EdgeTextFont, EdgeTextFontSize)
	gc.SetFontFace(face)
	gc.SetStrokeColor(EdgeTextColor)
	gc.SetFillColor(BackColor)
	for i, edge := range g.edges {
		p0 := g.points[edge.pi0]
		p1 := g.points[edge.pi1]
		pt := p0.Interpolate(p1, 0.5)
		ptNew := matrix.Transform(pt)
		lbl := fmt.Sprintf("%d", i)
		lblW, lblH := gc.MeasureString(lbl)
		gc.DrawCircle(ptNew.X, ptNew.Y, max(lblW/2, lblH/2))
		gc.Fill()
		gc.DrawStringAnchored(lbl, ptNew.X, ptNew.Y, 0.5, 0.5)
	}
	face = fonts.NewFace(PointTextFont, PointTextFontSize)
	gc.SetFontFace(face)
	gc.SetStrokeColor(PointTextColor)
	for i, pt := range g.points {
		ptNew := matrix.Transform(pt)
		gc.DrawString(fmt.Sprintf("%d", i), ptNew.X+4, ptNew.Y-4)
	}

	gc.Pop()
}

// Typ fuer ein Dreieck im Graph.
type Triangle struct {
	p0i, p1i, p2i int
}

// Typ fuer eine Kante im Graph.
type Edge struct {
	pi0, pi1 int
	length   float64
}

//----------------------------------------------------------------------------

// Zeichnet ein Achsenkreuz (oder so aehnlich).
func DrawAxes(gc *gg.Context) {
	gc.SetStrokeColor(AxesColor)
	gc.SetStrokeWidth(1.0)
	gc.DrawLine(-1, 1, 1, -1)
	gc.DrawLine(-1, -1, 1, 1)
	gc.DrawLine(-1, 0, 1, 0)
	gc.DrawLine(0, -1, 0, 1)
	gc.DrawRectangle(-1, -1, 2, 2)
	gc.DrawCircle(0, 0, 1)
	gc.Stroke()
}

// Hauptfunktion.
func main() {
	rand.Seed(RandomSeed)

	gc := gg.NewContext(ImageSize, ImageSize)
	gc.SetFillColor(BackColor)
	gc.Clear()

	gc.Translate(Margin, ImageSize-Margin)
	gc.Scale((ImageSize-2*Margin)/2.0, -(ImageSize-2*Margin)/2.0)
	gc.Translate(1, 1)

	grid := NewGrid()
	for i := 0; i < NumPoints; i++ {
		for {
			x := 2.0*rand.Float64() - 1.0
			y := 2.0*rand.Float64() - 1.0
			p := geom.Point{x, y}
			if grid.MinDist(p) < MinPointDist {
				continue
			}
			grid.AddPoint(p)
			break
		}
	}

	for _, edge := range grid.AllEdgesByLength() {
		pi0, pi1 := edge.pi0, edge.pi1
		if edge.length > MaxPointDist {
			continue
		}
		// fmt.Printf("trying to add an edge for [%d] and [%d]\n", from, to)
		if nok, _ := grid.Intersect(pi0, pi1); nok {
			// fmt.Printf("  > intersection with edge <%d>\n", otherEdge)
			continue
		}
		grid.AddEdge(pi0, pi1, edge.length)
		// fmt.Printf("  > add edge <%d> (from [%d] to [%d])\n", idx, from, to)
	}

	pi0 := 9
	edges := grid.GetEdges(pi0)
	sort.SliceStable(edges, func(i, j int) bool {
		return grid.Angle(edges[i]) < grid.Angle(edges[j])
	})
	fmt.Printf("edges, originating from point %d: %v\n", pi0, edges)
	for i := 0; i < len(edges); i++ {
		pi1 := edges[i].pi1
		pi2 := edges[(i+1)%len(edges)].pi1
		if grid.FindEdge(pi1, pi2) {
			fmt.Printf(">>> triangle %d, %d, %d\n", pi0, pi1, pi2)
		}
	}

	DrawAxes(gc)

	grid.Draw(gc, DrawLabels)

	gc.SavePNG(OutFileName)
}
