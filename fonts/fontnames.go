package fonts

import (
    "embed"
    "golang.org/x/image/font/opentype"
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
)

//go:embed FontFiles/*.ttf FontFiles/*.otf
var fontFiles embed.FS

var (
    elegantebold, _ = fontFiles.ReadFile("FontFiles/Elegante-Bold.ttf")
    elegante, _ = fontFiles.ReadFile("FontFiles/Elegante.ttf")
    elzevier, _ = fontFiles.ReadFile("FontFiles/Elzevier.ttf")
    floralcapitals, _ = fontFiles.ReadFile("FontFiles/FloralCapitals.ttf")
    garamondbold, _ = fontFiles.ReadFile("FontFiles/Garamond-Bold.otf")
    garamonditalic, _ = fontFiles.ReadFile("FontFiles/Garamond-Italic.otf")
    garamond, _ = fontFiles.ReadFile("FontFiles/Garamond.otf")
    goudyinitialen, _ = fontFiles.ReadFile("FontFiles/GoudyInitialen.ttf")
    leipzigfrakturbold, _ = fontFiles.ReadFile("FontFiles/LeipzigFraktur-Bold.ttf")
    leipzigfraktur, _ = fontFiles.ReadFile("FontFiles/LeipzigFraktur.ttf")
    lucidablackletter, _ = fontFiles.ReadFile("FontFiles/LucidaBlackletter.ttf")
    lucidabrightdemibolditalic, _ = fontFiles.ReadFile("FontFiles/LucidaBright-Demibold-Italic.ttf")
    lucidabrightdemibold, _ = fontFiles.ReadFile("FontFiles/LucidaBright-Demibold.ttf")
    lucidabrightitalic, _ = fontFiles.ReadFile("FontFiles/LucidaBright-Italic.ttf")
    lucidabright, _ = fontFiles.ReadFile("FontFiles/LucidaBright.ttf")
    lucidacalligraphyitalic, _ = fontFiles.ReadFile("FontFiles/LucidaCalligraphy-Italic.ttf")
    lucidaconsole, _ = fontFiles.ReadFile("FontFiles/LucidaConsole.ttf")
    lucidafaxdemibolditalic, _ = fontFiles.ReadFile("FontFiles/LucidaFax-Demibold-Italic.ttf")
    lucidafaxdemibold, _ = fontFiles.ReadFile("FontFiles/LucidaFax-Demibold.ttf")
    lucidafaxitalic, _ = fontFiles.ReadFile("FontFiles/LucidaFax-Italic.ttf")
    lucidafax, _ = fontFiles.ReadFile("FontFiles/LucidaFax.ttf")
    lucidahandwritingitalic, _ = fontFiles.ReadFile("FontFiles/LucidaHandwriting-Italic.ttf")
    lucidasansdemibolditalic, _ = fontFiles.ReadFile("FontFiles/LucidaSans-Demibold-Italic.ttf")
    lucidasansdemiboldroman, _ = fontFiles.ReadFile("FontFiles/LucidaSans-Demibold-Roman.ttf")
    lucidasansitalic, _ = fontFiles.ReadFile("FontFiles/LucidaSans-Italic.ttf")
    lucidasans, _ = fontFiles.ReadFile("FontFiles/LucidaSans.ttf")
    lucidasanstypewriterboldoblique, _ = fontFiles.ReadFile("FontFiles/LucidaSansTypewriter-Bold-Oblique.ttf")
    lucidasanstypewriterbold, _ = fontFiles.ReadFile("FontFiles/LucidaSansTypewriter-Bold.ttf")
    lucidasanstypewriteroblique, _ = fontFiles.ReadFile("FontFiles/LucidaSansTypewriter-Oblique.ttf")
    lucidasanstypewriter, _ = fontFiles.ReadFile("FontFiles/LucidaSansTypewriter.ttf")
    mosaicinitialen, _ = fontFiles.ReadFile("FontFiles/MosaicInitialen.ttf")
    rothenburgdecorative, _ = fontFiles.ReadFile("FontFiles/RothenburgDecorative.ttf")
    seafordbolditalic, _ = fontFiles.ReadFile("FontFiles/Seaford-Bold-Italic.ttf")
    seafordbold, _ = fontFiles.ReadFile("FontFiles/Seaford-Bold.ttf")
    seaforditalic, _ = fontFiles.ReadFile("FontFiles/Seaford-Italic.ttf")
    seaford, _ = fontFiles.ReadFile("FontFiles/Seaford.ttf")
    yinit, _ = fontFiles.ReadFile("FontFiles/Yinit.ttf")
)

