package tuple

import (
	"testing"

	"../util"
)

func TestIsPoint(t *testing.T) {
	tup := Tuple{4.3, -4.2, 3.1, 1.0}
	if tup.X != 4.3 {
		t.Errorf("Expected tuple.X to be %f", 4.3)
	}
	if tup.Y != -4.2 {
		t.Errorf("Expected tuple.Y to be %f", -4.2)
	}
	if tup.Z != 3.1 {
		t.Errorf("Expected tuple.Z to be %f", 3.1)
	}
	if tup.W != 1.0 {
		t.Errorf("Expected tuple to be a point")
	}
}

func TestIsVector(t *testing.T) {
	tup := Tuple{4.3, -4.2, 3.1, 0.0}
	if tup.X != 4.3 {
		t.Errorf("tuple.X should be %f", 4.3)
	}
	if tup.Y != -4.2 {
		t.Errorf("tuple.Y should be %f", -4.2)
	}
	if tup.Z != 3.1 {
		t.Errorf("tuple.Z should be %f", 3.1)
	}
	if tup.W != 0.0 {
		t.Errorf("tuple should be a vector")
	}
}

func TestPoint(t *testing.T) {
	tup := Point(4, -4, 3)
	if tup.W != 1.0 {
		t.Errorf("tuple should be a point")
	}
}

func TestVector(t *testing.T) {
	tup := Vector(4, -4, 3)
	if tup.W != 0.0 {
		t.Errorf("tuple should be a vector")
	}
}

func TestAdd(t *testing.T) {
	t1 := Tuple{3, -2, 5, 1}
	t2 := Tuple{-2, 3, 1, 0}
	t3 := Add(t1, t2)
	if (t3 != Tuple{1, 1, 6, 1}) {
		t.Errorf("tuples should add to (1,1,6,1)")
	}
}
func TestSubtract(t *testing.T) {
	t1 := Point(3, 2, 1)
	t2 := Point(5, 6, 7)
	t3 := Subtract(t1, t2)
	if t3 != (Tuple{-2, -4, -6, 0}) {
		t.Errorf("tuples should subtract to give vector(-2,-4,-6)")
	}
}

func TestSubtractVectorFromPoint(t *testing.T) {
	p := Point(3, 2, 1)
	v := Vector(5, 6, 7)
	tup := Subtract(p, v)
	if tup != (Tuple{-2, -4, -6, 1}) {
		t.Errorf("should be point(-2,-4,-6)")
	}
}

func TestSubtractVectors(t *testing.T) {
	v1 := Vector(3, 2, 1)
	v2 := Vector(5, 6, 7)
	tup := Subtract(v1, v2)
	if (tup != Tuple{-2, -4, -6, 0}) {
		t.Errorf("should be vector(-2,-4,-6)")
	}
}

func TestSubtractVectorFromZero(t *testing.T) {
	v1 := Vector(0, 0, 0)
	v2 := Vector(1, -2, 3)
	tup := Subtract(v1, v2)
	if (tup != Tuple{-1, 2, -3, 0}) {
		t.Errorf("should be vector(-1,2,-3)")
	}
}

func TestNegate(t *testing.T) {
	tup := Tuple{1, -2, 3, -4}
	if (Negate(tup) != Tuple{-1, 2, -3, 4}) {
		t.Errorf("should be (-1,2,-3,-4)")
	}
}

func TestMultiplyByScalar(t *testing.T) {
	tup := Tuple{1, -2, 3, -4}
	if (MultiplyByScalar(tup, 3.5) != Tuple{3.5, -7, 10.5, -14}) {
		t.Errorf("should be (3.5,-7,10.5,-14)")
	}
	if (MultiplyByScalar(tup, 0.5) != Tuple{0.5, -1, 1.5, -2}) {
		t.Errorf("should be (0.5, -1, 1.5, -2)")
	}
}

