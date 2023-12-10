package main

import (
	"math/rand"
	"github.com/stefan-muehlebach/gg/fonts"
    	"github.com/stefan-muehlebach/gg"
	"github.com/stefan-muehlebach/gg/color"
	"github.com/stefan-muehlebach/gg/colornames"
)

const (
	Width, Height = 512.0, 512.0
    Message = "Wir sind verloren!"
)

type ConfigData struct {
    color color.Color
    dist, size float64
}

var (
    	BackColor = colornames.Teal.Dark(0.7)
    
    ConfigList = []ConfigData{
        { colornames.Powderblue, -192, 28 },
        { colornames.Cornsilk,   -128, 24 },
        { colornames.Peachpuff,   -96, 22 },
    }
)

func main() {
	gc := gg.NewContext(Width, Height)
	gc.SetFillColor(BackColor)
	gc.Clear()

    for _, conf := range ConfigList {
        gc.SetFontFace(fonts.NewFace(fonts.GoRegular, conf.size))
        gc.Identity()
        gc.Translate(Width/2, Height/2)    
        gc.SetStrokeColor(conf.color.Alpha(0.5))
        for i:=0; i<60; i++ {
            angle := 0.6*rand.Float64()
            scale := 0.9 + 0.15*rand.Float64()
            gc.Rotate(angle)
            gc.Scale(scale, scale)
            gc.DrawStringAnchored(Message, 0.0, conf.dist, 0.5, 0.5)
        }
    }
	gc.SavePNG("coord-transf.png")
}
