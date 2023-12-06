package main

import (
	"fmt"
    "math"
    "github.com/stefan-muehlebach/gg"
    "github.com/stefan-muehlebach/gg/colornames"
    "github.com/stefan-muehlebach/gg/fonts"
    "github.com/stefan-muehlebach/gg/geom"
)

type LinSpace struct {
	startVal, endVal, stepSize float64
	numSteps                   int
}

type LinSpaceNode struct {
	T, Value float64
	Step     int
}

func NewLinSpace(startVal, endVal float64) *LinSpace {
	ls := &LinSpace{}
	ls.startVal = startVal
	ls.endVal   = endVal
	return ls
}

func (ls *LinSpace) Range(numSteps int) <-chan LinSpaceNode {
	ch := make(chan LinSpaceNode, bufferSize)
	ls.numSteps = numSteps
	ls.stepSize = (ls.endVal - ls.startVal) / float64(numSteps)
	go func() {
		for step := 0; step <= ls.numSteps; step++ {
			t := float64(step) / float64(numSteps)
			val := ls.startVal + float64(step)*ls.stepSize
			ch <- LinSpaceNode{T: t, Value: val, Step: step}
		}
		close(ch)
	}()
	return ch
}

func (ls *LinSpace) RelPos(val float64) float64 {
	return (val - ls.startVal) / (ls.endVal - ls.startVal)
}

func (ls *LinSpace) Value(t float64) float64 {
	return (1.0-t)*ls.startVal + t*ls.endVal
}

func Spiral(gc *gg.Context) {
	fillColor := colornames.Black
	fillColor.A = 0.1
	gc.SetFillColor(fillColor)
	linSpace := NewLinSpace(0.0, 360.0)
	for pt := range linSpace.Range(24) {
		// fmt.Printf("received %.3f from channel\n", alpha)
		gc.Push()
		gc.RotateAbout(gg.Radians(pt.Value), Width/2, Height/2)
		gc.DrawEllipse(Width/2, Height/2, Width*7/16, Height/8)
		gc.Fill()
		gc.Pop()
	}
}

func ScaledText(gc *gg.Context) {
	var path *geom.Path

	path = geom.NewPath()
	path.MoveTo(geom.Point{Width/2, 5 * Height/6})
	path.LineTo(geom.Point{Width/2, Height/6})

	gc.SetFillColor(colornames.Black)
	gc.Clear()
	face := fonts.NewFace(fonts.Map["LucidaBlackletter"], 100)
	textColor := colornames.White
	textColor.A = 0.3
	gc.SetStrokeColor(textColor)
	gc.SetFontFace(face)
	linSpace := NewLinSpace(1.0, -1.0)
	alphaSpace := NewLinSpace(0.9, 0.1)
	for pt := range linSpace.Range(11) {
		gc.Push()
		textColor.A = alphaSpace.Value(pt.T)
		gc.SetStrokeColor(textColor)
		sp := path.PointNorm(pt.T)
		fmt.Printf("%.3f: %v\n", pt.T, sp)
		gc.ScaleAbout(pt.Value, pt.Value, sp.X, sp.Y)
		gc.DrawStringAnchored("Here we are!", sp.X, sp.Y, 0.5, 0.5)
		gc.Pop()
	}
}

func RotatedText(gc *gg.Context) {
	var path *geom.Path

	path = geom.NewPath()
	path.MoveTo(geom.Point{Width/5, 2*Height/3})
	path.BezierTo(geom.Point{Width/5, Height/3}, geom.Point{4*Width/5, Height/3}, geom.Point{4*Width/5, 2*Height/3})

	gc.SetFillColor(colornames.Black)
	gc.Clear()
	face := fonts.NewFace(fonts.Map["SeafordBold"], 96)
	textColor := colornames.White
	textColor.A = 0.7
	gc.SetFontFace(face)
	linSpace := NewLinSpace(0.0, float64(path.Segments()))
	for pt := range linSpace.Range(8) {
		gc.Push()
		rp := path.Point(pt.Value)
        dir := path.Dir(pt.Value)
        alpha := dir.Angle()
		// fmt.Printf("t: %.4f, angle: %2.4f, rotPt: %v, dir: (%v): %2.5f\n", pt.T, pt.Value, rp, dir, 180.0*alpha/math.Pi)
		gc.RotateAbout(alpha+math.Pi/2.0, rp.X, rp.Y)
        
        // gc.SetStrokeColor(colornames.White)
        // gc.DrawLine(rp.X-20.0, rp.Y, rp.X+20.0, rp.Y)
        // gc.Stroke()
        // gc.DrawRectangle(rp.X-50.0, rp.Y-50.0, 100.0, 100.0)
        // gc.Stroke()
        
		gc.SetFillColor(colornames.Red)
		gc.DrawPoint(rp.X, rp.Y, 5.0)
		gc.Fill()
        
        	gc.SetStrokeColor(textColor)
		gc.DrawStringAnchored("KREISE", rp.X, rp.Y, 0.5, 0.5)
		gc.Pop()
	}
    gc.SetStrokeColor(colornames.White)
    gc.MoveTo(Width/5, 2*Height/3)
    gc.CubicTo(Width/5, Height/3, 4*Width/5, Height/3, 4*Width/5, 2*Height/3)
    gc.Stroke()
}

