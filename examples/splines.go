package main

import (
	// "math/rand"
	// "github.com/stefan-muehlebach/gg/fonts"
    	"github.com/stefan-muehlebach/gg"
    	"github.com/stefan-muehlebach/gg/geom"
	// "github.com/stefan-muehlebach/gg/color"
	"github.com/stefan-muehlebach/gg/colornames"
)

const (
	Width, Height = 512.0, 512.0
    MarginSize = 32.0
    FieldSize = Width - 2*MarginSize
    Message = "Wir sind verloren!"
    LineWidth = 10.0
)

var (
    	BackColor = colornames.Teal.Dark(0.7)
    LineColor = colornames.Whitesmoke.Alpha(0.5)
)

func main() {
	gc := gg.NewContext(Width, Height)
	gc.SetFillColor(BackColor)
	gc.Clear()
    

    pA := geom.Point{MarginSize, MarginSize}
    pB := pA.AddXY(FieldSize, 0.0)
    pC := pA.AddXY(0.0, FieldSize)
    pD := pC.AddXY(FieldSize, 0.0)

    p2 := geom.Point{Width/2.0, Height/2.0}

    gc.SetStrokeColor(LineColor)
    gc.SetStrokeWidth(LineWidth)
    for i:=0; i<10; i++ {
        t := float64(i)/10.0
        pa := pA.Interpolate(pC, t/2.0)
        pb := pB.Interpolate(pA, -t)
        pc := pC.Interpolate(pD, -t)        
        pd := pD.Interpolate(pB, t/2.0)
        p1 := pC.Interpolate(pd, t)
        p3 := pB.Interpolate(pa, t)
    
        gc.MoveTo(p1.AsCoord())
        gc.CubicTo(pa.X, pa.Y, pb.X, pb.Y, p2.X, p2.Y)
        gc.CubicTo(pc.X, pc.Y, pd.X, pd.Y, p3.X, p3.Y)
        gc.SetStrokeColor(LineColor)
        gc.SetStrokeWidth(LineWidth)
        gc.StrokePreserve()
        gc.SetStrokeColor(BackColor.Alpha(0.5))
        gc.SetStrokeWidth(7.0)
        gc.Stroke()
    }

	gc.SavePNG("splines.png")
}
