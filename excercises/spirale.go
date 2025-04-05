package main

import (
	"math"

	"github.com/stefan-muehlebach/gg"
	"github.com/stefan-muehlebach/gg/colors"
	"github.com/stefan-muehlebach/gg/geom"
)

const (
	Width, Height = 512.0, 512.0
	MarginSize    = 24.0
	CanvasSize    = Width - 2*MarginSize
	NumMajorTicks = 20
	TickWidth     = CanvasSize / NumMajorTicks
	LineWidth     = 3.0
	AxesLineWidth = 0.8
	OutFileName   = "spirale.png"
)

var (
	BackColor = colors.RGBAF{0.851, 0.811, 0.733, 1.0}
	LineColor = colors.RGBAF{0.153, 0.157, 0.133, 1.0}
)

func main() {
	gc := gg.NewContext(Width, Height)
	gc.SetFillColor(BackColor)
	gc.Clear()

	gc.SetStrokeColor(LineColor)
	gc.SetStrokeWidth(LineWidth)
	mp1 := geom.Point{Width / 2.0, Height / 2.0}
	mp2 := mp1.AddXY(TickWidth/2.0, 0.0)
	for i := 0; i < NumMajorTicks/2-1; i += 1 {
		radius1 := float64(i+1) * TickWidth
		radius2 := radius1 + (TickWidth / 2.0)
		gc.DrawArc(mp1.X, mp1.Y, radius1, 2*math.Pi, math.Pi)
		gc.DrawArc(mp2.X, mp2.Y, radius2, math.Pi, 0.0)
	}
	gc.Stroke()

	gc.SetStrokeWidth(AxesLineWidth)
	gc.DrawLine(Width/2, MarginSize, Width/2, MarginSize+CanvasSize)
	gc.DrawLine(MarginSize, Height/2, MarginSize+CanvasSize, Height/2)
	gc.Stroke()

	gc.SetFillColor(colors.Crimson.Alpha(0.7))
	gc.DrawPoint(mp1.X, mp1.Y, 1.2*LineWidth)
	gc.DrawPoint(mp2.X, mp2.Y, 1.2*LineWidth)
	gc.Fill()

	gc.SavePNG(OutFileName)
}
