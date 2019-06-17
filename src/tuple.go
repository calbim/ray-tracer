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

func Point(x, y, z float32) Tuple {
	return Tuple{x, y, z, 1.0}
}

func Vector(x, y, z float32) Tuple {
	return Tuple{x, y, z, 0.0}
}
