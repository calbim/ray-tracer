package world

import (
	"sort"

	"github.com/calbim/ray-tracer/src/color"
	"github.com/calbim/ray-tracer/src/light"
	"github.com/calbim/ray-tracer/src/material"
	"github.com/calbim/ray-tracer/src/ray"
	"github.com/calbim/ray-tracer/src/shape"
	"github.com/calbim/ray-tracer/src/transforms"
	"github.com/calbim/ray-tracer/src/tuple"
)

//World is a world object containing a light source and objects
type World struct {
	Objects []shape.Shape
	Light   *light.Light
}

//Object interface
type Object interface {
}

//New returns a default world object
func New() World {
	l := light.PointLight(tuple.Point(-10, 10, -10), color.New(1, 1, 1))

	w := World{
		Light: &l,
	}
	s1 := shape.NewSphere()
	s2 := shape.NewSphere()
	m := material.New()
	s1M := m
	s1M.Color = color.New(0.8, 1.0, 0.6)
	s1M.Diffuse = 0.7
	s1M.Specular = 0.2
	s1.Material = s1M
	s2.Material = m
	s2.SetTransform(transforms.Scaling(0.5, 0.5, 0.5))
	w.Objects = []shape.Shape{s1, s2}
	return w
}

//Intersect returns a list of intersections in sorted order when a ray passes through a world
func (w World) Intersect(r ray.Ray) []shape.Intersection {
	list := []shape.Intersection{}
	for _, o := range w.Objects {
		intersection := o.Intersect(r)
		list = append(list, intersection...)
	}
	sort.Sort(ByIntersectionValue(list))
	return list
}

//ShadeHit computes the color of an intersection in a world
func (w *World) ShadeHit(comp shape.Computation) color.Color {
	return comp.Object.GetMaterial().Lighting(*w.Light,
		comp.Point, comp.Eyev, comp.Normal)
}

//ColorAt returns the color of the intersection of ray r with world w
func (w *World) ColorAt(r ray.Ray) color.Color {
	ints := w.Intersect(r)
	hit := shape.Hit(ints)
	if hit == nil {
		return color.Black
	}
	comp := hit.PrepareComputations(r)
	shade := w.ShadeHit(comp)
	return shade
}

// ByIntersectionValue implements sort.Interface for []Intersection based on the value
type ByIntersectionValue []shape.Intersection

func (a ByIntersectionValue) Len() int           { return len(a) }
func (a ByIntersectionValue) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByIntersectionValue) Less(i, j int) bool { return a[i].Value < a[j].Value }
