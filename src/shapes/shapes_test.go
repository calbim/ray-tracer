package shapes

import (
	"testing"

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

func TestAssignMaterialToSphere(t *testing.T) {
	s := NewTestShape()
	m := material.New()
	m.Ambient = 1
	s.SetMaterial(m)
	if m != s.GetMaterial() {
		t.Errorf("Shape material should be %v, but is %v", m, s.GetMaterial())
	}
}
