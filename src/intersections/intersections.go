package intersections

import "math"

//Intersection encapsulates an object and the intersection point
type Intersection struct {
	Value  float64
	Object interface{}
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
