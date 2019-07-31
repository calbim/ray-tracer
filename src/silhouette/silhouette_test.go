package silhouette

import (
	"fmt"
	"os"
	"testing"

	"github.com/calbim/ray-tracer/src/canvas"
	"github.com/calbim/ray-tracer/src/light"
	"github.com/calbim/ray-tracer/src/material"
	"github.com/calbim/ray-tracer/src/ray"
	"github.com/calbim/ray-tracer/src/shapes"
	"github.com/calbim/ray-tracer/src/sphere"
	"github.com/calbim/ray-tracer/src/tuple"
)

func TestSilhouette(t *testing.T) {
	rayOrigin := tuple.Point(0, 0, -5)
	c := canvas.New(200, 200)
	shape, err := sphere.New()
	shape.Material.Color = tuple.Color(0, 0.2, 1)
	light := light.PointLight{Intensity: tuple.Color(1, 1, 1), Position: tuple.Point(-10, 10, -10)}
	if err != nil {
		t.Errorf("Error while creating sphere")
	}
	wallZ := 20.0
	wallSize := 14.0
	half := wallSize / 2
	pixelSize := wallSize / 200
	for y := 0; y < 200; y++ {
		worldY := float64(half - pixelSize*float64(y))
		for x := 0; x < 200; x++ {
			worldX := float64(-half + pixelSize*float64(x))
			position := tuple.Point(worldX, worldY, wallZ)
			r := ray.Ray{Origin: rayOrigin, Direction: tuple.Normalize(tuple.Subtract(position, rayOrigin))}
			xs, err := shapes.Intersect(shape, r)
			if err != nil {
				t.Errorf("Error while calculating intersection")
			}
			hit := shapes.Hit(xs)
			if hit != nil {
				p := ray.Position(r, hit.Value)
				normalv, err := hit.Object.Normal(p)
				if err != nil {
					t.Errorf("Could not find normal at point %v on sphere", p)
				}
				eyev := tuple.Negate(r.Direction)
				color := material.Lighting(shape.Material, light, p, eyev, *normalv, false)
				canvas.WritePixel(&c, x, y, color)
			}
		}
	}
	ppm := canvas.ToPPM(c)
	file, err := os.Create("file.ppm")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintf(file, ppm)
}
