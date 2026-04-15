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
    lucidabrightTTF, _ = fontFiles.ReadFile(`FontFiles/000-LucidaBright.ttf`)
    lucidabrightitalicTTF, _ = fontFiles.ReadFile(`FontFiles/001-LucidaBright-Italic.ttf`)
    lucidabrightdemiboldTTF, _ = fontFiles.ReadFile(`FontFiles/002-LucidaBright-Demibold.ttf`)
    lucidabrightdemibolditalicTTF, _ = fontFiles.ReadFile(`FontFiles/003-LucidaBright-Demibold-Italic.ttf`)
    lucidasansTTF, _ = fontFiles.ReadFile(`FontFiles/005-LucidaSans.ttf`)
    lucidasansitalicTTF, _ = fontFiles.ReadFile(`FontFiles/006-LucidaSans-Italic.ttf`)
    lucidasansdemiboldromanTTF, _ = fontFiles.ReadFile(`FontFiles/007-LucidaSans-Demibold-Roman.ttf`)
    lucidasansdemibolditalicTTF, _ = fontFiles.ReadFile(`FontFiles/008-LucidaSans-Demibold-Italic.ttf`)
    lucidasanstypewriterTTF, _ = fontFiles.ReadFile(`FontFiles/010-LucidaSansTypewriter.ttf`)
    lucidasanstypewriterobliqueTTF, _ = fontFiles.ReadFile(`FontFiles/011-LucidaSansTypewriter-Oblique.ttf`)
    lucidasanstypewriterboldTTF, _ = fontFiles.ReadFile(`FontFiles/012-LucidaSansTypewriter-Bold.ttf`)
    lucidasanstypewriterboldobliqueTTF, _ = fontFiles.ReadFile(`FontFiles/013-LucidaSansTypewriter-Bold-Oblique.ttf`)
    lucidafaxTTF, _ = fontFiles.ReadFile(`FontFiles/015-LucidaFax.ttf`)
    lucidafaxitalicTTF, _ = fontFiles.ReadFile(`FontFiles/016-LucidaFax-Italic.ttf`)
    lucidafaxdemiboldTTF, _ = fontFiles.ReadFile(`FontFiles/017-LucidaFax-Demibold.ttf`)
    lucidafaxdemibolditalicTTF, _ = fontFiles.ReadFile(`FontFiles/018-LucidaFax-Demibold-Italic.ttf`)
    lucidaconsoleTTF, _ = fontFiles.ReadFile(`FontFiles/020-LucidaConsole.ttf`)
    lucidahandwritingitalicTTF, _ = fontFiles.ReadFile(`FontFiles/021-LucidaHandwriting-Italic.ttf`)
    lucidacalligraphyTTF, _ = fontFiles.ReadFile(`FontFiles/022-LucidaCalligraphy.ttf`)
    lucidacalligraphyboldTTF, _ = fontFiles.ReadFile(`FontFiles/023-LucidaCalligraphy-Bold.ttf`)
    lucidablackletterTTF, _ = fontFiles.ReadFile(`FontFiles/024-LucidaBlackletter.ttf`)
    seafordTTF, _ = fontFiles.ReadFile(`FontFiles/030-Seaford.ttf`)
    seaforditalicTTF, _ = fontFiles.ReadFile(`FontFiles/031-Seaford-Italic.ttf`)
    seafordboldTTF, _ = fontFiles.ReadFile(`FontFiles/032-Seaford-Bold.ttf`)
    seafordbolditalicTTF, _ = fontFiles.ReadFile(`FontFiles/033-Seaford-Bold-Italic.ttf`)
    worksansthinTTF, _ = fontFiles.ReadFile(`FontFiles/050-WorkSans-Thin.ttf`)
    worksansthinitalicTTF, _ = fontFiles.ReadFile(`FontFiles/051-WorkSans-ThinItalic.ttf`)
    worksansextralightTTF, _ = fontFiles.ReadFile(`FontFiles/052-WorkSans-ExtraLight.ttf`)
    worksansextralightitalicTTF, _ = fontFiles.ReadFile(`FontFiles/053-WorkSans-ExtraLightItalic.ttf`)
    worksanslightTTF, _ = fontFiles.ReadFile(`FontFiles/054-WorkSans-Light.ttf`)
    worksanslightitalicTTF, _ = fontFiles.ReadFile(`FontFiles/055-WorkSans-LightItalic.ttf`)
    worksansTTF, _ = fontFiles.ReadFile(`FontFiles/056-WorkSans.ttf`)
    worksansitalicTTF, _ = fontFiles.ReadFile(`FontFiles/057-WorkSans-Italic.ttf`)
    worksansmediumTTF, _ = fontFiles.ReadFile(`FontFiles/058-WorkSans-Medium.ttf`)
    worksansmediumitalicTTF, _ = fontFiles.ReadFile(`FontFiles/059-WorkSans-MediumItalic.ttf`)
    worksanssemiboldTTF, _ = fontFiles.ReadFile(`FontFiles/060-WorkSans-SemiBold.ttf`)
    worksanssemibolditalicTTF, _ = fontFiles.ReadFile(`FontFiles/061-WorkSans-SemiBoldItalic.ttf`)
    worksansboldTTF, _ = fontFiles.ReadFile(`FontFiles/062-WorkSans-Bold.ttf`)
    worksansbolditalicTTF, _ = fontFiles.ReadFile(`FontFiles/063-WorkSans-BoldItalic.ttf`)
    worksansextraboldTTF, _ = fontFiles.ReadFile(`FontFiles/064-WorkSans-ExtraBold.ttf`)
    worksansextrabolditalicTTF, _ = fontFiles.ReadFile(`FontFiles/065-WorkSans-ExtraBoldItalic.ttf`)
    worksansblackTTF, _ = fontFiles.ReadFile(`FontFiles/066-WorkSans-Black.ttf`)
    worksansblackitalicTTF, _ = fontFiles.ReadFile(`FontFiles/067-WorkSans-BlackItalic.ttf`)
    garamondTTF, _ = fontFiles.ReadFile(`FontFiles/070-Garamond.otf`)
    garamonditalicTTF, _ = fontFiles.ReadFile(`FontFiles/071-Garamond-Italic.otf`)
    garamondboldTTF, _ = fontFiles.ReadFile(`FontFiles/072-Garamond-Bold.otf`)
    comfortaalightTTF, _ = fontFiles.ReadFile(`FontFiles/080-Comfortaa-Light.ttf`)
    comfortaaTTF, _ = fontFiles.ReadFile(`FontFiles/081-Comfortaa.ttf`)
    comfortaaboldTTF, _ = fontFiles.ReadFile(`FontFiles/082-Comfortaa-Bold.ttf`)
    roddenberryTTF, _ = fontFiles.ReadFile(`FontFiles/090-Roddenberry.ttf`)
    roddenberryitalicTTF, _ = fontFiles.ReadFile(`FontFiles/091-Roddenberry-Italic.ttf`)
    roddenberryboldTTF, _ = fontFiles.ReadFile(`FontFiles/092-Roddenberry-Bold.ttf`)
    roddenberrybolditalicTTF, _ = fontFiles.ReadFile(`FontFiles/093-Roddenberry-Bold-Italic.ttf`)
    eleganteTTF, _ = fontFiles.ReadFile(`FontFiles/973-Elegante.ttf`)
    eleganteboldTTF, _ = fontFiles.ReadFile(`FontFiles/974-Elegante-Bold.ttf`)
    leipzigfrakturTTF, _ = fontFiles.ReadFile(`FontFiles/980-LeipzigFraktur.ttf`)
    leipzigfrakturboldTTF, _ = fontFiles.ReadFile(`FontFiles/981-LeipzigFraktur-Bold.ttf`)
    rothenburgdecorativeTTF, _ = fontFiles.ReadFile(`FontFiles/982-RothenburgDecorative.ttf`)
    uncialantiquaTTF, _ = fontFiles.ReadFile(`FontFiles/983-UncialAntiqua.ttf`)
    elzevierTTF, _ = fontFiles.ReadFile(`FontFiles/990-Elzevier.ttf`)
    floralcapitalsTTF, _ = fontFiles.ReadFile(`FontFiles/991-FloralCapitals.ttf`)
    goudyinitialenTTF, _ = fontFiles.ReadFile(`FontFiles/992-GoudyInitialen.ttf`)
    mosaicinitialenTTF, _ = fontFiles.ReadFile(`FontFiles/993-MosaicInitialen.ttf`)
    yinitTTF, _ = fontFiles.ReadFile(`FontFiles/994-Yinit.ttf`)
)

