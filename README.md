# Go Graphics

`gg` ist eine Sammlung von Typen und Funktionen zur Erstellung von 2D
Pixel- oder Rasterbildern.

![Stars](http://i.imgur.com/CylQIJt.png)

## Installation

    go install github.com/stefan-muehlebach/gg@latest

## Dokumentation

- pkg.go.dev: https://pkg.go.dev/github.com/stefan-muehlebach/gg?tab=doc

## Hello, Circle!

Der Klassiker unter den Erstlingswerken...

```go
package main

import (
    "github.com/stefan-muehlebach/gg"
)

func main() {
    dc := gg.NewContext(512, 512)
    dc.DrawCircle(256, 256, 224)
    dc.SetFillColor(gg.NewRGB(0, 0, 0))
    dc.Fill()
    dc.SavePNG("circle.png")
}
```

## Beispiele

Ein wichtiger Bestandteil dieser Sammlung sind die [Beispiele](https://github.com/stefan-muehlebach/gg/tree/main/examples) included.
Auch wenn sie ursprünglich zum Testen der Software erstellt wurden,
können sie auch verwendet werden, um die einzelnen Funktionen besser zu
verstehen.

![Examples](http://i.imgur.com/tMFoyzu.png)


## Graphische Kontexte

Stehen im Zentrum jeder Anwendung und können nur über wenige Funktionen
erstellt werden.

```go
NewContext(width, height int) *Context
NewContextForImage(im image.Image) *Context
NewContextForRGBA(im *image.RGBA) *Context
```

## Zeichenfunktionen

Was wäre eine Graphikumgebung ohne Werkzeuge, um Kreise, Geraden oder
Rechtecke (und vieles mehr) zu erstellen,

```go
DrawPoint(x, y, r float64)
DrawLine(x1, y1, x2, y2 float64)
DrawRectangle(x, y, w, h float64)
DrawRoundedRectangle(x, y, w, h, r float64)
DrawCircle(x, y, r float64)
DrawArc(x, y, r, angle1, angle2 float64)
DrawEllipse(x, y, rx, ry float64)
DrawEllipticalArc(x, y, rx, ry, angle1, angle2 float64)
DrawRegularPolygon(n int, x, y, r, rotation float64)
DrawImage(im image.Image, x, y int)
DrawImageAnchored(im image.Image, x, y int, ax, ay float64)
SetPixel(x, y int)

MoveTo(x, y float64)
LineTo(x, y float64)
QuadraticTo(x1, y1, x2, y2 float64)
CubicTo(x1, y1, x2, y2, x3, y3 float64)
ClosePath()
ClearPath()
NewSubPath()

Clear()
Stroke()
Fill()
FillStroke()
StrokePreserve()
FillPreserve()
```

Bei der Darstellung von Bildern oder Schriften kann es hilfreich sein, diese
zentriert an einer bestimmten Stelle auszugeben. Verwende dazu die Funktion
`DrawImageAnchored` mit 0.5 für die Parameter `ax` und `ay`. Mit 0 wird das
Bild links oben ausgerichtet, mit 1 rechts unten. `DrawStringAnchored`
bietet die gleiche Funktionalität für Texte an, so dass der Aufruf von
`MeasureString` dazu nicht notwendig sein sollte.

## Textfunktionen

Übernehmen sogar den Zeilenumbruch für dich!

```go
DrawString(s string, x, y float64)
DrawStringAnchored(s string, x, y, ax, ay float64)
DrawStringWrapped(s string, x, y, ax, ay, width, lineSpacing float64, align Align)
MeasureString(s string) (w, h float64)
MeasureMultilineString(s string, lineSpacing float64) (w, h float64)
WordWrap(s string, w float64) []string
SetFontFace(fontFace font.Face)
LoadFontFace(path string, points float64) error
```

## Farbfunktionen

Für die Erstellung und Veränderung von Farben stehen im Package `gg/color`
eine Anzahl von neuen Farbtypen (HSL, HSV, etc) zur Verfügung.
Das Package `gg/colornames` dagegen bietet die Farben aus SVG 1.1 als
vorbereitete Variablen an.

```go
SetLineColor(c color.Color)
SetFillColor(c color.Color)
```

## Optionen für Linen und Flächen

```go
SetLineWidth(lineWidth float64)
SetLineCap(lineCap LineCap)
SetLineJoin(lineJoin LineJoin)
SetDash(dashes ...float64)
SetDashOffset(offset float64)
SetFillRule(fillRule FillRule)
```

## Verläufe und Muster

`gg` unterstützt lineare, radiale und konische Verläufe sowie Muster um
graphische Objekte damit zu füllen. Es können sogar eigene Muster definiert
werden.

```go
SetFillStyle(pattern Pattern)
SetStrokeStyle(pattern Pattern)
NewSolidPattern(color color.Color)
NewLinearGradient(x0, y0, x1, y1 float64)
NewRadialGradient(x0, y0, r0, x1, y1, r1 float64)
NewConicGradient(cx, cy, deg float64)
NewSurfacePattern(im image.Image, op RepeatOp)
```

## Koordinatentransformationen

```go
Identity()
Translate(x, y float64)
Scale(x, y float64)
Rotate(angle float64)
ScaleAbout(sx, sy, x, y float64)
RotateAbout(angle, x, y float64)
TransformPoint(x, y float64) (tx, ty float64)
```

Die Funktionen `RotateAbout` und `ScaleAbout` haben ihren Bezugspunkt an der
Stelle `x`, `y` und nicht beim Ursprung.

## Stackfunktionen

Die aktuellen Einstellungen können gesichert oder wiederhergestellt werden.
Der Aufruf ist mehrfach und geschachtelt möglich.

```go
Push()
Pop()
```

## Funktionen für den Zuschnitt 

Mit folgenden Funktionen sind Zeicheoperationen nur in einem bestimmten
Gebiet zu sehen.

```go
Clip()
ClipPreserve()
ResetClip()
AsMask() *image.Alpha
SetMask(mask *image.Alpha)
InvertMask()
```

## Hilfsfunktionen

Manchmal will man sie einfach nicht selber schreiben müssen!

```go
Radians(degrees float64) float64
Degrees(radians float64) float64
LoadImage(path string) (image.Image, error)
LoadPNG(path string) (image.Image, error)
SavePNG(path string, im image.Image) error
```

![Separator](http://i.imgur.com/fsUvnPB.png)

## Ein weiteres Beispiel

Das Bild dazu siehst du weiter unten.

```go
package main

import "github.com/stefan-muehlebach/gg"

func main() {
	const S = 1024
	dc := gg.NewContext(S, S)
	dc.SetFillColor(gg.NewRGBA(0.0, 0.0, 0.0, 0.2))
	for i := 0; i < 360; i += 15 {
		dc.Push()
		dc.RotateAbout(gg.Radians(float64(i)), S/2, S/2)
		dc.DrawEllipse(S/2, S/2, S*7/16, S/8)
		dc.Fill()
		dc.Pop()
	}
	dc.SavePNG("ellipses.png")
}
```

![Ellipses](http://i.imgur.com/J9CBZef.png)
