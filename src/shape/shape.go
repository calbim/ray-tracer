package shape

import (
	"github.com/calbim/ray-tracer/src/material"
	"github.com/calbim/ray-tracer/src/matrix"
	"github.com/calbim/ray-tracer/src/ray"
	"github.com/calbim/ray-tracer/src/tuple"
)

//Shape interface
type Shape interface {
	Intersect(r ray.Ray) []Intersection
	Normal(p tuple.Tuple) *tuple.Tuple
	GetMaterial() *material.Material
	SetMaterial(material.Material)
	GetTransform() *matrix.Matrix
}
