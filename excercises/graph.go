package main

import (
	"log"
	"math"

	"github.com/stefan-muehlebach/gg"
	"github.com/stefan-muehlebach/gg/color"
)

const (
	Width, Height = 512.0, 512.0
	MarginSize    = 24.0
	GraphSize     = Width - 2*MarginSize

	// Parameter fuer die Halbkugel
	xMin, xMax = -1.0, 1.0
	yMin, yMax = -1.0, 1.0
	zMin, zMax = 0.0, 0.7
)

var (
	BackColor = color.RGBAF{0.851, 0.811, 0.733, 1.0}
	LineColor = color.RGBAF{0.153, 0.157, 0.133, 1.0}
)

// Alle Graphikobjekte muessen im Minimum die Methode Z implementieren,
// welche den maximalen Z-Wert des Objektes an der Stelle (x,y) berechnet.
type GraphObj interface {
	Z(x, y float64) float64
}

type Sphere struct {
	xm, ym, zm float64
	r      float64
}

func (s *Sphere) Z(x, y float64) float64 {
	x2 := (x - s.xm) * (x - s.xm)
	y2 := (y - s.ym) * (y - s.ym)
	r2 := s.r * s.r
    if r2 < x2 + y2 {
        return 0.0
    }
	return math.Sqrt(r2 - x2 - y2) + s.zm
}

type Tube struct {
	pos, zm        float64
	isVertical bool
	r          float64
}

func (t *Tube) Z(x, y float64) float64 {
	if t.isVertical {
		if x < t.pos-t.r || x > t.pos+t.r {
			return 0.0
		}
		return math.Sqrt(t.r*t.r - (x-t.pos)*(x-t.pos)) + t.zm
	} else {
		if y < t.pos-t.r || y > t.pos+t.r {
			return 0.0
		}
	}
	return math.Sqrt(t.r*t.r - (y-t.pos)*(y-t.pos)) + t.zm
}

func scaleZ(z float64) float64 {
	z = z - zMin
	z = z / (zMax - zMin)
	return z
}
func colToX(col int) float64 {
	return xMin + (float64(col)/float64(GraphSize))*(xMax-xMin)
}
func rowToY(row int) float64 {
	return yMax - (float64(row)/float64(GraphSize))*(yMax-yMin)
}

func main() {
	gc := gg.NewContext(Width, Height)
	gc.SetFillColor(BackColor)
	gc.Clear()

	objList := []GraphObj{
        &Tube{pos: 0.2, zm: 0.5, isVertical: true, r: 0.2},
        &Tube{pos: -0.7, zm: 0.8, isVertical: true, r: 0.2},
		&Sphere{xm: 0.4, ym: 0.4, zm: 0.1, r: 0.5},
		&Sphere{xm: 0.0, ym: 0.0, zm: 0.2, r: 0.5},
	}

	col0, row0 := int(MarginSize), int(Height-MarginSize)
	for row := 0; row < GraphSize; row++ {
		y := rowToY(row)
		for col := 0; col < GraphSize; col++ {
			x := colToX(col)
			z := 0.0
			for _, obj := range objList {
				z = max(z, obj.Z(x, y))
			}
			// fmt.Printf("(%.3f, %.3f): %f\n", x, y, z)
			z = scaleZ(z)
			if z > 1.0 || z < 0.0 {
				log.Printf("z is out of range: %f\n", z)
			}
			c := col0 + col
			r := row0 - row
			color := color.RGBAF{z, z, z, 1.0}
			gc.SetPixel(c, r, color)
		}
	}
	gc.SavePNG("graph.png")
}
