package colors

import (
	"encoding/json"
	"fmt"
	"image/color"
	"math"
	"math/rand"
	"strings"
	"testing"

	"github.com/stefan-muehlebach/gg"
)

// Dies ist eine Hilfsfunktion, mit welcher zwei Fliesskommazahlen auf
// "Gleichheit" überprüft werden können. Mit der Konstanten 'eps' wird
// eine Grenze definiert, wie weit zwei Zahlen auseinanderliegen dürfen,
// um noch als gleich behandelt zu werden.
const (
	eps = 0.005
)

func eq(f1, f2 float64) bool {
	if math.Abs(f1-f2) < eps {
		return true
	} else {
		return false
	}
}

const (
	numConversions = 10_000
	uint8Eps       = 1
	floatEps       = 0.005
)

func abs(i int) int {
	if i < 0 {
		return -i
	} else {
		return i
	}
}

func CompRGBA(c1, c2 RGBA) bool {
	if abs(int(c1.R)-int(c2.R)) > uint8Eps {
		return false
	}
	if abs(int(c1.G)-int(c2.G)) > uint8Eps {
		return false
	}
	if abs(int(c1.B)-int(c2.B)) > uint8Eps {
		return false
	}
	if abs(int(c1.A)-int(c2.A)) > uint8Eps {
		return false
	}
	return true
}

func CompRGBAF(c1, c2 RGBAF) bool {
	if !eq(c1.R, c2.R) {
		return false
	}
	if !eq(c1.G, c2.G) {
		return false
	}
	if !eq(c1.B, c2.B) {
		return false
	}
	if !eq(c1.A, c2.A) {
		return false
	}
	return true
}

func CompHSP(c1, c2 HSP) bool {
	if !eq(c1.H, c2.H) {
		return false
	}
	if !eq(c1.S, c2.S) {
		return false
	}
	if !eq(c1.P, c2.P) {
		return false
	}
	if !eq(c1.A, c2.A) {
		return false
	}
	return true
}

func CompHSV(c1, c2 HSV) bool {
	if !eq(c1.H, c2.H) {
		return false
	}
	if !eq(c1.S, c2.S) {
		return false
	}
	if !eq(c1.V, c2.V) {
		return false
	}
	if !eq(c1.A, c2.A) {
		return false
	}
	return true
}

func CompHSL(c1, c2 HSL) bool {
	if !eq(c1.H, c2.H) {
		return false
	}
	if !eq(c1.S, c2.S) {
		return false
	}
	if !eq(c1.L, c2.L) {
		return false
	}
	if !eq(c1.A, c2.A) {
		return false
	}
	return true
}

func CompHSI(c1, c2 HSI) bool {
	if !eq(c1.H, c2.H) {
		return false
	}
	if !eq(c1.S, c2.S) {
		return false
	}
	if !eq(c1.I, c2.I) {
		return false
	}
	if !eq(c1.A, c2.A) {
		return false
	}
	return true
}

