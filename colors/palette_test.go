package colors

import (
	"math/rand"
	"testing"

	"github.com/stefan-muehlebach/gg"
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
	_, _, err := ReadPaletteFile("palette.json")
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

func TestDrawPalette(t *testing.T) {
	imgFileName := "palette.png"
	palFileName := "palette.json"

	palNames, palMap, err := ReadPaletteFile(palFileName)
	if err != nil {
		t.Fatalf("error: %v", err)
	}

	face1, _ := fonts.NewFace(RegularFont, float64(FontSize))
	face2, _ := fonts.NewFace(ItalicFont, float64(FontSize))

	gc := gg.NewContext(BoxWidth, len(palNames)*BoxHeight)
	gc.SetFillColor(White)
	gc.Clear()
	gc.SetTextColor(Black)
	for i, name := range palNames {
		x0, y0 := 0, i*BoxHeight
		gc.SetStrokeColor(BoxLineColor)
		gc.SetStrokeWidth(BoxLineWidth)
		gc.DrawRectangle(float64(x0), float64(y0), float64(BoxWidth), float64(BoxHeight))
		gc.Stroke()
		x0 += BoxPadding
		y0 += BoxPadding
		pal := palMap[name]
		sPal := NewPaletteFromOther(name+" (Simple)", SimplePalSize, pal)
		for col := range ColorBarWidth {
			x := x0 + col
			c1 := pal.Color(float64(col) / float64(ColorBarWidth-1))
			c2 := sPal.Color(float64(col) / float64(ColorBarWidth-1))
			for row := range ColorBarHeight {
				y := y0 + row
				if row < ColorBarHeight/2 {
					gc.SetPixel(x, y, c1)
				} else if row > ColorBarHeight/2 {
					gc.SetPixel(x, y, c2)
				}
			}
		}
		gc.SetStrokeColor(ColorBarLineColor)
		gc.SetStrokeWidth(ColorBarLineWidth)
		gc.DrawRectangle(float64(x0), float64(y0), float64(ColorBarWidth), float64(ColorBarHeight))
		gc.Stroke()
		y0 += ColorBarHeight + BoxPadding/2
		gc.SetFontFace(face1)
		gc.DrawStringAnchored(pal.Name(), float64(x0), float64(y0), 0.0, 1.0)
		gc.SetFontFace(face2)
		gc.DrawStringAnchored(pal.Type().String(), float64(x0+ColorBarWidth), float64(y0), 1.0, 1.0)
	}
	gc.SavePNG(imgFileName)
}

func BenchmarkSimplePalette(b *testing.B) {
	palFileName := "palette.json"
	palName := "Plasma"

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
	palFileName := "palette.json"
	palName := "Plasma"

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
	palFileName := "palette.json"
	palName := "EarthAndSky"

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
	palFileName := "palette.json"
	palName := "Neon"

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
