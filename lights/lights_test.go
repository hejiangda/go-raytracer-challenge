package lights

import (
	"testing"
	"tuples"
)

func TestPointLight(t *testing.T) {
	intensity := tuples.Color(1, 1, 1)
	position := tuples.Point(0, 0, 0)
	light := NewPointLight(position, intensity)
	if !position.Equal(light.Position) {
		t.Fatal("light.position != position")
	}
	if !intensity.Equal(light.Intensity) {
		t.Fatal("light.intensity != intensity")
	}
}

func TestDefaultMaterial(t *testing.T) {
	m := NewMaterial()
	if !tuples.Color(1, 1, 1).Equal(m.Color) {
		t.Fatal("m.color != color(1, 1, 1)")
	}
	if !tuples.AlmostEqual(m.Ambient, 0.1, tuples.Eps) {
		t.Fatal("m.ambient != 0.1")
	}
	if !tuples.AlmostEqual(m.Diffuse, 0.9, tuples.Eps) {
		t.Fatal("m.diffuse != 0.9")
	}
	if !tuples.AlmostEqual(m.Specular, 0.9, tuples.Eps) {
		t.Fatal("m.specular != 0.9")
	}
	if !tuples.AlmostEqual(m.Shininess, 200, tuples.Eps) {
		t.Fatal("m.shininess != 200.0")
	}
}
