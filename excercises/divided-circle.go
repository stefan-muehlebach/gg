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
	Width, Height = 512.0, 512.0
	MarginSize    = Width / 32.0
	NumDivisions  = 97
	LineWidth     = 1.5
	FontSize      = 14.0
)

type GraphData struct {
	stepSize  int
	lineWidth float64
	color     color.Color
}

var (
	GraphList = []GraphData{
		//{54, 0.4, colornames.Blue.Dark(0.5)},
		{44, 0.4, colornames.Green.Dark(0.7)},
		//{33, 0.4, colornames.Red.Dark(0.7)},
	}
	BackColor = color.RGBAF{0.851, 0.811, 0.733, 1.0}
	LineColor = color.RGBAF{0.153, 0.157, 0.133, 1.0}
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
	gc.SetFillColor(LineColor)
	for _, p := range pointList {
		gc.DrawPoint(p.X, p.Y, 3.0)
		gc.Fill()
	}

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

	gc.SavePNG("divided-circle.png")
}
