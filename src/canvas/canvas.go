package canvas

import (
	"fmt"
	"math"
	"strconv"
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
	for i := 0; i < c.height; i++ {
		for j := 0; j < c.width; j++ {
			pix := ToPixel(c.pixels[i][j])
			b.WriteString(pix)
			if j == c.width-1 {
				b.WriteString("\n")
			} else {
				b.WriteString(" ")
			}
		}
	}
	s := b.String()
	return s
	// res := ""
	// line := 0
	// for i := 0; i < len(s); {
	// 	pix := string(s[i : i+3])
	// 	res = res + pix
	// 	line = line + 3
	// 	i = i + 3
	// 	if line >= 67 {
	// 		res = res + "\n"
	// 	}
	// 	wspace := string(s[i])
	// 	if wspace == " " {
	// 		res = res + wspace
	// 	}

	// }
	// return b.String()
}

// ToPixel converts a color tuple to an string of 3 integers
// each lying between 0 and 255
func ToPixel(color tuple.Tuple) string {
	r := ColorToString(color.X)
	b := ColorToString(color.Y)
	g := ColorToString(color.Z)

	return r + " " + b + " " + g
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
