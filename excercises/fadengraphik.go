package main

import (
	"math"
	// "fmt"
	"github.com/stefan-muehlebach/gg"
	"github.com/stefan-muehlebach/gg/color"
	"github.com/stefan-muehlebach/gg/colornames"

	// "github.com/stefan-muehlebach/gg/fonts"
	"github.com/stefan-muehlebach/gg/geom"
	// "math"
)

const (
	// Breite und Höhe des Bildes in Pixel
	Width, Height = 512.0, 512.0
	// Freizulassender Rand
	MarginSize = 20.0

	NumPoints = 360

	// Grösse der Punkte.
	PointSize = 3.0
	// Grösse der Beschriftungen. Hier muss man etwas experimentieren, bis
	// die ideale Einstellung gefunden ist.
	FontSize = 24.0

	OutFileName = "kardioide.png"
)

type ThreadData struct {
	width float64
	color color.Color
}

var (
	// BackColor und LineColor sind die Standardfarben für den Hintergrund
	// (das Papier) und den Vordergrund.
	// BackColor = color.RGBAF{0.851, 0.811, 0.733, 1.0}
	// BackColor = colornames.DarkRed.Dark(0.6)
	// LineColor = color.RGBAF{0.153, 0.157, 0.133, 1.0}

	BackColor = color.RGBAF{0.851, 0.811, 0.733, 1.0}
	LineColor = color.RGBAF{0.153, 0.157, 0.133, 1.0}

	ThreadList = []ThreadData{
		{1.0, colornames.DarkSlateBlue.Dark(0.4)},
		{1.0, colornames.DarkViolet.Bright(0.2)},
		{1.0, colornames.DarkBlue.Bright(0.1)},
		{1.0, colornames.DarkGreen},
		{1.0, colornames.DarkOliveGreen.Dark(0.1)},
		{1.0, colornames.DarkOrange.Dark(0.2)},
	}
)

// func DrawLines(gc *gg.Context, pl []geom.Point, data GraphData) {
// 	var idx0, idx1 int

// 	idx0 = 0
// 	gc.SetStrokeWidth(data.lineWidth)
// 	gc.SetStrokeColor(data.color)
// 	for {
// 		idx1 = (idx0 + data.stepSize) % len(pl)
// 		p0 := pl[idx0]
// 		p1 := pl[idx1]
// 		gc.DrawLine(p0.X, p0.Y, p1.X, p1.Y)
// 		gc.Stroke()
// 		idx0 = idx1
// 		if idx0 == 0 {
// 			break
// 		}
// 	}
// }

func LinearPointList(p1, p2 geom.Point, n int) []geom.Point {
	pl := make([]geom.Point, n)
	for i := 0; i < n; i++ {
		t := float64(i) / float64(n-1)
		pl[i] = p1.Interpolate(p2, t)
	}
	return pl
}

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

func BezierPointList(p1, p2, p3, p4 geom.Point, n int) []geom.Point {
	pl := make([]geom.Point, n)
	for i := 0; i < n; i++ {
		t := float64(i) / float64(n-1)
		t2 := t * t
		t3 := t2 * t
		u := 1.0 - t
		u2 := u * u
		u3 := u2 * u
		pl[i] = p1.Mul(u3).Add(p2.Mul(3 * u2 * t)).Add(p3.Mul(3 * u * t2)).Add(p4.Mul(t3))
	}
	return pl
}

func DrawPointList(gc *gg.Context, pl []geom.Point) {
	gc.SetFillColor(LineColor)
	for _, p := range pl {
		gc.DrawPoint(p.X, p.Y, PointSize)
		gc.Fill()
	}
}

