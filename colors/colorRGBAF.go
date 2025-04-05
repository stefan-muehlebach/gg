package colors

import (
	"image/color"
	"log"
)

// RGBAF entspricht dem NRGBA-Typ aus image/color, verwendet fuer die
// einzelnen Komponenten jedoch Fliesskommazahlen im Intervall [0,1].
// Beachte: Der Typ in diesem Package heisst FRGBA, die Werte R, G, B und A
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
	return c.Interpolate(RGBAF{1, 1, 1, 1}, t)
}

func (c RGBAF) Dark(t float64) Color {
	return c.Interpolate(RGBAF{0, 0, 0, 1}, t)
}

func (c RGBAF) Alpha(a float64) Color {
	a = setIn(a, 0, 1)
	return RGBAF{c.R, c.G, c.B, a}
}

func (c1 RGBAF) Interpolate(col Color, t float64) Color {
	t = ipf(setIn(t, 0, 1))
	c2 := RGBAFModel.Convert(col).(RGBAF)
	r := (1.0-t)*c1.R + t*c2.R
	g := (1.0-t)*c1.G + t*c2.G
	b := (1.0-t)*c1.B + t*c2.B
	a := (1.0-t)*c1.A + t*c2.A
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

// func (c *RGBAF) UnmarshalJSON(j []byte) error {
//     var col RGBAF

//     err := json.Unmarshal(j, &col)
//     if err != nil {
//         return err
//     }
//     *c = col
//     return nil
// }

// Modell fuer den neuen Farbtyp, d.h. fuer die Konvertierung von einer
// beliebigen Farbe nach RGBAF.
var (
	RGBAFModel color.Model = color.ModelFunc(rgbafModel)
)

func rgbafModel(c color.Color) color.Color {
	if _, ok := c.(RGBAF); ok {
		return c
	}
	r, g, b, a := c.RGBA()
	if a == 0xffff {
		return RGBAF{float64(r) / 65535.0, float64(g) / 65535.0, float64(b) / 65535.0, float64(a) / 65535.0}
	}
	if a == 0 {
		return RGBAF{}
	}
	r = (r * 0xffff) / a
	g = (g * 0xffff) / a
	b = (b * 0xffff) / a
	return RGBAF{float64(r) / 65535.0, float64(g) / 65535.0, float64(b) / 65535.0, float64(a) / 65535.0}
}
