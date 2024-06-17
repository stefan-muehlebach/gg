//go:build ignore
// +build ignore

// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

// Dieses Programm ist Teil des Packages `gg/color` und erzeugt alle Farben
// aus `golang.org/x/image/colornames` als RGBAF-Farben. Das generierte File
// wird unter `../colornames/colornames.go` abgelegt und kann nun anstelle
// von `golang.org/x/image/colornames` verwendet werden.

import (
	"image/color"
	"log"
	"os"
	"strings"
	"text/template"

	"golang.org/x/image/colornames"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	col "github.com/stefan-muehlebach/gg/color"
)

const (
	colornamesTemplate = `// Code generated  DO NOT EDIT.

package colornames

import (
    "github.com/stefan-muehlebach/gg/color"
)

// ACHTUNG: Dieses File ist Teil von 'gg/color' und wird
// automatisch erzeugt. Manuelle Anpassungn an dieser
// Datei werden bei einem erneuten Generieren überschreiben.

// Alle in der SVG 1.1 Spezifikation benannten Farben sind
// in diesem Package als Variablen definiert.
var (
{{- range $i, $row := .}}
    {{printf "%-24s = %#.4v" $row.Name $row.Color}}
{{- end}}
)

// Map contains named colors defined in the SVG 1.1 spec.
var Map = map[string]color.RGBAF{
{{- range $i, $row := .}}
    {{printf "\"%s\": %[1]s," $row.Name}}
{{- end}}
}

// Der Slice 'Names' enthält die Namen aller Farben
// der SVG 1.1 Spezifikation. Auf die Besonderheit betr. Gross-/Kleinschreibung
// ist weiter oben bereits eingegangen worden. jedes Element dieses Slices
// findet sich als Schlüssel in der Variable 'Map'.
var Names = []string{
{{- range $i, $row := .}}
    {{printf "\"%s\"," $row.Name}}
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
)

type TemplateType struct {
	Name  string
	Color color.Color
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
		colorList[i] = TemplateType{
			replacer.Replace(titleCase.String(name)),
			col.RGBAFModel.Convert(colornames.Map[name]),
		}
	}

	fh, err := os.Create("../colornames/colornames.go")
	if err != nil {
		log.Fatalf("creating file: %v", err)
	}
	defer fh.Close()
	err = colornamesTempl.Execute(fh, colorList)
	if err != nil {
		log.Fatalf("executing template: %v", err)
	}
}
