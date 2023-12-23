// Erweiterung des Packages 'image/color' um neue Farbtypen.
//
// Die bestehende Implementation von Farben in 'image/color' bietet keine
// Methoden, um Farben heller, resp. dunkler zu schattieren oder um zwischen
// zwei beliebigen Farben eine lineare Interpolation durchzuführen.
// Die in diesem Package definierten Farben implementieren alle das
// Interface 'Color' aus 'image/color'. Der Name des Packages und die
package color

import (
	"log"
	"math"
    "image/color"
)

// Da dieses Package als vollwertiger Ersatz für 'image/color' gedacht ist,
// sind einige Standardfarben auch hier zu finden.
var (
	Black       = RGBAF{0.0, 0.0, 0.0, 1.0}
	White       = RGBAF{1.0, 1.0, 1.0, 1.0}
	Transparent = RGBAF{0.0, 0.0, 0.0, 0.0}
	Opaque      = RGBAF{1.0, 1.0, 1.0, 1.0}
)

// Das Interface Color basiert auf dem gleichnamigen Interface der
// Standard-Bibliothek, verlangt jedoch Methoden, um eine Farbe aufzuhellen,
// resp. abzudunkeln, den Alpha-Wert als Fliesskommazahl in [0,1] direkt
// anzugeben und um zwischen zwei Farben eine lineare Interpolation
// durchzuführen.
type Color interface {
    color.Color
	Bright(t float64) Color
	Dark(t float64) Color
	Alpha(a float64) Color
	Interpolate(c2 Color, t float64) Color
}

func setIn(v, a, b float64) (float64) {
    if v < a {
        return a
    }
    if v > b {
        return b
    }
    return v
}
    
// RGBAF entspricht dem NRGBA-Typ aus image/color, verwendet fuer die
// einzelnen Komponenten jedoch Fliesskommazahlen im Intervall [0,1].
// Beachte: Der Typ in diesem Package heisst RGBA, die Werte R, G, B und A
// werden jedoch als _nicht_ normierte Werte interpretiert, meine praktischen
// Erfahrungen haben gezeigt, dass dies intuitiver ist.
type RGBAF struct {
	R, G, B, A float64
}

func (c RGBAF) RGBA() (r, g, b, a uint32) {
	r = uint32(65535.0 * c.R * c.A)
	g = uint32(65535.0 * c.G * c.A)
	b = uint32(65535.0 * c.B * c.A)
	a = uint32(65535.0 * c.A)
	return
}

func (c RGBAF) Bright(t float64) Color {
    t = setIn(t, 0, 1)
	u := 1.0 - t
	return RGBAF{u*c.R + t, u*c.G + t, u*c.B + t, c.A}
}

func (c RGBAF) Dark(t float64) Color {
    t = setIn(t, 0, 1)
	u := 1.0 - t
	return RGBAF{u * c.R, u * c.G, u * c.B, c.A}
}

func (c RGBAF) Alpha(a float64) Color {
    a = setIn(a, 0, 1)
	return RGBAF{c.R, c.G, c.B, a}
}

func (c1 RGBAF) Interpolate(col Color, t float64) Color {
    t = setIn(t, 0, 1)
	c2 := col.(RGBAF)

	r := (1-t)*c1.R + t*c2.R
	g := (1-t)*c1.G + t*c2.G
	b := (1-t)*c1.B + t*c2.B
	a := (1-t)*c1.A + t*c2.A
	return RGBAF{r, g, b, a}
}

func (c1 RGBAF) Less(c2 RGBAF, key SortField) bool {
	switch key {
	case SortByRed:
		return c1.R < c2.R
	case SortByGreen:
		return c1.G < c2.G
	case SortByBlue:
		return c1.B < c2.B
	default:
		log.Fatalf("invalid sort field specified: '%v'", key)
		return false
	}
}

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
    t = setIn(t, 0, 1)
    u := 1-t
	r := c
	r.S = u*c.S
	r.V = u*c.V + t
	return r
}

func (c HSV) Dark(t float64) Color {
    t = setIn(t, 0, 1)
    u := 1-t
	r := c
	r.V = u*c.V
	return r
}

func (c HSV) Alpha(a float64) Color {
    a = setIn(a, 0, 1)
	return HSV{c.H, c.S, c.V, a}
}

func (c1 HSV) Interpolate(col Color, t float64) Color {
    t = setIn(t, 0, 1)
	c2 := col.(HSV)

	h := (1-t)*c1.H + t*c2.H
	s := (1-t)*c1.S + t*c2.S
	v := (1-t)*c1.V + t*c2.V
	a := (1-t)*c1.A + t*c2.A
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
    t = setIn(t, 0, 1)
	r := c
	r.L = (1-t)*c.L + t*1.0
	return r
}

func (c HSL) Dark(t float64) Color {
    t = setIn(t, 0, 1)
	r := c
	r.L = (1 - t) * c.L
	return r
}

func (c HSL) Alpha(a float64) Color {
    a = setIn(a, 0, 1)
	return HSL{c.H, c.S, c.L, a}
}

func (c1 HSL) Interpolate(col Color, t float64) Color {
    t = setIn(t, 0, 1)
	c2 := col.(HSL)
	h := (1-t)*c1.H + t*c2.H
	s := (1-t)*c1.S + t*c2.S
	l := (1-t)*c1.L + t*c2.L
	a := (1-t)*c1.A + t*c2.A
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
    if R > 1.0 { R = 1.0 }
    if G > 1.0 { G = 1.0 }
    if B > 1.0 { B = 1.0 }
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
