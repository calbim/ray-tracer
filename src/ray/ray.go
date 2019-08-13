package ray

import (
	"github.com/calbim/ray-tracer/src/matrix"
	"github.com/calbim/ray-tracer/src/tuple"
)

// Ray represents a ray
type Ray struct {
	Origin    tuple.Tuple
	Direction tuple.Tuple
}

//New returns a ray with given origin and direction
func New(o tuple.Tuple, d tuple.Tuple) Ray {
	return Ray{
		Origin:    o,
		Direction: d,
	}
}

// Position returns the position of a point on a ray at distance t
func (r Ray) Position(t float64) tuple.Tuple {
	p := r.Direction.Multiply(t)
	return p.Add(r.Origin)
}

// Transform applies a transform to a ray and returns the transformed ray
func (r Ray) Transform(m *matrix.Matrix) Ray {
	origin := m.MultiplyTuple(r.Origin)
	direction := m.MultiplyTuple(r.Direction)
	return New(origin, direction)
}