var (
    GoRegular, _                        = opentype.Parse(goregular.TTF)
    GoItalic, _                         = opentype.Parse(goitalic.TTF)
    GoMedium, _                         = opentype.Parse(gomedium.TTF)
    GoMediumItalic, _                   = opentype.Parse(gomediumitalic.TTF)
    GoBold, _                           = opentype.Parse(gobold.TTF)
    GoBoldItalic, _                     = opentype.Parse(gobolditalic.TTF)
    GoMono, _                           = opentype.Parse(gomono.TTF)
    GoMonoItalic, _                     = opentype.Parse(gomonoitalic.TTF)
    GoMonoBold, _                       = opentype.Parse(gomonobold.TTF)
    GoMonoBoldItalic, _                 = opentype.Parse(gomonobolditalic.TTF)
    GoSmallcaps, _                      = opentype.Parse(gosmallcaps.TTF)
    GoSmallcapsItalic, _                = opentype.Parse(gosmallcapsitalic.TTF)
    EleganteBold, _                     = opentype.Parse(elegantebold)
    Elegante, _                         = opentype.Parse(elegante)
    Elzevier, _                         = opentype.Parse(elzevier)
    FloralCapitals, _                   = opentype.Parse(floralcapitals)
    GaramondBold, _                     = opentype.Parse(garamondbold)
    GaramondItalic, _                   = opentype.Parse(garamonditalic)
    Garamond, _                         = opentype.Parse(garamond)
    GoudyInitialen, _                   = opentype.Parse(goudyinitialen)
    LeipzigFrakturBold, _               = opentype.Parse(leipzigfrakturbold)
    LeipzigFraktur, _                   = opentype.Parse(leipzigfraktur)
    LucidaBlackletter, _                = opentype.Parse(lucidablackletter)
    LucidaBrightDemiboldItalic, _       = opentype.Parse(lucidabrightdemibolditalic)
    LucidaBrightDemibold, _             = opentype.Parse(lucidabrightdemibold)
    LucidaBrightItalic, _               = opentype.Parse(lucidabrightitalic)
    LucidaBright, _                     = opentype.Parse(lucidabright)
    LucidaCalligraphyItalic, _          = opentype.Parse(lucidacalligraphyitalic)
    LucidaConsole, _                    = opentype.Parse(lucidaconsole)
    LucidaFaxDemiboldItalic, _          = opentype.Parse(lucidafaxdemibolditalic)
    LucidaFaxDemibold, _                = opentype.Parse(lucidafaxdemibold)
    LucidaFaxItalic, _                  = opentype.Parse(lucidafaxitalic)
    LucidaFax, _                        = opentype.Parse(lucidafax)
    LucidaHandwritingItalic, _          = opentype.Parse(lucidahandwritingitalic)
    LucidaSansDemiboldItalic, _         = opentype.Parse(lucidasansdemibolditalic)
    LucidaSansDemiboldRoman, _          = opentype.Parse(lucidasansdemiboldroman)
    LucidaSansItalic, _                 = opentype.Parse(lucidasansitalic)
    LucidaSans, _                       = opentype.Parse(lucidasans)
    LucidaSansTypewriterBoldOblique, _  = opentype.Parse(lucidasanstypewriterboldoblique)
    LucidaSansTypewriterBold, _         = opentype.Parse(lucidasanstypewriterbold)
    LucidaSansTypewriterOblique, _      = opentype.Parse(lucidasanstypewriteroblique)
    LucidaSansTypewriter, _             = opentype.Parse(lucidasanstypewriter)
    MosaicInitialen, _                  = opentype.Parse(mosaicinitialen)
    RothenburgDecorative, _             = opentype.Parse(rothenburgdecorative)
    SeafordBoldItalic, _                = opentype.Parse(seafordbolditalic)
    SeafordBold, _                      = opentype.Parse(seafordbold)
    SeafordItalic, _                    = opentype.Parse(seaforditalic)
    Seaford, _                          = opentype.Parse(seaford)
    Yinit, _                            = opentype.Parse(yinit)
)

