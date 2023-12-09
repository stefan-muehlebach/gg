package geom

import (
	"fmt"
	"golang.org/x/image/math/fixed"
	"image"
	"math"
)

// Der Datentyp Point ist für Koordinaten in einem Koordinatensystem mit
// Fliesskommawerten gedacht. Dass der Typ Point und nicht Coord oder Vector
// heisst, hat eher historische Gründe.
type Point struct {
	X, Y float64
}

// Erstellt einen neuen Punkt mit den angebenen X- und Y-Koordianten.
func NewPoint(x, y float64) Point {
	return Point{x, y}
}

// Erstellt einen neuen Punkt aus den Koordianten des angegebenen Punktes
// aus dem [image]-Paket.
func NewPointIMG(p image.Point) Point {
	return Point{float64(p.X), float64(p.Y)}
}

// Addiert die jeweiligen X- und Y-Koordianten der Punkte p und q.
func (p Point) Add(q Point) Point {
	return Point{p.X + q.X, p.Y + q.Y}
}

func (p Point) AddXY(x, y float64) Point {
	return Point{p.X + x, p.Y + y}
}

// Subtrahiert vom Punkt q die jeweiligen X- und Y-Koordinaten des Punktes q.
func (p Point) Sub(q Point) Point {
	return Point{p.X - q.X, p.Y - q.Y}
}

func (p Point) SubXY(x, y float64) Point {
	return Point{p.X - x, p.Y - y}
}

// Streckt die Koordinaten des Punktes p um den Faktor k.
func (p Point) Mul(k float64) Point {
	return Point{p.X * k, p.Y * k}
}

// Dividiert die Koordinaten des Punktes p durch den Wert k.
// Eliminiert, kann durch Mul abgedeckt werden!
func (p Point) Div(k float64) Point {
	return Point{p.X / k, p.Y / k}
}

// Move verschiebt einen Punkt um die X- und Y-Werte des Punktes dp. Diese
// Verschiebung wirkt sich auf das Objekt selber aus!
func (p *Point) Move(dp Point) {
	p.X += dp.X
	p.Y += dp.Y
}

func (p Point) Neg() Point {
	return Point{-p.X, -p.Y}
}

// Prüft, ob p und q exakt die gleichen Koordinaten haben.
func (p Point) Eq(q Point) bool {
	return p == q
}

// Prüft, ob der Punkt p im Rechteck r liegt.
func (p Point) In(r Rectangle) bool {
	return r.Min.X <= p.X && p.X < r.Max.X &&
		r.Min.Y <= p.Y && p.Y < r.Max.Y
}

// Berechnet den euklidischen Abstand zwischen den Punkten p und q.
// Will man nur Abstände vergleichen, ist es oft effizienter, die Methode
// Dist2 zu verwenden, welche den quadrierten Abstand ermittelt (also ohne
// Quadratwurzel) und damit schneller ist.
func (p Point) Distance(q Point) float64 {
	return math.Hypot(p.X-q.X, p.Y-q.Y)
}

// Berechnet den quadrierten Abstand zwischen den Punkten p und q.
// Siehe auch den kommentar bei Distance.
func (p Point) Dist2(q Point) float64 {
	return (p.X-q.X)*(p.X-q.X) + (p.Y-q.Y)*(p.Y-q.Y)
}

func (p Point) Abs() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

func (p Point) Norm() Point {
	l := p.Abs()
	if l == 0.0 {
		return p
	} else {
		return p.Div(l)
	}
}

func (p Point) Angle() float64 {
	return math.Atan2(p.Y, p.X)
}

// Berechnet einen neuen Punkt, der linear zwischen p und q liegt.
// Es gilt:
//   - t=0.0: das Resultat ist p
//   - t=1.0: das Resultat ist q
//
// t kann auch ausserhalb des Intervalls [0,1] liegen. In diesem Fall erhalten
// wir Punkte, welche nicht auf der Strecke zwischen p und q liegen aber auf
// der Geraden durch p und q.
func (p Point) Interpolate(q Point, t float64) Point {
	return p.Mul(1.0 - t).Add(q.Mul(t))
}

// Vergleicht die X- sowie die Y-Werte der Punkte p und q und retourniert einen
// neuen Punkt mit den jeweils grösseren Werten.
func (p Point) Max(q Point) Point {
	return Point{max(p.X, q.X), max(p.Y, q.Y)}
}

// Vergleicht die X- sowie die Y-Werte der Punkte p und q und retourniert einen
// neuen Punkt mit den jeweils kleineren Werten.
func (p Point) Min(q Point) Point {
	return Point{min(p.X, q.X), min(p.Y, q.Y)}
}

// Liefert die X- und Y-Koordinate als einzelne Werte zurück. Gut zu verwenden
// in Funktionen, welche die Koordinaten als separate Werte erwarten.
func (p Point) AsCoord() (x, y float64) {
	return p.X, p.Y
}

// Konvertiert den Punkt in einen Datentyp aus dem [image]-Package.
func (p Point) Int() image.Point {
	return image.Pt(int(p.X), int(p.Y))
}

// Konvertiert den Punkt in einen Datentyp aus dem [fixed]-Package.
func (p Point) Fixed() fixed.Point26_6 {
	return fixp(p.X, p.Y)
}

// Gibt die Koordinaten des Punktes in der Form '(x; y)' zurück. Implementiert
// das Stringer-Interface.
func (p Point) String() string {
	return fmt.Sprintf("(%.4f; %.4f)", p.X, p.Y)
}

// Damit können Punkte (resp. die Koordinaten dazu) auch über ein Textfile
// oder die Kommandozeile eingelesen werden. Mit String zusammen implementiert
// Point somit das Getter-Interface aus flag.
func (p *Point) Set(s string) error {
	var x, y float64

	_, err := fmt.Sscanf(s, "(%f;%f)", &x, &y)
	p.X, p.Y = x, y
	return err
}

//----------------------------------------------------------------------------

func max(a, b float64) float64 {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b float64) float64 {
	if a < b {
		return a
	} else {
		return b
	}
}

func fixp(x, y float64) fixed.Point26_6 {
	return fixed.Point26_6{X: fix(x), Y: fix(y)}
}

func fix(x float64) fixed.Int26_6 {
	return fixed.Int26_6(math.Round(x * 64))
}
