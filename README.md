# Go Graphics for TEKO

`gg` ist eine Sammlung von Typen und Funktionen zur Erstellung von 2D
Pixel- oder Rasterbildern. Die Sammlung dient im Rahmen des Unterrichtes im
Fach "Mathematik und SW-Tools" als Werkzeugkasten.
Sie basiert massgeblich auf dem gleichnamigen
Paket [gg](https://github.com/fogleman/gg), wurde geringfügig modifiziert
und um ein paar Unterpakete ergänzt - so finden sich bspw. in 'gg/color'
Farben (analog 'image/color').

## Installation

Mit folgedem Befehl wird das Paket in der neusten Version, inkl. aller
Unterpakete installiert.

    go install github.com/stefan-muehlebach/gg@latest

## Dokumentation

https://pkg.go.dev/github.com/stefan-muehlebach/gg?tab=doc

## Hello, Circle!

Der Klassiker unter den Erstlingswerken. Zwar in einer graphischen aber noch
recht unbunten Variante.

```go
package main

import (
    "github.com/stefan-muehlebach/gg"
    "github.com/stefan-muehlebach/gg/colornames"
)

func main() {
    gc := gg.NewContext(512, 512)
    gc.DrawCircle(256.0, 256.0, 224.0)
    gc.SetFillColor(colornames.Black)
    gc.Fill()
    gc.SavePNG("circle.png")
}
```

## Beispiele

Ein wichtiger Bestandteil dieser Sammlung sind die
[Beispiele](https://github.com/stefan-muehlebach/gg/tree/main/examples).
Auch wenn sie ursprünglich zum Testen der Software erstellt wurden,
können sie auch verwendet werden, um die einzelnen Funktionen besser zu
verstehen.

## Graphische Umgebungen

Stehen im Zentrum jeder Anwendung und werden über folgende Funktionen erstellt.

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
DrawImage(im image.Image, x, y float64)
DrawImageAnchored(im image.Image, x, y, ax, ay float64)
SetPixel(x, y int, c color.Color)

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
SetFontFace(face font.Face)
LoadFontFace(path string, points float64) error
```

## Farbfunktionen

Für die Erstellung und Veränderung von Farben stehen im Package `gg/color`
eine Anzahl von neuen Farbtypen (RGBAF, HSL, HSV, etc) zur Verfügung.
Im Package `gg/colornames` sind alle Farben aus SVG 1.1 als
vorbereitete Variablen zu finden.

```go
SetStrokeColor(c color.Color)
SetFillColor(c color.Color)
```

## Optionen für Linen und Flächen

```go
SetStrokeWidth(lineWidth float64)
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

## Koordinatensysteme

Wie das Koordinatensystem interpretiert werden soll, kann mit diesen
Funktionen beeinflusst werden.

```go
Identity()
Translate(tx, ty float64)
Scale(sx, sy float64)
ScaleAbout(sx, sy, x, y float64)
Rotate(angle float64)
RotateAbout(angle, x, y float64)
TransformPoint(x, y float64) (tx, ty float64)

Matrix() (geom.Matrix)
SetMatrix(m geom.Matrix)
```

Die Funktionen `RotateAbout` und `ScaleAbout` haben ihren Bezugspunkt an der
Stelle `x`, `y` und nicht beim Ursprung.

## Stackfunktionen

Die aktuellen Einstellungen (insbesondere die Koordiantentransformationen)
können auf einem Stack gesichert und wiederhergestellt werden.
Der Aufruf ist mehrfach und damit geschachtelt möglich.

```go
Push()
Pop()
```

## Ausschnittsfunktionen

Mit folgenden Funktionen werden die Zeicheoperationen auf ein bestimtes Gebiet
beschränkt. Der aktuelle Pfad wird dabei als Rand des Gebietes verwendet.

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

## Ein weiteres Beispiel

Das Bild dazu siehst du weiter unten.

```go
package main

import (
    "github.com/stefan-muehlebach/gg"
    "github.com/stefan-muehlebach/gg/color"
)

func main() {
    const S = 512
    dc := gg.NewContext(S, S)
    dc.SetFillColor(color.RGBAF{0, 0, 0, 0.2})
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

## Hall of fame

Ohne die Unterstützung folgender Personen wäre das Werk nicht, viel später
oder weniger freudig zustande gekommen.

* Weber, Nicolas - für die allererste Installation dieses Paketes ausserhalb meines Computers.

