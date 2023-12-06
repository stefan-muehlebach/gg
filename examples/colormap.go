package main

import (
	"github.com/stefan-muehlebach/gg"
	"github.com/stefan-muehlebach/gg/color"
	"github.com/stefan-muehlebach/gg/colornames"
	"github.com/stefan-muehlebach/gg/fonts"
	"math"
	"regexp"
	"sort"
)

type NamedColor struct {
	name  string
	color color.Color
}

var (
	ColorList     []NamedColor
	NumFadeSteps  int     = 9
	FadeStep      float64 = 1.0 / float64(NumFadeSteps+1)
	Padding       float64 = 10
	SampleWidth   float64 = 250
	UniformHeight float64 = 50.0
	FadeHeight    float64 = 20.0
	FadeWidth     float64 = SampleWidth / float64(NumFadeSteps)
	SampleHeight  float64 = 2.0*FadeHeight + UniformHeight
)

func DrawColorSample(gc *gg.Context, col, row int, color color.Color, name string) {
	gc.SetFillColor(color)
	gc.DrawRectangle(float64(col)*(SampleWidth+Padding), float64(row)*(SampleHeight+Padding)+FadeHeight, SampleWidth, UniformHeight)
	gc.Fill()

	for l := 0; l < NumFadeSteps; l++ {
		t := FadeStep * float64(l+1)
		colBright := color.Bright(1.0 - t)
		colDark := color.Dark(t)
		x := float64(col)*(SampleWidth+Padding) + float64(l)*FadeWidth
		y := float64(row) * (SampleHeight + Padding)
		gc.SetFillColor(colBright)
		gc.DrawRectangle(x, y, FadeWidth, FadeHeight)
		gc.Fill()
		y += FadeHeight + UniformHeight
		gc.SetFillColor(colDark)
		gc.DrawRectangle(x, y, FadeWidth, FadeHeight)
		gc.Fill()
	}
	r, g, b, _ := color.RGBA()
	mean := (r + g + b) / 3
	if mean > 3*(math.MaxInt16/2) {
		gc.SetStrokeColor(colornames.Black)
	} else {
		gc.SetStrokeColor(colornames.Whitesmoke)
	}
	gc.DrawStringAnchored(name, float64(col)*(SampleWidth+Padding)+SampleWidth/2.0, float64(row)*(SampleHeight+Padding)+SampleHeight/2.0, 0.5, 0.5)
}

func colorMapRGBAF() {
	Columns := 5
	Rows := len(ColorList) / Columns
	face := fonts.NewFace(fonts.LucidaBrightDemibold, 20.0)

	Width := float64(Columns)*SampleWidth + float64(Columns-1)*Padding
	Height := float64(Rows)*SampleHeight + float64(Rows-1)*Padding

	gc := gg.NewContext(int(Width), int(Height))
	gc.SetFontFace(face)
	gc.SetStrokeWidth(0.0)
	gc.SetFillColor(colornames.White)
	gc.Clear()

	for i, namedColor := range ColorList {
		row := i % Rows
		col := i / Rows

		color := color.RGBAFModel.Convert(namedColor.color).(color.RGBAF)
		DrawColorSample(gc, col, row, color, namedColor.name)
	}
	gc.SavePNG("colormap.png")
}

func main() {
	ColorList = make([]NamedColor, 0)
	for _, name := range colornames.Names {
		if ok, _ := regexp.MatchString("[Gg]ray", name); ok {
			continue
		}
		ColorList = append(ColorList, NamedColor{name, colornames.Map[name]})
	}
	sort.Slice(ColorList, func(i, j int) bool {
		c1 := color.HSLModel.Convert(ColorList[i].color).(color.HSL)
		c2 := color.HSLModel.Convert(ColorList[j].color).(color.HSL)
		return c1.Less(c2, color.SortByHue) ||
			(!c2.Less(c1, color.SortByHue) && c1.Less(c2, color.SortBySaturation)) ||
			(!c2.Less(c1, color.SortByHue) && !c2.Less(c1, color.SortBySaturation) && c1.Less(c2, color.SortByLightness))
	})

	colorMapRGBAF()
}
