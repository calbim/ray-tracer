package light

import (
	"github.com/calbim/ray-tracer/src/color"

	"github.com/calbim/ray-tracer/src/tuple"
)

//Light represents a light of given intensity at a position
type Light struct {
	Intensity color.Color
	Position  tuple.Tuple
}

//PointLight returns a light originating at point p and intensity i
func PointLight(p tuple.Tuple, i color.Color) Light {
	return Light{
		Intensity: i,
		Position:  p,
	}
}
