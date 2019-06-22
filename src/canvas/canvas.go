package canvas

import (
	"fmt"
	"strings"

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

// WritePixel writes pixel p at position x,y on a canvas c
func WritePixel(c *Canvas, x int, y int, p tuple.Tuple) {
	c.pixels[x][y] = p
}

// ToPPM returns the PPM string representation of a canvas
func ToPPM(c Canvas) string {
	var b strings.Builder
	b.WriteString("P3\n")
	b.WriteString(fmt.Sprintf("%d %d\n", c.width, c.height))
	b.WriteString(fmt.Sprintf("%d\n", 255))
	return b.String()
}
