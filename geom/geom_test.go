package geom

import (
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
    d, t, tx, ty float64
    x, y, z float64
    points []Point
    rectangles []Rectangle
    isIn bool
)

func init() {

}

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
    str string         // Die Textdarstellung des Punktes
    ok bool            // Ist dies ueberhaupt eine korrekte Darstellung
    pt Point           // Die erwartete Darstellung als Point
    dist float64       // Der erwartete Betrag
}

var (
    testList = []testCase{
        {"(0;0)",   true, Point{0, 0}, 0.0},
        {"(1;0)",   true, Point{1, 0}, 1.0},
        {"(1;1)",   true, Point{1, 1}, math.Sqrt(2)},
        {"(0;1)",   true, Point{0, 1}, 1.0},
        {"(-1;1)",  true, Point{-1, 1}, math.Sqrt(2)},
        {"(-1;0)",  true, Point{-1, 0}, 1.0},
        {"(-1;-1)", true, Point{-1, -1}, math.Sqrt(2)},
        {"(0;-1)",  true, Point{0, -1}, 1.0},
        {"(1;-1)",  true, Point{1, -1}, math.Sqrt(2)},

        {"(-.5;.5)", true, Point{-0.5, 0.5}, math.Sqrt(0.5)},

        {"1;1",     false, Point{}, 0.0},
        {"(1,1)",   false, Point{}, 0.0},
        {"1,1)",    false, Point{}, 0.0},
        {"(1,1",    false, Point{}, 0.0},
   
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
        
        pNorm := p0.Norm()
        lNorm := pNorm.Abs()
        if math.Abs(1.0-lNorm) > eps {
            t.Errorf("%d: length of normalized vector should be 1.0, is %f", i, lNorm)
        }
    }
}

