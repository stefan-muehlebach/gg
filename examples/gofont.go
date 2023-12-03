package main

import (
	"log"

	"github.com/golang/freetype/truetype"
	"github.com/stefan-muehlebach/gg"
	"golang.org/x/image/font/gofont/goregular"
)

const (
	outFile = "gofont.png"
)

func main() {
	font, err := truetype.Parse(goregular.TTF)
	if err != nil {
		log.Fatal(err)
	}

	face := truetype.NewFace(font, &truetype.Options{Size: 48})

	dc := gg.NewContext(1024, 1024)
	dc.SetFontFace(face)
	dc.SetFillColor(gg.NewRGB(1, 1, 1))
	dc.Clear()
	dc.SetStrokeColor(gg.NewRGB(0, 0, 0))
	dc.DrawStringAnchored("Hello, world!", 512, 512, 0.5, 0.5)
	dc.SavePNG(outFile)
}