var (
	R, G, B, A     uint8
	r, g, b, a     uint32
	rf, gf, bf, af float64
	h, s, v, p, l  float64

	convColor color.Color

	rgbaColorList = []RGBA{
		{0, 0, 0, 255},
		{255, 255, 255, 255},
		{255, 0, 0, 255},
		{0, 255, 0, 255},
		{0, 0, 255, 255},
		{255, 255, 0, 255},
		{0, 255, 255, 255},
		{255, 0, 255, 255},
		{191, 191, 191, 255},
		{128, 128, 128, 255},
		{128, 0, 0, 255},
		{128, 128, 0, 255},
		{0, 128, 0, 255},
		{128, 0, 128, 255},
		{0, 128, 128, 255},
		{0, 0, 128, 255},
	}

	rgbafColorList = []RGBAF{
		{0, 0, 0, 1.0},
		{1.0, 1.0, 1.0, 1.0},
		{1.0, 0, 0, 1.0},
		{0, 1.0, 0, 1.0},
		{0, 0, 1.0, 1.0},
		{1.0, 1.0, 0, 1.0},
		{0, 1.0, 1.0, 1.0},
		{1.0, 0, 1.0, 1.0},
		{0.75, 0.75, 0.75, 1.0},
		{0.5, 0.5, 0.5, 1.0},
		{0.5, 0, 0, 1.0},
		{0.5, 0.5, 0, 1.0},
		{0, 0.5, 0, 1.0},
		{0.5, 0, 0.5, 1.0},
		{0, 0.5, 0.5, 1.0},
		{0, 0, 0.5, 1.0},
	}

	hspColorList = []HSP{
		{H: 0, S: 0, P: 0, A: 1},
		{H: 0, S: 0, P: 1, A: 1},
		{H: 0, S: 1, P: 0.49092, A: 1},
		{H: 120, S: 1, P: 0.83126, A: 1},
		{H: 240, S: 1, P: 0.26077, A: 1},
		{H: 60, S: 1, P: 0.9654, A: 1},
		{H: 180, S: 1, P: 0.87121, A: 1},
		{H: 300, S: 1, P: 0.55588, A: 1},
		{H: 0, S: 0, P: 0.74902, A: 1},
		{H: 0, S: 0, P: 0.50196, A: 1},
		{H: 0, S: 1, P: 0.24642, A: 1},
		{H: 60, S: 1, P: 0.48459, A: 1},
		{H: 120, S: 1, P: 0.41726, A: 1},
		{H: 300, S: 1, P: 0.27903, A: 1},
		{H: 180, S: 1, P: 0.43731, A: 1},
		{H: 240, S: 1, P: 0.1309, A: 1},
	}

	hsvColorList = []HSV{
		{0, 0, 0, 1},
		{0, 0, 1, 1},
		{0, 1, 1, 1},
		{120, 1, 1, 1},
		{240, 1, 1, 1},
		{60, 1, 1, 1},
		{180, 1, 1, 1},
		{300, 1, 1, 1},
		{0, 0, 0.75, 1},
		{0, 0, 0.5, 1},
		{0, 1, 0.5, 1},
		{60, 1, 0.5, 1},
		{120, 1, 0.5, 1},
		{300, 1, 0.5, 1},
		{180, 1, 0.5, 1},
		{240, 1, 0.5, 1},
	}

	hslColorList = []HSL{
		{0, 0, 0, 1},
		{0, 0, 1, 1},
		{0, 1, 0.5, 1},
		{120, 1, 0.5, 1},
		{240, 1, 0.5, 1},
		{60, 1, 0.5, 1},
		{180, 1, 0.5, 1},
		{300, 1, 0.5, 1},
		{0, 0, 0.75, 1},
		{0, 0, 0.5, 1},
		{0, 1, 0.25, 1},
		{60, 1, 0.25, 1},
		{120, 1, 0.25, 1},
		{300, 1, 0.25, 1},
		{180, 1, 0.25, 1},
		{240, 1, 0.25, 1},
	}

	hsiColorList = []HSI{
		{0, 0, 0, 1},
		{0, 0, 1, 1},
		{0, 1, 0.333, 1},
		{120, 1, 0.333, 1},
		{240, 1, 0.333, 1},
		{60, 1, 0.666, 1},
		{180, 1, 0.666, 1},
		{300, 1, 0.666, 1},
		{0, 0, 0.749, 1},
		{0, 0, 0.502, 1},
		{0, 1, 0.167, 1},
		{60, 1, 0.335, 1},
		{120, 1, 0.167, 1},
		{300, 1, 0.335, 1},
		{180, 1, 0.335, 1},
		{240, 1, 0.167, 1},
	}

	rnd *rand.Rand
)

func init() {
	seed := rand.NewSource(123_456_789_012)
	rnd = rand.New(seed)

	R = uint8(rnd.Intn(256))
	G = uint8(rnd.Intn(256))
	B = uint8(rnd.Intn(256))
	A = uint8(rnd.Intn(256))
}

// func BenchmarkRGBAF2RGBA(bench *testing.B) {
// 	c := RGBAF{rnd.Float64(), rnd.Float64(), rnd.Float64(), rnd.Float64()}
// 	for i := 0; i < bench.N; i++ {
// 		r, g, b, a = c.RGBA()
// 	}
// }
// func BenchmarkRGBA2RGBAF(bench *testing.B) {
// 	c := RGBA{uint8(rnd.Intn(256)), uint8(rnd.Intn(256)), uint8(rnd.Intn(256)), uint8(rnd.Intn(256))}
// 	for i := 0; i < bench.N; i++ {
// 		convColor = RGBAFModel.Convert(c)
// 	}
// }

// func BenchmarkHSP2RGBA(bench *testing.B) {
// 	c := HSP{360.0 * rnd.Float64(), rnd.Float64(), rnd.Float64(), rnd.Float64()}
// 	for i := 0; i < bench.N; i++ {
// 		r, g, b, a = c.RGBA()
// 	}
// }
// func BenchmarkRGBA2HSP(bench *testing.B) {
// 	c := RGBA{uint8(rnd.Intn(256)), uint8(rnd.Intn(256)), uint8(rnd.Intn(256)), uint8(rnd.Intn(256))}
// 	for i := 0; i < bench.N; i++ {
// 		convColor = HSPModel.Convert(c)
// 	}
// }

