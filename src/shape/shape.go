package shape

import (
	"github.com/calbim/ray-tracer/src/material"
	"github.com/calbim/ray-tracer/src/matrix"
	"github.com/calbim/ray-tracer/src/ray"
	"github.com/calbim/ray-tracer/src/tuple"
)

//Shape interface
type Shape interface {
	LocalIntersect(ray.Ray) []Intersection
	LocalNormalAt(tuple.Tuple) *tuple.Tuple
	GetMaterial() *material.Material
	SetMaterial(*material.Material)
	SetTransform(*matrix.Matrix)
	GetTransform() *matrix.Matrix
}

//NormalAt returns the normal of a shape at a point
func NormalAt(s Shape, p tuple.Tuple) *tuple.Tuple {
	transform := s.GetTransform()
	inv, _ := transform.Inverse()
	localPoint := inv.MultiplyTuple(p)
	localNormal := s.LocalNormalAt(localPoint)
	transpose := inv.Transpose()
	worldNormal := transpose.MultiplyTuple(*localNormal)
	worldNormal.W = 0
	n := worldNormal.Normalize()
	return &n
}

// Intersect intersects a shape with a ray
func Intersect(s Shape, r ray.Ray) []Intersection {
	transform := s.GetTransform()
	inv, _ := transform.Inverse()
	localRay := r.Transform(inv)
	return s.LocalIntersect(localRay)
}
