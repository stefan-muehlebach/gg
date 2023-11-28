package font

import (
//    "path/filepath"
//    "io/ioutil"
    "log"

    "golang.org/x/image/font"
    "golang.org/x/image/font/opentype"

/*
    "golang.org/x/image/font/gofont/gobold"
    "golang.org/x/image/font/gofont/gobolditalic"
    "golang.org/x/image/font/gofont/goitalic"
    "golang.org/x/image/font/gofont/gomedium"
    "golang.org/x/image/font/gofont/gomediumitalic"
    "golang.org/x/image/font/gofont/gomono"
    "golang.org/x/image/font/gofont/gomonobold"
    "golang.org/x/image/font/gofont/gomonobolditalic"
    "golang.org/x/image/font/gofont/gomonoitalic"
    "golang.org/x/image/font/gofont/goregular"
    "golang.org/x/image/font/gofont/gosmallcaps"
    "golang.org/x/image/font/gofont/gosmallcapsitalic"    
*/
)

const (
    //pkgBaseDir = "/home/dietpi/Go/src/mju.net/font"
)

/*
var (
    Names []string
    Map   map[string]*opentype.Font
)

func init() {    
    Names = make([]string, 0)
    Map = make(map[string]*opentype.Font)
    
    ParseFont("Go", goregular.TTF)
    ParseFont("GoItalic", goitalic.TTF)
    ParseFont("GoMedium", gomedium.TTF)
    ParseFont("GoMediumItalic", gomediumitalic.TTF)
    ParseFont("GoBold", gobold.TTF)
    ParseFont("GoBoldItalic", gobolditalic.TTF)

    ParseFont("GoMono", gomono.TTF)
    ParseFont("GoMonoItalic", gomonoitalic.TTF)
    ParseFont("GoMonoBold", gomonobold.TTF)
    ParseFont("GoMonoBoldItalic", gomonobolditalic.TTF)

    ParseFont("GoSmallcaps", gosmallcaps.TTF)
    ParseFont("GoSmallcapsItalic", gosmallcapsitalic.TTF)

    LoadFont("LucidaBright", "FontFiles/LucidaBright.ttf")
    LoadFont("LucidaBrightItalic", "FontFiles/LucidaBright-Italic.ttf")
    LoadFont("LucidaBrightDemibold", "FontFiles/LucidaBright-Demibold.ttf")
    LoadFont("LucidaBrightDemiboldItalic", "FontFiles/LucidaBright-Demibold-Italic.ttf")

    LoadFont("LucidaSans", "FontFiles/LucidaSans.ttf")
    LoadFont("LucidaSansItalic", "FontFiles/LucidaSans-Italic.ttf")
    LoadFont("LucidaSansDemiboldRoman", "FontFiles/LucidaSans-Demibold-Roman.ttf")
    LoadFont("LucidaSansDemiboldItalic", "FontFiles/LucidaSans-Demibold-Italic.ttf")

    LoadFont("LucidaSansTypewriter", "FontFiles/LucidaSansTypewriter.ttf")
    LoadFont("LucidaSansTypewriterOblique", "FontFiles/LucidaSansTypewriter-Oblique.ttf")
    LoadFont("LucidaSansTypewriterBold", "FontFiles/LucidaSansTypewriter-Bold.ttf")
    LoadFont("LucidaSansTypewriterBoldOblique", "FontFiles/LucidaSansTypewriter-Bold-Oblique.ttf")

    LoadFont("LucidaFax", "FontFiles/LucidaFax.ttf")
    LoadFont("LucidaFaxItalic", "FontFiles/LucidaFax-Italic.ttf")
    LoadFont("LucidaFaxDemibold", "FontFiles/LucidaFax-Demibold.ttf")
    LoadFont("LucidaFaxDemiboldItalic", "FontFiles/LucidaFax-Demibold-Italic.ttf")

    LoadFont("LucidaConsole", "FontFiles/LucidaConsole.ttf")
    LoadFont("LucidaBlackletter", "FontFiles/LucidaBlackletter.ttf")
    LoadFont("LucidaCalligraphyItalic", "FontFiles/LucidaCalligraphy-Italic.ttf")
    LoadFont("LucidaHandwritingItalic", "FontFiles/LucidaHandwriting-Italic.ttf")

    LoadFont("Seaford", "FontFiles/Seaford.ttf")
    LoadFont("SeafordItalic", "FontFiles/Seaford-Italic.ttf")
    LoadFont("SeafordBold", "FontFiles/Seaford-Bold.ttf")
    LoadFont("SeafordBoldItalic", "FontFiles/Seaford-Bold-Italic.ttf")

    LoadFont("Garamond", "FontFiles/Garamond.otf")
    LoadFont("GaramondItalic", "FontFiles/Garamond-Italic.otf")
    LoadFont("GaramondBold", "FontFiles/Garamond-Bold.otf")

    LoadFont("Elegante", "FontFiles/Elegante.ttf")
    LoadFont("EleganteBold", "FontFiles/Elegante-Bold.ttf")

    LoadFont("LeipzigFraktur", "FontFiles/LeipzigFraktur.ttf")
    LoadFont("LeipzigFrakturBold", "FontFiles/LeipzigFraktur-Bold.ttf")
    LoadFont("RothenburgDecorative", "FontFiles/RothenburgDecorative.ttf")
    
    LoadFont("GoudyInitialen", "FontFiles/GoudyInitialen.ttf")
    LoadFont("MosaicInitialen", "FontFiles/MosaicInitialen.ttf")
    LoadFont("FloralCapitals", "FontFiles/FloralCapitals.ttf")
    LoadFont("Elzevier", "FontFiles/Elzevier.ttf")
    LoadFont("Yinit", "FontFiles/Yinit.ttf")
}
*/

/*
func ParseFont(name string, src []byte) (*opentype.Font) {
    fnt, err := opentype.Parse(src)
    if err != nil {
        log.Fatalf("error loading font %s: %v", name, err)
    }
    Map[name] = fnt
    Names = append(Names, name)
    return fnt
}

func LoadFont(name string, fileName string) (*opentype.Font) {
    fontBytes, err := ioutil.ReadFile(filepath.Join(pkgBaseDir, fileName))
    if err != nil {
        log.Fatalf("error loading font %s: %v", name, err)
    }
    return ParseFont(name, fontBytes)
}
*/

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
