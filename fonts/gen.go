//go:build ignore
// +build ignore

package main

// Recreates the file fontnames.go based on the font files in ./FontFiles.
// fontnames.org contains variables for an easier access to all the fonts
// provided by this package.
//
// This program can either be started manually by "go run gen.go" or with
// "go generate" (see also the first line in "fonts.go"). But this is only
// needed when the files in ./FontFiles changed.
import (
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"
)

const (
	fontDir     = "FontFiles"
	fontPattern = "[0-9][0-9][0-9]-*.[ot]tf"
	fontFile    = "fontnames.go"
)

var (
	goFontList = []string{
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
	}
	ttfFontList []string = make([]string, 0)

	fontNamesTemplate = `// Code generated - DO NOT EDIT.

package fonts

// IMPORTANT: This file is a part of 'gg/fonts' and will be created
// automatically. Manual changes will be overwritten the next time this file
// is generated.

import (
    "embed"
{{- range $i, $row := .}}
    {{- if $row.IsGoFont}}
    {{printf "\"golang.org/x/image/font/gofont/%s\"" $row.LowerName}}
    {{- end}}
{{- end}}
)

//go:embed FontFiles/*.ttf FontFiles/*.otf
var fontFiles embed.FS

var (
{{- range $i, $row := .}}
    {{- if not $row.IsGoFont}}
    {{printf "%sTTF, _ = fontFiles.ReadFile(` + "`%s`" + `)" $row.LowerName $row.FileName}}
    {{- end}}
{{- end}}
)

var (
{{- range $i, $row := .}}
    {{- if $row.IsGoFont}}
    {{printf "%-35s = NewFont(%d, %s.TTF)" $row.FontName $row.ID $row.LowerName}}
    {{- else}}
    {{printf "%-35s = NewFont(%d, %sTTF)" $row.FontName $row.ID $row.LowerName}}
    {{- end}}
{{- end}}
)

var Map = map[string]*Font{
{{- range $i, $row := .}}
    {{printf "\"%s\": %s," $row.FontName $row.FontName}}
{{- end}}
}

var Names = []string{
{{- range $i, $row := .}}
    {{printf "\"%s\"," $row.FontName}}
{{- end}}
}
`

	fontNamesTempl = template.Must(template.New("fontNames").Parse(fontNamesTemplate))
)

type TemplateData struct {
	ID                            int
	FontName, LowerName, FileName string
	IsGoFont                      bool
}

func main() {
	fileList, err := filepath.Glob(filepath.Join(fontDir, fontPattern))
    if err != nil {
		log.Fatal(err)
    }

	dataList := make([]TemplateData, 0)
	for _, name := range goFontList {
		data := TemplateData{
			FontName:  name,
			LowerName: strings.ToLower(name),
			IsGoFont:  true,
		}
		dataList = append(dataList, data)
	}
	for _, fileName := range fileList {
		baseName := filepath.Base(fileName)
		id, err := strconv.Atoi(baseName[:3])
		if err != nil {
			log.Fatalf("Couldn't parse '%s': %v", baseName[:3], err)
		}
		name := baseName[3 : len(baseName)-len(filepath.Ext(baseName))]
		name = strings.Replace(name, "-", "", -1)
		data := TemplateData{
			ID:        id,
			FontName:  name,
			LowerName: strings.ToLower(name),
			FileName:  fileName,
			IsGoFont:  false,
		}
		dataList = append(dataList, data)
	}

	fh, err := os.Create(fontFile)
    if err != nil {
		log.Fatal(err)
    }
	defer fh.Close()
	err = fontNamesTempl.Execute(fh, dataList)
    if err != nil {
		log.Fatal(err)
    }
}
