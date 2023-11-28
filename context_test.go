package gg

import (
    "path"
    "crypto/md5"
    "flag"
    "fmt"
    "image/color"
    "math/rand"
    "os"
    "testing"
)

var save bool
var dirName string = "test"

func init() {
    flag.BoolVar(&save, "save", true, "save PNG output for each test case")
    os.Mkdir(dirName, 0755)
}

func hash(dc *Context) string {
    return fmt.Sprintf("%x", md5.Sum(dc.im.Pix))
}

func checkHash(t *testing.T, dc *Context, expected string) {
    actual := hash(dc)
    if actual != expected {
        t.Fatalf("expected hash: %s != actual hash: %s", expected, actual)
    }
}

func saveImage(dc *Context, name string) error {
    if save {
        return SavePNG(path.Join(dirName, name+".png"), dc.Image())
    }
    return nil
}

func TestBlank(t *testing.T) {
    dc := NewContext(100, 100)
    saveImage(dc, "TestBlank")
    checkHash(t, dc, "4e0a293a5b638f0aba2c4fe2c3418d0e")
}

func TestGrid(t *testing.T) {
    dc := NewContext(100, 100)
    dc.SetFillColor(color.White)
    dc.Clear()
    for i := 10; i < 100; i += 10 {
        x := float64(i) + 0.5
        dc.DrawLine(x, 0, x, 100)
        dc.DrawLine(0, x, 100, x)
    }
    dc.SetStrokeColor(color.Black)
    dc.Stroke()
    saveImage(dc, "TestGrid")
    checkHash(t, dc, "78606adda71d8abfbd8bb271087e4d69")
}

func TestLines(t *testing.T) {
    dc := NewContext(100, 100)
    dc.SetFillColor(NewRGB(0.5, 0.5, 0.5))
    dc.Clear()
    rnd := rand.New(rand.NewSource(99))
    for i := 0; i < 100; i++ {
        x1 := rnd.Float64() * 100
        y1 := rnd.Float64() * 100
        x2 := rnd.Float64() * 100
        y2 := rnd.Float64() * 100
        dc.DrawLine(x1, y1, x2, y2)
        dc.SetLineWidth(rnd.Float64() * 3)
        dc.SetStrokeColor(NewRGB(rnd.Float64(), rnd.Float64(), rnd.Float64()))
        dc.Stroke()
    }
    saveImage(dc, "TestLines")
    checkHash(t, dc, "1dfa8c19ba801bf5af2e464aa1f46518")
}

func TestCircles(t *testing.T) {
    dc := NewContext(100, 100)
    dc.SetFillColor(color.White)
    dc.Clear()
    rnd := rand.New(rand.NewSource(99))
    for i := 0; i < 10; i++ {
        x := rnd.Float64() * 100
        y := rnd.Float64() * 100
        r := rnd.Float64()*10 + 5
        dc.DrawCircle(x, y, r)
        dc.SetFillColor(NewRGB(rnd.Float64(), rnd.Float64(), rnd.Float64()))
        dc.SetStrokeColor(NewRGB(rnd.Float64(), rnd.Float64(), rnd.Float64()))
        dc.SetLineWidth(rnd.Float64() * 3)
        dc.FillStroke()
    }
    saveImage(dc, "TestCircles")
    checkHash(t, dc, "f7d9d71b15f21ee0a808489e92f16cc3")
}

func TestQuadratic(t *testing.T) {
    dc := NewContext(100, 100)
    dc.SetFillColor(NewRGB(0.25, 0.25, 0.25))
    dc.Clear()
    rnd := rand.New(rand.NewSource(99))
    for i := 0; i < 100; i++ {
        x1 := rnd.Float64() * 100
        y1 := rnd.Float64() * 100
        x2 := rnd.Float64() * 100
        y2 := rnd.Float64() * 100
        x3 := rnd.Float64() * 100
        y3 := rnd.Float64() * 100
        dc.MoveTo(x1, y1)
        dc.QuadraticTo(x2, y2, x3, y3)
        dc.SetLineWidth(rnd.Float64() * 3)
        dc.SetStrokeColor(NewRGB(rnd.Float64(), rnd.Float64(), rnd.Float64()))
        dc.Stroke()
    }
    saveImage(dc, "TestQuadratic")
    checkHash(t, dc, "e5f6caacbd63600e98d08250071d0e34")
}

func TestQuadraticSingle(t *testing.T) {
    dc := NewContext(100, 100)
    dc.SetFillColor(NewRGB(0.25, 0.25, 0.25))
    dc.Clear()
    rnd := rand.New(rand.NewSource(99))
    x1 := 25.0
    y1 := 75.0
    x2 := 50.0
    y2 := 25.0
    x3 := 75.0
    y3 := 75.0
    dc.MoveTo(x1, y1)
    dc.QuadraticTo(x2, y2, x3, y3)
    dc.SetLineWidth(rnd.Float64() * 3)
    dc.SetStrokeColor(NewRGB(rnd.Float64(), rnd.Float64(), rnd.Float64()))
    dc.Stroke()
    saveImage(dc, "TestQuadraticSingle")
}


