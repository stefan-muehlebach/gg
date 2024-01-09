package main

import (
    "github.com/stefan-muehlebach/gg"
    "github.com/stefan-muehlebach/gg/colornames"
)

func main() {
	gc := gg.NewContext(256, 256)                //* \label{src:newcontext}

	gc.DrawCircle(128.0, 88.0, 64.0)             //* \label{src:drawcirc}
	gc.SetFillColor(colornames.Red.Alpha(0.5))   //* \label{src:fillcolor}
	gc.Fill()                                    //* \label{src:fill}
	gc.DrawCircle(86.0, 160.0, 64.0)
	gc.SetFillColor(colornames.Green.Alpha(0.5))
	gc.Fill()
	gc.DrawCircle(170.0, 160.0, 64.0)
	gc.SetFillColor(colornames.Blue.Alpha(0.5))
	gc.Fill()

	gc.SavePNG("circles.png")                    //* \label{src:savepng}
}
