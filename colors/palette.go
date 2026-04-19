package colors

import (
	"cmp"
	"encoding/json"
	"io"
	"log"
	"math"
	"os"
	"slices"
)

type Palette interface {
	Name() string
	Type() PaletteType
	Color(t float64) RGBA
}

type PaletteType byte

const (
	LinearType PaletteType = iota
	ColorStopsType
	FunctionType
	SimpleType
)

func (t PaletteType) String() string {
	switch t {
	case LinearType:
		return "Linear"
	case ColorStopsType:
		return "ColorStops"
	case FunctionType:
		return "Function"
	case SimpleType:
		return "Simple"
	default:
		return "(unknown type)"
	}
}

// ---------------------------------------------------------------------------

type paletteEmb struct {
	name   string
	typ    PaletteType
}

func (p *paletteEmb) Name() string {
	return p.name
}

func (p *paletteEmb) Type() PaletteType {
	return p.typ
}

// ---------------------------------------------------------------------------

type ColorIdent byte
type ColorMask byte

const (
	RedId ColorIdent = iota
	GreenId
	BlueId
	AlphaId
	NumColorIds
)

const (
	RedMask ColorMask = (1 << iota)
	GreenMask
	BlueMask
	AlphaMask
)

func (c ColorMask) String() (s string) {
	if (c & RedMask) != 0 {
		s += "R"
	}
	if (c & GreenMask) != 0 {
		s += "G"
	}
	if (c & BlueMask) != 0 {
		s += "B"
	}
	if (c & AlphaMask) != 0 {
		s += "A"
	}
	return s
}

func (c *ColorMask) UnmarshalText(text []byte) error {
	for _, ch := range text {
		switch ch {
		case 'R':
			*c |= RedMask
		case 'G':
			*c |= GreenMask
		case 'B':
			*c |= BlueMask
		case 'A':
			*c |= AlphaMask
		default:
			log.Fatalf("no such color identifier: %c", ch)
		}
	}
	return nil
}

// ---------------------------------------------------------------------------

type SimplePalette struct {
	paletteEmb
	colors []RGBA
}

func NewSimplePalette(name string, n int) *SimplePalette {
	p := &SimplePalette{}
	p.name = name
	p.typ = 0
	p.colors = make([]RGBA, n)
	return p
}

func NewPaletteFromOther(name string, n int, pal Palette) *SimplePalette {
	p := NewSimplePalette(name, n)
	for i := range n {
		t := float64(i) / float64(n-1)
		p.SetColor(i, pal.Color(t))
	}
	return p
}

func (p *SimplePalette) SetColor(i int, c RGBA) {
	p.colors[i] = c
}

func (p *SimplePalette) Color(t float64) RGBA {
	if t == 0.0 {
		return p.colors[0]
	}
	if t == 1.0 {
		return p.colors[len(p.colors)-1]
	}
	i := int(math.Round(t * float64(len(p.colors)-1)))
	return p.colors[i]
}

// ---------------------------------------------------------------------------

// Bei einer Palette des Typs 'Linear' werden eine feste Anzahl
// RGB-Farben gleichmaessig, d.h. aequidistant uber das Interval [0, 1]
// verteilt. Beim Ermitteln einer Farbe mit der Funktion Color(t) wird
// zwischen zwei hinterlegen Farben linear interpoliert.
type LinearPalette struct {
	paletteEmb
	colors []RGBA
}

func NewLinearPalette(name string) *LinearPalette {
	p := &LinearPalette{}
	p.name = name
	p.typ = LinearType
	p.colors = make([]RGBA, 0)
	return p
}

func NewPaletteByColors(name string, colors ...RGBA) *LinearPalette {
	p := NewLinearPalette(name)
	for _, c := range colors {
		p.AddColor(c)
	}
	return p
}

func (p *LinearPalette) AddColor(c RGBA) {
	p.colors = append(p.colors, c)
}

