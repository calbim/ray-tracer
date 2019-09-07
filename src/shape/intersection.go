package shape

import (
	"math"

	"github.com/calbim/ray-tracer/src/ray"
	"github.com/calbim/ray-tracer/src/tuple"
	"github.com/calbim/ray-tracer/src/util"
)

//Intersection represents a point where an object is intersected
type Intersection struct {
	Value  float64
	Object Shape
}

//NewIntersection returns a new intersection object
func NewIntersection(t float64, o Shape) Intersection {
	return Intersection{
		Value:  t,
		Object: o,
	}
}

//Intersections returns a an array of intersections
func Intersections(intersections ...Intersection) []Intersection {
	arr := make([]Intersection, len(intersections))
	for i, val := range intersections {
		arr[i] = val
	}
	return arr
}

//Hit returns the lowest nonnegative point of intersection
func Hit(intersections []Intersection) *Intersection {
	if intersections == nil {
		return nil
	}
	min := math.MaxFloat64
	index := 0
	hit := false
	for i, val := range intersections {
		if val.Value >= 0 && val.Value < min {
			min = val.Value
			index = i
			hit = true
		}
	}
	if hit == false {
		return nil
	}
	return &intersections[index]
}

// Computation object that contains data about an intersection
type Computation struct {
	Value     float64
	Object    Shape
	Point     tuple.Tuple
	Overpoint tuple.Tuple
	Eyev      tuple.Tuple
	Normal    tuple.Tuple
	Reflectv  tuple.Tuple
	Inside    bool
}

// PrepareComputations calculates the Computation object for an intersection
func (i *Intersection) PrepareComputations(r ray.Ray) Computation {
	tValue := i.Value
	object := i.Object
	point := r.Position(tValue)
	normal := NormalAt(object, point)
	eyev := r.Direction.Negate()
	inside := false
	if normal.DotProduct(eyev) < 0 {
		inside = true
		tmp := normal.Negate()
		normal = &tmp
	}

	return Computation{
		Value:     tValue,
		Object:    object,
		Point:     point,
		Overpoint: (point.Add(normal.Multiply(util.Eps))),
		Eyev:      eyev,
		Normal:    *normal,
		Inside:    inside,
		Reflectv:  r.Direction.Reflect(*normal),
	}
}
