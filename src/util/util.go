package util

import "math"

// Equals checks if two floats are almost equal
func Equals(f1, f2 float64) bool {
	eps := 0.000001
	if math.Abs(f1-f2) < eps {
		return true
	}
	return false
}
