// Bietet einfachen Zugriff auf die Go-Fonts aber auch auf eine Reihen
// OpenSource-Schriten.
package fonts

import (
    "log"

    "golang.org/x/image/font"
    "golang.org/x/image/font/opentype"
)

func NewFace(textFont *opentype.Font, size float64) font.Face {
    face, _ := opentype.NewFace(textFont,
        &opentype.FaceOptions{
            Size:    size,
            DPI:     72,
            Hinting: font.HintingFull,
        })
    return face
}

func check(fontName string, err error) {
    if err != nil {
        log.Fatalf("error loading font %s: %v", fontName, err)
    }
}
