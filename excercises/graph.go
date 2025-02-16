package main

import (
	"math"

	"github.com/stefan-muehlebach/gg"
	"github.com/stefan-muehlebach/gg/color"
)

var (
	// Parametrierung des Output-Bildes (PNG)
	Width, Height = 512, 512
	MarginSize    = 24
	GraphSize     = Width - 2*MarginSize

	// Parameter fuer die mathem. Groessen des Koordinatensystems.
	xMin, xMax = -1.0, 1.0
	yMin, yMax = -1.0, 1.0
	zMin, zMax = 1.0, -1.0

	BackColor = color.RGBAF{0.851, 0.811, 0.733, 1.0}
	LineColor = color.RGBAF{0.153, 0.157, 0.133, 1.0}
)

// Alle anzuzeigenden Objekte muessen folgendes Interface implementieren.
type GraphObj interface {
	// Mit Z wird die Hoehe des Objektes beim Punkt (x,y) ermittelt. Falls
	// sich das Objekt gar nicht an dieser Stelle befindet, muss 0.0
	// retourniert werden.
	Z(x, y float64) float64
	// Mit ZRange kann das Intervall in Z bestimmt werden, welches vom Objekt
	// belegt wird. Damit lassen sich bspw. Skalierungen vornehmen.
	ZRange() (float64, float64)
}

// Eine erste, einfache Objektart sind Kugeln. Sie werden durch einen
// Mittelpunkt in R^3 und einen Radius spezifiziert.
type Sphere struct {
	xm, ym, zm float64
	r          float64
}

func (s *Sphere) Z(x, y float64) float64 {
	x2 := (x - s.xm) * (x - s.xm)
	y2 := (y - s.ym) * (y - s.ym)
	r2 := s.r * s.r
	if r2 < x2+y2 {
		return math.Inf(-1)
	}
	return math.Sqrt(r2-x2-y2) + s.zm
}

func (s *Sphere) ZRange() (float64, float64) {
	return s.zm, s.zm + s.r
}

type Ellipse struct {
	xm, ym, zm float64
	a, b, c    float64
}

func (s *Ellipse) Z(x, y float64) float64 {
	x2 := (x - s.xm) / s.a
	x2 *= x2
	y2 := (y - s.ym) / s.b
	y2 *= y2
	fz := x2 + y2
	if fz > 1.0 {
		return math.Inf(-1)
	}
	return math.Sqrt(1.0-fz)*s.c + s.zm
}

func (s *Ellipse) ZRange() (float64, float64) {
	return s.zm, s.zm + s.c
}

// Roehren oder Tubes sind eine weitere Objektart, die sehr gut zur Geltung
// kommen. Sie haben eine Anfangs- und Endkoordinate (in R^3, die ausserhalb
// des definierten Bereiches liegen sollten. Ausserdem haben sie einen Radius.
type Tube struct {
	x0, y0, z0 float64
	x1, y1, z1 float64
	r          float64
}

func (t *Tube) Z(x, y float64) float64 {
	d := 0.0
	if t.y0 == t.y1 {
		d = y - t.y0
	} else {
		d = x - t.x0
	}
	if d > t.r || d < -t.r {
		return math.Inf(-1)
	}
	return math.Sqrt(t.r*t.r-d*d) + t.z0
}

func (t *Tube) ZRange() (float64, float64) {
	return t.z0, t.z0 + t.r
}

func scaleZ(z float64) float64 {
	return (z - zMin) / (zMax - zMin)
}
func colToX(col int) float64 {
	return xMin + (float64(col)/float64(GraphSize))*(xMax-xMin)
}
func rowToY(row int) float64 {
	return yMax - (float64(row)/float64(GraphSize))*(yMax-yMin)
}

var (
	tubes = []GraphObj{
		&Tube{-0.2, -1.0, 0.7, -0.2, 1.0, 0.7, 0.15},
		&Tube{0.6, -1.0, 0.7, 0.6, 1.0, 0.7, 0.2},
		&Tube{-1.0, -0.5, 0.7, 1.0, -0.5, 0.7, 0.25},
		&Tube{-1.0, 0.15, 0.7, 1.0, 0.15, 0.7, 0.2},
	}

	ellipses = []GraphObj{
		&Ellipse{xm: 0.0, ym: 0.0, zm: 0.5, a: 0.9, b: 0.3, c: 0.3},
		&Ellipse{xm: 0.0, ym: 0.0, zm: 0.45, a: 0.3, b: 0.9, c: 0.3},
		&Sphere{xm: 0.6, ym: 0.6, zm: 0.5, r: 0.2},
		&Sphere{xm: -0.6, ym: 0.6, zm: 0.5, r: 0.2},
		&Sphere{xm: 0.6, ym: -0.6, zm: 0.5, r: 0.2},
		&Sphere{xm: -0.6, ym: -0.6, zm: 0.5, r: 0.2},
	}

	tubedSpheres = []GraphObj{
		&Tube{-1.0, 0.0, 0.5, 1.0, 0.0, 0.5, 0.2},
		&Sphere{xm: -0.6, ym: 0.1, zm: 0.4, r: 0.3},
		&Sphere{xm: 0.0, ym: 0.2, zm: 0.4, r: 0.3},
		&Sphere{xm: 0.6, ym: 0.3, zm: 0.4, r: 0.3},
    }

	spheres = []GraphObj{
		&Sphere{xm: -0.2, ym: -0.2, zm: 0.5, r: 0.7},
		&Sphere{xm:  0.3, ym:  0.3, zm: 0.5, r: 0.6},
        &Sphere{xm:  0.4, ym: -0.4, zm: 0.5, r: 0.55},
	}


)

func main() {
	gc := gg.NewContext(Width, Height)
	gc.SetFillColor(BackColor)
	gc.Clear()

	objList := spheres
    fileName := "graph04.png"

	for _, obj := range objList {
		zmin, zmax := obj.ZRange()
		zMin = min(zMin, zmin)
		zMax = max(zMax, zmax)
	}

	// fmt.Printf("zMin, zMax: %f, %f\n", zMin, zMax)

	pixelCol := color.RGBAF{}
	col0, row0 := int(MarginSize), int(Height-MarginSize)
	for row := 0; row < GraphSize; row++ {
		y := rowToY(row)
		for col := 0; col < GraphSize; col++ {
			x := colToX(col)
			z := math.Inf(-1)
			for _, obj := range objList {
				z = max(z, obj.Z(x, y))
			}
			if math.IsInf(z, 0) {
				pixelCol = color.RGBAF{0.0, 0.0, 0.0, 1.0}
			} else {
				z = scaleZ(z)
				pixelCol = color.RGBAF{z, z, z, 1.0}
			}
			c := col0 + col
			r := row0 - row

			gc.SetPixel(c, r, pixelCol)
		}
	}
	gc.SavePNG(fileName)
}
