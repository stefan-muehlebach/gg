package main

import (
	"image/color"
	_ "log"
	"math"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"

	"github.com/stefan-muehlebach/gg"
)

const (
	sampleText = "Hello, world!"
	outFile    = "transform.png"
	width      = 800
	height     = 800
)

var (
	dc          *gg.Context
	face, face2 font.Face
)

func drawRect(dc *gg.Context, mx, my, w, h float64) {
	// Halbtransparentes Rechteck zeichnen.
	dc.DrawRectangle(mx-w/2.0, my-h/2.0, w, h)
	dc.SetLineWidth(10.0)
	dc.SetFillColor(gg.NewRGBA(0, 0, 0, 0.5))
	dc.SetStrokeColor(color.White)
	dc.FillStroke()

	// Schrift laden und Text im Zentrum des Rechtecks zeichnen.
	dc.SetFontFace(face)
	dc.SetStrokeColor(color.Black)
	dc.DrawStringAnchored(sampleText, mx, my, 0.5, 0.5)

	// Auffaelliger Rotationspunkt zeichnen.
	dc.DrawPoint(mx, my, 5.0)
	dc.SetLineWidth(2.0)
	dc.SetFillColor(gg.NewRGBA(1, 0, 1, 0.5))
	dc.SetStrokeColor(gg.NewRGBA(1, 0, 1, 1))
	dc.FillStroke()
}

func drawGrid(dc *gg.Context, mainTick float64) {
	b := dc.Bounds()

	dc.SetLineWidth(1.0)
	dc.SetStrokeColor(color.Black)
	for tick := mainTick; (b.Min.X < -tick) || (tick < b.Max.X) || (b.Min.Y < -tick) || (tick < b.Max.Y); tick += mainTick {
		if b.Min.X < -tick {
			dc.DrawLine(-tick, b.Min.Y, -tick, b.Max.Y)
		}
		if tick < b.Max.X {
			dc.DrawLine(tick, b.Min.Y, tick, b.Max.Y)
		}
		if b.Min.Y < -tick {
			dc.DrawLine(b.Min.X, -tick, b.Max.X, -tick)
		}
		if tick < b.Max.Y {
			dc.DrawLine(b.Min.X, tick, b.Max.X, tick)
		}
	}
	dc.Stroke()

	dc.SetLineWidth(3.0)
	dc.SetStrokeColor(color.Black)
	dc.DrawLine(b.Min.X, 0, b.Max.X, 0)
	dc.DrawLine(0, b.Min.Y, 0, b.Max.Y)
	dc.Stroke()
}

func main() {
	font, err := truetype.Parse(goregular.TTF)
	if err != nil {
		panic("")
	}
	face = truetype.NewFace(font, &truetype.Options{
		Size: 60,
		DPI:  72,
	})
	face2 = truetype.NewFace(font, &truetype.Options{
		Size: 60,
		DPI:  72,
	})

	dc = gg.NewContext(width, height)

	// Zeichnen ohne Koordinatentransformationen.
	dc.Push()
	drawRect(dc, 3*width/8, 3*height/8, width/2, width/2)
	dc.Pop()

	// Zeichnen mit lokalen Koordinatentransformationen.
	// Hier stimmt die Welt noch komplett!!!
	dc.Push()
	dc.Translate(5*width/8, 3*height/8)
	dc.Scale(1.2, 1.2)
	dc.Rotate(math.Pi / 16.0)
	drawRect(dc, 0, 0, width/2, width/2)
	dc.Pop()

	// Zeichnung mit globaler Koordinatentransformation.
	dc.Push()
	dc.Translate(width/2, width/2)

	drawRect(dc, 0, 0, width/2, width/2)
	dc.Rotate(math.Pi / 16.0)
	drawRect(dc, 0, 0, width/2, width/2)
	dc.Pop()

	dc.SavePNG(outFile)
}
