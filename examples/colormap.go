package main

import (
	"fmt"
	"image/color"
	"regexp"

	"github.com/stefan-muehlebach/gg"
	"github.com/stefan-muehlebach/gg/colors"
	"github.com/stefan-muehlebach/gg/fonts"
)

type NamedColor struct {
	name  string
	color colors.Color
}

type NamedGroup struct {
	name string
	list []NamedColor
}

var (
	// Die Breite des Feldes fuer eine einzelne Farbe.
	SampleWidth = 250.0
	// Die Hoehe des Teils, in dem die Farbe uni dargestellt wird.
	UniformHeight = 60.0
	// Die Hoehe eines Streifens mit einem Farbverlauf.
	FadeHeight = 30.0
	// Die Anzahl Verlaufsschritte.
	NumFadeSteps = 9
	// Der Rand, welcher um jedes Farbfeld gelegt wird.
	Padding = 5.0
	// Font und Groesse der Farbgruppentitel
	TitleFont     = fonts.GoBold
	TitleFontSize = 40.0
	// Font und Groesse der Farbbezeichnungen.
	TextFont     = fonts.GoMedium
	TextFontSize = 20.0

	// Abgeleitete Groessen (sollten nicht angepasst werden muessen)
	SampleHeight  = 2.0*FadeHeight + UniformHeight
	FadeStep      = 1.0 / float64(NumFadeSteps+1)
	FadeWidth     = SampleWidth / float64(NumFadeSteps)
	TitleFontFace = fonts.NewFace(TitleFont, TitleFontSize)
	TextFontFace  = fonts.NewFace(TextFont, TextFontSize)
)

var (
	GoColorGroup = NamedGroup{
		name: "GoColors",
		list: []NamedColor{
			{"GoGopherBlue", colors.RGBAF{0.004, 0.678, 0.847, 1}},
			{"GoLightBlue", colors.RGBAF{0.369, 0.788, 0.890, 1}},
			{"GoAqua", colors.RGBAF{0.000, 0.635, 0.622, 1}},
			{"GoBlack", colors.RGBAF{0.000, 0.000, 0.000, 1}},
			{"GoFuchsia", colors.RGBAF{0.808, 0.188, 0.384, 1}},
			{"GoYellow", colors.RGBAF{0.992, 0.867, 0.000, 1}},
			{"GoTeal", colors.RGBAF{0.000, 0.520, 0.553, 1}},
			{"GoDimGray", colors.RGBAF{0.333, 0.341, 0.349, 1}},
			{"GoIndigo", colors.RGBAF{0.251, 0.169, 0.337, 1}},
			{"GoLightGray", colors.RGBAF{0.859, 0.851, 0.839, 1}},
		},
	}
)

func DrawColorSample(gc *gg.Context, x0, y0 float64, namedCol NamedColor) {
	gc.SetFillColor(namedCol.color)
	gc.DrawRectangle(x0, y0+FadeHeight, SampleWidth, UniformHeight)
	gc.Fill()
	col := namedCol.color
    hspCol := colors.HSPModel.Convert(namedCol.color).(colors.HSP)

	for l := 0; l < NumFadeSteps; l++ {
		t := FadeStep * float64(l+1)
		brightCol := col.Bright(1.0 - t)
		darkCol := col.Dark(t)
		x := x0 + float64(l)*FadeWidth
        DrawFadeField(gc, x, y0, brightCol)
        DrawFadeField(gc, x, y0+FadeHeight+UniformHeight, darkCol)
	}
	if hspCol.P < 0.6 {
		gc.SetStrokeColor(colors.WhiteSmoke)
	} else {
		gc.SetStrokeColor(colors.Black)
	}
	gc.SetFontFace(TextFontFace)
	gc.DrawStringAnchored(namedCol.name, x0+SampleWidth/2.0, y0+SampleHeight/2.0, 0.5, 0.5)
}

