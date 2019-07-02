package sphere

import (
	"testing"

	"../ray"
	"../tuple"
)

func TestIntersection(t *testing.T) {
	s := New()
	r := ray.Ray{tuple.Point(0, 0, -5), tuple.Vector(0, 0, 1)}
	xs := Intersect(s, r)
	if len(xs) != 2 || xs[0] != 4 || xs[1] != 6 {
		t.Errorf("Ray should intersect sphere at distance %f and %f from the center", 4, 6)
	}
}
