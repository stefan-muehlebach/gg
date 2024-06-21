package main

import (
	"math"

	"github.com/stefan-muehlebach/gg"
	"github.com/stefan-muehlebach/gg/color"
)

type PlayerType int

const (
	Player1 PlayerType = iota
	Player2
)

var (
	ImageSize       = 256.0
	MarginSize      = 20.0
	SymbolSize      = 52.0
	SymbolPadding   = 15.0
	OuterFieldSize  = 67.0
	InnerFieldSize  = 82.0
	BackColor       = color.Beige
	LineColor       = color.DarkSlateGray
	Player1Color    = color.DarkGreen
	Player2Color    = color.DarkRed
	GridLineWidth   = 7.0
	SymbolLineWidth = 10.0
	PNGFileName     = "tictactoe03.png"

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

func DrawSymbol(gc *gg.Context, col, row int, player PlayerType) {
	x := MarginSize + SymbolSize/2 + float64(col)*(SymbolSize+2*SymbolPadding)
	y := MarginSize + SymbolSize/2 + float64(row)*(SymbolSize+2*SymbolPadding)
	dx := (SymbolSize / 2) * math.Sqrt(3) / 2
	switch player {
	case Player1:
		gc.SetStrokeColor(Player1Color)
		gc.SetStrokeWidth(SymbolLineWidth)
		gc.DrawLine(x-dx, y-dx, x+dx, y+dx)
		gc.DrawLine(x-dx, y+dx, x+dx, y-dx)
		gc.Stroke()
	case Player2:
		gc.SetStrokeColor(Player2Color)
		gc.SetStrokeWidth(SymbolLineWidth)
		gc.DrawCircle(x, y, SymbolSize/2)
		gc.Stroke()
	}
}

func main() {
	gc := gg.NewContext(int(ImageSize), int(ImageSize))
	gc.SetFillColor(BackColor)
	gc.Clear()

	DrawGrid(gc)

	DrawSymbol(gc, 0, 0, Player2)
	DrawSymbol(gc, 2, 2, Player2)

	DrawSymbol(gc, 0, 1, Player1)
	DrawSymbol(gc, 1, 1, Player1)
	DrawSymbol(gc, 2, 1, Player1)

	gc.SavePNG(PNGFileName)
}
