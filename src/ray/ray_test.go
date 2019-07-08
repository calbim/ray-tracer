package ray

import (
	"testing"

	"github.com/calbim/ray-tracer/src/transformations"
	"github.com/calbim/ray-tracer/src/tuple"
)

func TestRay(t *testing.T) {
	o := tuple.Point(1, 2, 3)
	d := tuple.Vector(4, 5, 6)
	r := Ray{o, d}
	if !tuple.Equals(r.Origin, o) || !tuple.Equals(r.Direction, d) {
		t.Errorf("ray was not initiliazed correctly")
	}
}

func TestPositionAfterT(t *testing.T) {
	r := Ray{tuple.Point(2, 3, 4), tuple.Vector(1, 0, 0)}
	if !tuple.Equals(Position(r, 0), tuple.Point(2, 3, 4)) {
		t.Errorf("after t 0, point should be at position %v, is at %v", Position(r, 0), tuple.Point(2, 3, 4))
	}
	if !tuple.Equals(Position(r, 1), tuple.Point(3, 3, 4)) {
		t.Errorf("after t 1, point should be at position %v, is at %v", Position(r, 1), tuple.Point(3, 3, 4))
	}
	if !tuple.Equals(Position(r, -1), tuple.Point(1, 3, 4)) {
		t.Errorf("before t 1, point should be at position %v, is at %v", Position(r, -1), tuple.Point(1, 3, 4))
	}
	if !tuple.Equals(Position(r, 2.5), tuple.Point(4.5, 3, 4)) {
		t.Errorf("after t 2.5, point should be at position %v, is at %v", Position(r, 2.5), tuple.Point(4.5, 3, 4))
	}
}

func TestTranslateRay(t *testing.T) {
	r := Ray{tuple.Point(1, 2, 3), tuple.Vector(0, 1, 0)}
	m := transformations.NewTranslation(3, 4, 5)
	r2 := Transform(r, m)
	if !tuple.Equals(r2.Origin, tuple.Point(4, 6, 8)) {
		t.Errorf("ray origin after transform should be %v", tuple.Point(4, 6, 8))
	}
	if !tuple.Equals(r2.Direction, tuple.Vector(0, 1, 0)) {
		t.Errorf("ray direction after transform should be %v", tuple.Vector(0, 1, 0))
	}
}

func TestScalingRay(t *testing.T) {
	r := Ray{tuple.Point(1, 2, 3), tuple.Vector(0, 1, 0)}
	m := transformations.NewScaling(2, 3, 4)
	r2 := Transform(r, m)
	if !tuple.Equals(r2.Origin, tuple.Point(2, 6, 12)) {
		t.Errorf("ray origin after transform should be %v", tuple.Point(2, 6, 12))
	}
	if !tuple.Equals(r2.Direction, tuple.Vector(0, 3, 0)) {
		t.Errorf("ray direction after transform should be %v", tuple.Vector(0, 3, 0))
	}
}
