package hexagon

import (
	"fmt"
	"math"
	"os"
	"testing"

	"github.com/calbim/ray-tracer/src/camera"
	"github.com/calbim/ray-tracer/src/color"
	"github.com/calbim/ray-tracer/src/light"
	"github.com/calbim/ray-tracer/src/transforms"
	"github.com/calbim/ray-tracer/src/tuple"
	"github.com/calbim/ray-tracer/src/world"

	"github.com/calbim/ray-tracer/src/shape"
)

//todo: finish hexagon
func TestHexagon(t *testing.T) {
	p1 := shape.NewPlane()
	t1 := transforms.Chain(transforms.Translation(0, 0, -20), transforms.Rotation(math.Pi/2))
	p1.SetTransform(t1)
	p2 := shape.NewPlane()
	t2 := transforms.Chain(transforms.Translation(2, 0, 2), transforms.RotationX(math.Pi/2))
	p2.SetTransform(t2)

	w := world.World{}
	l := light.PointLight(tuple.Point(-10, 10, -10), color.New(1, 1, 1))
	w.Light = &l
	w.Objects = []shape.Shape{p1}
	camera := camera.New(100, 50, math.Pi/2)
	camera.Transform = transforms.ViewTransform(tuple.Point(0, 100, -0.5), tuple.Point(0, 1, 0), tuple.Vector(0, 1, 0))
	can := camera.Render(w)
	ppm := can.ToPPM()
	file, err := os.Create("hex.ppm")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintf(file, ppm)
}
