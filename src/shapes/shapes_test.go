package shapes

import (
	"math"
	"testing"

	"github.com/calbim/ray-tracer/src/ray"
	"github.com/calbim/ray-tracer/src/tuple"

	"github.com/calbim/ray-tracer/src/material"

	"github.com/calbim/ray-tracer/src/transformations"

	"github.com/calbim/ray-tracer/src/matrix"
)

func TestDefaultTransformation(t *testing.T) {
	s := NewTestShape()
	if !matrix.Equals(s.GetTransform(), matrix.NewIdentity(), 4, 4, 4, 4) {
		t.Errorf("Default transform should be identity matrix")
	}
}

func TestAssigningTransformation(t *testing.T) {
	s := NewTestShape()
	s.SetTransform(transformations.NewTranslation(2, 3, 4))
	if !matrix.Equals(s.GetTransform(), transformations.NewTranslation(2, 3, 4), 4, 4, 4, 4) {
		t.Errorf("Shape transform should be %v but it is %v", transformations.NewTranslation(2, 3, 4), s.GetTransform())
	}
}

func TestShapeHasDefaultMaterial(t *testing.T) {
	s := NewTestShape()
	m := s.GetMaterial()
	if m != material.New() {
		t.Errorf("A shape should have a default material")
	}
}

func TestAssignMaterialToShape(t *testing.T) {
	s := NewTestShape()
	m := material.New()
	m.Ambient = 1
	s.SetMaterial(m)
	if m != s.GetMaterial() {
		t.Errorf("Shape material should be %v, but is %v", m, s.GetMaterial())
	}
}

func TestIntersectScaledShape(t *testing.T) {
	r := ray.Ray{
		Origin:    tuple.Point(0, 0, -5),
		Direction: tuple.Vector(0, 0, 1),
	}
	s := NewTestShape()
	s.SetTransform(transformations.NewScaling(2, 2, 2))
	_, err := Intersect(s, r)
	if err != nil {
		t.Errorf("Error while calculating intersection")
	}
	sr := s.GetSavedRay()
	if !tuple.Equals(sr.Origin, tuple.Point(0, 0, -2.5)) {
		t.Errorf("Ray origin should be %v but it is %v", tuple.Point(0, 0, -2.5), sr.Origin)
	}
}

func TestTranslatedShape(t *testing.T) {
	r := ray.Ray{
		Origin:    tuple.Point(0, 0, -5),
		Direction: tuple.Vector(0, 0, 1),
	}
	s := NewTestShape()
	s.SetTransform(transformations.NewTranslation(5, 0, 0))
	_, err := Intersect(s, r)
	if err != nil {
		t.Errorf("Error while calculating intersection")
	}
	sr := s.GetSavedRay()
	if !tuple.Equals(sr.Origin, tuple.Point(-5, 0, -5)) {
		t.Errorf("Ray origin should be %v but it is %v", tuple.Point(-5, 0, -5), sr.Origin)
	}
}

func TestNormalOnTranslatedSphere(t *testing.T) {
	s := NewTestShape()
	s.SetTransform(transformations.NewTranslation(0, 1, 0))
	n, err := Normal(s, tuple.Point(0, 1.70711, -0.70711))
	if err != nil {
		t.Errorf("Could not find normal due to error %v", err)
	}
	if !tuple.Equals(*n, tuple.Vector(0, 0.70711, -0.70711)) {
		t.Errorf("Normal should be %v", tuple.Vector(0, 0.70711, -1.70711))
	}
}

func TestNormalOnTransformedSphere(t *testing.T) {
	s := NewTestShape()
	m := transformations.RotationZ(math.Pi / 5)
	s.SetTransform(matrix.Multiply(transformations.NewScaling(1, 0.5, 1), m))
	n, err := Normal(s, tuple.Point(0, math.Sqrt2/2, -math.Sqrt2/2))
	if err != nil {
		t.Errorf("Could not find normal due to error %v", err)
	}
	if !tuple.Equals(*n, tuple.Vector(0, 0.97014, -0.24254)) {
		t.Errorf("Normal should be %v", tuple.Vector(0, 0.97014, -0.24254))
	}
}
