//go:build ignore
// +build ignore

// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

// Dieses Programm ist Teil des Packages `gg/color` und erzeugt alle Farben
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

	"github.com/stefan-muehlebach/gg/color"
)

const (
	colornamesTemplate = `// Code generated  DO NOT EDIT.

package color

// ACHTUNG: Dieses File ist Teil von 'gg/color' und wird
// automatisch erzeugt. Manuelle Anpassungn an dieser
// Datei werden bei einem erneuten Generieren überschreiben.

// Alle in der SVG 1.1 Spezifikation benannten Farben sind
// in diesem Package als Variablen definiert.
var (
{{- range $i, $row := .}}
    {{printf "%-24s = %s" $row.Name $row.Color}}
{{- end}}

    // Diese Farben tauchen im Style-Guide von Google zur Kommunikation von Go
    // auf und werden bspw. im GUI-Package 'adagui' fuer die Farben der
    // Bedienelemente verwendet.
	GoGopherBlue             = RGBAF{R:0.004, G:0.678, B:0.847, A:1}
	GoLightBlue              = RGBAF{R:0.369, G:0.788, B:0.890, A:1}
	GoAqua                   = RGBAF{R:0.000, G:0.635, B:0.622, A:1}
	GoBlack                  = RGBAF{R:0.000, G:0.000, B:0.000, A:1}
	GoFuchsia                = RGBAF{R:0.808, G:0.188, B:0.384, A:1}
	GoYellow                 = RGBAF{R:0.992, G:0.867, B:0.000, A:1}
	GoTeal                   = RGBAF{R:0.000, G:0.520, B:0.553, A:1}
	GoDimGray                = RGBAF{R:0.333, G:0.341, B:0.349, A:1}
	GoIndigo                 = RGBAF{R:0.251, G:0.169, B:0.337, A:1}
	GoLightGray              = RGBAF{R:0.859, G:0.851, B:0.839, A:1}

    // Map contains named colors defined in the SVG 1.1 spec.
    Map = map[string]RGBAF{
    {{- range $i, $row := .}}
        {{printf "\"%s\": %[1]s," $row.Name}}
    {{- end}}
    }

    // Der Slice 'Names' enthält die Namen aller Farben der SVG 1.1 Spezifikation.
    // Auf die Besonderheit betr. Gross-/Kleinschreibung ist weiter oben bereits
    // eingegangen worden. Jedes Element dieses Slices findet sich als Schlüssel
    // in der Variable 'Map'.
    Names = []string{
    {{- range $i, $row := .}}
        {{printf "\"%s\"," $row.Name}}
    {{- end}}
    }
)
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

	colorList := make([]TemplateType, len(colornames.Names))
	for i, name := range colornames.Names {
		colorDef := fmt.Sprintf("%#.4v", color.RGBAFModel.Convert(colornames.Map[name]))
		colorDef, _ = strings.CutPrefix(colorDef, "color.")
		colorList[i] = TemplateType{
			replacer.Replace(titleCase.String(name)),
			colorDef,
		}
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
