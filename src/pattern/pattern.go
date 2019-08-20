package pattern

import (
	"math"

	"github.com/calbim/ray-tracer/src/color"
	"github.com/calbim/ray-tracer/src/matrix"

	"github.com/calbim/ray-tracer/src/tuple"
)

//Pattern interface
type Pattern interface {
	GetTransform() *matrix.Matrix
	PatternAt(point tuple.Tuple) *color.Color
}

//Stripe pattern
type Stripe struct {
	a         color.Color
	b         color.Color
	Transform *matrix.Matrix
}

//Gradient pattern
type Gradient struct {
	a         color.Color
	b         color.Color
	Transform *matrix.Matrix
}

//Ring pattern
type Ring struct {
	a         color.Color
	b         color.Color
	Transform *matrix.Matrix
}

//Object struct
type Object struct {
	Transform *matrix.Matrix
}

//NewObject returns a new object
func NewObject() Object {
	return Object{
		Transform: matrix.Identity,
	}
}

//SetTransform sets an object's transform
func (o *Object) SetTransform(transform *matrix.Matrix) {
	o.Transform = transform
}

//NewStripe returns a stripe pattern
func NewStripe(a color.Color, b color.Color) *Stripe {
	return &Stripe{a: a, b: b, Transform: matrix.Identity}
}

//GetTransform returns a stripe's transformation matrix
func (p *Stripe) GetTransform() *matrix.Matrix {
	return p.Transform
}

//PatternAt returns the color of a stripe at a point
func (p *Stripe) PatternAt(point tuple.Tuple) *color.Color {
	if int(math.Floor(point.X))%2 == 0 {
		return &p.a
	}
	return &p.b
}

//NewGradient returns a gradient pattern
func NewGradient(a color.Color, b color.Color) *Gradient {
	return &Gradient{a: a, b: b, Transform: matrix.Identity}
}

//GetTransform returns a gradient's transformation matrix
func (p *Gradient) GetTransform() *matrix.Matrix {
	return p.Transform
}

//PatternAt returns the color of a gradient at a point
func (p *Gradient) PatternAt(point tuple.Tuple) *color.Color {
	diff := p.b.Subtract(p.a)
	c := p.a.Add(diff.Multiply(point.X - math.Floor(point.X)))
	return &c
}

//NewRing returns a gradient pattern
func NewRing(a color.Color, b color.Color) *Ring {
	return &Ring{a: a, b: b, Transform: matrix.Identity}
}

//GetTransform returns a gradient's transformation matrix
func (p *Ring) GetTransform() *matrix.Matrix {
	return p.Transform
}

//PatternAt returns the color of a gradient at a point
func (p *Ring) PatternAt(point tuple.Tuple) *color.Color {
	v := int(math.Floor(math.Sqrt(point.X*point.X + point.Z*point.Z)))
	if v%2 == 0 {
		return &p.a
	}
	return &p.b
}

//AtObject returns the color of a stripe at a point on an object
func AtObject(p Pattern, o Object, point tuple.Tuple) *color.Color {
	oInv, _ := o.Transform.Inverse()
	point = oInv.MultiplyTuple(point)
	pInv, _ := p.GetTransform().Inverse()
	point = pInv.MultiplyTuple(point)
	return p.PatternAt(point)
}
