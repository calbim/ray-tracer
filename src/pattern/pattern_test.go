package pattern

import (
	"testing"

	"github.com/calbim/ray-tracer/src/matrix"

	"github.com/calbim/ray-tracer/src/color"
	"github.com/calbim/ray-tracer/src/transforms"
	"github.com/calbim/ray-tracer/src/tuple"
)

func TestStripePattern(t *testing.T) {
	pattern := Stripe(color.White, color.Black)
	if !pattern.a.Equals(color.White) {
		t.Errorf("wanted stripe a=%v, got %v", color.White, pattern.a)
	}
	if !pattern.b.Equals(color.Black) {
		t.Errorf("wanted stripe a=%v, got %v", color.Black, pattern.b)
	}
}

func TestStripeIsConstantInY(t *testing.T) {
	pattern := Stripe(color.White, color.Black)
	if !pattern.StripeAt(tuple.Point(0, 0, 0)).Equals(color.White) {
		t.Errorf("wanted stripe at point %v=%v, got %v", tuple.Point(0, 0, 0), color.White, pattern.StripeAt(tuple.Point(0, 0, 0)))
	}
	if !pattern.StripeAt(tuple.Point(0, 1, 0)).Equals(color.White) {
		t.Errorf("wanted stripe at point %v=%v, got %v", tuple.Point(0, 1, 0), color.White, pattern.StripeAt(tuple.Point(0, 1, 0)))
	}
	if !pattern.StripeAt(tuple.Point(0, 2, 0)).Equals(color.White) {
		t.Errorf("wanted stripe at point %v=%v, got %v", tuple.Point(0, 2, 0), color.White, pattern.StripeAt(tuple.Point(0, 2, 0)))
	}
}

func TestStripeIsConstantInZ(t *testing.T) {
	pattern := Stripe(color.White, color.Black)
	if !pattern.StripeAt(tuple.Point(0, 0, 0)).Equals(color.White) {
		t.Errorf("wanted stripe at point %v=%v, got %v", tuple.Point(0, 0, 0), color.White, pattern.StripeAt(tuple.Point(0, 0, 0)))
	}
	if !pattern.StripeAt(tuple.Point(0, 0, 1)).Equals(color.White) {
		t.Errorf("wanted stripe at point %v=%v, got %v", tuple.Point(0, 0, 1), color.White, pattern.StripeAt(tuple.Point(0, 0, 1)))
	}
	if !pattern.StripeAt(tuple.Point(0, 0, 2)).Equals(color.White) {
		t.Errorf("wanted stripe at point %v=%v, got %v", tuple.Point(0, 0, 2), color.White, pattern.StripeAt(tuple.Point(0, 0, 2)))
	}
}

func TestStripeAlternatesInX(t *testing.T) {
	pattern := Stripe(color.White, color.Black)
	if !pattern.StripeAt(tuple.Point(0, 0, 0)).Equals(color.White) {
		t.Errorf("wanted stripe at point %v=%v, got %v", tuple.Point(0, 0, 0), color.White, pattern.StripeAt(tuple.Point(0, 0, 0)))
	}
	if !pattern.StripeAt(tuple.Point(0.9, 0, 0)).Equals(color.White) {
		t.Errorf("wanted stripe at point %v=%v, got %v", tuple.Point(0.9, 0, 0), color.White, pattern.StripeAt(tuple.Point(0.9, 0, 0)))
	}
	if !pattern.StripeAt(tuple.Point(1, 0, 0)).Equals(color.Black) {
		t.Errorf("wanted stripe at point %v=%v, got %v", tuple.Point(1, 0, 0), color.Black, pattern.StripeAt(tuple.Point(1, 0, 0)))
	}
	if !pattern.StripeAt(tuple.Point(-0.1, 0, 0)).Equals(color.Black) {
		t.Errorf("wanted stripe at point %v=%v, got %v", tuple.Point(-0.1, 0, 0), color.Black, pattern.StripeAt(tuple.Point(-0.1, 0, 0)))
	}
	if !pattern.StripeAt(tuple.Point(-1, 0, 0)).Equals(color.Black) {
		t.Errorf("wanted stripe at point %v=%v, got %v", tuple.Point(-1, 0, 0), color.Black, pattern.StripeAt(tuple.Point(-1, 0, 0)))
	}
	if !pattern.StripeAt(tuple.Point(-1.1, 0, 0)).Equals(color.White) {
		t.Errorf("wanted stripe at point %v=%v, got %v", tuple.Point(-1.1, 0, 0), color.White, pattern.StripeAt(tuple.Point(-1.1, 0, 0)))
	}
}

func TestStripesWithObjectTransformation(t *testing.T) {
	o := NewObject()
	o.Transform = transforms.Scaling(2, 2, 2)
	p := Stripe(color.White, color.Black)
	c := p.StripeAtObject(o, tuple.Point(1.5, 0, 0))
	if !c.Equals(color.White) {
		t.Errorf("wanted c=%v, got %v", color.White, c)
	}
}

func TestStripesWithPatternTransformation(t *testing.T) {
	o := NewObject()
	p := Stripe(color.White, color.Black)
	p.Transform = transforms.Scaling(2, 2, 2)
	c := p.StripeAtObject(o, tuple.Point(1.5, 0, 0))
	if !c.Equals(color.White) {
		t.Errorf("wanted c=%v, got %v", color.White, c)
	}
}

func TestStripesWithObjectAndPatternTransformation(t *testing.T) {
	o := NewObject()
	o.Transform = transforms.Scaling(2, 2, 2)
	p := Stripe(color.White, color.Black)
	p.Transform = transforms.Translation(0.5, 0, 0)
	c := p.StripeAtObject(o, tuple.Point(2.5, 0, 0))
	if !c.Equals(color.White) {
		t.Errorf("wanted c=%v, got %v", color.White, c)
	}
}

type TestPattern struct {
	Transform *matrix.Matrix
}

func NewTestPattern() *TestPattern {
	return &TestPattern{
		Transform: matrix.Identity,
	}
}
func (tp *TestPattern) GetTransform() *matrix.Matrix {
	return tp.Transform
}

func (tp *TestPattern) SetTransform(transform *matrix.Matrix) {
	tp.Transform = transform
}
func TestDefaultPatternTransformation(t *testing.T) {
	p := NewTestPattern()
	if !p.GetTransform().Equals(matrix.Identity) {
		t.Errorf("wanted pattern default transformation to be %v, got %v", matrix.Identity, p.GetTransform())
	}
}

func TestAssignPatternTransformation(t *testing.T) {
	p := NewTestPattern()
	p.SetTransform(transforms.Translation(1, 2, 3))
	if !p.GetTransform().Equals(transforms.Translation(1, 2, 3)) {
		t.Errorf("wanted assigned pattern transformation to be %v, got %v", transforms.Translation(1, 2, 3), p.GetTransform())
	}
}
