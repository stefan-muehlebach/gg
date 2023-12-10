package main

import (
	"github.com/stefan-muehlebach/gg/geom"
	"github.com/stefan-muehlebach/gg"
	"github.com/stefan-muehlebach/gg/colornames"
)

const (
	Width, Height   = 512.0, 512.0
	MarginSize      = 16.0
    NumSteps        = 60
    PartitionFactor = 0.05
	GridLineWidth   = 1.5
)

var (
    	BackColor = colornames.Indigo.Dark(0.5)
    GridLineColor = colornames.Lightgray
)

func main() {
    var p0, p1, p2, p3, p4 geom.Point
    
    p1 = geom.NewPoint(MarginSize, MarginSize)
    p2 = p1.AddXY(Width-2*MarginSize, 0.0)
    p3 = p2.AddXY(0.0, Height-2*MarginSize)
    p4 = p1.AddXY(0.0, Height-2*MarginSize)
    
	gc := gg.NewContext(Width, Height)
	gc.SetFillColor(BackColor)
	gc.Clear()
	gc.SetStrokeWidth(GridLineWidth)
	gc.SetStrokeColor(GridLineColor)
	gc.MoveTo(p1.X, p1.Y)
    gc.LineTo(p2.X, p2.Y)
    gc.LineTo(p3.X, p3.Y)
    gc.LineTo(p4.X, p4.Y)
    gc.ClosePath()
    gc.Stroke()
    for i:=0; i<NumSteps; i++ {
        p0 = p1
        p1 = p1.Interpolate(p2, PartitionFactor)
        p2 = p2.Interpolate(p3, PartitionFactor)
        p3 = p3.Interpolate(p4, PartitionFactor)
        p4 = p4.Interpolate(p0, PartitionFactor)
        	gc.MoveTo(p1.X, p1.Y)
        gc.LineTo(p2.X, p2.Y)
        gc.LineTo(p3.X, p3.Y)
        gc.LineTo(p4.X, p4.Y)
        gc.ClosePath()
        gc.Stroke()
    }    
	gc.SavePNG("vertigo.png")
}
