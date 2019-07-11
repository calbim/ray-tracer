package light

import (
	"github.com/calbim/ray-tracer/src/tuple"
)

//PointLight represents a light of given intensity at given position
type PointLight struct{
	Intensity tuple.Tuple
	Position tuple.Tuple
}