package tuple

// A Tuple is a set of coordinates (x,y,z) that represent a  point or direction in space.
// w = 1.0 for a point
// w = 0.0 for a vector
type Tuple struct {
	x float32
	y float32
	z float32
	w float32
}

// Point is a factory method that returns a point Tuple
func Point(x, y, z float32) Tuple {
	return Tuple{x, y, z, 1.0}
}

// Vector is a factory method that returns a point Vector
func Vector(x, y, z float32) Tuple {
	return Tuple{x, y, z, 0.0}
}

// Add adds two tuples
func Add(t1, t2 Tuple) Tuple {
	return Tuple{t1.x + t2.x, t1.y + t2.y, t1.z + t2.z, t1.w + t2.w}
}

// Subtract subtracts two tuples
func Subtract(t1, t2 Tuple) Tuple {
	return Tuple{t1.x - t2.x, t1.y - t2.y, t1.z - t2.z, t1.w - t2.w}
}