func (p *LinearPalette) Color(t float64) (c RGBA) {
	if t == 0.0 {
		return p.colors[0]
	}
	if t == 1.0 {
		return p.colors[len(p.colors)-1]
	}
	i, u := math.Modf(t * float64(len(p.colors)-1))
	col1 := p.colors[int(i)]
	col2 := p.colors[int(i)+1]
	return col1.Interpolate(col2, u)
}

// ---------------------------------------------------------------------------

// ColorStops-Paletten basieren auf einer Anzahl Farben (Stuetzstellen)
// zwischen denen eine Farbe interpoliert werden kann. Jede Stuetzstelle
// besteht aus einer Position (Zahl im Intervall [0,1]) und einem dazu
// gehoerenden Farbwert fuer die insgesamt 4 Farbkanaele (R, G, B und A)
type ColorStopsPalette struct {
	paletteEmb
	stops [][]ValueStop
}

type ValueStop struct {
	Pos   float64
	Value uint8
}

// Dieser Typ wird nur fuer das Einlesen der Paletten-Daten aus einem
// JSON-File verwendet.
type ColorStop struct {
	Pos    float64   `json:"pos"`
	Color  RGBA      `json:"color"`
	Ignore ColorMask `json:"ignore,omitempty"`
}

func NewColorStopsPalette(name string) *ColorStopsPalette {
	p := &ColorStopsPalette{}
	p.name = name
	p.typ = ColorStopsType
	p.stops = make([][]ValueStop, 4)
	for i := range NumColorIds {
		value := uint8(0)
		if i == AlphaId {
			value = 0xFF
		}
		p.stops[i] = []ValueStop{
			{Pos: 0.0, Value: value},
			{Pos: 1.0, Value: value},
		}
	}
	return p
}

func NewPaletteByStops(name string, stops ...ColorStop) *ColorStopsPalette {
	if len(stops) < 2 {
		log.Fatalf("at least two stops must be provided: got only %d", len(stops))
	}
	p := NewColorStopsPalette(name)
	for _, stop := range stops {
		p.SetColorStop(stop)
	}
	return p
}

func (p *ColorStopsPalette) SetColorStop(colStop ColorStop) {
	if colStop.Pos < 0.0 || colStop.Pos > 1.0 {
		log.Fatalf("Position must be in [0,1]; is: %f", colStop.Pos)
	}
	for colId := range NumColorIds {
		colMask := ColorMask(1 << colId)
		if colStop.Ignore&colMask != 0 {
			continue
		}
		value := uint8(0)
		switch colId {
		case RedId:
			value = colStop.Color.R
		case GreenId:
			value = colStop.Color.G
		case BlueId:
			value = colStop.Color.B
		case AlphaId:
			value = colStop.Color.A
		}
		p.SetValueStop(colId, ValueStop{colStop.Pos, value})
	}
}

func (p *ColorStopsPalette) SetValueStop(colId ColorIdent, valStop ValueStop) {
	if valStop.Pos < 0.0 || valStop.Pos > 1.0 {
		log.Fatalf("Position must be in [0,1]; is: %f", valStop.Pos)
	}
	for i, stop := range p.stops[colId] {
		if stop.Pos == valStop.Pos {
			p.stops[colId][i].Value = valStop.Value
			return
		}
		if stop.Pos > valStop.Pos {
			p.stops[colId] = slices.Insert(p.stops[colId], i, valStop)
			return
		}
	}
}

// Hier nun spielt die Musik: aufgrund des Wertes t (muss im Intervall [0,1]
// liegen) wird eine neue Farbe interpoliert.
func (p *ColorStopsPalette) Color(t float64) (c RGBA) {
	c.R = p.Value(RedId, t)
	c.G = p.Value(GreenId, t)
	c.B = p.Value(BlueId, t)
	c.A = p.Value(AlphaId, t)
	return c
}

