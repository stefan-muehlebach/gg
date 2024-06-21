package main

import (
	"github.com/stefan-muehlebach/gg"
	"github.com/stefan-muehlebach/gg/color"
	"github.com/stefan-muehlebach/gg/fonts"
	"math/rand"
)

const (
	Width, Height = 512.0, 512.0
	Message       = "Lorem ipsum"
	OutFileName   = "text-transform.png"
)

type ConfigData struct {
	color      color.Color
	dist, size float64
}

var (
	BackColor  = color.RGBAF{0.851, 0.811, 0.733, 1.0}
	TextFont   = fonts.GoBold
	ConfigList = []ConfigData{
		{color.Blue.Dark(0.5).Alpha(0.5), -80, 18.0},
		{color.Green.Dark(0.7).Alpha(0.6), -250, 24.0},
		{color.Red.Dark(0.7).Alpha(0.7), -300, 32.0},
	}
)

func main() {
	gc := gg.NewContext(Width, Height)
	gc.SetFillColor(BackColor)
	gc.Clear()

	for _, conf := range ConfigList {
		gc.SetFontFace(fonts.NewFace(TextFont, conf.size))
		gc.Identity()
		gc.Translate(Width/2, Height/2)
		gc.SetStrokeColor(conf.color)
		for i := 0; i < 60; i++ {
			angle := 0.4 + 0.1*rand.NormFloat64()
			scale := 0.97 + 0.08*rand.NormFloat64()
			gc.Rotate(angle)
			gc.Scale(scale, scale)
			gc.DrawStringAnchored(Message, 0.0, conf.dist, 0.5, 0.5)
		}
	}
	gc.SavePNG(OutFileName)
}
