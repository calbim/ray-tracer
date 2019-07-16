package world

import (
	"errors"

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
		Light: &light.PointLight{Intensity: tuple.Point(-10, -10, -10), Position: tuple.Point(1, 1, 1)},
	}
	m := material.Material{Color: tuple.Color(0.8, 1.0, 0.6), Diffuse: 0.7, Specular: 0.2}
	s1, err := sphere.New()
	if err != nil {
		return nil, errors.New("Could not create a sphere")
	}
	s1.Material = m
	s2, err := sphere.New()
	if err != nil {
		return nil, errors.New("Could not create a sphere")
	}
	s2.Material = m
	s2.SetTransform(transformations.NewScaling(0.5, 0.5, 0.5))
	w.Objects = []intersections.Object{s1, s2}
	return &w, nil
}
