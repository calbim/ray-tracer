package sphere

import (
	"math"
	"testing"

	"github.com/calbim/ray-tracer/src/material"
	"github.com/calbim/ray-tracer/src/matrix"
	"github.com/calbim/ray-tracer/src/ray"
	"github.com/calbim/ray-tracer/src/transformations"
	"github.com/calbim/ray-tracer/src/tuple"
)

func TestIntersectionTwoPoints(t *testing.T) {
	s, err := New()
	if err != nil {
		t.Errorf("Could not create new sphere")
	}
	r := ray.Ray{Origin: tuple.Point(0, 0, -5), Direction: tuple.Vector(0, 0, 1)}
	xs, err := s.Intersect(r)
	if err != nil {
		t.Errorf("Error while calculating intersection")
	}
	if len(xs) != 2 || xs[0] != 4 || xs[1] != 6 {
		t.Errorf("Ray should intersect sphere at distance %f and %f from the center", 4.0, 6.0)
	}
}

func TestIntersectionTangent(t *testing.T) {
	s, err := New()
	if err != nil {
		t.Errorf("Could not create new sphere")
	}
	r := ray.Ray{Origin: tuple.Point(0, 1, -5), Direction: tuple.Vector(0, 0, 1)}
	xs, err := s.Intersect(r)
	if err != nil {
		t.Errorf("Error while calculating intersection")
	}
	if len(xs) != 2 || xs[0] != 5 || xs[1] != 5 {
		t.Errorf("Ray should intersect sphere at distance %f and %f from the center", 5.0, 5.0)
	}
}

func TestRayMisses(t *testing.T) {
	s, err := New()
	if err != nil {
		t.Errorf("Could not create new sphere")
	}
	r := ray.Ray{Origin: tuple.Point(0, 2, -5), Direction: tuple.Vector(0, 0, 1)}
	xs, err := s.Intersect(r)
	if err != nil {
		t.Errorf("Error while calculating intersection")
	}
	if len(xs) != 0 {
		t.Errorf("Ray should miss sphere")
	}
}

func TestRayInsideSphere(t *testing.T) {
	s, err := New()
	if err != nil {
		t.Errorf("Could not create new sphere")
	}
	r := ray.Ray{Origin: tuple.Point(0, 0, 0), Direction: tuple.Vector(0, 0, 1)}
	xs, err := s.Intersect(r)
	if err != nil {
		t.Errorf("Error while calculating intersection")
	}
	if len(xs) != 2 || xs[0] != -1 || xs[1] != 1 {
		t.Errorf("Ray should intersect sphere at 2 points")
	}
}

func TestSphereBehindRay(t *testing.T) {
	s, err := New()
	if err != nil {
		t.Errorf("Could not create new sphere")
	}
	r := ray.Ray{Origin: tuple.Point(0, 0, 5), Direction: tuple.Vector(0, 0, 1)}
	xs, err := s.Intersect(r)
	if err != nil {
		t.Errorf("Error while calculating intersection")
	}
	if len(xs) != 2 || xs[0] != -6 || xs[1] != -4 {
		t.Errorf("Ray should intersect sphere at 2 points")
	}
}

func TestDefaultTransformation(t *testing.T) {
	s, err := New()
	if err != nil {
		t.Errorf("Could not create new sphere")
	}
	if !matrix.Equals(s.Transformation, matrix.NewIdentity(), 4, 4, 4, 4) {
		t.Errorf("Default transformation for sphere should be identity matrix")
	}
}

func TestChangeTransformation(t *testing.T) {
	s, err := New()
	if err != nil {
		t.Errorf("Could not create new sphere")
	}
	transformation := transformations.NewTranslation(2, 3, 4)
	s.SetTransform(transformation)
	if !matrix.Equals(s.Transformation, transformation, 4, 4, 4, 4) {
		t.Errorf("Transformation for sphere should be %v", transformation)
	}
}

func TestIntersectScaledSphere(t *testing.T) {
	r := ray.Ray{
		Origin:    tuple.Point(0, 0, -5),
		Direction: tuple.Vector(0, 0, 1),
	}
	s, err := New()
	if err != nil {
		t.Errorf("Could not create new sphere")
	}
	s.SetTransform(transformations.NewScaling(2, 2, 2))
	xs, err := s.Intersect(r)
	if err != nil {
		t.Errorf("Error while calculating intersection")
	}
	if len(xs) != 2 {
		t.Errorf("There should be 2 intersections")
	}
	if xs[0] != 3 || xs[1] != 7 {
		t.Errorf("Intersection points should be 3 and 7")
	}
}

