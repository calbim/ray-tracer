package intersections

import (
	"testing"
	"github.com/calbim/ray-tracer/src/ray"
	"github.com/calbim/ray-tracer/src/tuple"
)
type dummy struct {

}

func (d dummy) NormalAt(t tuple.Tuple)(*tuple.Tuple, error) {
	return &tuple.Tuple{}, nil
}

func (d dummy) Intersect(r ray.Ray)([]Intersection, error){
	return nil, nil
}

func  (d dummy)SetTransform([][]float64){
}
func TestIntersectionObject(t *testing.T) {
	d:=dummy{}
	i := Intersection{3.5, d}
	if i.Value != 3.5 {
		t.Errorf("The intersection point should be %f", 3.5)
	}
	if i.Object != d {
		t.Errorf("The intersected object should be %v", d)
	}
}

func TestIntersectionCollection(t *testing.T) {
	d:=dummy{}
	i1 := Intersection{1, d}
	i2 := Intersection{2, d}
	xs := Intersections(i1, i2)
	if len(xs) != 2 {
		t.Errorf("Number of intersections should be 2")
	}
	if xs[0].Value != 1 || xs[1].Value != 2 {
		t.Errorf("Intersection points should be 1 and 2 respectively")
	}
}

func TestHitAllPositive(t *testing.T) {
	d:=dummy{}
	i1 := Intersection{1, d}
	i2 := Intersection{2, d}
	xs := Intersections(i1, i2)
	i := Hit(xs)
	if *i != i1 {
		t.Errorf("hit should be %v", i1)
	}
}

func TestHitSomePositive(t *testing.T) {
	d:=dummy{}
	i1 := Intersection{-1, d}
	i2 := Intersection{1, d}
	xs := Intersections(i1, i2)
	i := Hit(xs)
	if *i != i2 {
		t.Errorf("hit should be %v", i2)
	}
}

func TestHitAllNegative(t *testing.T) {
	d:=dummy{}
	i1 := Intersection{-2, d}
	i2 := Intersection{-1, d}
	xs := Intersections(i1, i2)
	i := Hit(xs)
	if i != nil {
		t.Errorf("hit should be %v", nil)
	}
}

func TestHitMultipleIntersections(t *testing.T) {
	d:=dummy{}
	i1 := Intersection{5,d}
	i2 := Intersection{7,d}
	i3 := Intersection{-3,d}
	i4 := Intersection{2,d}
	xs := Intersections(i1,i2,i3,i4)
	i := Hit(xs)
	if *i != i4 {
		t.Errorf("hit should be %v", i4)
	}
}
