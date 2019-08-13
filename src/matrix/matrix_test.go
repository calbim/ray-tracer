package matrix

import (
	"testing"

	"github.com/calbim/ray-tracer/src/tuple"
	"github.com/calbim/ray-tracer/src/util"
)

func TestMatrixConstruction(t *testing.T) {
	m := New([]float64{1, 2, 3, 4, 5.5, 6.5, 7.5, 8.5,
		9, 10, 11, 12, 13.5, 14.5, 15.5, 16.5})
	if m.At(0, 0) != 1 {
		t.Errorf("wanted m(0,0)=%v,got %v", 1, m.At(0, 0))
	}
	if m.At(0, 3) != 4 {
		t.Errorf("wanted m(0,3)=%v,got %v", 4, m.At(0, 3))
	}
	if m.At(1, 0) != 5.5 {
		t.Errorf("wanted m(1,0)=%v,got %v", 5.5, m.At(1, 0))
	}
	if m.At(1, 2) != 7.5 {
		t.Errorf("wanted m(1,2)=%v,got %v", 7.5, m.At(1, 2))
	}
	if m.At(2, 2) != 11 {
		t.Errorf("wanted m(2,2)=%v,got %v", 11, m.At(2, 2))
	}
	if m.At(3, 0) != 13.5 {
		t.Errorf("wanted m(3,0)=%v,got %v", 13.5, m.At(3, 0))
	}
	if m.At(3, 2) != 15.5 {
		t.Errorf("wanted (3,2)=%v,got %v", 15.5, m.At(3, 2))
	}
}

func TestTwoByTwoMatrix(t *testing.T) {
	m := New([]float64{-3, 5, 1, -2})
	if m.At(0, 0) != -3 {
		t.Errorf("m(0,0) is %v, should be -3", m.At(0, 0))
	}
	if m.At(0, 1) != 5 {
		t.Errorf("m(0,1) is %v, should be 5", m.At(0, 1))
	}
	if m.At(1, 0) != 1 {
		t.Errorf("m(1,0) is %v, should be 1", m.At(1, 0))
	}
	if m.At(1, 1) != -2 {
		t.Errorf("m(1,1) is %v, should be 7.5", m.At(1, 1))
	}
}

func TestThreeByThreeMatrix(t *testing.T) {
	m := New([]float64{-3, 5, 0, 1, -2, 7, 0, 1, 1})
	if m.At(0, 0) != -3 {
		t.Errorf("m(0,0) is %v, should be -3", m.At(0, 0))
	}
	if m.At(1, 1) != -2 {
		t.Errorf("m(1,1) is %v, should be -2", m.At(1, 1))
	}
	if m.At(2, 2) != 1 {
		t.Errorf("m(2,2) is %v, should be 1", m.At(2, 2))
	}
}

func TestIdentical(t *testing.T) {
	m := New([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9,
		8, 7, 6, 5, 4, 3, 2})
	n := New([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9,
		8, 7, 6, 5, 4, 3, 2})
	if !m.Equals(n) {
		t.Errorf("wanted m==n, got m!=n")
	}
}

func TestNonIdentical(t *testing.T) {
	m := New([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9,
		8, 7, 6, 5, 4, 3, 2})
	n := New([]float64{2, 3, 4, 5, 6, 7, 8, 9,
		8, 7, 6, 5, 4, 3, 2, 1})
	if m.Equals(n) {
		t.Errorf("wanted m!=n, got m=n")
	}
}

func TestMultiply(t *testing.T) {
	m := New([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9,
		8, 7, 6, 5, 4, 3, 2})
	n := New([]float64{-2, 1, 2, 3, 3, 2, 1, -1, 4,
		3, 6, 5, 1, 2, 7, 8})
	product := m.Multiply(n)
	expected := New([]float64{20, 22, 50, 48, 44,
		54, 114, 108, 40, 58, 110, 102, 16, 26, 46, 42})
	if !product.Equals(expected) {
		t.Errorf("wanted m*n=%v, got %v", expected, product)
	}
}

