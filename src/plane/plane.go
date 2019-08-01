package plane

import (
	"math"

	"github.com/calbim/ray-tracer/src/matrix"
	"github.com/calbim/ray-tracer/src/util"

	"github.com/calbim/ray-tracer/src/material"
	"github.com/calbim/ray-tracer/src/ray"
	"github.com/calbim/ray-tracer/src/shapes"
	"github.com/calbim/ray-tracer/src/tuple"
)

// Plane object
type Plane struct {
	Transformation [][]float64
	savedRay       ray.Ray
	Material       material.Material
}

// New returns a new plane
func New() *Plane {
	return &Plane{
		Material:       material.New(),
		Transformation: matrix.NewIdentity(),
	}
}

// GetTransform sets given transform for plane
func (p *Plane) GetTransform() [][]float64 {
	return p.Transformation
}

// SetTransform sets given transform for plane
func (p *Plane) SetTransform(t [][]float64) {
	p.Transformation = t
}

// GetMaterial returns the material of the plane
func (p *Plane) GetMaterial() material.Material {
	return p.Material
}

// SetMaterial returns the material of the plane
func (p *Plane) SetMaterial(m material.Material) {
	p.Material = m
}

// GetSavedRay returns the material of the plane
func (p *Plane) GetSavedRay() ray.Ray {
	return p.savedRay
}

// SetSavedRay returns the material of the plane
func (p *Plane) SetSavedRay(r ray.Ray) {
	p.savedRay = r
}

//Intersect returns the points at which a ray intersects a plane
func (p *Plane) Intersect() ([]shapes.Intersection, error) {
	r := p.savedRay
	points := []shapes.Intersection{}
	if math.Abs(r.Direction.Y) < util.Eps {
		return points, nil
	}
	t := -r.Origin.Y / r.Direction.Y
	i := shapes.Intersection{Value: t, Object: p}
	return []shapes.Intersection{i}, nil
}

//Normal returns the normal vector at point P on a plane
func (p *Plane) Normal(point tuple.Tuple) (*tuple.Tuple, error) {
	v := tuple.Vector(0, 1, 0)
	return &v, nil
}