func TestCubic(t *testing.T) {
    dc := NewContext(100, 100)
    dc.SetFillColor(NewRGB(0.75, 0.75, 0.75))
    dc.Clear()
    rnd := rand.New(rand.NewSource(99))
    for i := 0; i < 100; i++ {
        x1 := rnd.Float64() * 100
        y1 := rnd.Float64() * 100
        x2 := rnd.Float64() * 100
        y2 := rnd.Float64() * 100
        x3 := rnd.Float64() * 100
        y3 := rnd.Float64() * 100
        x4 := rnd.Float64() * 100
        y4 := rnd.Float64() * 100
        dc.MoveTo(x1, y1)
        dc.CubicTo(x2, y2, x3, y3, x4, y4)
        dc.SetLineWidth(rnd.Float64() * 3)
        dc.SetStrokeColor(NewRGB(rnd.Float64(), rnd.Float64(), rnd.Float64()))
        dc.Stroke()
    }
    saveImage(dc, "TestCubic")
    checkHash(t, dc, "6d8ac746291488a6cad99c53e415982e")
}

func TestCubicSingle(t *testing.T) {
    dc := NewContext(100, 100)
    dc.SetFillColor(NewRGB(0.25, 0.25, 0.25))
    dc.Clear()
    rnd := rand.New(rand.NewSource(99))
    x1 := 25.0
    y1 := 75.0
    x2 := 25.0
    y2 := 25.0
    x3 := 75.0
    y3 := 25.0
    x4 := 75.0
    y4 := 75.0
    dc.MoveTo(x1, y1)
    dc.CubicTo(x2, y2, x3, y3, x4, y4)
    dc.SetLineWidth(rnd.Float64() * 3)
    dc.SetStrokeColor(NewRGB(rnd.Float64(), rnd.Float64(), rnd.Float64()))
    dc.Stroke()
    saveImage(dc, "TestCubicSingle")
}


func TestFill(t *testing.T) {
    dc := NewContext(100, 100)
    dc.SetFillColor(color.White)
    dc.Clear()
    rnd := rand.New(rand.NewSource(99))
    for i := 0; i < 10; i++ {
        dc.NewSubPath()
        for j := 0; j < 10; j++ {
            x := rnd.Float64() * 100
            y := rnd.Float64() * 100
            dc.LineTo(x, y)
        }
        dc.ClosePath()
        dc.SetFillColor(NewRGB(rnd.Float64(), rnd.Float64(), rnd.Float64()))
        dc.Fill()
    }
    saveImage(dc, "TestFill")
    checkHash(t, dc, "3a2e05e39000f9a0761e67d1049aa486")
}

func TestClip(t *testing.T) {
    dc := NewContext(100, 100)
    dc.SetFillColor(NewRGB(1, 1, 1))
    dc.Clear()
    dc.DrawCircle(50, 50, 40)
    dc.Clip()
    rnd := rand.New(rand.NewSource(99))
    for i := 0; i < 1000; i++ {
        x := rnd.Float64() * 100
        y := rnd.Float64() * 100
        r := rnd.Float64()*10 + 5
        dc.DrawCircle(x, y, r)
        dc.SetFillColor(NewRGBA(rnd.Float64(), rnd.Float64(), rnd.Float64(), rnd.Float64()))
        dc.Fill()
    }
    saveImage(dc, "TestClip")
    checkHash(t, dc, "bb4dd1b0fbbfb82fa27862cbea015582")
}

func TestPushPop(t *testing.T) {
    const S = 100
    dc := NewContext(S, S)
    dc.SetFillColor(NewRGBA(0, 0, 0, 0.1))
    for i := 0; i < 360; i += 15 {
        dc.Push()
        dc.RotateAbout(Radians(float64(i)), S/2, S/2)
        dc.DrawEllipse(S/2, S/2, S*7/16, S/8)
        dc.Fill()
        dc.Pop()
    }
    saveImage(dc, "TestPushPop")
    checkHash(t, dc, "98813dcbd31ca163aed034743cdb1918")
}

func TestDrawStringWrapped(t *testing.T) {
    dc := NewContext(100, 100)
    dc.SetFillColor(color.White)
    dc.Clear()
    dc.SetStrokeColor(color.Black)
    dc.DrawStringWrapped("Hello, world! How are you?", 50, 50, 0.5, 0.5, 90, 1.5, AlignCenter)
    saveImage(dc, "TestDrawStringWrapped")
    checkHash(t, dc, "8d92f6aae9e8b38563f171abd00893f8")
}

func TestDrawImage(t *testing.T) {
    src := NewContext(100, 100)
    src.SetFillColor(color.White)
    src.Clear()
    for i := 10; i < 100; i += 10 {
        x := float64(i) + 0.5
        src.DrawLine(x, 0, x, 100)
        src.DrawLine(0, x, 100, x)
    }
    src.SetStrokeColor(color.Black)
    src.Stroke()

    dc := NewContext(200, 200)
    dc.SetFillColor(color.Black)
    dc.Clear()
    dc.DrawImage(src.Image(), 50, 50)
    saveImage(dc, "TestDrawImage")
    checkHash(t, dc, "282afbc134676722960b6bec21305b15")
}