func TestNormalXAxis(t *testing.T) {
	s, err := New()
	if err != nil {
		t.Errorf("Could not create sphere")
	}
	n, err := s.NormalAt(tuple.Point(1, 0, 0))
	if err != nil {
		t.Errorf("Error while computing normal")
	}
	if !tuple.Equals(*n, tuple.Vector(1, 0, 0)) {
		t.Errorf("Normal vector should be %v", tuple.Vector(1, 0, 0))
	}
}

func TestNormalYAxis(t *testing.T) {
	s, err := New()
	if err != nil {
		t.Errorf("Could not create sphere")
	}
	n, err := s.NormalAt(tuple.Point(0, 1, 0))
	if err != nil {
		t.Errorf("Error while computing normal")
	}
	if !tuple.Equals(*n, tuple.Vector(0, 1, 0)) {
		t.Errorf("Normal vector should be %v", tuple.Vector(0, 1, 0))
	}
}

func TestNormalZAxis(t *testing.T) {
	s, err := New()
	if err != nil {
		t.Errorf("Could not create sphere")
	}
	n, err := s.NormalAt(tuple.Point(0, 0, 1))
	if err != nil {
		t.Errorf("Error while computing normal")
	}
	if !tuple.Equals(*n, tuple.Vector(0, 0, 1)) {
		t.Errorf("Normal vector should be %v", tuple.Vector(0, 0, 1))
	}
}

func TestNormalNonAxial(t *testing.T) {
	s, err := New()
	if err != nil {
		t.Errorf("Could not create sphere")
	}
	n, err := s.NormalAt(tuple.Point(math.Sqrt(3)/3, math.Sqrt(3)/3, math.Sqrt(3)/3))
	if err != nil {
		t.Errorf("Error while computing normal")
	}
	if !tuple.Equals(*n, tuple.Vector(math.Sqrt(3)/3, math.Sqrt(3)/3, math.Sqrt(3)/3)) {
		t.Errorf("Normal vector should be %v", tuple.Vector(math.Sqrt(3)/3, math.Sqrt(3)/3, math.Sqrt(3)/3))
	}
}
func TestNormalIsNormalized(t *testing.T) {
	s, err := New()
	if err != nil {
		t.Errorf("Could not create sphere")
	}
	n, err := s.NormalAt(tuple.Point(math.Sqrt(3)/3, math.Sqrt(3)/3, math.Sqrt(3)/3))
	if err != nil {
		t.Errorf("Error while computing normal")
	}
	if !tuple.Equals(*n, tuple.Normalize(*n)) {
		t.Errorf("Normal should be normalized")
	}
}

func TestNormalForTranslatedSphere(t *testing.T) {
	s, err := New()
	if err != nil {
		t.Errorf("Could not create sphere")
	}
	s.SetTransform(transformations.NewTranslation(0, 1, 0))
	n, err := s.NormalAt(tuple.Point(0, 1.70711, -0.70711))
	if err != nil {
		t.Errorf("Error while computing normal")
	}
	if !tuple.Equals(*n, tuple.Vector(0, 0.70711, -0.70711)) {
		t.Errorf("Normal at point of translated sphere should be %v", tuple.Vector(0, 0.70711, -0.70711))
	}
}

func TestNormalForTransformedSphere(t *testing.T) {
	s, err := New()
	if err != nil {
		t.Errorf("Could not create sphere")
	}
	m := matrix.Multiply(transformations.NewScaling(1, 0.5, 1), transformations.RotationZ(math.Pi/5))
	s.SetTransform(m)
	n, err := s.NormalAt(tuple.Point(0, math.Sqrt(2)/2, -math.Sqrt(2)/2))
	if err != nil {
		t.Errorf("Error while computing normal")
	}
	if !tuple.Equals(*n, tuple.Vector(0, 0.97014, -0.24254)) {
		t.Errorf("Normal should be %v", tuple.Vector(0, 0.97014, -0.24254))
	}
}

func TestSphereHasDefaultMaterial(t *testing.T) {
	s, err := New()
	if err != nil {
		t.Errorf("Error while creating sphere %v", err)
	}
	m := s.Material
	if m != material.New() {
		t.Errorf("A sphere should have a default material")
	}
}

func TestAssignMaterialToSphere(t *testing.T) {
	s, err := New()
	if err != nil {
		t.Errorf("Error while creating sphere %v", err)
	}
	m := material.New()
	m.Ambient = 1
	s.Material = m
	if m != s.Material {
		t.Errorf("Sphere material should be %v, but is %v", m, s.Material)
	}
}
