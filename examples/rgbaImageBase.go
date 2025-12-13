package main

import (
	"fmt"
	"image"
	"image/color"

	"github.com/stefan-muehlebach/gg"
	"github.com/stefan-muehlebach/gg/fonts"
)

const (
	outFile = "rgbaImageBase.png"
)

func main() {
	min := image.Point{128, 128}
	size := image.Point{256, 256}
	max := min.Add(size)
	rect := image.Rectangle{min, max}
	img := image.NewRGBA(rect)
	fmt.Printf("img.Bounds(): %v\n", img.Bounds())
	fmt.Printf("img.Bounds().Size(): %v\n", img.Bounds().Size())
	gc := gg.NewContextForRGBA(img)
	fmt.Printf("gc.Bounds() : %v\n", gc.Bounds())

	gc.SetFillColor(color.RGBA{0x7F, 0x00, 0x00, 0x7F})
	gc.SetStrokeColor(color.RGBA{0x00, 0x7F, 0x00, 0x7F})
	gc.DrawRectangle(256-64, 256-64, 128, 128)
	gc.SetStrokeWidth(8.0)
	gc.FillStroke()

	gc.SetStrokeColor(color.RGBA{0x7F, 0x00, 0xFF, 0xFF})
	gc.SetStrokeWidth(16.0)
	gc.DrawLine(192, 160, 320, 160)
	gc.DrawLine(160, 192, 160, 320)
	gc.Stroke()

	font := fonts.NewFace(fonts.GoRegular, 64.0)
	gc.SetFontFace(font)
	gc.SetTextColor(color.RGBA{A: 0xFF})
	gc.DrawStringAnchored("LowerRight", 256, 256, 1.0, 0.0)
	gc.DrawStringAnchored("UpperLeft", 256, 256, 0.0, 1.0)

	for _, pt := range []image.Point{{256, 256}, {192, 192}, {320, 320}} {
		for i := range 32 {
			col := pt.X + (i - 16)
			row := pt.Y + (i - 16)
			img.SetRGBA(col, pt.Y, color.RGBA{R: 0xFF, A: 0xFF})
			img.SetRGBA(pt.X, row, color.RGBA{R: 0xFF, A: 0xFF})
		}
	}

	gc.SavePNG(outFile)
}