// func BenchmarkHSV2RGBA(bench *testing.B) {
// 	c := HSV{360.0 * rnd.Float64(), rnd.Float64(), rnd.Float64(), rnd.Float64()}
// 	for i := 0; i < bench.N; i++ {
// 		r, g, b, a = c.RGBA()
// 	}
// }
// func BenchmarkRGBA2HSV(bench *testing.B) {
// 	c := RGBA{uint8(rnd.Intn(256)), uint8(rnd.Intn(256)), uint8(rnd.Intn(256)), uint8(rnd.Intn(256))}
// 	for i := 0; i < bench.N; i++ {
// 		convColor = HSVModel.Convert(c)
// 	}
// }

// func BenchmarkHSL2RGBA(bench *testing.B) {
// 	c := HSL{360.0 * rnd.Float64(), rnd.Float64(), rnd.Float64(), rnd.Float64()}
// 	for i := 0; i < bench.N; i++ {
// 		r, g, b, a = c.RGBA()
// 	}
// }
// func BenchmarkRGBA2HSL(bench *testing.B) {
// 	c := RGBA{uint8(rnd.Intn(256)), uint8(rnd.Intn(256)), uint8(rnd.Intn(256)), uint8(rnd.Intn(256))}
// 	for i := 0; i < bench.N; i++ {
// 		convColor = HSLModel.Convert(c)
// 	}
// }

// func BenchmarkHSI2RGBA(bench *testing.B) {
// 	c := HSI{360.0 * rnd.Float64(), rnd.Float64(), rnd.Float64(), rnd.Float64()}
// 	for i := 0; i < bench.N; i++ {
// 		r, g, b, a = c.RGBA()
// 	}
// }
// func BenchmarkRGBA2HSI(bench *testing.B) {
// 	c := RGBA{uint8(rnd.Intn(256)), uint8(rnd.Intn(256)), uint8(rnd.Intn(256)), uint8(rnd.Intn(256))}
// 	for i := 0; i < bench.N; i++ {
// 		convColor = HSIModel.Convert(c)
// 	}
// }

func ExampleRGBA() {
	r1, g1, b1 := uint8(0xFF), uint8(0x7F), uint8(0x3F)
	c1 := RGBA{r1, g1, b1, 0xFF}
	fmt.Printf("%v", c1)
	// Output:
	// {0xFF, 0x7F, 0x3F, 0xFF}
}

func ExampleRGBAF() {
	r1, g1, b1 := 1.0, 0.5, 0.25
	c1 := RGBAF{r1, g1, b1, 1.0}
	fmt.Printf("%v", c1)
	// Output:
	// {1 0.5 0.25 1}
}

func ExampleHSP() {
	h1, s1, p1 := 0.0, 1.0, 1.0
	c1 := HSP{h1, s1, p1, 1.0}
	fmt.Printf("%v", c1)
	// Output:
	// {0 1 1 1}
}

func ExampleHSV() {
	h1, s1, v1 := 0.0, 1.0, 1.0
	c1 := HSV{h1, s1, v1, 1.0}
	fmt.Printf("%v", c1)
	// Output:
	// {0 1 1 1}
}

func ExampleHSL() {
	h1, s1, l1 := 0.0, 1.0, 0.5
	c1 := HSL{h1, s1, l1, 1.0}
	fmt.Printf("%v", c1)
	// Output:
	// {0 1 0.5 1}
}

func ExampleHSI() {
	h1, s1, i1 := 0.0, 0.5, 0.1
	c1 := HSI{h1, s1, i1, 1.0}
	fmt.Printf("%v", c1)
	// Output:
	// {0 0.5 0.1 1}
}

// Check if every colorname in the slice 'Names' is also found
// in the map 'Map'.
func TestColorNames(t *testing.T) {
	for _, colorName := range Names {
		if _, ok := Map[colorName]; !ok {
			t.Errorf("Colorname '%s': missing in 'Map'.", colorName)
		}
	}
}

// Check if every colorname in the map 'Map' is also found
// in the slice 'Names'.
func TestColorMap(t *testing.T) {
	for colorName, _ := range Map {
		nameFound := false
		for _, name := range Names {
			if colorName == name {
				nameFound = true
				break
			}
		}
		if !nameFound {
			t.Errorf("Colorname '%s': missing in 'Names'.", colorName)
		}
	}
}

