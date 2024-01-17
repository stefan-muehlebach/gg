package color

import (
	"fmt"
	"image/color"
    // "log"
	"math"
	"math/rand"
	"testing"
)

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

func CompRGBA(c1, c2 color.RGBA) bool {
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
    if math.Abs(c1.R-c2.R) > floatEps {
        return false
    }
    if math.Abs(c1.G-c2.G) > floatEps {
        return false
    }
    if math.Abs(c1.B-c2.B) > floatEps {
        return false
    }
    if math.Abs(c1.A-c2.A) > floatEps {
        return false
    }
    return true
}

func CompHSV(c1, c2 HSV) bool {
	if math.Abs(c1.H-c2.H) > floatEps {
		return false
	}
	if math.Abs(c1.S-c2.S) > floatEps {
		return false
	}
	if math.Abs(c1.V-c2.V) > floatEps {
		return false
	}
	if math.Abs(c1.A-c2.A) > floatEps {
		return false
	}
	return true
}

func CompHSL(c1, c2 HSL) bool {
	if math.Abs(c1.H-c2.H) > floatEps {
		return false
	}
	if math.Abs(c1.S-c2.S) > floatEps {
		return false
	}
	if math.Abs(c1.L-c2.L) > floatEps {
		return false
	}
	if math.Abs(c1.A-c2.A) > floatEps {
		return false
	}
	return true
}

func CompHSI(c1, c2 HSI) bool {
	if math.Abs(c1.H-c2.H) > floatEps {
		return false
	}
	if math.Abs(c1.S-c2.S) > floatEps {
		return false
	}
	if math.Abs(c1.I-c2.I) > floatEps {
		return false
	}
	if math.Abs(c1.A-c2.A) > floatEps {
		return false
	}
	return true
}

