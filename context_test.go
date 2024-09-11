package gg

import (
	"crypto/md5"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"path"
	"testing"

	"github.com/stefan-muehlebach/gg/color"
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
	dc.SetFillColor(color.RGBAF{0.5, 0.5, 0.5, 1.0})
	dc.Clear()
	rnd := rand.New(rand.NewSource(99))
	for i := 0; i < 100; i++ {
		x1 := rnd.Float64() * 100
		y1 := rnd.Float64() * 100
		x2 := rnd.Float64() * 100
		y2 := rnd.Float64() * 100
		dc.DrawLine(x1, y1, x2, y2)
		dc.SetStrokeWidth(rnd.Float64() * 3)
		dc.SetStrokeColor(color.RGBAF{rnd.Float64(), rnd.Float64(), rnd.Float64(), 1.0})
		dc.Stroke()
	}
	saveImage(dc, "TestLines")
	checkHash(t, dc, "72735880e6b28b6351ea8f7d51c10193")
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
		dc.SetFillColor(color.RGBAF{rnd.Float64(), rnd.Float64(), rnd.Float64(), 1.0})
		dc.SetStrokeColor(color.RGBAF{rnd.Float64(), rnd.Float64(), rnd.Float64(), 1.0})
		dc.SetStrokeWidth(rnd.Float64() * 3)
		dc.FillStroke()
	}
	saveImage(dc, "TestCircles")
	checkHash(t, dc, "26986ecaac3136c56764c3c3c60ac12a")
}

func TestQuadratic(t *testing.T) {
	dc := NewContext(100, 100)
	dc.SetFillColor(color.RGBAF{0.25, 0.25, 0.25, 1.0})
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
		dc.SetStrokeWidth(rnd.Float64() * 3)
		dc.SetStrokeColor(color.RGBAF{rnd.Float64(), rnd.Float64(), rnd.Float64(), 1.0})
		dc.Stroke()
	}
	saveImage(dc, "TestQuadratic")
	checkHash(t, dc, "1fb834e5e648f8bbd592f2c07fbf608e")
}

func TestQuadraticSingle(t *testing.T) {
	dc := NewContext(100, 100)
	dc.SetFillColor(color.RGBAF{0.25, 0.25, 0.25, 1.0})
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
	dc.SetStrokeWidth(rnd.Float64() * 3)
	dc.SetStrokeColor(color.RGBAF{rnd.Float64(), rnd.Float64(), rnd.Float64(), 1.0})
	dc.Stroke()
	saveImage(dc, "TestQuadraticSingle")
}

func TestCubic(t *testing.T) {
	dc := NewContext(100, 100)
	dc.SetFillColor(color.RGBAF{0.75, 0.75, 0.75, 1.0})
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
		dc.SetStrokeWidth(rnd.Float64() * 3)
		dc.SetStrokeColor(color.RGBAF{rnd.Float64(), rnd.Float64(), rnd.Float64(), 1.0})
		dc.Stroke()
	}
	saveImage(dc, "TestCubic")
	checkHash(t, dc, "35b720d2ed0de31f58a000e1ea060901")
}

func TestCubicSingle(t *testing.T) {
	dc := NewContext(100, 100)
	dc.SetFillColor(color.RGBAF{0.25, 0.25, 0.25, 1.0})
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
	dc.SetStrokeWidth(rnd.Float64() * 3)
	dc.SetStrokeColor(color.RGBAF{rnd.Float64(), rnd.Float64(), rnd.Float64(), 1.0})
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
		dc.SetFillColor(color.RGBAF{rnd.Float64(), rnd.Float64(), rnd.Float64(), 1.0})
		dc.Fill()
	}
	saveImage(dc, "TestFill")
	checkHash(t, dc, "9b7600ba47a8a2a10657dea77a171f3b")
}

func TestClip(t *testing.T) {
	dc := NewContext(100, 100)
	dc.SetFillColor(color.RGBAF{1, 1, 1, 1.0})
	dc.Clear()
	dc.DrawCircle(50, 50, 40)
	dc.Clip()
	rnd := rand.New(rand.NewSource(99))
	for i := 0; i < 1000; i++ {
		x := rnd.Float64() * 100
		y := rnd.Float64() * 100
		r := rnd.Float64()*10 + 5
		dc.DrawCircle(x, y, r)
		dc.SetFillColor(color.RGBAF{rnd.Float64(), rnd.Float64(), rnd.Float64(), rnd.Float64()})
		dc.Fill()
	}
	saveImage(dc, "TestClip")
	checkHash(t, dc, "508f01f35249c17fb06b43a0ba819faa")
}

func TestPushPop(t *testing.T) {
	const S = 100
	dc := NewContext(S, S)
	dc.SetFillColor(color.RGBAF{0, 0, 0, 0.1})
	for i := 0; i < 360; i += 15 {
		dc.Push()
		dc.RotateAbout(Radians(float64(i)), S/2, S/2)
		dc.DrawEllipse(S/2, S/2, S*7/16, S/8)
		dc.Fill()
		dc.Pop()
	}
	saveImage(dc, "TestPushPop")
	checkHash(t, dc, "bb36bd0b91b2f10e9e55bad88e18f437")
}

