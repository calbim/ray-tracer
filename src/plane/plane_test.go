package plane

import (
	"fmt"
	"math"
	"os"
	"testing"

	"github.com/calbim/ray-tracer/src/camera"
	"github.com/calbim/ray-tracer/src/canvas"
	"github.com/calbim/ray-tracer/src/light"
	"github.com/calbim/ray-tracer/src/material"
	"github.com/calbim/ray-tracer/src/matrix"
	"github.com/calbim/ray-tracer/src/ray"
	"github.com/calbim/ray-tracer/src/shapes"
	"github.com/calbim/ray-tracer/src/sphere"
	"github.com/calbim/ray-tracer/src/transformations"
	"github.com/calbim/ray-tracer/src/tuple"
	"github.com/calbim/ray-tracer/src/world"
)

func TestNormal(t *testing.T) {
	p := Plane{}
	n1, err := p.Normal(tuple.Point(0, 0, 0))
	if err != nil {
		t.Errorf("Could not compute normal because %v", err)
	}
	n2, err := p.Normal(tuple.Point(10, 0, -10))
	if err != nil {
		t.Errorf("Could not compute normal because %v", err)
	}
	n3, err := p.Normal(tuple.Point(-5, 0, -150))
	if err != nil {
		t.Errorf("Could not compute normal because %v", err)
	}
	if *n1 != tuple.Vector(0, 1, 0) || *n2 != tuple.Vector(0, 1, 0) || *n3 != tuple.Vector(0, 1, 0) {
		t.Errorf("Plane normal should be (0,1,0)")
	}
}

func TestIntersectParallelRay(t *testing.T) {
	p := Plane{}
	r := ray.Ray{
		Origin:    tuple.Point(0, 10, 0),
		Direction: tuple.Vector(0, 0, 1),
	}
	p.SetSavedRay(r)
	xs, err := p.Intersect()
	if err != nil {
		t.Errorf("Error during calculating intersection points")
	}
	if len(xs) != 0 {
		t.Errorf("Ray should not intersect with plane because they are parallel")
	}
}

func TestIntersectCoplanarRay(t *testing.T) {
	p := Plane{}
	r := ray.Ray{
		Origin:    tuple.Point(0, 0, 0),
		Direction: tuple.Vector(0, 0, 1),
	}
	p.SetSavedRay(r)
	xs, err := p.Intersect()
	if err != nil {
		t.Errorf("Error during calculating intersection points")
	}
	if len(xs) != 0 {
		t.Errorf("Ray should not intersect with plane because they are parallel")
	}
}

func TestIntersectRayAbove(t *testing.T) {
	p := Plane{}
	r := ray.Ray{
		Origin:    tuple.Point(0, 1, 0),
		Direction: tuple.Vector(0, -1, 0),
	}
	p.SetSavedRay(r)
	xs, err := p.Intersect()
	if err != nil {
		t.Errorf("Error during calculating intersection points")
	}
	if len(xs) != 1 {
		t.Errorf("There should be one intersection point")
	}
	if xs[0].Value != 1 {
		t.Errorf("Intersection value should be 1")
	}
	if xs[0].Object != &p {
		t.Errorf("Intersected objet should be %v", p)
	}
}

func TestIntersectRayBelow(t *testing.T) {
	p := Plane{}
	r := ray.Ray{
		Origin:    tuple.Point(0, -1, 0),
		Direction: tuple.Vector(0, 01, 0),
	}
	p.SetSavedRay(r)
	xs, err := p.Intersect()
	if err != nil {
		t.Errorf("Error during calculating intersection points")
	}
	if len(xs) != 1 {
		t.Errorf("There should be one intersection point")
	}
	if xs[0].Value != 1 {
		t.Errorf("Intersection value should be 1")
	}
	if xs[0].Object != &p {
		t.Errorf("Intersected objet should be %v", p)
	}
}

func TestPlaneIsShape(t *testing.T) {
	var shape shapes.Shape
	shape = &Plane{}
	_, ok := shape.(shapes.Shape)
	if !ok {
		t.Errorf("Plane does not implement shape")
	}
}

func TestSceneWithPlanes(t *testing.T) {
	floor := New()
	floor.Material.Color = tuple.Color(1, 0.9, 0.9)
	floor.Material.Specular = 0

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
	transform := transformations.NewScaling(0.5, 0.5, 0.5)
	right.SetTransform(matrix.Multiply(transformations.NewTranslation(1.5, 0.5, -0.75), transform))
	right.SetMaterial(material.New())
	right.Material.Color = tuple.ColorFromHex("ff4785ff")
	right.Material.Diffuse = 1
	right.Material.Specular = 0.4

	w := world.World{}
	w.Objects = []shapes.Shape{floor, middle, right}
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
