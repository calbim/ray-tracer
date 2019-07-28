package intersections

import (
	"fmt"
	"math"

	"github.com/calbim/ray-tracer/src/util"

	"github.com/calbim/ray-tracer/src/material"
	"github.com/calbim/ray-tracer/src/ray"

	"github.com/calbim/ray-tracer/src/tuple"
)

//Object interface
type Object interface {
	GetMaterial() material.Material
	SetMaterial(material.Material)
	Intersect(ray.Ray) ([]float64, error)
	NormalAt(tuple.Tuple) (*tuple.Tuple, error)
	SetTransform([][]float64)
}

// Intersection encapsulates an object and the intersection point
type Intersection struct {
	Value  float64
	Object Object
}

// Computation object that contains more data about an intersection
type Computation struct {
	Value     float64 // t value
	Object    Object
	Point     tuple.Tuple
	Overpoint tuple.Tuple
	Eyev      tuple.Tuple
	Normal    tuple.Tuple
	Inside    bool
}

//Intersections returns a collection of intersection objects
func Intersections(intersections ...Intersection) []Intersection {
	arr := make([]Intersection, len(intersections))
	for i, val := range intersections {
		arr[i] = val
	}
	return arr
}

//Hit takes a collection of intersection points and returns the hit point
// which is the smallest nonnegative value
func Hit(intersections []Intersection) *Intersection {
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

// NormalAt returns the normal for the underlying object at point p
func NormalAt(o Object, p tuple.Tuple) (*tuple.Tuple, error) {
	return o.NormalAt(p)
}

// Material returns the object's material
func Material(o Object) material.Material {
	return o.GetMaterial()
}

//SetMaterial sets material m in object o
func SetMaterial(o Object, m material.Material) {
	o.SetMaterial(m)
}

// Intersect returns the intersections of object o with ray r
func Intersect(o Object, r ray.Ray) ([]Intersection, error) {
	points, err := o.Intersect(r)
	if err != nil {
		return nil, fmt.Errorf("Could not find intersection points due to error %v", err)
	}
	intersections := []Intersection{}
	for _, val := range points {
		intersections = append(intersections, Intersection{Value: val, Object: o})
	}
	return intersections, nil
}

// PrepareComputations calculates the Computation object for an intersection
func PrepareComputations(i Intersection, r ray.Ray) (*Computation, error) {
	tValue := i.Value
	object := i.Object
	point := ray.Position(r, tValue)
	normal, err := NormalAt(object, point)
	eyev := tuple.Negate(r.Direction)
	inside := false
	if tuple.DotProduct(*normal, eyev) < 0 {
		inside = true
		tmp := tuple.Negate(*normal)
		normal = &tmp
	}
	if err != nil {
		return nil, fmt.Errorf("Could not calculate computation because of error %v", err)
	}

	comps := Computation{
		Value:     tValue,
		Object:    object,
		Point:     point,
		Overpoint: tuple.Add(point, tuple.MultiplyByScalar(*normal, util.Eps)),
		Eyev:      eyev,
		Normal:    *normal,
		Inside:    inside,
	}
	return &comps, nil
}
