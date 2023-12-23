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

func RandGroupColor(group ColorGroup) (color.Color) {
    nameList, ok := Groups[group]
    if !ok {
        return Black
    }
    name := nameList[rand.Int() % len(nameList)]
    return Map[name]
}

