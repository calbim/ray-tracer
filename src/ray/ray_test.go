package ray

import (
	"github.com/calbim/ray-tracer/src/transforms"
	"testing"

	"github.com/calbim/ray-tracer/src/tuple"
)

func TestRay(t *testing.T) {
	o := tuple.Point(1, 2, 3)
	d := tuple.Vector(4, 5, 6)
	r := New(o, d)
	if !r.Direction.Equals(d) || !r.Origin.Equals(o) {
		t.Errorf("wanted r=%v, got %v", New(o, d), r)
	}
}

func TestPositionAtT(t *testing.T) {
	r := New(tuple.Point(2, 3, 4), tuple.Vector(1, 0, 0))
	pos := r.Position(0)
	if !pos.Equals(tuple.Point(2, 3, 4)) {
		t.Errorf("wanted position at 0=%v, got %v", tuple.Point(2, 3, 4), pos)
	}
	pos = r.Position(1)
	if !pos.Equals(tuple.Point(3, 3, 4)) {
		t.Errorf("wanted position at 1=%v, got %v", tuple.Point(3, 3, 4), pos)
	}
	pos = r.Position(-1)
	if !pos.Equals(tuple.Point(1, 3, 4)) {
		t.Errorf("wanted position at -1=%v, got %v", tuple.Point(1, 3, 4), pos)
	}
	pos = r.Position(2.5)
	if !pos.Equals(tuple.Point(4.5, 3, 4)) {
		t.Errorf("wanted position at 2.5=%v, got %v", tuple.Point(4.5, 3, 4), pos)
	}
}


func TestTranslateRay(t *testing.T) {
	r := New(tuple.Point(1, 2, 3), tuple.Vector(0, 1, 0))
	m := transforms.Translation(3, 4, 5)
	r2 := r.Transform(m)
	if !r2.Origin.Equals(tuple.Point(4, 6, 8)) {
		t.Errorf("ray origin after transform should be %v, got %v", tuple.Point(4, 6, 8), r2.Origin)
	}
	if !r2.Direction.Equals(tuple.Vector(0, 1, 0)) {
		t.Errorf("ray direction after transform should be %v, got %v", tuple.Vector(0, 1, 0), r2.Direction)
	}
}
func TestScalingRay(t *testing.T) {
	r := New(tuple.Point(1, 2, 3), tuple.Vector(0, 1, 0))
	m := transforms.Scaling(2, 3, 4)
	r2 := r.Transform(m)
	if !r2.Origin.Equals(tuple.Point(2, 6, 12)) {
		t.Errorf("ray origin after transform should be %v, got %v", tuple.Point(2, 6, 12), r2.Origin)
	}
	if !r2.Direction.Equals(tuple.Vector(0, 3, 0)) {
		t.Errorf("ray direction after transform should be %v, got %v", tuple.Vector(0, 3, 0), r2.Direction)
	}
}
