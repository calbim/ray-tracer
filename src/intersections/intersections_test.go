package intersections

import (
	"testing"

	"github.com/calbim/ray-tracer/src/sphere"

	"github.com/calbim/ray-tracer/src/ray"
	"github.com/calbim/ray-tracer/src/tuple"
)

func TestIntersectionObject(t *testing.T) {
	s, err := sphere.New()
	if err != nil {
		t.Errorf("Error creating sphere due to %v", err)
	}
	i := Intersection{3.5, s}
	if i.Value != 3.5 {
		t.Errorf("The intersection point should be %f", 3.5)
	}
	if i.Object != s {
		t.Errorf("The intersected object should be %v", s)
	}
}

func TestIntersectionCollection(t *testing.T) {
	s, err := sphere.New()
	if err != nil {
		t.Errorf("Error creating sphere due to %v", err)
	}
	i1 := Intersection{1, s}
	i2 := Intersection{2, s}
	xs := Intersections(i1, i2)
	if len(xs) != 2 {
		t.Errorf("Number of intersections should be 2")
	}
	if xs[0].Value != 1 || xs[1].Value != 2 {
		t.Errorf("Intersection points should be 1 and 2 respectively")
	}
}

func TestHitAllPositive(t *testing.T) {
	s, err := sphere.New()
	if err != nil {
		t.Errorf("Error creating sphere due to %v", err)
	}
	i1 := Intersection{1, s}
	i2 := Intersection{2, s}
	xs := Intersections(i1, i2)
	i := Hit(xs)
	if *i != i1 {
		t.Errorf("hit should be %v", i1)
	}
}

func TestHitSomePositive(t *testing.T) {
	s, err := sphere.New()
	if err != nil {
		t.Errorf("Error creating sphere due to %v", err)
	}
	i1 := Intersection{-1, s}
	i2 := Intersection{1, s}
	xs := Intersections(i1, i2)
	i := Hit(xs)
	if *i != i2 {
		t.Errorf("hit should be %v", i2)
	}
}

func TestHitAllNegative(t *testing.T) {
	s, err := sphere.New()
	if err != nil {
		t.Errorf("Error creating sphere due to %v", err)
	}
	i1 := Intersection{-2, s}
	i2 := Intersection{-1, s}
	xs := Intersections(i1, i2)
	i := Hit(xs)
	if i != nil {
		t.Errorf("hit should be %v", nil)
	}
}

func TestHitMultipleIntersections(t *testing.T) {
	s, err := sphere.New()
	if err != nil {
		t.Errorf("Error creating sphere due to %v", err)
	}
	i1 := Intersection{5, s}
	i2 := Intersection{7, s}
	i3 := Intersection{-3, s}
	i4 := Intersection{2, s}
	xs := Intersections(i1, i2, i3, i4)
	i := Hit(xs)
	if *i != i4 {
		t.Errorf("hit should be %v", i4)
	}
}

func TestPrepareComputations(t *testing.T) {
	r := ray.Ray{
		Origin:    tuple.Point(0, 0, -5),
		Direction: tuple.Vector(0, 0, 1),
	}
	s, err := sphere.New()
	if err != nil {
		t.Errorf("Error %v while creating sphere", err)
	}
	i := Intersection{
		Value:  4,
		Object: s,
	}
	comps, err := PrepareComputations(i, r)
	if err != nil {
		t.Errorf("Could not calculate computation due to error %v", err)
	}
	if comps.Value != i.Value {
		t.Errorf("Computation value should be %v, but it is %v", i.Value, comps.Value)
	}
	if comps.Object != i.Object {
		t.Errorf("Computation value should be %v, but it is %v", i.Object, comps.Object)
	}
	if comps.Point != tuple.Point(0, 0, -1) {
		t.Errorf("Computation point should be %v, but it is %v", tuple.Point(0, 0, -1), comps.Point)
	}
	if comps.Eyev != tuple.Vector(0, 0, -1) {
		t.Errorf("Computation eyev should be %v, but it is %v", tuple.Point(0, 0, -1), comps.Eyev)
	}
	if comps.Normal != tuple.Vector(0, 0, -1) {
		t.Errorf("Computation normal should be %v, but it is %v", tuple.Point(0, 0, -1), comps.Normal)
	}
}
