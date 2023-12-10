// Eine Sammlung von Funktionen um Pixel-, resp. Rasterbilder zu erstellen.
package gg

import (
	"errors"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"io"
	_ "log"
	"math"
	"strings"

	"github.com/stefan-muehlebach/gg/geom"

	"github.com/golang/freetype/raster"
	"golang.org/x/image/draw"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/f64"
)

type LineCap int

const (
	LineCapRound LineCap = iota
	LineCapButt
	LineCapSquare
)

type LineJoin int

const (
	LineJoinRound LineJoin = iota
	LineJoinBevel
)

type FillRule int

const (
	FillRuleWinding FillRule = iota
	FillRuleEvenOdd
)

type Align int

const (
	AlignLeft Align = iota
	AlignCenter
	AlignRight
)

var (
	defaultFillStyle   = NewSolidPattern(color.White)
	defaultStrokeStyle = NewSolidPattern(color.Black)
)

type Context struct {
	width         int
	height        int
	bounds        geom.Rectangle
	rasterizer    *raster.Rasterizer
	im            *image.RGBA
	mask          *image.Alpha
	path          raster.Path
	strokePattern Pattern
	fillPattern   Pattern
	fillRule      FillRule
	start         geom.Point
	current       geom.Point
	hasCurrent    bool
	dashes        []float64
	dashOffset    float64
	lineWidth     float64
	lineCap       LineCap
	lineJoin      LineJoin
	fontFace      font.Face
	fontHeight    float64
	matrix        geom.Matrix
	stack         []*Context
}

// NewContext creates a new image.RGBA with the specified width and height
// and prepares a context for rendering onto that image.
func NewContext(width, height int) *Context {
	return NewContextForRGBA(image.NewRGBA(image.Rect(0, 0, width, height)))
}

// NewContextForImage copies the specified image into a new image.RGBA
// and prepares a context for rendering onto that image.
func NewContextForImage(im image.Image) *Context {
	return NewContextForRGBA(imageToRGBA(im))
}

// NewContextForRGBA prepares a context for rendering onto the specified image.
// No copy is made.
func NewContextForRGBA(im *image.RGBA) *Context {
	w := im.Bounds().Size().X
	h := im.Bounds().Size().Y
	return &Context{
		width:         w,
		height:        h,
		bounds:        geom.Rectangle{Max: geom.Point{X: float64(w), Y: float64(h)}},
		rasterizer:    raster.NewRasterizer(w, h),
		im:            im,
		fillPattern:   defaultFillStyle,
		strokePattern: defaultStrokeStyle,
		lineWidth:     1,
		fillRule:      FillRuleWinding,
		fontFace:      basicfont.Face7x13,
		fontHeight:    13,
		matrix:        geom.Identity(),
	}
}

// GetCurrentPoint will return the current point and if there is a current point.
// The point will have been transformed by the context's transformation matrix.
func (dc *Context) GetCurrentPoint() (geom.Point, bool) {
	if dc.hasCurrent {
		return dc.current, true
	}
	return geom.Point{}, false
}

// Image returns the image that has been drawn by this context.
func (dc *Context) Image() image.Image {
	return dc.im
}

// Width returns the width of the image in pixels.
func (dc *Context) Width() int {
	return dc.width
}

// Height returns the height of the image in pixels.
func (dc *Context) Height() int {
	return dc.height
}

// Bounds returns the coordinates of the visible range as a rectangle.
func (dc *Context) Bounds() geom.Rectangle {
	return dc.bounds
}

// SavePNG codiert das Bild als PNG Daten und schreibt diese in die Datei
// path.
func (dc *Context) SavePNG(path string) error {
	return SavePNG(path, dc.im)
}

// SaveJPG codiert das Bild as JPEG Daten und schreibt diese in die Datei
// path. Mit quality kann die Kompressionsrate vorgegeben werden.
func (dc *Context) SaveJPG(path string, quality int) error {
	return SaveJPG(path, dc.im, quality)
}

// EncodePNG codiert das Bild als PNG Daten und schreibt sie über den
// angegebenen io.Writer.
func (dc *Context) EncodePNG(w io.Writer) error {
	return png.Encode(w, dc.im)
}

// EncodeJPG encodes the image as a JPG and writes it to the provided io.Writer
// in JPEG 4:2:0 baseline format with the given options.
// Default parameters are used if a nil *jpeg.Options is passed.
func (dc *Context) EncodeJPG(w io.Writer, o *jpeg.Options) error {
	return jpeg.Encode(w, dc.im, o)
}

