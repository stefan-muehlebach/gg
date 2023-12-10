package main

import (
    "log"
	"github.com/stefan-muehlebach/gg"
	"github.com/stefan-muehlebach/gg/colornames"
	"github.com/stefan-muehlebach/gg/fonts"
)

var (
	Width      = 2048.0
	Height     = 4096.0
	bufferSize = 10
)

func main() {
	fontSize := 96.0
	margin := 40.0
	lineSize := fontSize + margin/2.0
	textColor := colornames.White.Alpha(0.7)
	// textColor.A = 0.7

	Height = float64(len(fonts.Names))*lineSize + 2*margin

	gc := gg.NewContext(int(Width), int(Height))
	gc.SetFillColor(colornames.Black)
	gc.Clear()
	for i, fontName := range fonts.Names {
		// col := i / maxRows
		// row := i % maxRows
		x := margin
		y := margin/2.0 + float64(i+1)*lineSize
		face := fonts.NewFace(fonts.Map[fontName], fontSize)
		gc.SetFontFace(face)
		gc.SetStrokeColor(textColor)
		gc.DrawString(fontName, x, y)

        w, h := gc.MeasureString(fontName)
        log.Printf("w, h: %f, %f\n", w, h)

		gc.SetStrokeColor(colornames.Lightyellow)
		gc.SetStrokeWidth(2.0)
        
        // Links unten
		gc.MoveTo(x, y-10.0)
		gc.LineTo(x, y)
		gc.LineTo(x+10.0, y)
		// Links oben
		gc.MoveTo(x+10.0, y-fontSize)
		gc.LineTo(x, y-fontSize)
		gc.LineTo(x, y-fontSize+10.0)
        // Rechts unten
		gc.MoveTo(x+w-10.0, y)
		gc.LineTo(x+w, y)
		gc.LineTo(x+w, y-10.0)
        // Rechts oben
		gc.MoveTo(x+w-10.0, y-fontSize)
		gc.LineTo(x+w, y-fontSize)
		gc.LineTo(x+w, y-fontSize+10.0)
            
		gc.Stroke()
	}
	gc.SavePNG("fontmap.png")
}
