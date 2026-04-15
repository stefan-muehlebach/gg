package main

import (
	"fmt"
	"log"
	"math"

	"golang.org/x/image/math/fixed"

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

	NumberFont     = fonts.GoMono
	NumberFontSize = 18.0

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

	numberFace, err := fonts.NewFace(NumberFont, NumberFontSize)
	if err != nil {
		log.Fatalf("couldn't create face for number font: %v", err)
	}
	gc.SetFontFace(numberFace)
	numberWidth, _ := gc.MeasureString("999")

    var y float64
	for i, fontName := range fonts.Names {
		col, row := i/NumRows, i%NumRows
		x := float64(col)*(MarginSize+LineWidth) + MarginSize
		font := fonts.Map[fontName]

        fontSize := FontSize
        if font.Id >= 900 {
            fontSize *= 2.0
        }
		face, err := fonts.NewFace(font, fontSize)
		if err != nil {
			log.Fatalf("couldn't create face for font '%s': %v", fontName, err)
		}
		metr := face.Metrics()

		log.Printf("%s: %+v", fontName, metr)
		if row == 0 {
			y = 0.0
		}
        // height := fix2flt(metr.Height)
        capHeight := math.Abs(fix2flt(metr.CapHeight))
		// log.Printf("height: %f, capHeight: %f", height, capHeight)

		y += MarginSize + capHeight
		log.Printf("y: %f", y)

		str := fmt.Sprintf("%03d", font.Id)
		gc.SetFontFace(numberFace)
		gc.DrawString(str, x, y)
		x += 1.5 * float64(numberWidth)

		gc.SetFontFace(face)
		gc.DrawString(fontName, x, y)

		w, _ := gc.MeasureString(fontName)

		// Links unten
		gc.MoveTo(x, y-10.0)
		gc.LineTo(x, y)
		gc.LineTo(x+10.0, y)
		// Links oben
		gc.MoveTo(x+10.0, y-capHeight)
		gc.LineTo(x, y-capHeight)
		gc.LineTo(x, y-capHeight+10.0)
		// Rechts unten
		gc.MoveTo(x+w-10.0, y)
		gc.LineTo(x+w, y)
		gc.LineTo(x+w, y-10.0)
		// Rechts oben
		gc.MoveTo(x+w-10.0, y-capHeight)
		gc.LineTo(x+w, y-capHeight)
		gc.LineTo(x+w, y-capHeight+10.0)

		gc.Stroke()
	}
	gc.SavePNG(FileName)
}

func flt2fix(x float64) fixed.Int26_6 {
	return fixed.Int26_6(math.Round(x * 64))
}

func fix2flt(x fixed.Int26_6) float64 {
	const shift, mask = 6, 1<<6 - 1
	if x >= 0 {
		return float64(x>>shift) + float64(x&mask)/64
	}
	x = -x
	if x >= 0 {
		return -(float64(x>>shift) + float64(x&mask)/64)
	}
	return 0
}
