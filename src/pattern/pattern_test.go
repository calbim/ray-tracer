package pattern

import (
	"testing"

	"github.com/calbim/ray-tracer/src/matrix"

	"github.com/calbim/ray-tracer/src/color"
	"github.com/calbim/ray-tracer/src/transforms"
	"github.com/calbim/ray-tracer/src/tuple"
)

func TestStripePattern(t *testing.T) {
	pattern := NewStripe(color.White, color.Black)
	if !pattern.a.Equals(color.White) {
		t.Errorf("wanted stripe a=%v, got %v", color.White, pattern.a)
	}
	if !pattern.b.Equals(color.Black) {
		t.Errorf("wanted stripe a=%v, got %v", color.Black, pattern.b)
	}
}

func TestStripeIsConstantInY(t *testing.T) {
	pattern := NewStripe(color.White, color.Black)
	if !pattern.PatternAt(tuple.Point(0, 0, 0)).Equals(color.White) {
		t.Errorf("wanted stripe at point %v=%v, got %v", tuple.Point(0, 0, 0), color.White, pattern.PatternAt(tuple.Point(0, 0, 0)))
	}
	if !pattern.PatternAt(tuple.Point(0, 1, 0)).Equals(color.White) {
		t.Errorf("wanted stripe at point %v=%v, got %v", tuple.Point(0, 1, 0), color.White, pattern.PatternAt(tuple.Point(0, 1, 0)))
	}
	if !pattern.PatternAt(tuple.Point(0, 2, 0)).Equals(color.White) {
		t.Errorf("wanted stripe at point %v=%v, got %v", tuple.Point(0, 2, 0), color.White, pattern.PatternAt(tuple.Point(0, 2, 0)))
	}
}

func TestStripeIsConstantInZ(t *testing.T) {
	pattern := NewStripe(color.White, color.Black)
	if !pattern.PatternAt(tuple.Point(0, 0, 0)).Equals(color.White) {
		t.Errorf("wanted stripe at point %v=%v, got %v", tuple.Point(0, 0, 0), color.White, pattern.PatternAt(tuple.Point(0, 0, 0)))
	}
	if !pattern.PatternAt(tuple.Point(0, 0, 1)).Equals(color.White) {
		t.Errorf("wanted stripe at point %v=%v, got %v", tuple.Point(0, 0, 1), color.White, pattern.PatternAt(tuple.Point(0, 0, 1)))
	}
	if !pattern.PatternAt(tuple.Point(0, 0, 2)).Equals(color.White) {
		t.Errorf("wanted stripe at point %v=%v, got %v", tuple.Point(0, 0, 2), color.White, pattern.PatternAt(tuple.Point(0, 0, 2)))
	}
}

func TestStripeAlternatesInX(t *testing.T) {
	pattern := NewStripe(color.White, color.Black)
	if !pattern.PatternAt(tuple.Point(0, 0, 0)).Equals(color.White) {
		t.Errorf("wanted stripe at point %v=%v, got %v", tuple.Point(0, 0, 0), color.White, pattern.PatternAt(tuple.Point(0, 0, 0)))
	}
	if !pattern.PatternAt(tuple.Point(0.9, 0, 0)).Equals(color.White) {
		t.Errorf("wanted stripe at point %v=%v, got %v", tuple.Point(0.9, 0, 0), color.White, pattern.PatternAt(tuple.Point(0.9, 0, 0)))
	}
	if !pattern.PatternAt(tuple.Point(1, 0, 0)).Equals(color.Black) {
		t.Errorf("wanted stripe at point %v=%v, got %v", tuple.Point(1, 0, 0), color.Black, pattern.PatternAt(tuple.Point(1, 0, 0)))
	}
	if !pattern.PatternAt(tuple.Point(-0.1, 0, 0)).Equals(color.Black) {
		t.Errorf("wanted stripe at point %v=%v, got %v", tuple.Point(-0.1, 0, 0), color.Black, pattern.PatternAt(tuple.Point(-0.1, 0, 0)))
	}
	if !pattern.PatternAt(tuple.Point(-1, 0, 0)).Equals(color.Black) {
		t.Errorf("wanted stripe at point %v=%v, got %v", tuple.Point(-1, 0, 0), color.Black, pattern.PatternAt(tuple.Point(-1, 0, 0)))
	}
	if !pattern.PatternAt(tuple.Point(-1.1, 0, 0)).Equals(color.White) {
		t.Errorf("wanted stripe at point %v=%v, got %v", tuple.Point(-1.1, 0, 0), color.White, pattern.PatternAt(tuple.Point(-1.1, 0, 0)))
	}
}

