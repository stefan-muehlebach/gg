package color

import (
    "log"
    "math"
)

type Color interface {
    RGBA() (r, g, b, a uint32)
    Bright(t float64) (Color)
    Dark(t float64) (Color)
    // Bright(l int) (Color)
    // Dark(l int) (Color)
    // Transp() (Color)
    // Opaque() (Color)
    Alpha(a float64) (Color)
    Interpolate(c2 Color, t float64) (Color)
}

// RGBAF64 entspricht dem RGBA-Typ aus image/color, verwendet fuer die
// einzelnen Komponenten jedoch Fliesskommazahlen im Intervall [0,1].
type RGBAF64 struct {
    R, G, B, A float64
}

func (c RGBAF64) RGBA() (r, g, b, a uint32) {
    r = uint32(65535.0 * c.R)
    g = uint32(65535.0 * c.G)
    b = uint32(65535.0 * c.B)
    a = uint32(65535.0 * c.A)
    return
}

func (c RGBAF64) Bright(t float64) (Color) {
    return c
}

func (c RGBAF64) Dark(t float64) (Color) {
    return c
}

// func (c RGBAF64) Transp() (Color) {
//     r := c
//     r.A = 0.6
//     r.R *= r.A
//     r.G *= r.A
//     r.B *= r.A
//     return r
// }

// func (c RGBAF64) Opaque() (Color) {
//     r := c
//     r.R /= r.A
//     r.G /= r.A
//     r.B /= r.A
//     r.A = 1.0
//     return r
// }

func (c RGBAF64) Alpha(a float64) (Color) {
    r := c
    r.R *= a/r.A
    r.G *= a/r.A
    r.B *= a/r.A
    r.A  = a
    return r
}

func (c1 RGBAF64) Interpolate(col Color, t float64) (Color) {
    c2 := col.(RGBAF64)

    r := (1-t)*c1.R + t*c2.R
    g := (1-t)*c1.G + t*c2.G
    b := (1-t)*c1.B + t*c2.B
    a := (1-t)*c1.A + t*c2.A
    return RGBAF64{r, g, b, a}
}

