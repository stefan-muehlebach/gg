package colors

import (
	"encoding/json"
	"log"
	"math"
	"os"
	"slices"
)

type Palette interface {
	Name() string
    Type() string
	Color(t float64) Color
}

var (
    ipf = LinearInterp
)

func LinearInterp(t float64) float64 {
    return t
}

func PowerInterp(t float64) float64 {
	a := 1.4
	t1 := math.Pow(2, a-1.0)
	if t <= 0.5 {
		return t1 * math.Pow(t, a)
	} else {
		return 1.0 - t1*math.Pow(1.0-t, a)
	}
}

func CubicInterp(t float64) float64 {
    return 3.0*t*t - 2.0*t*t*t
}

// Dieser (interne) Typ wird verwendet, um einen bestimmten Wert im Interval
// [0,1] mit einer Farbe zu assoziieren.
type ColorIdent byte

const (
	RedColor = (1 << iota)
	GreenColor
	BlueColor
)

func (c ColorIdent) String() (s string) {
	if (c & RedColor) != 0 {
		s += "R"
	}
	if (c & GreenColor) != 0 {
		s += "G"
	}
	if (c & BlueColor) != 0 {
		s += "B"
	}
	return s
}

func (c *ColorIdent) UnmarshalText(text []byte) error {
	for _, ch := range text {
		switch ch {
		case 'R':
			*c |= RedColor
		case 'G':
			*c |= GreenColor
		case 'B':
			*c |= BlueColor
		default:
			log.Fatalf("no such color identifier: %c", ch)
		}
	}
	return nil
}

type ColorStop struct {
	Pos    float64
	Color  RGBA
	Ignore ColorIdent
}

// ---------------------------------------------------------------------------

// Gradienten-Paletten basieren auf einer Anzahl Farben (Stuetzstellen)
// zwischen denen eine Farbe interpoliert werden kann. Jede Stuetzstelle
// besteht aus einer Position (Zahl im Intervall [0,1]) und einer dazu
// gehoerenden Farbe.
type GradPalette struct {
	name  string
	stops []ColorStop
}

func newGradPalette(name string) *GradPalette {
	p := &GradPalette{}
	p.name = name
	p.stops = []ColorStop{
		{Pos: 0.0, Color: Black},
		{Pos: 1.0, Color: Black},
	}
	return p
}

func NewPaletteByStops(name string, stops ...ColorStop) *GradPalette {
	p := newGradPalette(name)
	for _, stop := range stops {
		p.SetColorStop(stop)
	}
	return p
}

func NewPaletteByColors(name string, colors ...RGBA) *GradPalette {
	p := newGradPalette(name)
	posStep := 1.0 / (float64(len(colors) - 1))
	for i, color := range colors {
		stop := ColorStop{Pos: float64(i) * posStep, Color: RGBAModel.Convert(color).(RGBA)}
		p.SetColorStop(stop)
	}
	return p
}

func (p *GradPalette) SetColorStop(colStop ColorStop) {
	if colStop.Pos < 0.0 || colStop.Pos > 1.0 {
		log.Fatalf("Position must be in [0,1]; is: %f", colStop.Pos)
	}
	for i, stop := range p.stops {
		if stop.Pos == colStop.Pos {
			p.stops[i].Color = colStop.Color
			return
		}
		if stop.Pos > colStop.Pos {
			p.stops = slices.Insert(p.stops, i, colStop)
			return
		}
	}
}

func (p *GradPalette) Name() string {
	return p.name
}

func (p *GradPalette) Type() string {
	return "Gradient Palette"
}

// Hier nun spielt die Musik: aufgrund des Wertes t (muss im Intervall [0,1]
// liegen) wird eine neue Farbe interpoliert.
func (p *GradPalette) Color(t float64) (c Color) {
	var i int
	var stop ColorStop

	if t < 0.0 || t > 1.0 {
		t = max(0.0, min(1.0, t))
	}
	for i, stop = range p.stops[1:] {
		if stop.Pos > t {
			break
		}
	}
	t = (t - p.stops[i].Pos) / (p.stops[i+1].Pos - p.stops[i].Pos)
	c = p.stops[i].Color.Interpolate(p.stops[i+1].Color, ipf(t))
	return c
}

// ----------------------------------------------------------------------------

type ProcPalette struct {
	name   string
	params []ParamSet
}

func NewPaletteByParams(name string, params ...ParamSet) *ProcPalette {
    p := &ProcPalette{}
    p.name = name
    p.params = make([]ParamSet, 3)
    for i, param := range params {
        p.params[i] = param
    }
    return p
}

func (p *ProcPalette) Name() string {
	return p.name
}

func (p *ProcPalette) Type() string {
	return "Procedural Palette"
}

func (p *ProcPalette) Color(t float64) (c Color) {
    r := p.params[0].Value(t)
    g := p.params[1].Value(t)
    b := p.params[2].Value(t)
	return RGBA{uint8(255.0 * r), uint8(255.0 * g), uint8(255.0 * b), 0xFF}
}

type ParamSet struct {
	Y0, Y1, X0, X1 float64
}

func (p ParamSet) Value(t float64) float64 {
    return p.Y0 + p.Y1*math.Cos(2*math.Pi*(t*p.X0+p.X1))
}

// ---------------------------------------------------------------------------

type JsonPalette struct {
	Name   string
	Colors []RGBA
	Stops  []ColorStop
	Params []ParamSet
}

func ReadPaletteFile(fileName string) ([]string, map[string]Palette, error) {
	var jsonPalList []JsonPalette
    var palNames []string
	var palMap map[string]Palette
	var pal Palette

    fh, err := os.Open(fileName)
    if err != nil {
        return nil, nil, err
    }
    defer fh.Close()

    dec := json.NewDecoder(fh)
    dec.DisallowUnknownFields()
    if err := dec.Decode(&jsonPalList); err != nil {
        return nil, nil, err
    }

    palNames = make([]string, 0)
	palMap = make(map[string]Palette, 0)
	for _, jsonPal := range jsonPalList {
        palNames = append(palNames, jsonPal.Name)
		switch {
		case len(jsonPal.Stops) > 0:
			pal = NewPaletteByStops(jsonPal.Name, jsonPal.Stops...)
		case len(jsonPal.Colors) > 0:
			pal = NewPaletteByColors(jsonPal.Name, jsonPal.Colors...)
		case len(jsonPal.Params) > 0:
            pal = NewPaletteByParams(jsonPal.Name, jsonPal.Params...)
		}
		palMap[pal.Name()] = pal
	}
	return palNames, palMap, nil
}
