package pattern

import (
	"math"

	"github.com/calbim/ray-tracer/src/color"
	"github.com/calbim/ray-tracer/src/matrix"

	"github.com/calbim/ray-tracer/src/tuple"
)

//Pattern struct
type Pattern struct {
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

//Stripe returns a stripe pattern alternating between colors a and b
func Stripe(a color.Color, b color.Color) *Pattern {
	return &Pattern{a: a, b: b, Transform: matrix.Identity}
}

//StripeAt returns the color of a stripe at a point
func (p *Pattern) StripeAt(point tuple.Tuple) *color.Color {
	if int(math.Floor(point.X))%2 == 0 {
		return &p.a
	}
	return &p.b
}

//StripeAtObject returns the color of a stripe at a point on an object
func (p *Pattern) StripeAtObject(o Object, point tuple.Tuple) *color.Color {
	oInv, _ := o.Transform.Inverse()
	point = oInv.MultiplyTuple(point)
	pInv, _ := p.Transform.Inverse()
	point = pInv.MultiplyTuple(point)
	return p.StripeAt(point)
}
