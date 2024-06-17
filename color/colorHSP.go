package color

import (
	"image/color"
	"math"
)

const (
	perRed   = 0.241
	perGreen = 0.691
	perBlue  = 0.068
)

// Der neue Farbtyp HSP mit 'P' fuer "perceived brightness".
type HSP struct {
	H, S, P, A float64
}

func (c HSP) RGBA() (r, g, b, a uint32) {
	var part, minOverMax, h float64

	R, G, B := 0.0, 0.0, 0.0
	h = c.H / 360.0

	minOverMax = 1. - c.S

	if minOverMax > 0. {
		if h < 1./6. { //  R>G>B
			h = 6. * (h - 0./6.)
			part = 1. + h*(1./minOverMax-1.)
			B = c.P / math.Sqrt(perRed/minOverMax/minOverMax+perGreen*part*part+perBlue)
			R = (B) / minOverMax
			G = (B) + h*((R)-(B))
		} else if h < 2./6. { //  G>R>B
			h = 6. * (-h + 2./6.)
			part = 1. + h*(1./minOverMax-1.)
			B = c.P / math.Sqrt(perGreen/minOverMax/minOverMax+perRed*part*part+perBlue)
			G = (B) / minOverMax
			R = (B) + h*((G)-(B))
		} else if h < 3./6. { //  G>B>R
			h = 6. * (h - 2./6.)
			part = 1. + h*(1./minOverMax-1.)
			R = c.P / math.Sqrt(perGreen/minOverMax/minOverMax+perBlue*part*part+perRed)
			G = (R) / minOverMax
			B = (R) + h*((G)-(R))
		} else if h < 4./6. { //  B>G>R
			h = 6. * (-h + 4./6.)
			part = 1. + h*(1./minOverMax-1.)
			R = c.P / math.Sqrt(perBlue/minOverMax/minOverMax+perGreen*part*part+perRed)
			B = (R) / minOverMax
			G = (R) + h*((B)-(R))
		} else if h < 5./6. { //  B>R>G
			h = 6. * (h - 4./6.)
			part = 1. + h*(1./minOverMax-1.)
			G = c.P / math.Sqrt(perBlue/minOverMax/minOverMax+perRed*part*part+perGreen)
			B = (G) / minOverMax
			R = (G) + h*((B)-(G))
		} else { //  R>B>G
			h = 6. * (-h + 6./6.)
			part = 1. + h*(1./minOverMax-1.)
			G = c.P / math.Sqrt(perRed/minOverMax/minOverMax+perBlue*part*part+perGreen)
			R = (G) / minOverMax
			B = (G) + h*((R)-(G))
		}
	} else {
		if h < 1./6. { //  R>G>B
			h = 6. * (h - 0./6.)
			R = math.Sqrt(c.P * c.P / (perRed + perGreen*h*h))
			G = (R) * h
			B = 0.
		} else if h < 2./6. { //  G>R>B
			h = 6. * (-h + 2./6.)
			G = math.Sqrt(c.P * c.P / (perGreen + perRed*h*h))
			R = (G) * h
			B = 0.
		} else if h < 3./6. { //  G>B>R
			h = 6. * (h - 2./6.)
			G = math.Sqrt(c.P * c.P / (perGreen + perBlue*h*h))
			B = (G) * h
			R = 0.
		} else if h < 4./6. { //  B>G>R
			h = 6. * (-h + 4./6.)
			B = math.Sqrt(c.P * c.P / (perBlue + perGreen*h*h))
			G = (B) * h
			R = 0.
		} else if h < 5./6. { //  B>R>G
			h = 6. * (h - 4./6.)
			B = math.Sqrt(c.P * c.P / (perBlue + perRed*h*h))
			R = (B) * h
			G = 0.
		} else { //  R>B>G
			h = 6. * (-h + 6./6.)
			R = math.Sqrt(c.P * c.P / (perRed + perBlue*h*h))
			B = (R) * h
			G = 0.
		}
	}
	R, G, B = min(R, 1.0), min(G, 1.0), min(B, 1.0)
	r = uint32(65535.0 * R * c.A)
	g = uint32(65535.0 * G * c.A)
	b = uint32(65535.0 * B * c.A)
	a = uint32(65535.0 * c.A)
	return
}

func (c HSP) Bright(t float64) Color {
	t = setIn(t, 0, 1)
	r := c
	r.S = (1-t)* c.S
	r.P = (1-t)*c.P + t
	return r
}

func (c HSP) Dark(t float64) Color {
	t = setIn(t, 0, 1)
	r := c
	r.P = (1 - t) * c.P
	return r
}

func (c HSP) Alpha(a float64) Color {
	a = setIn(a, 0, 1)
	return HSP{c.H, c.S, c.P, a}
}

func (c1 HSP) Interpolate(col Color, t float64) Color {
	return c1
}

// Modell fuer den neuen Farbtyp, d.h. fuer die Konvertierung von einer
// beliebigen Farbe nach HSP.
var (
	HSPModel color.Model = color.ModelFunc(hspModel)
)

func hspModel(c color.Color) color.Color {
	if _, ok := c.(HSP); ok {
		return c
	}
	red, green, blue, alpha := c.RGBA()
	if alpha == 0 {
		return HSP{0.0, 0.0, 0.0, 0.0}
	}
	r := float64((red*0xffff)/alpha) / 65535.0
	g := float64((green*0xffff)/alpha) / 65535.0
	b := float64((blue*0xffff)/alpha) / 65535.0
	a := float64(alpha) / 65535.0

	h, s, p := 0.0, 0.0, 0.0

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

	p = math.Sqrt(r*r*perRed + g*g*perGreen + b*b*perBlue)

	return HSP{h, s, p, a}
}