// SetDash sets the current dash pattern to use. Call with zero arguments to
// disable dashes. The values specify the lengths of each dash, with
// alternating on and off lengths.
func (dc *Context) SetDash(dashes ...float64) {
	dc.dashes = dashes
}

// SetDashOffset sets the initial offset into the dash pattern to use when
// stroking dashed paths.
func (dc *Context) SetDashOffset(offset float64) {
	dc.dashOffset = offset
}

// Legt die Breite fest, mit der ein Pfad gezeichnet werden soll.
func (dc *Context) SetStrokeWidth(lineWidth float64) {
	dc.lineWidth = lineWidth
}

func (dc *Context) SetLineCap(lineCap LineCap) {
	dc.lineCap = lineCap
}

func (dc *Context) SetLineCapRound() {
	dc.lineCap = LineCapRound
}

func (dc *Context) SetLineCapButt() {
	dc.lineCap = LineCapButt
}

func (dc *Context) SetLineCapSquare() {
	dc.lineCap = LineCapSquare
}

func (dc *Context) SetLineJoin(lineJoin LineJoin) {
	dc.lineJoin = lineJoin
}

func (dc *Context) SetLineJoinRound() {
	dc.lineJoin = LineJoinRound
}

func (dc *Context) SetLineJoinBevel() {
	dc.lineJoin = LineJoinBevel
}

func (dc *Context) SetFillRule(fillRule FillRule) {
	dc.fillRule = fillRule
}

func (dc *Context) SetFillRuleWinding() {
	dc.fillRule = FillRuleWinding
}

func (dc *Context) SetFillRuleEvenOdd() {
	dc.fillRule = FillRuleEvenOdd
}

// Farbfunktionen

func (dc *Context) SetFillColor(c color.Color) {
	dc.fillPattern = NewSolidPattern(c)
}

func (dc *Context) SetStrokeColor(c color.Color) {
	dc.strokePattern = NewSolidPattern(c)
}

// SetFillStyle sets current fill style
func (dc *Context) SetFillStyle(pattern Pattern) {
	dc.fillPattern = pattern
}

// SetStrokeStyle sets current stroke style
func (dc *Context) SetStrokeStyle(pattern Pattern) {
	dc.strokePattern = pattern
}

// Path Manipulation

// MoveTo starts a new subpath within the current path starting at the
// specified point.
func (dc *Context) MoveTo(x, y float64) {
	x, y = dc.TransformPoint(x, y)
	p := geom.Point{X: x, Y: y}
	dc.path.Start(p.Fixed())
	dc.start = p
	dc.current = p
	dc.hasCurrent = true
}

// LineTo adds a line segment to the current path starting at the current
// point. If there is no current point, it is equivalent to MoveTo(x, y)
func (dc *Context) LineTo(x, y float64) {
	if !dc.hasCurrent {
		dc.MoveTo(x, y)
	} else {
		x, y = dc.TransformPoint(x, y)
		p := geom.Point{X: x, Y: y}
		dc.path.Add1(p.Fixed())
		dc.current = p
	}
}

// QuadraticTo adds a quadratic bezier curve to the current path starting at
// the current point. If there is no current point, it first performs
// MoveTo(x1, y1)
func (dc *Context) QuadraticTo(x1, y1, x2, y2 float64) {
	if !dc.hasCurrent {
		dc.MoveTo(x1, y1)
	}
	x1, y1 = dc.TransformPoint(x1, y1)
	x2, y2 = dc.TransformPoint(x2, y2)
	p1 := geom.Point{X: x1, Y: y1}
	p2 := geom.Point{X: x2, Y: y2}
	dc.path.Add2(p1.Fixed(), p2.Fixed())
	dc.current = p2
}

// CubicTo adds a cubic bezier curve to the current path starting at the
// current point. If there is no current point, it first performs
// MoveTo(x1, y1). Because freetype/raster does not support cubic beziers,
// this is emulated with many small line segments.
func (dc *Context) CubicTo(x1, y1, x2, y2, x3, y3 float64) {
	if !dc.hasCurrent {
		dc.MoveTo(x1, y1)
	}
	x0, y0 := dc.current.X, dc.current.Y
	x1, y1 = dc.TransformPoint(x1, y1)
	x2, y2 = dc.TransformPoint(x2, y2)
	x3, y3 = dc.TransformPoint(x3, y3)
	pts := CubicBezier(x0, y0, x1, y1, x2, y2, x3, y3)
	for _, pt := range pts[1:] {
		f := pt.Fixed()
		dc.path.Add1(f)
		dc.current = pt
	}
}

