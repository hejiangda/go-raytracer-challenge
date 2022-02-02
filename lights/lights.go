package lights

import (
	"math"
	"tuples"
)

type PointLight struct {
	Position  *tuples.Tuple
	Intensity *tuples.Tuple
}
type Material struct {
	Color     *tuples.Tuple
	Ambient   float64
	Diffuse   float64
	Specular  float64
	Shininess float64
}

func NewPointLight(p, i *tuples.Tuple) *PointLight {
	return &PointLight{p, i}
}
func NewMaterial() *Material {
	return &Material{tuples.Color(1, 1, 1), 0.1, 0.9, 0.9, 200}
}
func (m *Material) Equal(n *Material) bool {
	if m.Color.Equal(n.Color) &&
		tuples.AlmostEqual(m.Ambient, n.Ambient, tuples.Eps) &&
		tuples.AlmostEqual(m.Diffuse, n.Diffuse, tuples.Eps) &&
		tuples.AlmostEqual(m.Specular, n.Specular, tuples.Eps) &&
		tuples.AlmostEqual(m.Shininess, n.Shininess, tuples.Eps) {
		return true
	}
	return false
}
func Lighting(m *Material, light *PointLight, position *tuples.Tuple, eyev *tuples.Tuple, normalv *tuples.Tuple) *tuples.Tuple {
	effectiveColor := tuples.MultiplyColors(m.Color, light.Intensity)
	lightv := tuples.Normalize(tuples.Subtract(light.Position, position))
	ambient := effectiveColor.Multiply(m.Ambient)
	lightDotNormal := tuples.Dot(lightv, normalv)

	var diffuse, specular *tuples.Tuple

	if lightDotNormal < 0 {
		diffuse = tuples.Color(0, 0, 0)
		specular = tuples.Color(0, 0, 0)
	} else {
		diffuse = effectiveColor.Multiply(m.Diffuse).Multiply(lightDotNormal)
		reflectv := tuples.Reflect(lightv.Negate(), normalv)
		reflectDotEye := tuples.Dot(reflectv, eyev)
		if reflectDotEye <= 0 {
			specular = tuples.Color(0, 0, 0)
		} else {
			factor := math.Pow(reflectDotEye, m.Shininess)
			specular = light.Intensity.Multiply(m.Specular).Multiply(factor)
		}
	}
	return ambient.Add(diffuse).Add(specular)
}
