package main

import (
	"github.com/stefan-muehlebach/gg"
	"github.com/stefan-muehlebach/gg/colornames"
)

var (
	ImageSize      = 256.0
	MarginSize     = 20.0
	OuterFieldSize = 67.0
	InnerFieldSize = 82.0

	BackColor     = colornames.Beige
	LineColor     = colornames.DarkSlateGray
	Player1Color  = colornames.DarkGreen
	Player2Color  = colornames.DarkRed
	GridLineWidth = 7.0
	IconLineWidth = 10.0
	PNGFileName   = "tictactoe02.png"

	GridPos1 = MarginSize + OuterFieldSize
	GridPos2 = GridPos1 + InnerFieldSize
)

func DrawGrid(gc *gg.Context) {
	gc.SetStrokeColor(LineColor)
	gc.SetStrokeWidth(GridLineWidth)
	gc.DrawLine(MarginSize, GridPos1, ImageSize-MarginSize, GridPos1)
	gc.DrawLine(MarginSize, GridPos2, ImageSize-MarginSize, GridPos2)
	gc.DrawLine(GridPos1, MarginSize, GridPos1, ImageSize-MarginSize)
	gc.DrawLine(GridPos2, MarginSize, GridPos2, ImageSize-MarginSize)
	gc.Stroke()
}

func main() {
	gc := gg.NewContext(int(ImageSize), int(ImageSize))
	gc.SetFillColor(BackColor)
	gc.Clear()

	DrawGrid(gc)

	gc.SetStrokeColor(Player2Color)
	gc.SetStrokeWidth(IconLineWidth)
	gc.DrawCircle(46, 46, 26)
	gc.DrawCircle(210, 210, 26)
	gc.Stroke()

	gc.SetStrokeColor(Player1Color)
	gc.DrawLine(106, 106, 150, 150)
	gc.DrawLine(106, 150, 150, 106)
	gc.DrawLine(24, 106, 68, 150)
	gc.DrawLine(24, 150, 68, 106)
	gc.DrawLine(188, 106, 232, 150)
	gc.DrawLine(188, 150, 232, 106)
	gc.Stroke()

	gc.SavePNG(PNGFileName)
}
