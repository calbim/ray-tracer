package matrix

// New Returns a new matrix with r rows and c columns
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
