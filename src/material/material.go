package material

import (
	"math"

	"github.com/calbim/ray-tracer/src/pattern"

	"github.com/calbim/ray-tracer/src/tuple"

	"github.com/calbim/ray-tracer/src/color"
	"github.com/calbim/ray-tracer/src/light"
)

//Material represents the properties of a material as per the phong reflection model
type Material struct {
	Color      color.Color
	Ambient    float64
	Diffuse    float64
	Specular   float64
	Shininess  float64
	Pattern    *pattern.Pattern
	hasPattern bool
	Reflective float64
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

//SetPattern sets a pattern
func (m *Material) SetPattern(p pattern.Pattern) {
	m.Pattern = &p
	m.hasPattern = true
}

//Lighting returns the shade of an object under various light properties
func (m *Material) Lighting(object pattern.Object, light light.Light, point tuple.Tuple, eyev tuple.Tuple, normalv tuple.Tuple, inShadow bool) color.Color {
	c := m.Color
	if m.hasPattern {
		tmp := pattern.AtObject(*m.Pattern, object, point)
		c = *tmp
	}
	effectiveColor := c.MultiplyColor(light.Intensity)
	lightv := light.Position.Subtract(point)
	lightv = lightv.Normalize()
	ambient := effectiveColor.Multiply(m.Ambient)
	lightDotNormal := lightv.DotProduct(normalv)
	diffuse, specular := color.Black, color.Black
	if lightDotNormal >= 0 && !inShadow {
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
