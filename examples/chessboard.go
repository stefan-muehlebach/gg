package main

import (
    "github.com/stefan-muehlebach/gg"
    "github.com/stefan-muehlebach/gg/colornames"
)

const ( 
    Width, Height = 512.0, 512.0
    MarginSize    = 16.0
    FieldSize     = (Width - 2*MarginSize) / 8.0
)

var (
    BackColor   = colornames.Midnightblue
    FirstColor  = colornames.Whitesmoke
    SecondColor = colornames.Black
)

func main() {
    gc := gg.NewContext(Width, Height)
    gc.SetFillColor(BackColor)
    gc.Clear()
    for row:=0; row<8; row++ {
        y := MarginSize + float64(row)*FieldSize
        for col:=0; col<8; col++ {
            x := MarginSize + float64(col)*FieldSize
            
            if (row+col) % 2 == 0 {
                gc.SetFillColor(FirstColor)
            } else {
                gc.SetFillColor(SecondColor)
            }
            gc.DrawRectangle(x, y, FieldSize, FieldSize)
            gc.Fill()
        }
    }
    gc.SavePNG("chessboard.png")
}
