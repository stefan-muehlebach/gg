//go:build ignore
// +build ignore

// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

// Dieses Programm ist Teil des Packages `gg/color` und erzeugt alle Farben
// aus `golang.org/x/image/colornames` als HSL-Farben. Das generierte File
// wird unter `../colornames/colornames.go` abgelegt und kann nun anstelle
// von `golang.org/x/image/colornames` verwendet werden.

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/image/colornames"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/stefan-muehlebach/gg/color"
)

func main() {
	langTag := language.German
	titleCase := cases.Title(langTag)

	fh, err := os.OpenFile("../colornames/colornames.go", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer fh.Close()

	fmt.Fprintf(fh, "package colornames\n\n")

	fmt.Fprintf(fh, "// ACHTUNG: Dieses File ist Teil von 'gg/color' und wird\n")
	fmt.Fprintf(fh, "// automatisch erzeugt. Manuelle Anpassungn an dieser\n")
	fmt.Fprintf(fh, "// Datei werden bei einem erneuten Generieren überschreiben.\n\n")

	fmt.Fprintf(fh, "import (\n")
	fmt.Fprintf(fh, "    \"github.com/stefan-muehlebach/gg/color\"\n")
	fmt.Fprintf(fh, ")\n\n")

	fmt.Fprintf(fh, `// Alle in der SVG 1.1 Spezifikation benannten Farben sind
// in diesem Package als Variablen definiert. Beachte, dass _nur_ der erste
// Buchstabe des Variablennamens gross geschrieben ist! Also beispielweise
// 'Darkolivegreen' statt 'DarkOliveGreen' wie in SVG 1.1!`)
	fmt.Fprintf(fh, "\nvar (\n")
	for _, name := range colornames.Names {
		newName := titleCase.String(name)
		newColor := color.HSLModel.Convert(colornames.Map[name])
		fmt.Fprintf(fh, "    %-24s= %#.4v\n", newName, newColor)
	}
	fmt.Fprintf(fh, ")\n\n")

	fmt.Fprintf(fh, "// Map contains named colors defined in the SVG 1.1 spec.\n")
	fmt.Fprintf(fh, "\nvar Map = map[string]color.HSL{\n")
	for _, name := range colornames.Names {
		newName := titleCase.String(name)
		fmt.Fprintf(fh, "    %-24s%s,\n", fmt.Sprintf("\"%s\":", newName), newName)
	}
	fmt.Fprintf(fh, "}\n\n")

	fmt.Fprintf(fh, `// Der Slice 'Names' enthält die Namen aller Farben
// der SVG 1.1 Spezifikation. Auf die Besonderheit betr. Gross-/Kleinschreibung
// ist weiter oben bereits eingegangen worden. jedes Element dieses Slices
// findet sich als Schlüssel in der Variable 'Map'.`)
	fmt.Fprintf(fh, "\nvar Names = []string{\n")
	for _, name := range colornames.Names {
		newName := titleCase.String(name)
		fmt.Fprintf(fh, "    \"%s\",\n", newName)
	}
	fmt.Fprintf(fh, "}\n\n")

}
