package scene

import (
	"fmt"
	"math"
	"os"
	"testing"

	"github.com/calbim/ray-tracer/src/pattern"

	"github.com/calbim/ray-tracer/src/camera"

	"github.com/calbim/ray-tracer/src/light"
	"github.com/calbim/ray-tracer/src/tuple"

	"github.com/calbim/ray-tracer/src/color"
	"github.com/calbim/ray-tracer/src/material"
	"github.com/calbim/ray-tracer/src/shape"
	"github.com/calbim/ray-tracer/src/transforms"
	"github.com/calbim/ray-tracer/src/world"
)

func TestScene(t *testing.T) {
	floor := shape.NewSphere()
	floor.Transform = transforms.Scaling(10, 0.01, 10)
	mFloor := material.New()
	floor.Material = &mFloor
	floor.Material.Color = color.New(1, 0.9, 0.9)
	floor.Material.Specular = 0

	leftWall := shape.NewSphere()
	leftWall.Transform = transforms.Chain(transforms.Scaling(10, 0.01, 10), transforms.RotationX(math.Pi/2),
		transforms.RotationY(-math.Pi/4), transforms.Translation(0, 0, 5))
	leftWall.Material = floor.Material
	leftWall.Material.Color = color.FromHex("ffffffff")

	rightWall := shape.NewSphere()
	rightWall.Transform = transforms.Chain(transforms.Scaling(10, 0.01, 10),
		transforms.RotationX(math.Pi/2), transforms.RotationY(math.Pi/4), transforms.Translation(0, 0, 5))
	rightWall.Material = floor.Material

	left := shape.NewSphere()
	left.Transform = transforms.Chain(transforms.Scaling(0.33, 0.33, 0.33), transforms.Translation(-1.5, 0.5, -0.75))
	lMaterial := material.New()
	left.Material = &lMaterial
	left.Material.Color = color.FromHex("a64263ff")
	left.Material.Diffuse = 0.85
	left.Material.Specular = 0.6

	middle := shape.NewSphere()
	middle.SetTransform(transforms.Translation(-0.5, 1, 0.5))
	mMaterial := material.New()
	middle.Material = &mMaterial
	middle.Material.SetPattern(pattern.NewStripe(color.FromHex("f6abb6ff"), color.White))
	middle.Material.Color = color.FromHex("f6abb6ff")
	middle.Material.Diffuse = 0.9
	middle.Material.Specular = 0.3

	right := shape.NewSphere()
	right.Transform = transforms.Chain(transforms.Scaling(0.5, 0.5, 0.5), transforms.Translation(1.5, 0.5, -0.75))
	rMaterial := material.New()
	right.Material = &rMaterial
	right.Material.Color = color.FromHex("ff4785ff")
	middle.Material.SetPattern(pattern.NewStripe(color.FromHex("ff4785ff"), color.White))

	right.Material.Diffuse = 0.7
	right.Material.Specular = 0.4

	w := world.World{}
	l := light.PointLight(tuple.Point(-10, 10, -10), color.New(1, 1, 1))
	w.Light = &l
	w.Objects = []shape.Shape{leftWall, floor, left, right, middle, rightWall}
	camera := camera.New(300, 150, math.Pi/3)
	camera.Transform = transforms.ViewTransform(tuple.Point(0, 1.5, -5), tuple.Point(0, 1, 0), tuple.Vector(0, 1, 0))
	can := camera.Render(w)
	ppm := can.ToPPM()
	file, err := os.Create("scene-shadow.ppm")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintf(file, ppm)
}
