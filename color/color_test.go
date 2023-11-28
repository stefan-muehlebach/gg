package color

import (
    "fmt"
    "image/color"
    "math"
    "math/rand"
    "testing"
)

const (
    numConversions = 10_000
    uint8Eps       = 1
    floatEps       = 0.005
)

func abs(i int) (int) {
    if i < 0.0 {
        return -i
    } else {
        return  i
    }
}

func CompRGBA(c1, c2 color.RGBA) (bool) {
    if abs(int(c1.R) - int(c2.R)) > uint8Eps {
        return false
    }
    if abs(int(c1.G) - int(c2.G)) > uint8Eps {
        return false
    }
    if abs(int(c1.B) - int(c2.B)) > uint8Eps {
        return false
    }
    if abs(int(c1.A) - int(c2.A)) > uint8Eps {
        return false
    }
    return true
}

func CompHSV(c1, c2 HSV) (bool) {
    if math.Abs(c1.H - c2.H) > floatEps {
        return false
    }
    if math.Abs(c1.S - c2.S) > floatEps {
        return false
    }
    if math.Abs(c1.V - c2.V) > floatEps {
        return false
    }
    if math.Abs(c1.A - c2.A) > floatEps {
        return false
    }
    return true
}

func CompHSL(c1, c2 HSL) (bool) {
    if math.Abs(c1.H - c2.H) > floatEps {
        return false
    }
    if math.Abs(c1.S - c2.S) > floatEps {
        return false
    }
    if math.Abs(c1.L - c2.L) > floatEps {
        return false
    }
    if math.Abs(c1.A - c2.A) > floatEps {
        return false
    }
    return true
}

func CompHSI(c1, c2 HSI) (bool) {
    if math.Abs(c1.H - c2.H) > floatEps {
        return false
    }
    if math.Abs(c1.S - c2.S) > floatEps {
        return false
    }
    if math.Abs(c1.I - c2.I) > floatEps {
        return false
    }
    if math.Abs(c1.A - c2.A) > floatEps {
        return false
    }
    return true
}



var (
    R, G, B, A uint8
    r, g, b, a uint32
    h, s, v, l float64

    rgbColorList = []color.RGBA{
        color.RGBA{  0,   0,   0, 255},
        color.RGBA{255, 255, 255, 255},
        color.RGBA{255,   0,   0, 255},
        color.RGBA{  0, 255,   0, 255},
        color.RGBA{  0,   0, 255, 255},
        color.RGBA{255, 255,   0, 255},
        color.RGBA{  0, 255, 255, 255},
        color.RGBA{255,   0, 255, 255},
        color.RGBA{191, 191, 191, 255},
        color.RGBA{128, 128, 128, 255},
        color.RGBA{128,   0,   0, 255},
        color.RGBA{128, 128,   0, 255},
        color.RGBA{  0, 128,   0, 255},
        color.RGBA{128,   0, 128, 255},
        color.RGBA{  0, 128, 128, 255},
        color.RGBA{  0,   0, 128, 255},
    }

    hsvColorList = []HSV{
        HSV{  0, 0, 0, 1},
        HSV{  0, 0, 1, 1},
        HSV{  0, 1, 1, 1},
        HSV{120, 1, 1, 1},
        HSV{240, 1, 1, 1},
        HSV{ 60, 1, 1, 1},
        HSV{180, 1, 1, 1},
        HSV{300, 1, 1, 1},
        HSV{  0, 0, 0.75, 1},
        HSV{  0, 0, 0.5, 1},
        HSV{  0, 1, 0.5, 1},
        HSV{ 60, 1, 0.5, 1},
        HSV{120, 1, 0.5, 1},
        HSV{300, 1, 0.5, 1},
        HSV{180, 1, 0.5, 1},
        HSV{240, 1, 0.5, 1},
    }

    hslColorList = []HSL{
        HSL{  0, 0, 0, 1},
        HSL{  0, 0, 1, 1},
        HSL{  0, 1, 0.5, 1},
        HSL{120, 1, 0.5, 1},
        HSL{240, 1, 0.5, 1},
        HSL{ 60, 1, 0.5, 1},
        HSL{180, 1, 0.5, 1},
        HSL{300, 1, 0.5, 1},
        HSL{  0, 0, 0.75, 1},
        HSL{  0, 0, 0.5, 1},
        HSL{  0, 1, 0.25, 1},
        HSL{ 60, 1, 0.25, 1},
        HSL{120, 1, 0.25, 1},
        HSL{300, 1, 0.25, 1},
        HSL{180, 1, 0.25, 1},
        HSL{240, 1, 0.25, 1},
    }

    hsiColorList = []HSI{
        HSI{  0, 0, 0    , 1},
        HSI{  0, 0, 1    , 1},
        HSI{  0, 1, 0.333, 1},
        HSI{120, 1, 0.333, 1},
        HSI{240, 1, 0.333, 1},
        HSI{ 60, 1, 0.666, 1},
        HSI{180, 1, 0.666, 1},
        HSI{300, 1, 0.666, 1},
        HSI{  0, 0, 0.749, 1},
        HSI{  0, 0, 0.502, 1},
        HSI{  0, 1, 0.167, 1},
        HSI{ 60, 1, 0.335, 1},
        HSI{120, 1, 0.167, 1},
        HSI{300, 1, 0.335, 1},
        HSI{180, 1, 0.335, 1},
        HSI{240, 1, 0.167, 1},
    }

    rnd *rand.Rand
)

