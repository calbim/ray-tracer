package matrix

import (
	"errors"

	"github.com/calbim/ray-tracer/src/tuple"
	"github.com/calbim/ray-tracer/src/util"
)

// New Returns a new matrix with r rows and c columns
// Input values are in array a
func New(a []float64, r, c int) [][]float64 {
	m := make([][]float64, r)
	for i := 0; i < r; i++ {
		m[i] = make([]float64, c)
	}
	k := 0
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			m[i][j] = a[k]
			k++
		}
	}
	return m
}

// NewIdentity returns a new 4X4 identity matrix
func NewIdentity() [][]float64 {
	m := make([][]float64, 4)
	for i := 0; i < 4; i++ {
		m[i] = make([]float64, 4)
	}
	m[0][0], m[1][1], m[2][2], m[3][3] = 1, 1, 1, 1
	return m
}

// Equals checks matrix equality for two matrices
func Equals(a, b [][]float64, arow, acol, brow, bcol int) bool {
	if arow != brow || acol != bcol {
		return false
	}
	for i := 0; i < arow; i++ {
		for j := 0; j < acol; j++ {
			if !util.Equals(a[i][j], b[i][j]) {
				return false
			}
		}
	}
	return true
}

// Multiply multiplies two 4x4 matrices and returns the result
func Multiply(m, n [][]float64) [][]float64 {
	res := make([][]float64, 4)
	for i := 0; i < 4; i++ {
		res[i] = make([]float64, 4)
	}
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			res[i][j] = m[i][0]*n[0][j] + m[i][1]*n[1][j] + m[i][2]*n[2][j] + m[i][3]*n[3][j]
		}
	}
	return res
}

// MultiplyWithTuple returns the product of a matrix and tuple
func MultiplyWithTuple(m [][]float64, t tuple.Tuple) tuple.Tuple {
	return tuple.Tuple{
		X: m[0][0]*t.X + m[0][1]*t.Y + m[0][2]*t.Z + m[0][3]*t.W,
		Y: m[1][0]*t.X + m[1][1]*t.Y + m[1][2]*t.Z + m[1][3]*t.W,
		Z: m[2][0]*t.X + m[2][1]*t.Y + m[2][2]*t.Z + m[2][3]*t.W,
		W: m[3][0]*t.X + m[3][1]*t.Y + m[3][2]*t.Z + m[3][3]*t.W,
	}
}

// Transpose computes the transpose of matrix
// by swapping the rows and columns
func Transpose(m [][]float64) [][]float64 {
	res := make([][]float64, 4, 4)
	for i := 0; i < 4; i++ {
		res[i] = make([]float64, 4)
	}
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			res[j][i] = m[i][j]
		}
	}
	return res
}

// Determinant returns the determinant of a NxN matrix
func Determinant(m [][]float64, N int) float64 {
	if N == 2 {
		return m[0][0]*m[1][1] - m[0][1]*m[1][0]
	}
	if N == 3 {
		return m[0][0]*Cofactor(m, 0, 0, 3) + m[0][1]*Cofactor(m, 0, 1, 3) + m[0][2]*Cofactor(m, 0, 2, 3)
	}
	if N == 4 {
		return m[0][0]*Cofactor(m, 0, 0, 4) + m[0][1]*Cofactor(m, 0, 1, 4) + m[0][2]*Cofactor(m, 0, 2, 4) + m[0][3]*Cofactor(m, 0, 3, 4)
	}
	return 0
}

// Submatrix returns a copy of a given matrix with r rows and
// c columns after deleting row rd and column cd for passed to it
// Submatrix of a NxN is a N-1xN-1 matrix
func Submatrix(m [][]float64, r, c, rd, cd int) [][]float64 {
	res := make([]float64, (r-1)*(c-1))
	k := 0
	for i := 0; i < r; {
		if i == rd {
			i++
		}
		for j := 0; i < r && j < c; {
			if j == cd {
				j++
			}
			if j < c {
				res[k] = m[i][j]
				k++
				j++
			}
		}
		i++
	}
	return New(res, r-1, c-1)
}

// Minor returns the minor computed at row i and column j
// of a NxN matrix
func Minor(m [][]float64, i, j, N int) float64 {
	return Determinant(Submatrix(m, N, N, i, j), N-1)
}

// Cofactor computes the cofactor of a NxN matrix at (i,j)
func Cofactor(m [][]float64, i, j, N int) float64 {
	minor := Minor(m, i, j, N)
	if (i+j)%2 == 0 {
		return minor
	}
	return -minor
}

// IsInvertible determines if a NxN is invertible
func IsInvertible(m [][]float64, N int) bool {
	return Determinant(m, N) != 0
}

// Inverse returns the inverse of a matrix
func Inverse(m [][]float64, N int) ([][]float64, error) {
	if !IsInvertible(m, N) {
		return nil, errors.New("matrix is not invertible")
	}
	b := New([]float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, N, N)
	det := Determinant(m, N)
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			c := Cofactor(m, i, j, N)
			b[j][i] = c / det
		}
	}
	return b, nil
}
