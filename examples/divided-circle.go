package main

import (
	"github.com/stefan-muehlebach/gg/color"
    "fmt"
	"github.com/stefan-muehlebach/gg/fonts"
    "math"
	"github.com/stefan-muehlebach/gg/geom"
	"github.com/stefan-muehlebach/gg"
	"github.com/stefan-muehlebach/gg/colornames"
)

const (
	Width, Height   = 1024.0, 1024.0
	MarginSize      = 32.0
    NumDivisions    = 113
	LineWidth       = 1.5
    FontSize        = 28.0
)

var (
    StepSizeList    = []int{ 55, 44, 33 }
    LineWidthList   = []float64 { 0.8, 0.8, 0.8 }
    ColorList       = []color.Color{colornames.Linen, colornames.Lightblue, colornames.Lemonchiffon }
	BackColor       = colornames.Indigo.Dark(0.5)
	LineColor       = colornames.Silver
)

func DrawLines(gc *gg.Context, pl []geom.Point, ss int, lw float64, c color.Color) {
    var idx0, idx1 int
    
    idx0 = 0    
    gc.SetStrokeWidth(lw)
    gc.SetStrokeColor(c)
    for {
        idx1 = (idx0 + ss) % len(pl)
        p0 := pl[idx0]
        p1 := pl[idx1]
        gc.DrawLine(p0.X, p0.Y, p1.X, p1.Y)
        gc.Stroke()
        idx0 = idx1
        if idx0 == 0 {
            break
        }
    }
}

func main() {
    var pointList []geom.Point
    var mp geom.Point
    var angle, step, radius float64
    
    pointList = make([]geom.Point, NumDivisions)
    step = (2.0 * math.Pi) / float64(NumDivisions)
    angle = 0.0
    radius = Width/2 - MarginSize
    mp = geom.Point{Width/2, Height/2}
    
    for i:=0; i<NumDivisions; i++ {
        angle = float64(i) * step
        p := geom.Point{math.Cos(angle), math.Sin(angle)}
        pointList[i] = p.Mul(radius).Add(mp)
    }
    
	gc := gg.NewContext(Width, Height)
	gc.SetFillColor(BackColor)
	gc.Clear()
	gc.SetFillColor(LineColor)
    for _, p := range pointList {
        gc.DrawPoint(p.X, p.Y, 3.0)
        gc.Fill()
	}
    
    face := fonts.NewFace(fonts.GoMedium, FontSize)
    str := fmt.Sprintf("n = %d", NumDivisions)
    textPos := geom.Point{MarginSize, MarginSize}
    gc.SetFontFace(face)
    gc.SetStrokeColor(LineColor)
    gc.DrawStringAnchored(str, textPos.X, textPos.Y, 0.0, 1.0)
    
    for i, stepSize := range StepSizeList {
        DrawLines(gc, pointList, stepSize, LineWidthList[i], ColorList[i])
        str = fmt.Sprintf("s = %d", stepSize)
        textPos = textPos.AddXY(0.0, 1.2*FontSize)
        gc.SetStrokeColor(ColorList[i])
        gc.DrawStringAnchored(str, textPos.X, textPos.Y, 0.0, 1.0)
    }

	gc.SavePNG("divided-circle.png")    
}
