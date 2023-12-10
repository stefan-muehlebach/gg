package main

import (
	"math"
    "fmt"
	"github.com/stefan-muehlebach/gg/geom"
	"github.com/stefan-muehlebach/gg/fonts"
	"github.com/stefan-muehlebach/gg"
	"github.com/stefan-muehlebach/gg/colornames"
)

const (
	Width, Height = 1024.0, 1024.0
	MarginSize    = 40.0
	FieldSize     = (Width - 2*MarginSize) / 8.0
)

var (
	BackColor   = colornames.Midnightblue.Dark(0.4)
	FirstColor  = colornames.Whitesmoke
	SecondColor = colornames.Black
)

func main() {
	gc := gg.NewContext(Width, Height)
	gc.SetFillColor(BackColor)
	gc.Clear()
	for row := 0; row < 8; row++ {
		y := MarginSize + float64(row)*FieldSize
		for col := 0; col < 8; col++ {
			x := MarginSize + float64(col)*FieldSize

			if (row+col)%2 == 0 {
				gc.SetFillColor(FirstColor)
			} else {
				gc.SetFillColor(SecondColor)
			}
			gc.DrawRectangle(x, y, FieldSize, FieldSize)
			gc.Fill()
		}
	}
    face := fonts.NewFace(fonts.LucidaHandwritingItalic, 22.0)
    gc.SetFontFace(face)
    gc.SetStrokeColor(colornames.Gold)
    gc.SetStrokeWidth(2.0)
    
    pos := geom.Point{MarginSize+FieldSize/2.0, Height-MarginSize/2.0}
    for k:=0; k<2; k++ {
        p1 := pos.AddXY(-FieldSize/2.0, -10.0)
        p2 := p1.AddXY(0.0, 20.0)
        gc.DrawLine(p1.X, p1.Y, p2.X, p2.Y)
        gc.Stroke()
        for i:=0; i<8; i++ {
            p := pos.AddXY(float64(i)*FieldSize, 0.0)
            lbl := fmt.Sprintf("%c", 'A'+i)
            gc.DrawStringAnchored(lbl, p.X, p.Y, 0.5, 0.5)
            p1 = p.AddXY(FieldSize/2.0, -10.0)
            p2 = p1.AddXY(0.0, 20.0)
            gc.DrawLine(p1.X, p1.Y, p2.X, p2.Y)
            gc.Stroke()            
        }
        gc.RotateAbout(math.Pi, Width/2, Height/2)
    }
    gc.RotateAbout(math.Pi/2.0, Width/2, Height/2)
    for k:=0; k<2; k++ {
        p1 := pos.AddXY(-FieldSize/2.0, -10.0)
        p2 := p1.AddXY(0.0, 20.0)
        gc.DrawLine(p1.X, p1.Y, p2.X, p2.Y)
        gc.Stroke()
        for i:=0; i<8; i++ {
            p := pos.AddXY(float64(i)*FieldSize, 0.0)
            lbl := fmt.Sprintf("%d", i+1)
            gc.DrawStringAnchored(lbl, p.X, p.Y, 0.5, 0.5)
            p1 = p.AddXY(FieldSize/2.0, -10.0)
            p2 = p1.AddXY(0.0, 20.0)
            gc.DrawLine(p1.X, p1.Y, p2.X, p2.Y)
            gc.Stroke()            
        }
        gc.RotateAbout(math.Pi, Width/2, Height/2)
    }
	gc.SavePNG("chessboard.png")
}
