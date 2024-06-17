package main

import (
	"flag"
	"fmt"
	"image"
	gocolor "image/color"
	"log"
	"math"
	"os"

	"github.com/stefan-muehlebach/gg"
	"github.com/stefan-muehlebach/gg/color"
	"github.com/stefan-muehlebach/gg/geom"
)

type RasterFormat int

const (
	SquareRasterFormat RasterFormat = iota
	TriangleRasterFormat
	NumRasterFormats
)

const (
	Width, Height = 512, 512
	MarginSize    = 20
	FieldSize     = Width - 2*MarginSize
	DefInFile     = "skull.png"
	DefFormat     = SquareRasterFormat
)

var (
	BackColor = color.RGBAF{0.851, 0.811, 0.733, 1.0}
	LineColor = color.RGBAF{0.153, 0.157, 0.133, 1.0}

	Sqrt3                   = math.Sqrt(3.0)
	Sqrt3_2                 = math.Sqrt(3.0) / 2.0
	NumPointsX, NumPointsY  = 100, 0
	GapX, GapY              float64
	BlockWidth, BlockHeight float64
	MinPointRad             float64
	MaxPointRad             float64
)

// Das Pixel in img in der Spalte col und Zeile row wird als Graustufenwert
// im Interval [0,1] zurueckgegeben.
func GrayValue(img image.Image, col, row int) float64 {
	pixel := img.At(col, row)
	grayPixel := gocolor.Gray16Model.Convert(pixel).(gocolor.Gray16)
	return float64(grayPixel.Y) / float64(0xFFFF)
}

// Bestimmt den Mittelwert der Grauwerte im Bild img und zwar ab Koordinate
// (col0, row0) bis (col0+width, row0+heigt). Alle Pixel erhalten das gleiche
// Gewicht - eine gewichtete Bestimmung des Mittelwertes ist an- aber
// noch nicht zu Ende gedacht.
func AverageGrayValue(img image.Image, col0, row0, width, height int) float64 {
	s := 0.0
	for row := row0; row < row0+height; row++ {
		for col := col0; col < col0+width; col++ {
			s += GrayValue(img, col, row)
		}
	}
	return s / float64(width*height)
}

func RasterImage(gc *gg.Context, fileName string, format RasterFormat) {
	var RasterCoord func(col, row int) geom.Point
	var PixelCoord func(col, row int) image.Point

	img, err := gg.LoadImage(fileName)
	if err != nil {
		log.Fatalf("couldn't load image: %v", err)
	}

	switch format {
	case SquareRasterFormat:
		NumPointsY = NumPointsX
		GapX = FieldSize / float64(NumPointsX-1)
		GapY = FieldSize / float64(NumPointsY-1)
		MinPointRad = 0.0
		MaxPointRad = GapX * math.Sqrt2 / 2.0
		BlockWidth = float64(img.Bounds().Dx()) / float64(NumPointsX)
		BlockHeight = float64(img.Bounds().Dy()) / float64(NumPointsY)
		RasterCoord = func(col, row int) geom.Point {
			return geom.Point{MarginSize + float64(col)*GapX,
				MarginSize + float64(row)*GapY}
		}
		PixelCoord = func(col, row int) image.Point {
			return geom.Point{float64(col) * BlockWidth,
				float64(row) * BlockHeight}.Int()
		}

	case TriangleRasterFormat:
		NumPointsY = int(math.Ceil(float64(NumPointsX) / Sqrt3_2))
		GapX = FieldSize / float64(NumPointsX-1)
		GapY = FieldSize / float64(NumPointsY-1)
		MinPointRad = 0.0
		MaxPointRad = GapX / Sqrt3
		BlockWidth = float64(img.Bounds().Dx()) / (float64(NumPointsX) + 0.5)
		BlockHeight = float64(img.Bounds().Dy()) / float64(NumPointsY)
		RasterCoord = func(col, row int) geom.Point {
			return geom.Point{MarginSize + float64(col)*GapX +
				float64(row%2)*(GapX/2.0), MarginSize + float64(row)*GapY}
		}
		PixelCoord = func(col, row int) image.Point {
			return geom.Point{float64(col)*BlockWidth + float64(row%2)*(BlockWidth/2.0),
				float64(row) * BlockHeight}.Int()
		}
	}

	gc.SetFillColor(LineColor)
	for row := 0; row < NumPointsY; row++ {
		//y0 := float64(row) * BlockHeight
		for col := 0; col < NumPointsX; col++ {
			//x0 := float64(col) * BlockWidth
			pp := PixelCoord(col, row)
			rp := RasterCoord(col, row)
			t := AverageGrayValue(img, pp.X, pp.Y,
				int(BlockWidth), int(BlockHeight))
			pointRad := MinPointRad*(t) + MaxPointRad*(1-t)
			gc.DrawPoint(rp.X, rp.Y, pointRad)
			gc.Fill()
		}
	}
}

func main() {
	var inFileName, outFileName string
	var format RasterFormat

	flag.IntVar((*int)(&format), "format", int(DefFormat), "Raster Format")
	flag.Parse()

	switch len(os.Args) {
	case 1:
		inFileName = DefInFile
	case 2:
		inFileName = os.Args[1]
	default:
		log.Fatalf("usage: %s [<file>]", os.Args[0])
	}
	outFileName = fmt.Sprintf("raster-%s", inFileName)
	gc := gg.NewContext(Width, Height)
	gc.SetFillColor(BackColor)
	gc.Clear()
	RasterImage(gc, inFileName, format)
	gc.SavePNG(outFileName)
}
