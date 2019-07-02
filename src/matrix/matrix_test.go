package matrix

import (
	"fmt"
	"testing"

	"github.com/calbim/ray-tracer/src/tuple"
	"github.com/calbim/ray-tracer/src/util"
)

func TestMatrixConstruction(t *testing.T) {
	m := New([]float64{1, 2, 3, 4, 5.5, 6.5, 7.5, 8.5,
		9, 10, 11, 12, 13.5, 14.5, 15.5, 16.5}, 4, 4)
	if m[0][0] != 1 {
		t.Errorf("m[0][0] is %f, should be 1", m[0][0])
	}
	if m[0][3] != 4 {
		t.Errorf("m[0][3] is %f, should be 4", m[0][3])
	}
	if m[1][0] != 5.5 {
		t.Errorf("m[1][0] is %f, should be 5.5", m[1][0])
	}
	if m[1][2] != 7.5 {
		t.Errorf("m[1][2] is %f, should be 7.5", m[1][2])
	}
	if m[2][2] != 11 {
		t.Errorf("m[2][2] is %f, should be 11", m[0][0])
	}
	if m[3][0] != 13.5 {
		t.Errorf("m[3][0] is %f, should be 13.5", m[0][0])
	}
	if m[3][2] != 15.5 {
		t.Errorf("m[3][2] is %f, should be 15.5", m[0][0])
	}
}

func TestTwoByTwoMatrix(t *testing.T) {
	m := New([]float64{-3, 5, 1, -2}, 2, 2)
	if m[0][0] != -3 {
		t.Errorf("m[0][0] is %f, should be -3", m[0][0])
	}
	if m[0][1] != 5 {
		t.Errorf("m[0][1] is %f, should be 5", m[0][1])
	}
	if m[1][0] != 1 {
		t.Errorf("m[1][0] is %f, should be 1", m[1][0])
	}
	if m[1][1] != -2 {
		t.Errorf("m[1][1] is %f, should be 7.5", m[1][1])
	}
}
func TestThreeByThreeMatrix(t *testing.T) {
	m := New([]float64{-3, 5, 0, 1, -2, 7, 0, 1, 1}, 3, 3)
	if m[0][0] != -3 {
		t.Errorf("m[0][0] is %f, should be -3", m[0][0])
	}
	if m[1][1] != -2 {
		t.Errorf("m[1][1] is %f, should be -2", m[1][1])
	}
	if m[2][2] != 1 {
		t.Errorf("m[2][2] is %f, should be 1", m[2][2])
	}
}

func TestEqualsForIdentical(t *testing.T) {
	m := New([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9,
		8, 7, 6, 5, 4, 3, 2}, 4, 4)
	n := New([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9,
		8, 7, 6, 5, 4, 3, 2}, 4, 4)
	res := Equals(m, n, 4, 4, 4, 4)
	if res == false {
		t.Errorf("Matrices are equal")
	}
}

func TestEqualsForNonIdentical(t *testing.T) {
	m := New([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9,
		8, 7, 6, 5, 4, 3, 2}, 4, 4)
	n := New([]float64{2, 3, 4, 5, 6, 7, 8, 9,
		8, 7, 6, 5, 4, 3, 2, 1}, 4, 4)
	res := Equals(m, n, 4, 4, 4, 4)
	if res == true {
		t.Errorf("Matrices are not equal")
	}
}

func TestMultiply(t *testing.T) {
	m := New([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9,
		8, 7, 6, 5, 4, 3, 2}, 4, 4)
	n := New([]float64{-2, 1, 2, 3, 3, 2, 1, -1, 4,
		3, 6, 5, 1, 2, 7, 8}, 4, 4)
	if !Equals(Multiply(m, n), New([]float64{20, 22, 50, 48, 44,
		54, 114, 108, 40, 58, 110, 102, 16, 26, 46, 42}, 4, 4), 4, 4, 4, 4) {
		t.Errorf("Multiplication is incorrect")
	}
}

func TestMultiplyWithTuple(t *testing.T) {
	m := New([]float64{1, 2, 3, 4, 2, 4, 4, 2, 8,
		6, 4, 1, 0, 0, 0, 1}, 4, 4)
	b := tuple.Tuple{X: 1, Y: 2, Z: 3, W: 1}
	if !tuple.Equals(MultiplyWithTuple(m, b), tuple.Tuple{X: 18, Y: 24, Z: 33, W: 1}) {
		t.Errorf("Multiplication with tuple is incorrect")
	}

}

