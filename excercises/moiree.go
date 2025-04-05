package main

import (
	"github.com/stefan-muehlebach/gg"
	"github.com/stefan-muehlebach/gg/colors"
	// "github.com/stefan-muehlebach/gg/colors"
)

const (
	Width, Height = 512.0, 512.0
	MarginSize    = 24.0
	CanvasSize    = Width - 2*MarginSize
	FieldSize     = (CanvasSize - MarginSize) / 2
	NumGridLines  = 30
	GridLineSep   = FieldSize / float64(NumGridLines-1)
	LineWidth     = 1.5
)

var (
	BackColor = colors.RGBAF{0.851, 0.811, 0.733, 1.0}
	LineColor = colors.RGBAF{0.153, 0.157, 0.133, 1.0}
)

func DrawGrid(gc *gg.Context, x0, y0, w, h float64, numCols int, rot float64) {
	if rot != 0.0 {
		gc.Push()
		gc.RotateAbout(rot, x0+(w/2), y0+(h/2))
		defer gc.Pop()
	}
	sep := w / float64(numCols-1)
	for col := 0; col < numCols; col++ {
		x := x0 + float64(col)*sep
		gc.DrawLine(x, y0, x, y0+h)
		gc.Stroke()
	}
}

func main() {
	gc := gg.NewContext(Width, Height)
	gc.SetFillColor(BackColor)
	gc.Clear()
	gc.SetStrokeWidth(LineWidth)
	gc.SetStrokeColor(LineColor)
	for row := 0; row < 2; row++ {
		y0 := MarginSize + float64(row)*(FieldSize+MarginSize)
		for col := 0; col < 2; col++ {
			x0 := MarginSize + float64(col)*(FieldSize+MarginSize)
			angle := 0.025 * float64(2*row+col+1)
			DrawGrid(gc, x0, y0, FieldSize, FieldSize, 40, 0.0)
			DrawGrid(gc, x0, y0, FieldSize, FieldSize, 40, angle)
		}
	}
	gc.SavePNG("moiree.png")
}
