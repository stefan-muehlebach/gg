package main

import (
	"image"
	"image/png"
	"log"
	"math"
	"os"
)

var (
	imgWidth    = 512
	imgHeight   = 512
	pngFileName = "colorWheel.png"

	imgBounds = image.Rect(0, 0, imgWidth, imgHeight)
	midPoint  = image.Point{imgWidth / 2, imgHeight / 2}
	maxDist   = math.Hypot(float64(midPoint.X), float64(midPoint.Y))
)

//----------------------------------------------------------------------------

type HSV struct {
	H, S, V, A float64
}

func (c HSV) RGBA() (r, g, b, a uint32) {
	C := c.V * c.S
	X := C * (1.0 - math.Abs(math.Mod(c.H/60.0, 2.0)-1.0))
	m := c.V - C
	R, G, B := 0.0, 0.0, 0.0
	switch {
	case c.H < 60.0:
		R, G, B = C, X, 0.0
	case c.H < 120.0:
		R, G, B = X, C, 0.0
	case c.H < 180.0:
		R, G, B = 0.0, C, X
	case c.H < 240.0:
		R, G, B = 0.0, X, C
	case c.H < 300.0:
		R, G, B = X, 0.0, C
	case c.H < 360.0:
		R, G, B = C, 0.0, X
	}
	r = uint32(65535.0 * (R + m) * c.A)
	g = uint32(65535.0 * (G + m) * c.A)
	b = uint32(65535.0 * (B + m) * c.A)
	a = uint32(65535.0 * c.A)
	return
}

//----------------------------------------------------------------------------

func main() {
	img := image.NewRGBA(imgBounds)
	for row := range imgHeight {
		for col := range imgWidth {
			diffVec := image.Point{col, row}.Sub(midPoint)
			dist := math.Hypot(float64(diffVec.X), float64(diffVec.Y))
			angleRad := math.Atan2(float64(diffVec.Y), float64(diffVec.X))
			if angleRad < 0.0 {
				angleRad = 2.0*math.Pi + angleRad
			}
			if angleRad > 2.0*math.Pi {
				angleRad = angleRad - 2.0*math.Pi
			}
			H := 360.0 * angleRad / (2.0 * math.Pi)
			S := dist / maxDist
			V := 1.0
			color := HSV{H, S, V, 1.0}
			img.Set(col, row, color)
		}
	}

	fh, err := os.Create(pngFileName)
	if err != nil {
		log.Fatalf("Couldn't create file: %v", err)
	}
	defer fh.Close()
	if err := png.Encode(fh, img); err != nil {
		log.Fatalf("Couldn't encode image: %v", err)
	}
}
