package world

import (
	"testing"

	"github.com/calbim/ray-tracer/src/ray"

	"github.com/calbim/ray-tracer/src/color"
	"github.com/calbim/ray-tracer/src/light"
	"github.com/calbim/ray-tracer/src/material"
	"github.com/calbim/ray-tracer/src/shape"
	"github.com/calbim/ray-tracer/src/transforms"
	"github.com/calbim/ray-tracer/src/tuple"
)

func TestCreateWorld(t *testing.T) {
	w := World{}
	if w.Objects != nil || w.Light != nil {
		t.Errorf("wanted empty world to contain no objects or light source")
	}
}

func TestDefaultWorld(t *testing.T) {
	w := Default()
	light := light.PointLight(tuple.Point(-10, 10, -10), color.New(1, 1, 1))

	m := material.Material{Color: color.New(0.8, 1.0, 0.6), Diffuse: 0.7, Specular: 0.2, Ambient: 0.1, Shininess: 200}
	s1 := shape.NewSphere()
	s1.Material = m

	s2 := shape.NewSphere()
	s2.Material = material.New()
	s2.SetTransform(transforms.Scaling(0.5, 0.5, 0.5))

	if *w.Light != light {
		t.Errorf("wanted world light source to be %v, got %v", light, w.Light)
	}

	if !contains(w.Objects, s1) || !contains(w.Objects, s2) {
		t.Errorf("wanted world to contain objects s1 %v and s2 %v, got %v and %v", s1, s2, w.Objects[0], w.Objects[1])
	}
}

func TestIntersectWorld(t *testing.T) {
	w := Default()
	r := ray.New(tuple.Point(0.0, 0.0, -5.0), tuple.Vector(0.0, 0.0, 1.0))
	xs := w.Intersect(r)
	if len(xs) != 4 {
		t.Errorf("wanted %v points of intersections, got %v", 4, len(xs))
	}
	if xs[0].Value != 4 || xs[1].Value != 4.5 || xs[2].Value != 5.5 || xs[3].Value != 6 {
		t.Errorf("wanted interesection points to be 4,4.5,5,6, got %v", xs)
	}
}

func TestShadingIntersection(t *testing.T) {
	w := Default()
	r := ray.New(tuple.Point(0, 0, -5), tuple.Vector(0, 0, 1))
	s := w.Objects[0]
	i := &shape.Intersection{Value: 4.0, Object: s}
	comp := i.PrepareComputations(r)
	c := w.ShadeHit(comp)
	if !c.Equals(color.New(0.38066, 0.47583, 0.2855)) {
		t.Errorf("wanted hit shade=%v, got %v", color.New(0.38066, 0.47583, 0.2855), c)
	}
}

func TestShadingIntersectionInside(t *testing.T) {
	w := Default()
	l := light.PointLight(tuple.Point(0, 0.25, 0), color.New(1, 1, 1))
	w.Light = &l
	r := ray.New(tuple.Point(0, 0, 0), tuple.Vector(0, 0, 1))
	s := w.Objects[1]
	i := &shape.Intersection{Value: 0.5, Object: s}
	comp := i.PrepareComputations(r)
	c := w.ShadeHit(comp)
	if !(c.Equals(color.New(0.90498, 0.90498, 0.90498))) {
		t.Errorf("wanted hit shade=%v, got %v", color.New(0.90498, 0.90498, 0.90498), c)
	}
}

func TestColorWhenRayMisses(t *testing.T) {
	w := Default()
	r := ray.New(tuple.Point(0, 0, -5), tuple.Vector(0, 1, 0))
	c := w.ColorAt(r)
	if !c.Equals(color.New(0, 0, 0)) {
		t.Errorf("wanted color=%v, got %v", color.Black, c)
	}
}

func TestDefaultColorWhenRayHits(t *testing.T) {
	w := Default()
	r := ray.New(tuple.Point(0, 0, -5), tuple.Vector(0, 0, 1))
	c := w.ColorAt(r)
	if !c.Equals(color.New(0.38066, 0.47583, 0.2855)) {
		t.Errorf("wanted color=%v, got %v", color.New(0.38066, 0.47583, 0.2855), c)
	}
}

func TestColorWhenIntersectionIsBehindRay(t *testing.T) {
	w := Default()
	m := w.Objects[0].GetMaterial()
	m.Ambient = 1
	outer := w.Objects[0]
	outer.SetMaterial(*m)
	inner := w.Objects[1]
	mInner := inner.GetMaterial()
	mInner.Ambient = 1
	inner.SetMaterial(*mInner)

	r := ray.New(tuple.Point(0, 0, 0.075), tuple.Vector(0, 0, -1))
	c := w.ColorAt(r)
	if !c.Equals(inner.GetMaterial().Color) {
		t.Errorf("wanted color=%v got %v", inner.GetMaterial().Color, c)
	}
}

func contains(list []shape.Shape, s shape.Shape) bool {
	for _, obj := range list {
		trans := obj.GetTransform()
		if (*obj.GetMaterial() == *s.GetMaterial()) && trans.Equals(s.GetTransform()) {
			return true
		}
	}
	return false
}
