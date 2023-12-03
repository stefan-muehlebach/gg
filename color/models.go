package color

import (
	"image/color"
    "math"
)

// Modelle für die neuen Farbtypen, d.h. für die Konvertierung eines beliebigen
// color.Color-Wertes in den jewiligen Farbtyp.
var (
    //RGBAF64Model  color.Model = color.ModelFunc(rgbaf64Model)
    RGBAFModel    color.Model = color.ModelFunc(rgbafModel)
    HSVModel      color.Model = color.ModelFunc(hsvModel)
    HSLModel      color.Model = color.ModelFunc(hslModel)
    HSIModel      color.Model = color.ModelFunc(hsiModel)
)

/*
func rgbaf64Model(c color.Color) color.Color {
    if _, ok := c.(RGBAF64); ok {
        return c
    }
    r, g, b, a := c.RGBA()
    return RGBAF64{float64(r)/65535.0, float64(g)/65535.0, float64(b)/65535.0, float64(a)/65535.0}
}
*/

func rgbafModel(c color.Color) color.Color {
    if _, ok := c.(RGBAF); ok {
        return c
    }
    r, g, b, a := c.RGBA()
    if a == 0xffff {
        return RGBAF{float64(r)/65535.0, float64(g)/65535.0, float64(b)/65535.0, float64(a)/65535.0}
    }
    if a == 0 {
        return RGBAF{0.0, 0.0, 0.0, 0.0}
    }
    r = (r * 0xffff) / a
    g = (g * 0xffff) / a
    b = (b * 0xffff) / a
    return RGBAF{float64(r)/65535.0, float64(g)/65535.0, float64(b)/65535.0, float64(a)/65535.0}
}

func hsvModel(c color.Color) color.Color {
    if _, ok := c.(HSV); ok {
        return c
    }
    red, green, blue, alpha := c.RGBA()
    if alpha == 0 {
        return HSV{0.0, 0.0, 0.0, 0.0}
    }
    r := float64((red   * 0xffff) / alpha) / 65535.0
    g := float64((green * 0xffff) / alpha) / 65535.0
    b := float64((blue  * 0xffff) / alpha) / 65535.0
    a := float64(alpha) / 65535.0
    
    h, s, v := 0.0, 0.0, 0.0
    
    max := math.Max(r, math.Max(g, b))
    min := math.Min(r, math.Min(g, b))
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
        h = 2.0 + (b - r) / d
    case b:
        h = 4.0 + (r - g) / d
    }
    h *= 60.0

    if max == 0.0 {
        s = 0.0
    } else {
        s = d/max
    }
    
    v = max
    
    return HSV{h, s, v, a}
}

func hslModel(c color.Color) color.Color {
    if _, ok := c.(HSL); ok {
        return c
    }
    red, green, blue, alpha := c.RGBA()
    if alpha == 0 {
        return HSL{0.0, 0.0, 0.0, 0.0}
    }
    r   := float64((red * 0xffff) / alpha) / 65535.0
    g   := float64((green * 0xffff) / alpha) / 65535.0
    b   := float64((blue * 0xffff) / alpha) / 65535.0
    a   := float64(alpha) / 65535.0
    
    h, s, l := 0.0, 0.0, 0.0
    
    max := math.Max(r, math.Max(g, b))
    min := math.Min(r, math.Min(g, b))
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
        h = 2.0 + (b - r) / d
    case b:
        h = 4.0 + (r - g) / d
    }
    h *= 60.0
    
    l = (max + min) / 2.0

    if d == 0.0 {
        s = 0.0
    } else {
        s = d / (1.0 - math.Abs(2.0 * l - 1.0))
    }

    return HSL{h, s, l, a}
}

func hsiModel(c color.Color) color.Color {
    if _, ok := c.(HSI); ok {
        return c
    }
    red, green, blue, alpha := c.RGBA()
    if alpha == 0 {
        return HSI{0.0, 0.0, 0.0, 0.0}
    }
    r   := float64((red   * 0xffff) / alpha) / 65535.0
    g   := float64((green * 0xffff) / alpha) / 65535.0
    b   := float64((blue  * 0xffff) / alpha) / 65535.0
    a   := float64(alpha) / 65535.0

    h, s, i := 0.0, 0.0, 0.0
    
    i = (r+g+b)/3.0
    m := math.Min(r, math.Min(g, b))
    if i > 0.0 {
        s = 1.0 - m/i
    } else {
        s = 0.0
    }
    h = (180.0 / math.Pi) * math.Acos((r - 0.5*g - 0.5*b) / math.Sqrt(r*r + g*g + b*b - r*g - r*b - g*b))
    if b > g {
        h = 360.0 - h
    }

    return HSI{h, s, i, a}
}
