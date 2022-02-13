package raytracer

import (
	"math"
)

type PointLight struct {
	Position  *Tuple
	Intensity *Tuple
}
type Material struct {
	Color     *Tuple
	Ambient   float64
	Diffuse   float64
	Specular  float64
	Shininess float64
}

func NewPointLight(p, i *Tuple) *PointLight {
	return &PointLight{p, i}
}
func NewMaterial() *Material {
	return &Material{Color(1, 1, 1), 0.1, 0.9, 0.9, 200}
}
func (m *Material) Equal(n *Material) bool {
	if m.Color.Equal(n.Color) &&
		AlmostEqual(m.Ambient, n.Ambient, Eps) &&
		AlmostEqual(m.Diffuse, n.Diffuse, Eps) &&
		AlmostEqual(m.Specular, n.Specular, Eps) &&
		AlmostEqual(m.Shininess, n.Shininess, Eps) {
		return true
	}
	return false
}
func Lighting(m *Material, light *PointLight, position *Tuple, eyev *Tuple, normalv *Tuple) *Tuple {
	effectiveColor := MultiplyColors(m.Color, light.Intensity)
	lightv := Normalize(Subtract(light.Position, position))
	ambient := effectiveColor.Multiply(m.Ambient)
	lightDotNormal := Dot(lightv, normalv)

	var diffuse, specular *Tuple

	if lightDotNormal < 0 {
		diffuse = Color(0, 0, 0)
		specular = Color(0, 0, 0)
	} else {
		diffuse = effectiveColor.Multiply(m.Diffuse).Multiply(lightDotNormal)
		reflectv := Reflect(lightv.Negate(), normalv)
		reflectDotEye := Dot(reflectv, eyev)
		if reflectDotEye <= 0 {
			specular = Color(0, 0, 0)
		} else {
			factor := math.Pow(reflectDotEye, m.Shininess)
			specular = light.Intensity.Multiply(m.Specular).Multiply(factor)
		}
	}
	return ambient.Add(diffuse).Add(specular)
}
