package pattern_test

import (
	"testing"

	. "github.com/calbim/ray-tracer/src/pattern"

	"github.com/calbim/ray-tracer/src/light"
	"github.com/calbim/ray-tracer/src/material"

	"github.com/calbim/ray-tracer/src/tuple"
)

func TestCreateStripe(t *testing.T) {
	p := StripePattern(tuple.Black, tuple.White)
	if p.A != tuple.Black {
		t.Errorf("pattern.a should be black")
	}
	if p.B != tuple.White {
		t.Errorf("pattern.b should be white")
	}
}

func TestStripeIsConstantInY(t *testing.T) {
	p := StripePattern(tuple.White, tuple.Black)
	if StripeAt(p, tuple.Point(0, 0, 0)) != tuple.White {
		t.Errorf("Stripe color at (0,0,0) should be %v", tuple.White)
	}
	if StripeAt(p, tuple.Point(0, 1, 0)) != tuple.White {
		t.Errorf("Stripe color at (0,0,0) should be %v", tuple.White)
	}
	if StripeAt(p, tuple.Point(0, 2, 0)) != tuple.White {
		t.Errorf("Stripe color at (0,0,0) should be %v", tuple.White)
	}
}

func TestStripeIsConstantInZ(t *testing.T) {
	p := StripePattern(tuple.White, tuple.Black)
	if StripeAt(p, tuple.Point(0, 0, 0)) != tuple.White {
		t.Errorf("Stripe color at (0,0,0) should be %v", tuple.White)
	}
	if StripeAt(p, tuple.Point(0, 0, 1)) != tuple.White {
		t.Errorf("Stripe color at (0,0,1) should be %v", tuple.White)
	}
	if StripeAt(p, tuple.Point(0, 0, 2)) != tuple.White {
		t.Errorf("Stripe color at (0,0,2) should be %v", tuple.White)
	}
}

func TestStripePatternAlternatesInX(t *testing.T) {
	p := StripePattern(tuple.White, tuple.Black)
	if StripeAt(p, tuple.Point(0, 0, 0)) != tuple.White {
		t.Errorf("Stripe color at (0,0,0) should be %v", tuple.White)
	}
	if StripeAt(p, tuple.Point(0.9, 0, 0)) != tuple.White {
		t.Errorf("Stripe color at (0.9,0,0) should be %v", tuple.White)
	}
	if StripeAt(p, tuple.Point(1.0, 0, 0)) != tuple.Black {
		t.Errorf("Stripe color at (1.0,0,0) should be %v", tuple.Black)
	}
	if StripeAt(p, tuple.Point(-0.1, 0, 0)) != tuple.Black {
		t.Errorf("Stripe color at (-0.1,0,0) should be %v", tuple.Black)
	}
	if StripeAt(p, tuple.Point(-1, 0, 0)) != tuple.Black {
		t.Errorf("Stripe color at (-1,0,0) should be %v", tuple.Black)
	}
	if StripeAt(p, tuple.Point(-1.1, 0, 0)) != tuple.White {
		t.Errorf("Stripe color at (-1.1,0,0) should be %v", tuple.White)
	}
}

func TestLightingWithPattern(t *testing.T) {
	m := material.New()
	p := StripePattern(tuple.White, tuple.Black)
	m.Pattern = &p
	m.Ambient = 1
	m.Specular = 0
	m.Diffuse = 0
	eyeV := tuple.Vector(0, 0, -1)
	normalV := tuple.Vector(0, 0, -1)
	light := light.PointLight{Position: tuple.Point(0, 0, -10), Intensity: tuple.White}
	c1 := material.Lighting(m, light, tuple.Point(0.9, 0, 0), eyeV, normalV, false)
	c2 := material.Lighting(m, light, tuple.Point(1.1, 0, 0), eyeV, normalV, false)
	if c1 != tuple.White {
		t.Errorf("c1 should be %v but it is %v", tuple.White, c1)
	}
	if c2 != tuple.Black {
		t.Errorf("c2 should be %v but it is %v", tuple.Black, c2)
	}
}
