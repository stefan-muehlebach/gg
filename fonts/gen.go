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
	fontPattern = "[0-9][0-9]-*.[ot]tf"
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
    {{printf "%s, _ = fontFiles.ReadFile(\"%s\")" $row.LowerName $row.FileName}}
    {{- end}}
{{- end}}
)

var (
{{- range $i, $row := .}}
    {{- if $row.IsGoFont}}
    {{printf "%-35s = Parse(%s.TTF)" $row.FontName $row.LowerName}}
    {{- else}}
    {{printf "%-35s = Parse(%s)" $row.FontName $row.LowerName}}
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

// func Main() {
// 	fileList, err := filepath.Glob(filepath.Join(fontDir, fontPattern))
// 	check(err)

// 	fh, err := os.OpenFile(fontFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
// 	check(err)
// 	defer fh.Close()

// 	fmt.Fprintf(fh, `// Code generated  DO NOT EDIT.

// package fonts

// // WICHTIG: Diese Datei sollte nicht manuell angepasst werden!
// // Sie wird automatisch per Script neu erzeugt. Allfaellige manuelle
// // Anpassungen werden damit ueberschrieben.`)

// 	fmt.Fprintf(fh, "\n\nimport (\n")
// 	fmt.Fprintf(fh, "    \"embed\"\n")
// 	// fmt.Fprintf(fh, "    \"golang.org/x/image/font/opentype\"\n")
// 	for _, goFont := range goFontList {
// 		pkgName := strings.ToLower(goFont)
// 		fmt.Fprintf(fh, "    \"golang.org/x/image/font/gofont/%s\"\n", pkgName)
// 	}
// 	fmt.Fprintf(fh, ")\n\n")

// 	fmt.Fprintf(fh, "//go:embed FontFiles/*.ttf FontFiles/*.otf\n")
// 	fmt.Fprintf(fh, "var fontFiles embed.FS\n\n")

// 	fmt.Fprintf(fh, "var (\n")
// 	for _, pathName := range fileList {
// 		baseName := filepath.Base(pathName)
// 		fontName := baseName[3 : len(baseName)-len(filepath.Ext(baseName))]
// 		fontName = strings.Replace(fontName, "-", "", -1)
// 		varName := strings.ToLower(fontName)
// 		ttfFontList = append(ttfFontList, fontName)
// 		fmt.Fprintf(fh, "    %s, _ = fontFiles.ReadFile(\"%s\")\n", varName, pathName)
// 	}
// 	fmt.Fprintf(fh, ")\n\n")

// 	fmt.Fprintf(fh, "var (\n")
// 	for _, name := range goFontList {
// 		varName := strings.ToLower(name)
// 		fmt.Fprintf(fh, "    %-35s = Parse(%s.TTF)\n", fmt.Sprintf("%s", name), varName)
// 		// fmt.Fprintf(fh, "    %-35s = opentype.Parse(%s.TTF)\n", fmt.Sprintf("%s, _", name), varName)
// 	}
// 	for _, name := range ttfFontList {
// 		varName := strings.ToLower(name)
// 		fmt.Fprintf(fh, "    %-35s = Parse(%s)\n", fmt.Sprintf("%s", name), varName)
// 		// fmt.Fprintf(fh, "    %-35s = opentype.Parse(%s)\n", fmt.Sprintf("%s, _", name), varName)
// 	}
// 	fmt.Fprintf(fh, ")\n\n")

// 	fmt.Fprintf(fh, "var Map = map[string]*Font{\n")
// 	// fmt.Fprintf(fh, "var Map = map[string]*opentype.Font{\n")
// 	for _, name := range goFontList {
// 		fmt.Fprintf(fh, "    %-35s %s,\n", fmt.Sprintf("\"%s\":", name), name)
// 	}
// 	for _, name := range ttfFontList {
// 		fmt.Fprintf(fh, "    %-35s %s,\n", fmt.Sprintf("\"%s\":", name), name)
// 	}
// 	fmt.Fprintf(fh, "}\n\n")

// 	fmt.Fprintf(fh, "var Names = []string{\n")
// 	for _, name := range goFontList {
// 		fmt.Fprintf(fh, "    \"%s\",\n", name)
// 	}
// 	for _, name := range ttfFontList {
// 		fmt.Fprintf(fh, "    \"%s\",\n", name)
// 	}
// 	fmt.Fprintf(fh, "}\n\n")
// }

// // func do(ttfName string) {
// //     fontName := fontName(ttfName)
// //     pkgName := pkgName(ttfName)
// //     if err := os.Mkdir(pkgName, 0777); err != nil && !os.IsExist(err) {
// //         log.Fatal(err)
// //     }
// //     src, err := ioutil.ReadFile(filepath.Join(fontDir, ttfName))
// //     if err != nil {
// //         log.Fatal(err)
// //     }

// //     desc := "a proportional-width, sans-serif"
// //     if strings.Contains(ttfName, "Mono") {
// //         desc = "a fixed-width, slab-serif"
// //     }

// //     b := new(bytes.Buffer)
// //     fmt.Fprintf(b, "// generated by go run gen.go; DO NOT EDIT\n\n")
// //     fmt.Fprintf(b, "// Package %s provides the %q TrueType font\n", pkgName, fontName)
// //     fmt.Fprintf(b, "// from the Go font family. It is %s font.\n", desc)
// //     fmt.Fprintf(b, "//\n")
// //     fmt.Fprintf(b, "// See https://blog.golang.org/go-fonts for details.\n")
// //     fmt.Fprintf(b, "package %s\n\n", pkgName)
// //     fmt.Fprintf(b, "// TTF is the data for the %q TrueType font.\n", fontName)
// //     fmt.Fprintf(b, "var TTF = []byte{")
// //     for i, x := range src {
// //         if i&15 == 0 {
// //             b.WriteByte('\n')
// //         }
// //         fmt.Fprintf(b, "%#02x,", x)
// //     }
// //     fmt.Fprintf(b, "\n}\n")

// //     dst, err := format.Source(b.Bytes())
// //     if err != nil {
// //         log.Fatal(err)
// //     }
// //     if err := ioutil.WriteFile(filepath.Join(pkgName, "data.go"), dst, 0666); err != nil {
// //         log.Fatal(err)
// //     }
// // }

// // // fontName maps "Go-Regular.ttf" to "Go Regular".
// // func fontName(ttfName string) string {
// //     extLen := len(filepath.Ext(ttfName))
// //     s := ttfName[:len(ttfName)-extLen]
// //     s = strings.Replace(s, "-", " ", -1)
// //     return s
// // }

// // // pkgName maps "Go-Regular.ttf" to "goregular".
// // func pkgName(ttfName string) string {
// //     extLen := len(filepath.Ext(ttfName))
// //     s := ttfName[:len(ttfName)-extLen]
// //     s = strings.Replace(s, "-", "", -1)
// //     s = strings.ToLower(s)
// //     return s
// }

func check(err error) {
	if err != nil {
		log.Fatalf("%v", err)
	}
}
