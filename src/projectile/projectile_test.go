package projectile

import (
	"fmt"
	"testing"

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
