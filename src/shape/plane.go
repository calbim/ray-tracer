package shape

import (
	"math"

	"github.com/calbim/ray-tracer/src/material"
	"github.com/calbim/ray-tracer/src/matrix"
	"github.com/calbim/ray-tracer/src/ray"
	"github.com/calbim/ray-tracer/src/tuple"
	"github.com/calbim/ray-tracer/src/util"
)

//Plane is a flat surface that extends infinitely in xz
type Plane struct {
	Transform *matrix.Matrix
	Material  *material.Material
}

//NewPlane returns a sphere with a unique ID
func NewPlane() *Plane {
	m := material.New()
	return &Plane{
		Transform: matrix.Identity,
		Material:  &m,
	}
}

//LocalIntersect returns the points at which a ray intersects a plane
func (p *Plane) LocalIntersect(r ray.Ray) []Intersection {
	if math.Abs(r.Direction.Y) < util.Eps {
		return []Intersection{}
	}
	t := -r.Origin.Y / r.Direction.Y
	i := NewIntersection(t, p)
	return []Intersection{i}
}

// SetTransform sets given transform for sphere
func (p *Plane) SetTransform(m *matrix.Matrix) {
	p.Transform = m
}

//LocalNormalAt returns the normal vector at point P on a plane
func (p *Plane) LocalNormalAt(point tuple.Tuple) *tuple.Tuple {
	n := tuple.Vector(0, 1, 0)
	return &n
}

//GetMaterial returns the material of the plane
func (p *Plane) GetMaterial() *material.Material {
	return p.Material
}

//SetMaterial returns the material of the plane
func (p *Plane) SetMaterial(m *material.Material) {
	p.Material = m
}

//GetTransform returns the transform of the plane
func (p *Plane) GetTransform() *matrix.Matrix {
	return p.Transform
}
