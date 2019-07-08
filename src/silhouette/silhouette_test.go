package silhouette

import (
	"fmt"
	"os"
	"testing"

	"github.com/calbim/ray-tracer/src/canvas"
	"github.com/calbim/ray-tracer/src/intersections"
	"github.com/calbim/ray-tracer/src/ray"
	"github.com/calbim/ray-tracer/src/sphere"
	"github.com/calbim/ray-tracer/src/tuple"
)

func TestSilhouette(t *testing.T) {
	rayOrigin := tuple.Point(0, 0, -5)
	c := canvas.New(100, 100)
	color := tuple.Point(1, 0, 1)
	shape, err := sphere.New()
	if err != nil {
		t.Errorf("Error while creating sphere")
	}
	wallZ := 10.0
	wallSize := 7.0
	half := wallSize / 2
	pixelSize := wallSize / 100
	for y := 0; y < 100; y++ {
		worldY := float64(half - pixelSize*float64(y))
		for x := 0; x < 100; x++ {
			worldX := float64(-half + pixelSize*float64(x))
			position := tuple.Point(worldX, worldY, wallZ)
			r := ray.Ray{Origin: rayOrigin, Direction: tuple.Subtract(position, rayOrigin)}
			xs, err := sphere.Intersect(shape, r)
			if err != nil {
				t.Errorf("Error while calculating intersection")
			}
			if intersections.Hit(xs) != nil {
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
