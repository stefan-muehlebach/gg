package main

import (
	"math/rand"

	"github.com/stefan-muehlebach/gg"
	"github.com/stefan-muehlebach/gg/colors"
	"github.com/stefan-muehlebach/gg/fonts"
)

const (
	Width, Height = 512.0, 512.0
	Message       = "Lorem ipsum"
	OutFileName   = "text-transform.png"
)

type ConfigData struct {
	color      colors.Color
	dist, size float64
}

var (
	BackColor  = colors.RGBAF{0.851, 0.811, 0.733, 1.0}
	TextFont   = fonts.LucidaBrightDemibold
	ConfigList = []ConfigData{
		{colors.Blue.Dark(0.5).Alpha(0.5), -100.0, 15.0},
		{colors.Green.Dark(0.7).Alpha(0.4), -220.0, 24.0},
		{colors.Red.Dark(0.7).Alpha(0.3), -300.0, 35.0},
	}
)

func main() {
	gc := gg.NewContext(Width, Height)
	gc.SetFillColor(BackColor)
	gc.Clear()

	for _, conf := range ConfigList {
		face, _ := fonts.NewFace(TextFont, conf.size)
		gc.SetFontFace(face)
		gc.Identity()
		gc.Translate(Width/2, Height/2)
		gc.SetTextColor(conf.color)
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
