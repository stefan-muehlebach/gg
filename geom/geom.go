// Package mit Typen für grundlegende, geometrische 2D-Operationen.
//
// Das Package geht von einer 2 dimensionalen Ebene aus, in welcher Punkte
// mittels X- und Y-Koordinaten (beides float64-Werte) bezeichnet werden.
// Es gibt zwei Datentypen für geometrische Objekte: Point und Rectangle
// und daneben den Datentyp Matrix, mit welchem linear affine Transformationen
// beschrieben und auf die Typen Point und Rectangle angewendet werden können.
//
// # Point
//
// Ein Punkt ist ein struct mit den Feldern X und Y für die X-, resp. die
// Y-Koordinate des Punktes. Beide Felder sind exportiert. Nahezu alle
// Methoden des Types Point haben einen Value-Receiver, d.h. sie verändern die
// Werte des Receivers nicht. Ich habe mich für den Begriff Point entschieden,
// obwohl es sich bei dieser Struktur genaugenommen um Vektoren handelt und
// viele der implementierten Methoden stark an das Rechnen mit Vektoren
// erinnern.
//
// # Erzeugen von Punkten
//
// Über die folgenden zwei Funktionen können Punkte erstellt werden:
//
//	NewPoint(x, y float64) (Point)
//	NewPointIMG(p image.Point) (Point)
//
// Beide Funktionen sind eigentlich überflüssig: dadurch dass die Felder X und
// Y in Point exportiert sind, können neue Punkte jederzeit über
// Struct-Literals erzeugt werden (z.B. geom.Point{1.0, 2.0}).
//
// # Modifikation von Punkten
//
// Mit den folgenden Methoden werden neue Punkte durch Modifikation bestehender
// Punkte erzeugt:
//
//	Add(q Point) (Point)
//	AddXY(x, y float64) (Point)
//	Sub(q Point) (Point)
//	SubXY(x, y float64) (Point)
//	Mul(k float64) (Point)
//	Neg() (Point)
//	Move(dp Point)           // Pointer receiver! Verschiebt den Punkt selber
//
// Die XY-Varianten der Methoden Add und Sub sind entstanden, da es oft
// einfacher ist zu schreiben:
//
//	p2 = p1.AddXY(1, 2)
//
// als
//
//	p2 = p1.Add(geom.Point{1, 2})
//
// # Abstände
//
// Zwei Methoden stehen zur Bestimmung des euklidischen Abstandes zwischen
// zwei Punkten zur Verfügung. Dist2 liefert den Abstand _im Quadrat_,
// verzichtet aus Perf.gründen auf das Anwenden der Wurzel. Will man aber nur
// Abstände nur miteinander _vergleichen_, so reicht Dist2, denn:
// wenn a > b, dann ist auch a^2 > b^2 und umgekehrt.
//
//	Distance(q Point) (float64)
//	Dist2(q Point) (float64)
//
// # Checks
//
// Mit Eq werden zwei Punkte auf Identität geprüft. Mit In kann ermittelt
// werden, ob der Punkt innerhalb des Rechtecks r liegt. Beachte hierzu, wie
// die Koordinaten des Typs Rectangle zu verstehen sind.
//
//	Eq(q Point) (bool)
//	In(r Rectangle) (bool)
//
// # Weitere Berechnungen
//
// Folgende Methoden stehen für weitergehende Berechnungen zur Verfügung:
//
//	Interpolate(q Point, t float64) (Point)
//	Max(q Point) (Point)
//	Min(q Point) (Point)
//
// # Konvertierung
//
// Folgende Methoden stehen zur Konvertierung (im weitesten Sinne) von Punkten
// zur Verfügung. Mit der Methode Set kann man bspw. die Koordinaten von
// Punkten via Flag oder Kommandozeile einlesen (Getter-Interface in flags).
//
//	AsCoord() (x, y float64)
//	Int() (image.Point)
//	String() (string)
//	Set(s string) (error)   // Pointer receiver!
//
// # Rectangle
//
// Der zweite geometrische Type ist Rectangle mit welchem ein rechteckiges
// Gebiet in einer 2-dimensionalen Ebene definiert werden kann. Der Typ besteht
// aus zwei (exportierten) Feldern Min und Max, welche die Eckpunkte eines
// Rechtecks enthaten. Dabei muss beachtet werden, dass immer folgendes gilt:
// Min.X <= Max.X UND Min.Y <= Max.Y.
//
// # Erzeugen von Rechtecken
//
//	NewRectangle(x0, y0, x1, y1 float64) (Rectangle)
//	NewRectangleWH(x, y, w, h float64) (Rectangle)
//	NewRectangleCWH(mx, my, w, h float64) (Rectangle)
//	NewRectangleIMG(r image.Rectangle) (Retangle)
//
// # Modifikation von Rechtecken
//
//	Add(p Point) (Rectangle)
//	Sub(p Point) (Rectangle)
//	Move(dp Point) (Rectangle)    // Pointer-Receiver
//
// # Checks und Vergleiche
//
//	Empty() (bool)
//	Eq(s Rectangle) (bool)
//	In(s Rectangle) (bool)
//	Overlaps(s Rectangle) (bool)
//
// # Grössen und spezielle Punkte
//
//	Dx() (float64)
//	Dy() (float64)
//	Size() (Point)
//
// Die folgenden Methoden dienen dazu, bestimmte Punkte auf dem Rand des
// Rechtecks einfacher zu ermitteln. Die Bezeichnungen entsprechen dabei
// den Angaben auf einem virtuellen Kompass. Bspw. bedeutet 'NW' Nordwest,
// bezeichnet also den linken oberen Punkt des Rechtecks.
//
//	NW() (Point)
//	N() (Point)
//	NE() (Point)
//	W() (Point)
//	C() (Point)
//	E() (Point)
//	SW() (Point)
//	S() (Point)
//	SE() (Point)
//
// # Weitere Berechnungen
//
//	Intersect(s Rectangle) (Rectangle)
//	Union(s Rectangle) (Rectangle)
//	Inset(dx, dy float64) (Rectangle)
//	PosRel(p Point) (fx, fy float64)
//	RelPos(fx, fy float64) (Point)
//	SetInside(p Point) (Point)
//	Canon() (Rectangle)
//
// # Konvertierung
//
//	AsCoord() (x, y, w, h float64)
//	Int() (image.Rectangle)
//	String() (string)
//	Set(s string) (error)
//
// # Matrizen
//
// Der letzte Datentyp in diesem Package ist Matrix, mit welchem eine
// linear-affine Transformation in 2D dargestellt werden kann. Der Typ ist
// wiederum ein struct mit den Elementen m_11 bis m_23 einer 3x3 Matrix.
//
// # Basis-Matrizen
//
//	Identity() (Matrix)
//	Translate(d Point) (Matrix)
//	Rotate(a float64) (Matrix)
//	RotateAbout(rp Point, a float64) (Matrix)
//	Scale(sx, sy float64) (Matrix)
//	ScaleAbout(sp Point, sx, sy float64) (Matrix)
//
// # Verknüpfungen und Invertierung von Matrizen
//
//	Multiply(b Matrix) (Matrix)
//	Inv() (Matrix)
//
// # Transformation von Matrizen
//
//	Translate(d Point) (Matrix)
//	Scale(sx, sy float64) (Matrix)
//	ScaleAbout(sp Point, sx, sy float64) (Matrix)
//	Rotate(a float64) (Matrix)
//	RotateAbout(rp Point, a float64) (Matrix)
//
// # Transformation von Punkten und Rechtecken
//
//	Transform(p Point) (Point)
//	TransformRect(r Rectangle) (Rectangle)
//
// # Diverse Methoden
//
//	String() (string)
package geom
