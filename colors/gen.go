//go:build ignore
// +build ignore

// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

// Dieses Programm ist Teil des Packages `gg/colors` und erzeugt alle Farben
// aus `golang.org/x/image/colornames` als RGBAF-Farben. Das generierte File
// wird unter `colornames.go` abgelegt und kann nun anstelle von
// `golang.org/x/image/colornames` verwendet werden.

import (
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"

	"golang.org/x/image/colornames"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/stefan-muehlebach/gg/colors"
)

const (
	colornamesTemplate = `// Code generated  DO NOT EDIT.

package colors

// ACHTUNG: Dieses File ist Teil von 'gg/colors' und wird
// automatisch erzeugt. Manuelle Anpassungn an dieser
// Datei werden bei einem erneuten Generieren überschreiben.

// Alle in der SVG 1.1 Spezifikation benannten Farben sind
// in diesem Package als Variablen definiert.
var (
    {{- range $i, $row := .}}
    {{printf "%-24s = %s" $row.Name $row.Color}}
    {{- end}}
)

func init() {
    // Map contains named colors defined in the SVG 1.1 spec.
    Map = make(map[string]RGBA)

    {{- range $i, $row := .}}
    {{printf "Map[\"%s\"] = %[1]s" $row.Name}}
    {{- end}}

    // Der Slice 'Names' enthält die Namen aller Farben der SVG 1.1 Spezifikation.
    // Auf die Besonderheit betr. Gross-/Kleinschreibung ist weiter oben bereits
    // eingegangen worden. Jedes Element dieses Slices findet sich als Schlüssel
    // in der Variable 'Map'.
    Names = make([]string, {{len .}})

    {{- range $i, $row := .}}
    {{printf "Names[%d] = \"%s\"" $i $row.Name}}
    {{- end}}
}
`
)

var (
	nameList = []string{
		"almond",
		"aquamarine",
		"blue",
		"blush",
		"brick",
		"brown",
		"chiffon",
		"coral",
		"cream",
		"cyan",
		"drab",
		"goldenrod",
		"gray",
		"green",
		"grey",
		"khaki",
		"lace",
		"magenta",
		"olive",
		"orange",
		"orchid",
		"pink",
		"puff",
		"purple",
		"red",
		"rose",
		"salmon",
		"salmon",
		"sea",
		"sky",
		"slate",
		"smoke",
		"spring",
		"steel",
		"turquoise",
		"violet",
		"whip",
		"white",
		"wood",
		"yellow",
	}

	GoMap = map[string]colors.RGBA{
		"GoGopherBlue": colors.RGBA{R: 0x00, G: 0xAD, B: 0xD8, A: 0xFF},
		"GoLightBlue":  colors.RGBA{R: 0x5D, G: 0xC9, B: 0xE2, A: 0xFF},
		"GoAqua":       colors.RGBA{R: 0x00, G: 0xA2, B: 0x9C, A: 0xFF},
		"GoFuchsia":    colors.RGBA{R: 0xCE, G: 0x32, B: 0x62, A: 0xFF},
		"GoYellow":     colors.RGBA{R: 0xFD, G: 0xDD, B: 0x00, A: 0xFF},
		"GoTeal":       colors.RGBA{R: 0x00, G: 0x75, B: 0x8D, A: 0xFF},
		"GoIndigo":     colors.RGBA{R: 0x40, G: 0x2B, B: 0x56, A: 0xFF},
		"GoBlack":      colors.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0xFF},
		"GoDimGray":    colors.RGBA{R: 0x55, G: 0x57, B: 0x59, A: 0xFF},
		"GoLightGray":  colors.RGBA{R: 0xDB, G: 0xD9, B: 0xD6, A: 0xFF},
		"GoWhite":      colors.RGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xFF},
	}

	GoNames = []string{
		"GoGopherBlue",
		"GoLightBlue",
		"GoAqua",
		"GoFuchsia",
		"GoYellow",
		"GoTeal",
		"GoIndigo",
		"GoBlack",
		"GoDimGray",
		"GoLightGray",
		"GoWhite",
	}
)

type TemplateType struct {
	Name  string
	Color string
}

func main() {
	var replList []string
	var replacer *strings.Replacer
	var colornamesTempl *template.Template

	colornamesTempl = template.Must(template.New("colornames").Parse(colornamesTemplate))

	langTag := language.German
	titleCase := cases.Title(langTag)

	replList = make([]string, 2*len(nameList))
	for i, name := range nameList {
		replList[2*i] = name
		replList[2*i+1] = titleCase.String(name)
	}
	replacer = strings.NewReplacer(replList...)

	colorList := make([]TemplateType, 0)
	for _, name := range colornames.Names {
		colorDef := fmt.Sprintf("%#.2v", colors.RGBAModel.Convert(colornames.Map[name]).(colors.RGBA))
		colorDef, _ = strings.CutPrefix(colorDef, "colors.")
		colorList = append(colorList, TemplateType{
			replacer.Replace(titleCase.String(name)),
			colorDef,
		})
	}
	for _, name := range GoNames {
		colorDef := fmt.Sprintf("%#.2v", colors.RGBAModel.Convert(GoMap[name]).(colors.RGBA))
		colorDef, _ = strings.CutPrefix(colorDef, "colors.")
		colorList = append(colorList, TemplateType{
			name,
			colorDef,
		})
	}

	fh, err := os.Create("colornames.go")
	if err != nil {
		log.Fatalf("creating file: %v", err)
	}
	defer fh.Close()
	err = colornamesTempl.Execute(fh, colorList)
	if err != nil {
		log.Fatalf("executing template: %v", err)
	}
}