// func (dc *Context) CubicTo(x1, y1, x2, y2, x3, y3 float64) {
//     if !dc.hasCurrent {
//         dc.MoveTo(x1, y1)
//     }
//     x0, y0 := dc.current.X, dc.current.Y
//     x1, y1 = dc.TransformPoint(x1, y1)
//     x2, y2 = dc.TransformPoint(x2, y2)
//     x3, y3 = dc.TransformPoint(x3, y3)
//     points := CubicBezier(x0, y0, x1, y1, x2, y2, x3, y3)
//     previous := dc.current.Fixed()
//     for _, p := range points[1:] {
//         f := p.Fixed()
//         if f == previous {
//             // TODO: this fixes some rendering issues but not all
//             continue
//         }
//         previous = f
//         dc.path.Add1(f)
//         // dc.strokePath.Add1(f)
//         // dc.fillPath.Add1(f)
//         dc.current = p
//     }
// }

// ClosePath adds a line segment from the current point to the beginning
// of the current subpath. If there is no current point, this is a no-op.
func (dc *Context) ClosePath() {
	if dc.hasCurrent {
		dc.path.Add1(dc.start.Fixed())
		dc.current = dc.start
	}
}

// ClearPath clears the current path. There is no current point after this
// operation.
func (dc *Context) ClearPath() {
	dc.path.Clear()
	dc.hasCurrent = false
}

// NewSubPath starts a new subpath within the current path. There is no current
// point after this operation.
func (dc *Context) NewSubPath() {
	dc.hasCurrent = false
}

// Path Drawing

func (dc *Context) capper() raster.Capper {
	switch dc.lineCap {
	case LineCapButt:
		return raster.ButtCapper
	case LineCapRound:
		return raster.RoundCapper
	case LineCapSquare:
		return raster.SquareCapper
	}
	return nil
}

func (dc *Context) joiner() raster.Joiner {
	switch dc.lineJoin {
	case LineJoinBevel:
		return raster.BevelJoiner
	case LineJoinRound:
		return raster.RoundJoiner
	}
	return nil
}

func (dc *Context) stroke(painter raster.Painter) {
	path := dc.path
	if len(dc.dashes) > 0 {
		path = dashed(path, dc.dashes, dc.dashOffset)
	} else {
		// TODO: this is a temporary workaround to remove tiny segments
		// that result in rendering issues
		// path = rasterPath(flattenPath(path))
	}
	r := dc.rasterizer
	r.UseNonZeroWinding = true
	r.Clear()
	r.AddStroke(path, fix(dc.lineWidth), dc.capper(), dc.joiner())
	r.Rasterize(painter)
}

func (dc *Context) fill(painter raster.Painter) {
	path := dc.path
	if dc.hasCurrent {
		path = make(raster.Path, len(dc.path))
		copy(path, dc.path)
		path.Add1(dc.start.Fixed())
	}
	r := dc.rasterizer
	r.UseNonZeroWinding = dc.fillRule == FillRuleWinding
	r.Clear()
	r.AddPath(path)
	r.Rasterize(painter)
}

// StrokePreserve strokes the current path with the current color, line width,
// line cap, line join and dash settings. The path is preserved after this
// operation.
func (dc *Context) StrokePreserve() {
	var painter raster.Painter
	if dc.mask == nil {
		if pattern, ok := dc.strokePattern.(*solidPattern); ok {
			// with a nil mask and a solid color pattern, we can be more efficient
			// TODO: refactor so we don't have to do this type assertion stuff?
			p := raster.NewRGBAPainter(dc.im)
			p.SetColor(pattern.color)
			painter = p
		}
	}
	if painter == nil {
		painter = newPatternPainter(dc.im, dc.mask, dc.strokePattern)
	}
	dc.stroke(painter)
}

// Stroke strokes the current path with the current color, line width,
// line cap, line join and dash settings. The path is cleared after this
// operation.
func (dc *Context) Stroke() {
	dc.StrokePreserve()
	dc.ClearPath()
}

