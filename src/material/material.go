package material

import (
	"math"

	"github.com/calbim/ray-tracer/src/tuple"

	"github.com/calbim/ray-tracer/src/color"
	"github.com/calbim/ray-tracer/src/light"
)

//Material represents the properties of a material as per the phong reflection model
type Material struct {
	Color     color.Color
	Ambient   float64
	Diffuse   float64
	Specular  float64
	Shininess float64
}

//New returns a default material
func New() Material {
	return Material{
		Color:     color.New(1, 1, 1),
		Ambient:   0.1,
		Diffuse:   0.9,
		Specular:  0.9,
		Shininess: 200,
	}
}

//Lighting returns the shade of an object under various light properties
func (m *Material) Lighting(light light.Light, point tuple.Tuple, eyev tuple.Tuple, normalv tuple.Tuple) color.Color {
	c := m.Color
	effectiveColor := c.MultiplyColor(light.Intensity)
	lightv := light.Position.Subtract(point)
	lightv = lightv.Normalize()
	ambient := effectiveColor.Multiply(m.Ambient)
	lightDotNormal := lightv.DotProduct(normalv)
	diffuse, specular := color.Black, color.Black
	if lightDotNormal >= 0 {
		diffuse = effectiveColor.Multiply(m.Diffuse * lightDotNormal)
		reflectv := lightv.Negate()
		reflectv = reflectv.Reflect(normalv)
		reflectDotEye := reflectv.DotProduct(eyev)
		if reflectDotEye > 0 {
			factor := math.Pow(reflectDotEye, m.Shininess)
			specular = light.Intensity.Multiply(m.Specular * factor)
		}
	}
	sum := diffuse.Add(specular)
	return sum.Add(ambient)
}
