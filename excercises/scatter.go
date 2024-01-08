package main

import (
	"github.com/stefan-muehlebach/gg"
	"github.com/stefan-muehlebach/gg/color"
	"github.com/stefan-muehlebach/gg/colornames"
	"github.com/stefan-muehlebach/gg/fonts"
	"github.com/stefan-muehlebach/gg/geom"
	"math"
	// "image/color"
	"math/rand"
)

const (
	Width, Height = 512.0, 512.0
	Padding       = 24
	OutFileName   = "scatter.png"
	NumPoints     = 1000
	TitleFontSize = 12.0
	TextFontSize  = 10.0
)

var (
	BackColor = color.RGBAF{0.851, 0.811, 0.733, 1.0}
	LineColor = color.RGBAF{0.153, 0.157, 0.133, 1.0}
	TitleFont = fonts.GoBold
	TextFont  = fonts.GoRegular
)

func NewPointList(n int) []geom.Point {
	points := make([]geom.Point, n)
	for i := 0; i < n; i++ {
		x := 0.5 + rand.NormFloat64()*0.1
		y := x + rand.NormFloat64()*0.1
		points[i] = geom.Point{x, y}
	}
	return points
}

func main() {
	dc := gg.NewContext(Width, Height)
	dc.Scale(1.0, -1.0)
	dc.Translate(0.0, -Height)
	dc.SetFillColor(BackColor)
	dc.Clear()
	points := NewPointList(NumPoints)
	dc.Translate(Padding, Padding)
	dc.Scale(Width-Padding*2, Height-Padding*2)
	// draw minor grid
	for i := 1; i <= 10; i++ {
		x := float64(i) / 10
		dc.MoveTo(x, 0)
		dc.LineTo(x, 1)
		dc.MoveTo(0, x)
		dc.LineTo(1, x)
	}
	dc.SetStrokeColor(LineColor)
	dc.SetStrokeWidth(1)
	dc.Stroke()
	// draw axes
	dc.MoveTo(0, 0)
	dc.LineTo(1, 0)
	dc.MoveTo(0, 0)
	dc.LineTo(0, 1)
	dc.SetStrokeColor(LineColor)
	dc.SetStrokeWidth(4)
	dc.Stroke()
	// draw points
	dc.SetFillColor(colornames.Blue.Alpha(0.5))
	for _, p := range points {
		dc.DrawPoint(p.X, p.Y, 1.5)
		dc.Fill()
	}
	// draw text
	dc.Identity()
	dc.SetStrokeColor(LineColor)
	dc.SetFontFace(fonts.NewFace(TitleFont, TitleFontSize))
	dc.DrawStringAnchored("Beispiel für ein Streudiagramm (Scatter-Plot)", Width/2, Padding/2, 0.5, 0.5)
	dc.SetFontFace(fonts.NewFace(TextFont, TextFontSize))
	dc.DrawStringAnchored("X-Achse und Beschriftung", Width/2, Height-Padding/2, 0.5, 0.5)
	dc.Push()
	dc.RotateAbout(-math.Pi/2.0, Padding/2, Height/2)
	dc.DrawStringAnchored("Y-Achse und Beschriftung", Padding/2, Height/2, 0.5, 0.5)
	dc.Pop()
	dc.SavePNG(OutFileName)
}
