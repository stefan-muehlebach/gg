package colors

import (
	"image/color"
	"log"
	"math"
)

// Der Typ HSL schliesslich enthaelt die Werte fuer Hue, Saturation und
// Lightness gespeichert.
type HSL struct {
	H, S, L, A float64
}

func (c HSL) RGBA() (r, g, b, a uint32) {
	C := c.S * (1.0 - math.Abs(2.0*c.L-1.0))
	X := C * (1.0 - math.Abs(math.Mod(c.H/60.0, 2.0)-1.0))
	m := c.L - C/2.0
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

func (c HSL) Bright(t float64) Color {
    white := HSL{c.H, 0.0, 1.0, 1.0}
    return c.Interpolate(white, t)
	// t = setIn(t, 0, 1)
	// r := c
	// r.L = (1-t)*c.L + t
	// return r
}

func (c HSL) Dark(t float64) Color {
    black := HSL{c.H, 0.0, 0.0, 1.0}
    return c.Interpolate(black, t)
	// t = setIn(t, 0, 1)
	// r := c
	// r.L = (1 - t) * c.L
	// return r
}

func (c HSL) Alpha(a float64) Color {
	a = setIn(a, 0, 1)
	return HSL{c.H, c.S, c.L, a}
}

func (c1 HSL) Interpolate(col Color, t float64) Color {
	t = ipf(setIn(t, 0, 1))
	c2 := HSLModel.Convert(col).(HSL)
	h := (1.0-t)*c1.H + t*c2.H
	s := (1.0-t)*c1.S + t*c2.S
	l := (1.0-t)*c1.L + t*c2.L
	a := (1.0-t)*c1.A + t*c2.A
	return HSL{h, s, l, a}
}

func (c1 HSL) Less(c2 HSL, key SortField) bool {
	switch key {
	case SortByHue:
		return c1.H < c2.H
	case SortBySaturation:
		return c1.S < c2.S
	case SortByLightness:
		return c1.L < c2.L
	default:
		log.Fatalf("invalid sort field specified: '%v'", key)
		return false
	}
}

// Modell fuer den neuen Farbtyp, d.h. fuer die Konvertierung von einer
// beliebigen Farbe nach HSL.
var (
	HSLModel color.Model = color.ModelFunc(hslModel)
)

func hslModel(c color.Color) color.Color {
	if _, ok := c.(HSL); ok {
		return c
	}
	red, green, blue, alpha := c.RGBA()
	if alpha == 0 {
		return HSL{0.0, 0.0, 0.0, 0.0}
	}
	r := float64((red*0xffff)/alpha) / 65535.0
	g := float64((green*0xffff)/alpha) / 65535.0
	b := float64((blue*0xffff)/alpha) / 65535.0
	a := float64(alpha) / 65535.0

	h, s, l := 0.0, 0.0, 0.0

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

	l = (max + min) / 2.0

	if d == 0.0 {
		s = 0.0
	} else {
		s = d / (1.0 - math.Abs(2.0*l-1.0))
	}

	return HSL{h, s, l, a}
}