func TestStripesWithObjectTransformation(t *testing.T) {
	o := NewObject()
	o.Transform = transforms.Scaling(2, 2, 2)
	p := NewStripe(color.White, color.Black)
	c := AtObject(p, o, tuple.Point(1.5, 0, 0))
	if !c.Equals(color.White) {
		t.Errorf("wanted c=%v, got %v", color.White, c)
	}
}

func TestStripesWithPatternTransformation(t *testing.T) {
	o := NewObject()
	p := NewStripe(color.White, color.Black)
	p.Transform = transforms.Scaling(2, 2, 2)
	c := AtObject(p, o, tuple.Point(1.5, 0, 0))
	if !c.Equals(color.White) {
		t.Errorf("wanted c=%v, got %v", color.White, c)
	}
}

func TestStripesWithObjectAndPatternTransformation(t *testing.T) {
	o := NewObject()
	o.Transform = transforms.Scaling(2, 2, 2)
	p := NewStripe(color.White, color.Black)
	p.Transform = transforms.Translation(0.5, 0, 0)
	c := AtObject(p, o, tuple.Point(2.5, 0, 0))
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

func (tp *TestPattern) PatternAtObject(o Object, point tuple.Tuple) color.Color {
	oInv, _ := o.Transform.Inverse()
	point = oInv.MultiplyTuple(point)
	pInv, _ := tp.Transform.Inverse()
	point = pInv.MultiplyTuple(point)
	return color.New(point.X, point.Y, point.Z)
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

func TestPatternWithObjectTransformation(t *testing.T) {
	o := NewObject()
	o.Transform = transforms.Scaling(2, 2, 2)
	p := NewTestPattern()
	c := p.PatternAtObject(o, tuple.Point(2, 3, 4))
	if !c.Equals(color.New(1, 1.5, 2)) {
		t.Errorf("wanted color=%v, got %v", color.New(1, 1.5, 2), c)
	}
}

func TestPatternWithPatterntTransformation(t *testing.T) {
	o := NewObject()
	p := NewTestPattern()
	p.SetTransform(transforms.Scaling(2, 2, 2))
	c := p.PatternAtObject(o, tuple.Point(2, 3, 4))
	if !c.Equals(color.New(1, 1.5, 2)) {
		t.Errorf("wanted color=%v, got %v", color.New(1, 1.5, 2), c)
	}
}

func TestPatternWithObjectAndPatterntTransformation(t *testing.T) {
	o := NewObject()
	o.SetTransform(transforms.Scaling(2, 2, 2))
	p := NewTestPattern()
	p.SetTransform(transforms.Translation(0.5, 1, 1.5))
	c := p.PatternAtObject(o, tuple.Point(2.5, 3, 3.5))
	if !c.Equals(color.New(0.75, 0.5, 0.25)) {
		t.Errorf("wanted color=%v, got %v", color.New(0.75, 0.5, 0.25), c)
	}
}

func TestGradientLinearlyInterpolatesBetweenColors(t *testing.T) {
	p := NewGradient(color.White, color.Black)
	c := p.PatternAt(tuple.Point(0, 0, 0))
	if !c.Equals(color.White) {
		t.Errorf("wanted color to be %v, got %v", color.White, c)
	}
	c = p.PatternAt(tuple.Point(0.25, 0, 0))
	if !c.Equals(color.New(0.75, 0.75, 0.75)) {
		t.Errorf("wanted color to be %v, got %v", color.New(0.75, 0.75, 0.75), c)
	}
	c = p.PatternAt(tuple.Point(0.5, 0, 0))
	if !c.Equals(color.New(0.5, 0.5, 0.5)) {
		t.Errorf("wanted color to be %v, got %v", color.New(0.5, 0.5, 0.5), c)
	}
	c = p.PatternAt(tuple.Point(0.75, 0, 0))
	if !c.Equals(color.New(0.25, 0.25, 0.25)) {
		t.Errorf("wanted color to be %v, got %v", color.New(0.25, 0.25, 0.25), c)
	}
}

func TestRingExtendsInXAndZ(t *testing.T) {
	p := NewRing(color.White, color.Black)
	c := p.PatternAt(tuple.Point(0, 0, 0))
	if !c.Equals(color.White) {
		t.Errorf("wanted color to be %v, got %v", color.White, c)
	}
	c = p.PatternAt(tuple.Point(1, 0, 0))
	if !c.Equals(color.Black) {
		t.Errorf("wanted color to be %v, got %v", color.Black, c)
	}
	c = p.PatternAt(tuple.Point(0, 0, 1))
	if !c.Equals(color.Black) {
		t.Errorf("wanted color to be %v, got %v", color.Black, c)
	}
	c = p.PatternAt(tuple.Point(0.708, 0, 0.708))
	if !c.Equals(color.Black) {
		t.Errorf("wanted color to be %v, got %v", color.Black, c)
	}
}
