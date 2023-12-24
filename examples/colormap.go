package main

import (
	"math"
	"regexp"

	"github.com/stefan-muehlebach/gg"
	"github.com/stefan-muehlebach/gg/color"
	"github.com/stefan-muehlebach/gg/colornames"
	"github.com/stefan-muehlebach/gg/fonts"
)

type NamedColor struct {
	name  string
	color color.Color
}

type NamedGroup struct {
	name string
	list []NamedColor
}

var (
	SampleWidth   = 250.0
	SampleHeight  = 2.0*FadeHeight + UniformHeight
	UniformHeight = 50.0
	FadeHeight    = 20.0
	NumFadeSteps  = 9
	FadeStep      = 1.0 / float64(NumFadeSteps+1)
	FadeWidth     = SampleWidth / float64(NumFadeSteps)

	Padding       = 10.0

	TextFontSize  = 20.0
	TextFont      = fonts.LucidaBrightDemibold
	TextFontFace  = fonts.NewFace(TextFont, TextFontSize)
	TitleFontSize = 40.0
	TitleFont     = fonts.LucidaBrightDemibold
	TitleFontFace = fonts.NewFace(TitleFont, TitleFontSize)
)

func DrawColorSample(gc *gg.Context, x0, y0 float64, namedCol NamedColor) {
	gc.SetFillColor(namedCol.color)
	gc.DrawRectangle(x0, y0+FadeHeight, SampleWidth, UniformHeight)
	gc.Fill()
	col := color.RGBAFModel.Convert(namedCol.color).(color.RGBAF)

	for l := 0; l < NumFadeSteps; l++ {
		t := FadeStep * float64(l+1)
		colBright := col.Bright(1.0 - t)
		colDark := col.Dark(t)
		x := x0 + float64(l)*FadeWidth
		y := y0
		gc.SetFillColor(colBright)
		gc.DrawRectangle(x, y, FadeWidth, FadeHeight)
		gc.Fill()
		y += FadeHeight + UniformHeight
		gc.SetFillColor(colDark)
		gc.DrawRectangle(x, y, FadeWidth, FadeHeight)
		gc.Fill()
	}
	r, g, b, _ := col.RGBA()
	max := max3(r, g, b)
	min := min3(r, g, b)
	mid := (max + min) / 2
	if mid >= math.MaxInt16 {
		gc.SetStrokeColor(colornames.Black)
	} else {
		gc.SetStrokeColor(colornames.WhiteSmoke)
	}
	gc.SetFontFace(TextFontFace)
	gc.DrawStringAnchored(namedCol.name, x0+SampleWidth/2.0, y0+SampleHeight/2.0, 0.5, 0.5)
}

func DrawColorMap(groupList []NamedGroup) {
	Columns := 6
	numSlots := len(groupList)
	for _, namedGroup := range groupList {
		numSlots += len(namedGroup.list)
	}
	Rows := numSlots/Columns + 1

	Width := float64(Columns)*(SampleWidth+Padding) - Padding
	Height := float64(Rows)*(SampleHeight+Padding) - Padding

	gc := gg.NewContext(int(Width), int(Height))
	gc.SetStrokeWidth(0.0)
	gc.SetFillColor(colornames.White)
	gc.Clear()

	slotIndex := 0
	for _, namedGroup := range groupList {
		for j, namedColor := range namedGroup.list {
			if j == 0 {
                for Rows - (slotIndex % Rows) < 3 {
                    slotIndex += 1
                }
				x0 := float64(slotIndex / Rows) * (SampleWidth + Padding)
				y0 := float64(slotIndex % Rows) * (SampleHeight + Padding)   
				gc.SetFontFace(TitleFontFace)
				gc.SetStrokeColor(colornames.Black)
				gc.DrawStringAnchored(namedGroup.name, x0+SampleWidth/2, y0+SampleHeight/2, 0.5, 0.5)
				slotIndex += 1
			}
        		x0 := float64(slotIndex / Rows) * (SampleWidth + Padding)
			y0 := float64(slotIndex % Rows) * (SampleHeight + Padding)   
			DrawColorSample(gc, x0, y0, namedColor)
			slotIndex += 1
		}
	}
	gc.SavePNG("colormap.png")
}

func max(a, b uint32) uint32 {
	if a > b {
		return a
	} else {
		return b
	}
}
func max3(a, b, c uint32) uint32 {
	return max(a, max(b, c))
}

func min(a, b uint32) uint32 {
	if a < b {
		return a
	} else {
		return b
	}
}
func min3(a, b, c uint32) uint32 {
	return min(a, min(b, c))
}

func main() {
	var groupIndex colornames.ColorGroup
    var groupList []NamedGroup

	groupList = make([]NamedGroup, colornames.NumColorGroups)
	for i := range groupList {
		groupList[i].list = make([]NamedColor, 0)
	}

	for groupIndex = 0; groupIndex < colornames.NumColorGroups; groupIndex++ {
		groupList[groupIndex].name = groupIndex.String()
		for _, colorName := range colornames.Groups[groupIndex] {
			if ok, _ := regexp.MatchString("[Gg]rey", colorName); ok {
				continue
			}
			groupList[groupIndex].list = append(groupList[groupIndex].list, NamedColor{colorName, colornames.Map[colorName]})
		}
	}
	DrawColorMap(groupList)
}