func PathedText(gc *gg.Context) {
	var path *geom.Path
    var p0, c0, c1, p1 geom.Point

    path = geom.NewPath()

    p0 = geom.Point{Width/6, Height/6}
    c0 = geom.Point{Width/2, Height/6}
    c1 = geom.Point{Width/6, Height/2}
    p1 = geom.Point{Width/2, Height/2}
	
	path.MoveTo(p0)
	path.BezierTo(c0, c1, p1)
    
    c0 = geom.Point{2*Width/3, Height/2}
    c1 = geom.Point{2*Width/3, 2*Height/3}
    p1 = geom.Point{Width/2, 2*Height/3}

	path.BezierTo(c0, c1, p1)

	gc.SetFillColor(colornames.Black)
	gc.Clear()
    
	face := fonts.NewFace(fonts.Map["SeafordBold"], 96)
	gc.SetFontFace(face)
    
    fmt.Printf("length of path: %f\n", path.ArcLength())
    fmt.Printf("  spline 1: %f\n", path.Segment(0).ArcLength())
    fmt.Printf("  spline 2: %f\n", path.Segment(1).ArcLength())

    gc.SetStrokeColor(colornames.Grey)
    gc.SetStrokeWidth(2.5)
    gc.MoveTo(path.Start().AsCoord())
    for i:=0; i<path.Segments(); i++ {
        segm := path.Segment(i)
        switch s := segm.(type) {
        case geom.LinearSegment:
            gc.LineTo(s.End().AsCoord())
        // case geom.BezierSegment:
        //     gc.CubicTo(s.C0.X, s.C0.Y, s.C1.X, s.C1.Y, s.P1.X, s.P1.Y)
        }            
    }
    gc.Stroke()

	textColor := colornames.White
	textColor.A = 0.7
	linSpace := NewLinSpace(0.0, 2.0)
	for pt := range linSpace.Range(20) {
		refPt := path.Point(pt.Value)
        dir := path.Dir(pt.Value).Neg()
        alpha := dir.Angle() + math.Pi/2.0
        fmt.Printf("t: %f, val: %f, pt: %v, dir: %v, alpha: %f\n", pt.T, pt.Value, refPt, dir, alpha)
        
        endPt := refPt.Add(dir.Mul(20.0))
        gc.SetStrokeColor(colornames.White)
        gc.DrawLine(refPt.X, refPt.Y, endPt.X, endPt.Y)
        gc.Stroke()

		gc.SetFillColor(colornames.Lightgreen)
		gc.DrawPoint(refPt.X, refPt.Y, 5.0)
		gc.Fill()

		gc.Push()
        gc.RotateAbout(alpha, refPt.X, refPt.Y)        
        
        // gc.DrawRectangle(refPt.X-50.0, refPt.Y-50.0, 100.0, 100.0)
        // gc.SetStrokeColor(colornames.White)
        // gc.Stroke()
        // 	gc.SetStrokeColor(textColor)
		// gc.DrawStringAnchored("W", refPt.X, refPt.Y, 0.5, 0.5)
		gc.Pop()
	}
}

var (
	Width  = 2048.0
    Height = 4096.0
	bufferSize = 10
)

func keys[K comparable, V any](m map[K]V) ([]K) {
    keys := make([]K, 0, len(m))
    for k := range m {
        keys = append(keys, k)
    }
    return keys
}

func FontSamples(gc *gg.Context) {
    fontSize := 96.0
    margin := 40.0
    lineSize := fontSize + margin/2.0
	textColor := colornames.White
	textColor.A = 0.7
    
    Height = float64(len(fonts.Names))*lineSize + 2*margin

    gc = gg.NewContext(int(Width), int(Height))
	gc.SetFillColor(colornames.Black)
	gc.Clear()
    for i, fontName := range fonts.Names {
        // col := i / maxRows
        // row := i % maxRows
        x := margin
        y := margin/2.0 + float64(i+1)*lineSize
        face := fonts.NewFace(fonts.Map[fontName], fontSize)    
        	gc.SetFontFace(face)
    	    gc.SetStrokeColor(textColor)
        gc.DrawString(fontName, x, y)
        
        gc.SetStrokeColor(colornames.Lightyellow)
        gc.SetStrokeWidth(2.0)
        gc.MoveTo(x, y-10.0)
        gc.LineTo(x, y)
        gc.LineTo(x+10.0, y)
        // gc.Stroke()
        gc.MoveTo(x+10.0, y-fontSize)
        gc.LineTo(x, y-fontSize)
        gc.LineTo(x, y-fontSize+10.0)
        gc.Stroke()
    }
	gc.SavePNG("fontmap.png")
}

func main() {
	gc := gg.NewContext(int(Width), int(Height))
	FontSamples(gc)
	// gc.SavePNG("fonts.png")
}
