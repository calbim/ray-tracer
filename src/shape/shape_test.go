package shape

import (
	"math"
	"testing"

	"github.com/calbim/ray-tracer/src/material"
	"github.com/calbim/ray-tracer/src/transforms"

	"github.com/calbim/ray-tracer/src/matrix"

	"github.com/calbim/ray-tracer/src/ray"
	"github.com/calbim/ray-tracer/src/tuple"
)

func TestSphereIntersection(t *testing.T) {
	s := NewSphere()
	r := ray.New(tuple.Point(0, 0, -5), tuple.Vector(0, 0, 1))
	xs := s.Intersect(r)
	if len(xs) != 2 || xs[0].Value != 4 || xs[1].Value != 6 {
		t.Errorf("wanted intersection points to be %v and %v, got %v and %v", 4.0, 6.0, xs[0], xs[1])
	}
}

func TestSphereIntersectionTangent(t *testing.T) {
	s := NewSphere()
	r := ray.New(tuple.Point(0, 1, -5), tuple.Vector(0, 0, 1))
	xs := s.Intersect(r)
	if len(xs) != 2 || xs[0].Value != 5 || xs[1].Value != 5 {
		t.Errorf("wanted intersection points to be %v and %v, got %v and %v", 5.0, 5.0, xs[0], xs[1])
	}
}

func TestSphereRayMisses(t *testing.T) {
	s := NewSphere()
	r := ray.New(tuple.Point(0, 2, -5), tuple.Vector(0, 0, 1))
	xs := s.Intersect(r)
	if len(xs) != 0 {
		t.Errorf("wamted 0 intersections, got %v", len(xs))
	}
}

func TestRayInsideSphere(t *testing.T) {
	s := NewSphere()
	r := ray.New(tuple.Point(0, 0, 0), tuple.Vector(0, 0, 1))
	xs := s.Intersect(r)
	if len(xs) != 2 || xs[0].Value != -1 || xs[1].Value != 1 {
		t.Errorf("wanted intersections points to be %v and %v, got %v and %v", -1, 1, xs[0], xs[1])
	}
}
func TestSphereBehindRay(t *testing.T) {
	s := NewSphere()
	r := ray.New(tuple.Point(0, 0, 5), tuple.Vector(0, 0, 1))
	xs := s.Intersect(r)
	if len(xs) != 2 || xs[0].Value != -6 || xs[1].Value != -4 {
		t.Errorf("wanted intersections points to be %v and %v, got %v and %v", -6, -4, xs[0], xs[1])
	}
}

func TestIntersection(t *testing.T) {
	s := NewSphere()
	i := NewIntersection(3.5, s)
	if i.Value != 3.5 {
		t.Errorf("wanted intersection point=%v, got %v", 3.5, i.Value)
	}
	if i.Object != s {
		t.Errorf("wanted intersected object=%v, got %v", s, i.Object)
	}
}
func TestIntersections(t *testing.T) {
	s := NewSphere()
	i1 := Intersection{1, s}
	i2 := Intersection{2, s}
	xs := Intersections(i1, i2)
	if len(xs) != 2 {
		t.Errorf("wanted %v intersections, got %v", 2, len(xs))
	}
	if xs[0].Value != 1 || xs[1].Value != 2 {
		t.Errorf("wanted %v and %v, got %v and %v", 1, 2, xs[0], xs[1])
	}
}

func TestIntersectionSetsObject(t *testing.T) {
	r := ray.New(tuple.Point(0, 0, 5), tuple.Vector(0, 0, 1))
	s := NewSphere()
	xs := s.Intersect(r)
	if len(xs) != 2 {
		t.Errorf("wanted %v intersections, got %v", 2, len(xs))
	}
	if xs[0].Object != s || xs[1].Object != s {
		t.Errorf("wanted object to be %v, got %v", s, xs[0])
	}
}

func TestHitAllPositive(t *testing.T) {
	s := NewSphere()
	i1 := Intersection{1, s}
	i2 := Intersection{2, s}
	xs := Intersections(i1, i2)
	i := Hit(xs)
	if *i != i1 {
		t.Errorf("wanted hit=%v, got %v", i1, i)
	}
}

