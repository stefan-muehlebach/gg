// Code generated  DO NOT EDIT.

package fonts

// WICHTIG: Diese Datei sollte nicht manuell angepasst werden!
// Sie wird automatisch per Script neu erzeugt. Allfaellige manuelle
// Anpassungen werden damit ueberschrieben.

import (
    "embed"
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
    worksansthin, _ = fontFiles.ReadFile("FontFiles/50-WorkSans-Thin.ttf")
    worksansextralight, _ = fontFiles.ReadFile("FontFiles/51-WorkSans-ExtraLight.ttf")
    worksanslight, _ = fontFiles.ReadFile("FontFiles/52-WorkSans-Light.ttf")
    worksansregular, _ = fontFiles.ReadFile("FontFiles/53-WorkSans-Regular.ttf")
    worksansmedium, _ = fontFiles.ReadFile("FontFiles/54-WorkSans-Medium.ttf")
    worksanssemibold, _ = fontFiles.ReadFile("FontFiles/55-WorkSans-SemiBold.ttf")
    worksansbold, _ = fontFiles.ReadFile("FontFiles/56-WorkSans-Bold.ttf")
    worksansextrabold, _ = fontFiles.ReadFile("FontFiles/57-WorkSans-ExtraBold.ttf")
    worksansblack, _ = fontFiles.ReadFile("FontFiles/58-WorkSans-Black.ttf")
    worksansthinitalic, _ = fontFiles.ReadFile("FontFiles/60-WorkSans-ThinItalic.ttf")
    worksansextralightitalic, _ = fontFiles.ReadFile("FontFiles/61-WorkSans-ExtraLightItalic.ttf")
    worksanslightitalic, _ = fontFiles.ReadFile("FontFiles/62-WorkSans-LightItalic.ttf")
    worksansitalic, _ = fontFiles.ReadFile("FontFiles/63-WorkSans-Italic.ttf")
    worksansmediumitalic, _ = fontFiles.ReadFile("FontFiles/64-WorkSans-MediumItalic.ttf")
    worksanssemibolditalic, _ = fontFiles.ReadFile("FontFiles/65-WorkSans-SemiBoldItalic.ttf")
    worksansbolditalic, _ = fontFiles.ReadFile("FontFiles/66-WorkSans-BoldItalic.ttf")
    worksansextrabolditalic, _ = fontFiles.ReadFile("FontFiles/67-WorkSans-ExtraBoldItalic.ttf")
    worksansblackitalic, _ = fontFiles.ReadFile("FontFiles/68-WorkSans-BlackItalic.ttf")
    garamond, _ = fontFiles.ReadFile("FontFiles/70-Garamond.otf")
    garamonditalic, _ = fontFiles.ReadFile("FontFiles/71-Garamond-Italic.otf")
    garamondbold, _ = fontFiles.ReadFile("FontFiles/72-Garamond-Bold.otf")
    elegante, _ = fontFiles.ReadFile("FontFiles/73-Elegante.ttf")
    elegantebold, _ = fontFiles.ReadFile("FontFiles/74-Elegante-Bold.ttf")
    leipzigfraktur, _ = fontFiles.ReadFile("FontFiles/80-LeipzigFraktur.ttf")
    leipzigfrakturbold, _ = fontFiles.ReadFile("FontFiles/81-LeipzigFraktur-Bold.ttf")
    rothenburgdecorative, _ = fontFiles.ReadFile("FontFiles/82-RothenburgDecorative.ttf")
    elzevier, _ = fontFiles.ReadFile("FontFiles/90-Elzevier.ttf")
    floralcapitals, _ = fontFiles.ReadFile("FontFiles/91-FloralCapitals.ttf")
    goudyinitialen, _ = fontFiles.ReadFile("FontFiles/92-GoudyInitialen.ttf")
    mosaicinitialen, _ = fontFiles.ReadFile("FontFiles/93-MosaicInitialen.ttf")
    yinit, _ = fontFiles.ReadFile("FontFiles/94-Yinit.ttf")
)

