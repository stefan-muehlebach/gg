package gg

import (
	"image/color"
	"image"
	"github.com/golang/freetype/raster"
)

type Painter struct {
    painter *raster.RGBAPainter
    changed image.Rectangle
}

func newPainter(im *image.RGBA) *Painter {
    p := &Painter{
        painter: raster.NewRGBAPainter(im),
        changed: image.Rectangle{},
    }
    return p
}

func (p *Painter) Paint(ss []raster.Span, done bool) {
    p.painter.Paint(ss, done)
    if len(ss) == 0 {
        return
    }
    if p.changed.Empty() {
        p.changed = image.Rect(ss[0].X0, ss[0].Y, ss[0].X1, ss[0].Y+1)
    }
    for _, s := range ss {
        if s.Y < p.changed.Min.Y {
            p.changed.Min.Y = s.Y
        }
        if s.Y > p.changed.Max.Y {
            p.changed.Max.Y = s.Y
        }
        if s.X0 < p.changed.Min.X {
            p.changed.Min.X = s.X0
        }
        if s.X1 > p.changed.Max.X {
            p.changed.Max.X = s.X1
        }
    }
}

func (p *Painter) SetColor(c color.Color) {
    p.painter.SetColor(c)
}

func (p *Painter) Changed() image.Rectangle {
    return p.changed
}

func (p *Painter) Clear() {
    p.changed = image.Rectangle{}
}

