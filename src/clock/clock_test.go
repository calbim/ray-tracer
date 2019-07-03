package clock

import (
	"fmt"
	"os"
	"testing"

	"github.com/calbim/ray-tracer/src/canvas"
	"github.com/calbim/ray-tracer/src/matrix"
	"github.com/calbim/ray-tracer/src/transformations"
	"github.com/calbim/ray-tracer/src/tuple"
)

func TestDrawClock(t *testing.T) {
	var pX, pY int
	c := canvas.New(500, 500)
	p := tuple.Point(100, 100, 0)
	o := tuple.Point(250, 250, 0)
	color := tuple.Point(1, 0, 1)
	rotation := transformations.RotationZ(transformations.Pi / 6)
	for i := 0; i < 12; i++ {
		pX = int(p.X + o.X)
		pY = int(p.Y + o.Y)
		fmt.Println(pX, pY)
		canvas.WritePixel(&c, pX, pY, color)
		p = matrix.MultiplyWithTuple(rotation, p)
	}
	ppm := canvas.ToPPM(c)
	file, err := os.Create("clock.ppm")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintf(file, ppm)
}
