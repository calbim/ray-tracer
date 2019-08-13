package transforms

import (
	"math"
	"testing"

	"github.com/calbim/ray-tracer/src/matrix"

	"github.com/calbim/ray-tracer/src/tuple"
)

func TestTranslation(t *testing.T) {
	transform := Translation(5, -3, 2)
	p := tuple.Point(-3, 4, 5)
	product := transform.MultiplyTuple(p)
	if !product.Equals(tuple.Point(2, 1, 7)) {
		t.Errorf("wanted product=%v, got%v", tuple.Point(2, 1, 7), product)
	}
}

func TestTranslationInverse(t *testing.T) {
	transform := Translation(5, -3, 2)
	inv, err := transform.Inverse()
	if err != nil {
		t.Errorf("could not inverse translate matrix")
	}
	p := tuple.Point(-3, 4, 5)
	product := inv.MultiplyTuple(p)
	if !product.Equals(tuple.Point(-8, 7, 3)) {
		t.Errorf("wanted %v, got %v", tuple.Point(-8, 7, 3), product)
	}
}

func TestTranslateVector(t *testing.T) {
	transform := Translation(5, -3, 2)
	v := tuple.Vector(-3, 4, 5)
	product := transform.MultiplyTuple(v)
	if !product.Equals(v) {
		t.Errorf("wanted product=%v, got %v", v, product)
	}
}

func TestScalePoint(t *testing.T) {
	transform := Scaling(2, 3, 4)
	p := tuple.Point(-4, 6, 8)
	product := transform.MultiplyTuple(p)
	if !product.Equals(tuple.Point(-8, 18, 32)) {
		t.Errorf("wanted product=%v, got %v", tuple.Point(-8, 18, 32), product)
	}
}

func TestScaleVector(t *testing.T) {
	transform := Scaling(2, 3, 4)
	v := tuple.Vector(-4, 6, 8)
	product := transform.MultiplyTuple(v)
	if !product.Equals(tuple.Vector(-8, 18, 32)) {
		t.Errorf("wanted product=%v, got %v", tuple.Vector(-8, 18, 32), product)
	}
}

func TestScalingInverse(t *testing.T) {
	transform := Scaling(2, 3, 4)
	inv, err := transform.Inverse()
	if err != nil {
		t.Errorf("could not inverse scaling matrix")
	}
	v := tuple.Vector(-4, 6, 8)
	product := inv.MultiplyTuple(v)
	if !product.Equals(tuple.Vector(-2, 2, 2)) {
		t.Errorf("wanted product=%v, got %v", tuple.Vector(-2, 2, 2), product)
	}
}

func TestReflection(t *testing.T) {
	transform := Scaling(-1, 1, 1)
	p := tuple.Point(2, 3, 4)
	product := transform.MultiplyTuple(p)
	if !product.Equals(tuple.Point(-2, 3, 4)) {
		t.Errorf("wanted product=%v,got %v", product, tuple.Point(-2, 3, 4))
	}
}

func TestRotationX(t *testing.T) {
	p := tuple.Point(0, 1, 0)
	halfQuarter := RotationX(math.Pi / 4)
	fullQuarter := RotationX(math.Pi / 2)
	hR := halfQuarter.MultiplyTuple(p)
	if !hR.Equals(tuple.Point(0, math.Sqrt(2)/2, math.Sqrt(2)/2)) {
		t.Errorf("wanted rotated point=%v, got %v", tuple.Point(0, math.Sqrt(2)/2, math.Sqrt(2)/2), hR)
	}
	fR := fullQuarter.MultiplyTuple(p)
	if !fR.Equals(tuple.Point(0, 0, 1)) {
		t.Errorf("wanted rotated point=%v, got %v", tuple.Point(0, 0, 1), fR)
	}
}

func TestInverseRotationX(t *testing.T) {
	p := tuple.Point(0, 1, 0)
	halfQuarter := RotationX(math.Pi / 4)
	inv, err := halfQuarter.Inverse()
	if err != nil {
		t.Errorf("could not compute inverse")
	}
	hR := inv.MultiplyTuple(p)
	if !hR.Equals(tuple.Point(0, math.Sqrt(2)/2, -math.Sqrt(2)/2)) {
		t.Errorf("wanted rotated point=%v, got %v", tuple.Point(0, math.Sqrt(2)/2, -math.Sqrt(2)/2), hR)
	}
}

func TestYRotation(t *testing.T) {
	p := tuple.Point(0, 0, 1)
	halfQuarter := RotationY(math.Pi / 4)
	fullQuarter := RotationY(math.Pi / 2)
	hR := halfQuarter.MultiplyTuple(p)
	if !hR.Equals(tuple.Point(math.Sqrt(2)/2, 0, math.Sqrt(2)/2)) {
		t.Errorf("wanted rotated point=%v, got %v", tuple.Point(0, math.Sqrt(2)/2, math.Sqrt(2)/2), hR)
	}
	fR := fullQuarter.MultiplyTuple(p)
	if !fR.Equals(tuple.Point(1, 0, 0)) {
		t.Errorf("wanted rotated point=%v, got %v", tuple.Point(1, 0, 0), fR)
	}
}

func TestZRotation(t *testing.T) {
	p := tuple.Point(0, 1, 0)
	halfQuarter := RotationZ(math.Pi / 4)
	fullQuarter := RotationZ(math.Pi / 2)
	hR := halfQuarter.MultiplyTuple(p)
	if !hR.Equals(tuple.Point(-math.Sqrt(2)/2, math.Sqrt(2)/2, 0)) {
		t.Errorf("wanted rotated point=%v, got %v", tuple.Point(-math.Sqrt(2)/2, math.Sqrt(2)/2, 0), hR)
	}
	fR := fullQuarter.MultiplyTuple(p)
	if !fR.Equals(tuple.Point(-1, 0, 0)) {
		t.Errorf("wanted rotated point=%v, got %v", tuple.Point(-1, 0, 0), fR)
	}
}

