package main

import (
	"github.com/stefan-muehlebach/gg"
	"github.com/stefan-muehlebach/gg/color"
)

const (
	Width, Height     = 512.0, 512.0
	MarginSize        = 32.0
	ColorColumnWidth  = (Width - 2.0*MarginSize) / 3.0
	ColorColumnHeight = Height - 2.0*MarginSize
	NumShades         = 10
	LineWidth         = 32.0
)

func main() {
	gc := gg.NewContext(512.0, 512.0)

	gc.DrawRectangle(MarginSize, MarginSize, ColorColumnWidth, ColorColumnHeight)
	gc.SetFillColor(color.Black)
	gc.Fill()

	gc.DrawRectangle(MarginSize+2*ColorColumnWidth, MarginSize, ColorColumnWidth, ColorColumnHeight)
	gc.SetFillColor(color.White)
	gc.Fill()

	gc.SetStrokeWidth(LineWidth)
	x1, x2, y := MarginSize+ColorColumnWidth/2.0, Width-MarginSize-ColorColumnWidth/2.0, MarginSize+ColorColumnWidth/2.0
	gc.DrawLine(x1, y, x2, y)
	gc.SetStrokeColor(color.White.Alpha(0.5))
	gc.Stroke()
	y += LineWidth
	gc.DrawLine(x1, y, x2, y)
	gc.SetStrokeColor(color.RGBAF{0.5, 0.5, 0.5, 0.5})
	gc.Stroke()

	gc.SavePNG("colors.png")
}
