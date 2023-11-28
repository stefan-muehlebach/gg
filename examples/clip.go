package main

import "github.com/stefan-muehlebach/gg"

const (
    outFile = "clip.png"
)

func main() {
	dc := gg.NewContext(1000, 1000)
	dc.DrawCircle(350, 500, 300)
	dc.Clip()
	dc.DrawCircle(650, 500, 300)
	dc.Clip()
	dc.DrawRectangle(0, 0, 1000, 1000)
	dc.SetFillColor(gg.NewRGBA(0, 0, 0, 0.5))
	dc.Fill()
	dc.SavePNG(outFile)
}
