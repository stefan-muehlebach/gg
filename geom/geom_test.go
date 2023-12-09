package geom

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"testing"
	"time"
)

const (
	eps = 1.0e-12
)

var (
	p0, p1, p2, p3, p4, p5, p6, p7, p8, p9 Point
	r0, r1, r2, r3, r4, r5, r6, r7, r8, r9 Rectangle
	d, t, tx, ty                           float64
	x, y, z                                float64
	points                                 []Point
	rectangles                             []Rectangle
	isIn                                   bool
    
    	rectList = []Rectangle{
		Rect(0, 0, 10, 10),
		Rect(10, 0, 20, 10),
		Rect(1, 2, 3, 4),
		Rect(4, 6, 10, 10),
		Rect(2, 3, 12, 5),
		Rect(-1, -2, 0, 0),
		Rect(-1, -2, 4, 6),
		Rect(-10, -20, 30, 40),
		Rect(8, 8, 8, 8),
		Rect(88, 88, 88, 88),
		Rect(6, 5, 4, 3),
        Rectangle{Point{-1, -1}, Point{1, 1}},
        Rectangle{Point{-1,  1}, Point{1, -1}},
        Rectangle{Point{ 1,  1}, Point{-1, -1}},
        Rectangle{Point{ 1, -1}, Point{-1, 1}},
        Rectangle{Point{-1,  0}, Point{1, 0}},
        Rectangle{Point{1,  0}, Point{-1, 0}},
        Rectangle{Point{0,  -1}, Point{0, 1}},
        Rectangle{Point{0,  1}, Point{0, -1}},
	}
)

func init() {

}

//----------------------------------------------------------------------------
//
// Tests for the Point datatype.
//

func TestPointAdd(t *testing.T) {
	p0 = Point{2.0, -2.0}
	p1 = Point{-4.0, 4.0}
	p3 = p0.Add(p1)
	p4 = Point{-2.0, 2.0}
	if p3 != p4 {
		t.Errorf("%v.Add(%v) = %v; want %v", p0, p1, p3, p4)
	}
}

func TestPointDistance(t *testing.T) {
	p0 = Point{0.0, 0.0}
	p1 = Point{1.0, 1.0}
	dr := p0.Distance(p0)
	dt := 0.0
	if dr != dt {
		t.Errorf("%v.Distance(%v) = %f; want %f", p0, p0, dr, dt)
	}
	dr = p0.Distance(p1)
	dt = math.Sqrt(2.0)
	if dr != dt {
		t.Errorf("%v.Distance(%v) = %f; want %f", p0, p1, dr, dt)
	}
}

func TestPointIn(t *testing.T) {
	r0 = Rectangle{Point{-1, -1}, Point{1, 1}}
	p0 = Point{0, 0}
	r := p0.In(r0)
	if !r {
		t.Errorf("%v.In(%v) = %v; want true", p0, r0, r)
	}
	p0 = Point{-1, 0}
	r = p0.In(r0)
	if !r {
		t.Errorf("%v.In(%v) = %v; want true", p0, r0, r)
	}
	p0 = Point{0, -1}
	r = p0.In(r0)
	if !r {
		t.Errorf("%v.In(%v) = %v; want true", p0, r0, r)
	}
	p0 = Point{1, 0}
	r = p0.In(r0)
	if r {
		t.Errorf("%v.In(%v) = %v; want false", p0, r0, r)
	}
	p0 = Point{0, 1}
	r = p0.In(r0)
	if r {
		t.Errorf("%v.In(%v) = %v; want false", p0, r0, r)
	}
}

type testCase struct {
	str  string  // Die Textdarstellung des Punktes
	ok   bool    // Ist dies ueberhaupt eine korrekte Darstellung
	pt   Point   // Die erwartete Darstellung als Point
	dist float64 // Der erwartete Betrag
}

var (
	testList = []testCase{
		{"(0;0)", true, Point{0, 0}, 0.0},
		{"(1;0)", true, Point{1, 0}, 1.0},
		{"(1;1)", true, Point{1, 1}, math.Sqrt(2)},
		{"(0;1)", true, Point{0, 1}, 1.0},
		{"(-1;1)", true, Point{-1, 1}, math.Sqrt(2)},
		{"(-1;0)", true, Point{-1, 0}, 1.0},
		{"(-1;-1)", true, Point{-1, -1}, math.Sqrt(2)},
		{"(0;-1)", true, Point{0, -1}, 1.0},
		{"(1;-1)", true, Point{1, -1}, math.Sqrt(2)},

		{"(-.5;.5)", true, Point{-0.5, 0.5}, math.Sqrt(0.5)},

		{"1;1", false, Point{}, 0.0},
		{"(1,1)", false, Point{}, 0.0},
		{"1,1)", false, Point{}, 0.0},
		{"(1,1", false, Point{}, 0.0},
	}
)

