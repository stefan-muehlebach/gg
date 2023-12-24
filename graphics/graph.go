package main

import (
	"github.com/stefan-muehlebach/gg"
	"github.com/stefan-muehlebach/gg/color"
	"log"
	"math"
)

const (
	Width, Height = 512.0, 512.0
	MarginSize    = 24.0
	GraphSize     = Width - 2*MarginSize

	// Parameter fuer die Halbkugel
	xMin, xMax = -1.0, 1.0
	yMin, yMax = -1.0, 1.0
	zMin, zMax = 0.0, 0.7

	// Parameter fuer den Kardinal-Sinus
	// xMin, xMax = -15.0, 15.0
	// yMin, yMax = -15.0, 15.0
	// zMin, zMax = -0.25,  1.0
)

var (
	BackColor = color.RGBAF{0.851, 0.811, 0.733, 1.0}
	LineColor = color.RGBAF{0.153, 0.157, 0.133, 1.0}
)

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

func sphereFunc(x, y float64) float64 {
	return math.Sqrt(4.0/9.0 - x*x - y*y)
}

func cardSinus(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}

func main() {
	gc := gg.NewContext(Width, Height)
	gc.SetFillColor(BackColor)
	gc.Clear()

	col0, row0 := int(MarginSize), int(Height-MarginSize)
	for row := 0; row < GraphSize; row++ {
		y := rowToY(row)
		for col := 0; col < GraphSize; col++ {
			x := colToX(col)
			z := sphereFunc(x, y)
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