func TestShearing(t *testing.T) {
	transform := Shearing(1, 0, 0, 0, 0, 0)
	p := transform.MultiplyTuple(tuple.Point(2, 3, 4))
	if !p.Equals(tuple.Point(5, 3, 4)) {
		t.Errorf("wanted shearing to be %v, got %v", tuple.Point(5, 3, 4), p)
	}
	transform = Shearing(0, 1, 0, 0, 0, 0)
	p = transform.MultiplyTuple(tuple.Point(2, 3, 4))
	if !p.Equals(tuple.Point(6, 3, 4)) {
		t.Errorf("wanted shearing to be %v, got %v", tuple.Point(6, 3, 4), p)
	}
	transform = Shearing(0, 0, 1, 0, 0, 0)
	p = transform.MultiplyTuple(tuple.Point(2, 3, 4))
	if !p.Equals(tuple.Point(2, 5, 4)) {
		t.Errorf("wanted shearing to be %v, got %v", tuple.Point(2, 5, 4), p)
	}
	transform = Shearing(0, 0, 0, 1, 0, 0)
	p = transform.MultiplyTuple(tuple.Point(2, 3, 4))
	if !p.Equals(tuple.Point(2, 7, 4)) {
		t.Errorf("wanted shearing to be %v, got %v", tuple.Point(2, 7, 4), p)
	}
	transform = Shearing(0, 0, 0, 0, 1, 0)
	p = transform.MultiplyTuple(tuple.Point(2, 3, 4))
	if !p.Equals(tuple.Point(2, 3, 6)) {
		t.Errorf("wanted shearing to be %v, got %v", tuple.Point(2, 3, 6), p)
	}
	transform = Shearing(0, 0, 0, 0, 0, 1)
	p = transform.MultiplyTuple(tuple.Point(2, 3, 4))
	if !p.Equals(tuple.Point(2, 3, 7)) {
		t.Errorf("wanted shearing to be %v, got %v", tuple.Point(2, 3, 7), p)
	}
}

func TestMultipleTransformations(t *testing.T) {
	p := tuple.Point(1, 0, 1)
	A := RotationX(math.Pi / 2)
	B := Scaling(5, 5, 5)
	C := Translation(10, 5, 7)
	p2 := A.MultiplyTuple(p)
	if !p2.Equals(tuple.Point(1, -1, 0)) {
		t.Errorf("wanted rotated p=%v, got %v", tuple.Point(1, -1, 0), p2)
	}
	p3 := B.MultiplyTuple(p2)
	if !p3.Equals(tuple.Point(5, -5, 0)) {
		t.Errorf("wanted scaled p=%v, got %v", tuple.Point(5, -5, 0), p3)
	}
	p4 := C.MultiplyTuple(p3)
	if !p4.Equals(tuple.Point(15, 0, 7)) {
		t.Errorf("wanted translated p=%v, got %v", tuple.Point(15, 0, 7), p4)
	}
}

func TestChainingTransformationsOrder(t *testing.T) {
	p := tuple.Point(1, 0, 1)
	A := RotationX(math.Pi / 2)
	B := Scaling(5, 5, 5)
	C := Translation(10, 5, 7)
	chain := Chain(A, B, C)
	product := chain.MultiplyTuple(p)
	if !product.Equals(tuple.Point(15, 0, 7)) {
		t.Errorf("transformations chained in the wrong order")
	}
}

func TestViewTransformDefaultOrientation(t *testing.T) {
	from := tuple.Point(0, 0, 0)
	to := tuple.Point(0, 0, -1)
	up := tuple.Vector(0, 1, 0)
	transform := ViewTransform(from, to, up)
	if !transform.Equals(matrix.Identity) {
		t.Errorf("wanted default transform=%v, got %v", transform, matrix.Identity)
	}
}

func TestTransformationMatrixforPositiveZAxis(t *testing.T) {
	from := tuple.Point(0, 0, 0)
	to := tuple.Point(0, 0, 1)
	up := tuple.Vector(0, 1, 0)
	transform := ViewTransform(from, to, up)
	if !transform.Equals(Scaling(-1, 1, -1)) {
		t.Errorf("wanted transform=%v, got %v", Scaling(-1, 1, -1), transform)
	}
}

func TestTransformationMovestheWorld(t *testing.T) {
	from := tuple.Point(0, 0, 8)
	to := tuple.Point(0, 0, 0)
	up := tuple.Vector(0, 1, 0)
	transform := ViewTransform(from, to, up)
	if !transform.Equals(Translation(0, 0, -8)) {
		t.Errorf("wanted transform=%v, got %v", Translation(0, 0, -8), transform)
	}
}

func TestArbitraryTransformation(t *testing.T) {
	from := tuple.Point(1, 3, 2)
	to := tuple.Point(4, -2, 8)
	up := tuple.Vector(1, 1, 0)
	transform := ViewTransform(from, to, up)
	expected := matrix.New([]float64{
		-0.50709, 0.50709, 0.67612, -2.36643,
		0.76772, 0.60609, 0.12122, -2.82843,
		-0.35857, 0.59761, -0.71714, 0.00000,
		0, 0, 0, 1})
	if !transform.Equals(expected) {
		t.Errorf("wanted transform=%v,got %v", Translation(0, 0, -8), expected)
	}
}
