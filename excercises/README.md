# Graphische Trainingseinheiten

In diesem Verzeichnis sind mehrere Aufgaben enthalten, die graphisch
implementiert werden müssen. Zu den einzelnen Aufgaben gibt es nur wenige
Erklärungen in schriftlicher Form, die zentrale Aufgabenstellung ist die
Bildvorlage. Das Ziel ist, jedes gezeigte Bild mit einem eigenständigen
Go-Programm nachzuzeichnen. Dabei gilt es auf Folgendes zu achten:

* Wie stark gleichen sich die beiden Bilder (die Vorlage und das selbst
  erstellte?
* Ist der Code so sparsam/kurz/ökonomisch wie möglich?
* Sind alle Redundanzen eliminiert?
* Wurden zur Steuerung, resp. Modifikation des Bildes Parameter eingesetzt?

## TicTacToe

Anhand eines Beispiels möchte ich zeigen, wie ich mir das genau vorstelle.
Für das folgende Bild (ein angefangenes TicTacToe) werde ich nacheinander
mehrere Go-Programme erstellen. In jeder Iteration wird der Code kompakter,
generischer und besser.

Die Vorlagen sind immer quadratisch und als Breite empfiehlt sich eine Zahl
zwischen 300 und 1000. Am Schönsten sind natürlich 2er Potenzen wie 256, 512
oder 1024.

![TicTacToe](tictactoe01.png)

```go
package main

import (
	"github.com/stefan-muehlebach/gg"
	"github.com/stefan-muehlebach/gg/colornames"
)

func main() {
    gc := gg.NewContext(256, 256)
    gc.SetFillColor(colornames.Beige)
    gc.Clear()

    gc.SetStrokeColor(colornames.DarkSlateGray)
    gc.SetStrokeWidth(7.0)
    gc.DrawLine( 20, 87, 236, 87)
    gc.DrawLine( 20, 169, 236, 169)
    gc.DrawLine( 87, 20,  87, 236)
    gc.DrawLine(169, 20, 169, 236)
    gc.Stroke()

    gc.SetStrokeColor(colornames.DarkRed)
    gc.SetStrokeWidth(10.0)
    gc.DrawCircle( 46,  46, 26)
    gc.DrawCircle(210, 210, 26)
    gc.Stroke()

    gc.SetStrokeColor(colornames.DarkGreen)
    gc.DrawLine(106, 106, 150, 150)
    gc.DrawLine(106, 150, 150, 106)
    gc.DrawLine( 24, 106,  68, 150)
    gc.DrawLine( 24, 150,  68, 106)
    gc.DrawLine(188, 106, 232, 150)
    gc.DrawLine(188, 150, 232, 106)
    gc.Stroke()

    gc.SavePNG("tictactoe01.png")
}
```
Wenn man das Programm ausführt, erscheint im aktuellen Verzeichnis die Datei
`tictactoe01.png` mit dem gewünschten Output.
Grundsätzlich könnte man alle Aufgaben auf diese Weise lösen - was jedoch
überhaupt nicht meine Intention ist!

Es gibt ein paar Mängel im obigen Programm



Sind Wiederholungen im Bild durch Schlaufen implementiert?

(Stand 24.12.2023) 16 Bilder, welche mit Hilfe
von Go und [gg](https://github.com/stefan-muehlebach/gg) _und so wenig Code
wie möglich_ nachzubilden sind.
Mit entsprechendem Aufwand können alle Bilder durch ein Go-Programm in
Form einer laaaaaaangen Liste von Zeichenbfehlen erstellt werden - das ist
jedoch nicht das Ziel!
Allen Graphiken liegt ein bestimmter algorithmischer Aufbau zugrunde.
Das heisst, mit geschickter Verwendung von Schlaufen, Iterationen und weiteren
Programmiertechniken, lässt sich der Aufwand massiv verkleinern.

## Erste und wichtigste Schritte

Dieser Abschnitt erklärt, wie die Programme zur Erzeugung von PNG-Dateien
grundsätzlich aufgebaut sind, welche Einstellungen (Grössen, Farben, etc) zu
verwenden sind und zeigt Schritt für Schritt wie ein einfaches Bild erstellt
werden kann.

Die Programme habe alle

## List der Aufgaben

![Irrgarten](irrgarten.png)    ![Spirale](spirale.png)
![Schachbrett](chessboard.png) ![Moiree](moiree.png)
![Quadrate](vertigo.png)       ![Bauhaus](bauhaus.png)

![](divided-circle.png)
![](graph.png)
![](text-transform.png)
![](farn.png)
![](raster.png)
![](color-cells.png)
![](relief.png)

![](raster-image.png)
![](flower-of-life.png)