func TestSetPixel(t *testing.T) {
    dc := NewContext(100, 100)
    dc.SetFillColor(color.Black)
    dc.Clear()
    dc.SetStrokeColor(NewRGB(0, 1, 0))
    i := 0
    for y := 0; y < 100; y++ {
        for x := 0; x < 100; x++ {
            if i%31 == 0 {
                dc.SetPixel(x, y)
            }
            i++
        }
    }
    saveImage(dc, "TestSetPixel")
    checkHash(t, dc, "27dda6b4b1d94f061018825b11982793")
}

func TestDrawPoint(t *testing.T) {
    dc := NewContext(100, 100)
    dc.SetFillColor(color.Black)
    dc.Clear()
    dc.SetFillColor(NewRGB(0, 1, 0))
    dc.Scale(10, 10)
    for y := 0; y <= 10; y++ {
        for x := 0; x <= 10; x++ {
            dc.DrawPoint(float64(x), float64(y), 3)
            dc.Fill()
        }
    }
    saveImage(dc, "TestDrawPoint")
    checkHash(t, dc, "a4e7546ed558cdf186e00fb6716b91bc")
}

func TestLinearGradient(t *testing.T) {
    dc := NewContext(100, 100)
    g := NewLinearGradient(0, 0, 100, 100)
    g.AddColorStop(0, color.RGBA{0, 255, 0, 255})
    g.AddColorStop(1, color.RGBA{0, 0, 255, 255})
    g.AddColorStop(0.5, color.RGBA{255, 0, 0, 255})
    dc.SetFillStyle(g)
    dc.DrawRectangle(0, 0, 100, 100)
    dc.Fill()
    saveImage(dc, "TestLinearGradient")
    checkHash(t, dc, "75eb9385c1219b1d5bb6f4c961802c7a")
}

func TestRadialGradient(t *testing.T) {
    dc := NewContext(100, 100)
    g := NewRadialGradient(30, 50, 0, 70, 50, 50)
    g.AddColorStop(0, color.RGBA{0, 255, 0, 255})
    g.AddColorStop(1, color.RGBA{0, 0, 255, 255})
    g.AddColorStop(0.5, color.RGBA{255, 0, 0, 255})
    dc.SetFillStyle(g)
    dc.DrawRectangle(0, 0, 100, 100)
    dc.Fill()
    saveImage(dc, "TestRadialGradient")
    checkHash(t, dc, "f170f39c3f35c29de11e00428532489d")
}

func TestDashes(t *testing.T) {
    dc := NewContext(100, 100)
    dc.SetFillColor(color.White)
    dc.Clear()
    rnd := rand.New(rand.NewSource(99))
    for i := 0; i < 100; i++ {
        x1 := rnd.Float64() * 100
        y1 := rnd.Float64() * 100
        x2 := rnd.Float64() * 100
        y2 := rnd.Float64() * 100
        dc.SetDash(rnd.Float64()*3+1, rnd.Float64()*3+3)
        dc.DrawLine(x1, y1, x2, y2)
        dc.SetLineWidth(rnd.Float64() * 3)
        dc.SetStrokeColor(NewRGB(rnd.Float64(), rnd.Float64(), rnd.Float64()))
        dc.Stroke()
    }
    saveImage(dc, "TestDashes")
    checkHash(t, dc, "2c4c6e23ae4219f1dbfd8c3ea5a8be68")
}

func BenchmarkCircles(b *testing.B) {
    dc := NewContext(1000, 1000)
    dc.SetFillColor(color.White)
    dc.Clear()
    rnd := rand.New(rand.NewSource(99))
    for i := 0; i < b.N; i++ {
        x := rnd.Float64() * 1000
        y := rnd.Float64() * 1000
        dc.DrawCircle(x, y, 10)
        if i%2 == 0 {
            dc.SetFillColor(color.Black)
        } else {
            dc.SetFillColor(color.White)
        }
        dc.Fill()
    }
    b.StopTimer()
    saveImage(dc, "BenchmarkCircles")
}

func BenchmarkMathCircles(b *testing.B) {
    dc := NewContext(1000, 1000)
    dc.Translate(500, 500)
    dc.Scale(500, -500)
    dc.SetFillColor(color.White)
    dc.Clear()
    rnd := rand.New(rand.NewSource(99))
    for i := 0; i < b.N; i++ {
        x := rnd.Float64() * 2.0 - 1.0
        y := rnd.Float64() * 2.0 - 1.0
        dc.DrawCircle(x, y, 0.02)
        if i%2 == 0 {
            dc.SetFillColor(color.Black)
        } else {
            dc.SetFillColor(color.White)
        }
        dc.Fill()
    }
    b.StopTimer()
    saveImage(dc, "BenchmarkMathCircles")
}
