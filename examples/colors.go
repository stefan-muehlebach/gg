package main

import (
	"github.com/stefan-muehlebach/gg"
	"github.com/stefan-muehlebach/gg/colors"
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
	gc.SetFillColor(colors.Black)
	gc.Fill()

	gc.DrawRectangle(MarginSize+2*ColorColumnWidth, MarginSize, ColorColumnWidth, ColorColumnHeight)
	gc.SetFillColor(colors.White)
	gc.Fill()

	gc.SetStrokeWidth(LineWidth)
	x1, x2, y := MarginSize+ColorColumnWidth/2.0, Width-MarginSize-ColorColumnWidth/2.0, MarginSize+ColorColumnWidth/2.0
	gc.DrawLine(x1, y, x2, y)
	gc.SetStrokeColor(colors.White.Alpha(0.333))
	gc.Stroke()
	y += LineWidth
	gc.DrawLine(x1, y, x2, y)
	gc.SetStrokeColor(colors.RGBAF{0.5, 0.5, 0.5, 0.333})
	gc.Stroke()
	y += LineWidth
	gc.DrawLine(x1, y, x2, y)
	gc.SetStrokeColor(colors.Black.Alpha(0.333))
	gc.Stroke()

	y += 2 * LineWidth
	gc.DrawLine(x1, y, x2, y)
	gc.SetStrokeColor(colors.White.Alpha(0.666))
	gc.Stroke()
	y += LineWidth
	gc.DrawLine(x1, y, x2, y)
	gc.SetStrokeColor(colors.RGBAF{0.5, 0.5, 0.5, 0.666})
	gc.Stroke()
	y += LineWidth
	gc.DrawLine(x1, y, x2, y)
	gc.SetStrokeColor(colors.Black.Alpha(0.666))
	gc.Stroke()

	y += 2 * LineWidth
	gc.DrawLine(x1, y, x2, y)
	gc.SetStrokeColor(colors.White)
	gc.Stroke()
	y += LineWidth
	gc.DrawLine(x1, y, x2, y)
	gc.SetStrokeColor(colors.RGBAF{0.5, 0.5, 0.5, 1.0})
	gc.Stroke()
	y += LineWidth
	gc.DrawLine(x1, y, x2, y)
	gc.SetStrokeColor(colors.Black)
	gc.Stroke()

	gc.SavePNG("colors.png")
}
