//go:generate rm -f colornames.go
//go:generate go run gen.go

// Erweiterung des Packages 'image/color' um neue Farbtypen.
//
// Dieses Package versteht sich als Erweiterung von 'image/color'
// in Zusammenhang mit dem Package 'gg'.
// Die bestehende Implementation von Farben in 'image/color' bietet keine
// Methoden, um Farben heller, resp. dunkler zu schattieren oder um zwischen
// zwei beliebigen Farben eine lineare Interpolation durchzuführen.
// Die in diesem Package definierten Farben implementieren alle das
// Interface 'Color' aus 'image/color'. Der Name des Packages und die
package colors

import (
	"image/color"
	"math"
	"math/rand/v2"
)

// Da dieses Package anstelle von 'image/color' verwendet werden kann,
// sind einige Standardfarben auch hier definiert.
var (
	Transparent = RGBA{0x00, 0x00, 0x00, 0x00}
	Opaque      = RGBA{0xFF, 0xFF, 0xFF, 0xFF}

	Map   map[string]RGBA
	Names []string

	black = RGBA{0x00, 0x00, 0x00, 0xFF}
	white = RGBA{0xFF, 0xFF, 0xFF, 0xFF}
)

// Mit folgenden Konstanten kann das Verfahren bestimmt werden, welches beim
// Mischen von Farben verwendet werden soll (siehe auch Methode Mix).
type ColorMixType int

const (
	// Ersetzt die Hintergundfarbe durch die Vordergrundfarbe.
	Replace ColorMixType = iota
	// Ueberblendet die Hintergrundfarbe mit der Vordergrundfarbe anhand
	// des Alpha-Wertes.
	Blend
	// Bestimmt die neue Farbe durch das Maximum von RGB zwischen Hinter- und
	// Vordergrundfarbe.
	Max
	// Analog zu Max, nimmt jedoch den Mittelwert von jeweils R, G und B.
	Average
	// Analog zu Max, nimmt jedoch das Minimum von jeweils R, G und B.
	Min
)

// Mischt die Farben c (Vordergrundfarbe) und bg (Hintergrundfarbe) nach einem
// Verfahren, welches in mix spezifiziert ist. Siehe auch ColorMixType.

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
	Mix(bg Color, mix ColorMixType) Color
}

// Mit RandColor kann zufällig eine aus dem gesamten Sortiment der hier
// definierten Farben gewählt werden. Hilfreich für Tests, Beispielprogramme
// oder anderes.
func RandColor() RGBA {
	name := Names[rand.IntN(len(Names))]
	return Map[name]
}

// Mit RandGroupColor kann der Zufall eine bestimmte Farbgruppe beschraenkt
// werden.
func RandColorByGroup(group ColorGroup) RGBA {
	nameList, ok := Groups[group]
	if !ok {
		return RGBA{A: 0xff}
	}
	name := nameList[rand.IntN(len(nameList))]
	return Map[name]
}

// ---------------------------------------------------------------------------

var (
	ipf = cubicInterp
    gamma = 1.2
)

//go:inline
func linearInterp(t float64) float64 {
	return t
}

func powerInterp(t float64) float64 {
	t1 := math.Pow(2, gamma-1.0)
	if t < 0.5 {
		return t1 * math.Pow(t, gamma)
	} else {
		return 1.0 - t1*math.Pow(1.0-t, gamma)
	}
}

func cubicInterp(t float64) float64 {
	return 3.0*t*t - 2.0*t*t*t
}

// Hilfsfunktion, mit welcher sichergestellt werden kann, dass der Wert v
// zwingend zwischen a und b zu liegen kommt. Falls a groesser ist als b,
// dann wird v unveraendert zurueckgegeben.
func setIn(v, a, b float64) float64 {
	if a > b {
		return v
	}
	return min(max(v, a), b)
}
