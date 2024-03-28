package fonts

// WICHTIG: Diese Datei sollte nicht manuell angepasst werden!
// Sie wird automatisch per Script neu erzeugt. Allfaellige manuelle
// Anpassungen werden damit ueberschrieben.

import (
    "embed"
//    "golang.org/x/image/font/opentype"
    "golang.org/x/image/font/gofont/goregular"
    "golang.org/x/image/font/gofont/goitalic"
    "golang.org/x/image/font/gofont/gomedium"
    "golang.org/x/image/font/gofont/gomediumitalic"
    "golang.org/x/image/font/gofont/gobold"
    "golang.org/x/image/font/gofont/gobolditalic"
    "golang.org/x/image/font/gofont/gomono"
    "golang.org/x/image/font/gofont/gomonoitalic"
    "golang.org/x/image/font/gofont/gomonobold"
    "golang.org/x/image/font/gofont/gomonobolditalic"
    "golang.org/x/image/font/gofont/gosmallcaps"
    "golang.org/x/image/font/gofont/gosmallcapsitalic"
)

//go:embed FontFiles/*.ttf FontFiles/*.otf
var fontFiles embed.FS

var (
    lucidabright, _ = fontFiles.ReadFile("FontFiles/00-LucidaBright.ttf")
    lucidabrightitalic, _ = fontFiles.ReadFile("FontFiles/01-LucidaBright-Italic.ttf")
    lucidabrightdemibold, _ = fontFiles.ReadFile("FontFiles/02-LucidaBright-Demibold.ttf")
    lucidabrightdemibolditalic, _ = fontFiles.ReadFile("FontFiles/03-LucidaBright-Demibold-Italic.ttf")
    lucidasans, _ = fontFiles.ReadFile("FontFiles/05-LucidaSans.ttf")
    lucidasansitalic, _ = fontFiles.ReadFile("FontFiles/06-LucidaSans-Italic.ttf")
    lucidasansdemiboldroman, _ = fontFiles.ReadFile("FontFiles/07-LucidaSans-Demibold-Roman.ttf")
    lucidasansdemibolditalic, _ = fontFiles.ReadFile("FontFiles/08-LucidaSans-Demibold-Italic.ttf")
    lucidasanstypewriter, _ = fontFiles.ReadFile("FontFiles/10-LucidaSansTypewriter.ttf")
    lucidasanstypewriteroblique, _ = fontFiles.ReadFile("FontFiles/11-LucidaSansTypewriter-Oblique.ttf")
    lucidasanstypewriterbold, _ = fontFiles.ReadFile("FontFiles/12-LucidaSansTypewriter-Bold.ttf")
    lucidasanstypewriterboldoblique, _ = fontFiles.ReadFile("FontFiles/13-LucidaSansTypewriter-Bold-Oblique.ttf")
    lucidafax, _ = fontFiles.ReadFile("FontFiles/15-LucidaFax.ttf")
    lucidafaxitalic, _ = fontFiles.ReadFile("FontFiles/16-LucidaFax-Italic.ttf")
    lucidafaxdemibold, _ = fontFiles.ReadFile("FontFiles/17-LucidaFax-Demibold.ttf")
    lucidafaxdemibolditalic, _ = fontFiles.ReadFile("FontFiles/18-LucidaFax-Demibold-Italic.ttf")
    lucidaconsole, _ = fontFiles.ReadFile("FontFiles/20-LucidaConsole.ttf")
    lucidahandwritingitalic, _ = fontFiles.ReadFile("FontFiles/21-LucidaHandwriting-Italic.ttf")
    lucidacalligraphy, _ = fontFiles.ReadFile("FontFiles/22-LucidaCalligraphy.ttf")
    lucidacalligraphybold, _ = fontFiles.ReadFile("FontFiles/23-LucidaCalligraphy-Bold.ttf")
    lucidablackletter, _ = fontFiles.ReadFile("FontFiles/24-LucidaBlackletter.ttf")
    seaford, _ = fontFiles.ReadFile("FontFiles/30-Seaford.ttf")
    seaforditalic, _ = fontFiles.ReadFile("FontFiles/31-Seaford-Italic.ttf")
    seafordbold, _ = fontFiles.ReadFile("FontFiles/32-Seaford-Bold.ttf")
    seafordbolditalic, _ = fontFiles.ReadFile("FontFiles/33-Seaford-Bold-Italic.ttf")
    garamond, _ = fontFiles.ReadFile("FontFiles/35-Garamond.otf")
    garamonditalic, _ = fontFiles.ReadFile("FontFiles/36-Garamond-Italic.otf")
    garamondbold, _ = fontFiles.ReadFile("FontFiles/37-Garamond-Bold.otf")
    elegante, _ = fontFiles.ReadFile("FontFiles/38-Elegante.ttf")
    elegantebold, _ = fontFiles.ReadFile("FontFiles/39-Elegante-Bold.ttf")
    leipzigfraktur, _ = fontFiles.ReadFile("FontFiles/40-LeipzigFraktur.ttf")
    leipzigfrakturbold, _ = fontFiles.ReadFile("FontFiles/41-LeipzigFraktur-Bold.ttf")
    rothenburgdecorative, _ = fontFiles.ReadFile("FontFiles/42-RothenburgDecorative.ttf")
    elzevier, _ = fontFiles.ReadFile("FontFiles/45-Elzevier.ttf")
    floralcapitals, _ = fontFiles.ReadFile("FontFiles/46-FloralCapitals.ttf")
    goudyinitialen, _ = fontFiles.ReadFile("FontFiles/47-GoudyInitialen.ttf")
    mosaicinitialen, _ = fontFiles.ReadFile("FontFiles/48-MosaicInitialen.ttf")
    yinit, _ = fontFiles.ReadFile("FontFiles/49-Yinit.ttf")
)

