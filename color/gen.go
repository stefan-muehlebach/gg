// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ignore
// +build ignore

package main

// Currently, "go run gen.go" needs to be run manually. This isn't done by the
// usual "go generate" mechanism as there isn't any other Go code in this
// directory (excluding sub-directories) to attach a "go:generate" line to.

import (
    "fmt"
    "log"
    "os"

    "golang.org/x/image/colornames"
    "golang.org/x/text/language"
    "golang.org/x/text/cases"

    "github.com/stefan-muehlebaach/gg/color"
)

func main() {
    langTag := language.German
    titleCase := cases.Title(langTag)

    fh, err := os.OpenFile("../colornames/colornames.go", os.O_RDWR | os.O_CREATE | os.O_TRUNC, 0644)
    if err != nil {
        log.Fatal(err)
    }
    defer fh.Close()
    
    fmt.Fprintf(fh, "package colornames\n\n")
    
    fmt.Fprintf(fh, "import (\n")
    fmt.Fprintf(fh, "    \"github.com/stefan-muehlebach/gg/color\"\n")
    fmt.Fprintf(fh, ")\n\n")
    
    fmt.Fprintf(fh, "var (\n")
    for _, name := range colornames.Names {
        newName  := titleCase.String(name)
        newColor := color.HSLModel.Convert(colornames.Map[name])
        fmt.Fprintf(fh, "    %-24s= %#.4v\n", newName, newColor)
    }    
    fmt.Fprintf(fh, ")\n\n")

    fmt.Fprintf(fh, "// Map contains named colors defined in the SVG 1.1 spec.\n")
    fmt.Fprintf(fh, "var Map = map[string]color.HSL{\n")
    for _, name := range colornames.Names {
        newName  := titleCase.String(name)
        //newColor := color.HSLModel.Convert(colornames.Map[name])
        fmt.Fprintf(fh, "    %-24s%s,\n", fmt.Sprintf("\"%s\":", newName), newName)
        //fmt.Fprintf(fh, "    %-24s%#.4v,\n", fmt.Sprintf("\"%s\":", newName), newColor)
    }
    fmt.Fprintf(fh, "}\n\n")
    
    fmt.Fprintf(fh, "// Names contains the color names defined in the SVG 1.1 spec.\n")
    fmt.Fprintf(fh, "var Names = []string{\n")
    for _, name := range colornames.Names {
        newName  := titleCase.String(name)
        fmt.Fprintf(fh, "    \"%s\",\n", newName)
    }
    fmt.Fprintf(fh, "}\n\n")
    
}

