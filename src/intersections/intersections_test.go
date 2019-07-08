package intersections

import (
	"testing"
)

func TestIntersectionObject(t *testing.T) {
	var sphere interface{}
	i := Intersection{3.5, sphere}
	if i.Value != 3.5 {
		t.Errorf("The intersection point should be %f", 3.5)
	}
	if i.Object != sphere {
		t.Errorf("The intersected object should be %v", sphere)
	}
}

func TestIntersectionCollection(t *testing.T) {
	var sphere interface{}
	i1 := Intersection{1, sphere}
	i2 := Intersection{2, sphere}
	xs := Intersections(i1, i2)
	if len(xs) != 2 {
		t.Errorf("Number of intersections should be 2")
	}
	if xs[0].Value != 1 || xs[1].Value != 2 {
		t.Errorf("Intersection points should be 1 and 2 respectively")
	}
}

func TestHitAllPositive(t *testing.T) {
	var sphere interface{}
	i1 := Intersection{1, sphere}
	i2 := Intersection{2, sphere}
	xs := Intersections(i1, i2)
	i := Hit(xs)
	if *i != i1 {
		t.Errorf("hit should be %v", i1)
	}
}

func TestHitSomePositive(t *testing.T) {
	var sphere interface{}
	i1 := Intersection{-1, sphere}
	i2 := Intersection{1, sphere}
	xs := Intersections(i1, i2)
	i := Hit(xs)
	if *i != i2 {
		t.Errorf("hit should be %v", i2)
	}
}

func TestHitAllNegative(t *testing.T) {
	var sphere interface{}
	i1 := Intersection{-2, sphere}
	i2 := Intersection{-1, sphere}
	xs := Intersections(i1, i2)
	i := Hit(xs)
	if i != nil {
		t.Errorf("hit should be %v", nil)
	}
}

func TestHitMultipleIntersections(t *testing.T) {
	var sphere interface{}
	i1 := Intersection{5,sphere}
	i2 := Intersection{7,sphere}
	i3 := Intersection{-3,sphere}
	i4 := Intersection{2,sphere}
	xs := Intersections(i1,i2,i3,i4)
	i := Hit(xs)
	if *i != i4 {
		t.Errorf("hit should be %v", i4)
	}
}
