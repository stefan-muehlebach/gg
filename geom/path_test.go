package geom

import (
	"testing"
    "fmt"
)

func TestPathPoint(t *testing.T) {
    pth := NewPath()
    p0 := Point{-1,1}
    c0 := Point{1,1}
    c1 := Point{-1,-1}
    p1 := Point{1,-1}
    
    pth.MoveTo(p0)
    pth.BezierTo(c0, c1, p1)
    pth.BezierTo(c0, c1, p0)
    fmt.Printf("len: %d\n", pth.Segments())
    fmt.Printf("nodes: %d\n", pth.Nodes())
    for r:=0.0; r<=1.0; r+=0.1 {
        pt := pth.PointNorm(r)
        fmt.Printf("%0.4f: %4v\n", r, pt)
    }
    
    pth.Close()
    fmt.Printf("len: %d\n", pth.Segments())
    fmt.Printf("nodes: %d\n", pth.Nodes())
    for r:=0.0; r<=1.0; r+=0.1 {
        pt := pth.PointNorm(r)
        fmt.Printf("%0.4f: %4v\n", r, pt)
    }
    
}