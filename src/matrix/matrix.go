package matrix

import (
	"errors"
	"math"

	"github.com/calbim/ray-tracer/src/tuple"
	"github.com/calbim/ray-tracer/src/util"
)

// Matrix represents a NxN matrix
type Matrix struct {
	N      int
	Values [][]float64
}

// Identity is a 4x4 identity matrix
var Identity = New([]float64{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1})

// New returns a new NxN matrix from an array of N*N values
func New(arr []float64) *Matrix {
	N := int(math.Sqrt(float64(len(arr))))
	m := make([][]float64, N)
	for i := 0; i < N; i++ {
		m[i] = make([]float64, N)
	}
	k := 0
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			m[i][j] = arr[k]
			k++
		}
	}
	return &Matrix{
		N:      N,
		Values: m,
	}
}

// At returns matrix value at (i,J)
func (m *Matrix) At(i, j int) float64 {
	return m.Values[i][j]
}

// Equals reports whether two matrices are equal
func (m *Matrix) Equals(n *Matrix) bool {
	if m.N != n.N {
		return false
	}
	for i := 0; i < m.N; i++ {
		for j := 0; j < n.N; j++ {
			if !util.Equals(m.At(i, j), n.At(i, j)) {
				return false
			}
		}
	}
	return true
}

// Multiply returns the product of two 4x4 matrices
func (m *Matrix) Multiply(n *Matrix) *Matrix {
	res := make([][]float64, 4)
	for i := 0; i < 4; i++ {
		res[i] = make([]float64, 4)
	}
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			res[i][j] = m.At(i, 0)*n.At(0, j) + m.At(i, 1)*n.At(1, j) + m.At(i, 2)*n.At(2, j) + m.At(i, 3)*n.At(3, j)
		}
	}
	return &Matrix{
		N:      4,
		Values: res,
	}
}

// MultiplyTuple returns the product of a matrix and tuple
func (m *Matrix) MultiplyTuple(t tuple.Tuple) tuple.Tuple {
	return tuple.Tuple{
		X: m.At(0, 0)*t.X + m.At(0, 1)*t.Y + m.At(0, 2)*t.Z + m.At(0, 3)*t.W,
		Y: m.At(1, 0)*t.X + m.At(1, 1)*t.Y + m.At(1, 2)*t.Z + m.At(1, 3)*t.W,
		Z: m.At(2, 0)*t.X + m.At(2, 1)*t.Y + m.At(2, 2)*t.Z + m.At(2, 3)*t.W,
		W: m.At(3, 0)*t.X + m.At(3, 1)*t.Y + m.At(3, 2)*t.Z + m.At(3, 3)*t.W,
	}
}

// Transpose returns the transpose of a matrix
func (m *Matrix) Transpose() Matrix {
	res := make([][]float64, 4, 4)
	for i := 0; i < 4; i++ {
		res[i] = make([]float64, 4)
	}
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			res[j][i] = m.At(i, j)
		}
	}
	return Matrix{
		N:      4,
		Values: res,
	}
}

// Determinant returns the determinant of a NxN matrix
func (m *Matrix) Determinant() float64 {
	if m.N == 2 {
		return m.At(0, 0)*m.At(1, 1) - m.At(0, 1)*m.At(1, 0)
	}
	if m.N == 3 {
		return m.At(0, 0)*m.Cofactor(0, 0) + m.At(0, 1)*m.Cofactor(0, 1) + m.At(0, 2)*m.Cofactor(0, 2)
	}
	if m.N == 4 {
		return m.At(0, 0)*m.Cofactor(0, 0) + m.At(0, 1)*m.Cofactor(0, 1) + m.At(0, 2)*m.Cofactor(0, 2) + m.At(0, 3)*m.Cofactor(0, 3)
	}
	return 0
}

// Submatrix returns a copy of m after deleting row r and column c from it
func (m *Matrix) Submatrix(r, c int) *Matrix {
	subN := (m.N - 1)
	res := make([]float64, subN*subN)
	k := 0
	for i := 0; i < m.N; {
		if i == r {
			i++
		}
		for j := 0; i < m.N && j < m.N; {
			if j == c {
				j++
			}
			if j < m.N {
				res[k] = m.At(i, j)
				k++
				j++
			}
		}
		i++
	}
	return New(res)
}

// Minor returns the minor of a matrix at row i and column j
func (m *Matrix) Minor(i, j int) float64 {
	sub := m.Submatrix(i, j)
	return sub.Determinant()
}

// Cofactor returns the cofactor of a matrix at row i and column j
func (m *Matrix) Cofactor(i, j int) float64 {
	minor := m.Minor(i, j)
	if (i+j)%2 == 0 {
		return minor
	}
	return -minor
}

//Invertible reports if a matrix is invertible
func (m *Matrix) Invertible() bool {
	return m.Determinant() != 0
}

// Inverse returns the inverse of a matrix
func (m *Matrix) Inverse() (*Matrix, error) {
	if !m.Invertible() {
		return nil, errors.New("matrix is not invertible")
	}
	b := []float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	det := m.Determinant()
	for i := 0; i < m.N; i++ {
		for j := 0; j < m.N; j++ {
			c := m.Cofactor(i, j)
			b[j*(m.N)+i] = c / det
		}
	}
	inv := New(b)
	return inv, nil
}
