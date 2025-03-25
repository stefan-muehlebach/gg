package geom

import (
	"fmt"
	"image"
)

// Datentype für Rechtecke in einem 2-dimensionalen Feld. Gespeichert werden
// die zwei Eckpunkte: Min für den Punkt links oben (d.h. mit den kleineren
// Werten für die x-, resp. y-Koordinate) und Max für den gegenüberliegenen
// Punkt. Die Felder Min und Max sind exportiert. Ein veränderndes Programm
// muss darauf achten, dass die Bedingung (Min.X <= Max.X UND Min.Y <= Max.Y)
// eingehalten wird!
type Rectangle struct {
	Min, Max Point
}

// Neues Rechteck mit den einzelnen Koordinaten für die Eckpunkte mit den
// kleinsten (Min) und grössten Werten (Max).
func NewRectangle(x0, y0, x1, y1 float64) Rectangle {
	if x0 > x1 {
		x0, x1 = x1, x0
	}
	if y0 > y1 {
		y0, y1 = y1, y0
	}
	return Rectangle{Point{x0, y0}, Point{x1, y1}}
}

func Rect(x0, y0, x1, y1 float64) Rectangle {
	return NewRectangle(x0, y0, x1, y1)
}

// Neues Rechteck mit den Koordinaten des Eckpunktes mit den kleinsten Werten,
// der Breite (d.h. Ausdehung in x-Richtung) und der Höhe (Ausdehung in
// y-Richtung).
func NewRectangleWH(x, y, w, h float64) Rectangle {
	return NewRectangle(x, y, x+w, y+h)
}

// Neues Rechteck mit den Koordinaten des Mittelpunktes, der Breite und der
// Höhe.
func NewRectangleCWH(mx, my, w, h float64) Rectangle {
	return NewRectangle(mx-w/2.0, my-h/2.0, mx+w/2.0, my+h/2.0)
}

// Erstellt ein neues Rechteck aus dem Datentyp Rectangle des [image]-Paketes.
func NewRectangleIMG(r image.Rectangle) Rectangle {
	return Rectangle{NewPointIMG(r.Min), NewPointIMG(r.Max)}
}

// Add verschiebt das Rechteck r um die Koordinaten des Punktes p.
func (r Rectangle) Add(p Point) Rectangle {
	return Rectangle{
		r.Min.Add(p),
		r.Max.Add(p),
	}
}

// Sub subtrahiert von allen Koordinaten des Rechtecks r die Werte vom Punkt p.
func (r Rectangle) Sub(p Point) Rectangle {
	return Rectangle{
		r.Min.Sub(p),
		r.Max.Sub(p),
	}
}

// Verschiebt das Rechteck r um die Koordinaten aus dem Punkt dp. Die
// Verschiebung wirkt sich direkt das Objekt aus.
func (r *Rectangle) Move(dp Point) {
	r.Min.Move(dp)
	r.Max.Move(dp)
}

// Prüft, ob das Rechteck leer ist. Im Fall, dass Min >= Max ist, gilt das
// Rechteck als leer.
func (r Rectangle) Empty() bool {
	return r.Min.X >= r.Max.X || r.Min.Y >= r.Max.Y
}

// Prüft, ob zwei Rechteck gleich sind, dh. die exakt gleichen Koordinaten
// haben.
func (r Rectangle) Eq(s Rectangle) bool {
	return r == s || r.Empty() && s.Empty()
}

// Prüft, ob sich das Rechteck r vollständig im Rechteck s befindet.
func (r Rectangle) In(s Rectangle) bool {
	if r.Empty() {
		return true
	}
	return r.Min.X >= s.Min.X && r.Max.X <= s.Max.X &&
		r.Min.Y >= s.Min.Y && r.Max.Y <= s.Max.Y
}

// Prüft, ob sich die Rechtecke r und s überlappen.
func (r Rectangle) Overlaps(s Rectangle) bool {
	return !r.Empty() && !s.Empty() &&
		r.Min.X < s.Max.X && s.Min.X < r.Max.X &&
		r.Min.Y < s.Max.Y && s.Min.Y < r.Max.Y
}

// Liefert die Breite (Ausdehung in X-Richtung) des Rechteckes.
func (r Rectangle) Dx() float64 {
	return r.Max.X - r.Min.X
}

// Liefert die Höhe (Ausdehung in Y-Richtung) des Rechteckes.
func (r Rectangle) Dy() float64 {
	return r.Max.Y - r.Min.Y
}

func (r Rectangle) Size() Point {
	return Point{r.Max.X - r.Min.X, r.Max.Y - r.Min.Y}
}

// Liefert den Mittelpunkt des Rechteckes.
func (r Rectangle) Center() Point {
	return Point{(r.Min.X + r.Max.X) / 2.0, (r.Min.Y + r.Max.Y) / 2.0}
}

// Berechnet ein neues Rechteck, welches der Schnitt der Rechtecke r und s
// ist. Das Resultat kann leer sein, falls sich r und s nicht überlappen.
func (r Rectangle) Intersect(s Rectangle) Rectangle {
	if r.Min.X < s.Min.X {
		r.Min.X = s.Min.X
	}
	if r.Min.Y < s.Min.Y {
		r.Min.Y = s.Min.Y
	}
	if r.Max.X > s.Max.X {
		r.Max.X = s.Max.X
	}
	if r.Max.Y > s.Max.Y {
		r.Max.Y = s.Max.Y
	}
	if r.Empty() {
		return Rectangle{}
	}
	return r
}

