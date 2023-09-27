package gg

import (
    "fmt"
    "math"
)

type Matrix struct {
    M11, M12, M13 float64
    M21, M22, M23 float64
}

// ----------------------------------------------------------------------------

func Identity() (Matrix) {
    return Matrix{1.0, 0.0, 0.0,
                  0.0, 1.0, 0.0}
}

func Translate(dx, dy float64) (Matrix) {
    return Matrix{1.0, 0.0, dx,
                  0.0, 1.0, dy}
}

func Rotate(a float64) (Matrix) {
    s := math.Sin(a)
    c := math.Cos(a)
    return Matrix{c, -s, 0.0,
                  s,  c, 0.0}
}

func Scale(x, y float64) Matrix {
    return Matrix{x, 0.0, 0.0,
                  0.0, y, 0.0}
}

// ----------------------------------------------------------------------------

func (a Matrix) Multiply(b Matrix) Matrix {
    return Matrix{a.M11*b.M11 + a.M12*b.M21,
                  a.M11*b.M12 + a.M12*b.M22,
                  a.M11*b.M13 + a.M12*b.M23 + a.M13,

                  a.M21*b.M11 + a.M22*b.M21,
                  a.M21*b.M12 + a.M22*b.M22,
                  a.M21*b.M13 + a.M22*b.M23 + a.M23}
}

func (a Matrix) Inv() Matrix {
    det := a.M11*a.M22 - a.M12*a.M21
    return Matrix{ a.M22 / det,
                  -a.M12 / det,
                  (a.M12*a.M23 - a.M22*a.M13) / det,
        
                  -a.M21 / det,
                   a.M11 / det,
                  (a.M21*a.M13 - a.M11*a.M23) / det}
}

// ----------------------------------------------------------------------------

func (m Matrix) Translate(dx, dy float64) Matrix {
    return m.Multiply(Translate(dx, dy))
}

func (m Matrix) Rotate(angle float64) Matrix {
    return m.Multiply(Rotate(angle))
}

func (m Matrix) Scale(x, y float64) Matrix {
    return m.Multiply(Scale(x, y))
}

// ----------------------------------------------------------------------------

func (m Matrix) TransformVector(x, y float64) (tx, ty float64) {
    return m.M11*x + m.M12*y, m.M21*x + m.M22*y
}

func (m Matrix) TransformPoint(x, y float64) (tx, ty float64) {
    return m.M11*x + m.M12*y + m.M13, m.M21*x + m.M22*y + m.M23
}

func (m Matrix) String() (string) {
    return fmt.Sprintf("[ % .4f  % .4f  % .4f ]\n[ % .4f  % .4f  % .4f ]\n[ % .4f  % .4f  % .4f ]",
            m.M11, m.M12, m.M13,
            m.M21, m.M22, m.M23,
              0.0,   0.0, 1.0)
}