// Check if every colorname in the map 'Groups' is also found
// in the map 'Map'.
func TestColorGroups(t *testing.T) {
	for group, colorList := range Groups {
		for _, colorName := range colorList {
			if _, ok := Map[colorName]; !ok {
				t.Errorf("Colorname '%s' in group '%v': missing in 'Map'.",
					colorName, group)
			}
		}
	}
}

// Check if every group name can be converted between string and ColorGroup.
func TestColorGroup(t *testing.T) {
	var g ColorGroup
	var str string
	var err error

	for group := range NumColorGroups {
		str = group.String()
		if err = g.Set(str); err != nil {
			t.Errorf("Color group name '%s': not convertable back", str)
		}
		if g != group {
			t.Errorf("Color group name '%s': wrong value", str)
		}
	}
}

func TestFade(test *testing.T) {
	padding := 10
	fieldSize := 256
	fileName := "color_fade.png"

	width := 4*fieldSize + 5*padding
	height := fieldSize + 2*padding

	gc := gg.NewContext(width, height)

	xOff := padding
	yOff := padding + fieldSize
	hsvColor := HSV{1.0, 0.0, 0.0, 1.0}
	for j := range fieldSize {
		row := yOff - j
		hsvColor.V = float64(j) / 255.0
		for i := range j + 1 {
			col := xOff + i
			hsvColor.S = float64(i) / float64(j+1)
			gc.SetPixel(col, row, hsvColor)
		}
	}

	xOff += fieldSize + padding
	hsiColor := HSI{1.0, 0.0, 0.0, 1.0}
	for j := range fieldSize {
		row := yOff - j
		hsiColor.I = float64(j) / 255.0
		for i := range j + 1 {
			col := xOff + i
			hsiColor.S = float64(i) / float64(j+1)
			gc.SetPixel(col, row, hsiColor)
		}
	}

	xOff += fieldSize + padding
	hslColor := HSL{1.0, 0.0, 0.0, 1.0}
	for j := range fieldSize {
		row := yOff - j
		hslColor.L = float64(j) / 255.0
		max := 2 * j
		if max >= 256 {
			max -= 2 * (max % 256)
		}
		for i := range max + 1 {
			col := xOff + i
			hslColor.S = float64(i) / float64(max+1)
			gc.SetPixel(col, row, hslColor)
		}
	}

	xOff += fieldSize + padding
	hspColor := HSP{1.0, 0.0, 0.0, 1.0}
	for j := range fieldSize {
		row := yOff - j
		hspColor.P = float64(j) / 255.0
		for i := range j + 1 {
			col := xOff + i
			hspColor.S = float64(i) / float64(j+1)
			gc.SetPixel(col, row, hspColor)
		}
	}

	err := gc.SavePNG(fileName)
	if err != nil {
		test.Error(err)
	}
}

func TestRGBAF(test *testing.T) {
	for i := range rgbaColorList {
		rgbaColor := rgbaColorList[i]
		rgbafColor := rgbafColorList[i]

		convRgbaColor := RGBAModel.Convert(rgbafColor).(RGBA)
		convRgbafColor := RGBAFModel.Convert(rgbaColor).(RGBAF)

		if !CompRGBA(rgbaColor, convRgbaColor) {
			test.Errorf("[%d]\n", i)
			test.Errorf("  want: %#v\n", rgbaColor)
			test.Errorf("  got : %#v\n", convRgbaColor)
		}
		if !CompRGBAF(rgbafColor, convRgbafColor) {
			test.Errorf("[%d]\n", i)
			test.Errorf("  want: %#v\n", rgbafColor)
			test.Errorf("  got : %#v\n", convRgbafColor)
		}
	}
}

func TestHSP(test *testing.T) {
	for i := range rgbafColorList {
		rgbaColor := rgbaColorList[i]
		hspColor := hspColorList[i]

		convRgbaColor := RGBAModel.Convert(hspColor).(RGBA)
		convHspColor := HSPModel.Convert(rgbaColor).(HSP)

		if !CompRGBA(rgbaColor, convRgbaColor) {
			test.Errorf("[%d]\n", i)
			test.Errorf("  want: %#v\n", rgbaColor)
			test.Errorf("  got : %#v\n", convRgbaColor)
		}
		if !CompHSP(hspColor, convHspColor) {
			test.Errorf("[%d]\n", i)
			test.Errorf("  want: %#v\n", hspColor)
			test.Errorf("  got : %#v\n", convHspColor)
		}
	}
}

