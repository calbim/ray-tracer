package util

import "math"

var eps = 0.00001

// Equals checks if two floats are almost equal
func Equals(f1, f2 float64) bool {
	if math.Abs(f1-f2) < eps {
		return true
	}
	return false
}
