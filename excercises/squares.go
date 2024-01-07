package main

import (
	"github.com/stefan-muehlebach/gg"
	"github.com/stefan-muehlebach/gg/colornames"
)

const (
	ImageSize  = 256.0
	MarginSize = ImageSize / 20
	CanvasSize = ImageSize - 2*MarginSize
    FieldSize  = CanvasSize / 3
    LineWidth  = 5.0

	OutFileName = "squares.png"
)

var (
	BackColor = colornames.Beige
	LineColor = colornames.DarkSlateGrey
)

func main() {
	gc := gg.NewContext(ImageSize, ImageSize)
	gc.SetFillColor(BackColor)
	gc.Clear()

    gc.SetStrokeColor(LineColor)
    gc.SetStrokeWidth(LineWidth)

    gc.DrawLine(MarginSize, MarginSize+FieldSize, MarginSize+CanvasSize, MarginSize+FieldSize)
    gc.DrawLine(MarginSize, MarginSize+2*FieldSize, MarginSize+CanvasSize, MarginSize+2*FieldSize)
    gc.DrawLine(MarginSize+FieldSize, MarginSize, MarginSize+FieldSize, MarginSize+CanvasSize)
    gc.DrawLine(MarginSize+2*FieldSize, MarginSize, MarginSize+2*FieldSize, MarginSize+CanvasSize)
    gc.Stroke()

    gc.SetStrokeColor(colornames.DarkRed)
    gc.SetStrokeWidth(2*LineWidth)
    gc.DrawCircle(MarginSize+FieldSize/2, MarginSize+FieldSize/2, FieldSize/3)
    gc.Stroke()

    gc.SetStrokeColor(colornames.DarkGreen)
    gc.SetStrokeWidth(2*LineWidth)
    gc.DrawLine(ImageSize/2-FieldSize/3, ImageSize/2-FieldSize/3, ImageSize/2+FieldSize/3, ImageSize/2+FieldSize/3)
    gc.DrawLine(ImageSize/2-FieldSize/3, ImageSize/2+FieldSize/3, ImageSize/2+FieldSize/3, ImageSize/2-FieldSize/3)
    gc.Stroke()

	gc.SavePNG(OutFileName)
}