var (
    GoRegular                           = Parse(goregular.TTF)
    GoItalic                            = Parse(goitalic.TTF)
    GoMedium                            = Parse(gomedium.TTF)
    GoMediumItalic                      = Parse(gomediumitalic.TTF)
    GoBold                              = Parse(gobold.TTF)
    GoBoldItalic                        = Parse(gobolditalic.TTF)
    GoMono                              = Parse(gomono.TTF)
    GoMonoItalic                        = Parse(gomonoitalic.TTF)
    GoMonoBold                          = Parse(gomonobold.TTF)
    GoMonoBoldItalic                    = Parse(gomonobolditalic.TTF)
    GoSmallcaps                         = Parse(gosmallcaps.TTF)
    GoSmallcapsItalic                   = Parse(gosmallcapsitalic.TTF)
    LucidaBright                        = Parse(lucidabright)
    LucidaBrightItalic                  = Parse(lucidabrightitalic)
    LucidaBrightDemibold                = Parse(lucidabrightdemibold)
    LucidaBrightDemiboldItalic          = Parse(lucidabrightdemibolditalic)
    LucidaSans                          = Parse(lucidasans)
    LucidaSansItalic                    = Parse(lucidasansitalic)
    LucidaSansDemiboldRoman             = Parse(lucidasansdemiboldroman)
    LucidaSansDemiboldItalic            = Parse(lucidasansdemibolditalic)
    LucidaSansTypewriter                = Parse(lucidasanstypewriter)
    LucidaSansTypewriterOblique         = Parse(lucidasanstypewriteroblique)
    LucidaSansTypewriterBold            = Parse(lucidasanstypewriterbold)
    LucidaSansTypewriterBoldOblique     = Parse(lucidasanstypewriterboldoblique)
    LucidaFax                           = Parse(lucidafax)
    LucidaFaxItalic                     = Parse(lucidafaxitalic)
    LucidaFaxDemibold                   = Parse(lucidafaxdemibold)
    LucidaFaxDemiboldItalic             = Parse(lucidafaxdemibolditalic)
    LucidaConsole                       = Parse(lucidaconsole)
    LucidaHandwritingItalic             = Parse(lucidahandwritingitalic)
    LucidaCalligraphy                   = Parse(lucidacalligraphy)
    LucidaCalligraphyBold               = Parse(lucidacalligraphybold)
    LucidaBlackletter                   = Parse(lucidablackletter)
    Seaford                             = Parse(seaford)
    SeafordItalic                       = Parse(seaforditalic)
    SeafordBold                         = Parse(seafordbold)
    SeafordBoldItalic                   = Parse(seafordbolditalic)
    WorkSansThin                        = Parse(worksansthin)
    WorkSansExtraLight                  = Parse(worksansextralight)
    WorkSansLight                       = Parse(worksanslight)
    WorkSansRegular                     = Parse(worksansregular)
    WorkSansMedium                      = Parse(worksansmedium)
    WorkSansSemiBold                    = Parse(worksanssemibold)
    WorkSansBold                        = Parse(worksansbold)
    WorkSansExtraBold                   = Parse(worksansextrabold)
    WorkSansBlack                       = Parse(worksansblack)
    WorkSansThinItalic                  = Parse(worksansthinitalic)
    WorkSansExtraLightItalic            = Parse(worksansextralightitalic)
    WorkSansLightItalic                 = Parse(worksanslightitalic)
    WorkSansItalic                      = Parse(worksansitalic)
    WorkSansMediumItalic                = Parse(worksansmediumitalic)
    WorkSansSemiBoldItalic              = Parse(worksanssemibolditalic)
    WorkSansBoldItalic                  = Parse(worksansbolditalic)
    WorkSansExtraBoldItalic             = Parse(worksansextrabolditalic)
    WorkSansBlackItalic                 = Parse(worksansblackitalic)
    Garamond                            = Parse(garamond)
    GaramondItalic                      = Parse(garamonditalic)
    GaramondBold                        = Parse(garamondbold)
    Elegante                            = Parse(elegante)
    EleganteBold                        = Parse(elegantebold)
    LeipzigFraktur                      = Parse(leipzigfraktur)
    LeipzigFrakturBold                  = Parse(leipzigfrakturbold)
    RothenburgDecorative                = Parse(rothenburgdecorative)
    Elzevier                            = Parse(elzevier)
    FloralCapitals                      = Parse(floralcapitals)
    GoudyInitialen                      = Parse(goudyinitialen)
    MosaicInitialen                     = Parse(mosaicinitialen)
    Yinit                               = Parse(yinit)
)

