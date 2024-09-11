//go:generate rm -f colornames.go
//go:generate go run gen.go

// Erweiterung des Packages 'image/color' um neue Farbtypen.
//
// Dieses Package versteht sich als Erweiterung von 'image/color'
// in Zusammenhang mit dem Package 'gg'.
//
// Im Wesentlichen
// Die bestehende Implementation von Farben in 'image/color' bietet keine
// Methoden, um Farben heller, resp. dunkler zu schattieren oder um zwischen
// zwei beliebigen Farben eine lineare Interpolation durchzuführen.
// Die in diesem Package definierten Farben implementieren alle das
// Interface 'Color' aus 'image/color'. Der Name des Packages und die
package color

import (
	"image/color"
	"math"
	"math/rand"
)

// Da dieses Package anstelle von 'image/color' verwendet werden kann,
// sind einige Standardfarben auch hier definiert.
var (
	Transparent = RGBAF{0.0, 0.0, 0.0, 0.0}
	Opaque      = RGBAF{1.0, 1.0, 1.0, 1.0}
	//Map         map[string]RGBAF
	//Names       []string
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

// Mit RandColor kann zufällig eine aus dem gesamten Sortiment der hier
// definierten Farben gewählt werden. Hilfreich für Tests, Beispielprogramme
// oder anderes.
func RandColor() Color {
	name := Names[rand.Int()%len(Names)]
	return Map[name]
}

// Mit RandGroupColor kann der Zufall eine bestimmte Farbgruppe beschraenkt
// werden.
func RandGroupColor(group ColorGroup) Color {
	nameList, ok := Groups[group]
	if !ok {
		return RGBAF{0, 0, 0, 1}
	}
	name := nameList[rand.Int()%len(nameList)]
	return Map[name]
}

type interpolFuncType func(float64) float64

var (
	ipf   = LinearInterpol
	gamma = 1.5
)

func SetInterpolFunc(fnc interpolFuncType) {
	ipf = fnc
}

func LinearInterpol(t float64) float64 {
	return t
}

func GammaInterpol(t float64) float64 {
	return math.Pow(t, gamma)
}

func CubicInterpol(t float64) float64 {
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
