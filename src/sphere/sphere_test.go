package sphere

import (
	"math"
	"testing"

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
	xs, err := Intersect(s, r)
	if err != nil {
		t.Errorf("Error while calculating intersection")
	}
	if len(xs) != 2 || xs[0].Value != 4 || xs[1].Value != 6 {
		t.Errorf("Ray should intersect sphere at distance %f and %f from the center", 4.0, 6.0)
	}
}

func TestIntersectionTangent(t *testing.T) {
	s, err := New()
	if err != nil {
		t.Errorf("Could not create new sphere")
	}
	r := ray.Ray{Origin: tuple.Point(0, 1, -5), Direction: tuple.Vector(0, 0, 1)}
	xs, err := Intersect(s, r)
	if err != nil {
		t.Errorf("Error while calculating intersection")
	}
	if len(xs) != 2 || xs[0].Value != 5 || xs[1].Value != 5 {
		t.Errorf("Ray should intersect sphere at distance %f and %f from the center", 5.0, 5.0)
	}
}

func TestRayMisses(t *testing.T) {
	s, err := New()
	if err != nil {
		t.Errorf("Could not create new sphere")
	}
	r := ray.Ray{Origin: tuple.Point(0, 2, -5), Direction: tuple.Vector(0, 0, 1)}
	xs, err := Intersect(s, r)
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
	xs, err := Intersect(s, r)
	if err != nil {
		t.Errorf("Error while calculating intersection")
	}
	if len(xs) != 2 || xs[0].Value != -1 || xs[1].Value != 1 {
		t.Errorf("Ray should intersect sphere at 2 points")
	}
}

func TestSphereBehindRay(t *testing.T) {
	s, err := New()
	if err != nil {
		t.Errorf("Could not create new sphere")
	}
	r := ray.Ray{Origin: tuple.Point(0, 0, 5), Direction: tuple.Vector(0, 0, 1)}
	xs, err := Intersect(s, r)
	if err != nil {
		t.Errorf("Error while calculating intersection")
	}
	if len(xs) != 2 || xs[0].Value != -6 || xs[1].Value != -4 {
		t.Errorf("Ray should intersect sphere at 2 points")
	}
}

func TestDefaultTransformation(t *testing.T) {
	s, err := New()
	if err != nil {
		t.Errorf("Could not create new sphere")
	}
	if !matrix.Equals(s.transformation, matrix.NewIdentity(), 4, 4, 4, 4) {
		t.Errorf("Default transformation for sphere should be identity matrix")
	}
}

func TestChangeTransformation(t *testing.T) {
	s, err := New()
	if err != nil {
		t.Errorf("Could not create new sphere")
	}
	transformation := transformations.NewTranslation(2, 3, 4)
	SetTransform(s, transformation)
	if !matrix.Equals(s.transformation, transformation, 4, 4, 4, 4) {
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
	SetTransform(s, transformations.NewScaling(2, 2, 2))
	xs, err := Intersect(s, r)
	if err != nil {
		t.Errorf("Error while calculating intersection")
	}
	if len(xs) != 2 {
		t.Errorf("There should be 2 intersections")
	}
	if xs[0].Value != 3 || xs[1].Value != 7 {
		t.Errorf("Intersection points should be 3 and 7")
	}
}

func TestNormalXAxis(t *testing.T) {
	s, err := New()
	if err != nil {
		t.Errorf("Could not create sphere")
	}
	n, err := NormalAt(s, tuple.Point(1, 0, 0))
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
	n, err := NormalAt(s, tuple.Point(0, 1, 0))
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
	n, err := NormalAt(s, tuple.Point(0, 0, 1))
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
	n, err := NormalAt(s, tuple.Point(math.Sqrt(3)/3, math.Sqrt(3)/3, math.Sqrt(3)/3))
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
	n, err := NormalAt(s, tuple.Point(math.Sqrt(3)/3, math.Sqrt(3)/3, math.Sqrt(3)/3))
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
	SetTransform(s, transformations.NewTranslation(0, 1, 0))
	n, err := NormalAt(s, tuple.Point(0, 1.70711, -0.70711))
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
	SetTransform(s, m)
	n, err := NormalAt(s, tuple.Point(0, math.Sqrt(2)/2, -math.Sqrt(2)/2))
	if err != nil {
		t.Errorf("Error while computing normal")
	}
	if !tuple.Equals(*n, tuple.Vector(0, 0.97014, -0.24254)) {
		t.Errorf("Normal should be %v", tuple.Vector(0, 0.97014, -0.24254))
	}
}
