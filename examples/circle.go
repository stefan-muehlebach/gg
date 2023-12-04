package main

import (
	"github.com/stefan-muehlebach/gg"
	"github.com/stefan-muehlebach/gg/colornames"
)

func main() {
	gc := gg.NewContext(512, 512)      //* \label{src:newcontext}
	gc.DrawCircle(256.0, 256.0, 224.0) //* \label{src:drawcirc}
	gc.SetFillColor(colornames.Black)  //* \label{src:fillcolor}
	gc.Fill()                          //* \label{src:fill}
	gc.SavePNG("circle.png")           //* \label{src:savepng}
}