func DrawFadeField(gc *gg.Context, x, y float64, col colors.Color) {
    gc.SetFillColor(col)
    gc.DrawRectangle(x, y, FadeWidth, FadeHeight)
    gc.Fill()
}

func DrawColorMap(groupList []NamedGroup, fileName string) {
    gc := CreateCanvas(groupList)

	for column, namedGroup := range groupList {
		x0 := float64(column) * (SampleWidth + Padding)
		for row, namedColor := range namedGroup.list {
			if row == 0 {
				y0 := float64(row) * (SampleHeight + Padding)
				gc.SetStrokeColor(colors.Black)
				gc.SetStrokeWidth(2.0)
				gc.SetFillColor(colors.Silver)
				gc.DrawRectangle(x0, y0, SampleWidth, SampleHeight)
				gc.Fill()
				gc.SetFontFace(TitleFontFace)
				gc.SetStrokeColor(colors.Black)
				gc.DrawStringWrapped(namedGroup.name, x0+SampleWidth/2, y0+SampleHeight/2, 0.5, 0.5, SampleWidth, 1.0, gg.AlignCenter)
			}
			y0 := float64(row+1) * (SampleHeight + Padding)
			DrawColorSample(gc, x0, y0, namedColor)
		}
	}
	gc.SavePNG(fileName)
}

func CreateCanvas(groupList []NamedGroup) *gg.Context {
	columns := len(groupList)
	rows := 0
	for _, namedGroup := range groupList {
		rows = max(rows, len(namedGroup.list))
	}
	rows += 1

	width := float64(columns)*(SampleWidth+Padding) - Padding
	height := float64(rows)*(SampleHeight+Padding) - Padding

	gc := gg.NewContext(int(width), int(height))
	gc.SetStrokeWidth(0.0)
	gc.SetFillColor(colors.White)
	gc.Clear()

    return gc
}

func PrepareColorList() []NamedGroup {
	var groupIndex colors.ColorGroup
	var groupList []NamedGroup

	groupList = make([]NamedGroup, colors.NumColorGroups)
	for i := range groupList {
		groupList[i].list = make([]NamedColor, 0)
	}

	for groupIndex = 0; groupIndex < colors.NumColorGroups; groupIndex++ {
		groupList[groupIndex].name = groupIndex.String()
		for _, colorName := range colors.Groups[groupIndex] {
			if ok, _ := regexp.MatchString("[Gg]rey", colorName); ok {
				continue
			}
			groupList[groupIndex].list = append(groupList[groupIndex].list, NamedColor{colorName, colors.Map[colorName]})
		}
	}
	return groupList
}

func PrepareFadeList(groupIndex colors.ColorGroup, modelList []color.Model) []NamedGroup {
    var groupList []NamedGroup

    groupList = make([]NamedGroup, len(modelList))
    for i, model := range modelList {
        groupList[i] = NamedGroup{
            fmt.Sprintf("%s\n%T", groupIndex, model),
            make([]NamedColor, len(colors.Groups[groupIndex])),
        }
        for j, colorName := range colors.Groups[groupIndex] {
            groupList[i].list[j] = NamedColor{
                colorName,
                model.Convert(colors.Map[colorName]).(colors.Color),
            }
        }
    }

    return groupList
}

func main() {
    // modelList := []gocolors.Model{
    //     colors.RGBAFModel,
    //     colors.HSPModel,
    // }
    // groupList := PrepareFadeList(colors.Greens, modelList)
    // colors.SetInterpolFunc(colors.LinearInterpol)
	// DrawColorMap(groupList, "colormap-linear.png")
    // colors.SetInterpolFunc(colors.CubicInterpol)
	// DrawColorMap(groupList, "colormap-cubic.png")
    // colors.SetInterpolFunc(colors.GammaInterpol)
	// DrawColorMap(groupList, "colormap-gamma.png")

    	groupList := PrepareColorList()
	groupList = append(groupList, GoColorGroup)
    DrawColorMap(groupList, "colormap.png")

}
