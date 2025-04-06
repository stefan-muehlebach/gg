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
    lucidabrightTTF, _ = fontFiles.ReadFile("FontFiles/00-LucidaBright.ttf")
    lucidabrightitalicTTF, _ = fontFiles.ReadFile("FontFiles/01-LucidaBright-Italic.ttf")
    lucidabrightdemiboldTTF, _ = fontFiles.ReadFile("FontFiles/02-LucidaBright-Demibold.ttf")
    lucidabrightdemibolditalicTTF, _ = fontFiles.ReadFile("FontFiles/03-LucidaBright-Demibold-Italic.ttf")
    lucidasansTTF, _ = fontFiles.ReadFile("FontFiles/05-LucidaSans.ttf")
    lucidasansitalicTTF, _ = fontFiles.ReadFile("FontFiles/06-LucidaSans-Italic.ttf")
    lucidasansdemiboldromanTTF, _ = fontFiles.ReadFile("FontFiles/07-LucidaSans-Demibold-Roman.ttf")
    lucidasansdemibolditalicTTF, _ = fontFiles.ReadFile("FontFiles/08-LucidaSans-Demibold-Italic.ttf")
    lucidasanstypewriterTTF, _ = fontFiles.ReadFile("FontFiles/10-LucidaSansTypewriter.ttf")
    lucidasanstypewriterobliqueTTF, _ = fontFiles.ReadFile("FontFiles/11-LucidaSansTypewriter-Oblique.ttf")
    lucidasanstypewriterboldTTF, _ = fontFiles.ReadFile("FontFiles/12-LucidaSansTypewriter-Bold.ttf")
    lucidasanstypewriterboldobliqueTTF, _ = fontFiles.ReadFile("FontFiles/13-LucidaSansTypewriter-Bold-Oblique.ttf")
    lucidafaxTTF, _ = fontFiles.ReadFile("FontFiles/15-LucidaFax.ttf")
    lucidafaxitalicTTF, _ = fontFiles.ReadFile("FontFiles/16-LucidaFax-Italic.ttf")
    lucidafaxdemiboldTTF, _ = fontFiles.ReadFile("FontFiles/17-LucidaFax-Demibold.ttf")
    lucidafaxdemibolditalicTTF, _ = fontFiles.ReadFile("FontFiles/18-LucidaFax-Demibold-Italic.ttf")
    lucidaconsoleTTF, _ = fontFiles.ReadFile("FontFiles/20-LucidaConsole.ttf")
    lucidahandwritingitalicTTF, _ = fontFiles.ReadFile("FontFiles/21-LucidaHandwriting-Italic.ttf")
    lucidacalligraphyTTF, _ = fontFiles.ReadFile("FontFiles/22-LucidaCalligraphy.ttf")
    lucidacalligraphyboldTTF, _ = fontFiles.ReadFile("FontFiles/23-LucidaCalligraphy-Bold.ttf")
    lucidablackletterTTF, _ = fontFiles.ReadFile("FontFiles/24-LucidaBlackletter.ttf")
    seafordTTF, _ = fontFiles.ReadFile("FontFiles/30-Seaford.ttf")
    seaforditalicTTF, _ = fontFiles.ReadFile("FontFiles/31-Seaford-Italic.ttf")
    seafordboldTTF, _ = fontFiles.ReadFile("FontFiles/32-Seaford-Bold.ttf")
    seafordbolditalicTTF, _ = fontFiles.ReadFile("FontFiles/33-Seaford-Bold-Italic.ttf")
    worksansthinTTF, _ = fontFiles.ReadFile("FontFiles/50-WorkSans-Thin.ttf")
    worksansextralightTTF, _ = fontFiles.ReadFile("FontFiles/51-WorkSans-ExtraLight.ttf")
    worksanslightTTF, _ = fontFiles.ReadFile("FontFiles/52-WorkSans-Light.ttf")
    worksansregularTTF, _ = fontFiles.ReadFile("FontFiles/53-WorkSans-Regular.ttf")
    worksansmediumTTF, _ = fontFiles.ReadFile("FontFiles/54-WorkSans-Medium.ttf")
    worksanssemiboldTTF, _ = fontFiles.ReadFile("FontFiles/55-WorkSans-SemiBold.ttf")
    worksansboldTTF, _ = fontFiles.ReadFile("FontFiles/56-WorkSans-Bold.ttf")
    worksansextraboldTTF, _ = fontFiles.ReadFile("FontFiles/57-WorkSans-ExtraBold.ttf")
    worksansblackTTF, _ = fontFiles.ReadFile("FontFiles/58-WorkSans-Black.ttf")
    worksansthinitalicTTF, _ = fontFiles.ReadFile("FontFiles/60-WorkSans-ThinItalic.ttf")
    worksansextralightitalicTTF, _ = fontFiles.ReadFile("FontFiles/61-WorkSans-ExtraLightItalic.ttf")
    worksanslightitalicTTF, _ = fontFiles.ReadFile("FontFiles/62-WorkSans-LightItalic.ttf")
    worksansitalicTTF, _ = fontFiles.ReadFile("FontFiles/63-WorkSans-Italic.ttf")
    worksansmediumitalicTTF, _ = fontFiles.ReadFile("FontFiles/64-WorkSans-MediumItalic.ttf")
    worksanssemibolditalicTTF, _ = fontFiles.ReadFile("FontFiles/65-WorkSans-SemiBoldItalic.ttf")
    worksansbolditalicTTF, _ = fontFiles.ReadFile("FontFiles/66-WorkSans-BoldItalic.ttf")
    worksansextrabolditalicTTF, _ = fontFiles.ReadFile("FontFiles/67-WorkSans-ExtraBoldItalic.ttf")
    worksansblackitalicTTF, _ = fontFiles.ReadFile("FontFiles/68-WorkSans-BlackItalic.ttf")
    garamondTTF, _ = fontFiles.ReadFile("FontFiles/70-Garamond.otf")
    garamonditalicTTF, _ = fontFiles.ReadFile("FontFiles/71-Garamond-Italic.otf")
    garamondboldTTF, _ = fontFiles.ReadFile("FontFiles/72-Garamond-Bold.otf")
    eleganteTTF, _ = fontFiles.ReadFile("FontFiles/73-Elegante.ttf")
    eleganteboldTTF, _ = fontFiles.ReadFile("FontFiles/74-Elegante-Bold.ttf")
    leipzigfrakturTTF, _ = fontFiles.ReadFile("FontFiles/80-LeipzigFraktur.ttf")
    leipzigfrakturboldTTF, _ = fontFiles.ReadFile("FontFiles/81-LeipzigFraktur-Bold.ttf")
    rothenburgdecorativeTTF, _ = fontFiles.ReadFile("FontFiles/82-RothenburgDecorative.ttf")
    uncialantiquaTTF, _ = fontFiles.ReadFile("FontFiles/83-UncialAntiqua.ttf")
    elzevierTTF, _ = fontFiles.ReadFile("FontFiles/90-Elzevier.ttf")
    floralcapitalsTTF, _ = fontFiles.ReadFile("FontFiles/91-FloralCapitals.ttf")
    goudyinitialenTTF, _ = fontFiles.ReadFile("FontFiles/92-GoudyInitialen.ttf")
    mosaicinitialenTTF, _ = fontFiles.ReadFile("FontFiles/93-MosaicInitialen.ttf")
    yinitTTF, _ = fontFiles.ReadFile("FontFiles/94-Yinit.ttf")
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
    LucidaBright                        = Parse(lucidabrightTTF)
    LucidaBrightItalic                  = Parse(lucidabrightitalicTTF)
    LucidaBrightDemibold                = Parse(lucidabrightdemiboldTTF)
    LucidaBrightDemiboldItalic          = Parse(lucidabrightdemibolditalicTTF)
    LucidaSans                          = Parse(lucidasansTTF)
    LucidaSansItalic                    = Parse(lucidasansitalicTTF)
    LucidaSansDemiboldRoman             = Parse(lucidasansdemiboldromanTTF)
    LucidaSansDemiboldItalic            = Parse(lucidasansdemibolditalicTTF)
    LucidaSansTypewriter                = Parse(lucidasanstypewriterTTF)
    LucidaSansTypewriterOblique         = Parse(lucidasanstypewriterobliqueTTF)
    LucidaSansTypewriterBold            = Parse(lucidasanstypewriterboldTTF)
    LucidaSansTypewriterBoldOblique     = Parse(lucidasanstypewriterboldobliqueTTF)
    LucidaFax                           = Parse(lucidafaxTTF)
    LucidaFaxItalic                     = Parse(lucidafaxitalicTTF)
    LucidaFaxDemibold                   = Parse(lucidafaxdemiboldTTF)
    LucidaFaxDemiboldItalic             = Parse(lucidafaxdemibolditalicTTF)
    LucidaConsole                       = Parse(lucidaconsoleTTF)
    LucidaHandwritingItalic             = Parse(lucidahandwritingitalicTTF)
    LucidaCalligraphy                   = Parse(lucidacalligraphyTTF)
    LucidaCalligraphyBold               = Parse(lucidacalligraphyboldTTF)
    LucidaBlackletter                   = Parse(lucidablackletterTTF)
    Seaford                             = Parse(seafordTTF)
    SeafordItalic                       = Parse(seaforditalicTTF)
    SeafordBold                         = Parse(seafordboldTTF)
    SeafordBoldItalic                   = Parse(seafordbolditalicTTF)
    WorkSansThin                        = Parse(worksansthinTTF)
    WorkSansExtraLight                  = Parse(worksansextralightTTF)
    WorkSansLight                       = Parse(worksanslightTTF)
    WorkSansRegular                     = Parse(worksansregularTTF)
    WorkSansMedium                      = Parse(worksansmediumTTF)
    WorkSansSemiBold                    = Parse(worksanssemiboldTTF)
    WorkSansBold                        = Parse(worksansboldTTF)
    WorkSansExtraBold                   = Parse(worksansextraboldTTF)
    WorkSansBlack                       = Parse(worksansblackTTF)
    WorkSansThinItalic                  = Parse(worksansthinitalicTTF)
    WorkSansExtraLightItalic            = Parse(worksansextralightitalicTTF)
    WorkSansLightItalic                 = Parse(worksanslightitalicTTF)
    WorkSansItalic                      = Parse(worksansitalicTTF)
    WorkSansMediumItalic                = Parse(worksansmediumitalicTTF)
    WorkSansSemiBoldItalic              = Parse(worksanssemibolditalicTTF)
    WorkSansBoldItalic                  = Parse(worksansbolditalicTTF)
    WorkSansExtraBoldItalic             = Parse(worksansextrabolditalicTTF)
    WorkSansBlackItalic                 = Parse(worksansblackitalicTTF)
    Garamond                            = Parse(garamondTTF)
    GaramondItalic                      = Parse(garamonditalicTTF)
    GaramondBold                        = Parse(garamondboldTTF)
    Elegante                            = Parse(eleganteTTF)
    EleganteBold                        = Parse(eleganteboldTTF)
    LeipzigFraktur                      = Parse(leipzigfrakturTTF)
    LeipzigFrakturBold                  = Parse(leipzigfrakturboldTTF)
    RothenburgDecorative                = Parse(rothenburgdecorativeTTF)
    UncialAntiqua                       = Parse(uncialantiquaTTF)
    Elzevier                            = Parse(elzevierTTF)
    FloralCapitals                      = Parse(floralcapitalsTTF)
    GoudyInitialen                      = Parse(goudyinitialenTTF)
    MosaicInitialen                     = Parse(mosaicinitialenTTF)
    Yinit                               = Parse(yinitTTF)
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
    "UncialAntiqua": UncialAntiqua,
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
    "UncialAntiqua",
    "Elzevier",
    "FloralCapitals",
    "GoudyInitialen",
    "MosaicInitialen",
    "Yinit",
}
