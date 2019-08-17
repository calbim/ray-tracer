package shape

import (
	"math"
	"time"

	"github.com/calbim/ray-tracer/src/material"
	"github.com/calbim/ray-tracer/src/matrix"

	"github.com/calbim/ray-tracer/src/tuple"

	"github.com/calbim/ray-tracer/src/ray"
)

//Sphere struct
type Sphere struct {
	ID        int64
	Transform *matrix.Matrix
	Material  *material.Material
}

//NewSphere returns a sphere with a unique ID
func NewSphere() *Sphere {
	m := material.New()
	return &Sphere{
		ID:        time.Now().Unix(),
		Transform: matrix.Identity,
		Material:  &m,
	}
}

//LocalIntersect returns the points at which a ray intersects a sphere
func (s *Sphere) LocalIntersect(r ray.Ray) []Intersection {
	sphereToRay := r.Origin.Subtract(tuple.Point(0, 0, 0))
	a := r.Direction.DotProduct(r.Direction)
	b := 2 * r.Direction.DotProduct(sphereToRay)
	c := sphereToRay.DotProduct(sphereToRay) - 1
	d := b*b - 4*a*c
	if d < 0 {
		return nil
	}
	i1 := NewIntersection(((-b - math.Sqrt(d)) / (2 * a)), s)
	i2 := NewIntersection(((-b + math.Sqrt(d)) / (2 * a)), s)
	return Intersections(i1, i2)
}

// SetTransform sets given transform for sphere
func (s *Sphere) SetTransform(m *matrix.Matrix) {
	s.Transform = m
}

//LocalNormalAt returns the normal vector at point P on a sphere
func (s *Sphere) LocalNormalAt(p tuple.Tuple) *tuple.Tuple {
	n := p.Subtract(tuple.Point(0, 0, 0))
	return &n
}

//GetMaterial returns the material of the sphere
func (s *Sphere) GetMaterial() *material.Material {
	return s.Material
}

//SetMaterial returns the material of the sphere
func (s *Sphere) SetMaterial(m *material.Material) {
	s.Material = m
}

//GetTransform returns the transform of the sphere
func (s *Sphere) GetTransform() *matrix.Matrix {
	return s.Transform
}