func TestMultiplyWithTuple(t *testing.T) {
	m := New([]float64{1, 2, 3, 4, 2, 4, 4, 2, 8,
		6, 4, 1, 0, 0, 0, 1})
	tup := tuple.Tuple{X: 1, Y: 2, Z: 3, W: 1}
	product := m.MultiplyTuple(tup)
	if !product.Equals(tuple.Tuple{X: 18, Y: 24, Z: 33, W: 1}) {
		t.Errorf("wanted m*tup=%v, got %v", product, tuple.Tuple{X: 18, Y: 24, Z: 33, W: 1})
	}

}

func TestMultiplyWithIdentity(t *testing.T) {
	m := New([]float64{0, 1, 2, 4, 1, 2, 4, 8,
		2, 4, 8, 16, 4, 8, 16, 32})
	expected := m.Multiply(Identity)
	if !m.Equals(expected) {
		t.Errorf("wanted m*Identity=m, got %v", m.Multiply(Identity))
	}
}
func TestTranspose(t *testing.T) {
	m := New([]float64{0, 9, 3, 0, 9, 8, 0, 8,
		1, 8, 5, 3, 0, 0, 5, 8})
	expected := New([]float64{0, 9, 1, 0, 9, 8, 8,
		0, 3, 0, 5, 5, 0, 8, 3, 8})
	transpose := m.Transpose()
	if !transpose.Equals(expected) {
		t.Errorf("wanted m.Transpose()=%v, got %v", expected, transpose)
	}
}

func TestTransposeIdentity(t *testing.T) {
	transpose := Identity.Transpose()
	if !transpose.Equals(Identity) {
		t.Errorf("wanted Identity.Transpose()=Identity")
	}
}

func TestDeterminant(t *testing.T) {
	m := New([]float64{1, 5, -3, 2})
	if m.Determinant() != 17 {
		t.Errorf("wanted determinant=%v, got %v", m.Determinant(), 17)
	}
}

func TestSubmatrix(t *testing.T) {
	m := New([]float64{1, 5, 0, -3, 2, 7, 0, 6, -3})
	n := m.Submatrix(0, 2)
	expected := New([]float64{-3, 2, 0, 6})
	if !n.Equals(expected) {
		t.Errorf("wanted submatrix=%v, got %v", expected, n)
	}

	m = New([]float64{-6, 1, 1, 6, -8, 5, 8, 6, 1, 0, 8, 2, -7, 1, -1, 1})
	n = m.Submatrix(2, 1)
	expected = New([]float64{-6, 1, 6, -8, 8, 6, -7, -1, 1})
	if !n.Equals(expected) {
		t.Errorf("wanted submatrix=%v, got %v", expected, n)
	}
}

func TestMinor(t *testing.T) {
	m := New([]float64{3, 5, 0, 2, -1, -7, 6, -1, 5})
	sub := m.Submatrix(1, 0)
	if sub.Determinant() != 25 {
		t.Errorf("wanted determinant=%v, got %v", 25, sub.Determinant())
	}
	if m.Minor(1, 0) != 25 {
		t.Errorf("wanted minor=%v, got %v", 25, m.Minor(1, 0))
	}
}