func TestMultiplyWithIdentity(t *testing.T) {
	m := New([]float64{0, 1, 2, 4, 1, 2, 4, 8,
		2, 4, 8, 16, 4, 8, 16, 32}, 4, 4)
	expected := Multiply(m, NewIdentity())
	if !Equals(m, expected, 4, 4, 4, 4) {
		t.Errorf("Product of a matrix and identity matrix should be the original matrix")
	}
}

func TestTranspose(t *testing.T) {
	m := New([]float64{0, 9, 3, 0, 9, 8, 0, 8,
		1, 8, 5, 3, 0, 0, 5, 8}, 4, 4)
	expected := New([]float64{0, 9, 1, 0, 9, 8, 8,
		0, 3, 0, 5, 5, 0, 8, 3, 8}, 4, 4)
	if !Equals(Transpose(m), expected, 4, 4, 4, 4) {
		t.Errorf("Transpose of matrix is incorrect")
	}
}

func TestIdentityTranspose(t *testing.T) {
	m := NewIdentity()
	if !Equals(Transpose(m), m, 4, 4, 4, 4) {
		t.Errorf("Transpose of identity matrix is incorrect")
	}
}

func TestDeterminant(t *testing.T) {
	m := New([]float64{1, 5, -3, 2}, 2, 2)
	if Determinant(m, 2) != 17 {
		t.Errorf("Determinant of 2X2 matrix should be 17, is %f", Determinant(m, 2))
	}
}

func TestSubmatrix(t *testing.T) {
	m := New([]float64{1, 5, 0, -3, 2, 7, 0, 6, -3}, 3, 3)
	n := Submatrix(m, 3, 3, 0, 2)
	expected := New([]float64{-3, 2, 0, 6}, 2, 2)
	if !Equals(n, expected, 2, 2, 2, 2) {
		t.Errorf("Submatrix should be %v", expected)
	}

	m = New([]float64{-6, 1, 1, 6, -8, 5, 8, 6, 1, 0, 8, 2, -7, 1, -1, 1}, 4, 4)
	n = Submatrix(m, 4, 4, 2, 1)
	expected = New([]float64{-6, 1, 6, -8, 8, 6, -7, -1, 1}, 3, 3)
	if !Equals(n, expected, 3, 3, 3, 3) {
		t.Errorf("Submatrix should be %v but is %v", expected, n)
	}
}

func TestMinor(t *testing.T) {
	m := New([]float64{3, 5, 0, 2, -1, -7, 6, -1, 5}, 3, 3)
	b := Submatrix(m, 3, 3, 1, 0)
	if Determinant(b, 2) != 25 {
		t.Errorf("Determinant should be 25")
	}
	if Minor(m, 1, 0, 3) != 25 {
		t.Errorf("Minor should be 25")
	}
}

func TestCofactor(t *testing.T) {
	m := New([]float64{3, 5, 0, 2, -1, -7, 6, -1, 5}, 3, 3)
	if Cofactor(m, 0, 0, 3) != -12 {
		t.Errorf("Cofactor is %f should be -12", Cofactor(m, 0, 0, 3))
	}
	if Cofactor(m, 1, 0, 3) != -25 {
		t.Errorf("Cofactor should be -25")
	}
}

func TestDeterminantThreeByThreeMatrix(t *testing.T) {
	m := New([]float64{1, 2, 6, -5, 8, -4, 2, 6, 4}, 3, 3)
	if Cofactor(m, 0, 0, 3) != 56 || Cofactor(m, 0, 1, 3) != 12 || Cofactor(m, 0, 2, 3) != -46 {
		t.Errorf("Cofactors are wrong")
	}
	if Determinant(m, 3) != -196 {
		t.Errorf("Determinant is %f, should be %d", Determinant(m, 3), -196)
	}
}

