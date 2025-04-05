package main

import (
	"fmt"
	"github.com/stefan-muehlebach/gg"
	"github.com/stefan-muehlebach/gg/colors"
	"github.com/stefan-muehlebach/gg/fonts"
	"github.com/stefan-muehlebach/gg/geom"
	"math"
)

const (
	Width, Height = 512.0, 512.0
	MarginSize    = Width / 32.0
	CanvasSize    = Width - 2*MarginSize
	BorderSize    = 20.0
	FieldSize     = (Width - 2*MarginSize - 2*BorderSize) / 8.0
	FontSize      = 11.0
	LineWidth     = 1.5
)

var (
	BackColor  = colors.RGBAF{0.851, 0.811, 0.733, 1.0}
	LineColor  = colors.RGBAF{0.153, 0.157, 0.133, 1.0}
	WhiteColor = colors.Ivory
	BlackColor = colors.RGBAF{0.153, 0.157, 0.133, 1.0}
	TextColor  = colors.RGBAF{0.153, 0.157, 0.133, 1.0}
)

func main() {
	gc := gg.NewContext(Width, Height)
	gc.SetFillColor(BackColor)
	gc.Clear()
	for row := 0; row < 8; row++ {
		y := MarginSize + BorderSize + float64(row)*FieldSize
		for col := 0; col < 8; col++ {
			x := MarginSize + BorderSize + float64(col)*FieldSize

			if (row+col)%2 == 0 {
				gc.SetFillColor(WhiteColor)
			} else {
				gc.SetFillColor(BlackColor)
			}
			gc.DrawRectangle(x, y, FieldSize, FieldSize)
			gc.Fill()
		}
	}
	face := fonts.NewFace(fonts.LucidaHandwritingItalic, FontSize)
	gc.SetFontFace(face)
	gc.SetStrokeColor(TextColor)
	gc.SetStrokeWidth(LineWidth)
	gc.DrawRectangle(MarginSize, MarginSize, CanvasSize, CanvasSize)
	gc.Stroke()

	pos := geom.Point{MarginSize + BorderSize + FieldSize/2.0, Height - MarginSize - BorderSize/2.0}
	for k := 0; k < 2; k++ {
		p1 := pos.AddXY(-FieldSize/2.0, -5.0)
		p2 := p1.AddXY(0.0, 10.0)
		gc.DrawLine(p1.X, p1.Y, p2.X, p2.Y)
		gc.Stroke()
		for i := 0; i < 8; i++ {
			p := pos.AddXY(float64(i)*FieldSize, 0.0)
			lbl := fmt.Sprintf("%c", 'A'+i)
			gc.DrawStringAnchored(lbl, p.X, p.Y, 0.5, 0.5)
			p1 = p.AddXY(FieldSize/2.0, -5.0)
			p2 = p1.AddXY(0.0, 10.0)
			gc.DrawLine(p1.X, p1.Y, p2.X, p2.Y)
			gc.Stroke()
		}
		gc.RotateAbout(math.Pi, Width/2, Height/2)
	}
	gc.RotateAbout(math.Pi/2.0, Width/2, Height/2)
	for k := 0; k < 2; k++ {
		p1 := pos.AddXY(-FieldSize/2.0, -5.0)
		p2 := p1.AddXY(0.0, 10.0)
		gc.DrawLine(p1.X, p1.Y, p2.X, p2.Y)
		gc.Stroke()
		for i := 0; i < 8; i++ {
			p := pos.AddXY(float64(i)*FieldSize, 0.0)
			lbl := fmt.Sprintf("%d", i+1)
			gc.DrawStringAnchored(lbl, p.X, p.Y, 0.5, 0.5)
			p1 = p.AddXY(FieldSize/2.0, -5.0)
			p2 = p1.AddXY(0.0, 10.0)
			gc.DrawLine(p1.X, p1.Y, p2.X, p2.Y)
			gc.Stroke()
		}
		gc.RotateAbout(math.Pi, Width/2, Height/2)
	}
	gc.SavePNG("chessboard.png")
}
