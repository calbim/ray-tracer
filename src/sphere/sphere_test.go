package sphere

import (
	"testing"

	"github.com/calbim/ray-tracer/src/ray"
	"github.com/calbim/ray-tracer/src/tuple"
)

func TestIntersectionTwoPoints(t *testing.T) {
	s := New()
	r := ray.Ray{Origin: tuple.Point(0, 0, -5), Direction: tuple.Vector(0, 0, 1)}
	xs := Intersect(s, r)
	if len(xs) != 2 || xs[0].Value != 4 || xs[1].Value != 6 {
		t.Errorf("Ray should intersect sphere at distance %f and %f from the center", 4.0, 6.0)
	}
}

func TestIntersectionTangent(t *testing.T) {
	s := New()
	r := ray.Ray{Origin: tuple.Point(0, 1, -5), Direction: tuple.Vector(0, 0, 1)}
	xs := Intersect(s, r)
	if len(xs) != 2 || xs[0].Value != 5 || xs[1].Value != 5 {
		t.Errorf("Ray should intersect sphere at distance %f and %f from the center", 5.0, 5.0)
	}
}

func TestRayMisses(t *testing.T) {
	s := New()
	r := ray.Ray{Origin: tuple.Point(0, 2, -5), Direction: tuple.Vector(0, 0, 1)}
	xs := Intersect(s, r)
	if len(xs) != 0 {
		t.Errorf("Ray should miss sphere")
	}
}

func TestRayInsideSphere(t *testing.T) {
	s := New()
	r := ray.Ray{Origin: tuple.Point(0, 0, 0), Direction: tuple.Vector(0, 0, 1)}
	xs := Intersect(s, r)
	if len(xs) != 2 || xs[0].Value != -1 || xs[1].Value != 1 {
		t.Errorf("Ray should intersect sphere at 2 points")
	}
}

func TestSphereBehindRay(t *testing.T) {
	s := New()
	r := ray.Ray{Origin: tuple.Point(0, 0, 5), Direction: tuple.Vector(0, 0, 1)}
	xs := Intersect(s, r)
	if len(xs) != 2 || xs[0].Value != -6 || xs[1].Value != -4 {
		t.Errorf("Ray should intersect sphere at 2 points")
	}
}