func TestPointSet(t *testing.T) {
	var err error

	for i, tst := range testList {
		err = p0.Set(tst.str)
		if tst.ok && err != nil {
			t.Errorf("%d: parsing %s failed: %v", i, tst.str, err)
			continue
		}
		if !tst.ok {
			if err == nil {
				t.Errorf("%d: parsing %s should fail!", i, tst.str)
			}
			continue
		}
		if !p0.Eq(tst.pt) {
			t.Errorf("%d: %v not equal to %v", i, p0, tst.pt)
			continue
		}
		d := p0.Abs()
		if d != tst.dist {
			t.Errorf("%d: distance to (0;0): have %f, want %f", i, d, tst.dist)
			continue
		}

		if (p0.X == 0.0) && (p0.Y == 0.0) {
			continue
		}
		pNorm := p0.Norm()
		lNorm := pNorm.Abs()
		if math.Abs(1.0-lNorm) > eps {
			t.Errorf("length of normalized vector %v should be 1.0, is %f", pNorm, lNorm)
		}
	}
}

//----------------------------------------------------------------------------
//
// Tests for the Rectangle datatype.
//

func TestRectangle(t *testing.T) {
	// in checks that every point in f is in g.
	in := func(f, g Rectangle) error {
		if !f.In(g) {
			return fmt.Errorf("f=%s, f.In(%s): got false, want true", f, g)
		}
		for y := f.Min.Y; y < f.Max.Y; y++ {
			for x := f.Min.X; x < f.Max.X; x++ {
				p := Point{x, y}
				if !p.In(g) {
					return fmt.Errorf("p=%s, p.In(%s): got false, want true", p, g)
				}
			}
		}
		return nil
	}

	// r.Eq(s) should be equivalent to every point in r being in s, and every
	// point in s being in r.
	for _, r := range rectList {
		for _, s := range rectList {
			got := r.Eq(s)
			want := in(r, s) == nil && in(s, r) == nil
			if got != want {
				t.Errorf("Eq: r=%s, s=%s: got %t, want %t", r, s, got, want)
			}
		}
	}

	// The intersection should be the largest rectangle a such that every point
	// in a is both in r and in s.
	for _, r := range rectList {
		for _, s := range rectList {
			a := r.Intersect(s)
			if err := in(a, r); err != nil {
				t.Errorf("Intersect: r=%s, s=%s, a=%s, a not in r: %v", r, s, a, err)
			}
			if err := in(a, s); err != nil {
				t.Errorf("Intersect: r=%s, s=%s, a=%s, a not in s: %v", r, s, a, err)
			}
			if isZero, overlaps := a == (Rectangle{}), r.Overlaps(s); isZero == overlaps {
				t.Errorf("Intersect: r=%s, s=%s, a=%s: isZero=%t same as overlaps=%t",
					r, s, a, isZero, overlaps)
			}
			largerThanA := [4]Rectangle{a, a, a, a}
			largerThanA[0].Min.X--
			largerThanA[1].Min.Y--
			largerThanA[2].Max.X++
			largerThanA[3].Max.Y++
			for i, b := range largerThanA {
				if b.Empty() {
					// b isn't actually larger than a.
					continue
				}
				if in(b, r) == nil && in(b, s) == nil {
					t.Errorf("Intersect: r=%s, s=%s, a=%s, b=%s, i=%d: intersection could be larger",
						r, s, a, b, i)
				}
			}
		}
	}

	// The union should be the smallest rectangle a such that every point in r
	// is in a and every point in s is in a.
	for _, r := range rectList {
		for _, s := range rectList {
			a := r.Union(s)
			if err := in(r, a); err != nil {
				t.Errorf("Union: r=%s, s=%s, a=%s, r not in a: %v", r, s, a, err)
			}
			if err := in(s, a); err != nil {
				t.Errorf("Union: r=%s, s=%s, a=%s, s not in a: %v", r, s, a, err)
			}
			if a.Empty() {
				// You can't get any smaller than a.
				continue
			}
			smallerThanA := [4]Rectangle{a, a, a, a}
			smallerThanA[0].Min.X++
			smallerThanA[1].Min.Y++
			smallerThanA[2].Max.X--
			smallerThanA[3].Max.Y--
			for i, b := range smallerThanA {
				if in(r, b) == nil && in(s, b) == nil {
					t.Errorf("Union: r=%s, s=%s, a=%s, b=%s, i=%d: union could be smaller",
						r, s, a, b, i)
				}
			}
		}
	}
}

