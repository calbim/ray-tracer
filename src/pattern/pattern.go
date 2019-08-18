package pattern

import (
	"math"

	"github.com/calbim/ray-tracer/src/color"

	"github.com/calbim/ray-tracer/src/tuple"
)

//Pattern struct
type Pattern struct {
	a color.Color
	b color.Color
}

//Stripe returns a stripe pattern alternating between colors a and b
func Stripe(a color.Color, b color.Color) *Pattern {
	return &Pattern{a: a, b: b}
}

//StripeAt returns the color of a stripe at a point
func (p *Pattern) StripeAt(point tuple.Tuple) *color.Color {
	if int(math.Floor(point.X))%2 == 0 {
		return &p.a
	}
	return &p.b
}
