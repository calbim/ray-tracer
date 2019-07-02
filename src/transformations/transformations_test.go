package transformations

import (
	"math"
	"testing"

	"../matrix"
	"../tuple"
)

func TestTranslation(t *testing.T) {
	transform := NewTranslation(5, -3, 2)
	p := tuple.Point(-3, 4, 5)
	if !tuple.Equals(matrix.MultiplyWithTuple(transform, p), tuple.Point(2, 1, 7)) {
		t.Errorf("Transformation is wrong!")
	}
}

func TestTranslationInverse(t *testing.T) {
	transform := NewTranslation(5, -3, 2)
	inv, err := matrix.Inverse(transform, 4)
	if err != nil {
		t.Errorf("could not inverse transformation matrix")
	}
	p := tuple.Point(-3, 4, 5)
	if !tuple.Equals(matrix.MultiplyWithTuple(inv, p), tuple.Point(-8, 7, 3)) {
		t.Errorf("Transformation is wrong!")
	}
}

func TestTranslateVector(t *testing.T) {
	transform := NewTranslation(5, -3, 2)
	v := tuple.Vector(-3, 4, 5)
	if !tuple.Equals(matrix.MultiplyWithTuple(transform, v), v) {
		t.Errorf("Transformation of a vector returns vector itself")
	}
}

func TestScalePoint(t *testing.T) {
	transform := NewScaling(2, 3, 4)
	p := tuple.Point(-4, 6, 8)
	if !tuple.Equals(matrix.MultiplyWithTuple(transform, p), tuple.Point(-8, 18, 32)) {
		t.Errorf("Scaling is incorrect")
	}
}

func TestScaleVector(t *testing.T) {
	transform := NewScaling(2, 3, 4)
	v := tuple.Vector(-4, 6, 8)
	if !tuple.Equals(matrix.MultiplyWithTuple(transform, v), tuple.Vector(-8, 18, 32)) {
		t.Errorf("Scaling is incorrect")
	}
}

func TestScalingInverse(t *testing.T) {
	transform := NewScaling(2, 3, 4)
	inv, err := matrix.Inverse(transform, 4)
	if err != nil {
		t.Errorf("could not inverse transformation matrix")
	}
	v := tuple.Vector(-4, 6, 8)
	if !tuple.Equals(matrix.MultiplyWithTuple(inv, v), tuple.Vector(-2, 2, 2)) {
		t.Errorf("Scaling is incorrect!")
	}
}

func TestReflection(t *testing.T) {
	transform := NewScaling(-1, 1, 1)
	p := tuple.Point(2, 3, 4)
	if !tuple.Equals(matrix.MultiplyWithTuple(transform, p), tuple.Point(-2, 3, 4)) {
		t.Errorf("Reflection is incorrect!")
	}
}

func TestXRotation(t *testing.T) {
	p := tuple.Point(0, 1, 0)
	halfQuarter := RotationX(Pi / 4)
	fullQuarter := RotationX(Pi / 2)
	if !tuple.Equals(matrix.MultiplyWithTuple(halfQuarter, p), tuple.Point(0, math.Sqrt(2)/2, math.Sqrt(2)/2)) {
		t.Errorf("rotation for p %v is incorrect", p)
	}
	if !tuple.Equals(matrix.MultiplyWithTuple(fullQuarter, p), tuple.Point(0, 0, 1)) {
		t.Errorf("rotation for p %v is incorrect", p)
	}
}

func TestXRotationInverse(t *testing.T) {
	p := tuple.Point(0, 1, 0)
	halfQuarter := RotationX(Pi / 4)
	inv, err := matrix.Inverse(halfQuarter, 4)
	if err != nil {
		t.Errorf("could not compute inverse")
	}
	if !tuple.Equals(matrix.MultiplyWithTuple(inv, p), tuple.Point(0, math.Sqrt(2)/2, -math.Sqrt(2)/2)) {
		t.Errorf("rotation for p %v is incorrect", p)
	}
}

func TestYRotation(t *testing.T) {
	p := tuple.Point(0, 0, 1)
	halfQuarter := RotationY(Pi / 4)
	fullQuarter := RotationY(Pi / 2)
	if !tuple.Equals(matrix.MultiplyWithTuple(halfQuarter, p), tuple.Point(math.Sqrt(2)/2, 0, math.Sqrt(2)/2)) {
		t.Errorf("rotation for p %v is incorrect", p)
	}
	if !tuple.Equals(matrix.MultiplyWithTuple(fullQuarter, p), tuple.Point(1, 0, 0)) {
		t.Errorf("rotation for p %v is incorrect", p)
	}
}