func TestDrawStringWrapped(t *testing.T) {
	gc := NewContext(100, 100)
	DrawStringWrapped(gc, 1)
	saveImage(gc, "TestDrawStringWrapped")
	checkHash(t, gc, "bfa8bd15395510b453e3b9d075a1a66a")
}

func BenchmarkDrawStringWrapped(b *testing.B) {
	gc := NewContext(100, 100)
	b.ResetTimer()
	DrawStringWrapped(gc, b.N)
}

func DrawStringWrapped(gc *Context, num int) {
	for range num {
		gc.SetFillColor(color.White)
		gc.Clear()
		gc.SetTextColor(color.Teal)
		gc.DrawStringWrapped("Hello, world! How are you?",
                    50, 50, 0.5, 0.5, 90, 1.5, AlignCenter)
	}
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
	i := 0
	for y := 0; y < 100; y++ {
		for x := 0; x < 100; x++ {
			if i%31 == 0 {
				dc.SetPixel(x, y, color.RGBAF{0, 1, 0, 1.0})
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
	dc.SetFillColor(color.RGBAF{0, 1, 0, 1.0})
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
	gc := NewContext(100, 100)
	DrawLinearGradient(gc, 1)
	saveImage(gc, "TestLinearGradient")
	checkHash(t, gc, "75eb9385c1219b1d5bb6f4c961802c7a")
}

func BenchmarkLinearGradient(b *testing.B) {
	gc := NewContext(100, 100)
	b.ResetTimer()
	DrawLinearGradient(gc, b.N)
}

func DrawLinearGradient(gc *Context, num int) {
	for range num {
		g := NewLinearGradient(0, 0, 100, 100)
		g.AddColorStop(0.0, color.RGBAF{0, 1, 0, 1})
		g.AddColorStop(0.5, color.RGBAF{1, 0, 0, 1})
		g.AddColorStop(1.0, color.RGBAF{0, 0, 1, 1})
		gc.SetFillStyle(g)
		gc.DrawRectangle(0, 0, 100, 100)
		gc.Fill()
	}
}

func TestRadialGradient(t *testing.T) {
	gc := NewContext(100, 100)
	DrawRadialGradient(gc, 1)
	saveImage(gc, "TestRadialGradient")
	checkHash(t, gc, "f170f39c3f35c29de11e00428532489d")
}

func BenchmarkRadialGradient(b *testing.B) {
	gc := NewContext(100, 100)
	b.ResetTimer()
	DrawRadialGradient(gc, b.N)
}

func DrawRadialGradient(gc *Context, num int) {
	for range num {
		g := NewRadialGradient(30, 50, 0, 70, 50, 50)
		g.AddColorStop(0.0, color.RGBAF{0, 1, 0, 1})
		g.AddColorStop(0.5, color.RGBAF{1, 0, 0, 1})
		g.AddColorStop(1.0, color.RGBAF{0, 0, 1, 1})
		gc.SetFillStyle(g)
		gc.DrawRectangle(0, 0, 100, 100)
		gc.Fill()
	}
}

func TestDashes(t *testing.T) {
	gc := NewContext(100, 100)
	DrawDashes(gc, 100)
	saveImage(gc, "TestDashes")
	checkHash(t, gc, "89669e7e03a08ca9fc7ba589a310f427")
}

func BenchmarkDashes(b *testing.B) {
	gc := NewContext(100, 100)
	b.ResetTimer()
	DrawDashes(gc, b.N)
}

func DrawDashes(gc *Context, num int) {
	gc.SetFillColor(color.White)
	gc.Clear()
	rnd := rand.New(rand.NewSource(99))
	for range num {
		x1 := rnd.Float64() * 100
		y1 := rnd.Float64() * 100
		x2 := rnd.Float64() * 100
		y2 := rnd.Float64() * 100
		gc.SetDash(rnd.Float64()*3+1, rnd.Float64()*3+3)
		gc.DrawLine(x1, y1, x2, y2)
		gc.SetStrokeWidth(rnd.Float64() * 3)
		gc.SetStrokeColor(color.RGBAF{rnd.Float64(), rnd.Float64(), rnd.Float64(), 1.0})
		gc.Stroke()
	}
}

func BenchmarkScreenCoordSystem(b *testing.B) {
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
	// saveImage(dc, "BenchmarkCircles")
}

func BenchmarkMathCoordSystem(b *testing.B) {
	dc := NewContext(1000, 1000)
	dc.Translate(500, 500)
	dc.Scale(500, -500)
	dc.SetFillColor(color.White)
	dc.Clear()
	rnd := rand.New(rand.NewSource(99))
	for i := 0; i < b.N; i++ {
		x := rnd.Float64()*2.0 - 1.0
		y := rnd.Float64()*2.0 - 1.0
		dc.DrawCircle(x, y, 0.02)
		if i%2 == 0 {
			dc.SetFillColor(color.Black)
		} else {
			dc.SetFillColor(color.White)
		}
		dc.Fill()
	}
	b.StopTimer()
	// saveImage(dc, "BenchmarkMathCircles")
}
