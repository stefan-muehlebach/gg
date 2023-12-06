package geom

import (
    "fmt"
    "math"
)

// Der Datentyp Matrix enthält die relevanten Felder einer Transformations-
// Matrix für linear-affine Abbildungen in der Ebene. Mij ist das Element
// auf der Zeile i, in der Spalte j. Die letzte Zeile der Matrix ist immer
// [0, 0, 1] und wird daher nicht mitgeführt.
type Matrix struct {
    M11, M12, M13 float64
    M21, M22, M23 float64
}

// Erzeugt die Einheitsmatrix.
func Identity() (Matrix) {
    return Matrix{1.0, 0.0, 0.0,
                  0.0, 1.0, 0.0}
}

// Erzeugt eine Translationsmatrix. Die Angaben für die Translation befinden
// sich als X-, resp. Y-Koordinate im Punkt d .
func Translate(d Point) (Matrix) {
    return Matrix{1.0, 0.0, d.X,
                  0.0, 1.0, d.Y}
}

// Erzeugt eine Rotationsmatrix um den Winkel a (im Bogenmass). Drehpunkt ist
// der Ursprung des Koordinatensystems, Drehrichtung ist im Gegenuhrzeigersinn.
func Rotate(a float64) (Matrix) {
    s := math.Sin(a)
    c := math.Cos(a)
    return Matrix{c, -s, 0.0,
                  s,  c, 0.0}
}

// Erzeugt eine Rotationsmatrix um den Winkel a (im Bogenmass) mit Drehpunkt
// bei rp.
func RotateAbout(rp Point, a float64) (Matrix) {
    m := Translate(rp)
    m  = m.Rotate(a)
    m  = m.Translate(rp.Neg())
    return m
}

// Erzeugt eine Skalierungsmatrix mit den Skalierungsfaktoren sx in X-Richtung
// und sy in Y-Richtung. Zentrum der Skalierung ist der Ursprung des
// Koordinatensystems.
func Scale(sx, sy float64) (Matrix) {
    return Matrix{ sx, 0.0, 0.0,
                  0.0,  sy, 0.0}
}

// Erzeugt eine Skalierungsmatrix mit den Skalierungsfaktoren sx in X-Richtung
// und sy in Y-Richtung. Zentrum der Skalierung ist der Punkt sp.
func ScaleAbout(sp Point, sx, sy float64) (Matrix) {
    m := Translate(sp)
    m  = m.Scale(sx, sy)
    m  = m.Translate(sp.Neg())
    return m
}

// Invertiert die Matrix a und liefert das Resultat als neue Matrix. Da wir
// es hier eigentlich nie singulären Matrizen zu tun haben (ausser jemand
// erstellt bewusst die Nullmatrix oder eine Skalierungsmatrix mit 0 als
// einem der beiden Faktoren) verzichten wir hier zugunsten der Performance
// auf einen entsprechenden Test.
func (a Matrix) Inv() Matrix {
    det := a.M11*a.M22 - a.M12*a.M21
    return Matrix{ a.M22 / det,
                  -a.M12 / det,
                  (a.M12*a.M23 - a.M22*a.M13) / det,

                  -a.M21 / det,
                   a.M11 / det,
                  (a.M21*a.M13 - a.M11*a.M23) / det}
}


// Multipliziert die Matrizen a und b (d.h. berechnet a*b) und liefert das
// Resultat als neue Matrix.
func (a Matrix) Multiply(b Matrix) Matrix {
    return Matrix{a.M11*b.M11 + a.M12*b.M21,
                  a.M11*b.M12 + a.M12*b.M22,
                  a.M11*b.M13 + a.M12*b.M23 + a.M13,

                  a.M21*b.M11 + a.M22*b.M21,
                  a.M21*b.M12 + a.M22*b.M22,
                  a.M21*b.M13 + a.M22*b.M23 + a.M23}
}

// Die Methoden Translate, Rotate, RotateAbout, Scale und ScaleAbout sind
// Hilfsmethoden, um eine bestehende Matrix m mit einer entsprechenden
// Transformationsmatrix zu ergänzen.
func (m Matrix) Translate(d Point) Matrix {
    return m.Multiply(Translate(d))
}

// Siehe Translate
func (m Matrix) Rotate(angle float64) Matrix {
    return m.Multiply(Rotate(angle))
}

// Siehe Translate
func (m Matrix) RotateAbout(rp Point, a float64) Matrix {
    return m.Multiply(RotateAbout(rp, a))
}

// Siehe Translate
func (m Matrix) Scale(sx, sy float64) Matrix {
    return m.Multiply(Scale(sx, sy))
}

// Siehe Translate
func (m Matrix) ScaleAbout(sp Point, sx, sy float64) Matrix {
    return m.Multiply(ScaleAbout(sp, sx, sy))
}

// Hier schliesslich spielt die Musik: eine Matrix wird für die Transformation
// eines Punktes verwendet. Es wird ein neuer Punkt erstellt, aktuell gibt
// noch keinen Bedarf nach 'in place' Transformation.
func (m Matrix) Transform(p Point) (Point) {
    return Point{m.M11*p.X + m.M12*p.Y + m.M13,
                 m.M21*p.X + m.M22*p.Y + m.M23}
}

// Will man den zweiten geometrischen Typ (Rectangle) transformieren, dann
// ist dies die Methode der Wahl.
func (m Matrix) TransformRect(r Rectangle) (Rectangle) {
    return Rectangle{m.Transform(r.Min), m.Transform(r.Max)}
}

// TransformPoint und TransformVector sind zwei Methoden, die aus dem
// gg-Package übernommen wurden: sie operieren auf Punkte, resp. Vektoren, die
// als Paar von Float-Werten angegeben werden.
func (m Matrix) TransformPoint(x, y float64) (float64, float64) {
    return m.M11*x + m.M12*y + m.M13, m.M21*x + m.M22*y + m.M23
}

// Während TransformPoint eine vollständige Transformation durchführt, wird
// bei TransformVector die Translation ignoriert.
func (m Matrix) TransformVector(x, y float64) (float64, float64) {
    return m.M11*x + m.M12*y, m.M21*x + m.M22*y
}

// Mit dieser Methode implementiert Matrix schliesslich das 'Stringer'
// Interface und kann bequem per Printf("%v", m) ausgegeben werden.
// Beachte, dass dieser String Zeilenumbrüche enthält!
func (m Matrix) String() (string) {
    return fmt.Sprintf("[ % .4f  % .4f  % .4f ]\n[ % .4f  % .4f  % .4f ]\n[ % .4f  % .4f  % .4f ]",
            m.M11, m.M12, m.M13,
            m.M21, m.M22, m.M23,
              0.0,   0.0,    1.0)
}

// Im Package golang.org/x/image/math/f64 sind affine Transformationen als
// Arrays von Float-Werten definiert. Für die Konvertierung in diesen 
// Typ steht die folgende Methode zur Verfügung.
func (m Matrix) AsAff3() ([6]float64) {
    return [...]float64{m.M11, m.M12, m.M13, m.M21, m.M22, m.M23}
}

func (m Matrix) AsMat3() ([9]float64) {
    return [...]float64{m.M11, m.M12, m.M13, m.M21, m.M22, m.M23, 0, 0, 1}
}
