package tuple

import (
	"testing"
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
	if t3.x != 1 || t3.y != 1 || t3.z != 6 {
		t.Errorf("tuples should add to (1,1,6,1)")
	}
}
func TestSubtract(t *testing.T) {
	t1 := Point(3, 2, 1)
	t2 := Point(5, 6, 7)
	t3 := Subtract(t1, t2)
	if t3.x != -2 || t3.y != -4 || t3.z != -6 || t3.w != 0 {
		t.Errorf("tuples should subtract to give a vector (-2,-4,-6)")
	}
}
