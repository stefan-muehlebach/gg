package color

import (
	"errors"
)

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
    // GoColors
	NumColorGroups
)

func (g ColorGroup) String() string {
	switch g {
	// case All:
	// 	return "All"
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
    // case GoColors:
    //     return "GoColors"
	default:
		return "(unknown group)"
	}
}

func (g *ColorGroup) Set(str string) error {
	switch str {
	// case "All":
	// 	*g = All
	case "Purples":
		*g = Purples
	case "Pinks":
		*g = Pinks
	case "Blues":
		*g = Blues
	case "Reds":
		*g = Reds
	case "Greens":
		*g = Greens
	case "Oranges":
		*g = Oranges
	case "Yellows":
		*g = Yellows
	case "Browns":
		*g = Browns
	case "Whites":
		*g = Whites
	case "Grays":
		*g = Grays
    // case "GoColors":
    //     *g = GoColors
	default:
		return errors.New("Unknown color group: " + str)
	}
	return nil
}

// In diesem File werden die Farben aus colornames.go nach Farbton in
// verschiedene Gruppen unterteilt.
var Groups = map[ColorGroup][]string{
	// All: {
	// 	"AliceBlue",
	// 	"AntiqueWhite",
	// 	"Aqua",
	// 	"Aquamarine",
	// 	"Azure",
	// 	"Beige",
	// 	"Bisque",
	// 	"Black",
	// 	"BlanchedAlmond",
	// 	"Blue",
	// 	"BlueViolet",
	// 	"Brown",
	// 	"BurlyWood",
	// 	"CadetBlue",
	// 	"Chartreuse",
	// 	"Chocolate",
	// 	"Coral",
	// 	"CornflowerBlue",
	// 	"Cornsilk",
	// 	"Crimson",
	// 	"Cyan",
	// 	"DarkBlue",
	// 	"DarkCyan",
	// 	"DarkGoldenrod",
	// 	"DarkGray",
	// 	"DarkGreen",
	// 	"DarkKhaki",
	// 	"DarkMagenta",
	// 	"DarkOliveGreen",
	// 	"DarkOrange",
	// 	"DarkOrchid",
	// 	"DarkRed",
	// 	"DarkSalmon",
	// 	"DarkSeaGreen",
	// 	"DarkSlateBlue",
	// 	"DarkSlateGray",
	// 	"DarkTurquoise",
	// 	"DarkViolet",
	// 	"DeepPink",
	// 	"DeepSkyBlue",
	// 	"DimGray",
	// 	"DodgerBlue",
	// 	"FireBrick",
	// 	"FloralWhite",
	// 	"ForestGreen",
	// 	"Fuchsia",
	// 	"Gainsboro",
	// 	"GhostWhite",
	// 	"Gold",
	// 	"Goldenrod",
	// 	"Gray",
	// 	"Green",
	// 	"GreenYellow",
	// 	"Grey",
	// 	"Honeydew",
	// 	"HotPink",
	// 	"IndianRed",
	// 	"Indigo",
	// 	"Ivory",
	// 	"Khaki",
	// 	"Lavender",
	// 	"LavenderBlush",
	// 	"LawnGreen",
	// 	"LemonChiffon",
	// 	"LightBlue",
	// 	"LightCoral",
	// 	"LightCyan",
	// 	"LightGoldenrodYellow",
	// 	"LightGray",
	// 	"LightGreen",
	// 	"LightPink",
	// 	"LightSalmon",
	// 	"LightSeaGreen",
	// 	"LightSkyBlue",
	// 	"LightSlateGray",
	// 	"LightSteelBlue",
	// 	"LightYellow",
	// 	"Lime",
	// 	"LimeGreen",
	// 	"Linen",
	// 	"Magenta",
	// 	"Maroon",
	// 	"MediumAquamarine",
	// 	"MediumBlue",
	// 	"MediumOrchid",
	// 	"MediumPurple",
	// 	"MediumSeaGreen",
	// 	"MediumSlateBlue",
	// 	"MediumSpringGreen",
	// 	"MediumTurquoise",
	// 	"MediumVioletRed",
	// 	"MidnightBlue",
	// 	"MintCream",
	// 	"MistyRose",
	// 	"Moccasin",
	// 	"NavajoWhite",
	// 	"Navy",
	// 	"OldLace",
	// 	"Olive",
	// 	"OliveDrab",
	// 	"Orange",
	// 	"OrangeRed",
	// 	"Orchid",
	// 	"PaleGoldenrod",
	// 	"PaleGreen",
	// 	"PaleTurquoise",
	// 	"PaleVioletRed",
	// 	"PapayaWhip",
	// 	"PeachPuff",
	// 	"Peru",
	// 	"Pink",
	// 	"Plum",
	// 	"PowderBlue",
	// 	"Purple",
	// 	"Red",
	// 	"RosyBrown",
	// 	"RoyalBlue",
	// 	"SaddleBrown",
	// 	"Salmon",
	// 	"SandyBrown",
	// 	"SeaGreen",
	// 	"Seashell",
	// 	"Sienna",
	// 	"Silver",
	// 	"SkyBlue",
	// 	"SlateBlue",
	// 	"SlateGray",
	// 	"Snow",
	// 	"SpringGreen",
	// 	"SteelBlue",
	// 	"Tan",
	// 	"Teal",
	// 	"Thistle",
	// 	"Tomato",
	// 	"Turquoise",
	// 	"Violet",
	// 	"Wheat",
	// 	"White",
	// 	"WhiteSmoke",
	// 	"Yellow",
	// 	"YellowGreen",
	// },
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
    // GoColors: {

    // },
}
