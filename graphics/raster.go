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
	FieldSize     = Width - 2*MarginSize
	NumPoints     = 49
	PointGap      = FieldSize / float64(NumPoints-1)
	MinPointRad   = 0.5
	MaxPointRad   = math.Sqrt2 * PointGap / 2.0
)

var (
	BackColor = color.RGBAF{0.851, 0.811, 0.733, 1.0}
	LineColor = color.RGBAF{0.153, 0.157, 0.133, 1.0}
)

// func drawPointsHexa(gc *gg.Context) {
// 	NumPointsY = int((2.0 / Sqrt3) * NumPointsX)
// 	GapX = FieldSize / float64(NumPointsX-1)
// 	GapY = GapX * Sqrt3 / 2.0
// 	MinPointRad = 0.5
// 	MaxPointRad = GapX / Sqrt3

// 	gc.SetFillColor(LineColor)
// 	for row := 0; row < NumPointsY; row++ {
// 		ty := float64(row) / float64(NumPointsY-1)
// 		for col := 0; col < NumPointsX-(row%2); col++ {
// 			xBias := float64(row%2) * GapX / 2.0
// 			t := ty
// 			pointRad := MinPointRad*(1-t) + MaxPointRad*t
// 			mp := geom.Point{MarginSize + xBias + float64(col)*GapX, MarginSize + float64(row)*GapY}
// 			gc.DrawPoint(mp.X, mp.Y, pointRad)
// 			gc.Fill()
// 		}
// 	}
// }

func main() {
	gc := gg.NewContext(Width, Height)
	gc.SetFillColor(BackColor)
	gc.Clear()
	gc.SetFillColor(LineColor)
	for row := 0; row < NumPoints; row++ {
		ty := float64(row) / float64(NumPoints-1)
		for col := 0; col < NumPoints; col++ {
			if row%2 != col%2 {
				continue
			}
			// tx := float64(col)/float64(NumPoints-1)
			t := ty
			pointRad := MinPointRad*(1-t) + MaxPointRad*t
			mp := geom.Point{MarginSize + float64(col)*PointGap, MarginSize + float64(row)*PointGap}
			gc.DrawPoint(mp.X, mp.Y, pointRad)
			gc.Fill()
		}
	}
	gc.SavePNG("raster.png")
}
