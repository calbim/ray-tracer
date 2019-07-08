package tuple

import (
	"math"

	"github.com/calbim/ray-tracer/src/util"
)

// A Tuple is a set of coordinates (x,y,z) that represent a  point or direction in space.
// w = 1.0 for a point
// w = 0.0 for a vector
type Tuple struct {
	X float64
	Y float64
	Z float64
	W float64
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
	return Tuple{t1.X + t2.X, t1.Y + t2.Y, t1.Z + t2.Z, t1.W + t2.W}
}

// Subtract subtracts t2 from t1 and returns a tuple
func Subtract(t1, t2 Tuple) Tuple {
	return Tuple{t1.X - t2.X, t1.Y - t2.Y, t1.Z - t2.Z, t1.W - t2.W}
}

// Negate returns the negative of a tuple
func Negate(tup Tuple) Tuple {
	return Tuple{-tup.X, -tup.Y, -tup.Z, -tup.W}
}

// MultiplyByScalar multiplies a tuple with a scalar and returns the result
func MultiplyByScalar(t Tuple, f float64) Tuple {
	return Tuple{t.X * f, t.Y * f, t.Z * f, t.W * f}
}

// DivideByScalar divides a tuple with a scalar and returns the result
func DivideByScalar(t Tuple, f float64) Tuple {
	return Tuple{t.X / f, t.Y / f, t.Z / f, t.W / f}
}

// Magnitude returns the magnitude of a vector
func Magnitude(t Tuple) float64 {
	return math.Sqrt(t.X*t.X + t.Y*t.Y + t.Z*t.Z + t.W*t.W)
}

// Normalize converts a vector to a unit vector while preserving the direction of the vector
func Normalize(t Tuple) Tuple {
	return DivideByScalar(t, Magnitude(t))
}

// DotProduct returns the scalar product of two vectors
func DotProduct(v1, v2 Tuple) float64 {
	return v1.X*v2.X + v1.Y*v2.Y + v1.Z*v2.Z + v1.W*v2.W
}

// CrossProduct returns the vector product of two vectors
func CrossProduct(v1, v2 Tuple) Tuple {
	return Tuple{v1.Y*v2.Z - v1.Z*v2.Y, v1.Z*v2.X - v2.Z*v1.X, v1.X*v2.Y - v2.X*v1.Y, 0}
}

// Color returns a tuple that denotes a color
// X = red; Y = blue; Z = green
func Color(r, b, g float64) Tuple {
	return Tuple{r, b, g, 0}
}

// HadamardProduct is what you do to blend two colors
func HadamardProduct(c1, c2 Tuple) Tuple {
	return Tuple{
		c1.X * c2.X, c1.Y * c2.Y, c1.Z * c2.Z, 0,
	}
}

//Equals checks for the equality of two tuples
func Equals(t1, t2 Tuple) bool {
	if util.Equals(t1.X, t2.X) && util.Equals(t1.Y, t2.Y) &&
		util.Equals(t1.Z, t2.Z) && util.Equals(t1.W, t2.W) {
		return true
	}
	return false
}

//Reflect returns the reflected vector corresponding to in around normal
func Reflect(in, normal Tuple) Tuple {
	return Subtract(in, MultiplyByScalar(normal, 2 * DotProduct(in, normal)))
}
