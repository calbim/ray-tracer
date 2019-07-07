package intersections

import "github.com/calbim/ray-tracer/src/sphere"

//Intersection encapsulates an object and the intersection point
type Intersection struct{
	Value float64
	Object sphere.Sphere
}