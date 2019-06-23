package projectile

import (
	"fmt"
	"os"
	"testing"

	"../canvas"
	"../tuple"
)

func TestProjectile(t *testing.T) {
	p := Projectile{
		tuple.Point(0, 1, 0),
		tuple.Normalize(tuple.Vector(1, 1, 0)),
	}
	e := Environment{
		tuple.Vector(0, -0.1, 0),
		tuple.Vector(-0.01, 0, 0),
	}
	np := tick(p, e)
	for np.point.Y > 0 {
		fmt.Printf("Point coordinates are %v\n", np.point)
		np = tick(np, e)
	}

}

func TestDrawProjectile(t *testing.T) {
	c := canvas.New(900, 550)
	p := Projectile{
		tuple.Point(0, 1, 0),
		tuple.MultiplyByScalar(tuple.Normalize(tuple.Vector(1, 1.8, 0)), 11.25),
	}
	e := Environment{
		tuple.Vector(0, -0.1, 0),
		tuple.Vector(-0.01, 0, 0),
	}
	np := tick(p, e)
	for np.point.Y > 0 {
		x := int(np.point.X)
		y := int(np.point.Y)
		if x <= 900 && y <= 550 {
			canvas.WritePixel(&c, x, y, tuple.Color(1, 0, 0))
		}
		np = tick(np, e)
	}
	ppm := canvas.ToPPM(c)
	file, err := os.Create("file.ppm")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintf(file, ppm)
}
