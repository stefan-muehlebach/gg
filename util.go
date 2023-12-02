package gg

import (
    "fmt"
    "image"
    "image/color"
    "image/draw"
    "image/jpeg"
    _ "image/jpeg"
    "image/png"
    "io/ioutil"
    "math"
    "os"
    "strings"

    "github.com/golang/freetype/truetype"

    "golang.org/x/image/font"
    "golang.org/x/image/math/fixed"
)

// Umrechnungsfunktionen zwischen Grad und Bogenmass.

func Radians(degrees float64) float64 {
    return degrees * math.Pi / 180
}

func Degrees(radians float64) float64 {
    return radians * 180 / math.Pi
}

// Funktionen zum Speichern und Laden von Bildern in den Formaten PNG und JPG

func LoadImage(path string) (image.Image, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()
    im, _, err := image.Decode(file)
    return im, err
}

func LoadPNG(path string) (image.Image, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()
    return png.Decode(file)
}

func SavePNG(path string, im image.Image) error {
    file, err := os.Create(path)
    if err != nil {
        return err
    }
    defer file.Close()
    return png.Encode(file, im)
}

func LoadJPG(path string) (image.Image, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()
    return jpeg.Decode(file)
}

func SaveJPG(path string, im image.Image, quality int) error {
    file, err := os.Create(path)
    if err != nil {
        return err
    }
    defer file.Close()

    var opt jpeg.Options
    opt.Quality = quality

    return jpeg.Encode(file, im, &opt)
}

// LoadFontFace is a helper function to load the specified font file with
// the specified point size. Note that the returned `font.Face` objects
// are not thread safe and cannot be used in parallel across goroutines.
// You can usually just use the Context.LoadFontFace function instead of
// this package-level function.
func LoadFontFace(path string, points float64) (font.Face, error) {
    fontBytes, err := ioutil.ReadFile(path)
    if err != nil {
        return nil, err
    }
    f, err := truetype.Parse(fontBytes)
    if err != nil {
        return nil, err
    }
    face := truetype.NewFace(f, &truetype.Options{
        Size: points,
        // Hinting: font.HintingFull,
    })
    return face, nil
}

// Hilfsfunktionen zum Erstellen von Farben.

// Erzeugt ein neues RGB-Farbobjekt, und verwendet die Werte r, g, b in [0,1]
// als Intensitätswerte der Farben Rot, Grün und Blau.
func NewRGB(r, g, b float64) (color.Color) {
    return NewRGBA(r, g, b, 1.0)
}

// Analog zu [NewRGB], jedoch mit a in [0,1] als Alpha-Wert (Deckung der
// Farbe). a=1 bedeutet volle Deckung und a=0 bedeutet vollständige
// Transparenz. Die Werte für r, g und b sind als nicht-normalisierte Werte
// anzugeben!
func NewRGBA(r, g, b, a float64) (color.Color) {
    return color.NRGBA{
        uint8(r * 255),
        uint8(g * 255),
        uint8(b * 255),
        uint8(a * 255),
    }
}

func NewRGB255(r, g, b int) (color.Color) {
    return NewRGBA255(r, g, b, 255)
}

func NewRGBA255(r, g, b, a int) (color.Color) {
    return color.NRGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
}

// Interne Funktionen, nicht fuer den oeffentlichen Gebrauch.

func imageToRGBA(src image.Image) *image.RGBA {
    bounds := src.Bounds()
    dst := image.NewRGBA(bounds)
    draw.Draw(dst, bounds, src, bounds.Min, draw.Src)
    return dst
}

func parseHexColor(x string) (r, g, b, a int) {
    x = strings.TrimPrefix(x, "#")
    a = 255
    if len(x) == 3 {
        format := "%1x%1x%1x"
        fmt.Sscanf(x, format, &r, &g, &b)
        r |= r << 4
        g |= g << 4
        b |= b << 4
    }
    if len(x) == 6 {
        format := "%02x%02x%02x"
        fmt.Sscanf(x, format, &r, &g, &b)
    }
    if len(x) == 8 {
        format := "%02x%02x%02x%02x"
        fmt.Sscanf(x, format, &r, &g, &b, &a)
    }
    return
}

func fixp(x, y float64) fixed.Point26_6 {
    return fixed.Point26_6{X: fix(x), Y: fix(y)}
}

func fix(x float64) fixed.Int26_6 {
    return fixed.Int26_6(math.Round(x * 64))
}

func unfix(x fixed.Int26_6) float64 {
    const shift, mask = 6, 1<<6 - 1
    if x >= 0 {
        return float64(x>>shift) + float64(x&mask)/64
    }
    x = -x
    if x >= 0 {
        return -(float64(x>>shift) + float64(x&mask)/64)
    }
    return 0
}