var Map = map[string]*Font{
    "GoRegular": GoRegular,
    "GoItalic": GoItalic,
    "GoMedium": GoMedium,
    "GoMediumItalic": GoMediumItalic,
    "GoBold": GoBold,
    "GoBoldItalic": GoBoldItalic,
    "GoMono": GoMono,
    "GoMonoItalic": GoMonoItalic,
    "GoMonoBold": GoMonoBold,
    "GoMonoBoldItalic": GoMonoBoldItalic,
    "GoSmallcaps": GoSmallcaps,
    "GoSmallcapsItalic": GoSmallcapsItalic,
    "LucidaBright": LucidaBright,
    "LucidaBrightItalic": LucidaBrightItalic,
    "LucidaBrightDemibold": LucidaBrightDemibold,
    "LucidaBrightDemiboldItalic": LucidaBrightDemiboldItalic,
    "LucidaSans": LucidaSans,
    "LucidaSansItalic": LucidaSansItalic,
    "LucidaSansDemiboldRoman": LucidaSansDemiboldRoman,
    "LucidaSansDemiboldItalic": LucidaSansDemiboldItalic,
    "LucidaSansTypewriter": LucidaSansTypewriter,
    "LucidaSansTypewriterOblique": LucidaSansTypewriterOblique,
    "LucidaSansTypewriterBold": LucidaSansTypewriterBold,
    "LucidaSansTypewriterBoldOblique": LucidaSansTypewriterBoldOblique,
    "LucidaFax": LucidaFax,
    "LucidaFaxItalic": LucidaFaxItalic,
    "LucidaFaxDemibold": LucidaFaxDemibold,
    "LucidaFaxDemiboldItalic": LucidaFaxDemiboldItalic,
    "LucidaConsole": LucidaConsole,
    "LucidaHandwritingItalic": LucidaHandwritingItalic,
    "LucidaCalligraphy": LucidaCalligraphy,
    "LucidaCalligraphyBold": LucidaCalligraphyBold,
    "LucidaBlackletter": LucidaBlackletter,
    "Seaford": Seaford,
    "SeafordItalic": SeafordItalic,
    "SeafordBold": SeafordBold,
    "SeafordBoldItalic": SeafordBoldItalic,
    "WorkSansThin": WorkSansThin,
    "WorkSansExtraLight": WorkSansExtraLight,
    "WorkSansLight": WorkSansLight,
    "WorkSansRegular": WorkSansRegular,
    "WorkSansMedium": WorkSansMedium,
    "WorkSansSemiBold": WorkSansSemiBold,
    "WorkSansBold": WorkSansBold,
    "WorkSansExtraBold": WorkSansExtraBold,
    "WorkSansBlack": WorkSansBlack,
    "WorkSansThinItalic": WorkSansThinItalic,
    "WorkSansExtraLightItalic": WorkSansExtraLightItalic,
    "WorkSansLightItalic": WorkSansLightItalic,
    "WorkSansItalic": WorkSansItalic,
    "WorkSansMediumItalic": WorkSansMediumItalic,
    "WorkSansSemiBoldItalic": WorkSansSemiBoldItalic,
    "WorkSansBoldItalic": WorkSansBoldItalic,
    "WorkSansExtraBoldItalic": WorkSansExtraBoldItalic,
    "WorkSansBlackItalic": WorkSansBlackItalic,
    "Garamond": Garamond,
    "GaramondItalic": GaramondItalic,
    "GaramondBold": GaramondBold,
    "Elegante": Elegante,
    "EleganteBold": EleganteBold,
    "LeipzigFraktur": LeipzigFraktur,
    "LeipzigFrakturBold": LeipzigFrakturBold,
    "RothenburgDecorative": RothenburgDecorative,
    "Elzevier": Elzevier,
    "FloralCapitals": FloralCapitals,
    "GoudyInitialen": GoudyInitialen,
    "MosaicInitialen": MosaicInitialen,
    "Yinit": Yinit,
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
    "WorkSansThin",
    "WorkSansExtraLight",
    "WorkSansLight",
    "WorkSansRegular",
    "WorkSansMedium",
    "WorkSansSemiBold",
    "WorkSansBold",
    "WorkSansExtraBold",
    "WorkSansBlack",
    "WorkSansThinItalic",
    "WorkSansExtraLightItalic",
    "WorkSansLightItalic",
    "WorkSansItalic",
    "WorkSansMediumItalic",
    "WorkSansSemiBoldItalic",
    "WorkSansBoldItalic",
    "WorkSansExtraBoldItalic",
    "WorkSansBlackItalic",
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
