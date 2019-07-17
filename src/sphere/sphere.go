package sphere

import (
	"errors"
	"math"

	"github.com/calbim/ray-tracer/src/intersections"
	"github.com/calbim/ray-tracer/src/material"
	"github.com/calbim/ray-tracer/src/matrix"
	"github.com/calbim/ray-tracer/src/ray"
	"github.com/calbim/ray-tracer/src/tuple"
	uuid "github.com/nu7hatch/gouuid"
)

// Sphere represents a unique sphere
type Sphere struct {
	id             string
	Transformation [][]float64
	Material       material.Material
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

// SetTransform sets given transform for sphere
func (s *Sphere) SetTransform(t [][]float64) {
	s.Transformation = t
}

// Intersect returns the points at which a ray intersects a sphere
func (s *Sphere) Intersect(r ray.Ray) ([]intersections.Intersection, error) {
	inverse, err := matrix.Inverse(s.Transformation, 4)
	if err != nil {
		return nil, errors.New("Could not invert sphere's transformation matrix")
	}
	r = ray.Transform(r, inverse)
	sphereToRay := tuple.Subtract(r.Origin, tuple.Point(0.0, 0.0, 0.0))
	a := tuple.DotProduct(r.Direction, r.Direction)
	b := 2 * tuple.DotProduct(r.Direction, sphereToRay)
	c := tuple.DotProduct(sphereToRay, sphereToRay) - 1
	d := b*b - 4*a*c
	if d < 0 {
		return []intersections.Intersection{}, nil
	}
	i1 := intersections.Intersection{Value: (-b - math.Sqrt(d)) / (2 * a), Object: s}
	i2 := intersections.Intersection{Value: (-b + math.Sqrt(d)) / (2 * a), Object: s}

	return intersections.Intersections(i1, i2), nil
}

//NormalAt returns the normal vector at point P on a sphere
func (s *Sphere) NormalAt(p tuple.Tuple) (*tuple.Tuple, error) {
	inverse, err := matrix.Inverse(s.Transformation, 4)
	if err != nil {
		return nil, errors.New("Could not compute object point for world point")
	}
	objectPoint := matrix.MultiplyWithTuple(inverse, p)
	objectNormal := tuple.Subtract(objectPoint, tuple.Point(0, 0, 0))
	worldNormal := matrix.MultiplyWithTuple(matrix.Transpose(inverse), objectNormal)
	worldNormal.W = 0
	normalized := tuple.Normalize(worldNormal)
	return &normalized, nil
}
