package main

import (
	"log"

	"mju.net/gg"
)

const (
    outFile = "mask.png"
)

func main() {
	im, err := gg.LoadImage("baboon.png")
	if err != nil {
		log.Fatal(err)
	}

	dc := gg.NewContext(512, 512)
	dc.DrawRoundedRectangle(0, 0, 512, 512, 64)
	dc.Clip()
	dc.DrawImage(im, 0, 0)
	dc.SavePNG(outFile)
}
