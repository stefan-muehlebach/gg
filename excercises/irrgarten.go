package main

import (
	"github.com/stefan-muehlebach/gg"
	"github.com/stefan-muehlebach/gg/color"
	"github.com/stefan-muehlebach/gg/geom"
)

const (
	Width, Height = 512.0, 512.0
	MarginSize    = 24.0
	CanvasSize    = Width - 2*MarginSize
	NumRaster     = 21
	RasterSize    = CanvasSize / NumRaster
	LineWidth     = 3.0
	OutFileName   = "irrgarten.png"
)

var (
	BackColor  = color.RGBAF{0.851, 0.811, 0.733, 1.0}
	LineColor  = color.RGBAF{0.153, 0.157, 0.133, 1.0}
	WhiteColor = color.Ivory
	BlackColor = color.RGBAF{0.153, 0.157, 0.133, 1.0}
	TextColor  = color.RGBAF{0.153, 0.157, 0.133, 1.0}
	DirList    = []geom.Point{
		{+1.0, 0.0},
		{0.0, +1.0},
		{-1.0, 0.0},
		{0.0, -1.0},
	}
)

func main() {
	gc := gg.NewContext(Width, Height)
	gc.SetFillColor(BackColor)
	gc.Clear()

	gc.SetStrokeColor(LineColor)
	gc.SetStrokeWidth(LineWidth)

	p := geom.Point{Width / 2.0, Height/2.0 - RasterSize/2.0}
	gc.MoveTo(p.AsCoord())
	for i := 0; i < NumRaster-1; i++ {
		p = p.Add(DirList[i%4].Mul(float64(i+1) * RasterSize))
		gc.LineTo(p.AsCoord())
	}
	gc.Stroke()

	p = geom.Point{Width / 2.0, Height/2.0 + RasterSize/2.0}
	gc.MoveTo(p.AsCoord())
	for i := 0; i < NumRaster-1; i++ {
		p = p.Sub(DirList[i%4].Mul(float64(i+1) * RasterSize))
		gc.LineTo(p.AsCoord())
	}
	gc.Stroke()

	gc.SavePNG(OutFileName)
}
