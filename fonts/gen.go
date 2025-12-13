//go:build ignore
// +build ignore

package main

// Erstellt die Datei fontnames.go. Diese enthält für die Go-Fonts, aber
// auch für alle Fonts, die als TrueType- oder OpenType-Dateien in ./FontFiles
// abgelegt sind, Variablen für eine einfachere Verwendung in
// Graphikprogrammen.
//
// Dieses Programm muss manuell mit "go run gen.go" gestartet werden!
// Dies ist allerdings nur dann notwendig, wenn es Anpassungen im Verzeichnis
// FontFiles gibt - was wohl eher selten passiert.

import (
	"log"
	"os"
	"path/filepath"
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

	fontNamesTemplate = `// Code generated  DO NOT EDIT.

package fonts

// WICHTIG: Diese Datei sollte nicht manuell angepasst werden!
// Sie wird automatisch per Script neu erzeugt. Allfaellige manuelle
// Anpassungen werden damit ueberschrieben.

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
    {{printf "%-35s = Parse(%s.TTF)" $row.FontName $row.LowerName}}
    {{- else}}
    {{printf "%-35s = Parse(%sTTF)" $row.FontName $row.LowerName}}
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
	FontName, LowerName, FileName string
	IsGoFont                      bool
}

func main() {
	fileList, err := filepath.Glob(filepath.Join(fontDir, fontPattern))
	check(err)

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
		name := baseName[3 : len(baseName)-len(filepath.Ext(baseName))]
		name = strings.Replace(name, "-", "", -1)
		data := TemplateData{
			FontName:  name,
			LowerName: strings.ToLower(name),
			FileName:  fileName,
			IsGoFont:  false,
		}
		dataList = append(dataList, data)
	}

	fh, err := os.Create(fontFile)
	check(err)
	defer fh.Close()
	err = fontNamesTempl.Execute(fh, dataList)
	check(err)
}

func check(err error) {
	if err != nil {
		log.Fatalf("%v", err)
	}
}
