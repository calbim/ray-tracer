package sphere

import (
	"errors"
	"math"

	"github.com/calbim/ray-tracer/src/material"
	"github.com/calbim/ray-tracer/src/matrix"
	"github.com/calbim/ray-tracer/src/ray"
	"github.com/calbim/ray-tracer/src/shapes"
	"github.com/calbim/ray-tracer/src/tuple"
	uuid "github.com/nu7hatch/gouuid"
)

// Sphere represents a unique sphere
type Sphere struct {
	id             string
	Transformation [][]float64
	Material       material.Material
	savedRay       ray.Ray
}

// New returns a new sphere centered at the origin and with radius 1 unit
func New() (*Sphere, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return nil, errors.New("failed to generate a unique identifier for sphere")
	}
	return &Sphere{
		id:             id.String(),
		Transformation: matrix.NewIdentity(),
		Material:       material.New(),
	}, nil
}

// GetTransform sets given transform for sphere
func (s *Sphere) GetTransform() [][]float64 {
	return s.Transformation
}

// SetTransform sets given transform for sphere
func (s *Sphere) SetTransform(t [][]float64) {
	s.Transformation = t
}

// GetMaterial returns the material of the sphere
func (s *Sphere) GetMaterial() material.Material {
	return s.Material
}

// SetMaterial returns the material of the sphere
func (s *Sphere) SetMaterial(m material.Material) {
	s.Material = m
}

// GetSavedRay returns the material of the sphere
func (s *Sphere) GetSavedRay() ray.Ray {
	return s.savedRay
}

// SetSavedRay returns the material of the sphere
func (s *Sphere) SetSavedRay(r ray.Ray) {
	s.savedRay = r
}

// Intersect returns the points at which a ray intersects a sphere
func (s *Sphere) Intersect() ([]shapes.Intersection, error) {
	r := s.savedRay
	sphereToRay := tuple.Subtract(r.Origin, tuple.Point(0.0, 0.0, 0.0))
	a := tuple.DotProduct(r.Direction, r.Direction)
	b := 2 * tuple.DotProduct(r.Direction, sphereToRay)
	c := tuple.DotProduct(sphereToRay, sphereToRay) - 1
	points := []shapes.Intersection{}
	d := b*b - 4*a*c
	if d < 0 {
		return points, nil
	}
	points = append(points, shapes.Intersection{Value: (-b - math.Sqrt(d)) / (2 * a), Object: s})
	points = append(points, shapes.Intersection{Value: (-b + math.Sqrt(d)) / (2 * a), Object: s})

	return points, nil
}

//Normal returns the normal vector at point P on a sphere
func (s *Sphere) Normal(p tuple.Tuple) (*tuple.Tuple, error) {
	n := tuple.Subtract(p, tuple.Point(0, 0, 0))
	return &n, nil
}