func TestHSV(test *testing.T) {
	for i := range rgbaColorList {
		rgbaColor := rgbaColorList[i]
		hsvColor := hsvColorList[i]

		convRgbaColor := RGBAModel.Convert(hsvColor).(RGBA)
		convHsvColor := HSVModel.Convert(rgbaColor).(HSV)

		if !CompRGBA(rgbaColor, convRgbaColor) {
			test.Errorf("[%d]\n", i)
			test.Errorf("  want: %#v\n", rgbaColor)
			test.Errorf("  got : %#v\n", convRgbaColor)
		}
		if !CompHSV(hsvColor, convHsvColor) {
			test.Errorf("[%d]\n", i)
			test.Errorf("  want: %#v\n", hsvColor)
			test.Errorf("  got : %#v\n", convHsvColor)
		}
	}
}

func TestHSL(test *testing.T) {
	for i := range rgbaColorList {
		rgbaColor := rgbaColorList[i]
		hslColor := hslColorList[i]

		convRgbaColor := RGBAModel.Convert(hslColor).(RGBA)
		convHslColor := HSLModel.Convert(rgbaColor).(HSL)

		if !CompRGBA(rgbaColor, convRgbaColor) {
			test.Errorf("[%d]\n", i)
			test.Errorf("  want: %#v\n", rgbaColor)
			test.Errorf("  got : %#v\n", convRgbaColor)
		}
		if !CompHSL(hslColor, convHslColor) {
			test.Errorf("[%d]\n", i)
			test.Errorf("  want: %#v\n", hslColor)
			test.Errorf("  got : %#v\n", convHslColor)
		}
	}
}

func TestHSI(test *testing.T) {
	for i := range rgbaColorList {
		rgbaColor := rgbaColorList[i]
		hsiColor := hsiColorList[i]

		convRgbaColor := RGBAModel.Convert(hsiColor).(RGBA)
		convHsiColor := HSIModel.Convert(rgbaColor).(HSI)

		if !CompRGBA(rgbaColor, convRgbaColor) {
			test.Errorf("[%d]\n", i)
			test.Errorf("  want: %#v\n", rgbaColor)
			test.Errorf("  got : %#v\n", convRgbaColor)
		}
		if !CompHSI(hsiColor, convHsiColor) {
			test.Errorf("[%d]\n", i)
			test.Errorf("  want: %#v\n", hsiColor)
			test.Errorf("  got : %#v\n", convHsiColor)
		}
	}
}

func TestUnmarshalColor(t *testing.T) {
	type NamedColor struct {
		Name   string  `json:"name"`
		Dark   float64 `json:"nark,omitempty"`
		Bright float64 `json:"bright,omitempty"`
		Alpha  float64 `json:"alpha,omitempty"`
	}
	type RGBAColor struct {
		R byte `json:"r"`
		G byte `json:"g"`
		B byte `json:"b"`
		A byte `json:"a"`
	}
	var j = []byte(`[
        {"r": 0, "g": 0, "b": 0, "a": 128},
        {"r": 128, "g": 128, "b": 128, "a": 128},
        {"name": "Indigo", "dark": 0.5, "alpha": 0.8},
        {"name": "Orange", "dark": 0.5},
        {"value": "0x121212"},
        {"value": "0xf0f0f0"}
    ]`)

	var jsonColors []json.RawMessage
	var colors []RGBA

	dec := json.NewDecoder(strings.NewReader(string(j)))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&jsonColors); err != nil {
		t.Fatalf("error decoding data: %v", err)
	}

	// if err := json.Unmarshal(j, &jsonColors); err != nil {
	// 	t.Fatalf("error: %v", err)
	// }

	for i, c := range jsonColors {
		var namedColor NamedColor = NamedColor{Alpha: 1.0}
		var rgbaColor RGBAColor
		var rgba RGBA

		t.Logf("%d: %s", i, string(c))

		if err := json.Unmarshal(c, &rgba); err != nil {
            t.Logf("  unmarshal(RGBA): %v", err)
            continue
        }
		colors = append(colors, rgba)

		if err := json.Unmarshal(c, &rgbaColor); err == nil {
            t.Logf("  found (RGBAColor)")
			colors = append(colors, RGBA{rgbaColor.R, rgbaColor.G, rgbaColor.B, rgbaColor.A})
			continue
		}

		if err := json.Unmarshal(c, &namedColor); err == nil {
            t.Logf("  found (NamedColor)")
			if col, ok := Map[namedColor.Name]; ok {
				colors = append(colors, col.Dark(namedColor.Dark).Bright(namedColor.Bright).Alpha(namedColor.Alpha))
			} else {
				t.Logf("color not found: %s", namedColor.Name)
			}
			continue
		}


		t.Fatalf("no type found for data '%s'", c)
	}
	t.Log(colors)
}
