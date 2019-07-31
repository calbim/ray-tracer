package world

import (
	"errors"
	"testing"

	"github.com/calbim/ray-tracer/src/light"
	"github.com/calbim/ray-tracer/src/material"
	"github.com/calbim/ray-tracer/src/matrix"
	"github.com/calbim/ray-tracer/src/ray"
	"github.com/calbim/ray-tracer/src/shapes"
	"github.com/calbim/ray-tracer/src/sphere"
	"github.com/calbim/ray-tracer/src/transformations"
	"github.com/calbim/ray-tracer/src/tuple"
	"github.com/calbim/ray-tracer/src/util"
)

func TestCreateWorld(t *testing.T) {
	w := World{}
	if w.Objects != nil || w.Light != nil {
		t.Errorf("An empty world object should not have any objects in it or a light source.")
	}
}
func TestDefaultWorld(t *testing.T) {
	w, err := NewDefault()
	if err != nil {
		t.Errorf("Could not create new world %v", err)
	}
	light := light.PointLight{Position: tuple.Point(-10, 10, -10), Intensity: tuple.Color(1, 1, 1)}

	m := material.Material{Color: tuple.Color(0.8, 1.0, 0.6), Diffuse: 0.7, Specular: 0.2, Ambient: 0.1, Shininess: 200}
	s1, err := sphere.New()
	if err != nil {
		t.Errorf("Could not create sphere s1")
	}
	s1.Material = m

	s2, err := sphere.New()
	if err != nil {
		t.Errorf("Could not create sphere s2")
	}
	s2.Material = material.New()
	s2.SetTransform(transformations.NewScaling(0.5, 0.5, 0.5))

	if *w.Light != light {
		t.Errorf("World light source should be %v, is %v", light, w.Light)
	}
	containsS1, err := contains(w.Objects, s1)
	containsS2, err := contains(w.Objects, s2)
	if !containsS1 || !containsS2 {
		t.Errorf("World should contain objects s1 %v and s2 %v but contains %v and %v", s1, s2, w.Objects[0], w.Objects[1])
	}

}

func TestWorldIntersect(t *testing.T) {
	w, err := NewDefault()
	if err != nil {
		t.Errorf("Error while creating world %v", err)
	}
	r := ray.Ray{Origin: tuple.Point(0.0, 0.0, -5.0), Direction: tuple.Vector(0.0, 0.0, 1.0)}
	xs, err := w.Intersect(r)
	if err != nil {
		t.Errorf("Error while intersection world with ray %v", err)
	}
	if len(xs) != 4 {
		t.Errorf("Expect 4 points of intersections, got %v", len(xs))
	}
	if xs[0].Value != 4 || xs[1].Value != 4.5 || xs[2].Value != 5.5 || xs[3].Value != 6 {
		t.Errorf("Expected interesection points to be 4,4.5,5,6, got %v", xs)
	}
}

func TestShadingIntersection(t *testing.T) {
	w, err := NewDefault()
	if err != nil {
		t.Errorf("Error %v creating world", err)
	}
	r := ray.Ray{Origin: tuple.Point(0, 0, -5), Direction: tuple.Vector(0, 0, 1)}
	shape := w.Objects[0]
	i := shapes.Intersection{Value: 4.0, Object: shape}
	comps, err := shapes.PrepareComputations(i, r)
	if err != nil {
		t.Error("Error preparing computations", err)
	}
	c := ShadeHit(*w, *comps)
	if !tuple.Equals(c, tuple.Color(0.38066, 0.47583, 0.2855)) {
		t.Errorf("Shade of Hit should be %v, but it is %v", tuple.Color(0.38066, 0.47583, 0.2855), c)
	}
}
func TestShadingIntersectionInside(t *testing.T) {
	w, err := NewDefault()
	if err != nil {
		t.Errorf("Error %v creating world", err)
	}
	w.Light = &light.PointLight{Intensity: tuple.Color(1, 1, 1), Position: tuple.Point(0, 0.25, 0)}
	r := ray.Ray{Origin: tuple.Point(0, 0, 0), Direction: tuple.Vector(0, 0, 1)}
	shape := w.Objects[1]
	i := shapes.Intersection{Value: 0.5, Object: shape}
	comps, err := shapes.PrepareComputations(i, r)
	if err != nil {
		t.Error("Error preparing computations", err)
	}
	c := ShadeHit(*w, *comps)
	if !tuple.Equals(c, tuple.Color(0.90498, 0.90498, 0.90498)) {
		t.Errorf("Shade of Hit should be %v, but it is %v", tuple.Color(0.90498, 0.90498, 0.90498), c)
	}
}

func TestShadeHitWhenIntersectionIsInShadow(t *testing.T) {
	w, err := NewDefault()
	if err != nil {
		t.Errorf("Error %v creating world", err)
	}
	w.Light = &light.PointLight{
		Intensity: tuple.Color(1, 1, 1),
		Position:  tuple.Point(0, 0, -10),
	}
	s1, err := sphere.New()
	if err != nil {
		t.Errorf("Error %v while creating sphere", err)
	}
	w.Objects = append(w.Objects, s1)
	s2, err := sphere.New()
	if err != nil {
		t.Errorf("Error %v while creating sphere", err)
	}
	s2.SetTransform(transformations.NewTranslation(0, 0, 10))
	w.Objects = append(w.Objects, s2)
	r := ray.Ray{Origin: tuple.Point(0, 0, 5), Direction: tuple.Vector(0, 0, 1)}
	i := shapes.Intersection{Value: 4, Object: s2}
	comps, err := shapes.PrepareComputations(i, r)
	if err != nil {
		t.Errorf("Error %v while preparing computation", err)
	}
	c := ShadeHit(*w, *comps)
	if !tuple.Equals(c, tuple.Color(0.1, 0.1, 0.1)) {
		t.Errorf("Color of hit should be %v but it is %v", tuple.Color(0.1, 0.1, 0.1), c)
	}
}