func TestDivideByScalar(t *testing.T) {
	tup := Tuple{1, -2, 3, -4}
	if (DivideByScalar(tup, 2) != Tuple{0.5, -1, 1.5, -2}) {
		t.Errorf("should be (0.5, -1, 1.5, -2)")
	}
}
func TestMagnitude(t *testing.T) {
	tup := Vector(1, 0, 0)
	if Magnitude(tup) != 1 {
		t.Errorf("should be 1")
	}
	tup = Vector(0, 1, 0)
	if Magnitude(tup) != 1 {
		t.Errorf("should be 1")
	}
	tup = Vector(0, 0, 1)
	if Magnitude(tup) != 1 {
		t.Errorf("should be 1")
	}
	tup = Vector(1, 2, 3)
	if !util.Equals(Magnitude(tup), 3.741657) {
		t.Errorf("should be 3.74165")
	}
	tup = Vector(-1, -2, -3)
	if !util.Equals(Magnitude(tup), 3.741657) {
		t.Errorf("should be 3.74165")
	}
}

func TestNormalize(t *testing.T) {
	tup := Vector(4, 0, 0)
	if Normalize(tup) != Vector(1, 0, 0) {
		t.Errorf("normalized vector should be a unit vector")
	}

	tup = Vector(1, 2, 3)
	nt := Normalize(tup)
	if !Equals(nt, Vector(0.26726, 0.53452, 0.80178)) {
		t.Errorf("normalized vector should be a unit vector %v", nt)
	}
}

func TestMagnitudeOfNormalizedVector(t *testing.T) {
	n := Normalize(Vector(1, 2, 3))
	if !util.Equals(Magnitude(n), 1) {
		t.Errorf("Magnitude of a normalized vector is 1")
	}
}

func TestDotProduct(t *testing.T) {
	v1 := Vector(1, 2, 3)
	v2 := Vector(2, 3, 4)
	if DotProduct(v1, v2) != 20 {
		t.Errorf("Dot product of vector(1,2,3) and vector(2,3,4) should be 20")
	}
}

func TestCrossProduct(t *testing.T) {
	v1 := Vector(1, 2, 3)
	v2 := Vector(2, 3, 4)
	if !Equals(CrossProduct(v1, v2), Vector(-1, 2, -1)) || !Equals(CrossProduct(v2, v1), (Vector(1, -2, 1))) {
		t.Errorf("Cross product of v1 and v2 should vector(-1,2,-1)")
	}
}

func TestColor(t *testing.T) {
	c := Color(-0.5, 0.4, 1.7)
	if c.X != -0.5 || c.Y != 0.4 || c.Z != 1.7 {
		t.Errorf("Colors are assigned incorrectly")
	}
}

func TestAddColor(t *testing.T) {
	c1 := Color(0.9, 0.6, 0.75)
	c2 := Color(0.7, 0.1, 0.25)
	if !Equals(Add(c1, c2), Color(1.6, 0.7, 1.0)) {
		t.Errorf("Sum of colours should be color(1.6, 0.7, 1.0)")
	}
}

func TestSubtractColor(t *testing.T) {
	c1 := Color(0.9, 0.6, 0.75)
	c2 := Color(0.7, 0.1, 0.25)
	if !Equals(Subtract(c1, c2), Color(0.2, 0.5, 0.5)) {
		t.Errorf("Difference of colours should be color(0.2, 0.5, 0.5)")
	}
}
func TestMultiplyColor(t *testing.T) {
	c := Color(0.2, 0.3, 0.4)
	if !Equals(MultiplyByScalar(c, 2), Color(0.4, 0.6, 0.8)) {
		t.Errorf(" of colours should be color(0.4, 0.6, 0.8)")
	}
}

func TestHadamardProduct(t *testing.T) {
	c1 := Color(1, 0.2, 0.4)
	c2 := Color(0.9, 1, 0.1)
	if !Equals(HadamardProduct(c1, c2), Color(0.9, 0.2, 0.04)) {
		t.Errorf("Hadamard product of colours should be color(0.9, 0.2, 0.04)")
	}
}
