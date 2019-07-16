package world

import (
	"github.com/calbim/ray-tracer/src/intersections"
	"github.com/calbim/ray-tracer/src/light"
)
//World contains a set of objects and a light source
type World struct {
	Objects *intersections.Object
	Light *light.PointLight
}