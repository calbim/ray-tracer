package color

import (
	"math"
	"strconv"

	"github.com/calbim/ray-tracer/src/util"
)

// Color represents a color in RGB notation
type Color struct {
	R float64
	G float64
	B float64
}

//Black color
var Black = New(0, 0, 0)

//White color
var White = New(1, 1, 1)

// New returns an RGB color
func New(r, g, b float64) Color {
	return Color{R: r, G: g, B: b}
}

// Add returns c1 + c2
func (c *Color) Add(c2 Color) Color {
	return Color{R: c.R + c2.R, G: c.G + c2.G, B: c.B + c2.B}
}

// Subtract returns c1 - c2
func (c *Color) Subtract(c2 Color) Color {
	return Color{R: c.R - c2.R, G: c.G - c2.G, B: c.B - c2.B}
}

// Multiply returns the product of a color and a scalar
func (c *Color) Multiply(x float64) Color {
	return Color{R: c.R * x, G: c.G * x, B: c.B * x}
}

// MultiplyColor returns the product of two colors
func (c *Color) MultiplyColor(c2 Color) Color {
	return Color{R: c.R * c2.R, G: c.G * c2.G, B: c.B * c2.B}
}

// Equals reports whether two colors are the same
func (c *Color) Equals(c2 Color) bool {
	return util.Equals(c.R, c2.R) && util.Equals(c.G, c2.G) && util.Equals(c.B, c2.B)
}

// ToPPMPixel represent a color as a PPM pixel
func (c *Color) ToPPMPixel() string {
	r := ToInt(c.R)
	g := ToInt(c.G)
	b := ToInt(c.B)
	return r + " " + g + " " + b + " "
}

// ToInt converts a float color to a number lying between 0 and 255
func ToInt(c float64) string {
	i := math.Round(c * 255)
	if i > 255 {
		i = 255
	}
	if i < 0 {
		i = 0
	}
	return strconv.Itoa(int(i))
}
