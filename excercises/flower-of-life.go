package main

import (
	"github.com/stefan-muehlebach/gg"
	"github.com/stefan-muehlebach/gg/color"
	"github.com/stefan-muehlebach/gg/geom"
	"math"
)

const (
	Width, Height = 512.0, 512.0
	MarginSize    = 24.0
	CanvasSize    = Width - 2*MarginSize
	RadiusBig     = CanvasSize / 2.0
	RadiusSmall   = CanvasSize / 6.0
	LineWidth     = 2.5
	OutFileName   = "flower-of-life.png"
)

var (
	BackColor = color.RGBAF{0.851, 0.811, 0.733, 1.0}
	LineColor = color.DarkGoldenrod.Dark(0.2)
)

func main() {
	gc := gg.NewContext(Width, Height)
	gc.SetFillColor(BackColor)
	gc.Clear()
	gc.Translate(Width/2.0, Height/2.0)
	gc.Scale(1.0, -1.0)

	gc.SetStrokeColor(LineColor)
	gc.SetStrokeWidth(LineWidth)
	mp := geom.Point{}
	gc.DrawCircle(mp.X, mp.Y, RadiusBig+LineWidth/2)
	gc.DrawCircle(mp.X, mp.Y, RadiusBig+MarginSize/3)
	gc.DrawCircle(mp.X, mp.Y, RadiusSmall)
	gc.Stroke()
	angle := math.Pi / 3.0

	dp1 := geom.Point{0.0, 1.0}.Mul(RadiusSmall)
	dp2 := geom.Point{math.Sqrt(3) / 2.0, 0.5}.Mul(RadiusSmall)

	gc.Push()
	for i := 0; i < 6; i++ {
		p := mp.Add(dp1)
		gc.DrawCircle(p.X, p.Y, RadiusSmall)
		gc.Stroke()
		p = p.Add(dp1)
		gc.DrawCircle(p.X, p.Y, RadiusSmall)
		gc.Stroke()
		p = p.Add(dp1)
		gc.DrawArc(p.X, p.Y, RadiusSmall, math.Pi+math.Pi/6.0, 2*math.Pi-math.Pi/6.0)
		gc.Stroke()

		p = mp.Add(dp1).Add(dp2)
		gc.DrawCircle(p.X, p.Y, RadiusSmall)
		gc.Stroke()
		p = p.Add(dp1)
		gc.DrawArc(p.X, p.Y, RadiusSmall, math.Pi-math.Pi/6.0, 2*math.Pi-math.Pi/6.0)
		gc.Stroke()
		p = p.Add(dp1)
		gc.DrawArc(p.X, p.Y, RadiusSmall, math.Pi+math.Pi/6.0, 3*math.Pi/2.0)
		gc.Stroke()

		p = mp.Add(dp1).Add(dp2).Add(dp2)
		gc.DrawArc(p.X, p.Y, RadiusSmall, math.Pi-math.Pi/6.0, 2*math.Pi-math.Pi/6.0)
		gc.Stroke()
		p = p.Add(dp1)
		gc.DrawArc(p.X, p.Y, RadiusSmall, math.Pi+math.Pi/6.0, 3*math.Pi/2.0)
		gc.Stroke()

		p = mp.Add(dp1).Add(dp2).Add(dp2).Add(dp2)
		gc.DrawArc(p.X, p.Y, RadiusSmall, math.Pi+math.Pi/6.0, 3*math.Pi/2.0)
		gc.Stroke()

		gc.Rotate(angle)
	}
	gc.Pop()
	gc.SavePNG(OutFileName)
}