func TestColorWhenRayMisses(t *testing.T) {
	w, err := NewDefault()
	if err != nil {
		t.Errorf("Error %v creating world", err)
	}
	r := ray.Ray{Origin: tuple.Point(0, 0, -5), Direction: tuple.Vector(0, 1, 0)}
	c, err := ColorAt(*w, r)
	if *c != tuple.Color(0, 0, 0) {
		t.Errorf("When a ray fails to intersect an object, the colour returned should be black")
	}
}

func TestDefaultColorWhenRayHits(t *testing.T) {
	w, err := NewDefault()
	if err != nil {
		t.Errorf("Error %v creating world", err)
	}
	r := ray.Ray{Origin: tuple.Point(0, 0, -5), Direction: tuple.Vector(0, 0, 1)}
	c, err := ColorAt(*w, r)
	if !tuple.Equals(*c, tuple.Color(0.38066, 0.47583, 0.2855)) {
		t.Errorf("Default colour for when a ray  to intersects an object should be %v", tuple.Color(0.38066, 0.47583, 0.2855))
	}
}

func TestColorWhenIntersectionIsBehindRay(t *testing.T) {
	w, err := NewDefault()
	if err != nil {
		t.Errorf("Error %v creating world", err)
	}
	m := w.Objects[0].GetMaterial()
	m.Ambient = 1
	outer := w.Objects[0]
	outer.SetMaterial(m)
	inner := w.Objects[1]
	mInner := inner.GetMaterial()
	mInner.Ambient = 1
	inner.SetMaterial(mInner)

	r := ray.Ray{Origin: tuple.Point(0, 0, 0.075), Direction: tuple.Vector(0, 0, -1)}
	c, err := ColorAt(*w, r)
	if !tuple.Equals(*c, inner.GetMaterial().Color) {
		t.Errorf("Intersection color should be %v but is %v", inner.GetMaterial().Color, c)
	}
}

func TestNoShadowWhenNothingIsCollinearWithPointAndLight(t *testing.T) {
	w, err := NewDefault()
	if err != nil {
		t.Errorf("Error %v creating world", err)
	}
	p := tuple.Point(0, 10, 0)
	shadow, err := w.IsShadowed(p)
	if err != nil {
		t.Errorf("Error while determining if point is shadowed")
	}
	if shadow {
		t.Errorf("There should be no shadow")
	}
}

func TestShadowPointIsBetweenObjectAndLight(t *testing.T) {
	w, err := NewDefault()
	if err != nil {
		t.Errorf("Error %v creating world", err)
	}
	p := tuple.Point(10, -10, 10)
	shadow, err := w.IsShadowed(p)
	if err != nil {
		t.Errorf("Error while determining if point is shadowed")
	}
	if !shadow {
		t.Errorf("There should be a shadow")
	}
}

func TestNoShadowWhenObjectIsBehindLight(t *testing.T) {
	w, err := NewDefault()
	if err != nil {
		t.Errorf("Error %v creating world", err)
	}
	p := tuple.Point(-20, 20, -20)
	shadow, err := w.IsShadowed(p)
	if err != nil {
		t.Errorf("Error while determining if point is shadowed")
	}
	if shadow {
		t.Errorf("There should be no shadow")
	}
}

func TestNoShadowWhenObjectIsBehindPoint(t *testing.T) {
	w, err := NewDefault()
	if err != nil {
		t.Errorf("Error %v creating world", err)
	}
	p := tuple.Point(-2, 2, -2)
	shadow, err := w.IsShadowed(p)
	if err != nil {
		t.Errorf("Error while determining if point is shadowed")
	}
	if shadow {
		t.Errorf("There should be no shadow")
	}
}

func TestHitShouldOffsetThePoint(t *testing.T) {
	r := ray.Ray{
		Origin:    tuple.Point(0, 0, -5),
		Direction: tuple.Vector(0, 0, 1),
	}
	s, err := sphere.New()
	if err != nil {
		t.Errorf("Error while creating sphere %v", err)
	}
	inv, err := matrix.Inverse(transformations.NewTranslation(0, 0, 1), 4)
	if err != nil {
		t.Errorf("Could not calculate inverse because %v", err)
	}
	r.Origin = matrix.MultiplyWithTuple(inv, r.Origin)
	r.Direction = matrix.MultiplyWithTuple(inv, r.Direction)
	s.SetSavedRay(r)
	i := shapes.Intersection{Value: 5, Object: s}
	comps, err := shapes.PrepareComputations(i, r)
	if err != nil {
		t.Errorf("Error while preparing computation due to %v", err)
	}
	if comps.Overpoint.Z >= -util.Eps/2 || comps.Point.Z <= comps.Overpoint.Z {
		t.Errorf("Overpoint should be a little less 0 and actual intersection point")
	}
}

func contains(list []shapes.Shape, s *sphere.Sphere) (bool, error) {
	for _, obj := range list {
		sphereObject, ok := obj.(*sphere.Sphere)
		if !ok {
			return false, errors.New("object list does not contain sphere type")
		}
		if (sphereObject.Material == s.Material) && matrix.Equals(sphereObject.Transformation, s.Transformation, 4, 4, 4, 4) {
			return true, nil
		}
	}
	return false, nil
}
