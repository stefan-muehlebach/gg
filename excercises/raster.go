package main

import (
	"fmt"
	"github.com/stefan-muehlebach/gg"
	"github.com/stefan-muehlebach/gg/color"
	"github.com/stefan-muehlebach/gg/geom"
	"math"
)

type RasterFormat int

const (
    SquareRasterFormat RasterFormat = iota
    DiamondRasterFormat
    TriangleRasterFormat
    numRasterFormats
)

const (
	Width, Height = 2048.0, 2048.0
	MarginSize    = Width / 16.0
	CanvasSize    = Width - 2*MarginSize
	NumPointsX    = 25
    OutFileName   = "raster-%d.png"
)

var (
	BackColor = color.RGBAF{0.851, 0.811, 0.733, 1.0}
	LineColor = color.RGBAF{0.153, 0.157, 0.133, 1.0}

    Sqrt3       = math.Sqrt(3.0)
    Sqrt3_2     = math.Sqrt(3.0)/2.0

    NumPointsY int
    PointGapX, PointGapY float64
    MinPointRad, MaxPointRad float64
)

func RasterGradient(gc *gg.Context, format RasterFormat) {
    var RasterCoord func(col, row int) (geom.Point, bool)

    switch format {
    case SquareRasterFormat:
        NumPointsY = NumPointsX
        PointGapX = CanvasSize / float64(NumPointsX-1)
        PointGapY = CanvasSize / float64(NumPointsY-1)
        MinPointRad = 0.5
        MaxPointRad = PointGapX / 2.0
        RasterCoord = func(col, row int) (geom.Point, bool) {
            return geom.Point{MarginSize + float64(col) * PointGapX,
                MarginSize + float64(row) * PointGapY}, true
        }
    case DiamondRasterFormat:
        NumPointsY = 2*NumPointsX - 1
        PointGapX = CanvasSize / float64(NumPointsX-1)
        PointGapY = CanvasSize / float64(NumPointsY-1)
        MinPointRad = 0.5
        MaxPointRad = PointGapX * math.Sqrt2 / 4.0
        RasterCoord = func(col, row int) (geom.Point, bool) {
            if (row%2)==1 && col==(NumPointsX-1) {
                return geom.Point{}, false
            } else {
                return geom.Point{MarginSize + float64(col) * PointGapX +
                    float64(row%2)*(PointGapX/2.0),
                    MarginSize + float64(row) * PointGapY}, true
            }
        }
    case TriangleRasterFormat:
        NumPointsY = int(math.Ceil(float64(NumPointsX)/Sqrt3_2))
        PointGapX = CanvasSize / float64(NumPointsX-1)
        PointGapY = CanvasSize / float64(NumPointsY-1)
        MinPointRad = 0.5
        MaxPointRad = PointGapX / 2.0
        RasterCoord = func(col, row int) (geom.Point, bool) {
            if (row%2)==1 && col==(NumPointsX-1) {
                return geom.Point{}, false
            } else {
                return geom.Point{MarginSize + float64(col)*PointGapX +
                    float64(row%2)*(PointGapX/2.0),
                    MarginSize + float64(row)*PointGapY}, true
            }
        }
    }

	gc.SetFillColor(LineColor)
	for row := 0; row < NumPointsY; row++ {
		t := float64(row) / float64(NumPointsY-1)
		for col := 0; col < NumPointsX; col++ {
            mp, draw := RasterCoord(col, row)
            if !draw {
                continue
            }
			pointRad := MinPointRad*(1-t) + MaxPointRad*t
			gc.DrawPoint(mp.X, mp.Y, pointRad)
			gc.Fill()
		}
	}
}

func main() {
    gc := gg.NewContext(Width, Height)
    for format := SquareRasterFormat; format < numRasterFormats; format += 1 {
        gc.SetFillColor(BackColor)
        gc.Clear()
        RasterGradient(gc, format)
        outFileName := fmt.Sprintf(OutFileName, format)
        gc.SavePNG(outFileName)
    }
}
