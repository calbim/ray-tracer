package sphere

import (
	"errors"

	"github.com/calbim/ray-tracer/src/ray"
	"github.com/calbim/ray-tracer/src/tuple"
	uuid "github.com/nu7hatch/gouuid"
)

// Sphere represents a unique sphere
type Sphere struct {
	id string
}

// New returns a new sphere
func New() Sphere {
	id, err := uuid.NewV4()
	if err != nil {
		errors.New("failed to generate a unique identifier for sphere")
	}
	return Sphere{
		id: id.String(),
	}
}

// Intersect returns the points at which a ray intersects a sphere
// func Intersect(s Sphere, r ray.Ray) []float64 {
// 	ray := ray.Ray{Origin: tuple.Point(0, 0, -5), Direction: tuple.Vector(0, 0, 1)}
// 	sphere := New()
// }
