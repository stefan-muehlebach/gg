package main

import (
	"github.com/stefan-muehlebach/gg"
	"github.com/stefan-muehlebach/gg/colornames"
)

const (
	Width, Height = 512.0, 512.0
	MarginSize    = 16.0
	FieldSize     = Width - 2*MarginSize
	NumGridLines  = 30
	GridLineSep   = FieldSize / float64(NumGridLines-1)
	GridLineWidth = 1.5
)

var (
	BackColor     = colornames.Midnightblue
	GridLineColor = colornames.Whitesmoke
)

func main() {
	gc := gg.NewContext(Width, Height)
	gc.SetFillColor(BackColor)
	gc.Clear()
	gc.SetLineWidth(GridLineWidth)
	gc.SetStrokeColor(GridLineColor)
	xa, ya := MarginSize, MarginSize
	xb, yb := Width-MarginSize, Height-MarginSize
	for n := 0; n < NumGridLines; n++ {
		d := float64(n) * GridLineSep
		x1, y1a, y1b := MarginSize+d, MarginSize, Height-MarginSize
		x2a, x2b, y2 := MarginSize, Width-MarginSize, MarginSize+d
		gc.DrawLine(xa, ya, x1, y1b)
		gc.DrawLine(xb, yb, x1, y1a)
		gc.DrawLine(xa, ya, x2b, y2)
		gc.DrawLine(xb, yb, x2a, y2)
		gc.Stroke()
	}
	gc.SavePNG("moiree.png")
}
