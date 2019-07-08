package ray

import (
	"github.com/calbim/ray-tracer/src/matrix"
	"github.com/calbim/ray-tracer/src/tuple"
)

// Ray starts from an origin and points towards a direction
type Ray struct {
	Origin    tuple.Tuple
	Direction tuple.Tuple
}

// Position returns the position of a point travelling along the ray after time t
func Position(r Ray, t float64) tuple.Tuple {
	return tuple.Add(tuple.MultiplyByScalar(r.Direction, t), r.Origin)
}

// Transform takes a ray and a transformation matrix m
// and returns a transformed ray
func Transform(r Ray, m [][]float64) Ray {
	origin := matrix.MultiplyWithTuple(m, r.Origin)
	direction := matrix.MultiplyWithTuple(m, r.Direction)
	return Ray{origin, direction}
}
