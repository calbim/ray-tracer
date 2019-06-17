package tuple

import (
	"testing"

	"../util"
)

func TestIsPoint(t *testing.T) {
	tup := Tuple{4.3, -4.2, 3.1, 1.0}
	if tup.x != 4.3 {
		t.Errorf("Expected tuple.X to be %f", 4.3)
	}
	if tup.y != -4.2 {
		t.Errorf("Expected tuple.Y to be %f", -4.2)
	}
	if tup.z != 3.1 {
		t.Errorf("Expected tuple.Z to be %f", 3.1)
	}
	if tup.w != 1.0 {
		t.Errorf("Expected tuple to be a point")
	}
}

func TestIsVector(t *testing.T) {
	tup := Tuple{4.3, -4.2, 3.1, 0.0}
	if tup.x != 4.3 {
		t.Errorf("tuple.X should be %f", 4.3)
	}
	if tup.y != -4.2 {
		t.Errorf("tuple.Y should be %f", -4.2)
	}
	if tup.z != 3.1 {
		t.Errorf("tuple.Z should be %f", 3.1)
	}
	if tup.w != 0.0 {
		t.Errorf("tuple should be a vector")
	}
}

func TestPoint(t *testing.T) {
	tup := Point(4, -4, 3)
	if tup.w != 1.0 {
		t.Errorf("tuple should be a point")
	}
}

func TestVector(t *testing.T) {
	tup := Vector(4, -4, 3)
	if tup.w != 0.0 {
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
	if !util.Equals(Magnitude(tup), 3.7416573) {
		t.Errorf("should be 3.7416573")
	}
	tup = Vector(-1, -2, -3)
	if !util.Equals(Magnitude(tup), 3.7416573) {
		t.Errorf("should be 3.7416573")
	}
}
