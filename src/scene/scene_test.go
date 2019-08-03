package scene

import (
	"fmt"
	"math"
	"os"
	"testing"

	"github.com/calbim/ray-tracer/src/matrix"
	"github.com/calbim/ray-tracer/src/shapes"

	"github.com/calbim/ray-tracer/src/light"
	"github.com/calbim/ray-tracer/src/world"

	"github.com/calbim/ray-tracer/src/camera"
	"github.com/calbim/ray-tracer/src/canvas"

	"github.com/calbim/ray-tracer/src/tuple"

	"github.com/calbim/ray-tracer/src/material"
	"github.com/calbim/ray-tracer/src/sphere"
	"github.com/calbim/ray-tracer/src/transformations"
)

func TestRenderScene(t *testing.T) {
	floor, err := sphere.New()
	if err != nil {
		t.Errorf("Error %v creating floor", err)
	}
	floor.SetTransform(transformations.NewScaling(10, 0.01, 10))
	floor.Material = material.New()
	floor.Material.Color = tuple.Color(1, 0.9, 0.9)
	floor.Material.Specular = 0

	leftWall, err := sphere.New()
	if err != nil {
		t.Errorf("Error %v creating left wall", err)
	}
	transform := transformations.NewScaling(10, 0.01, 10)
	transform = matrix.Multiply(transformations.RotationX(math.Pi/2), transform)
	transform = matrix.Multiply(transformations.RotationY(-math.Pi/4), transform)
	transform = matrix.Multiply(transformations.NewTranslation(0, 0, 5), transform)
	leftWall.SetTransform(transform)
	leftWall.Material = floor.Material

	rightWall, err := sphere.New()
	if err != nil {
		t.Errorf("Error %v creating right wall", err)
	}
	transform = transformations.NewScaling(10, 0.01, 10)
	transform = matrix.Multiply(transformations.RotationX(math.Pi/2), transform)
	transform = matrix.Multiply(transformations.RotationY(math.Pi/4), transform)
	transform = matrix.Multiply(transformations.NewTranslation(0, 0, 5), transform)
	rightWall.SetTransform(transform)
	rightWall.Material = floor.Material

	middle, err := sphere.New()
	if err != nil {
		t.Errorf("Error %v creating middle sphere", err)
	}
	middle.SetTransform(transformations.NewTranslation(-0.5, 1, 0.5))
	middle.SetMaterial(material.New())
	middle.Material.Color = tuple.ColorFromHex("f6abb6ff")
	middle.Material.Diffuse = 0.9
	middle.Material.Specular = 0.3

	right, err := sphere.New()
	if err != nil {
		t.Errorf("Error %v creating right sphere", err)
	}
	transform = transformations.NewScaling(0.5, 0.5, 0.5)
	right.SetTransform(matrix.Multiply(transformations.NewTranslation(1.5, 0.5, -0.75), transform))
	right.SetMaterial(material.New())
	right.Material.Color = tuple.ColorFromHex("ff4785ff")
	right.Material.Diffuse = 1
	right.Material.Specular = 0.4

	left, err := sphere.New()
	if err != nil {
		t.Errorf("Error %v creating left sphere", err)
	}
	transform = transformations.NewScaling(0.6, 0.6, 0.6)
	left.SetTransform(matrix.Multiply(transformations.NewTranslation(-1.5, 0.5, -0.75), transform))
	left.SetMaterial(material.New())
	left.Material.Color = tuple.ColorFromHex("a64263ff")
	// left.Material.Pattern = pattern.StripePattern(tuple.ColorFromHex("a64263ff"), tuple.White)
	left.Material.Diffuse = 0.85
	left.Material.Specular = 0.6

	w := world.World{}
	w.Objects = []shapes.Shape{floor, leftWall, rightWall, middle, right, left}
	if err != nil {
		t.Errorf("Could not create world due to error %v", err)
	}
	w.Light = &light.PointLight{Position: tuple.Point(-10, 10, -10), Intensity: tuple.Color(1, 1, 1)}

	camera := camera.New(200, 100, math.Pi/3)
	camera.Transform = transformations.ViewTransform(tuple.Point(4, 5, -5), tuple.Point(0, 1, 0), tuple.Vector(0, 1, 0))
	image, err := camera.Render(w)
	if err != nil {
		t.Errorf("Error while rendering camera to world %v", err)
	}
	ppm := canvas.ToPPM(*image)
	file, err := os.Create("file.ppm")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintf(file, ppm)
}
