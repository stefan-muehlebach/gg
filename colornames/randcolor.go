package colornames

import (
    "math/rand"
    "github.com/stefan-muehlebach/gg/color"
)

func RandColor() (color.Color) {
    name := Names[rand.Int() % len(Names)]
    return Map[name]
}

