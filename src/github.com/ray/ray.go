package ray

import (
	"../tuple"
)

// Ray starts from an origin and points towards a direction
type Ray struct {
	origin    tuple.Tuple
	direction tuple.Tuple
}

// Position returns the position of a point travelling along the ray after time t
func Position(r Ray, t float64) tuple.Tuple {
	return tuple.Add(tuple.MultiplyByScalar(r.direction, t), r.origin)
}