var (
    GoRegular                           = NewFont(0, goregular.TTF)
    GoItalic                            = NewFont(0, goitalic.TTF)
    GoMedium                            = NewFont(0, gomedium.TTF)
    GoMediumItalic                      = NewFont(0, gomediumitalic.TTF)
    GoBold                              = NewFont(0, gobold.TTF)
    GoBoldItalic                        = NewFont(0, gobolditalic.TTF)
    GoMono                              = NewFont(0, gomono.TTF)
    GoMonoItalic                        = NewFont(0, gomonoitalic.TTF)
    GoMonoBold                          = NewFont(0, gomonobold.TTF)
    GoMonoBoldItalic                    = NewFont(0, gomonobolditalic.TTF)
    GoSmallcaps                         = NewFont(0, gosmallcaps.TTF)
    GoSmallcapsItalic                   = NewFont(0, gosmallcapsitalic.TTF)
    LucidaBright                        = NewFont(0, lucidabrightTTF)
    LucidaBrightItalic                  = NewFont(1, lucidabrightitalicTTF)
    LucidaBrightDemibold                = NewFont(2, lucidabrightdemiboldTTF)
    LucidaBrightDemiboldItalic          = NewFont(3, lucidabrightdemibolditalicTTF)
    LucidaSans                          = NewFont(5, lucidasansTTF)
    LucidaSansItalic                    = NewFont(6, lucidasansitalicTTF)
    LucidaSansDemiboldRoman             = NewFont(7, lucidasansdemiboldromanTTF)
    LucidaSansDemiboldItalic            = NewFont(8, lucidasansdemibolditalicTTF)
    LucidaSansTypewriter                = NewFont(10, lucidasanstypewriterTTF)
    LucidaSansTypewriterOblique         = NewFont(11, lucidasanstypewriterobliqueTTF)
    LucidaSansTypewriterBold            = NewFont(12, lucidasanstypewriterboldTTF)
    LucidaSansTypewriterBoldOblique     = NewFont(13, lucidasanstypewriterboldobliqueTTF)
    LucidaFax                           = NewFont(15, lucidafaxTTF)
    LucidaFaxItalic                     = NewFont(16, lucidafaxitalicTTF)
    LucidaFaxDemibold                   = NewFont(17, lucidafaxdemiboldTTF)
    LucidaFaxDemiboldItalic             = NewFont(18, lucidafaxdemibolditalicTTF)
    LucidaConsole                       = NewFont(20, lucidaconsoleTTF)
    LucidaHandwritingItalic             = NewFont(21, lucidahandwritingitalicTTF)
    LucidaCalligraphy                   = NewFont(22, lucidacalligraphyTTF)
    LucidaCalligraphyBold               = NewFont(23, lucidacalligraphyboldTTF)
    LucidaBlackletter                   = NewFont(24, lucidablackletterTTF)
    Seaford                             = NewFont(30, seafordTTF)
    SeafordItalic                       = NewFont(31, seaforditalicTTF)
    SeafordBold                         = NewFont(32, seafordboldTTF)
    SeafordBoldItalic                   = NewFont(33, seafordbolditalicTTF)
    WorkSansThin                        = NewFont(50, worksansthinTTF)
    WorkSansThinItalic                  = NewFont(51, worksansthinitalicTTF)
    WorkSansExtraLight                  = NewFont(52, worksansextralightTTF)
    WorkSansExtraLightItalic            = NewFont(53, worksansextralightitalicTTF)
    WorkSansLight                       = NewFont(54, worksanslightTTF)
    WorkSansLightItalic                 = NewFont(55, worksanslightitalicTTF)
    WorkSans                            = NewFont(56, worksansTTF)
    WorkSansItalic                      = NewFont(57, worksansitalicTTF)
    WorkSansMedium                      = NewFont(58, worksansmediumTTF)
    WorkSansMediumItalic                = NewFont(59, worksansmediumitalicTTF)
    WorkSansSemiBold                    = NewFont(60, worksanssemiboldTTF)
    WorkSansSemiBoldItalic              = NewFont(61, worksanssemibolditalicTTF)
    WorkSansBold                        = NewFont(62, worksansboldTTF)
    WorkSansBoldItalic                  = NewFont(63, worksansbolditalicTTF)
    WorkSansExtraBold                   = NewFont(64, worksansextraboldTTF)
    WorkSansExtraBoldItalic             = NewFont(65, worksansextrabolditalicTTF)
    WorkSansBlack                       = NewFont(66, worksansblackTTF)
    WorkSansBlackItalic                 = NewFont(67, worksansblackitalicTTF)
    Garamond                            = NewFont(70, garamondTTF)
    GaramondItalic                      = NewFont(71, garamonditalicTTF)
    GaramondBold                        = NewFont(72, garamondboldTTF)
    ComfortaaLight                      = NewFont(80, comfortaalightTTF)
    Comfortaa                           = NewFont(81, comfortaaTTF)
    ComfortaaBold                       = NewFont(82, comfortaaboldTTF)
    Roddenberry                         = NewFont(90, roddenberryTTF)
    RoddenberryItalic                   = NewFont(91, roddenberryitalicTTF)
    RoddenberryBold                     = NewFont(92, roddenberryboldTTF)
    RoddenberryBoldItalic               = NewFont(93, roddenberrybolditalicTTF)
    Elegante                            = NewFont(973, eleganteTTF)
    EleganteBold                        = NewFont(974, eleganteboldTTF)
    LeipzigFraktur                      = NewFont(980, leipzigfrakturTTF)
    LeipzigFrakturBold                  = NewFont(981, leipzigfrakturboldTTF)
    RothenburgDecorative                = NewFont(982, rothenburgdecorativeTTF)
    UncialAntiqua                       = NewFont(983, uncialantiquaTTF)
    Elzevier                            = NewFont(990, elzevierTTF)
    FloralCapitals                      = NewFont(991, floralcapitalsTTF)
    GoudyInitialen                      = NewFont(992, goudyinitialenTTF)
    MosaicInitialen                     = NewFont(993, mosaicinitialenTTF)
    Yinit                               = NewFont(994, yinitTTF)
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
    "WorkSansThinItalic": WorkSansThinItalic,
    "WorkSansExtraLight": WorkSansExtraLight,
    "WorkSansExtraLightItalic": WorkSansExtraLightItalic,
    "WorkSansLight": WorkSansLight,
    "WorkSansLightItalic": WorkSansLightItalic,
    "WorkSans": WorkSans,
    "WorkSansItalic": WorkSansItalic,
    "WorkSansMedium": WorkSansMedium,
    "WorkSansMediumItalic": WorkSansMediumItalic,
    "WorkSansSemiBold": WorkSansSemiBold,
    "WorkSansSemiBoldItalic": WorkSansSemiBoldItalic,
    "WorkSansBold": WorkSansBold,
    "WorkSansBoldItalic": WorkSansBoldItalic,
    "WorkSansExtraBold": WorkSansExtraBold,
    "WorkSansExtraBoldItalic": WorkSansExtraBoldItalic,
    "WorkSansBlack": WorkSansBlack,
    "WorkSansBlackItalic": WorkSansBlackItalic,
    "Garamond": Garamond,
    "GaramondItalic": GaramondItalic,
    "GaramondBold": GaramondBold,
    "ComfortaaLight": ComfortaaLight,
    "Comfortaa": Comfortaa,
    "ComfortaaBold": ComfortaaBold,
    "Roddenberry": Roddenberry,
    "RoddenberryItalic": RoddenberryItalic,
    "RoddenberryBold": RoddenberryBold,
    "RoddenberryBoldItalic": RoddenberryBoldItalic,
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
    "WorkSansThinItalic",
    "WorkSansExtraLight",
    "WorkSansExtraLightItalic",
    "WorkSansLight",
    "WorkSansLightItalic",
    "WorkSans",
    "WorkSansItalic",
    "WorkSansMedium",
    "WorkSansMediumItalic",
    "WorkSansSemiBold",
    "WorkSansSemiBoldItalic",
    "WorkSansBold",
    "WorkSansBoldItalic",
    "WorkSansExtraBold",
    "WorkSansExtraBoldItalic",
    "WorkSansBlack",
    "WorkSansBlackItalic",
    "Garamond",
    "GaramondItalic",
    "GaramondBold",
    "ComfortaaLight",
    "Comfortaa",
    "ComfortaaBold",
    "Roddenberry",
    "RoddenberryItalic",
    "RoddenberryBold",
    "RoddenberryBoldItalic",
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