var (
    GoRegular, _                        = Parse(goregular.TTF)
    GoItalic, _                         = Parse(goitalic.TTF)
    GoMedium, _                         = Parse(gomedium.TTF)
    GoMediumItalic, _                   = Parse(gomediumitalic.TTF)
    GoBold, _                           = Parse(gobold.TTF)
    GoBoldItalic, _                     = Parse(gobolditalic.TTF)
    GoMono, _                           = Parse(gomono.TTF)
    GoMonoItalic, _                     = Parse(gomonoitalic.TTF)
    GoMonoBold, _                       = Parse(gomonobold.TTF)
    GoMonoBoldItalic, _                 = Parse(gomonobolditalic.TTF)
    GoSmallcaps, _                      = Parse(gosmallcaps.TTF)
    GoSmallcapsItalic, _                = Parse(gosmallcapsitalic.TTF)
    LucidaBright, _                     = Parse(lucidabright)
    LucidaBrightItalic, _               = Parse(lucidabrightitalic)
    LucidaBrightDemibold, _             = Parse(lucidabrightdemibold)
    LucidaBrightDemiboldItalic, _       = Parse(lucidabrightdemibolditalic)
    LucidaSans, _                       = Parse(lucidasans)
    LucidaSansItalic, _                 = Parse(lucidasansitalic)
    LucidaSansDemiboldRoman, _          = Parse(lucidasansdemiboldroman)
    LucidaSansDemiboldItalic, _         = Parse(lucidasansdemibolditalic)
    LucidaSansTypewriter, _             = Parse(lucidasanstypewriter)
    LucidaSansTypewriterOblique, _      = Parse(lucidasanstypewriteroblique)
    LucidaSansTypewriterBold, _         = Parse(lucidasanstypewriterbold)
    LucidaSansTypewriterBoldOblique, _  = Parse(lucidasanstypewriterboldoblique)
    LucidaFax, _                        = Parse(lucidafax)
    LucidaFaxItalic, _                  = Parse(lucidafaxitalic)
    LucidaFaxDemibold, _                = Parse(lucidafaxdemibold)
    LucidaFaxDemiboldItalic, _          = Parse(lucidafaxdemibolditalic)
    LucidaConsole, _                    = Parse(lucidaconsole)
    LucidaHandwritingItalic, _          = Parse(lucidahandwritingitalic)
    LucidaCalligraphy, _                = Parse(lucidacalligraphy)
    LucidaCalligraphyBold, _            = Parse(lucidacalligraphybold)
    LucidaBlackletter, _                = Parse(lucidablackletter)
    Seaford, _                          = Parse(seaford)
    SeafordItalic, _                    = Parse(seaforditalic)
    SeafordBold, _                      = Parse(seafordbold)
    SeafordBoldItalic, _                = Parse(seafordbolditalic)
    Garamond, _                         = Parse(garamond)
    GaramondItalic, _                   = Parse(garamonditalic)
    GaramondBold, _                     = Parse(garamondbold)
    Elegante, _                         = Parse(elegante)
    EleganteBold, _                     = Parse(elegantebold)
    LeipzigFraktur, _                   = Parse(leipzigfraktur)
    LeipzigFrakturBold, _               = Parse(leipzigfrakturbold)
    RothenburgDecorative, _             = Parse(rothenburgdecorative)
    Elzevier, _                         = Parse(elzevier)
    FloralCapitals, _                   = Parse(floralcapitals)
    GoudyInitialen, _                   = Parse(goudyinitialen)
    MosaicInitialen, _                  = Parse(mosaicinitialen)
    Yinit, _                            = Parse(yinit)
)

