// Enthält alle benannten Farben aus SVG 1.1 als HSV-Farben.
package colornames

import (
    "math/rand"
    "github.com/stefan-muehlebach/gg/color"
)

// Mit RandColor kann zufällig eine aus dem gesamten Sortiment der hier
// definierten Farben gewählt werden. Hilfreich für Tests, Beispielprogramme
// oder anderes.
func RandColor() (color.Color) {
    name := Names[rand.Int() % len(Names)]
    return Map[name]
}

