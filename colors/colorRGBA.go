package colors

import (
	"fmt"
	"image/color"
	"log"
	"strconv"
)

// Dieser Typ wird fuer die Farbwerte verwendet, welche via SPI zu den LED's
// gesendet werden. Die Daten sind _nicht_ gamma-korrigiert (dies wird erst
// auf dem Panel-Empfaenger gemacht) und entsprechen dem Typ color.NRGBA
// von Go. LedColor implementiert das color.Color Interface.
type RGBA struct {
	R, G, B, A uint8
}

// RGBA ist Teil des color.Color Interfaces und retourniert die Farbwerte
// als Alpha-korrigierte uint16-Werte.
func (c RGBA) RGBA() (r, g, b, a uint32) {
	r, g, b, a = uint32(c.R), uint32(c.G), uint32(c.B), uint32(c.A)
	r |= r << 8
	r *= a
	r /= 0xFF
	g |= g << 8
	g *= a
	g /= 0xFF
	b |= b << 8
	b *= a
	b /= 0xFF
	a |= a << 8
	return
}

// Dient dem schnelleren Zugriff auf die drei Farbwerte.
func (c RGBA) RGB() (r, g, b uint8) {
	return c.R, c.G, c.B
}

// Retourniert eine neue Farbe, welche eine Interpolation zwischen c und Weiss
// ist. t ist ein Wert in [0, 1] und bestimmt die Position der Interpolation.
// t=0 retourniert c, t=1 retourniert Weiss.
func (c RGBA) Bright(t float64) RGBA {
	t = setIn(t, 0.0, 1.0)
	return c.Interpolate(white, t)
}

// Retourniert eine neue Farbe, welche eine Interpolation zwischen c und
// Schwarz ist. t ist ein Wert in [0, 1] und bestimmt die Position der
// Interpolation. t=0 retourniert c, t=1 retourniert Schwarz.
func (c RGBA) Dark(t float64) RGBA {
	t = setIn(t, 0.0, 1.0)
	return c.Interpolate(black, t)
}

// Retourniert eine neue Farbe, basierend auf c, jedoch mit dem hier
// angegebenen Alpha-Wert (als Fliesskommazahl in [0, 1], wobei 0 voll
// transparent und 1 voll deckend bedeuten).
func (c RGBA) Alpha(t float64) RGBA {
	t = setIn(t, 0.0, 1.0)
	return RGBA{c.R, c.G, c.B, uint8(255.0 * t)}
}

// Berechnet eine RGB-Farbe, welche 'zwischen' den Farben c1 und c2 liegt,
// so dass bei t=0 der Farbwert c1 und bei t=1 der Farbwert c2 retourniert
// wird. t wird vorgaengig auf das Interval [0,1] eingeschraenkt.
func (c1 RGBA) Interpolate(c2 RGBA, t float64) RGBA {
	r := (1.0-t)*float64(c1.R) + t*float64(c2.R)
	g := (1.0-t)*float64(c1.G) + t*float64(c2.G)
	b := (1.0-t)*float64(c1.B) + t*float64(c2.B)
	a := (1.0-t)*float64(c1.A) + t*float64(c2.A)
	return RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
}

func (c1 RGBA) Less(c2 RGBA, key SortField) bool {
	switch key {
	case ByRed:
		return c1.R < c2.R
	case ByGreen:
		return c1.G < c2.G
	case ByBlue:
		return c1.B < c2.B
	default:
		log.Fatalf("invalid sort field specified: '%v'", key)
		return false
	}
}

// Erzeugt eine druckbare Variante der Farbe. Im Wesentlichen werden die Werte
// fuer Rot, Gruen, Blau und Alpha als Hex-Zahlen ausgegeben.
func (c RGBA) String() string {
	return fmt.Sprintf("{0x%02X, 0x%02X, 0x%02X, 0x%02X}", c.R, c.G, c.B, c.A)
}

// Damit koennen Farbwerte im Format 0xRRGGBB eingelesen werden, wie sie bspw.
// in JSON-Files verwendet werden.
func (c *RGBA) UnmarshalText(text []byte) error {
	hexStr := string(text[2:])
	hexVal, err := strconv.ParseUint(hexStr, 16, 32)
	if err != nil {
		log.Fatal(err)
	}
	c.R = uint8((hexVal & 0xFF0000) >> 16)
	c.G = uint8((hexVal & 0x00FF00) >> 8)
	c.B = uint8((hexVal & 0x0000FF))
	c.A = 0xFF
	return nil
}

// Mischt die Farben c (Vordergrundfarbe) und bg (Hintergrundfarbe) nach einem
// Verfahren, welches in mix spezifiziert ist. Siehe auch ColorMixType.
func (c RGBA) Mix(bg RGBA, mix ColorMixType) RGBA {
	// bg := RGBAModel.Convert(c2).(RGBA)

	switch mix {
	case Replace:
		return c
	case Blend:
		ca := float64(c.A) / 255.0
		da := float64(bg.A) / 255.0
		a := 1.0 - (1.0-ca)*(1.0-da)
		t1 := ca / a
		t2 := da * (1.0 - ca) / a
		r := float64(c.R)*t1 + float64(bg.R)*t2
		g := float64(c.G)*t1 + float64(bg.G)*t2
		b := float64(c.B)*t1 + float64(bg.B)*t2
		return RGBA{uint8(r), uint8(g), uint8(b), uint8(255.0 * a)}
	case Max:
		r := max(c.R, bg.R)
		g := max(c.G, bg.G)
		b := max(c.B, bg.B)
		a := max(c.A, bg.A)
		return RGBA{r, g, b, a}
	case Average:
		r := c.R/2 + bg.R/2
		g := c.G/2 + bg.G/2
		b := c.B/2 + bg.B/2
		a := c.A/2 + bg.A/2
		return RGBA{r, g, b, a}
	case Min:
		r := min(c.R, bg.R)
		g := min(c.G, bg.G)
		b := min(c.B, bg.B)
		a := min(c.A, bg.A)
		return RGBA{r, g, b, a}
	default:
		log.Fatalf("Unknown mixing function: '%d'", mix)
	}
	return RGBA{}
}

// Modell fuer den neuen Farbtyp, d.h. fuer die Konvertierung von einer
// beliebigen Farbe nach RGBAF.
var (
	RGBAModel color.Model = color.ModelFunc(rgbaModel)
)

func rgbaModel(c color.Color) color.Color {
	if _, ok := c.(RGBA); ok {
		return c
	}
	r, g, b, a := c.RGBA()
	if a == 0xffff {
		return RGBA{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), 0xFF}
	}
	if a == 0 {
		return RGBA{}
	}
	r = (r * 0xffff) / a
	g = (g * 0xffff) / a
	b = (b * 0xffff) / a
	return RGBA{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), uint8(a >> 8)}
}
