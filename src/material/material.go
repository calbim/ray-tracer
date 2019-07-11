package material

import (
	"math"

	"github.com/calbim/ray-tracer/src/light"
	"github.com/calbim/ray-tracer/src/tuple"
)

//Material represents a surface and components of the Phong reflection model
type Material struct {
	Color     tuple.Tuple
	Ambient   float64
	Diffuse   float64
	Specular  float64
	Shininess float64
}

//New returns a default material
func New() Material {
	return Material{
		Color:     tuple.Color(1, 1, 1),
		Ambient:   0.1,
		Diffuse:   0.9,
		Specular:  0.9,
		Shininess: 200,
	}
}

//Lighting returns the shade an observer sees
func Lighting(material Material, light light.PointLight, point tuple.Tuple, eyev tuple.Tuple, normalv tuple.Tuple) tuple.Tuple {
	effectiveColor := tuple.HadamardProduct(light.Intensity, material.Color)
	lightv := tuple.Subtract(light.Position, point)
	ambient := tuple.MultiplyByScalar(effectiveColor, material.Ambient)
	lightDotNormal := tuple.DotProduct(lightv, normalv)
	diffuse, specular := tuple.Color(0, 0, 0), tuple.Color(0, 0, 0)
	if lightDotNormal >= 0 {
		diffuse = tuple.MultiplyByScalar(effectiveColor, material.Diffuse*lightDotNormal)
		reflectv := tuple.Reflect(tuple.Negate(lightv), normalv)
		reflectDotEye := tuple.DotProduct(reflectv, eyev)

		if reflectDotEye > 0 {
			factor := math.Pow(reflectDotEye, material.Shininess)
			specular = tuple.MultiplyByScalar(light.Intensity, material.Specular*factor)
		}
	}
	return tuple.Add(ambient, tuple.Add(diffuse, specular))
}