var (
	R, G, B, A uint8
	r, g, b, a uint32
    rf, gf, bf, af float64
	h, s, v, p, l float64

    convColor color.Color

	rgbaColorList = []color.RGBA{
		color.RGBA{0, 0, 0, 255},
		color.RGBA{255, 255, 255, 255},
		color.RGBA{255, 0, 0, 255},
		color.RGBA{0, 255, 0, 255},
		color.RGBA{0, 0, 255, 255},
		color.RGBA{255, 255, 0, 255},
		color.RGBA{0, 255, 255, 255},
		color.RGBA{255, 0, 255, 255},
		color.RGBA{191, 191, 191, 255},
		color.RGBA{128, 128, 128, 255},
		color.RGBA{128, 0, 0, 255},
		color.RGBA{128, 128, 0, 255},
		color.RGBA{0, 128, 0, 255},
		color.RGBA{128, 0, 128, 255},
		color.RGBA{0, 128, 128, 255},
		color.RGBA{0, 0, 128, 255},
		color.RGBA{0, 0, 0, 255},
	}

	rgbafColorList = []RGBAF{
		RGBAF{0, 0, 0, 1.0},
		RGBAF{1.0, 1.0, 1.0, 1.0},
		RGBAF{1.0, 0, 0, 1.0},
		RGBAF{0, 1.0, 0, 1.0},
		RGBAF{0, 0, 1.0, 1.0},
		RGBAF{1.0, 1.0, 0, 1.0},
		RGBAF{0, 1.0, 1.0, 1.0},
		RGBAF{1.0, 0, 1.0, 1.0},
		RGBAF{0.75, 0.75, 0.75, 1.0},
		RGBAF{0.5, 0.5, 0.5, 1.0},
		RGBAF{0.5, 0, 0, 1.0},
		RGBAF{0.5, 0.5, 0, 1.0},
		RGBAF{0, 0.5, 0, 1.0},
		RGBAF{0.5, 0, 0.5, 1.0},
		RGBAF{0, 0.5, 0.5, 1.0},
		RGBAF{0, 0, 0.5, 1.0},
		RGBAF{0, 0, 0, 1.0},
	}

    hspColorList = []HSP{
        HSP{H:0, S:0, P:0, A:1},
        HSP{H:0, S:0, P:1, A:1},
        HSP{H:0, S:1, P:0.49092, A:1},
        HSP{H:120, S:1, P:0.83126, A:1},
        HSP{H:240, S:1, P:0.26077, A:1},
        HSP{H:60, S:1, P:0.9654, A:1},
        HSP{H:180, S:1, P:0.87121, A:1},
        HSP{H:300, S:1, P:0.55588, A:1},
        HSP{H:0, S:0, P:0.74902, A:1},
        HSP{H:0, S:0, P:0.50196, A:1},
        HSP{H:0, S:1, P:0.24642, A:1},
        HSP{H:60, S:1, P:0.48459, A:1},
        HSP{H:120, S:1, P:0.41726, A:1},
        HSP{H:300, S:1, P:0.27903, A:1},
        HSP{H:180, S:1, P:0.43731, A:1},
        HSP{H:240, S:1, P:0.1309, A:1},
        HSP{H:0, S:1, P:1, A:1},
    }

	hsvColorList = []HSV{
		HSV{0, 0, 0, 1},
		HSV{0, 0, 1, 1},
		HSV{0, 1, 1, 1},
		HSV{120, 1, 1, 1},
		HSV{240, 1, 1, 1},
		HSV{60, 1, 1, 1},
		HSV{180, 1, 1, 1},
		HSV{300, 1, 1, 1},
		HSV{0, 0, 0.75, 1},
		HSV{0, 0, 0.5, 1},
		HSV{0, 1, 0.5, 1},
		HSV{60, 1, 0.5, 1},
		HSV{120, 1, 0.5, 1},
		HSV{300, 1, 0.5, 1},
		HSV{180, 1, 0.5, 1},
		HSV{240, 1, 0.5, 1},
	}

	hslColorList = []HSL{
		HSL{0, 0, 0, 1},
		HSL{0, 0, 1, 1},
		HSL{0, 1, 0.5, 1},
		HSL{120, 1, 0.5, 1},
		HSL{240, 1, 0.5, 1},
		HSL{60, 1, 0.5, 1},
		HSL{180, 1, 0.5, 1},
		HSL{300, 1, 0.5, 1},
		HSL{0, 0, 0.75, 1},
		HSL{0, 0, 0.5, 1},
		HSL{0, 1, 0.25, 1},
		HSL{60, 1, 0.25, 1},
		HSL{120, 1, 0.25, 1},
		HSL{300, 1, 0.25, 1},
		HSL{180, 1, 0.25, 1},
		HSL{240, 1, 0.25, 1},
	}

	hsiColorList = []HSI{
		HSI{0, 0, 0, 1},
		HSI{0, 0, 1, 1},
		HSI{0, 1, 0.333, 1},
		HSI{120, 1, 0.333, 1},
		HSI{240, 1, 0.333, 1},
		HSI{60, 1, 0.666, 1},
		HSI{180, 1, 0.666, 1},
		HSI{300, 1, 0.666, 1},
		HSI{0, 0, 0.749, 1},
		HSI{0, 0, 0.502, 1},
		HSI{0, 1, 0.167, 1},
		HSI{60, 1, 0.335, 1},
		HSI{120, 1, 0.167, 1},
		HSI{300, 1, 0.335, 1},
		HSI{180, 1, 0.335, 1},
		HSI{240, 1, 0.167, 1},
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

// func inEpsilon(c1, c2 Color) (bool) {
//     return math.Abs(float64(c1.R-c2.R)) <= colorEps &&
//             math.Abs(float64(c1.G-c2.G)) <= colorEps &&
//             math.Abs(float64(c1.B-c2.B)) <= colorEps
// }

func BenchmarkRGBAF2RGBA(bench *testing.B) {
	c := RGBAF{rnd.Float64(), rnd.Float64(), rnd.Float64(), rnd.Float64()}
	for i := 0; i < bench.N; i++ {
		r, g, b, a = c.RGBA()
	}
}
func BenchmarkRGBA2RGBAF(bench *testing.B) {
	c := color.RGBA{uint8(rnd.Intn(256)), uint8(rnd.Intn(256)), uint8(rnd.Intn(256)), uint8(rnd.Intn(256))}
	for i := 0; i < bench.N; i++ {
		convColor = RGBAFModel.Convert(c)
	}
}

func BenchmarkHSL2RGBA(bench *testing.B) {
	c := HSL{360.0 * rnd.Float64(), rnd.Float64(), rnd.Float64(), rnd.Float64()}
	for i := 0; i < bench.N; i++ {
		r, g, b, a = c.RGBA()
	}
}
func BenchmarkRGBA2HSL(bench *testing.B) {
	c := color.RGBA{uint8(rnd.Intn(256)), uint8(rnd.Intn(256)), uint8(rnd.Intn(256)), uint8(rnd.Intn(256))}
	for i := 0; i < bench.N; i++ {
		convColor = HSLModel.Convert(c)
	}
}

func BenchmarkHSV2RGBA(bench *testing.B) {
	c := HSV{360.0 * rnd.Float64(), rnd.Float64(), rnd.Float64(), rnd.Float64()}
	for i := 0; i < bench.N; i++ {
		r, g, b, a = c.RGBA()
	}
}
func BenchmarkRGBA2HSV(bench *testing.B) {
	c := color.RGBA{uint8(rnd.Intn(256)), uint8(rnd.Intn(256)), uint8(rnd.Intn(256)), uint8(rnd.Intn(256))}
	for i := 0; i < bench.N; i++ {
		convColor = HSVModel.Convert(c)
	}
}

func BenchmarkHSI2RGBA(bench *testing.B) {
	c := HSI{360.0 * rnd.Float64(), rnd.Float64(), rnd.Float64(), rnd.Float64()}
	for i := 0; i < bench.N; i++ {
		r, g, b, a = c.RGBA()
	}
}
func BenchmarkRGBA2HSI(bench *testing.B) {
	c := color.RGBA{uint8(rnd.Intn(256)), uint8(rnd.Intn(256)), uint8(rnd.Intn(256)), uint8(rnd.Intn(256))}
	for i := 0; i < bench.N; i++ {
		convColor = HSIModel.Convert(c)
	}
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

func TestRGBAF(test *testing.T) {
	for i := range rgbaColorList {
		rgbaColor   := rgbaColorList[i]
		rgbafColor := rgbafColorList[i]

		convRgbaColor := color.RGBAModel.Convert(rgbafColor).(color.RGBA)
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
		rgbafColor := rgbafColorList[i]
		hspColor := hspColorList[i]

		convRgbafColor := RGBAFModel.Convert(hspColor).(RGBAF)
		convHspColor := HSPModel.Convert(rgbafColor).(HSP)

		if !CompRGBAF(rgbafColor, convRgbafColor) {
			test.Errorf("[%d]\n", i)
			test.Errorf("  want: %#v\n", rgbafColor)
			test.Errorf("  got : %#v\n", convRgbafColor)
		}
        test.Logf("%#v  vs  %#v\n", hspColor, convHspColor)
	}
}

func TestHSV(test *testing.T) {
	for i := range rgbaColorList {
		rgbaColor := rgbaColorList[i]
		hsvColor := hsvColorList[i]

		convRgbaColor := color.RGBAModel.Convert(hsvColor).(color.RGBA)
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

		convRgbaColor := color.RGBAModel.Convert(hslColor).(color.RGBA)
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

		convRgbaColor := color.RGBAModel.Convert(hsiColor).(color.RGBA)
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
