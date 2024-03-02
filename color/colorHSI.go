package color

import (
	"image/color"
	"log"
	"math"
)

// Der Typ HSI enthaelt die Werte fuer Hue, Saturation und Intensity gespeichert.
type HSI struct {
	H, S, I, A float64
}

func (c HSI) RGBA() (r, g, b, a uint32) {
	Z := 1.0 - math.Abs(math.Mod(c.H/60.0, 2.0)-1.0)
	C := (3 * c.I * c.S) / (1 + Z)
	X := C * Z
	m := c.I * (1 - c.S)
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
	R += m
	G += m
	B += m
	if R > 1.0 {
		R = 1.0
	}
	if G > 1.0 {
		G = 1.0
	}
	if B > 1.0 {
		B = 1.0
	}
	r = uint32(65535.0 * R * c.A)
	g = uint32(65535.0 * G * c.A)
	b = uint32(65535.0 * B * c.A)
	a = uint32(65535.0 * c.A)
	return
}

func (c HSI) Bright(t float64) Color {
	t = setIn(t, 0, 1)
	return c
}

func (c HSI) Dark(t float64) Color {
	t = setIn(t, 0, 1)
	return c
}

func (c HSI) Alpha(a float64) Color {
	a = setIn(a, 0, 1)
	return HSI{c.H, c.S, c.I, a}
}

func (c1 HSI) Interpolate(col Color, t float64) Color {
	t = setIn(t, 0, 1)
	c2 := col.(HSI)
	h := (1-t)*c1.H + t*c2.H
	s := (1-t)*c1.S + t*c2.S
	i := (1-t)*c1.I + t*c2.I
	a := (1-t)*c1.A + t*c2.A
	return HSI{h, s, i, a}
}

func (c1 HSI) Less(c2 HSI, key SortField) bool {
	switch key {
	case SortByHue:
		return c1.H < c2.H
	case SortBySaturation:
		return c1.S < c2.S
	case SortByIntensity:
		return c1.I < c2.I
	default:
		log.Fatalf("invalid sort field specified: '%v'", key)
		return false
	}
}

// Modell fuer den neuen Farbtyp, d.h. fuer die Konvertierung von einer
// beliebigen Farbe nach HSI.
var (
	HSIModel color.Model = color.ModelFunc(hsiModel)
)

func hsiModel(c color.Color) color.Color {
	if _, ok := c.(HSI); ok {
		return c
	}
	red, green, blue, alpha := c.RGBA()
	if alpha == 0 {
		return HSI{0.0, 0.0, 0.0, 0.0}
	}
	r := float64((red*0xffff)/alpha) / 65535.0
	g := float64((green*0xffff)/alpha) / 65535.0
	b := float64((blue*0xffff)/alpha) / 65535.0
	a := float64(alpha) / 65535.0

	h, s, i := 0.0, 0.0, 0.0

	max := max(r, g, b)
	min := min(r, g, b)
	d := max - min

	switch max {
	case min:
		h = 0.0
	case r:
		h = (g - b) / d
		if h < 0.0 {
			h += 6.0
		}
	case g:
		h = 2.0 + (b-r)/d
	case b:
		h = 4.0 + (r-g)/d
	}
	h *= 60.0

	i = (r + g + b) / 3.0

	if i > 0.0 {
		s = 1.0 - min/i
	} else {
		s = 0.0
	}

	return HSI{h, s, i, a}
}
