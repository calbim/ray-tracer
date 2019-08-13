package canvas

import (
	"fmt"
	"strings"

	"github.com/calbim/ray-tracer/src/color"
)

// Canvas is a collection of pixels
type Canvas struct {
	width  int
	height int
	Pixels [][]color.Color
}

// New returns a new Canvas with width w and height h
func New(w, h int) Canvas {
	pixels := make([][]color.Color, h)
	for i := 0; i < h; i++ {
		pixels[i] = make([]color.Color, w)
	}
	return Canvas{
		w, h, pixels,
	}
}

// WritePixel writes a color at  width x and height y
func (c *Canvas) WritePixel(x int, y int, col color.Color) {
	c.Pixels[y][x] = col
}

// ToPPM converts a canvas to a PPM image type
func (c *Canvas) ToPPM() string {
	var b strings.Builder
	b.WriteString("P3\n")
	b.WriteString(fmt.Sprintf("%d %d\n", c.width, c.height))
	b.WriteString(fmt.Sprintf("%d\n", 255))
	length := 0
	for i := 0; i < c.height; i++ {
		for j := 0; j < c.width; j++ {
			p := c.Pixels[i][j]
			pix := p.ToPPMPixel()
			b.WriteString(pix)
			length = length + len(pix)
			if length > 56 {
				b.WriteString("\n")
				length = 0
			} else if j == c.width-1 {
				b.WriteString("\n")
				length = 0
			}
		}
	}
	return b.String()
}