func Weave(gc *gg.Context, pl1, pl2 []geom.Point, thread ThreadData, cont bool, reversed bool) {
	gc.SetStrokeWidth(thread.width)
	gc.SetStrokeColor(thread.color)
	for i, p1 := range pl1 {
		j := i
		if reversed {
			j = len(pl2) - i - 1
		}
		p2 := pl2[j]
		gc.DrawLine(p1.X, p1.Y, p2.X, p2.Y)
		if cont && i < len(pl1)-1 {
			p1 = pl1[i+1]
			gc.DrawLine(p2.X, p2.Y, p1.X, p1.Y)
		}
	}
	gc.Stroke()
}

func Cardioide(gc *gg.Context, pl []geom.Point, thread ThreadData, f int) {
	gc.SetStrokeWidth(thread.width)
	gc.SetStrokeColor(thread.color)
	for i := 1; i < len(pl); i++ {
		p1 := pl[i]
		j := (f * i) % (len(pl))
		p2 := pl[j]
		gc.DrawLine(p1.X, p1.Y, p2.X, p2.Y)
	}
	gc.Stroke()
}

func main() {
	var pointList []geom.Point
	// var endPointList []geom.Point

	pointList = make([]geom.Point, 0)
	// endPointList = []geom.Point{
	//     geom.Point{MarginSize, MarginSize},
	//     geom.Point{Width-MarginSize, MarginSize},
	//     geom.Point{Width-MarginSize, Height-MarginSize},
	//     geom.Point{MarginSize, Height-MarginSize},
	// }

	// for i, startPoint := range endPointList {
	//     endPoint := endPointList[(i+1)%len(endPointList)]
	//     pointList = append(pointList, LinearPointList(startPoint, endPoint, NumPoints)...)
	// }

	// pointList = append(pointList, CircularPointList(geom.Point{MarginSize, Height/2}, Width/2-MarginSize, 0, math.Pi/2, NumPoints/2)...)
	// pointList = append(pointList, CircularPointList(geom.Point{Width/2, MarginSize}, Width/2-MarginSize, math.Pi/2, math.Pi, NumPoints/2)...)
	// pointList = append(pointList, CircularPointList(geom.Point{Width-MarginSize, Height/2}, Width/2-MarginSize, math.Pi, 3*math.Pi/2, NumPoints/2)...)
	// pointList = append(pointList, CircularPointList(geom.Point{Width/2, Height-MarginSize}, Width/2-MarginSize, 3*math.Pi/2, 2*math.Pi, NumPoints/2)...)

	// pointList = append(pointList,
	//     BezierPointList(geom.Point{MarginSize, Height/2},
	//             geom.Point{MarginSize, MarginSize},
	//             geom.Point{Width-MarginSize, Height-MarginSize},
	//             geom.Point{Width-MarginSize, Height/2}, NumPoints)...)

	pointList = append(pointList, CircularPointList(geom.Point{Width / 2, Height / 2}, Width/2-MarginSize, 0, 2*math.Pi, NumPoints)...)

	gc := gg.NewContext(Width, Height)
	gc.SetFillColor(BackColor)
	gc.Clear()

	// Weave(gc, pointList[:NumPoints/2], pointList[NumPoints:3*NumPoints/2], ThreadList[0], false, true)
	// Weave(gc, pointList[NumPoints/2:NumPoints], pointList[3*NumPoints/2:], ThreadList[1], false, true)

	// Weave(gc, pointList[:NumPoints/2], pointList[NumPoints/2:NumPoints], ThreadList[0], false, false)
	// Weave(gc, pointList[NumPoints/2:NumPoints], pointList[NumPoints:3*NumPoints/2], ThreadList[1], false, false)
	// Weave(gc, pointList[NumPoints:3*NumPoints/2], pointList[3*NumPoints/2:], ThreadList[2], false, false)
	// Weave(gc, pointList[3*NumPoints/2:2*NumPoints], pointList[:NumPoints/2], ThreadList[4], false, false)

	Cardioide(gc, pointList, ThreadList[0], 2)

	// DrawPointList(gc, pointList)

	gc.SavePNG(OutFileName)
}
