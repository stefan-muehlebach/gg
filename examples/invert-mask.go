package main

import (
    "image/color"
    "github.com/stefan-muehlebach/gg"
)

const (
    outFile = "invert-mask.png"
)

func main() {
	dc := gg.NewContext(1024, 1024)
	dc.DrawCircle(512, 512, 384)
	dc.Clip()
	dc.InvertMask()
	dc.DrawRectangle(0, 0, 1024, 1024)
	dc.SetFillColor(color.Black)
	dc.Fill()
	dc.SavePNG(outFile)
}
