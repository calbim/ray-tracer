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

func TestPatternScene(t *testing.T) {
	floor := shape.NewPlane()
	m := material.New()
	p := pattern.NewRing(color.FromHex("ffff00ff"), color.Black)
	m.SetPattern(p)
	floor.SetMaterial(&m)

	sphere := shape.NewSphere()
	m2 := material.New()
	p2 := pattern.NewCheckers(color.FromHex("ff69b4ff"), color.White)
	p2.SetTransform(transforms.Scaling(0.3, 0.3, 0.3))
	m2.SetPattern(p2)

	sphere.SetMaterial(&m2)
	sphere.SetTransform(transforms.Translation(5, 2, 5))

	w := world.World{}
	l := light.PointLight(tuple.Point(0, 10, -5), color.New(1, 1, 1))

	w.Light = &l
	w.Objects = []shape.Shape{floor, sphere}
	camera := camera.New(1000, 500, math.Pi/3)
	camera.Transform = transforms.ViewTransform(tuple.Point(0, 5, -5), tuple.Point(10, 0, 10), tuple.Vector(0, 1, 0))
	can := camera.Render(w)
	ppm := can.ToPPM()
	file, err := os.Create("pattern.ppm")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintf(file, ppm)
}
