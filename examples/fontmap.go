package main

import (
	"fmt"

	"github.com/stefan-muehlebach/gg"
	"github.com/stefan-muehlebach/gg/colors"
	"github.com/stefan-muehlebach/gg/fonts"
)

var (
	FileName = "fontmap.png"

	MarginSize = 16.0
	FontSize   = 36.0
	LineHeight = 1.5 * FontSize
	LineWidth  = 1024.0
	BackColor  = colors.RGBAF{0.851, 0.811, 0.733, 1.0}
	TextColor  = colors.Black.Alpha(0.7)

	MarkerColor = colors.Crimson
	MarkerWidth = 2.0

	NumRows = 30
	NumCols = (len(fonts.Names) / NumRows) + 1

	Width  = float64(NumCols)*LineWidth + float64(NumCols+1)*MarginSize
	Height = float64(NumRows-1)*LineHeight + FontSize + 2*MarginSize
)

func main() {
	gc := gg.NewContext(int(Width), int(Height))
	gc.SetFillColor(BackColor)
	gc.Clear()
	gc.SetTextColor(TextColor)
	gc.SetStrokeColor(MarkerColor)
	gc.SetStrokeWidth(MarkerWidth)

	for i, fontName := range fonts.Names {
		col, row := i/NumRows, i%NumRows
		x := float64(col)*(MarginSize+LineWidth) + MarginSize
		y := float64(row)*(LineHeight) + MarginSize + FontSize
		face, err := fonts.NewFace(fonts.Map[fontName], FontSize)
		if err != nil {
			fmt.Printf("Couldn't create face for font '%s'\n", fontName)
			continue
		}
		gc.SetFontFace(face)
		gc.DrawString(fontName, x, y)

		w, _ := gc.MeasureString(fontName)

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
