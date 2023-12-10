package main

import (
	"math"
    "log"
	"github.com/stefan-muehlebach/gg"
	"github.com/stefan-muehlebach/gg/color"
	"github.com/stefan-muehlebach/gg/colornames"
)

const (
	Width, Height = 512.0, 512.0
	MarginSize    = 16.0
    GraphSize     = Width - 2*MarginSize
    
    xMin, xMax = -1.0, 1.0
    yMin, yMax = -1.0, 1.0
)

type heightFunc func(x, y float64) float64

func f1(x, y float64) (float64) {
    return x*x*y*y
}

func f2(x, y float64) (float64) {
    return math.Sqrt(4.0/9.0-x*x-y*y)
}

var (
    	BackColor = colornames.Teal.Dark(0.7)
	GridLineColor = colornames.Lightgray
)

func main() {
	gc := gg.NewContext(Width, Height)
	gc.SetFillColor(BackColor)
	gc.Clear()

    col0, row0 := int(MarginSize), int(Height-MarginSize)
    for row:=0; row<GraphSize; row++ {
        y := (yMax-yMin) * float64(row)/float64(GraphSize) - 1.0
        for col:=0; col<GraphSize; col++ {
            x := (xMax-xMin) * float64(col)/float64(GraphSize) - 1.0
            z := f2(x, y)
            if z > 1.0 || z < 0.0 {
                log.Printf("z is out of range: %f\n", z)
            }
            c := col0 + col
            r := row0 - row
            color := color.RGBAF{z, z, z, 1.0}
            gc.SetPixel(c, r, color)
        }
    }

	// gc.SetStrokeWidth(LineWidth)
	// gc.SetStrokeColor(LineColor)
	// xa, ya := MarginSize, MarginSize
	// xb, yb := Width-MarginSize, Height-MarginSize
	// for n := 0; n < NumGridLines; n++ {
	// 	d := float64(n) * GridLineSep
	// 	x1, y1a, y1b := MarginSize+d, MarginSize, Height-MarginSize
	// 	x2a, x2b, y2 := MarginSize, Width-MarginSize, MarginSize+d
	// 	gc.DrawLine(xa, ya, x1, y1b)
	// 	gc.DrawLine(xb, yb, x1, y1a)
	// 	gc.DrawLine(xa, ya, x2b, y2)
	// 	gc.DrawLine(xb, yb, x2a, y2)
	// 	gc.Stroke()
	// }
	gc.SavePNG("graph.png")
}
