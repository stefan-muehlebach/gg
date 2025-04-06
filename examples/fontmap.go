package main

import (
	"github.com/stefan-muehlebach/gg"
	"github.com/stefan-muehlebach/gg/colors"
	"github.com/stefan-muehlebach/gg/fonts"
)

var (
	MarginSize = 16.0
	FontSize   = 48.0
	LineHeight = 1.3 * FontSize
	BackColor  = colors.RGBAF{0.851, 0.811, 0.733, 1.0}
	TextColor  = colors.Black.Alpha(0.7)
	MarkerColor = colors.Crimson
	MarkerWidth = 2.0

	FileName = "fontmap.png"

	Width  = 1024.0
	Height = 2*MarginSize + FontSize + float64(len(fonts.Names)-1)*LineHeight

)

func main() {
	gc := gg.NewContext(int(Width), int(Height))
	gc.SetFillColor(BackColor)
	gc.Clear()
	for i, fontName := range fonts.Names {
		x := MarginSize
		y := MarginSize + FontSize + float64(i)*LineHeight
		face := fonts.NewFace(fonts.Map[fontName], FontSize)
		gc.SetFontFace(face)
		gc.SetStrokeColor(TextColor)
		gc.DrawString(fontName, x, y)

		w, _ := gc.MeasureString(fontName)

		gc.SetStrokeColor(MarkerColor)
		gc.SetStrokeWidth(MarkerWidth)

		// Links unten
		gc.MoveTo(x, y-10.0)
		gc.LineTo(x, y)
		gc.LineTo(x+10.0, y)
		// Links oben
		gc.MoveTo(x+10.0, y-FontSize)
		gc.LineTo(x, y-FontSize)
		gc.LineTo(x, y-FontSize+10.0)
		// Rechts unten
		gc.MoveTo(x+w-10.0, y)
		gc.LineTo(x+w, y)
		gc.LineTo(x+w, y-10.0)
		// Rechts oben
		gc.MoveTo(x+w-10.0, y-FontSize)
		gc.LineTo(x+w, y-FontSize)
		gc.LineTo(x+w, y-FontSize+10.0)

		gc.Stroke()
	}
	gc.SavePNG(FileName)
}
