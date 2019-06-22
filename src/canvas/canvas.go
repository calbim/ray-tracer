package canvas

import (
	"../tuple"
)

// Canvas represents a collection of pixels
type Canvas struct {
	width  int
	height int
	pixels [][]tuple.Tuple
}

// New returns a new Canvas with width w and height h
func New(w, h int) Canvas {
	pixels := make([][]tuple.Tuple, w)
	for i := 0; i < w; i++ {
		pixels[i] = make([]tuple.Tuple, h)
	}
	return Canvas{
		w, h, pixels,
	}
}
