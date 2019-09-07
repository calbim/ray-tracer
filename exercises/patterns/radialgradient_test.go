package patterns

import (
	"fmt"
	"math"
	"os"
	"testing"

	"github.com/calbim/ray-tracer/src/camera"
	"github.com/calbim/ray-tracer/src/color"
	"github.com/calbim/ray-tracer/src/light"
	"github.com/calbim/ray-tracer/src/material"
	"github.com/calbim/ray-tracer/src/pattern"
	"github.com/calbim/ray-tracer/src/shape"
	"github.com/calbim/ray-tracer/src/transforms"
	"github.com/calbim/ray-tracer/src/tuple"
	"github.com/calbim/ray-tracer/src/world"
)

func TestRadialGradient(t *testing.T) {
	floor := shape.NewPlane()
	m := material.New()
	p := pattern.NewRadialGradient(color.FromHex("ff69b4ff"), color.Black)
	m.SetPattern(p)
	floor.SetMaterial(&m)

	w := world.World{}
	l := light.PointLight(tuple.Point(0, 10, -5), color.New(1, 1, 1))

	w.Light = &l
	w.Objects = []shape.Shape{floor}
	camera := camera.New(600, 300, math.Pi/2)
	camera.Transform = transforms.ViewTransform(tuple.Point(10, 3, 10), tuple.Point(10.5, 0, 10), tuple.Vector(0, 1, 0))
	can := camera.Render(w)
	ppm := can.ToPPM()
	file, err := os.Create("radialgradient.ppm")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintf(file, ppm)
}