/*
func TestPosRel(t *testing.T) {
    r0 = Rectangle{Point{-1, -1}, Point{1, 1}}
    
    for _, tc := range testList {
        tx, ty := r0.PosRel(tc.p)
        if tx != tc.t[0] || ty != tc.t[1] {
            t.Errorf("%v.PosRel(%v) = %f, %f; want %v", r0, tc.p, tx, ty, tc.t)
        }
    }
}

func TestRelPos(t *testing.T) {
    r0 = Rectangle{Point{-1, -1}, Point{1, 1}}
    
    for _, tc := range testList {
        p := r0.RelPos(tc.t[0], tc.t[1])
        if ! p.Eq(tc.p) {
            t.Errorf("%v.RelPos(%f, %f) = %v; want %v", r0, tc.t[0], tc.t[1],
                p, tc.p)
        }
    }
}
*/

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
        points[i] = Point{320*rnd.Float64(), 240*rnd.Float64()}
        x, y := 320*rnd.Float64(), 240*rnd.Float64()
        w, h := (320-x)*rnd.Float64(), (240-y)*rnd.Float64()
        rectangles[i] = NewRectangleWH(x, y, w, h)
    }

    id = 0

    b.Run("Add", func(b *testing.B) {
        for i:=0; i<b.N; i++ {
            points[id+2] = points[id].Add(points[id+1])
        }
    })
    id += 3

    b.Run("Sub", func(b *testing.B) {
        for i:=0; i<b.N; i++ {
            points[id+2] = points[id].Sub(points[id+1])
        }
    })
    id += 3

    t = rnd.Float64()
    b.Run("Mul", func(b *testing.B) {
        for i:=0; i<b.N; i++ {
            points[id+1] = points[id].Mul(t)
        }
    })
    id += 2

    t = rnd.Float64()
    b.Run("Div", func(b *testing.B) {
        for i:=0; i<b.N; i++ {
            points[id+1] = points[id].Div(t)
        }
    })
    id += 2

    b.Run("Distance", func(b *testing.B) {
        for i:=0; i<b.N; i++ {
            d = points[id].Distance(points[id+1])
        }
    })
    id += 2

    b.Run("Dist2", func(b *testing.B) {
        for i:=0; i<b.N; i++ {
            d = points[id].Dist2(points[id+1])
        }
    })
    id += 2

    t = rnd.Float64()
    b.Run("In", func(b *testing.B) {
        for i:=0; i<b.N; i++ {
            isIn = points[id].In(rectangles[id+1])
        }
    })
    id += 2

    t = rnd.Float64()
    b.Run("Interpolate", func(b *testing.B) {
        for i:=0; i<b.N; i++ {
            points[id+2] = points[id].Interpolate(points[id+1], t)
        }
    })
    id += 3

    b.Run("Move", func(b *testing.B) {
        for i:=0; i<b.N; i++ {
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
        points[i] = Point{320*rnd.Float64(), 240*rnd.Float64()}
        x, y := 320*rnd.Float64(), 240*rnd.Float64()
        w, h := (320-x)*rnd.Float64(), (240-y)*rnd.Float64()
        rectangles[i] = NewRectangleWH(x, y, w, h)
    }

    b.ReportAllocs()
    id = 0

    b.Run("Add", func(b *testing.B) {
        for i:=0; i<b.N; i++ {
            rectangles[id] = rectangles[id+1].Add(points[id+2])
        }
    })
    id += 3

    b.Run("Sub", func(b *testing.B) {
        for i:=0; i<b.N; i++ {
            rectangles[id] = rectangles[id+1].Sub(points[id+2])
        }
    })
    id += 3

    b.Run("Move", func(b *testing.B) {
        for i:=0; i<b.N; i++ {
            rectangles[id].Move(points[id+1])
        }
    })
    id += 2

    b.Run("PosRel", func(b *testing.B) {
        for i:=0; i<b.N; i++ {
            tx, ty = rectangles[id].PosRel(points[id+1])
        }
    })
    id += 2

    tx, ty = rand.Float64(), rand.Float64()
    b.Run("RelPos", func(b *testing.B) {
        for i:=0; i<b.N; i++ {
            points[id] = rectangles[id+1].RelPos(tx, ty)
        }
    })
    id += 2

    b.Run("Center", func(b *testing.B) {
        for i:=0; i<b.N; i++ {
            points[id] = rectangles[id+1].Center()
        }
    })
    id += 2

    b.Run("SetInside", func(b *testing.B) {
        for i:=0; i<b.N; i++ {
            points[id] = rectangles[id+1].SetInside(points[id+2])
        }
    })
    id += 3

    b.Run("Union", func(b *testing.B) {
        for i:=0; i<b.N; i++ {
            rectangles[id] = rectangles[id+1].Union(rectangles[id+2])
        }
    })
    id += 3

    b.Run("Intersect", func(b *testing.B) {
        for i:=0; i<b.N; i++ {
            rectangles[id] = rectangles[id+1].Intersect(rectangles[id+2])
        }
    })
    id += 3
}

func BenchmarkRandValue(b *testing.B) {
    for i:=0; i<b.N; i++ {
        t = rand.Float64()
    }
}

func BenchmarkRandPoint(b *testing.B) {
    b.ReportAllocs()
    for i:=0; i<b.N; i++ {
        p0 = NewPoint(320*rand.Float64(), 240*rand.Float64())
    }
}

func BenchmarkRandRectangle(b *testing.B) {
    b.ReportAllocs()
    for i:=0; i<b.N; i++ {
        r0 = NewRectangle(320*rand.Float64(), 240*rand.Float64(),
            320*rand.Float64(), 240*rand.Float64())
    }
}

func BenchmarkIEEE754(b *testing.B) {
    x = rand.Float64()
    y = rand.Float64()
    b.Run("Add", func(b *testing.B) {
        for i:=0; i<b.N; i++ {
            z = x + y
        }
    })
    x = rand.Float64()
    y = rand.Float64()
    b.Run("Sub", func(b *testing.B) {
        for i:=0; i<b.N; i++ {
            z = x - y
        }
    })
    x = rand.Float64()
    y = rand.Float64()
    b.Run("Mul", func(b *testing.B) {
        for i:=0; i<b.N; i++ {
            z = x * y
        }
    })
    x = rand.Float64()
    y = rand.Float64()
    b.Run("Div", func(b *testing.B) {
        for i:=0; i<b.N; i++ {
            z = x / y
        }
    })
    x = rand.Float64()
    y = 1/rand.Float64()
    b.Run("DivNew", func(b *testing.B) {
        for i:=0; i<b.N; i++ {
            z = x * y
        }
    })
}

func TestMain(m *testing.M) {
    // log.Printf("TestMain()")
    os.Exit(m.Run())
}