func TestCofactor(t *testing.T) {
	m := New([]float64{3, 5, 0, 2, -1, -7, 6, -1, 5})
	if m.Cofactor(0, 0) != -12 {
		t.Errorf("wanted cofactor at (0,0)=%v, got %v", -12, m.Cofactor(0, 0))
	}
	if m.Cofactor(1, 0) != -25 {
		t.Errorf("wanted cofactor at (1,0)=%v, got %v", -25, m.Cofactor(1, 0))
	}
}
func TestDeterminantThreeByThreeMatrix(t *testing.T) {
	m := New([]float64{1, 2, 6, -5, 8, -4, 2, 6, 4})
	if m.Cofactor(0, 0) != 56 || m.Cofactor(0, 1) != 12 || m.Cofactor(0, 2) != -46 {
		t.Errorf("cofactors are wrong")
	}
	if m.Determinant() != -196 {
		t.Errorf("wanted determinant=%v, got %v", -196, m.Determinant())
	}
}
func TestDeterminantFourByFourMatrix(t *testing.T) {
	m := New([]float64{-2, -8, 3, 5, -3, 1, 7, 3, 1, 2, -9, 6, -6, 7, 7, -9})
	if m.Cofactor(0, 0) != 690 || m.Cofactor(0, 1) != 447 || m.Cofactor(0, 2) != 210 || m.Cofactor(0, 3) != 51 {
		t.Errorf("cofactors are wrong")
	}
	if m.Determinant() != -4071 {
		t.Errorf("wanted determinant=%v, got %v", -4071, m.Determinant())
	}
}

func TestInvertibility(t *testing.T) {
	m := New([]float64{6, 4, 4, 4, 5, 5, 7, 6, 4, -9, 3, -7, 9, 1, 7, -6})
	if m.Determinant() != -2120 {
		t.Errorf("wanted determinant=%v, got %v", -2120, m.Determinant())
	}
	if !m.Invertible() {
		t.Errorf("wanted matrix to be invertible")
	}
	m = New([]float64{-4, 2, -2, 3, 9, 6, 2, 6, 0, -5, 1, -5, 0, 0, 0, 0})
	if m.Determinant() != 0 {
		t.Errorf("wanted determinant=%v, got %v", 0, m.Determinant())
	}
}

func TestInverse(t *testing.T) {
	m := New([]float64{-5, 2, 6, -8, 1, -5, 1, 8, 7, 7, -6, -7, 1, -3, 7, 4})
	b, err := m.Inverse()
	if err != nil {
		t.Error(err)
	}
	if m.Determinant() != 532 {
		t.Errorf("wanted determinant=%v, got %v", 532, m.Determinant())
	}
	if m.Cofactor(2, 3) != -160 || !util.Equals(b.At(3, 2), float64(-160)/532) || m.Cofactor(3, 2) != 105 || !util.Equals(b.At(2, 3), float64(105)/532) {
		t.Errorf("inverse is incorrect")
	}
	expected := New([]float64{0.21805, 0.45113, 0.24060, -0.04511, -0.80827, -1.45677, -0.44361, 0.52068, -0.07895, -0.22368, -0.05263, 0.19737, -0.52256, -0.81391, -0.30075, 0.30639})
	if !b.Equals(expected) {
		t.Errorf("wanted inverse=%v, got %v", expected, b)
	}

	m = New([]float64{8, -5, 9, 2, 7, 5, 6, 1, -6, 0, 9, 6, -3, 0, -9, -4})
	b, err = m.Inverse()
	if err != nil {
		t.Error(err)
	}
	expected = New([]float64{-0.15385, -0.15385, -0.28205, -0.53846, -0.07692, 0.12308, 0.02564, 0.03077, 0.35897, 0.35897, 0.43590, 0.92308, -0.69231, -0.69231, -0.76923, -1.92308})
	if !b.Equals(expected) {
		t.Errorf("wanted inverse=%v, got %v", expected, b)
	}

	m = New([]float64{9, 3, 0, 9, -5, -2, -6, -3, -4, 9, 6, 4, -7, 6, 6, 2})
	b, err = m.Inverse()
	if err != nil {
		t.Error(err)
	}
	expected = New([]float64{-0.04074, -0.07778, 0.14444, -0.22222, -0.07778, 0.03333, 0.36667, -0.33333, -0.02901, -0.14630, -0.10926, 0.12963, 0.17778, 0.06667, -0.26667, 0.33333})
	if !b.Equals(expected) {
		t.Errorf("wanted inverse=%v, got %v", expected, b)
	}
}
