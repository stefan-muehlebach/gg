package main

import (
	"fmt"
	"github.com/stefan-muehlebach/gg"
	"github.com/stefan-muehlebach/gg/color"
	"github.com/stefan-muehlebach/gg/colornames"
	"github.com/stefan-muehlebach/gg/fonts"
	"github.com/stefan-muehlebach/gg/geom"
	"math"
)

const (
    // Breite und Höhe des Bildes in Pixel
	Width, Height = 512.0, 512.0
    // Freizulassender Rand
	MarginSize    = 20.0
    // Anzahl Punkte auf dem Kreisrand. Von dieser Zahl hängt das Muster
    // massgeblich ab. Primzahlen versprechen meist gute Muster.
	NumDivisions  = 101
    // Grösse der Punkte.
    PointSize     = 2.0
    // Grösse der Beschriftungen. Hier muss man etwas experimentieren, bis
    // die ideale Einstellung gefunden ist.
	FontSize      = 12.0
)

type GraphData struct {
	stepSize  int
	lineWidth float64
	color     color.Color
}

var (
    // BackColor und LineColor sind die Standardfarben für den Hintergrund
    // (das Papier) und den Vordergrund.
	BackColor = color.RGBAF{0.851, 0.811, 0.733, 1.0}
	LineColor = color.RGBAF{0.153, 0.157, 0.133, 1.0}

	GraphList = []GraphData{
		{44, 1.0, colornames.DarkViolet.Bright(0.2)},
		{36, 1.0, colornames.DarkBlue.Bright(0.1)},
		{32, 1.0, colornames.DarkGreen},
		{24, 1.0, colornames.DarkOliveGreen.Dark(0.1)},
		{20, 1.0, colornames.DarkOrange.Dark(0.2)},
	}
)

func DrawLines(gc *gg.Context, pl []geom.Point, data GraphData) {
	var idx0, idx1 int

	idx0 = 0
	gc.SetStrokeWidth(data.lineWidth)
	gc.SetStrokeColor(data.color)
	for {
		idx1 = (idx0 + data.stepSize) % len(pl)
		p0 := pl[idx0]
		p1 := pl[idx1]
		gc.DrawLine(p0.X, p0.Y, p1.X, p1.Y)
		gc.Stroke()
		idx0 = idx1
		if idx0 == 0 {
			break
		}
	}
}

func DrawPointList(gc *gg.Context, pl []geom.Point) {
	gc.SetFillColor(LineColor)
	for _, p := range pl {
		gc.DrawPoint(p.X, p.Y, PointSize)
		gc.Fill()
	}
}

func main() {
	var pointList []geom.Point
	var mp geom.Point
	var angle, step, radius float64

	pointList = make([]geom.Point, NumDivisions)
	step = (2.0 * math.Pi) / float64(NumDivisions)
	angle = 0.0
	radius = Width/2 - MarginSize
	mp = geom.Point{Width / 2, Height / 2}

	for i := 0; i < NumDivisions; i++ {
		angle = float64(i) * step
		p := geom.Point{math.Cos(angle), math.Sin(angle)}
		pointList[i] = p.Mul(radius).Add(mp)
	}

	gc := gg.NewContext(Width, Height)
	gc.SetFillColor(BackColor)
	gc.Clear()

	face := fonts.NewFace(fonts.GoMedium, FontSize)
	str := fmt.Sprintf("n = %d", NumDivisions)
	textPos := geom.Point{MarginSize, MarginSize}
	gc.SetFontFace(face)
	gc.SetStrokeColor(LineColor)
	gc.DrawStringAnchored(str, textPos.X, textPos.Y, 0.0, 1.0)

	for _, data := range GraphList {
		DrawLines(gc, pointList, data)
		str = fmt.Sprintf("s = %d", data.stepSize)
		textPos = textPos.AddXY(0.0, 1.2*FontSize)
		gc.SetStrokeColor(data.color)
		gc.DrawStringAnchored(str, textPos.X, textPos.Y, 0.0, 1.0)
	}

    DrawPointList(gc, pointList)

	gc.SavePNG("divided-circle.png")
}