// FillPreserve fills the current path with the current color. Open subpaths
// are implicity closed. The path is preserved after this operation.
func (dc *Context) FillPreserve() {
	var painter raster.Painter
	if dc.mask == nil {
		if pattern, ok := dc.fillPattern.(*solidPattern); ok {
			// with a nil mask and a solid color pattern, we can be more efficient
			// TODO: refactor so we don't have to do this type assertion stuff?
			p := raster.NewRGBAPainter(dc.im)
			p.SetColor(pattern.color)
			painter = p
		}
	}
	if painter == nil {
		painter = newPatternPainter(dc.im, dc.mask, dc.fillPattern)
	}
	dc.fill(painter)
}

// Fill fills the current path with the current color. Open subpaths
// are implicity closed. The path is cleared after this operation.
func (dc *Context) Fill() {
	dc.FillPreserve()
	dc.ClearPath()
}

// FillStroke is a convenient functions which does a FillPreserve first and
// a Stroke afterwards. See also [StrokeFill].
func (dc *Context) FillStroke() {
	dc.FillPreserve()
	dc.Stroke()
}

// StrokeFill is a convenient functions which does a StrokePreserve first and
// a Fill afterwards. See also [FillStroke].
func (dc *Context) StrokeFill() {
	dc.StrokePreserve()
	dc.Fill()
}

// ClipPreserve updates the clipping region by intersecting the current
// clipping region with the current path as it would be filled by dc.Fill().
// The path is preserved after this operation.
func (dc *Context) ClipPreserve() {
	clip := image.NewAlpha(image.Rect(0, 0, dc.width, dc.height))
	painter := raster.NewAlphaOverPainter(clip)
	dc.fill(painter)
	if dc.mask == nil {
		dc.mask = clip
	} else {
		mask := image.NewAlpha(image.Rect(0, 0, dc.width, dc.height))
		draw.DrawMask(mask, mask.Bounds(), clip, image.Point{}, dc.mask, image.Point{}, draw.Over)
		dc.mask = mask
	}
}

// SetMask allows you to directly set the *image.Alpha to be used as a clipping
// mask. It must be the same size as the context, else an error is returned
// and the mask is unchanged.
func (dc *Context) SetMask(mask *image.Alpha) error {
	if mask.Bounds().Size() != dc.im.Bounds().Size() {
		return errors.New("mask size must match context size")
	}
	dc.mask = mask
	return nil
}

// AsMask returns an *image.Alpha representing the alpha channel of this
// context. This can be useful for advanced clipping operations where you first
// render the mask geometry and then use it as a mask.
func (dc *Context) AsMask() *image.Alpha {
	mask := image.NewAlpha(dc.im.Bounds())
	draw.Draw(mask, dc.im.Bounds(), dc.im, image.Point{}, draw.Src)
	return mask
}

func (dc *Context) Mask() *image.Alpha {
	return dc.mask
}

// InvertMask inverts the alpha values in the current clipping mask such that
// a fully transparent region becomes fully opaque and vice versa.
func (dc *Context) InvertMask() {
	if dc.mask == nil {
		dc.mask = image.NewAlpha(dc.im.Bounds())
	} else {
		for i, a := range dc.mask.Pix {
			dc.mask.Pix[i] = 255 - a
		}
	}
}

// Clip updates the clipping region by intersecting the current
// clipping region with the current path as it would be filled by dc.Fill().
// The path is cleared after this operation.
func (dc *Context) Clip() {
	dc.ClipPreserve()
	dc.ClearPath()
}

// ResetClip clears the clipping region.
func (dc *Context) ResetClip() {
	dc.mask = nil
}

// Convenient Drawing Functions

// Clear fills the entire image with the current color.
func (dc *Context) Clear() {
	src := image.NewUniform(dc.fillPattern.ColorAt(0, 0))
	draw.Draw(dc.im, dc.im.Bounds(), src, image.Point{}, draw.Src)
}

// SetPixel sets the color of the specified pixel using the current color.
func (dc *Context) SetPixel(x, y int, c color.Color) {
	dc.im.Set(x, y, c)
}

// DrawPoint is like DrawCircle but ensures that a circle of the specified
// size is drawn regardless of the current transformation matrix. The position
// is still transformed, but not the shape of the point.
func (dc *Context) DrawPoint(x, y, r float64) {
	dc.Push()
	tx, ty := dc.TransformPoint(x, y)
	dc.matrix = geom.Identity()
	dc.DrawCircle(tx, ty, r)
	dc.Pop()
}

