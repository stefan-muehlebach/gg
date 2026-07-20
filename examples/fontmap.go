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

	MarginSize   = 16.0
	FontSize     = 36.0
	DecoFontSize = 3.0 * FontSize
	LineWidth    = 1024.0
	BackColor    = colors.RGBAF{0.851, 0.811, 0.733, 1.0}
	TextColor    = colors.Black

	NumberFont     = fonts.GoMono
	NumberFontSize = 18.0

	MarkerColor = colors.Red
	MarkerWidth = 2.0
    MarkerLen   = 10.0

	NumRows = 0
	NumCols = 3

	Width, Height float64
)

func main() {

	DecoFontsHeight := 0.0
	numRegularFonts, numDecoFonts := 0, 0
	for _, fontName := range fonts.Names {
		font := fonts.Map[fontName]
		if font.Id < 900 {
			numRegularFonts++
		} else {
			numDecoFonts++
			face, err := fonts.NewFace(font, DecoFontSize)
			if err != nil {
				log.Fatalf("couldn't create face for font '%s': %v", fontName, err)
			}
			metr := face.Metrics()
			DecoFontsHeight += math.Abs(fix2flt(metr.CapHeight)) + MarginSize
		}
	}

	if NumCols != 0 && NumRows == 0 {
		NumRows = numRegularFonts/NumCols + 1
	} else if NumCols == 0 && NumRows != 0 {
		NumCols = numRegularFonts/NumRows + 1
	} else {
		log.Fatal("Either NumCols or NumRows must be zero!")
	}

	Width = float64(NumCols)*LineWidth + float64(NumCols+1)*MarginSize
	Height = float64(NumRows)*FontSize + float64(NumRows-1)*MarginSize + 2*MarginSize
	Height += DecoFontsHeight

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

	var col, row int
	var x, y float64
	var fontSize float64
	var isDecoFont bool

	for i, fontName := range fonts.Names {
		font := fonts.Map[fontName]
		if font.Id < 900 {
			col, row = i/NumRows, i%NumRows
			x = float64(col)*(MarginSize+LineWidth) + MarginSize
			if row == 0 {
				y = 0.0
			}
			fontSize = FontSize
		} else {
			x = MarginSize
			if !isDecoFont {
				isDecoFont = true
				y = float64(NumRows) * (FontSize + MarginSize)
			}
			fontSize = DecoFontSize
		}

		face, err := fonts.NewFace(font, fontSize)
		if err != nil {
			log.Fatalf("couldn't create face for font '%s': %v", fontName, err)
		}
		metr := face.Metrics()

		log.Printf("%s: %+v", fontName, metr)
		// height := fix2flt(metr.Height)
		capHeight := math.Abs(fix2flt(metr.CapHeight))
		// log.Printf("height: %f, capHeight: %f", height, capHeight)

		if font.Id < 900 {
			y += MarginSize + fontSize
		} else {
			y += MarginSize + capHeight
		}
		//y += MarginSize + capHeight
		log.Printf("y: %f", y)

		str := fmt.Sprintf("%03d", font.Id)
		gc.SetFontFace(numberFace)
		gc.DrawString(str, x, y)
		x += 1.5 * float64(numberWidth)

		gc.SetFontFace(face)
		gc.DrawString(fontName, x, y)

		w, _ := gc.MeasureString(fontName)

		// Links unten
		gc.MoveTo(x, y-MarkerLen)
		gc.LineTo(x, y)
		gc.LineTo(x+MarkerLen, y)
		// Links oben
		gc.MoveTo(x+MarkerLen, y-capHeight)
		gc.LineTo(x, y-capHeight)
		gc.LineTo(x, y-capHeight+MarkerLen)
		// Rechts unten
		gc.MoveTo(x+w-MarkerLen, y)
		gc.LineTo(x+w, y)
		gc.LineTo(x+w, y-MarkerLen)
		// Rechts oben
		gc.MoveTo(x+w-MarkerLen, y-capHeight)
		gc.LineTo(x+w, y-capHeight)
		gc.LineTo(x+w, y-capHeight+MarkerLen)

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
