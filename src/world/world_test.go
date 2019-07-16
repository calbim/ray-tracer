package world

import (
	"errors"
	"testing"

	"github.com/calbim/ray-tracer/src/intersections"
	"github.com/calbim/ray-tracer/src/light"
	"github.com/calbim/ray-tracer/src/material"
	"github.com/calbim/ray-tracer/src/matrix"
	"github.com/calbim/ray-tracer/src/sphere"
	"github.com/calbim/ray-tracer/src/transformations"
	"github.com/calbim/ray-tracer/src/tuple"
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
	light := light.PointLight{Intensity: tuple.Point(-10, -10, -10), Position: tuple.Point(1, 1, 1)}
	m := material.Material{Color: tuple.Color(0.8, 1.0, 0.6), Diffuse: 0.7, Specular: 0.2}
	s1, err := sphere.New()
	if err != nil {
		t.Errorf("Could not create sphere s1")
	}
	s1.Material = m

	s2, err := sphere.New()
	if err != nil {
		t.Errorf("Could not create sphere s2")
	}
	s2.Material = m
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

func contains(list []intersections.Object, s *sphere.Sphere) (bool, error) {
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
