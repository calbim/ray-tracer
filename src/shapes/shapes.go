package shapes

import (
	"github.com/calbim/ray-tracer/src/material"
	"github.com/calbim/ray-tracer/src/matrix"
)

// Shape interface
type Shape interface {
	SetTransform([][]float64)
	GetTransform() [][]float64
	GetMaterial() material.Material
	SetMaterial(material.Material)
}

// TestShape is an object of type Shape
type TestShape struct {
	transform [][]float64
	material  material.Material
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

// NewTestShape returns a new TestShape
func NewTestShape() *TestShape {
	return &TestShape{
		transform: matrix.NewIdentity(),
		material:  material.New(),
	}
}
