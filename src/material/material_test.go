package material

import (
	"math"
	"testing"

	"github.com/calbim/ray-tracer/src/light"
	"github.com/calbim/ray-tracer/src/tuple"
)

func TestMaterial(t *testing.T) {
	m := New()
	if m.Color != tuple.Color(1, 1, 1) {
		t.Errorf("Material colour is %v, should be %v", m.Color, tuple.Color(1, 1, 1))
	}
	if m.Ambient != 0.1 {
		t.Errorf("Material ambient is %v, should be %v", m.Ambient, 0.1)
	}
	if m.Diffuse != 0.9 {
		t.Errorf("Material diffuse is %v, should be %v", m.Diffuse, 0.9)
	}
	if m.Specular != 0.9 {
		t.Errorf("Material specular is %v, should be %v", m.Diffuse, 0.9)
	}
	if m.Shininess != 200 {
		t.Errorf("Material shininess is %v, should be %v", m.Shininess, 200)
	}
}

func TestLightingEyeBetweenLightAndSurface(t *testing.T) {
	m := New()
	position := tuple.Point(0, 0, 0)
	light := light.PointLight{Intensity: tuple.Color(1, 1, 1), Position: tuple.Point(0, 0, -10)}
	eyev := tuple.Vector(0, 0, -1)
	normalv := tuple.Vector(0, 0, -1)
	result := Lighting(m, light, position, eyev, normalv)
	if !tuple.Equals(result, tuple.Color(1.9, 1.9, 1.9)) {
		t.Errorf("Lighting is %v, should be %v", result, tuple.Color(1.9, 1.9, 1.9))
	}
}
func TestLightingEyeOffset45BetweenLightAndSurface(t *testing.T) {
	m := New()
	position := tuple.Point(0, 0, 0)
	light := light.PointLight{Intensity: tuple.Color(1, 1, 1), Position: tuple.Point(0, 0, -10)}
	eyev := tuple.Vector(0, math.Sqrt(2)/2, -math.Sqrt(2)/2)
	normalv := tuple.Vector(0, 0, -1)
	result := Lighting(m, light, position, eyev, normalv)
	if !tuple.Equals(result,tuple.Color(1, 1, 1)) {
		t.Errorf("Lighting is %v, should be %v", result, tuple.Color(1.9, 1.9, 1.9))
	}
}

func TestLightingOffset45EyeBetweenLightAndSurface(t *testing.T) {
	m := New()
	position := tuple.Point(0, 0, 0)
	light := light.PointLight{Intensity: tuple.Color(1, 1, 1), Position: tuple.Point(0, 10, -10)}
	eyev := tuple.Vector(0, 0, -1)
	normalv := tuple.Vector(0, 0, -1)
	result := Lighting(m, light, position, eyev, normalv)
	if !tuple.Equals(result, tuple.Color(0.7364, 0.7364, 0.7364)) {
		t.Errorf("Lighting is %v, should be %v", result, tuple.Color(0.7364, 0.7364, 0.7364))
	}
}

func TestLightingEyeInPathOfReflectionVector(t *testing.T) {
	m := New()
	position := tuple.Point(0, 0, 0)
	light := light.PointLight{Intensity: tuple.Color(1, 1, 1), Position: tuple.Point(0, 10, -10)}
	eyev := tuple.Vector(0, -math.Sqrt(2)/2, -math.Sqrt(2)/2)
	normalv := tuple.Vector(0, 0, -1)
	result := Lighting(m, light, position, eyev, normalv)
	if !tuple.Equals(result, tuple.Color(1.6364, 1.6364, 1.6364)) {
		t.Errorf("Lighting is %v, should be %v", result, tuple.Color(1.6364, 1.6364, 1.6364))
	}
}

func TestLightingBehindSurface(t *testing.T) {
	m := New()
	position := tuple.Point(0, 0, 0)
	light := light.PointLight{Intensity: tuple.Color(1, 1, 1), Position: tuple.Point(0, 0, 10)}
	eyev := tuple.Vector(0, 0, -1)
	normalv := tuple.Vector(0, 0, -1)
	result := Lighting(m, light, position, eyev, normalv)
	if !tuple.Equals(result, tuple.Color(0.1, 0.1, 0.1)) {
		t.Errorf("Lighting is %v, should be %v", result, tuple.Color(0.1, 0.1, 0.1))
	}
}
