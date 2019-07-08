package sphere

import (
	"errors"
	"math"

	"github.com/calbim/ray-tracer/src/intersections"
	"github.com/calbim/ray-tracer/src/matrix"
	"github.com/calbim/ray-tracer/src/ray"
	"github.com/calbim/ray-tracer/src/tuple"
	uuid "github.com/nu7hatch/gouuid"
)

// Sphere represents a unique sphere
type Sphere struct {
	id             string
	transformation [][]float64
}

// New returns a new sphere
func New() (*Sphere, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return nil, errors.New("failed to generate a unique identifier for sphere")
	}
	return &Sphere{
		id:             id.String(),
		transformation: matrix.NewIdentity(),
	}, nil
}

// SetTransform sets given transform for sphere
func SetTransform(s *Sphere, t [][]float64) *Sphere {
	s.transformation = t
	return s
}

// Intersect returns the points at which a ray intersects a sphere
func Intersect(s *Sphere, r ray.Ray) ([]intersections.Intersection, error) {
	inverse, err := matrix.Inverse(s.transformation, 4)
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
	i1 := intersections.Intersection{Value: (-b - math.Sqrt(d)) / (2 * a), Object: *s}
	i2 := intersections.Intersection{Value: (-b + math.Sqrt(d)) / (2 * a), Object: *s}

	return intersections.Intersections(i1, i2), nil
}

//NormalAt returns the normal vector at point P on a sphere
func NormalAt(s *Sphere, p tuple.Tuple) (*tuple.Tuple, error) {
	inverse, err := matrix.Inverse(s.transformation, 4)
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