var Map = map[string]*Font{
    "GoRegular":                        GoRegular,
    "GoItalic":                         GoItalic,
    "GoMedium":                         GoMedium,
    "GoMediumItalic":                   GoMediumItalic,
    "GoBold":                           GoBold,
    "GoBoldItalic":                     GoBoldItalic,
    "GoMono":                           GoMono,
    "GoMonoItalic":                     GoMonoItalic,
    "GoMonoBold":                       GoMonoBold,
    "GoMonoBoldItalic":                 GoMonoBoldItalic,
    "GoSmallcaps":                      GoSmallcaps,
    "GoSmallcapsItalic":                GoSmallcapsItalic,
    "LucidaBright":                     LucidaBright,
    "LucidaBrightItalic":               LucidaBrightItalic,
    "LucidaBrightDemibold":             LucidaBrightDemibold,
    "LucidaBrightDemiboldItalic":       LucidaBrightDemiboldItalic,
    "LucidaSans":                       LucidaSans,
    "LucidaSansItalic":                 LucidaSansItalic,
    "LucidaSansDemiboldRoman":          LucidaSansDemiboldRoman,
    "LucidaSansDemiboldItalic":         LucidaSansDemiboldItalic,
    "LucidaSansTypewriter":             LucidaSansTypewriter,
    "LucidaSansTypewriterOblique":      LucidaSansTypewriterOblique,
    "LucidaSansTypewriterBold":         LucidaSansTypewriterBold,
    "LucidaSansTypewriterBoldOblique":  LucidaSansTypewriterBoldOblique,
    "LucidaFax":                        LucidaFax,
    "LucidaFaxItalic":                  LucidaFaxItalic,
    "LucidaFaxDemibold":                LucidaFaxDemibold,
    "LucidaFaxDemiboldItalic":          LucidaFaxDemiboldItalic,
    "LucidaConsole":                    LucidaConsole,
    "LucidaHandwritingItalic":          LucidaHandwritingItalic,
    "LucidaCalligraphy":                LucidaCalligraphy,
    "LucidaCalligraphyBold":            LucidaCalligraphyBold,
    "LucidaBlackletter":                LucidaBlackletter,
    "Seaford":                          Seaford,
    "SeafordItalic":                    SeafordItalic,
    "SeafordBold":                      SeafordBold,
    "SeafordBoldItalic":                SeafordBoldItalic,
    "Garamond":                         Garamond,
    "GaramondItalic":                   GaramondItalic,
    "GaramondBold":                     GaramondBold,
    "Elegante":                         Elegante,
    "EleganteBold":                     EleganteBold,
    "LeipzigFraktur":                   LeipzigFraktur,
    "LeipzigFrakturBold":               LeipzigFrakturBold,
    "RothenburgDecorative":             RothenburgDecorative,
    "Elzevier":                         Elzevier,
    "FloralCapitals":                   FloralCapitals,
    "GoudyInitialen":                   GoudyInitialen,
    "MosaicInitialen":                  MosaicInitialen,
    "Yinit":                            Yinit,
}

var Names = []string{
    "GoRegular",
    "GoItalic",
    "GoMedium",
    "GoMediumItalic",
    "GoBold",
    "GoBoldItalic",
    "GoMono",
    "GoMonoItalic",
    "GoMonoBold",
    "GoMonoBoldItalic",
    "GoSmallcaps",
    "GoSmallcapsItalic",
    "LucidaBright",
    "LucidaBrightItalic",
    "LucidaBrightDemibold",
    "LucidaBrightDemiboldItalic",
    "LucidaSans",
    "LucidaSansItalic",
    "LucidaSansDemiboldRoman",
    "LucidaSansDemiboldItalic",
    "LucidaSansTypewriter",
    "LucidaSansTypewriterOblique",
    "LucidaSansTypewriterBold",
    "LucidaSansTypewriterBoldOblique",
    "LucidaFax",
    "LucidaFaxItalic",
    "LucidaFaxDemibold",
    "LucidaFaxDemiboldItalic",
    "LucidaConsole",
    "LucidaHandwritingItalic",
    "LucidaCalligraphy",
    "LucidaCalligraphyBold",
    "LucidaBlackletter",
    "Seaford",
    "SeafordItalic",
    "SeafordBold",
    "SeafordBoldItalic",
    "Garamond",
    "GaramondItalic",
    "GaramondBold",
    "Elegante",
    "EleganteBold",
    "LeipzigFraktur",
    "LeipzigFrakturBold",
    "RothenburgDecorative",
    "Elzevier",
    "FloralCapitals",
    "GoudyInitialen",
    "MosaicInitialen",
    "Yinit",
}