func (dc *Context) DrawLine(x1, y1, x2, y2 float64) {
	dc.MoveTo(x1, y1)
	dc.LineTo(x2, y2)
}

func (dc *Context) DrawRectangle(x, y, w, h float64) {
	dc.NewSubPath()
	dc.MoveTo(x, y)
	dc.LineTo(x+w, y)
	dc.LineTo(x+w, y+h)
	dc.LineTo(x, y+h)
	dc.ClosePath()
}

func (dc *Context) DrawRoundedRectangle(x, y, w, h, r float64) {
	x0, x1, x2, x3 := x, x+r, x+w-r, x+w
	y0, y1, y2, y3 := y, y+r, y+h-r, y+h
	dc.NewSubPath()
	dc.MoveTo(x1, y0)
	dc.LineTo(x2, y0)
	dc.DrawArc(x2, y1, r, Radians(270), Radians(360))
	dc.LineTo(x3, y2)
	dc.DrawArc(x2, y2, r, Radians(0), Radians(90))
	dc.LineTo(x1, y3)
	dc.DrawArc(x1, y2, r, Radians(90), Radians(180))
	dc.LineTo(x0, y1)
	dc.DrawArc(x1, y1, r, Radians(180), Radians(270))
	dc.ClosePath()
}

func (dc *Context) DrawEllipticalArc(x, y, rx, ry, angle1, angle2 float64) {
	const n = 16
	for i := 0; i < n; i++ {
		p1 := float64(i+0) / n
		p2 := float64(i+1) / n
		a1 := angle1 + (angle2-angle1)*p1
		a2 := angle1 + (angle2-angle1)*p2
		x0 := x + rx*math.Cos(a1)
		y0 := y + ry*math.Sin(a1)
		x1 := x + rx*math.Cos((a1+a2)/2)
		y1 := y + ry*math.Sin((a1+a2)/2)
		x2 := x + rx*math.Cos(a2)
		y2 := y + ry*math.Sin(a2)
		cx := 2*x1 - x0/2 - x2/2
		cy := 2*y1 - y0/2 - y2/2
		if i == 0 {
			if dc.hasCurrent {
				dc.LineTo(x0, y0)
			} else {
				dc.MoveTo(x0, y0)
			}
		}
		dc.QuadraticTo(cx, cy, x2, y2)
	}
}

func (dc *Context) DrawEllipse(x, y, rx, ry float64) {
	dc.NewSubPath()
	dc.DrawEllipticalArc(x, y, rx, ry, 0, 2*math.Pi)
	dc.ClosePath()
}

func (dc *Context) DrawArc(x, y, r, angle1, angle2 float64) {
	dc.DrawEllipticalArc(x, y, r, r, angle1, angle2)
}

func (dc *Context) DrawCircle(x, y, r float64) {
	dc.NewSubPath()
	dc.DrawEllipticalArc(x, y, r, r, 0, 2*math.Pi)
	dc.ClosePath()
}

func (dc *Context) DrawRegularPolygon(n int, x, y, r, rotation float64) {
	angle := 2 * math.Pi / float64(n)
	rotation -= math.Pi / 2
	if n%2 == 0 {
		rotation += angle / 2
	}
	dc.NewSubPath()
	for i := 0; i < n; i++ {
		a := rotation + angle*float64(i)
		dc.LineTo(x+r*math.Cos(a), y+r*math.Sin(a))
	}
	dc.ClosePath()
}

// DrawImage draws the specified image at the specified point.
func (dc *Context) DrawImage(im image.Image, x, y float64) {
	dc.DrawImageAnchored(im, x, y, 0, 0)
}

// DrawImageAnchored draws the specified image at the specified anchor point.
// The anchor point is x - w * ax, y - h * ay, where w, h is the size of the
// image. Use ax=0.5, ay=0.5 to center the image at the specified point.
func (dc *Context) DrawImageAnchored(im image.Image, x, y, ax, ay float64) {
	s := im.Bounds().Size()
	x -= ax * float64(s.X)
	y -= ay * float64(s.Y)
	transformer := draw.BiLinear
	// fx, fy := float64(x), float64(y)
	m := dc.matrix.Translate(geom.Point{X: x, Y: y})
	s2d := f64.Aff3{m.M11, m.M12, m.M13, m.M21, m.M22, m.M23}
	if dc.mask == nil {
		transformer.Transform(dc.im, s2d, im, im.Bounds(), draw.Over, nil)
	} else {
		transformer.Transform(dc.im, s2d, im, im.Bounds(), draw.Over, &draw.Options{
			DstMask:  dc.mask,
			DstMaskP: image.Point{},
		})
	}
}