func TestDeterminantFourByFourMatrix(t *testing.T) {
	m := New([]float64{-2, -8, 3, 5, -3, 1, 7, 3, 1, 2, -9, 6, -6, 7, 7, -9}, 4, 4)
	if Cofactor(m, 0, 0, 4) != 690 || Cofactor(m, 0, 1, 4) != 447 || Cofactor(m, 0, 2, 4) != 210 || Cofactor(m, 0, 3, 4) != 51 {
		t.Errorf("Cofactors are wrong")
	}
	if Determinant(m, 4) != -4071 {
		t.Errorf("Determinant is %f, should be %d", Determinant(m, 4), -4071)
	}
}

func TestInvertibility(t *testing.T) {
	m := New([]float64{6, 4, 4, 4, 5, 5, 7, 6, 4, -9, 3, -7, 9, 1, 7, -6}, 4, 4)
	if Determinant(m, 4) != -2120 || !IsInvertible(m, 4) {
		fmt.Printf("Matrix is invertible")
	}
	m = New([]float64{-4, 2, -2, 3, 9, 6, 2, 6, 0, -5, 1, -5, 0, 0, 0, 0}, 4, 4)
	if Determinant(m, 4) != 0 || IsInvertible(m, 4) {
		fmt.Printf("Matrix is not invertible")
	}
}

func TestInverse(t *testing.T) {
	m := New([]float64{-5, 2, 6, -8, 1, -5, 1, 8, 7, 7, -6, -7, 1, -3, 7, 4}, 4, 4)
	b, err := Inverse(m, 4)
	if err != nil {
		t.Errorf("error while inversing matrix")
	}
	if Determinant(m, 4) != 532 {
		t.Errorf("Determinant should be 532")
	}
	if Cofactor(m, 2, 3, 4) != -160 || !util.Equals(b[3][2], float64(-160)/532) || Cofactor(m, 3, 2, 4) != 105 || !util.Equals(b[2][3], float64(105)/532) {
		t.Errorf("Inverse in incorrect")
	}
	expected := New([]float64{0.21805, 0.45113, 0.24060, -0.04511, -0.80827, -1.45677, -0.44361, 0.52068, -0.07895, -0.22368, -0.05263, 0.19737, -0.52256, -0.81391, -0.30075, 0.30639}, 4, 4)
	if !Equals(b, expected, 4, 4, 4, 4) {
		t.Errorf("Matrix should be %v, is %v", expected, b)
	}

	m = New([]float64{8, -5, 9, 2, 7, 5, 6, 1, -6, 0, 9, 6, -3, 0, -9, -4}, 4, 4)
	b, err = Inverse(m, 4)
	if err != nil {
		t.Errorf("error while inversing matrix")
	}
	expected = New([]float64{-0.15385, -0.15385, -0.28205, -0.53846, -0.07692, 0.12308, 0.02564, 0.03077, 0.35897, 0.35897, 0.43590, 0.92308, -0.69231, -0.69231, -0.76923, -1.92308}, 4, 4)
	if !Equals(b, expected, 4, 4, 4, 4) {
		t.Errorf("Matrix should be %v, is %v", expected, b)
	}

	m = New([]float64{9, 3, 0, 9, -5, -2, -6, -3, -4, 9, 6, 4, -7, 6, 6, 2}, 4, 4)
	b, err = Inverse(m, 4)
	if err != nil {
		t.Errorf("error while inversing matrix")
	}
	expected = New([]float64{-0.04074, -0.07778, 0.14444, -0.22222, -0.07778, 0.03333, 0.36667, -0.33333, -0.02901, -0.14630, -0.10926, 0.12963, 0.17778, 0.06667, -0.26667, 0.33333}, 4, 4)
	if !Equals(b, expected, 4, 4, 4, 4) {
		t.Errorf("Matrix should be %v\n, is %v", expected, b)
	}
}

func TestMultipluMatrixWithInverse(t *testing.T) {
	m := New([]float64{3, -9, 7, 3, 3, -8, 2, -9, -4, 4, 4, 1, -6, 5, -1, 1}, 4, 4)
	b := New([]float64{8, 2, 2, 2, 3, -1, 7, 0, 7, 0, 5, 4, 6, -2, 0, 5}, 4, 4)
	c := Multiply(m, b)
	bInverse, err := Inverse(b, 4)
	if err != nil {
		t.Errorf("inverse of B is incorrect")
	}
	if !Equals(m, Multiply(c, bInverse), 4, 4, 4, 4) {
		t.Errorf("if A*B = C, then A should = C * inverse(B)")
	}
}
