package lights

import (
	"math"
	"testing"
	"tuples"
)

var m *Material
var position *tuples.Tuple

func init() {
	m = NewMaterial()
	position = tuples.Point(0, 0, 0)
}

func TestLightingEyeBetweenLightAndSurface(t *testing.T) {
	eyev := tuples.Vector(0, 0, -1)
	normalv := tuples.Vector(0, 0, -1)
	light := NewPointLight(tuples.Point(0, 0, -10), tuples.Color(1, 1, 1))
	result := Lighting(m, light, position, eyev, normalv)
	if !tuples.Color(1.9, 1.9, 1.9).Equal(result) {
		t.Fatal("result != color(1.9, 1.9, 1.9),result:", result)
	}
}

func TestLightingEyeBetweenLightAndSurfaceEyeOffset45(t *testing.T) {
	eyev := tuples.Vector(0, math.Sqrt2/2, math.Sqrt2/2)
	normalv := tuples.Vector(0, 0, -1)
	light := NewPointLight(tuples.Point(0, 0, -10), tuples.Color(1, 1, 1))
	result := Lighting(m, light, position, eyev, normalv)
	if !tuples.Color(1.0, 1.0, 1.0).Equal(result) {
		t.Fatal("result != color(1.0, 1.0, 1.0)")
	}
}

func TestLightingEyeBetweenLightAndSurfaceLightOffset45(t *testing.T) {
	eyev := tuples.Vector(0, 0, -1)
	normalv := tuples.Vector(0, 0, -1)
	light := NewPointLight(tuples.Point(0, 10, -10), tuples.Color(1, 1, 1))
	result := Lighting(m, light, position, eyev, normalv)
	if !tuples.Color(0.7364, 0.7364, 0.7364).Equal(result) {
		t.Fatal("result != color(0.7364, 0.7364, 0.7364)")
	}
}
func TestLightingWithEyeInPathOfReflection(t *testing.T) {
	eyev := tuples.Vector(0, -math.Sqrt2/2, -math.Sqrt2/2)
	normalv := tuples.Vector(0, 0, -1)
	light := NewPointLight(tuples.Point(0, 10, -10), tuples.Color(1, 1, 1))
	result := Lighting(m, light, position, eyev, normalv)
	if !tuples.Color(1.6364, 1.6364, 1.6364).Equal(result) {
		t.Fatal("result != color(1.6364, 1.6364, 1.6364)")
	}
}
func TestLightingWithLightBehindSurface(t *testing.T) {
	eyev := tuples.Vector(0, 0, -1)
	normalv := tuples.Vector(0, 0, -1)
	light := NewPointLight(tuples.Point(0, 0, 10), tuples.Color(1, 1, 1))
	result := Lighting(m, light, position, eyev, normalv)
	if !tuples.Color(0.1, 0.1, 0.1).Equal(result) {
		t.Fatal("result != color(0.1, 0.1, 0.1)")
	}
}
