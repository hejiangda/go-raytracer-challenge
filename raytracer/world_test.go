package raytracer

import (
	"fmt"
	"testing"
)

func TestNewWorld(t *testing.T) {

	w := NewWorld()

	pointLight := NewPointLight(Point(-10, 10, -10), Color(1, 1, 1))
	var ok = false
	for _, light := range w.Lights {
		if light.Position.Equal(pointLight.Position) && light.Intensity.Equal(pointLight.Intensity) {
			ok = true
		}
	}
	if ok == false {
		t.Fatal("Does not contain light:", pointLight.Position, pointLight.Intensity)
	}

	s1 := NewSphere()
	s1.Material.Color = Color(0.8, 1.0, 0.6)
	s1.Material.Diffuse = 0.7
	s1.Material.Specular = 0.2

	s2 := NewSphere()
	s2.Transform = Scaling(0.5, 0.5, 0.5)

	var ok1 = false
	var ok2 = false
	for _, object := range w.Objects {
		switch object.(type) {
		case *Sphere:
			s := object.(*Sphere)
			if IsSame(s.Transform, s1.Transform) && s.Material.Equal(s1.Material) {
				ok1 = true
			}
			if IsSame(s.Transform, s2.Transform) && s.Material.Equal(s2.Material) {
				ok2 = true
			}
		}

	}
	if !ok1 {
		t.Fatal("Does not contain the object s1:", s1.Transform, s1.Material)
	}
	if !ok2 {
		t.Fatal("Does not contain the object: s2", s2.Transform, s2.Material)
	}
}

func TestWorld_Intersect(t *testing.T) {
	w := NewWorld()
	r := NewRay(Point(0, 0, -5), Vector(0, 0, 1))
	xs := w.Intersect(r)
	fmt.Println("xs len", len(xs))
	if xs[0].T != 4 || xs[1].T != 4.5 || xs[2].T != 5.5 || xs[3].T != 6 {
		t.Fatal("Intersect failed! xs:", xs)
	}
}
