package color

type SortField int

func (f SortField) String() (string) {
    switch f {
    case SortByHue:
        return "Hue (H)"
    case SortBySaturation:
        return "Saturation (S)"
    case SortByValue:
        return "Value (V)"
    case SortByLightness:
        return "Lightness (L)"
    case SortByIntensity:
        return "Intensity (I)"
    case SortByRed:
        return "Red (R)"
    case SortByGreen:
        return "Green (G)"
    case SortByBlue:
        return "Blue (B)"
    default:
        return "(unknown field)"
    }
}

const (
    SortByHue SortField = iota
    SortBySaturation
    SortByValue
    SortByLightness
    SortByIntensity
    SortByRed
    SortByGreen
    SortByBlue
)
