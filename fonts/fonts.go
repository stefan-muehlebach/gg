// Bietet einen einfachen Zugriff auf die Go-Fonts aber auch auf eine Reihe
// von OpenSource-Schriten.
package fonts

import (
    "golang.org/x/image/font"
    "golang.org/x/image/font/opentype"
)

// Erstellt einen neuen Fontface, der bspw. bei der Methode [SetFontFace]
// verwendet werden kann. textFont ist ein Pointer auf einen OpenType-Font
// Siehe auch Array [Names] für eine Liste aller Fonts, die in diesem Package
// angeboten werden.
func NewFace(textFont *opentype.Font, size float64) font.Face {
    face, _ := opentype.NewFace(textFont,
        &opentype.FaceOptions{
            Size:    size,
            DPI:     72,
            Hinting: font.HintingFull,
        })
    return face
}
