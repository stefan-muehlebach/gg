package main

import (
	"math"

	"github.com/stefan-muehlebach/gg"
	"github.com/stefan-muehlebach/gg/colornames"
)

type PlayerType int

const (
	Player1 PlayerType = iota
	Player2
)

var (
	ImageSize      = 256.0
	MarginSize     = 20.0
	IconSize       = 52.0
	IconPadding    = 15.0
	OuterFieldSize = 67.0
	InnerFieldSize = 82.0
	BackColor      = colornames.Beige
	LineColor      = colornames.DarkSlateGray
	Player1Color   = colornames.DarkGreen
	Player2Color   = colornames.DarkRed
	GridLineWidth  = 7.0
	IconLineWidth  = 10.0
	PNGFileName    = "tictactoe.png"

	gc       *gg.Context
	GridPos1 = MarginSize + OuterFieldSize
	GridPos2 = GridPos1 + InnerFieldSize
)

func DrawGrid() {
	gc.SetStrokeColor(LineColor)
	gc.SetStrokeWidth(GridLineWidth)
	gc.DrawLine(MarginSize, GridPos1, ImageSize-MarginSize, GridPos1)
	gc.DrawLine(MarginSize, GridPos2, ImageSize-MarginSize, GridPos2)
	gc.DrawLine(GridPos1, MarginSize, GridPos1, ImageSize-MarginSize)
	gc.DrawLine(GridPos2, MarginSize, GridPos2, ImageSize-MarginSize)
	gc.Stroke()
}

func DrawIcon(col, row int, player PlayerType) {
	x := MarginSize + IconSize/2 + float64(col)*(IconSize+2*IconPadding)
	y := MarginSize + IconSize/2 + float64(row)*(IconSize+2*IconPadding)
	dx := (IconSize / 2) * math.Sqrt(3) / 2
	switch player {
	case Player1:
		gc.SetStrokeColor(Player1Color)
		gc.SetStrokeWidth(IconLineWidth)
		gc.DrawLine(x-dx, y-dx, x+dx, y+dx)
		gc.DrawLine(x-dx, y+dx, x+dx, y-dx)
		gc.Stroke()
	case Player2:
		gc.SetStrokeColor(Player2Color)
		gc.SetStrokeWidth(IconLineWidth)
		gc.DrawCircle(x, y, IconSize/2)
		gc.Stroke()
	}
}

func main() {
	gc = gg.NewContext(int(ImageSize), int(ImageSize))
	gc.SetFillColor(BackColor)
	gc.Clear()

	DrawGrid()

	DrawIcon(0, 0, Player2)
	DrawIcon(2, 0, Player2)
	DrawIcon(2, 2, Player2)

	DrawIcon(0, 1, Player1)
	DrawIcon(1, 1, Player1)
	DrawIcon(2, 1, Player1)

	gc.SavePNG(PNGFileName)
}
