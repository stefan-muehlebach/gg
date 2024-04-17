package main

import (
	"github.com/stefan-muehlebach/gg"
	"github.com/stefan-muehlebach/gg/colornames"
)

func main() {
	gc := gg.NewContext(256, 256)
	gc.SetFillColor(colornames.Beige)
	gc.Clear()

	gc.SetStrokeColor(colornames.DarkSlateGray)
	gc.SetStrokeWidth(7.0)
	gc.DrawLine(20, 87, 236, 87)
	gc.DrawLine(20, 169, 236, 169)
	gc.DrawLine(87, 20, 87, 236)
	gc.DrawLine(169, 20, 169, 236)
	gc.Stroke()

	gc.SetStrokeColor(colornames.DarkRed)
	gc.SetStrokeWidth(10.0)
	gc.DrawCircle(46, 46, 26)
	gc.DrawCircle(210, 210, 26)
	gc.Stroke()

	gc.SetStrokeColor(colornames.DarkGreen)
	gc.DrawLine(106, 106, 150, 150)
	gc.DrawLine(106, 150, 150, 106)
	gc.DrawLine(24, 106, 68, 150)
	gc.DrawLine(24, 150, 68, 106)
	gc.DrawLine(188, 106, 232, 150)
	gc.DrawLine(188, 150, 232, 106)
	gc.Stroke()

	gc.SavePNG("tictactoe01.png")
}
