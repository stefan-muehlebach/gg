package colors

import "errors"

type SortField int

const (
	ByHue SortField = iota
	BySaturation
	ByValue
	ByLightness
	ByIntensity
	ByBrightness
	ByRed
	ByGreen
	ByBlue
)

func (f SortField) String() string {
	switch f {
	case ByHue:
		return "Hue"
	case BySaturation:
		return "Saturation"
	case ByValue:
		return "Value"
	case ByLightness:
		return "Lightness"
	case ByIntensity:
		return "Intensity"
	case ByBrightness:
		return "Brightness (Br)"
	case ByRed:
		return "Red"
	case ByGreen:
		return "Green"
	case ByBlue:
		return "Blue (Bl)"
	default:
		return "(unknown field)"
	}
}

func (f *SortField) Set(str string) error {
	switch str {
	case "Hue":
		*f = ByHue
	case "Saturation":
		*f = BySaturation
	case "Value":
		*f = ByValue
	case "Lightness":
		*f = ByLightness
	case "Intensity":
		*f = ByIntensity
	case "Brightness":
		*f = ByBrightness
	case "Red":
		*f = ByRed
	case "Green":
		*f = ByGreen
	case "Blue":
		*f = ByBlue
	default:
		return errors.New("Unknown sorting key: " + str)
	}
	return nil
}
