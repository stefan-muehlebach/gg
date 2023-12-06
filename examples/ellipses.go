package main

import "github.com/stefan-muehlebach/gg"

func main() {
	const S = 1024
	dc := gg.NewContext(S, S)
	dc.SetFillColor(gg.NewRGBA(0.0, 0.0, 0.0, 0.2))
	for i := 0; i < 360; i += 15 {
		dc.Push()
		dc.RotateAbout(gg.Radians(float64(i)), S/2, S/2)
		dc.DrawEllipse(S/2, S/2, S*7/16, S/8)
		dc.Fill()
		dc.Pop()
	}
	dc.SavePNG("ellipses.png")
}
