package canvas

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/calbim/ray-tracer/src/tuple"
)

// Canvas represents a collection of pixels
type Canvas struct {
	width  int
	height int
	pixels [][]tuple.Tuple
}

// New returns a new Canvas with width w and height h
func New(w, h int) Canvas {
	pixels := make([][]tuple.Tuple, h)
	for i := 0; i < h; i++ {
		pixels[i] = make([]tuple.Tuple, w)
	}
	return Canvas{
		w, h, pixels,
	}
}

// WritePixel writes pixel p at position width x, height y on a canvas c
func WritePixel(c *Canvas, x int, y int, p tuple.Tuple) {
	c.pixels[y][x] = p
}

// ToPPM returns the PPM string representation of a canvas
func ToPPM(c Canvas) string {
	var b strings.Builder
	b.WriteString("P3\n")
	b.WriteString(fmt.Sprintf("%d %d\n", c.width, c.height))
	b.WriteString(fmt.Sprintf("%d\n", 255))
	length := 0
	for i := 0; i < c.height; i++ {
		for j := 0; j < c.width; j++ {
			pix := ToPixel(c.pixels[i][j])
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

//SplitPPM splits long rows in PPM files when they exceed 70 chars
func SplitPPM(s string) string {
	slice := strings.Fields(s)
	res := ""
	limit := 17
	for len(slice) > 0 {
		res = res + strings.Join(slice[:limit], " ") + "\n"
		slice = slice[limit:]
		if len(slice) < limit {
			limit = len(slice)
		}
	}
	return res
}

// ToPixel converts a color tuple to an string of 3 integers
// each lying between 0 and 255
func ToPixel(color tuple.Tuple) string {
	r := ColorToString(color.X)
	b := ColorToString(color.Y)
	g := ColorToString(color.Z)

	return r + " " + b + " " + g + " "
}

// ColorToString converts a float64 representing a hue (r, b or g)
// to a number lying between 0 and 255
func ColorToString(c float64) string {
	i := math.Round(c * 255)
	if i > 255 {
		i = 255
	}
	if i < 0 {
		i = 0
	}
	return strconv.Itoa(int(i))
}