func (p *ColorStopsPalette) Value(colId ColorIdent, t float64) (v uint8) {
	var pStop, nStop ValueStop

	if t == 0.0 {
		return p.stops[colId][0].Value
	}
	if t == 1.0 {
		return p.stops[colId][len(p.stops[colId])-1].Value
	}
	pStop = p.stops[colId][0]
	for _, nStop = range p.stops[colId][1:] {
		if nStop.Pos > t {
			break
		}
		pStop = nStop
	}
	t = (t - pStop.Pos) / (nStop.Pos - pStop.Pos)
	v = uint8(float64(pStop.Value) + ipf(t)*(float64(nStop.Value)-float64(pStop.Value)))
	return v
}

// ----------------------------------------------------------------------------

type FunctionPalette struct {
	paletteEmb
	params []ParamSet
}

type ParamSet struct {
	P0 float64 `json:"p0"`
	P1 float64 `json:"p1"`
	P2 float64 `json:"p2"`
	P3 float64 `json:"p3"`
}

func NewFunctionPalette(name string) *FunctionPalette {
	p := &FunctionPalette{}
	p.name = name
	p.typ = FunctionType
	p.params = []ParamSet{
		{0.0, 0.0, 0.0, 0.0},
		{0.0, 0.0, 0.0, 0.0},
		{0.0, 0.0, 0.0, 0.0},
		{1.0, 0.0, 0.0, 0.0},
	}
	return p
}

func NewPaletteByParams(name string, params ...ParamSet) *FunctionPalette {
	p := NewFunctionPalette(name)
	for i, param := range params {
		p.SetParam(ColorIdent(i), param)
	}
	return p
}

func (p *FunctionPalette) SetParam(colId ColorIdent, param ParamSet) {
	p.params[colId] = param
}

func (p *FunctionPalette) Color(t float64) (c RGBA) {
	r := p.params[0].Value(t)
	g := p.params[1].Value(t)
	b := p.params[2].Value(t)
	a := p.params[3].Value(t)
	return RGBA{
		uint8(255.0 * r),
		uint8(255.0 * g),
		uint8(255.0 * b),
		uint8(255.0 * a),
	}
}

func (p ParamSet) Value(t float64) float64 {
	return p.P0 + p.P1*math.Cos(2*math.Pi*(t*p.P2+p.P3))
}

// ---------------------------------------------------------------------------

type JsonPalette struct {
	Name   string      `json:"name"`
	Colors []RGBA      `json:"colors,omitempty"`
	Stops  []ColorStop `json:"stops,omitempty"`
	Params []ParamSet  `json:"params,omitempty"`
}

// ---------------------------------------------------------------------------

func ReadPaletteFile(fileName string) ([]string, map[string]Palette, error) {
	fh, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("couldn't open file %s: %v", fileName, err.Error())
	}
	defer fh.Close()
	return ReadPaletteData(fh)
}

func ReadPaletteData(fh io.Reader) ([]string, map[string]Palette, error) {
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
			log.Printf("palette '%s' is already defined", jsonPal.Name)
			continue
		}
		switch {
		case len(jsonPal.Colors) > 0:
			pal = NewPaletteByColors(jsonPal.Name, jsonPal.Colors...)
		case len(jsonPal.Stops) > 0:
			pal = NewPaletteByStops(jsonPal.Name, jsonPal.Stops...)
		case len(jsonPal.Params) > 0:
			pal = NewPaletteByParams(jsonPal.Name, jsonPal.Params...)
		default:
			log.Printf("palette '%s' has no known format", jsonPal.Name)
			continue
		}
		palNames = append(palNames, jsonPal.Name)
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

// Mapper
type Mapper func(t float64) float64

var (
	DefaultMapper = func(t float64) float64 {
		return t
	}
)

func NewMapper(minIn, maxIn, minOut, maxOut float64, wrap bool) Mapper {
	diffIn := maxIn - minIn
	diffOut := maxOut - minOut
	quot := diffOut / diffIn

	if wrap {
		return func(t float64) float64 {
			t = math.Mod(t-minIn, diffIn)
			if t < 0.0 {
				t += diffIn
			}
			return minOut + t*quot
		}
	} else {
		return func(t float64) float64 {
			return minOut + (t-minIn)*quot
		}
	}
}
