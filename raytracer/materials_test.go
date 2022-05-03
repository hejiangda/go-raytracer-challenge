package raytracer

import (
	"math"
	"testing"
)

var m *Material
var position *Tuple

func init() {
	m = NewMaterial()
	position = Point(0, 0, 0)
}

func TestLightingEyeBetweenLightAndSurface(t *testing.T) {
	eyev := Vector(0, 0, -1)
	normalv := Vector(0, 0, -1)
	light := NewPointLight(Point(0, 0, -10), Color(1, 1, 1))
	obj := NewSphere()
	result := Lighting(m, obj, light, position, eyev, normalv, false)
	if !Color(1.9, 1.9, 1.9).Equal(result) {
		t.Fatal("result != color(1.9, 1.9, 1.9),result:", result)
	}
}

func TestLightingEyeBetweenLightAndSurfaceEyeOffset45(t *testing.T) {
	eyev := Vector(0, math.Sqrt2/2, math.Sqrt2/2)
	normalv := Vector(0, 0, -1)
	light := NewPointLight(Point(0, 0, -10), Color(1, 1, 1))
	obj := NewSphere()
	result := Lighting(m, obj, light, position, eyev, normalv, false)
	if !Color(1.0, 1.0, 1.0).Equal(result) {
		t.Fatal("result != color(1.0, 1.0, 1.0)")
	}
}

func TestLightingEyeBetweenLightAndSurfaceLightOffset45(t *testing.T) {
	eyev := Vector(0, 0, -1)
	normalv := Vector(0, 0, -1)
	light := NewPointLight(Point(0, 10, -10), Color(1, 1, 1))
	obj := NewSphere()
	result := Lighting(m, obj, light, position, eyev, normalv, false)
	if !Color(0.7364, 0.7364, 0.7364).Equal(result) {
		t.Fatal("result != color(0.7364, 0.7364, 0.7364)")
	}
}
func TestLightingWithEyeInPathOfReflection(t *testing.T) {
	eyev := Vector(0, -math.Sqrt2/2, -math.Sqrt2/2)
	normalv := Vector(0, 0, -1)
	light := NewPointLight(Point(0, 10, -10), Color(1, 1, 1))
	obj := NewSphere()
	result := Lighting(m, obj, light, position, eyev, normalv, false)
	if !Color(1.6364, 1.6364, 1.6364).Equal(result) {
		t.Fatal("result != color(1.6364, 1.6364, 1.6364)")
	}
}
func TestLightingWithLightBehindSurface(t *testing.T) {
	eyev := Vector(0, 0, -1)
	normalv := Vector(0, 0, -1)
	light := NewPointLight(Point(0, 0, 10), Color(1, 1, 1))
	obj := NewSphere()
	result := Lighting(m, obj, light, position, eyev, normalv, false)
	if !Color(0.1, 0.1, 0.1).Equal(result) {
		t.Fatal("result != color(0.1, 0.1, 0.1)")
	}
}

func TestLightingWithTheSurfaceInShadow(t *testing.T) {
	eyev := Vector(0, 0, -1)
	normalv := Vector(0, 0, -1)
	light := NewPointLight(Point(0, 0, -10), Color(1, 1, 1))
	obj := NewSphere()
	result := Lighting(m, obj, light, position, eyev, normalv, true)
	if !result.Equal(Color(0.1, 0.1, 0.1)) {
		t.Fatal("failed")
	}
}

func TestTransparencyAndRefractiveIndex(t *testing.T) {
	m := NewMaterial()
	if !AlmostEqual(m.Transparency, 0, Eps) ||
		!AlmostEqual(m.RefractiveIndex, 1.0, Eps) {
		t.Fatal("failed")
	}
}