// Text Functions

func (dc *Context) SetFontFace(fontFace font.Face) {
	dc.fontFace = fontFace
	dc.fontHeight = (float64(fontFace.Metrics().Height) / 64) * 72 / 96
}

func (dc *Context) LoadFontFace(path string, points float64) error {
	face, err := LoadFontFace(path, points)
	if err == nil {
		dc.fontFace = face
		dc.fontHeight = points * 72 / 96
	}
	return err
}

func (dc *Context) FontHeight() float64 {
	return dc.fontHeight
}

func (dc *Context) drawString(im *image.RGBA, s string, x, y float64) {
	d := &font.Drawer{
		Dst:  im,
		Src:  image.NewUniform(dc.strokePattern.ColorAt(0, 0)),
		Face: dc.fontFace,
		Dot:  fixp(x, y),
	}

	prevC := rune(-1)
	for _, c := range s {
		if prevC >= 0 {
			d.Dot.X += d.Face.Kern(prevC, c)
		}
		dr, mask, maskp, advance, ok := d.Face.Glyph(d.Dot, c)
		if !ok {
			// TODO: is falling back on the U+FFFD glyph the responsibility of
			// the Drawer or the Face?
			// TODO: set prevC = '\ufffd'?
			continue
		}
		sr := dr.Sub(dr.Min)
		transformer := draw.ApproxBiLinear
		fx, fy := float64(dr.Min.X), float64(dr.Min.Y)
		ml := dc.matrix.Translate(geom.Point{X: fx, Y: fy})
		s2d := f64.Aff3{ml.M11, ml.M12, ml.M13, ml.M21, ml.M22, ml.M23}
		transformer.Transform(d.Dst, s2d, d.Src, sr, draw.Over, &draw.Options{
			SrcMask:  mask,
			SrcMaskP: maskp,
		})
		d.Dot.X += advance
		prevC = c
	}
}

// DrawString draws the specified text at the specified point.
func (dc *Context) DrawString(s string, x, y float64) {
	dc.DrawStringAnchored(s, x, y, 0, 0)
}

// DrawStringAnchored draws the specified text at the specified anchor point.
// The anchor point is x - w * ax, y - h * ay, where w, h is the size of the
// text. Use ax=0.5, ay=0.5 to center the text at the specified point.
func (dc *Context) DrawStringAnchored(s string, x, y, ax, ay float64) {
	w, h := dc.MeasureString(s)
	x -= ax * w
	y += ay * h

	if dc.mask == nil {
		dc.drawString(dc.im, s, x, y)
	} else {
		im := image.NewRGBA(image.Rect(0, 0, dc.width, dc.height))
		dc.drawString(im, s, x, y)
		draw.DrawMask(dc.im, dc.im.Bounds(), im, image.Point{}, dc.mask, image.Point{}, draw.Over)
	}
}

// DrawStringWrapped word-wraps the specified string to the given max width
// and then draws it at the specified anchor point using the given line
// spacing and text alignment.
func (dc *Context) DrawStringWrapped(s string, x, y, ax, ay, width, lineSpacing float64, align Align) {
	lines := dc.WordWrap(s, width)

	// sync h formula with MeasureMultilineString
	h := float64(len(lines)) * dc.fontHeight * lineSpacing
	h -= (lineSpacing - 1) * dc.fontHeight

	x -= ax * width
	y -= ay * h
	switch align {
	case AlignLeft:
		ax = 0
	case AlignCenter:
		ax = 0.5
		x += width / 2
	case AlignRight:
		ax = 1
		x += width
	}
	ay = 1
	for _, line := range lines {
		dc.DrawStringAnchored(line, x, y, ax, ay)
		y += dc.fontHeight * lineSpacing
	}
}

func (dc *Context) MeasureMultilineString(s string, lineSpacing float64) (width, height float64) {
	lines := strings.Split(s, "\n")

	height = float64(len(lines)) * dc.fontHeight * lineSpacing
	height -= (lineSpacing - 1) * dc.fontHeight

	d := &font.Drawer{
		Face: dc.fontFace,
	}

	for _, line := range lines {
		adv := d.MeasureString(line)
		currentWidth := float64(adv >> 6)
		if currentWidth > width {
			width = currentWidth
		}
	}

	return width, height
}

