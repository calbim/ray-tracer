package tuple

import "math"

// A Tuple is a set of coordinates (x,y,z) that represent a  point or direction in space.
// w = 1.0 for a point
// w = 0.0 for a vector
type Tuple struct {
	x float64
	y float64
	z float64
	w float64
}

// Point is a factory method that returns a point Tuple
func Point(x, y, z float64) Tuple {
	return Tuple{x, y, z, 1.0}
}

// Vector is a factory method that returns a point Vector
func Vector(x, y, z float64) Tuple {
	return Tuple{x, y, z, 0.0}
}

// Add adds t1 and t2 and returns a tuple
func Add(t1, t2 Tuple) Tuple {
	return Tuple{t1.x + t2.x, t1.y + t2.y, t1.z + t2.z, t1.w + t2.w}
}

// Subtract subtracts t2 from t1 and returns a tuple
func Subtract(t1, t2 Tuple) Tuple {
	return Tuple{t1.x - t2.x, t1.y - t2.y, t1.z - t2.z, t1.w - t2.w}
}

// Negate returns the negative of a tuple
func Negate(tup Tuple) Tuple {
	return Tuple{-tup.x, -tup.y, -tup.z, -tup.w}
}

// MultiplyByScalar multiplies a tuple with a scalar and returns the result
func MultiplyByScalar(t Tuple, f float64) Tuple {
	return Tuple{t.x * f, t.y * f, t.z * f, t.w * f}
}

// DivideByScalar divides a tuple with a scalar and returns the result
func DivideByScalar(t Tuple, f float64) Tuple {
	return Tuple{t.x / f, t.y / f, t.z / f, t.w / f}
}

func Magnitude(t Tuple) float64 {
	return math.Sqrt(t.x*t.x + t.y*t.y + t.z*t.z + t.w*t.w)
}
