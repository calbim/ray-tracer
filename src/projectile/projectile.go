package projectile

import (
	"github.com/calbim/ray-tracer/src/tuple"
)

//Projectile denotes the trajectory of an object in space
type Projectile struct {
	point    tuple.Tuple
	velocity tuple.Tuple
}

//Environment in which the projectile is travelling
type Environment struct {
	gravity tuple.Tuple
	wind    tuple.Tuple
}

func tick(p Projectile, e Environment) Projectile {
	position := tuple.Add(p.point, p.velocity)
	velocity := tuple.Add(p.velocity, tuple.Add(e.gravity, e.wind))
	return Projectile{point: position, velocity: velocity}
}
