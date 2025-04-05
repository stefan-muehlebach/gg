package main

import (
	"github.com/stefan-muehlebach/gg"
	"github.com/stefan-muehlebach/gg/colors"
	"github.com/stefan-muehlebach/gg/geom"
	"math"
)

const (
	Width, Height = 512.0, 512.0
	MarginSize    = 24.0
	CanvasSize    = Width - 2*MarginSize
	NumFields     = 4
	FieldSize     = CanvasSize / NumFields
	Rad           = FieldSize / 2.0
)

type CircleData struct {
	color  colors.Color
	radius float64
}

var (
	// Per GIMP Color-Picker aus der verd*** Fotographie entnommen.
	BackColor = colors.RGBAF{0.851, 0.811, 0.733, 1.0}
	DataList  = []CircleData{
		{colors.RGBAF{0.125, 0.166, 0.356, 1.0}, Rad},
		{colors.RGBAF{0.392, 0.264, 0.298, 1.0}, 0.66 * Rad},
		{colors.RGBAF{0.907, 0.340, 0.169, 1.0}, 0.48 * Rad},
		{colors.RGBAF{0.930, 0.554, 0.104, 1.0}, 0.30 * Rad},
	}
)

func main() {
	gc := gg.NewContext(Width, Height)
	gc.SetFillColor(BackColor)
	gc.Clear()
	xSign, ySign := 1.0, 1.0
	for row := 0; row < NumFields; row++ {
		y := MarginSize + (float64(row)+0.5)*FieldSize
		if row%2 == 0 {
			ySign = -1
		} else {
			ySign = +1
		}
		for col := 0; col < NumFields; col++ {
			x := MarginSize + (float64(col)+0.5)*FieldSize
			if col%2 == 0 {
				xSign = +1
			} else {
				xSign = -1
			}
			dp := geom.Point{xSign * math.Sqrt(2) / 2, ySign * math.Sqrt(2) / 2}
			mp := geom.Point{x, y}
			refPt := mp.Sub(dp.Mul(Rad))
			for _, data := range DataList {
				p := refPt.Add(dp.Mul(data.radius))
				gc.DrawCircle(p.X, p.Y, data.radius)
				gc.SetFillColor(data.color)
				gc.Fill()
			}
		}
	}

	gc.SavePNG("bauhaus.png")
}
