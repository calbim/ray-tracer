package util

import (
	"encoding/hex"
	"math"
)

// Eps is Epsilon
var Eps = 0.00001

// Equals checks if two floats are almost equal
func Equals(f1, f2 float64) bool {
	if math.Abs(f1-f2) < Eps {
		return true
	}
	return false
}

// HexToRGB converts a hex string to float rgb values
func HexToRGB(hx string) []float64 {
	num, _ := hex.DecodeString(hx)
	return []float64{float64(num[0]) / 255.0, float64(num[1]) / 255.0, float64(num[2]) / 255.0}
}