// Berechnet die Vereinigung der Rechtecke r und s. Das Resultat ist das
// kleinste Rechteck, in welchem sowohl r als auch s vollständig enthalten
// sind.
func (r Rectangle) Union(s Rectangle) Rectangle {
	if r.Empty() {
		return s
	}
	if s.Empty() {
		return r
	}
	if r.Min.X > s.Min.X {
		r.Min.X = s.Min.X
	}
	if r.Min.Y > s.Min.Y {
		r.Min.Y = s.Min.Y
	}
	if r.Max.X < s.Max.X {
		r.Max.X = s.Max.X
	}
	if r.Max.Y < s.Max.Y {
		r.Max.Y = s.Max.Y
	}
	return r
}

// Verschiebt die Ränder von r um d nach Innen und retourniert dieses neü
// Rechteck. d kann negativ sein, in diesem Fall wird das Rechteck
// vergrössert
func (r Rectangle) Inset(dx, dy float64) Rectangle {
	if r.Dx() < 2*dx {
		r.Min.X = (r.Min.X + r.Max.X) / 2.0
		r.Max.X = r.Min.X
	} else {
		r.Min.X += dx
		r.Max.X -= dx
	}
	if r.Dy() < 2*dy {
		r.Min.Y = (r.Min.Y + r.Max.Y) / 2.0
		r.Max.Y = r.Min.Y
	} else {
		r.Min.Y += dy
		r.Max.Y -= dy
	}
	return r
}

// Wozu diese Sammlung von Methoden gebraucht wird, kann ich echt nicht mehr
// sagen - wohl ein Interface, welches damit implementiert ist.

func (r Rectangle) X0() float64 { return r.Min.X }
func (r Rectangle) Y0() float64 { return r.Min.Y }
func (r Rectangle) X1() float64 { return r.Max.X }
func (r Rectangle) Y1() float64 { return r.Max.Y }

func (r Rectangle) NW() Point {
	return r.Min
}
func (r Rectangle) N() Point {
	return Point{(r.Min.X + r.Max.X) / 2.0, r.Min.Y}
}
func (r Rectangle) NE() Point {
	return Point{r.Max.X, r.Min.Y}
}
func (r Rectangle) W() Point {
	return Point{r.Min.X, (r.Min.Y + r.Max.Y) / 2.0}
}
func (r Rectangle) C() Point {
	return Point{(r.Min.X + r.Max.X) / 2.0, (r.Min.Y + r.Max.Y) / 2.0}
}
func (r Rectangle) E() Point {
	return Point{r.Max.X, (r.Min.Y + r.Max.Y) / 2.0}
}
func (r Rectangle) SW() Point {
	return Point{r.Min.X, r.Max.Y}
}
func (r Rectangle) S() Point {
	return Point{(r.Min.X + r.Max.X) / 2.0, r.Max.Y}
}
func (r Rectangle) SE() Point {
	return r.Max
}

// Liefert die relativen Positionsdaten des Punktes p gegenüber dem Rechteck r
// Also falls p=r.Min, dann wird (0.0, 0.0) retourniert; falls p=r.Max, wird
// (1.0, 1.0) retourniert, etc.
func (r Rectangle) PosRel(p Point) (float64, float64) {
	w, h := r.Max.X-r.Min.X, r.Max.Y-r.Min.Y
	p = p.Sub(r.Min)
	return p.X / w, p.Y / h
}

// Erstellt einen Punkt, dessen Koordianten relativ zum Rechteck r über die
// Grössen fx und fx bestimmt werden.
// Falls bspw. fx=1.0 und fy=0.0 sind, wird (r.Max.X, r.Min.Y) retourniert.
func (r Rectangle) RelPos(fx, fy float64) Point {
	w, h := r.Max.X-r.Min.X, r.Max.Y-r.Min.Y
	p := Point{fx * w, fy * h}.Add(r.Min)
	return p
}

// Retourniert ausgehend vom Punkt p einen neuen Punkt, der sich garantiert
// im Rechteck r befindet.
func (r Rectangle) SetInside(p Point) Point {
	q := p
	if p.X < r.Min.X {
		q.X = r.Min.X
	} else if p.X >= r.Max.X {
		q.X = r.Max.X - r.Dx()/1000.0
	}
	if p.Y < r.Min.Y {
		q.Y = r.Min.Y
	} else if p.Y >= r.Max.Y {
		q.Y = r.Max.Y - r.Dy()/1000.0
	}
	return q
}

// Kanonisiert das Rechteck r. Das heisst, dass im resultierenden Rechteck
// die Koordinaten des Punktes Min auf jeden Fall kleiner sind als die Koord.
// des Punktes Max.
func (r Rectangle) Canon() Rectangle {
	return Rectangle{Min: r.Min.Min(r.Max), Max: r.Min.Max(r.Max)}
}

// Liefert X- und Y-Koordinaten des Eckpunktes mit den kleinsten Koordinaten
// sowie die Breite und Höhe als eigenständige Werte.
func (r Rectangle) AsCoord() (x, y, w, h float64) {
	return r.Min.X, r.Min.Y, r.Dx(), r.Dy()
}

// Konvertiert das Rechteck in den Datentyp image.Rectangle aus der Standard-
// Library von Go.
func (r Rectangle) Int() image.Rectangle {
	return image.Rectangle{r.Min.Int(), r.Max.Int()}
}

// Produziert einen String mit den Eckkoordinaten des Rechtecks.
func (r Rectangle) String() string {
	return r.Min.String() + "-" + r.Max.String()
}

func (r *Rectangle) Set(s string) error {
	var x0, y0, x1, y1 float64

	_, err := fmt.Sscanf(s, "(%f;%f)-(%f;%f)", &x0, &y0, &x1, &y1)
	r.Min.X, r.Min.Y = x0, y0
	r.Max.X, r.Max.Y = x1, y1
	return err
}
