package transforms

import (
	"math"

	"github.com/calbim/ray-tracer/src/matrix"
	"github.com/calbim/ray-tracer/src/tuple"
)

//Translation returns a matrix representing a translation operation
func Translation(x, y, z float64) *matrix.Matrix {
	return matrix.New([]float64{1, 0, 0, x, 0, 1, 0, y, 0, 0, 1, z, 0, 0, 0, 1})
}

// Scaling returns a matrix representing a scaling operation
func Scaling(x, y, z float64) *matrix.Matrix {
	return matrix.New([]float64{x, 0, 0, 0, 0, y, 0, 0, 0, 0, z, 0, 0, 0, 0, 1})
}

// RotationX returns a matrix that represents a rotation by r radians around the X axis
func RotationX(r float64) *matrix.Matrix {
	return matrix.New([]float64{1, 0, 0, 0, 0, math.Cos(r), -math.Sin(r), 0, 0, math.Sin(r), math.Cos(r), 0, 0, 0, 0, 1})
}

// RotationY returns a matrix that represents a rotation by r radians around the Y axis
func RotationY(r float64) *matrix.Matrix {
	return matrix.New([]float64{math.Cos(r), 0, math.Sin(r), 0, 0, 1, 0, 0, -math.Sin(r), 0, math.Cos(r), 0, 0, 0, 0, 1})
}

// RotationZ returns a matrix that represents a rotation by r radians around the Z axis
func RotationZ(r float64) *matrix.Matrix {
	return matrix.New([]float64{math.Cos(r), -math.Sin(r), 0, 0, math.Sin(r), math.Cos(r), 0, 0, 0, 0, 1, 0, 0, 0, 0, 1})
}

//Shearing returns a matrix that represents a shearing (skew) operation
func Shearing(Xy, Xz, Yx, Yz, Zx, Zy float64) *matrix.Matrix {
	return matrix.New([]float64{1, Xy, Xz, 0, Yx, 1, Yz, 0, Zx, Zy, 1, 0, 0, 0, 0, 1})
}

//Chain chains several transformations into one
func Chain(transforms ...*matrix.Matrix) *matrix.Matrix {
	p := matrix.Identity
	for i := len(transforms) - 1; i >= 0; i-- {
		p = p.Multiply(transforms[i])
	}
	return p
}

//ViewTransform returns a matrix that represents viewing vector
func ViewTransform(from tuple.Tuple, to tuple.Tuple, up tuple.Tuple) *matrix.Matrix {
	forward := to.Subtract(from)
	forward = forward.Normalize()
	upNormalize := up.Normalize()
	left := forward.CrossProduct(upNormalize)
	trueUp := left.CrossProduct(forward)
	orientation := matrix.New([]float64{
		left.X, left.Y, left.Z, 0,
		trueUp.X, trueUp.Y, trueUp.Z, 0,
		-forward.X, -forward.Y, -forward.Z, 0,
		0, 0, 0, 1})
	return orientation.Multiply(Translation(-from.X, -from.Y, -from.Z))
}
