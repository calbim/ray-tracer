package world

import (
	"errors"
	"fmt"
	"sort"

	"github.com/calbim/ray-tracer/src/ray"

	"github.com/calbim/ray-tracer/src/intersections"
	"github.com/calbim/ray-tracer/src/light"
	"github.com/calbim/ray-tracer/src/material"
	"github.com/calbim/ray-tracer/src/sphere"
	"github.com/calbim/ray-tracer/src/transformations"
	"github.com/calbim/ray-tracer/src/tuple"
)

//World contains a set of objects and a light source
type World struct {
	Objects []intersections.Object
	Light   *light.PointLight
}

//NewDefault creates a world with a default config of one light source and two spheres
func NewDefault() (*World, error) {
	w := World{
		Light: &light.PointLight{Position: tuple.Point(-10, 10, -10), Intensity: tuple.Color(1, 1, 1)},
	}
	s1, err := sphere.New()
	if err != nil {
		return nil, errors.New("Could not create a sphere")
	}
	s2, err := sphere.New()
	if err != nil {
		return nil, errors.New("Could not create a sphere")
	}
	m := material.New()
	s1M := m
	s1M.Color = tuple.Color(0.8, 1.0, 0.6)
	s1M.Diffuse = 0.7
	s1M.Specular = 0.2
	s1.Material = s1M
	s2.Material = m
	s2.SetTransform(transformations.NewScaling(0.5, 0.5, 0.5))
	w.Objects = []intersections.Object{s1, s2}
	return &w, nil
}

// ByIntersectionValue implements sort.Interface for []Intersection based on
// the Value field.
type ByIntersectionValue []intersections.Intersection

func (a ByIntersectionValue) Len() int           { return len(a) }
func (a ByIntersectionValue) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByIntersectionValue) Less(i, j int) bool { return a[i].Value < a[j].Value }

//Intersect returns a list of intersections in sorted order when a ray passes through a world
func (w World) Intersect(r ray.Ray) ([]intersections.Intersection, error) {
	list := []intersections.Intersection{}
	for _, o := range w.Objects {
		intersection, err := intersections.Intersect(o, r)
		if err != nil {
			return nil, fmt.Errorf("Error while computing intersection for object %v", err)
		}
		list = append(list, intersection...)
	}
	sort.Sort(ByIntersectionValue(list))
	return list, nil
}

//ShadeHit computes the color of an intersection in a world
func ShadeHit(w World, comp intersections.Computation) tuple.Tuple {
	shadow, err := w.IsShadowed(comp.Overpoint)
	if err != nil {
		shadow = false
	}
	return material.Lighting(comp.Object.GetMaterial(), *w.Light,
		comp.Overpoint, comp.Eyev, comp.Normal, shadow)
}

//ColorAt returns the color of the intersection of ray r with world w
func ColorAt(w World, r ray.Ray) (*tuple.Tuple, error) {
	ints, err := w.Intersect(r)
	if err != nil {
		{
			return nil, errors.New("Error finding intersections of world with ray")
		}
	}
	hit := intersections.Hit(ints)
	color := tuple.Color(0, 0, 0)
	if hit == nil {
		return &color, nil
	}
	comps, err := intersections.PrepareComputations(*hit, r)
	if err != nil {
		return nil, fmt.Errorf("Error while preparing computations %v", err)
	}
	shade := ShadeHit(w, *comps)
	return &shade, nil
}

// IsShadowed determines if a point p lies within a shadow
func (w *World) IsShadowed(p tuple.Tuple) (bool, error) {
	dV := tuple.Subtract(w.Light.Position, p)
	d := tuple.Magnitude(dV)
	r := ray.Ray{
		Origin:    p,
		Direction: tuple.Normalize(dV),
	}
	is, err := w.Intersect(r)
	if err != nil {
		return false, fmt.Errorf("Error determining if point is shadowed due to error %v", err)
	}
	hit := intersections.Hit(is)
	if hit != nil && hit.Value < d {
		return true, nil
	}
	return false, nil
}