func TestHitSomePositive(t *testing.T) {
	s := NewSphere()
	i1 := Intersection{-1, s}
	i2 := Intersection{1, s}
	xs := Intersections(i1, i2)
	i := Hit(xs)
	if *i != i2 {
		t.Errorf("hit should be %v, gpt %v", i2, i)
	}
}

func TestHitAllNegative(t *testing.T) {
	s := NewSphere()
	i1 := Intersection{-2, s}
	i2 := Intersection{-1, s}
	xs := Intersections(i1, i2)
	i := Hit(xs)
	if i != nil {
		t.Errorf("hit should be %v, got %v", nil, i)
	}
}

func TestHitMultipleIntersections(t *testing.T) {
	s := NewSphere()
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

func TestSphereTransformation(t *testing.T) {
	s := NewSphere()
	if !s.Transform.Equals(matrix.Identity) {
		t.Errorf("wanted transform=%v, got %v", matrix.Identity, s.Transform)
	}
}

func TestChangeTransformation(t *testing.T) {
	s := NewSphere()
	transform := transforms.Translation(2, 3, 4)
	s.SetTransform(transform)
	if !s.Transform.Equals(transform) {
		t.Errorf("wanted transform=%v, got %v", transform, s.Transform)
	}
}

func TestIntersectScaledSphere(t *testing.T) {
	r := ray.New(tuple.Point(0, 0, -5), tuple.Vector(0, 0, 1))
	s := NewSphere()
	s.SetTransform(transforms.Scaling(2, 2, 2))
	xs := s.Intersect(r)
	if len(xs) != 2 {
		t.Errorf("wanted %v intersections, got %v", 2, len(xs))
	}
	if xs[0].Value != 3 || xs[1].Value != 7 {
		t.Errorf("wanted intersection points to be %v and %v, got %v and %v", 3, 7, xs[0].Value, xs[1].Value)
	}
}

func TestIntersectTranslatedSphere(t *testing.T) {
	r := ray.New(tuple.Point(0, 0, -5), tuple.Vector(0, 0, 1))
	s := NewSphere()
	s.SetTransform(transforms.Translation(5, 0, 0))
	xs := s.Intersect(r)
	if len(xs) != 0 {
		t.Errorf("wanted %v intersections, got %v", 0, len(xs))
	}
}

func TestNormalXAxis(t *testing.T) {
	s := NewSphere()
	n := s.Normal(tuple.Point(1, 0, 0))
	if n == nil {
		t.Errorf("normal is nil")
	}
	if !n.Equals(tuple.Vector(1, 0, 0)) {
		t.Errorf("wanted n=%v, got %v", tuple.Vector(1, 0, 0), n)
	}
}

func TestNormalYAxis(t *testing.T) {
	s := NewSphere()
	n := s.Normal(tuple.Point(0, 1, 0))
	if n == nil {
		t.Errorf("normal is nil")
	}
	if !n.Equals(tuple.Vector(0, 1, 0)) {
		t.Errorf("wanted n=%v, got %v", tuple.Vector(0, 1, 0), n)
	}
}

func TestNormalZAxis(t *testing.T) {
	s := NewSphere()
	n := s.Normal(tuple.Point(0, 0, 1))
	if n == nil {
		t.Errorf("normal is nil")
	}
	if !n.Equals(tuple.Vector(0, 0, 1)) {
		t.Errorf("wanted n=%v, got %v", tuple.Vector(0, 0, 1), n)
	}
}

func TestNormalNonAxial(t *testing.T) {
	s := NewSphere()
	n := s.Normal(tuple.Point(math.Sqrt(3)/3, math.Sqrt(3)/3, math.Sqrt(3)/3))
	if n == nil {
		t.Errorf("normal is nil")
	}
	v := tuple.Vector(math.Sqrt(3)/3, math.Sqrt(3)/3, math.Sqrt(3)/3)
	if !n.Equals(v) {
		t.Errorf("wanted n=%v, got %v", v, n)
	}
}
func TestNormalIsNormalized(t *testing.T) {
	s := NewSphere()
	p := tuple.Point(math.Sqrt(3)/3, math.Sqrt(3)/3, math.Sqrt(3)/3)
	n := s.Normal(p)
	if !n.Equals(n.Normalize()) {
		t.Errorf("wanted n=%v, got %v", n.Normalize(), n)
	}
}

func TestNormalTranslatedSphere(t *testing.T) {
	s := NewSphere()
	s.SetTransform(transforms.Translation(0, 1, 0))
	n := s.Normal(tuple.Point(0, 1.70711, -0.70711))
	if !n.Equals(tuple.Vector(0, 0.70711, -0.70711)) {
		t.Errorf("wanted normal=%v, got %v", tuple.Vector(0, 0.70711, -0.70711), n)
	}
}

func TestNormalTransformedSphere(t *testing.T) {
	s := NewSphere()
	s.SetTransform(transforms.Chain(transforms.RotationZ(math.Pi/5), transforms.Scaling(1, 0.5, 1)))
	n := s.Normal(tuple.Point(0, math.Sqrt(2)/2, -math.Sqrt(2)/2))
	expected := tuple.Vector(0, 0.97014, -0.24254)
	if !n.Equals(expected) {
		t.Errorf("wanted normal=%v, got %v", expected, n)
	}
}

func TestSphereHasDefaultMaterial(t *testing.T) {
	s := NewSphere()
	m := s.Material
	if m != material.New() {
		t.Errorf("wanted material=%v, got %v", material.New(), m)
	}
}

func TestSphereCanBeAssignedMaterial(t *testing.T) {
	s := NewSphere()
	m := material.New()
	m.Ambient = 1
	s.Material = m
	if s.Material != m {
		t.Errorf("wanted material to be %v, got %v", m, s.Material)
	}
}

func TestPrepareComputations(t *testing.T) {
	r := ray.New(tuple.Point(0, 0, -5), tuple.Vector(0, 0, 1))
	s := NewSphere()
	i := &Intersection{
		Value:  4,
		Object: s,
	}
	comps := i.PrepareComputations(r)
	if comps.Value != i.Value {
		t.Errorf("wanted value=%v, got %v", i.Value, comps.Value)
	}
	if comps.Object != i.Object {
		t.Errorf("wanted object value=%v, got %v", i.Object, comps.Object)
	}
	if comps.Point != tuple.Point(0, 0, -1) {
		t.Errorf("wanted point=%v, got %v", tuple.Point(0, 0, -1), comps.Point)
	}
	if comps.Eyev != tuple.Vector(0, 0, -1) {
		t.Errorf("wanted eyev=%v, got %v", tuple.Vector(0, 0, -1), comps.Eyev)
	}
	if comps.Normal != tuple.Vector(0, 0, -1) {
		t.Errorf("wanted normal=%v, got %v", tuple.Vector(0, 0, -1), comps.Normal)
	}
}

func TestInsideFlagWhenIntersectionOutside(t *testing.T) {
	r := ray.New(tuple.Point(0, 0, -5), tuple.Vector(0, 0, 1))
	s := NewSphere()
	i := &Intersection{Value: 4.0, Object: s}
	comp := i.PrepareComputations(r)
	if comp.Inside == true {
		t.Errorf("wanted hit to be outside the object")
	}
}

func TestHitWhenIntersectionInside(t *testing.T) {
	r := ray.New(tuple.Point(0, 0, 0), tuple.Vector(0, 0, 1))
	s := NewSphere()
	i := &Intersection{Value: 1.0, Object: s}
	comp := i.PrepareComputations(r)
	if comp.Inside != true {
		t.Errorf("wanted hit to be inside the object")
	}
	if comp.Point != tuple.Point(0, 0, 1) {
		t.Errorf("wanted point=%v, got %v", tuple.Point(0, 0, 1), comp.Point)
	}
	if comp.Eyev != tuple.Vector(0, 0, -1) {
		t.Errorf("wanted point=%v, got %v", tuple.Point(0, 0, -1), comp.Eyev)
	}
	if comp.Normal != tuple.Vector(0, 0, -1) {
		t.Errorf("wanted normal=%v, got %v", tuple.Point(0, 0, -1), comp.Point)
	}
}
