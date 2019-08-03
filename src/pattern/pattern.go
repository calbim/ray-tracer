package pattern

import (
	"math"

	"github.com/calbim/ray-tracer/src/tuple"
)

// Pattern object
type Pattern struct {
	A tuple.Tuple
	B tuple.Tuple
}

// StripePattern returns a new stripe pattern with 2 colors
func StripePattern(a tuple.Tuple, b tuple.Tuple) Pattern {
	return Pattern{
		A: a,
		B: b,
	}
}

//StripeAt returns the colour of a pattern at given point
func StripeAt(p Pattern, point tuple.Tuple) tuple.Tuple {
	if int(math.Floor(point.X))%2 == 0 {
		return p.A
	}
	return p.B
}