func TestRectangleCanon(t *testing.T) {
	for _, rect := range rectList {
		r0 = rect.Canon()
		if (r0.Min.X > r0.Max.X) || (r0.Min.Y > r0.Max.Y) {
			t.Errorf("rectangle %v is not canonical", r0)
		}
	}
}

func TestRectangleRelPos(t *testing.T) {
	r0 = Rect(-3, -2, 4, 3)
	relList := [][]float64{
		{0.0, 0.0},
		{1.0, 0.0},
		{0.0, 1.0},
		{1.0, 1.0},
		{-1.0, -1.0},
		{2.0, 2.0},
		{0.5, 0.5},
		{-0.5, -0.5},
	}
	posList := []Point{
		Point{-3.0, -2.0},
		Point{4.0, -2.0},
		Point{-3.0, 3.0},
		Point{4.0, 3.0},
		Point{-10.0, -7.0},
		Point{11.0, 8.0},
		Point{0.5, 0.5},
		Point{-6.5, -4.5},
	}

	for i, rel := range relList {
		p0 = r0.RelPos(rel[0], rel[1])
		if (p0.X != posList[i].X) || (p0.Y != posList[i].Y) {
			t.Errorf("%v.RelPos(%f, %f) = %v; want %v", r0, rel[0], rel[1], p0, posList[i])
		}
	}
	for i, pos := range posList {
		rx, ry := r0.PosRel(pos)
		if (rx != relList[i][0]) || (ry != relList[i][1]) {
			t.Errorf("%v.PosRel(%v) = %f, %f; want %f, %f", r0, pos, rx, ry, relList[i][0], relList[i][1])
		}
	}
}

func TestRectangleSetInside(t *testing.T) {
	r0 = Rect(-3, -2, 4, 3)
    posList := [][]Point{
        {Point{-4, -3}, Point{-3, -2}},
        {Point{-4, -2}, Point{-3, -2}},
        {Point{-4,  3}, Point{-3,  3}},
        {Point{-4,  4}, Point{-3,  3}},
        {Point{-3, -3}, Point{-3, -2}},
        {Point{-3, -2}, Point{-3, -2}},
        {Point{-3,  3}, Point{-3,  3}},
        {Point{-3,  4}, Point{-3,  3}},
        {Point{ 0, -3}, Point{ 0, -2}},
        {Point{ 0, -2}, Point{ 0, -2}},
        {Point{ 0,  3}, Point{ 0,  3}},
        {Point{ 0,  4}, Point{ 0,  3}},
        {Point{ 4, -3}, Point{ 4, -2}},
        {Point{ 4, -2}, Point{ 4, -2}},
        {Point{ 4,  3}, Point{ 4,  3}},
        {Point{ 4,  4}, Point{ 4,  3}},
    }
    for _, pos := range posList {
        p0 = r0.SetInside(pos[0])
        if !p0.Eq(pos[1]) {
            t.Errorf("SetInside of %v failed: got %v want %v", pos[0], p0, pos[1])
        }
    }
}

//----------------------------------------------------------------------------
//
// Tests for the Path datatype.
//

// func TestPathPoint(t *testing.T) {
//     pth := NewPath()
//     p0 := Point{-1,1}
//     c0 := Point{1,1}
//     c1 := Point{-1,-1}
//     p1 := Point{1,-1}

//     pth.MoveTo(p0)
//     pth.BezierTo(c0, c1, p1)
//     pth.BezierTo(c0, c1, p0)
//     fmt.Printf("len: %d\n", pth.Segments())
//     fmt.Printf("nodes: %d\n", pth.Nodes())
//     for r:=0.0; r<=1.0; r+=0.1 {
//         pt := pth.PointNorm(r)
//         fmt.Printf("%0.4f: %4v\n", r, pt)
//     }

//     pth.Close()
//     fmt.Printf("len: %d\n", pth.Segments())
//     fmt.Printf("nodes: %d\n", pth.Nodes())
//     for r:=0.0; r<=1.0; r+=0.1 {
//         pt := pth.PointNorm(r)
//         fmt.Printf("%0.4f: %4v\n", r, pt)
//     }

// }

//----------------------------------------------------------------------------
//
// Benchmarks
//

const (
	numObjs int = 10_001
)

