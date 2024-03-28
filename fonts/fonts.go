//go:generate go run gen.go

// Bietet einen einfachen Zugriff auf die Go-Fonts aber auch auf eine Reihe
// von OpenSource-Schriten.
package fonts

import (
    "errors"
    "golang.org/x/image/font"
    "golang.org/x/image/font/opentype"
)

/*
type Font struct {
    *opentype.Font
}
*/

type Font opentype.Font

func Parse(data []byte) (*Font, error) {
    f, err := opentype.Parse(data)
    return (*Font)(f), err
}

// Erstellt einen neuen Fontface, der bspw. bei der Methode [SetFontFace]
// verwendet werden kann. textFont ist ein Pointer auf einen OpenType-Font
// Siehe auch Array [Names] für eine Liste aller Fonts, die in diesem Package
// angeboten werden.
func NewFace(textFont *Font, size float64) font.Face {
    face, _ := opentype.NewFace((*opentype.Font)(textFont),
        &opentype.FaceOptions{
            Size:    size,
            DPI:     72,
            Hinting: font.HintingFull,
        })
    return face
}

func (f *Font) MarshalText() ([]byte, error) {
    for k, v := range Map {
        if v == f {
            return []byte(k), nil
        }
    }
    return []byte{}, errors.New("Fontname not found")
}

