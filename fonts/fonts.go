//go:generate go run gen.go

// Bietet einen einfachen Zugriff auf die Go-Fonts aber auch auf eine Reihe
// von OpenSource-Schriten.
package fonts

import (
	"errors"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

type Font struct {
	Id int
	opentype.Font
}

func NewFont(id int, data []byte) *Font {
    d, _ := opentype.Parse(data)
    f := &Font{id, *d}
    return f
}

// func Parse(data []byte) *Font {
func Parse(data []byte) *opentype.Font {
	f, _ := opentype.Parse(data)
    return f
	// return &Font{*f, }
	// return (*Font)(f)
}

func (f *Font) MarshalText() ([]byte, error) {
	for key, val := range Map {
		if (*Font)(val) == f {
			return []byte(key), nil
		}
	}
	return []byte{}, errors.New("Font not found")
}

func (f *Font) UnmarshalText(text []byte) error {
	key := string(text)
	if val, ok := Map[key]; ok {
		*f = (Font)(*val)
		return nil
	}
	return errors.New("Fontname not found")
}

// Erstellt einen neuen Fontface, der bspw. bei der Methode [SetFontFace]
// verwendet werden kann. textFont ist ein Pointer auf einen OpenType-Font
// Siehe auch Array [Names] für eine Liste aller Fonts, die in diesem Package
// angeboten werden.
func NewFace(textFont *Font, size float64) (font.Face, error) {
	face, err := opentype.NewFace(&(textFont.Font),
		&opentype.FaceOptions{
			Size:    size,
			DPI:     100,
			Hinting: font.HintingFull,
		})
	if err != nil {
		return nil, err
	}
	return face, nil
}
