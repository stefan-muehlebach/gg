package main

import (
	"github.com/stefan-muehlebach/gg"
	"github.com/stefan-muehlebach/gg/colors"
	"github.com/stefan-muehlebach/gg/geom"
)

const (
	Width, Height = 512.0, 512.0
	MarginSize    = 24.0
	NumCells      = 19
	CellSize      = (Width - 2*MarginSize) / NumCells
)

var (
	BackColor = colors.RGBAF{0.851, 0.811, 0.733, 1.0}

	color1 = colors.Crimson
	color2 = colors.Gold
	color3 = colors.LightSeaGreen
	color4 = colors.Black

	//color1 = colors.Red
	//color2 = colors.Yellow
	//color3 = colors.Cyan
	//color4 = colors.Black

	// color1 = colors.Cyan
	// color2 = colors.Yellow
	// color3 = colors.Magenta
	// color4 = colors.Black
)

func DrawCell(gc *gg.Context, pos, size geom.Point, tx, ty float64) {
	x1 := tx * size.X
	x2 := (1.0 - tx) * size.X
	y1 := ty * size.Y
	y2 := (1.0 - ty) * size.Y

	gc.SetFillColor(color1)
	gc.DrawRectangle(pos.X, pos.Y, x1, y1)
	gc.Fill()

	p := pos.AddXY(x1, 0.0)
	gc.SetFillColor(color2)
	gc.DrawRectangle(p.X, p.Y, x2, y1)
	gc.Fill()

	p = pos.AddXY(0.0, y1)
	gc.SetFillColor(color3)
	gc.DrawRectangle(p.X, p.Y, x1, y2)
	gc.Fill()

	p = pos.AddXY(x1, y1)
	gc.SetFillColor(color4)
	gc.DrawRectangle(p.X, p.Y, x2, y2)
	gc.Fill()

	// p1 := pos.Interpolate(pos.Add(size), tx)
	// p2 := p1
	// p1.Y = pos.Y
	// p2.Y = p1.Y+size.Y
	// gc.DrawLine(p1.X, p1.Y, p2.X, p2.Y)

	// p1 = pos.Interpolate(pos.Add(size), ty)
	// p2 = p1
	// p1.X = pos.X
	// p2.X = p1.X+size.X
	// gc.DrawLine(p1.X, p1.Y, p2.X, p2.Y)

	// gc.Stroke()
}

func main() {
	gc := gg.NewContext(Width, Height)
	gc.SetFillColor(BackColor)
	gc.Clear()

	dt := 1.0 / float64(NumCells+1)
	tx, ty := 0.0, 0.0
	size := geom.Point{CellSize, CellSize}
	for row := 0; row < NumCells; row++ {
		ty = 1.0 - dt*float64(row+1)
		for col := 0; col < NumCells; col++ {
			pos := geom.Point{float64(col)*CellSize + MarginSize, float64(row)*CellSize + MarginSize}
			tx = 1.0 - dt*float64(col+1)
			DrawCell(gc, pos, size, tx, ty)
		}
	}
	gc.SavePNG("color-cells.png")
}