func TestZRotation(t *testing.T) {
	p := tuple.Point(0, 1, 0)
	halfQuarter := RotationZ(Pi / 4)
	fullQuarter := RotationZ(Pi / 2)
	if !tuple.Equals(matrix.MultiplyWithTuple(halfQuarter, p), tuple.Point(-math.Sqrt(2)/2, math.Sqrt(2)/2, 0)) {
		t.Errorf("rotation for p %v is incorrect", p)
	}
	if !tuple.Equals(matrix.MultiplyWithTuple(fullQuarter, p), tuple.Point(-1, 0, 0)) {
		t.Errorf("rotation for p %v is incorrect", p)
	}
}

func TestShearing(t *testing.T) {
	transform := NewShearing(1, 0, 0, 0, 0, 0)
	p := tuple.Point(2, 3, 4)
	if !tuple.Equals(matrix.MultiplyWithTuple(transform, p), tuple.Point(5, 3, 4)) {
		t.Errorf("shearing transformation'moving x in proportion to y' for p %v is incorrect", p)
	}
	transform = NewShearing(0, 1, 0, 0, 0, 0)
	p = tuple.Point(2, 3, 4)
	if !tuple.Equals(matrix.MultiplyWithTuple(transform, p), tuple.Point(6, 3, 4)) {
		t.Errorf("shearing transformation for p %v is incorrect", p)
	}
	transform = NewShearing(0, 0, 1, 0, 0, 0)
	p = tuple.Point(2, 3, 4)
	if !tuple.Equals(matrix.MultiplyWithTuple(transform, p), tuple.Point(2, 5, 4)) {
		t.Errorf("shearing transformation for p %v is incorrect", p)
	}
	transform = NewShearing(0, 0, 0, 1, 0, 0)
	p = tuple.Point(2, 3, 4)
	if !tuple.Equals(matrix.MultiplyWithTuple(transform, p), tuple.Point(2, 7, 4)) {
		t.Errorf("shearing transformation for p %v is incorrect", p)
	}
	transform = NewShearing(0, 0, 0, 0, 1, 0)
	p = tuple.Point(2, 3, 4)
	if !tuple.Equals(matrix.MultiplyWithTuple(transform, p), tuple.Point(2, 3, 6)) {
		t.Errorf("shearing transformation for p %v is incorrect", p)
	}
	transform = NewShearing(0, 0, 0, 0, 0, 1)
	p = tuple.Point(2, 3, 4)
	if !tuple.Equals(matrix.MultiplyWithTuple(transform, p), tuple.Point(2, 3, 7)) {
		t.Errorf("shearing transformation for p %v is incorrect", p)
	}
}

func TestIndividualTransformation(t *testing.T) {
	p := tuple.Point(1, 0, 1)
	A := RotationX(Pi / 2)
	B := NewScaling(5, 5, 5)
	C := NewTranslation(10, 5, 7)
	p2 := matrix.MultiplyWithTuple(A, p)
	if !tuple.Equals(p2, tuple.Point(1, -1, 0)) {
		t.Errorf("rotation is incorrect")
	}
	p3 := matrix.MultiplyWithTuple(B, p2)
	if !tuple.Equals(p3, tuple.Point(5, -5, 0)) {
		t.Errorf("scaling is incorrect")
	}
	p4 := matrix.MultiplyWithTuple(C, p3)
	if !tuple.Equals(p4, tuple.Point(15, 0, 7)) {
		t.Errorf("translation is incorrect")
	}
}

func TestChainingTransformationsOrder(t *testing.T) {
	p := tuple.Point(1, 0, 1)
	A := RotationX(Pi / 2)
	B := NewScaling(5, 5, 5)
	C := NewTranslation(10, 5, 7)
	T := matrix.Multiply(matrix.Multiply(C, B), A)
	if !tuple.Equals(matrix.MultiplyWithTuple(T, p), tuple.Point(15, 0, 7)) {
		t.Errorf("chained transformations are in the wrong order")
	}
}
