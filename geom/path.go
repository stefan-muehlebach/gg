package geom

import (
	"math"
)

const (
    numSegments = 200
)

type Segment interface {
    Start() (Point)
    End() (Point)
    Point(t float64) (Point)
    Dir(t float64) (Point)
    ArcLength() (float64)
    ArcTime(len float64) (float64)
}




type LinearSegment struct {
    P0, P1 Point
}

func (s LinearSegment) Start() (Point) {
    return s.P0
}

func (s LinearSegment) End() (Point) {
    return s.P1
}

func (s LinearSegment) Point(t float64) (Point) {
    return s.P0.Mul(1.0-t).Add(s.P1.Mul(t))
}

func (s LinearSegment) Dir(t float64) (Point) {
    return s.P1.Sub(s.P0).Norm()
}

func (s LinearSegment) ArcLength() (float64) {
    return s.P1.Sub(s.P0).Abs()
}

func (s LinearSegment) ArcTime(len float64) (float64) {
    t := len/s.ArcLength();
    if t > 1.0 {
        return 1.0
    }
    return t
}



type QuadraSegment struct {
    P0, C0, P1 Point
    lenCached float64
}




type CubicSegment struct {
    P0, C0, C1, P1 Point   
    lenCached float64 
}

func (s CubicSegment) Start() (Point) {
    return s.P0
}

func (s CubicSegment) End() (Point) {
    return s.P1
}

func (s CubicSegment) Point(t float64) (Point) {
    omt := 1.0 - t
    omt2 := omt * omt
    omt3 := omt2 * omt
    t2 := t * t
    t3 := t2 * t
    
    return s.P0.Mul(omt3).Add(s.C0.Mul(3*omt2*t)).Add(s.C1.Mul(3*omt*t2)).Add(s.P1.Mul(t3))
}

func (s CubicSegment) Dir(t float64) (Point) {
    var q1, q2, q3, q Point
    t2 := t*t
    
    q1 = s.P1.Add(s.C1.Mul(-3.0).Add(s.C0.Mul(3.0).Add(s.P0.Mul(-1)))).Mul(t2)
    q2 = s.C1.Add(s.C0.Mul(-2.0).Add(s.P0)).Mul(2.0*t)
    q3 = s.C0.Sub(s.P0)
    q  = q1.Add(q2.Add(q3)).Mul(3.0)

    return q.Norm()
}

func (s CubicSegment) ArcLength() (float64) {
    if s.lenCached <= 0.0 {
         s.updateArcLength()       
    }
    return s.lenCached
}        

func (s *CubicSegment) updateArcLength() {
    totalLength := 0.0
    for i:=0; i<numSegments; i++ {
        t0 := float64(i) / float64(numSegments)
	    t1 := float64(i+1) / float64(numSegments)
        
        p0t := s.Point(t0)
        p1t := s.Point(t1)
        totalLength += p0t.Distance(p1t)
    }
    s.lenCached = totalLength
}

func (s CubicSegment) ArcTime(len float64) (float64) {
    return 0.0
}




type Path struct {
    startPoint Point
    segms []Segment
    cyclic bool
    lenCached float64
}

func NewPath() (*Path) {
    p := &Path{}
    p.lenCached = -1
    return p
}

func (p *Path) Start() (Point) {
    if p.segms != nil {
        return p.segms[0].Start()
    } else {
        return Point{0,0}
    }
}

func (p *Path) End() (Point) {
    if p.segms != nil {
        return p.segms[len(p.segms)-1].End()
    } else {
        return Point{0,0}
    }
}

func (p *Path) Point(t float64) (Point) {
    var segmId int
    var tNew float64
    
    if t == float64(p.Segments()) {
        segmId = p.Segments()-1
        tNew = 1.0
    } else {
        segmId = int(math.Floor(t))
        tNew = t - math.Floor(t)
    }
    return p.segms[segmId].Point(tNew)
}

func (p *Path) PointNorm(t float64) (Point) {
    return p.Point(t * float64(p.Segments()))
}

func (p *Path) Dir(t float64) (Point) {
    var segmId int
    var tNew float64
    
    if t == float64(p.Segments()) {
        segmId = p.Segments()-1
        tNew = 1.0
    } else {
        segmId = int(math.Floor(t))
        tNew = t - math.Floor(t)
    }
    return p.segms[segmId].Dir(tNew)
}

func (p *Path) DirNorm(t float64) (Point) {
    return p.Dir(t * float64(p.Segments()))
}

func (p *Path) ArcLength() (float64) {
    if p.lenCached < 0.0 {
         p.updateArcLength()       
    }
    return p.lenCached
}

func (p *Path) updateArcLength() {
    length := 0.0
    for _, segm := range p.segms {
        length += segm.ArcLength()
    }
    p.lenCached = length
}

func (p *Path) ArcTime(len float64) (float64) {
    return 0.0
}



func (p *Path) Segments() (int) {
    return len(p.segms)
}

func (p *Path) Segment(i int) (Segment) {
    return p.segms[i]
}

func (p *Path) Nodes() (int) {
    if p.cyclic {
        return len(p.segms)
    } else {
        return len(p.segms)+1
    }
}

func (p *Path) Node(i int) (Point) {
    if i == p.Segments() {
        return p.End()
    } else {
        return p.segms[i].Start()
    }
}

func (p *Path) IsCyclic() (bool) {
    return p.cyclic
}

//-----------

func (p *Path) MoveTo(p0 Point) {
    p.segms = make([]Segment, 0)
    p.startPoint = p0
}

func (p *Path) LineTo(p1 Point) {
    var p0 Point
    
    if len(p.segms) == 0 {
        p0 = p.startPoint
    } else {
        p0 = p.segms[len(p.segms)-1].End()
    }
    ls := LinearSegment{p0, p1}
    p.segms = append(p.segms, ls)
}

func (p *Path) BezierTo(c0, c1, p1 Point) {
    var p0 Point
    
    if len(p.segms) == 0 {
        p0 = p.startPoint
    } else {
        p0 = p.segms[len(p.segms)-1].End()
    }
    bs := CubicSegment{P0: p0, C0: c0, C1: c1, P1: p1}
    p.segms = append(p.segms, bs)
}

func (p *Path) Close() {
    if !p.Start().Eq(p.End()) {
        p.LineTo(p.Start())
    }
    p.cyclic = true
    p.updateArcLength()
}
