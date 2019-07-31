package shapes

import (
	"errors"
	"fmt"

	"github.com/calbim/ray-tracer/src/tuple"

	"github.com/calbim/ray-tracer/src/material"
	"github.com/calbim/ray-tracer/src/matrix"
	"github.com/calbim/ray-tracer/src/ray"
)

// Shape interface
type Shape interface {
	SetTransform([][]float64)
	GetTransform() [][]float64
	GetMaterial() material.Material
	SetMaterial(material.Material)
	GetSavedRay() ray.Ray
	SetSavedRay(r ray.Ray)
	Intersect() ([]Intersection, error)
	Normal(tuple.Tuple) (*tuple.Tuple, error)
}

// TestShape is an object of type Shape
type TestShape struct {
	transform [][]float64
	material  material.Material
	savedRay  ray.Ray
}

// GetTransform returns transform for shape
func (ts *TestShape) GetTransform() [][]float64 {
	return ts.transform
}

// SetTransform sets t transform for shape
func (ts *TestShape) SetTransform(t [][]float64) {
	ts.transform = t
}

// GetMaterial returns shape material
func (ts *TestShape) GetMaterial() material.Material {
	return ts.material
}

// SetMaterial sets material for shape
func (ts *TestShape) SetMaterial(m material.Material) {
	ts.material = m
}

// GetSavedRay returns transform for shape
func (ts *TestShape) GetSavedRay() ray.Ray {
	return ts.savedRay
}

// SetSavedRay sets t transform for shape
func (ts *TestShape) SetSavedRay(r ray.Ray) {
	ts.savedRay = r
}

// Intersect returns the intersections when ray intersects shape
func (ts *TestShape) Intersect() ([]Intersection, error) {
	return []Intersection{}, nil
}

// Normal calculates a normal at point p
func (ts *TestShape) Normal(p tuple.Tuple) (*tuple.Tuple, error) {
	v := tuple.Vector(p.X, p.Y, p.Z)
	return &v, nil
}

// NewTestShape returns a new TestShape
func NewTestShape() *TestShape {
	return &TestShape{
		transform: matrix.NewIdentity(),
		material:  material.New(),
	}
}

// Intersect returns the intersections when a ray intersects a shape
func Intersect(s Shape, r ray.Ray) ([]Intersection, error) {
	inv, err := matrix.Inverse(s.GetTransform(), 4)
	if err != nil {
		return nil, fmt.Errorf("Could not calculate matrix inverse %v", err)
	}
	s.SetSavedRay(ray.Ray{
		Origin:    matrix.MultiplyWithTuple(inv, r.Origin),
		Direction: matrix.MultiplyWithTuple(inv, r.Direction),
	})
	return s.Intersect()
}

// Normal calculates the normal at point p for shape s
func Normal(s Shape, p tuple.Tuple) (*tuple.Tuple, error) {
	inverse, err := matrix.Inverse(s.GetTransform(), 4)
	if err != nil {
		return nil, errors.New("Could not compute object point for world point")
	}
	localPoint := matrix.MultiplyWithTuple(inverse, p)
	localNormal, err := s.Normal(localPoint)
	if err != nil {
		return nil, fmt.Errorf("Could not calculate normal because %v", err)
	}
	worldNormal := matrix.MultiplyWithTuple(matrix.Transpose(inverse), *localNormal)
	worldNormal.W = 0
	normalized := tuple.Normalize(worldNormal)
	return &normalized, nil
}