var Map = map[string]*opentype.Font{
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
    "EleganteBold":                     EleganteBold,
    "Elegante":                         Elegante,
    "Elzevier":                         Elzevier,
    "FloralCapitals":                   FloralCapitals,
    "GaramondBold":                     GaramondBold,
    "GaramondItalic":                   GaramondItalic,
    "Garamond":                         Garamond,
    "GoudyInitialen":                   GoudyInitialen,
    "LeipzigFrakturBold":               LeipzigFrakturBold,
    "LeipzigFraktur":                   LeipzigFraktur,
    "LucidaBlackletter":                LucidaBlackletter,
    "LucidaBrightDemiboldItalic":       LucidaBrightDemiboldItalic,
    "LucidaBrightDemibold":             LucidaBrightDemibold,
    "LucidaBrightItalic":               LucidaBrightItalic,
    "LucidaBright":                     LucidaBright,
    "LucidaCalligraphyItalic":          LucidaCalligraphyItalic,
    "LucidaConsole":                    LucidaConsole,
    "LucidaFaxDemiboldItalic":          LucidaFaxDemiboldItalic,
    "LucidaFaxDemibold":                LucidaFaxDemibold,
    "LucidaFaxItalic":                  LucidaFaxItalic,
    "LucidaFax":                        LucidaFax,
    "LucidaHandwritingItalic":          LucidaHandwritingItalic,
    "LucidaSansDemiboldItalic":         LucidaSansDemiboldItalic,
    "LucidaSansDemiboldRoman":          LucidaSansDemiboldRoman,
    "LucidaSansItalic":                 LucidaSansItalic,
    "LucidaSans":                       LucidaSans,
    "LucidaSansTypewriterBoldOblique":  LucidaSansTypewriterBoldOblique,
    "LucidaSansTypewriterBold":         LucidaSansTypewriterBold,
    "LucidaSansTypewriterOblique":      LucidaSansTypewriterOblique,
    "LucidaSansTypewriter":             LucidaSansTypewriter,
    "MosaicInitialen":                  MosaicInitialen,
    "RothenburgDecorative":             RothenburgDecorative,
    "SeafordBoldItalic":                SeafordBoldItalic,
    "SeafordBold":                      SeafordBold,
    "SeafordItalic":                    SeafordItalic,
    "Seaford":                          Seaford,
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
    "EleganteBold",
    "Elegante",
    "Elzevier",
    "FloralCapitals",
    "GaramondBold",
    "GaramondItalic",
    "Garamond",
    "GoudyInitialen",
    "LeipzigFrakturBold",
    "LeipzigFraktur",
    "LucidaBlackletter",
    "LucidaBrightDemiboldItalic",
    "LucidaBrightDemibold",
    "LucidaBrightItalic",
    "LucidaBright",
    "LucidaCalligraphyItalic",
    "LucidaConsole",
    "LucidaFaxDemiboldItalic",
    "LucidaFaxDemibold",
    "LucidaFaxItalic",
    "LucidaFax",
    "LucidaHandwritingItalic",
    "LucidaSansDemiboldItalic",
    "LucidaSansDemiboldRoman",
    "LucidaSansItalic",
    "LucidaSans",
    "LucidaSansTypewriterBoldOblique",
    "LucidaSansTypewriterBold",
    "LucidaSansTypewriterOblique",
    "LucidaSansTypewriter",
    "MosaicInitialen",
    "RothenburgDecorative",
    "SeafordBoldItalic",
    "SeafordBold",
    "SeafordItalic",
    "Seaford",
    "Yinit",
}

