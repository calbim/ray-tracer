package tuple

import (
	"math"
	"testing"
)

func TestPoint(t *testing.T) {
	tup := Tuple{4.3, -4.2, 3.1, 1.0}
	if tup.X != 4.3 {
		t.Errorf("wanted x to be %v, got %v", 4.3, tup.X)
	}
	if tup.Y != -4.2 {
		t.Errorf("wanted y to be %v, got %v", -4.2, tup.Y)
	}
	if tup.Z != 3.1 {
		t.Errorf("wanted z to be %v, got %v", 3.1, tup.Z)
	}
	if tup.W != 1.0 {
		t.Errorf("wanted w to be %v, got %v", 1, tup.W)
	}
	if !tup.isPoint() {
		t.Error("wanted t to be a point, but it is a vector")
	}
}

func TestVector(t *testing.T) {
	tup := Tuple{4.3, -4.2, 3.1, 0.0}
	if tup.X != 4.3 {
		t.Errorf("wanted x to be %v, got %v", 4.3, tup.X)
	}
	if tup.Y != -4.2 {
		t.Errorf("wanted y to be %v, got %v", -4.2, tup.Y)
	}
	if tup.Z != 3.1 {
		t.Errorf("wanted z to be %v, got %v", 3.1, tup.Z)
	}
	if tup.W != 0.0 {
		t.Errorf("wanted w to be %v, got %v", 0.0, tup.W)
	}
	if !tup.isVector() {
		t.Error("wanted t to be a vector, but it is a point")
	}
}

func TestPointW(t *testing.T) {
	tup := Point(4, -4, 3)
	if (tup != Tuple{4, -4, 3, 1}) {
		t.Errorf("wanted %v, got %v", Tuple{4, -4, 3, 1}, tup)
	}
}

func TestVectorW(t *testing.T) {
	tup := Vector(4, -4, 3)
	if (tup != Tuple{4, -4, 3, 0}) {
		t.Errorf("wanted %v, got %v", Tuple{4, -4, 3, 0}, tup)
	}
}

func TestAdd(t *testing.T) {
	t1 := Tuple{3, -2, 5, 1}
	t2 := Tuple{-2, 3, 1, 0}
	if (t1.Add(t2) != Tuple{1, 1, 6, 1}) {
		t.Errorf("wanted t1+t2 = %v, got %v", Tuple{1, 1, 6, 1}, t1.Add(t2))
	}
}

func TestSubtract(t *testing.T) {
	t1 := Point(3, 2, 1)
	t2 := Point(5, 6, 7)
	if t1.Subtract(t2) != Vector(-2, -4, -6) {
		t.Errorf("wanted t1-t2 = %v, got %v", Vector(-2, -4, -6), t1.Subtract(t2))
	}
}

func TestSubtractVectorFromPoint(t *testing.T) {
	p := Point(3, 2, 1)
	v := Vector(5, 6, 7)
	if p.Subtract(v) != (Tuple{-2, -4, -6, 1}) {
		t.Errorf("wanted p-v = %v, got %v", Tuple{-2, -4, -6, 1}, p.Subtract(v))
	}
}

func TestSubtractTwoVectors(t *testing.T) {
	v1 := Vector(3, 2, 1)
	v2 := Vector(5, 6, 7)
	if (v1.Subtract(v2) != Tuple{-2, -4, -6, 0}) {
		t.Errorf("wanted p-v = %v, got %v", Tuple{-2, -4, -6, 0}, v1.Subtract(v2))
	}
}

func TestSubtractVectorFromZero(t *testing.T) {
	v1 := Vector(0, 0, 0)
	v2 := Vector(1, -2, 3)
	if (v1.Subtract(v2) != Tuple{-1, 2, -3, 0}) {
		t.Errorf("wanted v1-v2 = %v, got %v", Tuple{-1, 2, -3, 0}, v1.Subtract(v2))
	}
}

func TestNegate(t *testing.T) {
	tup := Tuple{1, -2, 3, -4}
	if tup.Negate() != (Tuple{-1, 2, -3, 4}) {
		t.Errorf("wanted -t = %v, got %v", Tuple{-1, 2, -3, 4}, tup.Negate())
	}
}

func TestMultiplication(t *testing.T) {
	tup := Tuple{1, -2, 3, -4}
	if tup.Multiply(3.5) != (Tuple{3.5, -7, 10.5, -14}) {
		t.Errorf("wanted t*3.5 = %v, got %v", Tuple{3.5, -7, 10.5, -14}, tup.Multiply(3.5))
	}
	if tup.Multiply(0.5) != (Tuple{0.5, -1, 1.5, -2}) {
		t.Errorf("wanted t*0.5 = %v, got %v", Tuple{0.5, -1, 1.5, -2}, tup.Multiply(0.5))
	}
}
func TestDivision(t *testing.T) {
	tup := Tuple{1, -2, 3, -4}
	if (tup.Divide(2) != Tuple{0.5, -1, 1.5, -2}) {
		t.Errorf("wanted t/2 = %v, got %v", Tuple{0.5, -1, 1.5, -2}, tup.Divide(2))
	}
}

func TestNormalize(t *testing.T) {
	tup := Vector(4, 0, 0)
	if tup.Normalize() != Vector(1, 0, 0) {
		t.Errorf("wanted normalized vector = %v, got %v", Vector(1, 0, 0), tup.Normalize())
	}

	tup = Vector(1, 2, 3)
	norm := tup.Normalize()
	if !norm.Equals(Vector(0.26726, 0.53452, 0.80178)) {
		t.Errorf("wanted normalized vector = %v, got %v", Vector(0.26726, 0.53452, 0.80178), tup.Normalize())
	}
}

func TestMagnitudeOfNormalizedVector(t *testing.T) {
	tup := Vector(1, 2, 3)
	n := tup.Normalize()
	if n.Magnitude() != 1 {
		t.Errorf("wanted magnitude = 1, got %v", n.Magnitude())
	}
}

func TestReflectVectorAt45Degrees(t *testing.T) {
	v := Vector(1, -1, 0)
	n := Vector(0, 1, 0)
	r := v.Reflect(n)
	if !r.Equals(Vector(1, 1, 0)) {
		t.Errorf("wanted r=%v, got %v", Vector(1, 1, 0), r)
	}
}

func TestReflectVectorSlantedSurface(t *testing.T) {
	v := Vector(0, -1, 0)
	n := Vector(math.Sqrt(2)/2, math.Sqrt(2)/2, 0)
	r := v.Reflect(n)
	if !r.Equals(Vector(1, 0, 0)) {
		t.Errorf("wanted r=%v, got %v", Vector(1, 1, 0), r)
	}
}
