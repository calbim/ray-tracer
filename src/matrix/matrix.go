package matrix

import (
	"math"

	"../tuple"
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

// NewIdentity returns a new identity matrix
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
			if a[i][j] != b[i][j] {
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

// DeterminantTwoByTwo returns the determinant of a 2X2 matrix
func DeterminantTwoByTwo(m [][]float64) float64 {
	return math.Abs(m[0][0]*m[1][1] - m[0][1]*m[1][0])
}