func (c1 RGBAF64) Less(c2 RGBAF64, key SortField) bool {
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

// NRGBAF64 entspricht dem NRGBA-Typ aus image/color, verwendet fuer die
// einzelnen Komponenten jedoch Fliesskommazahlen im Intervall [0,1].
type NRGBAF64 struct {
    R, G, B, A float64
}

func (c NRGBAF64) RGBA() (r, g, b, a uint32) {
    r = uint32(65535.0 * c.R * c.A)
    g = uint32(65535.0 * c.G * c.A)
    b = uint32(65535.0 * c.B * c.A)
    a = uint32(65535.0 * c.A)
    return
}

func (c NRGBAF64) Bright(t float64) (Color) {
    return c
}

func (c NRGBAF64) Dark(t float64) (Color) {
    return c
}

// func (c NRGBAF64) Transp() (Color) {
//     r := c
//     r.A = 0.6
//     return r
// }

// func (c NRGBAF64) Opaque() (Color) {
//     r := c
//     r.A = 1.0
//     return r
// }

func (c NRGBAF64) Alpha(a float64) (Color) {
    r := c
    r.A = a
    return r
}

func (c1 NRGBAF64) Interpolate(col Color, t float64) (Color) {
    c2 := col.(NRGBAF64)

    r := (1-t)*c1.R + t*c2.R
    g := (1-t)*c1.G + t*c2.G
    b := (1-t)*c1.B + t*c2.B
    a := (1-t)*c1.A + t*c2.A
    return NRGBAF64{r, g, b, a}
}

func (c1 NRGBAF64) Less(c2 NRGBAF64, key SortField) bool {
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
    X := C * (1.0 - math.Abs(math.Mod(c.H/60.0, 2.0) - 1.0))
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
    r = uint32(65535.0 * (R+m) * c.A)
    g = uint32(65535.0 * (G+m) * c.A)
    b = uint32(65535.0 * (B+m) * c.A)
    a = uint32(65535.0 * c.A)
    return    
}

func (c HSV) Bright(t float64) (Color) {
    // t := 0.2 * float64(level)
    r := c
    r.S = (1-t)*c.S
    r.V = (1-t)*c.V + t*1.0
    return r
}

func (c HSV) Dark(t float64) (Color) {
    // t := 0.2 * float64(level)
    r := c
    r.S = (1-t)*c.S + t*1.0
    r.V = (1-t)*c.V
    return r
}

// func (c HSV) Transp() (Color) {
//     r := c
//     r.A = 0.6
//     return r
// }

// func (c HSV) Opaque() (Color) {
//     r := c
//     r.A = 1.0
//     return r
// }

func (c HSV) Alpha(a float64) (Color) {
    r := c
    r.A = a
    return r
}

func (c1 HSV) Interpolate(col Color, t float64) (Color) {
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
    C := c.S * (1.0 - math.Abs(2.0*c.L - 1.0))
    X := C * (1.0 - math.Abs(math.Mod(c.H/60.0, 2.0) - 1.0))
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
    r = uint32(65535.0 * (R+m) * c.A)
    g = uint32(65535.0 * (G+m) * c.A)
    b = uint32(65535.0 * (B+m) * c.A)
    a = uint32(65535.0 * c.A)
    return    
}

func (c HSL) Bright(t float64) (Color) {
    // t := 0.2 * float64(level)
    r := c
    r.L = (1-t)*c.L + t*1.0
    return r
}

func (c HSL) Dark(t float64) (Color) {
    // t := 0.2 * float64(level)
    r := c
    r.L = (1-t)*c.L
    return r
}

// func (c HSL) Transp() (Color) {
//     r := c
//     r.A = 0.6
//     return r
// }

// func (c HSL) Opaque() (Color) {
//     r := c
//     r.A = 1.0
//     return r
// }

func (c HSL) Alpha(a float64) (Color) {
    r := c
    r.A = a
    return r
}

func (c1 HSL) Interpolate(col Color, t float64) (Color) {
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
    red, green, blue := 0.0, 0.0, 0.0

    switch {
    case c.H == 0.0:
        red   = c.I + 2*c.I*c.S
        green = c.I - c.I*c.S
        blue  = c.I - c.I*c.S
    case c.H < 120.0:
        red   = c.I + c.I*c.S*math.Cos(c.H*math.Pi/180.0)/math.Cos((60-c.H)*math.Pi/180.0)
        green = c.I + c.I*c.S*(1.0 - math.Cos(c.H*math.Pi/180.0)/math.Cos((60-c.H)*math.Pi/180.0))
        blue  = c.I - c.I*c.S
    case c.H == 120.0:
        red   = c.I - c.I*c.S
        green = c.I + 2*c.I*c.S
        blue  = c.I - c.I*c.S
    case c.H < 240.0:
        red   = c.I - c.I*c.S
        green = c.I + c.I*c.S*math.Cos((c.H-120.0)*math.Pi/180.0)/math.Cos((180.0-c.H)*math.Pi/180.0)
        blue  = c.I + c.I*c.S*(1.0 - math.Cos((c.H-120.0)*math.Pi/180.0)/math.Cos((180.0-c.H)*math.Pi/180.0))
    case c.H == 240.0:
        red   = c.I - c.I*c.S
        green = c.I - c.I*c.S
        blue  = c.I + 2*c.I*c.S
    case c.H < 360.0:
        red   = c.I + c.I*c.S*(1.0 - math.Cos((c.H-240.0)*math.Pi/180.0)/math.Cos((300.0-c.H)*math.Pi/180.0))
        green = c.I - c.I*c.S
        blue  = c.I + c.I*c.S*math.Cos((c.H-240.0)*math.Pi/180.0)/math.Cos((300.0-c.H)*math.Pi/180.0)
    }    

    r = uint32(65535.0 * red * c.A)
    g = uint32(65535.0 * green * c.A)
    b = uint32(65535.0 * blue * c.A)
    a = uint32(65535.0 * c.A)
    return    
}

func (c HSI) Bright(t float64) (Color) {
    return c
}

func (c HSI) Dark(t float64) (Color) {
    return c
}

// func (c HSI) Transp() (Color) {
//     return c
// }
    
// func (c HSI) Opaque() (Color) {
//     return c
// }

func (c HSI) Alpha(a float64) (Color) {
    r := c
    r.A = a
    return r
}

func (c1 HSI) Interpolate(col Color, t float64) (Color) {
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
