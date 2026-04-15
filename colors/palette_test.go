package colors

import (
	"math/rand"
	"testing"

	"github.com/stefan-muehlebach/gg/fonts"
)

var (
	ColorBarWidth     = 1024
	ColorBarHeight    = 100
	ColorBarLineColor = DarkSlateGray
	ColorBarLineWidth = 3.0
	FontSize          = 18
	TextHeight        = FontSize + BoxPadding
	BoxPadding        = 10
	BoxLineWidth      = 1.0
	BoxLineColor      = Black
	BoxWidth          = ColorBarWidth + 2*BoxPadding
	BoxHeight         = ColorBarHeight + TextHeight + 2*BoxPadding
	SimplePalSize     = 1024
	RegularFont       = fonts.LucidaBrightDemibold
	ItalicFont        = fonts.LucidaBrightItalic
	c                 RGBA
)

func TestReadPaletteFile(t *testing.T) {
	_, _, err := ReadPaletteFile("paletten.json")
	if err != nil {
		t.Fatalf("error: %v", err)
	}
}

var (
	LinearData = []RGBA{
		{0x00, 0xff, 0x00, 0xff},
		{0xff, 0x00, 0x7f, 0xff},
		{0x00, 0xff, 0xff, 0xff},
	}
	LinearPosList = []float64{
		0.0, 0.5, 1.0,
	}
)

func TestLinearPalette(t *testing.T) {
	p := NewLinearPalette("Linear")
	for _, c := range LinearData {
		p.AddColor(c)
	}
	for i, v := range LinearPosList {
		c1 := LinearData[i]
		c2 := p.Color(v)
		if c1 != c2 {
			t.Logf("[%f]: expected %#v but got %#v", v, c1, c2)
		}
	}
}

var (
	ColorStopsData = []ColorStop{
		{0.0, RGBA{0x00, 0xff, 0x00, 0xff}, 0},
		{0.1, RGBA{0xff, 0x00, 0x7f, 0xff}, 0},
		{1.0, RGBA{0x00, 0xff, 0xff, 0xff}, 0},
	}
	ColorStopsPosList = []float64{
		0.0, 0.1, 1.0,
	}
)

func TestColorStopsPalette(t *testing.T) {
	p := NewColorStopsPalette("ColorStops")
	for _, s := range ColorStopsData {
		p.SetColorStop(s)
	}
	for i, v := range ColorStopsPosList {
		c1 := ColorStopsData[i].Color
		c2 := p.Color(v)
		if c1 != c2 {
			t.Logf("[%f]: expected %#v but got %#v", v, c1, c2)
		}
	}
}

var (
	FunctionData = []ParamSet{
		{0.5, 0.5, 1.0, 0.0},
		{0.5, 0.5, 1.0, 0.0},
		{0.5, 0.5, 1.0, 0.0},
	}
	FunctionPosList = []float64{
		0.0, 0.5, 1.0,
	}
	FunctionColorList = []RGBA{
		{0xFF, 0xFF, 0xFF, 0xFF},
		{0x00, 0x00, 0x00, 0xFF},
		{0xFF, 0xFF, 0xFF, 0xFF},
	}
)

func TestFunctionPalette(t *testing.T) {
	p := NewFunctionPalette("Function")
	for i, param := range FunctionData {
		p.SetParam(ColorIdent(i), param)
	}
	for i, v := range FunctionPosList {
		c1 := FunctionColorList[i]
		c2 := p.Color(v)
		if c1 != c2 {
			t.Logf("[%f]: expected %#v but got %#v", v, c1, c2)
		}
	}
}

func BenchmarkSimplePalette(b *testing.B) {
	palFileName := "paletten.json"
	palName := "TestLinear"

	seed := rand.NewSource(123_456_789_012)
	rnd := rand.New(seed)

	_, palMap, err := ReadPaletteFile(palFileName)
	if err != nil {
		b.Fatalf("error: %v", err)
	}
	pal := NewPaletteFromOther("Plasma (Simple)", SimplePalSize, palMap[palName])
	for b.Loop() {
		c = pal.Color(rnd.Float64())
	}
}

func BenchmarkLinearPalette(b *testing.B) {
	palFileName := "paletten.json"
	palName := "TestLinear"

	seed := rand.NewSource(123_456_789_012)
	rnd := rand.New(seed)

	_, palMap, err := ReadPaletteFile(palFileName)
	if err != nil {
		b.Fatalf("error: %v", err)
	}
	pal := palMap[palName]
	for b.Loop() {
		c = pal.Color(rnd.Float64())
	}
}

func BenchmarkColorStopsPalette(b *testing.B) {
	palFileName := "paletten.json"
	palName := "TestColorStops"

	seed := rand.NewSource(123_456_789_012)
	rnd := rand.New(seed)

	_, palMap, err := ReadPaletteFile(palFileName)
	if err != nil {
		b.Fatalf("error: %v", err)
	}
	pal := palMap[palName]
	for b.Loop() {
		c = pal.Color(rnd.Float64())
	}
}

func BenchmarkFunctionPalette(b *testing.B) {
	palFileName := "paletten.json"
	palName := "TestFunction"

	seed := rand.NewSource(123_456_789_012)
	rnd := rand.New(seed)

	_, palMap, err := ReadPaletteFile(palFileName)
	if err != nil {
		b.Fatalf("error: %v", err)
	}
	pal := palMap[palName]
	for b.Loop() {
		c = pal.Color(rnd.Float64())
	}
}
