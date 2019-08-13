package tuple

import (
	"math"

	"github.com/calbim/ray-tracer/src/util"
)

// Tuple is set of coordinates
type Tuple struct {
	X float64
	Y float64
	Z float64
	W float64
}

// Point returns a new point
func Point(x, y, z float64) Tuple {
	return Tuple{x, y, z, 1}
}

// Vector returns a new vector
func Vector(x, y, z float64) Tuple {
	return Tuple{x, y, z, 0}
}

// Add returns t1 + t2
func (t *Tuple) Add(t2 Tuple) Tuple {
	return Tuple{t.X + t2.X, t.Y + t2.Y, t.Z + t2.Z, t.W + t2.W}
}

// Subtract returns t1 - t2
func (t *Tuple) Subtract(t2 Tuple) Tuple {
	return Tuple{t.X - t2.X, t.Y - t2.Y, t.Z - t2.Z, t.W - t2.W}
}

// isPoint reports whether t is a point
func (t *Tuple) isPoint() bool {
	return t.W == 1
}

// isVector reports whether t is a vector
func (t *Tuple) isVector() bool {
	return t.W == 0
}

// Negate returns the negation of a tuple
func (t *Tuple) Negate() Tuple {
	return Tuple{-t.X, -t.Y, -t.Z, -t.W}
}

// Multiply returns the product of a tuple and a number
func (t *Tuple) Multiply(n float64) Tuple {
	return Tuple{t.X * n, t.Y * n, t.Z * n, t.W * n}
}

// Divide returns the result of dividing a tuple by a number
func (t *Tuple) Divide(n float64) Tuple {
	return Tuple{t.X / n, t.Y / n, t.Z / n, t.W / n}
}

// Magnitude returns the magnitude of a vector
func (t *Tuple) Magnitude() float64 {
	return math.Sqrt(t.X*t.X + t.Y*t.Y + t.Z*t.Z + t.W*t.W)
}

// Normalize converts a vector to a unit vector
func (t *Tuple) Normalize() Tuple {
	return t.Divide(t.Magnitude())
}

// DotProduct returns the scalar product of two vectors
func (t *Tuple) DotProduct(t2 Tuple) float64 {
	return t.X*t2.X + t.Y*t2.Y + t.Z*t2.Z + t.W*t2.W
}

// CrossProduct returns the vector product of two vectors
func (t *Tuple) CrossProduct(t2 Tuple) Tuple {
	return Tuple{t.Y*t2.Z - t.Z*t2.Y, t.Z*t2.X - t2.Z*t.X, t.X*t2.Y - t2.X*t.Y, 0}
}

// Equals reports whether two tuples are the same
func (t *Tuple) Equals(t2 Tuple) bool {
	return util.Equals(t.X, t2.X) && util.Equals(t.Y, t2.Y) && util.Equals(t.Z, t2.Z) && util.Equals(t.W, t2.W)
}

//Reflect returns the vector after it reflects with respect to the given normal
func (t *Tuple) Reflect(normal Tuple) Tuple {
	return t.Subtract(normal.Multiply(2*(t.DotProduct(normal))))
}