func BenchmarkPoint(b *testing.B) {
	var id int

	points = make([]Point, numObjs)
	rectangles = make([]Rectangle, numObjs)

	seed := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(seed)
	for i := 0; i < numObjs; i++ {
		points[i] = Point{320 * rnd.Float64(), 240 * rnd.Float64()}
		x, y := 320*rnd.Float64(), 240*rnd.Float64()
		w, h := (320-x)*rnd.Float64(), (240-y)*rnd.Float64()
		rectangles[i] = NewRectangleWH(x, y, w, h)
	}

	id = 0

	b.Run("Add", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			points[id+2] = points[id].Add(points[id+1])
		}
	})
	id += 3

	b.Run("Sub", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			points[id+2] = points[id].Sub(points[id+1])
		}
	})
	id += 3

	t = rnd.Float64()
	b.Run("Mul", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			points[id+1] = points[id].Mul(t)
		}
	})
	id += 2

	t = rnd.Float64()
	b.Run("Div", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			points[id+1] = points[id].Div(t)
		}
	})
	id += 2

	b.Run("Distance", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			d = points[id].Distance(points[id+1])
		}
	})
	id += 2

	b.Run("Dist2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			d = points[id].Dist2(points[id+1])
		}
	})
	id += 2

	t = rnd.Float64()
	b.Run("In", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			isIn = points[id].In(rectangles[id+1])
		}
	})
	id += 2

	t = rnd.Float64()
	b.Run("Interpolate", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			points[id+2] = points[id].Interpolate(points[id+1], t)
		}
	})
	id += 3

	b.Run("Move", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			points[id].Move(points[id+1])
		}
	})
	id += 2
}

func BenchmarkRectangle(b *testing.B) {
	var id int

	points = make([]Point, numObjs)
	rectangles = make([]Rectangle, numObjs)

	seed := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(seed)
	for i := 0; i < numObjs; i++ {
		points[i] = Point{320 * rnd.Float64(), 240 * rnd.Float64()}
		x, y := 320*rnd.Float64(), 240*rnd.Float64()
		w, h := (320-x)*rnd.Float64(), (240-y)*rnd.Float64()
		rectangles[i] = NewRectangleWH(x, y, w, h)
	}

	b.ReportAllocs()
	id = 0

	b.Run("Add", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			rectangles[id] = rectangles[id+1].Add(points[id+2])
		}
	})
	id += 3

	b.Run("Sub", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			rectangles[id] = rectangles[id+1].Sub(points[id+2])
		}
	})
	id += 3

	b.Run("Move", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			rectangles[id].Move(points[id+1])
		}
	})
	id += 2

	b.Run("PosRel", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			tx, ty = rectangles[id].PosRel(points[id+1])
		}
	})
	id += 2

	tx, ty = rand.Float64(), rand.Float64()
	b.Run("RelPos", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			points[id] = rectangles[id+1].RelPos(tx, ty)
		}
	})
	id += 2

	b.Run("Center", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			points[id] = rectangles[id+1].Center()
		}
	})
	id += 2

	b.Run("SetInside", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			points[id] = rectangles[id+1].SetInside(points[id+2])
		}
	})
	id += 3

	b.Run("Union", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			rectangles[id] = rectangles[id+1].Union(rectangles[id+2])
		}
	})
	id += 3

	b.Run("Intersect", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			rectangles[id] = rectangles[id+1].Intersect(rectangles[id+2])
		}
	})
	id += 3
}

func BenchmarkRandValue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		t = rand.Float64()
	}
}

func BenchmarkRandPoint(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		p0 = NewPoint(320*rand.Float64(), 240*rand.Float64())
	}
}

func BenchmarkRandRectangle(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		r0 = NewRectangle(320*rand.Float64(), 240*rand.Float64(),
			320*rand.Float64(), 240*rand.Float64())
	}
}

func BenchmarkIEEE754(b *testing.B) {
	x = rand.Float64()
	y = rand.Float64()
	b.Run("Add", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			z = x + y
		}
	})
	x = rand.Float64()
	y = rand.Float64()
	b.Run("Sub", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			z = x - y
		}
	})
	x = rand.Float64()
	y = rand.Float64()
	b.Run("Mul", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			z = x * y
		}
	})
	x = rand.Float64()
	y = rand.Float64()
	b.Run("Div", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			z = x / y
		}
	})
	x = rand.Float64()
	y = 1 / rand.Float64()
	b.Run("DivNew", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			z = x * y
		}
	})
}

func TestMain(m *testing.M) {
	// log.Printf("TestMain()")
	os.Exit(m.Run())
}
