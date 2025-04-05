package colors

import (
	"image/color"
	"log"
	"math"
)

// Beim Typ HSV werden die Werte fuer Hue, Saturation und Value gespeichert.
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

func (c HSV) Bright(t float64) Color {
    white := HSV{c.H, 0.0, 1.0, 1.0}
    return c.Interpolate(white, t)
	// t = setIn(t, 0, 1)
	// r := c
	// r.S = (1 - t) * c.S
	// r.V = (1-t)*c.V + t
	// return r
}

func (c HSV) Dark(t float64) Color {
    black := HSV{c.H, 0.0, 0.0, 1.0}
    return c.Interpolate(black, t)
	// t = setIn(t, 0, 1)
	// r := c
	// r.V = (1 - t) * c.V
	// return r
}

func (c HSV) Alpha(a float64) Color {
	a = setIn(a, 0, 1)
	return HSV{c.H, c.S, c.V, a}
}

func (c1 HSV) Interpolate(col Color, t float64) Color {
	t = ipf(setIn(t, 0, 1))
	c2 := HSVModel.Convert(col).(HSV)
	h := (1.0-t)*c1.H + t*c2.H
	s := (1.0-t)*c1.S + t*c2.S
	v := (1.0-t)*c1.V + t*c2.V
	a := (1.0-t)*c1.A + t*c2.A
	return HSV{h, s, v, a}
}

func (c1 HSV) Less(c2 HSV, key SortField) bool {
	switch key {
	case SortByHue:
		return c1.H < c2.H
	case SortBySaturation:
		return c1.S < c2.S
	case SortByValue:
		return c1.V < c2.V
	default:
		log.Fatalf("invalid sort field specified: '%v'", key)
		return false
	}
}

// Modell fuer den neuen Farbtyp, d.h. fuer die Konvertierung von einer
// beliebigen Farbe nach HSV.
var (
	HSVModel color.Model = color.ModelFunc(hsvModel)
)

func hsvModel(c color.Color) color.Color {
	if _, ok := c.(HSV); ok {
		return c
	}
	red, green, blue, alpha := c.RGBA()
	if alpha == 0 {
		return HSV{0.0, 0.0, 0.0, 0.0}
	}
	r := float64((red*0xffff)/alpha) / 65535.0
	g := float64((green*0xffff)/alpha) / 65535.0
	b := float64((blue*0xffff)/alpha) / 65535.0
	a := float64(alpha) / 65535.0

	h, s, v := 0.0, 0.0, 0.0

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

	if max == 0.0 {
		s = 0.0
	} else {
		s = d / max
	}

	v = max

	return HSV{h, s, v, a}
}
