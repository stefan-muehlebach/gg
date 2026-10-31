package colors

import (
	"log"
	"math/rand"
	"os"
	"testing"

	"github.com/stefan-muehlebach/gg"
	"github.com/stefan-muehlebach/gg/fonts"
)

var (
	// Von den beiden Variablen NumCols und NumRows kann eine Null sein, dann
	// wird sie aufgrund der anderen Variable (die nicht Null sein darf)
	// berechnet.
	NumCols           = 4
	NumRows           = 0
	ColorBoxWidth     = 600
	ColorBoxHeight    = 100
	TextBoxHeight     = 30
	ColorPaddingHori  = 20
	ColorPaddingVert  = 20
	NameFont          = fonts.LucidaBrightDemibold
	TypeFont          = fonts.LucidaBrightItalic
	FontSize          = float64(2.0 * TextBoxHeight / 3.0)
	FontSizeSmall     = float64(TextBoxHeight / 2.0)
	JSONFileName      = "palettes.json"
	PNGFileName       = "palettes.png"
	RandName          = "Randomizer"
	RandSize          = 50
	// ColorBarWidth        = 1024
	// ColorBarHeight    = 100
	// ColorBarLineColor = DarkSlateGray
	// ColorBarLineWidth = 3.0
	// FontSize          = 18
	// TextHeight        = FontSize + BoxPadding
	// BoxPadding        = 10
	// BoxLineWidth      = 1.0
	// BoxLineColor      = Black
	// BoxWidth          = ColorBarWidth + 2*BoxPadding
	// BoxHeight         = ColorBarHeight + TextHeight + 2*BoxPadding
	SimplePalSize     = 1024
	// RegularFont       = fonts.LucidaBrightDemibold
	// ItalicFont        = fonts.LucidaBrightItalic
	c RGBA
)

func TestReadPaletteFile(t *testing.T) {
	_, _, err := ReadPaletteFile(JSONFileName)
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

func TestOverview(t *testing.T) {
	var palNames []string
	var palMap map[string]Palette
	var err error

	fh, err := os.Open(JSONFileName)
	if err != nil {
		log.Fatalf("%+v", err)
	}
	defer fh.Close()

	palNames, palMap, err = ReadPaletteData(fh)
	if err != nil {
		log.Fatalf("%+v", err)
	}

	// Baue eine randomisierte Palette zusammen
	randPal := NewLinearPalette(RandName)
	var col, oldCol RGBA
	for i := 0; i < RandSize; i++ {
		col = RandColor()
		for i > 0 && col.Dist(oldCol) < 128 {
			col = RandColor()
		}
		randPal.AddColor(col.Dark(0.3))
		oldCol = col
	}
	palNames = append(palNames, RandName)
	palMap[RandName] = randPal

	if NumCols == 0 {
		if NumRows == 0 {
			log.Fatal("NumCols and NumRows can't be both zero!")
		}
		NumCols = len(palNames) / NumRows
		if len(palNames)%NumRows != 0 {
			NumCols += 1
		}
	} else {
		if NumRows == 0 {
			NumRows = len(palNames) / NumCols
			if len(palNames)%NumCols != 0 {
				NumRows += 1
			}
		}
	}
	nameFace, _ := fonts.NewFace(NameFont, FontSize)
	typeFace, _ := fonts.NewFace(TypeFont, FontSizeSmall)
	gc := gg.NewContext(NumCols*(ColorBoxWidth)+(NumCols-1)*ColorPaddingHori,
		NumRows*ColorBoxHeight+NumRows*(ColorPaddingVert+TextBoxHeight))
	gc.SetTextColor(Black)
	gc.SetFillColor(WhiteSmoke)
	gc.Clear()
	for i, name := range palNames {
		col, row := i/NumRows, i%NumRows
		x0 := col * (ColorBoxWidth + ColorPaddingHori)
		y0 := row * (ColorBoxHeight + ColorPaddingVert + TextBoxHeight)
		pal := palMap[name]
		for w := range ColorBoxWidth + 1 {
			t := float64(w) / float64(ColorBoxWidth)
			color := pal.Color(t)
			x := x0 + w
			for h := range ColorBoxHeight {
				y := y0 + h
				gc.SetPixel(x, y, color)
			}
		}

		gc.SetFontFace(nameFace)
		gc.DrawStringAnchored(name, float64(x0), float64(y0+ColorBoxHeight+TextBoxHeight), 0.0, 0.0)
		gc.SetFontFace(typeFace)
		gc.DrawStringAnchored(pal.Type().String(), float64(x0+ColorBoxWidth), float64(y0+ColorBoxHeight+TextBoxHeight), 1.0, 0.0)
	}
	gc.SavePNG(PNGFileName)
}

func BenchmarkSimplePalette(b *testing.B) {
	palName := "TestLinear"

	seed := rand.NewSource(123_456_789_012)
	rnd := rand.New(seed)

	_, palMap, err := ReadPaletteFile(JSONFileName)
	if err != nil {
		b.Fatalf("error: %v", err)
	}
	pal := NewPaletteFromOther("Plasma (Simple)", SimplePalSize, palMap[palName])
	for b.Loop() {
		c = pal.Color(rnd.Float64())
	}
}

func BenchmarkLinearPalette(b *testing.B) {
	palName := "TestLinear"

	seed := rand.NewSource(123_456_789_012)
	rnd := rand.New(seed)

	_, palMap, err := ReadPaletteFile(JSONFileName)
	if err != nil {
		b.Fatalf("error: %v", err)
	}
	pal := palMap[palName]
	for b.Loop() {
		c = pal.Color(rnd.Float64())
	}
}

func BenchmarkColorStopsPalette(b *testing.B) {
	palName := "TestColorStops"

	seed := rand.NewSource(123_456_789_012)
	rnd := rand.New(seed)

	_, palMap, err := ReadPaletteFile(JSONFileName)
	if err != nil {
		b.Fatalf("error: %v", err)
	}
	pal := palMap[palName]
	for b.Loop() {
		c = pal.Color(rnd.Float64())
	}
}

func BenchmarkFunctionPalette(b *testing.B) {
	palName := "TestFunction"

	seed := rand.NewSource(123_456_789_012)
	rnd := rand.New(seed)

	_, palMap, err := ReadPaletteFile(JSONFileName)
	if err != nil {
		b.Fatalf("error: %v", err)
	}
	pal := palMap[palName]
	for b.Loop() {
		c = pal.Color(rnd.Float64())
	}
}
