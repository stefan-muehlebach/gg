package main

import (
	"fmt"
	// "fmt"
	"github.com/stefan-muehlebach/gg"
	"github.com/stefan-muehlebach/gg/color"
	"github.com/stefan-muehlebach/gg/geom"
	"image"
	gocolor "image/color"
	"log"
	"math"
	"os"
)

const (
	Width, Height = 512, 512
	MarginSize    = 24.0
	FieldSize     = Width - 2*MarginSize
	NumPointsX    = 100
	DefInFile     = "skull.png"
	DefOutFile    = "raster-image.png"
)

var (
	BackColor = color.RGBAF{0.851, 0.811, 0.733, 1.0}
	LineColor = color.RGBAF{0.153, 0.157, 0.133, 1.0}

	Sqrt3       = math.Sqrt(3.0)
	NumPointsY  int
	GapX, GapY  float64
	MinPointRad float64
	MaxPointRad float64
)

func GrayValue(img image.Image, col, row int) float64 {
	pixel := img.At(col, row)
	grayPixel := gocolor.Gray16Model.Convert(pixel).(gocolor.Gray16)
	return float64(grayPixel.Y) / float64(0xFFFF)
}

func AverageGrayValue(img image.Image, col0, row0, width, height int) float64 {
	s := 0.0
	for row := row0; row < row0+height; row++ {
		for col := col0; col < col0+width; col++ {
			s += GrayValue(img, col, row)
		}
	}
	return s / float64(width*height)
}

func RasterImage(gc *gg.Context, fileName string) {
	NumPointsY = NumPointsX
	GapX = FieldSize / float64(NumPointsX-1)
	GapY = FieldSize / float64(NumPointsX-1)
	MinPointRad = 0.5
	MaxPointRad = GapX * math.Sqrt2 / 2.0

	img, err := gg.LoadImage(fileName)
	if err != nil {
		log.Fatalf("couldn't load image: %v", err)
	}
	blockWidth := img.Bounds().Dx() / NumPointsX
	blockHeight := img.Bounds().Dy() / NumPointsY

	gc.SetFillColor(LineColor)
	for row := 0; row < NumPointsY; row++ {
		imgRow0 := row * blockHeight
		for col := 0; col < NumPointsX; col++ {
			imgCol0 := col * blockWidth
			t := AverageGrayValue(img, imgCol0, imgRow0, blockWidth, blockHeight)
			pointRad := MinPointRad*(t) + MaxPointRad*(1-t)
			mp := geom.Point{MarginSize + float64(col)*GapX, MarginSize + float64(row)*GapY}
			gc.DrawPoint(mp.X, mp.Y, pointRad)
			gc.Fill()
		}
	}
}

func main() {
	var inFileName, outFileName string

	switch len(os.Args) {
	case 1:
		inFileName = DefInFile
		outFileName = DefOutFile
	case 2:
		inFileName = os.Args[1]
		outFileName = fmt.Sprintf("raster-%s", inFileName)
	default:
		log.Fatalf("usage: %s [<file>]", os.Args[0])
	}
	gc := gg.NewContext(Width, Height)
	gc.SetFillColor(BackColor)
	gc.Clear()

	RasterImage(gc, inFileName)

	gc.SavePNG(outFileName)
}