func init() {
    seed := rand.NewSource(123_456_789_012)
    rnd   = rand.New(seed)

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

func BenchmarkHSL2RGB(bench *testing.B) {
    c := HSL{360.0*rnd.Float64(), rnd.Float64(), rnd.Float64(), rnd.Float64()}
    for i:=0; i<bench.N; i++ {
        r, g, b, a = c.RGBA()
    }
}

func BenchmarkHSV2RGB(bench *testing.B) {
    c := HSV{360.0*rnd.Float64(), rnd.Float64(), rnd.Float64(), rnd.Float64()}
    for i:=0; i<bench.N; i++ {
        r, g, b, a = c.RGBA()
    }
}

func BenchmarkHSI2RGB(bench *testing.B) {
    c := HSI{360.0*rnd.Float64(), rnd.Float64(), rnd.Float64(), rnd.Float64()}
    for i:=0; i<bench.N; i++ {
        r, g, b, a = c.RGBA()
    }
}

// func BenchmarkHSV(bench *testing.B) {
//     c := Color{R, G, B, A}
//     for i:=0; i<bench.N; i++ {
//         h, s, v = c.HSV()
//     }
// }

// func BenchmarkHSL(bench *testing.B) {
//     c := Color{R, G, B, A}
//     for i:=0; i<bench.N; i++ {
//         h, s, l = c.HSL()
//     }
// }

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


func TestHSV(test *testing.T) {
    for i := range rgbColorList {
        rgbColor := rgbColorList[i]
        hsvColor := hsvColorList[i]
        
        convRgbColor := color.RGBAModel.Convert(hsvColor).(color.RGBA)
        convHsvColor := HSVModel.Convert(rgbColor).(HSV)
        
        if ! CompRGBA(rgbColor, convRgbColor) {
            test.Errorf("[%d]\n", i)
            test.Errorf("  want: %#v\n", rgbColor)
            test.Errorf("  got : %#v\n", convRgbColor)
        }
        if ! CompHSV(hsvColor, convHsvColor) {
            test.Errorf("[%d]\n", i)
            test.Errorf("  want: %#v\n", hsvColor)
            test.Errorf("  got : %#v\n", convHsvColor)
        }
    }
}

func TestHSL(test *testing.T) {
    for i := range rgbColorList {
        rgbColor := rgbColorList[i]
        hslColor := hslColorList[i]
        
        convRgbColor := color.RGBAModel.Convert(hslColor).(color.RGBA)
        convHslColor := HSLModel.Convert(rgbColor).(HSL)
        
        if ! CompRGBA(rgbColor, convRgbColor) {
            test.Errorf("[%d]\n", i)
            test.Errorf("  want: %#v\n", rgbColor)
            test.Errorf("  got : %#v\n", convRgbColor)
        }
        if ! CompHSL(hslColor, convHslColor) {
            test.Errorf("[%d]\n", i)
            test.Errorf("  want: %#v\n", hslColor)
            test.Errorf("  got : %#v\n", convHslColor)
        }
    }
}

func TestHSI(test *testing.T) {
    for i := range rgbColorList {
        rgbColor := rgbColorList[i]
        hsiColor := hsiColorList[i]
        
        convRgbColor := color.RGBAModel.Convert(hsiColor).(color.RGBA)
        convHsiColor := HSIModel.Convert(rgbColor).(HSI)
        
        if ! CompRGBA(rgbColor, convRgbColor) {
            test.Errorf("[%d]\n", i)
            test.Errorf("  want: %#v\n", rgbColor)
            test.Errorf("  got : %#v\n", convRgbColor)
        }
        if ! CompHSI(hsiColor, convHsiColor) {
            test.Errorf("[%d]\n", i)
            test.Errorf("  want: %#v\n", hsiColor)
            test.Errorf("  got : %#v\n", convHsiColor)
        }
    }
}

