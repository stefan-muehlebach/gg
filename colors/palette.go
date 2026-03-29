package colors

import (
	"cmp"
	"encoding/json"
	"io"
	"log"
	"math"
	"slices"
)

type Palette interface {
	Name() string
	Type() PaletteType
	Color(t float64) RGBA
}

var (
	ipf = linearInterp
)

func linearInterp(t float64) float64 {
	return t
}

func powerInterp(t float64) float64 {
	a := 1.2
	t1 := math.Pow(2, a-1.0)
	if t <= 0.5 {
		return t1 * math.Pow(t, a)
	} else {
		return 1.0 - t1*math.Pow(1.0-t, a)
	}
}

func cubicInterp(t float64) float64 {
	return 3.0*t*t - 2.0*t*t*t
}

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

type PaletteType byte

const (
	LinearGradientType PaletteType = iota
	GradientType
	FunctionType
)

func (t PaletteType) String() string {
	switch t {
	case LinearGradientType:
		return "Linear Gradient"
	case GradientType:
		return "Gradient"
	case FunctionType:
		return "Function"
	default:
		return "(unknown type)"
	}
}

// ---------------------------------------------------------------------------

type LinGradPalette struct {
	name   string
	colors []RGBA
	dt     float64
}

func NewPaletteByColors(name string, colors ...RGBA) *LinGradPalette {
	p := &LinGradPalette{}
	p.name = name
	p.colors = make([]RGBA, len(colors))
	p.dt = 1.0 / float64(len(colors)-1)
	copy(p.colors, colors)
	return p
}

func (p *LinGradPalette) Name() string {
	return p.name
}

func (p *LinGradPalette) Type() PaletteType {
	return LinearGradientType
}

func (p *LinGradPalette) Color(t float64) (c RGBA) {
	if t == 0.0 {
		return p.colors[0]
	}
	if t == 1.0 {
		return p.colors[len(p.colors)-1]
	}
	i, u := math.Modf(t / p.dt)
	col1 := p.colors[int(i)]
	col2 := p.colors[int(i)+1]
	return col1.Interpolate(col2, ipf(u))
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

func NewPaletteByStops(name string, stops ...ColorStop) *GradPalette {
	if len(stops) < 2 {
		log.Fatalf("at least two stops must be provided: got only %d", len(stops))
	}
	p := &GradPalette{}
	p.name = name
	p.stops = []ColorStop{
		{Pos: 0.0, Color: black},
		{Pos: 1.0, Color: black},
	}
	for _, stop := range stops {
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

func (p *GradPalette) Type() PaletteType {
	return GradientType
}

// Hier nun spielt die Musik: aufgrund des Wertes t (muss im Intervall [0,1]
// liegen) wird eine neue Farbe interpoliert.
func (p *GradPalette) Color(t float64) (c RGBA) {
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

func (p *ProcPalette) Type() PaletteType {
	return FunctionType
}

func (p *ProcPalette) Color(t float64) (c RGBA) {
	r := p.params[0].Value(t)
	g := p.params[1].Value(t)
	b := p.params[2].Value(t)
	return RGBA{uint8(255.0 * r), uint8(255.0 * g), uint8(255.0 * b), 0xFF}
}

// ---------------------------------------------------------------------------

type JsonPalette struct {
	Name   string      `json:"name"`
	Colors []RGBA      `json:"colors,omitempty"`
	Stops  []ColorStop `json:"stops,omitempty"`
	Params []ParamSet  `json:"params,omitempty"`
}

type ColorStop struct {
	Pos    float64    `json:"pos"`
	Color  RGBA       `json:"color"`
	Ignore ColorIdent `json:"ignore,omitempty"`
}

type ParamSet struct {
	Y0 float64 `json:"y0"`
	Y1 float64 `json:"y1"`
	X0 float64 `json:"x0"`
	X1 float64 `json:"x1"`
}

func (p ParamSet) Value(t float64) float64 {
	return p.Y0 + p.Y1*math.Cos(2*math.Pi*(t*p.X0+p.X1))
}

func ReadPaletteData(fh io.Reader) ([]string, map[string]Palette, error) {

	// func ReadPaletteFile(fileName string) ([]string, map[string]Palette, error) {
	var jsonPalList []JsonPalette
	var palNames []string
	var palMap map[string]Palette
	var pal Palette

	dec := json.NewDecoder(fh)
	dec.DisallowUnknownFields()
	if err := dec.Decode(&jsonPalList); err != nil {
		return nil, nil, err
	}

	palNames = make([]string, 0)
	palMap = make(map[string]Palette)
	for _, jsonPal := range jsonPalList {
		if _, found := palMap[jsonPal.Name]; found {
			log.Fatalf("palette %s already defined", jsonPal.Name)
			continue
		}
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
	slices.SortFunc(palNames, func(nameA, nameB string) int {
		typeA := palMap[nameA].Type()
		typeB := palMap[nameB].Type()
		if typeA != typeB {
			return cmp.Compare(typeA, typeB)
		}
		return cmp.Compare(nameA, nameB)
	})
	return palNames, palMap, nil
}
