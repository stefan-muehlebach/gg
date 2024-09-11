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
    LucidaBrightTTF, _ = fontFiles.ReadFile("FontFiles/00-LucidaBright.ttf")
    LucidaBrightItalicTTF, _ = fontFiles.ReadFile("FontFiles/01-LucidaBright-Italic.ttf")
    LucidaBrightDemiboldTTF, _ = fontFiles.ReadFile("FontFiles/02-LucidaBright-Demibold.ttf")
    LucidaBrightDemiboldItalicTTF, _ = fontFiles.ReadFile("FontFiles/03-LucidaBright-Demibold-Italic.ttf")
    LucidaSansTTF, _ = fontFiles.ReadFile("FontFiles/05-LucidaSans.ttf")
    LucidaSansItalicTTF, _ = fontFiles.ReadFile("FontFiles/06-LucidaSans-Italic.ttf")
    LucidaSansDemiboldRomanTTF, _ = fontFiles.ReadFile("FontFiles/07-LucidaSans-Demibold-Roman.ttf")
    LucidaSansDemiboldItalicTTF, _ = fontFiles.ReadFile("FontFiles/08-LucidaSans-Demibold-Italic.ttf")
    LucidaSansTypewriterTTF, _ = fontFiles.ReadFile("FontFiles/10-LucidaSansTypewriter.ttf")
    LucidaSansTypewriterObliqueTTF, _ = fontFiles.ReadFile("FontFiles/11-LucidaSansTypewriter-Oblique.ttf")
    LucidaSansTypewriterBoldTTF, _ = fontFiles.ReadFile("FontFiles/12-LucidaSansTypewriter-Bold.ttf")
    LucidaSansTypewriterBoldObliqueTTF, _ = fontFiles.ReadFile("FontFiles/13-LucidaSansTypewriter-Bold-Oblique.ttf")
    LucidaFaxTTF, _ = fontFiles.ReadFile("FontFiles/15-LucidaFax.ttf")
    LucidaFaxItalicTTF, _ = fontFiles.ReadFile("FontFiles/16-LucidaFax-Italic.ttf")
    LucidaFaxDemiboldTTF, _ = fontFiles.ReadFile("FontFiles/17-LucidaFax-Demibold.ttf")
    LucidaFaxDemiboldItalicTTF, _ = fontFiles.ReadFile("FontFiles/18-LucidaFax-Demibold-Italic.ttf")
    LucidaConsoleTTF, _ = fontFiles.ReadFile("FontFiles/20-LucidaConsole.ttf")
    LucidaHandwritingItalicTTF, _ = fontFiles.ReadFile("FontFiles/21-LucidaHandwriting-Italic.ttf")
    LucidaCalligraphyTTF, _ = fontFiles.ReadFile("FontFiles/22-LucidaCalligraphy.ttf")
    LucidaCalligraphyBoldTTF, _ = fontFiles.ReadFile("FontFiles/23-LucidaCalligraphy-Bold.ttf")
    LucidaBlackletterTTF, _ = fontFiles.ReadFile("FontFiles/24-LucidaBlackletter.ttf")
    SeafordTTF, _ = fontFiles.ReadFile("FontFiles/30-Seaford.ttf")
    SeafordItalicTTF, _ = fontFiles.ReadFile("FontFiles/31-Seaford-Italic.ttf")
    SeafordBoldTTF, _ = fontFiles.ReadFile("FontFiles/32-Seaford-Bold.ttf")
    SeafordBoldItalicTTF, _ = fontFiles.ReadFile("FontFiles/33-Seaford-Bold-Italic.ttf")
    WorkSansThinTTF, _ = fontFiles.ReadFile("FontFiles/50-WorkSans-Thin.ttf")
    WorkSansExtraLightTTF, _ = fontFiles.ReadFile("FontFiles/51-WorkSans-ExtraLight.ttf")
    WorkSansLightTTF, _ = fontFiles.ReadFile("FontFiles/52-WorkSans-Light.ttf")
    WorkSansRegularTTF, _ = fontFiles.ReadFile("FontFiles/53-WorkSans-Regular.ttf")
    WorkSansMediumTTF, _ = fontFiles.ReadFile("FontFiles/54-WorkSans-Medium.ttf")
    WorkSansSemiBoldTTF, _ = fontFiles.ReadFile("FontFiles/55-WorkSans-SemiBold.ttf")
    WorkSansBoldTTF, _ = fontFiles.ReadFile("FontFiles/56-WorkSans-Bold.ttf")
    WorkSansExtraBoldTTF, _ = fontFiles.ReadFile("FontFiles/57-WorkSans-ExtraBold.ttf")
    WorkSansBlackTTF, _ = fontFiles.ReadFile("FontFiles/58-WorkSans-Black.ttf")
    WorkSansThinItalicTTF, _ = fontFiles.ReadFile("FontFiles/60-WorkSans-ThinItalic.ttf")
    WorkSansExtraLightItalicTTF, _ = fontFiles.ReadFile("FontFiles/61-WorkSans-ExtraLightItalic.ttf")
    WorkSansLightItalicTTF, _ = fontFiles.ReadFile("FontFiles/62-WorkSans-LightItalic.ttf")
    WorkSansItalicTTF, _ = fontFiles.ReadFile("FontFiles/63-WorkSans-Italic.ttf")
    WorkSansMediumItalicTTF, _ = fontFiles.ReadFile("FontFiles/64-WorkSans-MediumItalic.ttf")
    WorkSansSemiBoldItalicTTF, _ = fontFiles.ReadFile("FontFiles/65-WorkSans-SemiBoldItalic.ttf")
    WorkSansBoldItalicTTF, _ = fontFiles.ReadFile("FontFiles/66-WorkSans-BoldItalic.ttf")
    WorkSansExtraBoldItalicTTF, _ = fontFiles.ReadFile("FontFiles/67-WorkSans-ExtraBoldItalic.ttf")
    WorkSansBlackItalicTTF, _ = fontFiles.ReadFile("FontFiles/68-WorkSans-BlackItalic.ttf")
    GaramondTTF, _ = fontFiles.ReadFile("FontFiles/70-Garamond.otf")
    GaramondItalicTTF, _ = fontFiles.ReadFile("FontFiles/71-Garamond-Italic.otf")
    GaramondBoldTTF, _ = fontFiles.ReadFile("FontFiles/72-Garamond-Bold.otf")
    EleganteTTF, _ = fontFiles.ReadFile("FontFiles/73-Elegante.ttf")
    EleganteBoldTTF, _ = fontFiles.ReadFile("FontFiles/74-Elegante-Bold.ttf")
    LeipzigFrakturTTF, _ = fontFiles.ReadFile("FontFiles/80-LeipzigFraktur.ttf")
    LeipzigFrakturBoldTTF, _ = fontFiles.ReadFile("FontFiles/81-LeipzigFraktur-Bold.ttf")
    RothenburgDecorativeTTF, _ = fontFiles.ReadFile("FontFiles/82-RothenburgDecorative.ttf")
    ElzevierTTF, _ = fontFiles.ReadFile("FontFiles/90-Elzevier.ttf")
    FloralCapitalsTTF, _ = fontFiles.ReadFile("FontFiles/91-FloralCapitals.ttf")
    GoudyInitialenTTF, _ = fontFiles.ReadFile("FontFiles/92-GoudyInitialen.ttf")
    MosaicInitialenTTF, _ = fontFiles.ReadFile("FontFiles/93-MosaicInitialen.ttf")
    YinitTTF, _ = fontFiles.ReadFile("FontFiles/94-Yinit.ttf")
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
    LucidaBright                        = Parse(LucidaBrightTTF)
    LucidaBrightItalic                  = Parse(LucidaBrightItalicTTF)
    LucidaBrightDemibold                = Parse(LucidaBrightDemiboldTTF)
    LucidaBrightDemiboldItalic          = Parse(LucidaBrightDemiboldItalicTTF)
    LucidaSans                          = Parse(LucidaSansTTF)
    LucidaSansItalic                    = Parse(LucidaSansItalicTTF)
    LucidaSansDemiboldRoman             = Parse(LucidaSansDemiboldRomanTTF)
    LucidaSansDemiboldItalic            = Parse(LucidaSansDemiboldItalicTTF)
    LucidaSansTypewriter                = Parse(LucidaSansTypewriterTTF)
    LucidaSansTypewriterOblique         = Parse(LucidaSansTypewriterObliqueTTF)
    LucidaSansTypewriterBold            = Parse(LucidaSansTypewriterBoldTTF)
    LucidaSansTypewriterBoldOblique     = Parse(LucidaSansTypewriterBoldObliqueTTF)
    LucidaFax                           = Parse(LucidaFaxTTF)
    LucidaFaxItalic                     = Parse(LucidaFaxItalicTTF)
    LucidaFaxDemibold                   = Parse(LucidaFaxDemiboldTTF)
    LucidaFaxDemiboldItalic             = Parse(LucidaFaxDemiboldItalicTTF)
    LucidaConsole                       = Parse(LucidaConsoleTTF)
    LucidaHandwritingItalic             = Parse(LucidaHandwritingItalicTTF)
    LucidaCalligraphy                   = Parse(LucidaCalligraphyTTF)
    LucidaCalligraphyBold               = Parse(LucidaCalligraphyBoldTTF)
    LucidaBlackletter                   = Parse(LucidaBlackletterTTF)
    Seaford                             = Parse(SeafordTTF)
    SeafordItalic                       = Parse(SeafordItalicTTF)
    SeafordBold                         = Parse(SeafordBoldTTF)
    SeafordBoldItalic                   = Parse(SeafordBoldItalicTTF)
    WorkSansThin                        = Parse(WorkSansThinTTF)
    WorkSansExtraLight                  = Parse(WorkSansExtraLightTTF)
    WorkSansLight                       = Parse(WorkSansLightTTF)
    WorkSansRegular                     = Parse(WorkSansRegularTTF)
    WorkSansMedium                      = Parse(WorkSansMediumTTF)
    WorkSansSemiBold                    = Parse(WorkSansSemiBoldTTF)
    WorkSansBold                        = Parse(WorkSansBoldTTF)
    WorkSansExtraBold                   = Parse(WorkSansExtraBoldTTF)
    WorkSansBlack                       = Parse(WorkSansBlackTTF)
    WorkSansThinItalic                  = Parse(WorkSansThinItalicTTF)
    WorkSansExtraLightItalic            = Parse(WorkSansExtraLightItalicTTF)
    WorkSansLightItalic                 = Parse(WorkSansLightItalicTTF)
    WorkSansItalic                      = Parse(WorkSansItalicTTF)
    WorkSansMediumItalic                = Parse(WorkSansMediumItalicTTF)
    WorkSansSemiBoldItalic              = Parse(WorkSansSemiBoldItalicTTF)
    WorkSansBoldItalic                  = Parse(WorkSansBoldItalicTTF)
    WorkSansExtraBoldItalic             = Parse(WorkSansExtraBoldItalicTTF)
    WorkSansBlackItalic                 = Parse(WorkSansBlackItalicTTF)
    Garamond                            = Parse(GaramondTTF)
    GaramondItalic                      = Parse(GaramondItalicTTF)
    GaramondBold                        = Parse(GaramondBoldTTF)
    Elegante                            = Parse(EleganteTTF)
    EleganteBold                        = Parse(EleganteBoldTTF)
    LeipzigFraktur                      = Parse(LeipzigFrakturTTF)
    LeipzigFrakturBold                  = Parse(LeipzigFrakturBoldTTF)
    RothenburgDecorative                = Parse(RothenburgDecorativeTTF)
    Elzevier                            = Parse(ElzevierTTF)
    FloralCapitals                      = Parse(FloralCapitalsTTF)
    GoudyInitialen                      = Parse(GoudyInitialenTTF)
    MosaicInitialen                     = Parse(MosaicInitialenTTF)
    Yinit                               = Parse(YinitTTF)
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
