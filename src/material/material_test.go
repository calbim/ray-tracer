package material

import (
	"math"
	"testing"

	"github.com/calbim/ray-tracer/src/pattern"
	"github.com/calbim/ray-tracer/src/tuple"

	"github.com/calbim/ray-tracer/src/color"
	"github.com/calbim/ray-tracer/src/light"
)

func TestMaterial(t *testing.T) {
	m := New()
	if !m.Color.Equals(color.New(1, 1, 1)) {
		t.Errorf("wanted color=%v, should be %v", color.New(1, 1, 1), m.Color)
	}
	if m.Ambient != 0.1 {
		t.Errorf("wanted ambient=%v, should be %v", 0.1, m.Ambient)
	}
	if m.Diffuse != 0.9 {
		t.Errorf("wanted diffuse=%v, should be %v", 0.9, m.Diffuse)
	}
	if m.Specular != 0.9 {
		t.Errorf("wanted specular=%v, should be %v", 0.9, m.Specular)
	}
	if m.Shininess != 200 {
		t.Errorf("wanted shininess=%v, should be %v", 200, m.Shininess)
	}
}

func TestLightingEyeBetweenLightAndSurface(t *testing.T) {
	m := New()
	position := tuple.Point(0, 0, 0)
	light := light.PointLight(tuple.Point(0, 0, -10), color.White)
	eyev := tuple.Vector(0, 0, -1)
	normalv := tuple.Vector(0, 0, -1)
	result := m.Lighting(pattern.NewObject(), light, position, eyev, normalv, false)
	if !result.Equals(color.New(1.9, 1.9, 1.9)) {
		t.Errorf("wanted lighting=%v, got %v", color.New(1.9, 1.9, 1.9), result)
	}
}

func TestLightingEyeOffset45BetweenLightAndSurface(t *testing.T) {
	m := New()
	position := tuple.Point(0, 0, 0)
	light := light.PointLight(tuple.Point(0, 0, -10), color.White)
	eyev := tuple.Vector(0, math.Sqrt(2)/2, -math.Sqrt(2)/2)
	normalv := tuple.Vector(0, 0, -1)
	result := m.Lighting(pattern.NewObject(), light, position, eyev, normalv, false)
	if !result.Equals(color.White) {
		t.Errorf("wanted lighting=%v, got %v", color.White, result)
	}
}

func TestLightingEyeInPathOfReflectionVector(t *testing.T) {
	m := New()
	position := tuple.Point(0, 0, 0)
	light := light.PointLight(tuple.Point(0, 10, -10), color.New(1, 1, 1))
	eyev := tuple.Vector(0, -math.Sqrt(2)/2, -math.Sqrt(2)/2)
	normalv := tuple.Vector(0, 0, -1)
	result := m.Lighting(pattern.NewObject(), light, position, eyev, normalv, false)
	if !result.Equals(color.New(1.6364, 1.6364, 1.6364)) {
		t.Errorf("wanted lighting=%v, got %v", color.New(1.6364, 1.6364, 1.6364), result)
	}
}

func TestLightingBehindSurface(t *testing.T) {
	m := New()
	position := tuple.Point(0, 0, 0)
	light := light.PointLight(tuple.Point(0, 0, 10), color.New(1, 1, 1))
	eyev := tuple.Vector(0, 0, -1)
	normalv := tuple.Vector(0, 0, -1)
	result := m.Lighting(pattern.NewObject(), light, position, eyev, normalv, false)
	if !result.Equals(color.New(0.1, 0.1, 0.1)) {
		t.Errorf("wanted lighting=%v, got %v", color.New(0.1, 0.1, 0.1), result)
	}
}

func TestLightingSurfaceInShadow(t *testing.T) {
	m := New()
	position := tuple.Point(0, 0, 0)
	eyev := tuple.Vector(0, 0, -1)
	normalv := tuple.Vector(0, 0, -1)
	light := light.PointLight(tuple.Point(0, 0, -10), color.New(1, 1, 1))
	inShadow := true
	result := m.Lighting(pattern.NewObject(), light, position, eyev, normalv, inShadow)
	if !result.Equals(color.New(0.1, 0.1, 0.1)) {
		t.Errorf("wanted lighting=%v, got %v", result, color.New(0.1, 0.1, 0.1))
	}
}

func TestLightingWithPattern(t *testing.T) {
	m := New()
	m.SetPattern(pattern.NewStripe(color.White, color.Black))
	m.Ambient = 1
	m.Diffuse = 0
	m.Specular = 0
	eyev := tuple.Vector(0, 0, -1)
	normalv := tuple.Vector(0, 0, -1)
	l := light.PointLight(tuple.Point(0, 0, -10), color.White)
	c1 := m.Lighting(pattern.NewObject(), l, tuple.Point(0.9, 0, 0), eyev, normalv, false)
	c2 := m.Lighting(pattern.NewObject(), l, tuple.Point(1.1, 0, 0), eyev, normalv, false)
	if !c1.Equals(color.White) {
		t.Errorf("wanted c1=%v, got %v", color.White, c1)
	}
	if !c2.Equals(color.Black) {
		t.Errorf("wanted c2=%v, got %v", color.Black, c2)
	}
}

func TestLightingWithPatternApplied(t *testing.T) {
	m := New()
	m.SetPattern(pattern.NewStripe(color.White, color.Black))
	m.Ambient = 1
	m.Diffuse = 0
	m.Specular = 0
	eyev := tuple.Vector(0, 0, -1)
	normalv := tuple.Vector(0, 0, -1)
	l := light.PointLight(tuple.Point(0, 0, -10), color.White)
	c1 := m.Lighting(pattern.NewObject(), l, tuple.Point(0.9, 0, 0), eyev, normalv, false)
	c2 := m.Lighting(pattern.NewObject(), l, tuple.Point(1.1, 0, 0), eyev, normalv, false)
	if !c1.Equals(color.White) {
		t.Errorf("wanted c1=%v, got %v", color.White, c1)
	}
	if !c2.Equals(color.Black) {
		t.Errorf("wanted c2=%v, got %v", color.White, c2)
	}
}