// MeasureString returns the rendered width and height of the specified text
// given the current font face.
func (dc *Context) MeasureString(s string) (w, h float64) {
	d := &font.Drawer{
		Face: dc.fontFace,
	}
	a := d.MeasureString(s)
	return float64(a >> 6), dc.fontHeight
}

// WordWrap wraps the specified string to the given max width and current
// font face.
func (dc *Context) WordWrap(s string, w float64) []string {
	return wordWrap(dc, s, w)
}

// Koordinatentransformationen

// Stellt die Transformationsmatrix auf die Einheitsmatrix.
func (dc *Context) Identity() {
	dc.matrix = geom.Identity()
}

// Ergänzt die aktuelle Transformation um eine Translation um die angegebenen
// Werte.
func (dc *Context) Translate(tx, ty float64) {
	dc.matrix = dc.matrix.Translate(geom.Point{X: tx, Y: ty})
}

// Ergänzt die atuelle Transformation um eine Skalierung, wobei die Skalierung
// in X-, resp. Y-Richtung getrennt angegeben werden kann.
func (dc *Context) Scale(sx, sy float64) {
	dc.matrix = dc.matrix.Scale(sx, sy)
}

// Ergänzt die atuelle Transformation um eine Skalierung, wobei die Skalierung
// in X-, resp. Y-Richtung getrennt angegeben werden kann. Der Mittelpunkt
// der Streckung befindet sich beim Punkt 'x,y'.
func (dc *Context) ScaleAbout(sx, sy, x, y float64) {
	dc.matrix = dc.matrix.ScaleAbout(geom.Point{X: x, Y: y}, sx, sy)
}

// Ergänzt die aktuelle Transformation um eine Rotation im Gegenuhrzeigersinn.
// Der Winkel ist im Bogenmass anzugeben.
func (dc *Context) Rotate(angle float64) {
	dc.matrix = dc.matrix.Rotate(angle)
}

// Ergänzt die aktuelle Transformation um eine Rotation im Gegenuhrzeigersinn
// um einen bestimmten Rotationspunkt. Der Winkel ist im Bogenmass anzugeben.
func (dc *Context) RotateAbout(angle, x, y float64) {
	dc.matrix = dc.matrix.RotateAbout(geom.Point{X: x, Y: y}, angle)
}

// Hängt die Transformation im 'm' der aktuellen Transformationsmatrix an.
// Die neue Transformation wird also durch Rechtsmultiplikation mit 'm'
// gebildet.
func (dc *Context) Multiply(m geom.Matrix) {
	dc.matrix = dc.matrix.Multiply(m)
}

// Transformiert den Punkt '(x,y)' mit der aktuellen Transformatiomsmatrix und
// Liefert das Resultat als Zahlenpaar. Translationen werden durchgeführt.
// Vergleiche dazu auch die Methode [TransformVector].
func (dc *Context) TransformPoint(x, y float64) (tx, ty float64) {
	return dc.matrix.TransformPoint(x, y)
}

// Transformiert den Punkt '(x,y)' mit der aktuellen Transformatiomsmatrix und
// Liefert das Resultat als Zahlenpaar. Translationen werden _nicht_
// durchgeführt. Vergleiche dazu auch die Methode [TransformPoint].
func (dc *Context) TransformVector(x, y float64) (tx, ty float64) {
	return dc.matrix.TransformVector(x, y)
}

// Überschreibt die aktuelle Transformationsmatrix der Zeichenumgebung mit
// dem Parameter 'm'.
func (dc *Context) SetMatrix(m geom.Matrix) {
	dc.matrix = m
}

// Retourniert die aktuelle Transformationsmatrix.
func (dc *Context) Matrix() geom.Matrix {
	return dc.matrix
}

// Stack

// Sichert die aktuellen Einstellungen des graphischen Kontexts auf einem
// Stack. Siehe auch [Pop].
func (dc *Context) Push() {
	x := *dc
	dc.stack = append(dc.stack, &x)
}

// Holt die zuletzt gesicherten Einstellungen vom Stack. Siehe auch [Push].
func (dc *Context) Pop() {
	before := *dc
	s := dc.stack
	x, s := s[len(s)-1], s[:len(s)-1]
	*dc = *x
	dc.path = before.path
	dc.start = before.start
	dc.current = before.current
	dc.hasCurrent = before.hasCurrent
}
