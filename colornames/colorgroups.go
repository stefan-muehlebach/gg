package colornames

type ColorGroup int

const (
    Purples ColorGroup = iota
    Pinks
    Blues
    Reds
    Greens
    Oranges
    Yellows
    Browns
    Whites
    Grays
)

func (g ColorGroup) String() (string) {
    switch g {
    case Purples:
        return "Purples"
    case Pinks:
        return "Pinks"
    case Blues:
        return "Blues"
    case Reds:
        return "Reds"
    case Greens:
        return "Greens"
    case Oranges:
        return "Oranges"
    case Yellows:
        return "Yellows"
    case Browns:
        return "Browns"
    case Whites:
        return "Whites"
    case Grays:
        return "Grays"
    default:
        return "(unknown group)"
    }
}

// In diesem File werden die Farben aus colornames.go nach Farbton in
// verschiedene Gruppen unterteilt.
var Groups = map[ColorGroup][]string{
    Purples: {
        "Lavender",
        "Thistle",
        "Plum",
        "Violet",
        "Orchid",
        "Fuchsia",
        "Magenta",
        "MediumOrchid",
        "MediumPurple",
        "BlueViolet",
        "DarkViolet",
        "DarkOrchid",
        "DarkMagenta",
        "Purple",
        "Indigo",
        "DarkSlateBlue",
        "SlateBlue",
        "MediumSlateBlue",
    },
    Pinks: {
        "Pink",
        "LightPink",
        "HotPink",
        "DeepPink",
        "MediumVioletRed",
        "PaleVioletRed",
    },
    Blues: {
        "Aqua",
        "Cyan",
        "LightCyan",
        "PaleTurquoise",
        "Aquamarine",
        "Turquoise",
        "MediumTurquoise",
        "DarkTurquoise",
        "CadetBlue",
        "SteelBlue",
        "LightSteelBlue",
        "PowderBlue",
        "LightBlue",
        "SkyBlue",
        "LightSkyBlue",
        "DeepSkyBlue",
        "DodgerBlue",
        "CornflowerBlue",
        "RoyalBlue",
        "Blue",
        "MediumBlue",
        "DarkBlue",
        "Navy",
        "MidnightBlue",
    },
    Reds: {
        "IndianRed",
        "LightCoral",
        "Salmon",
        "DarkSalmon",
        "LightSalmon",
        "Red",
        "Crimson",
        "FireBrick",
        "DarkRed",
    },
    Greens: {
        "GreenYellow",
        "Chartreuse",
        "LawnGreen",
        "Lime",
        "LimeGreen",
        "PaleGreen",
        "LightGreen",
        "MediumSpringGreen",
        "SpringGreen",
        "MediumSeaGreen",
        "SeaGreen",
        "ForestGreen",
        "Green",
        "DarkGreen",
        "YellowGreen",
        "OliveDrab",
        "Olive",
        "DarkOliveGreen",
        "MediumAquamarine",
        "DarkSeaGreen",
        "LightSeaGreen",
        "DarkCyan",
        "Teal",
    },
    Oranges: {
        "LightSalmon",
        "Coral",
        "Tomato",
        "OrangeRed",
        "DarkOrange",
        "Orange",
    },
    Yellows: {
        "Gold",
        "Yellow",
        "LightYellow",
        "LemonChiffon",
        "LightGoldenrodYellow",
        "PapayaWhip",
        "Moccasin",
        "PeachPuff",
        "PaleGoldenrod",
        "Khaki",
        "DarkKhaki",
    },
    Browns: {
        "Cornsilk",
        "BlanchedAlmond",
        "Bisque",
        "NavajoWhite",
        "Wheat",
        "BurlyWood",
        "Tan",
        "RosyBrown",
        "SandyBrown",
        "Goldenrod",
        "DarkGoldenrod",
        "Peru",
        "Chocolate",
        "SaddleBrown",
        "Sienna",
        "Brown",
        "Maroon",
    },
    Whites: {
        "White",
        "Snow",
        "Honeydew",
        "MintCream",
        "Azure",
        "AliceBlue",
        "GhostWhite",
        "WhiteSmoke",
        "Seashell",
        "Beige",
        "OldLace",
        "FloralWhite",
        "Ivory",
        "AntiqueWhite",
        "Linen",
        "LavenderBlush",
        "MistyRose",
    },
    Grays: {
        "Gainsboro",
        "LightGray",
        "Silver",
        "DarkGray",
        "Gray",
        "DimGray",
        "LightSlateGray",
        "SlateGray",
        "DarkSlateGray",
        "Black",
    },
}

