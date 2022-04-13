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

func TestPrepareComputations(t *testing.T) {
	r := NewRay(Point(0, 0, -5), Vector(0, 0, 1))
	shape := NewSphere()
	i := Intersection{4, shape}
	comps := PrepareComputations(i, r)
	if comps.T != i.T || comps.Object != i.Obj || !comps.Point.Equal(Point(0, 0, -1)) || !comps.EyeV.Equal(Vector(0, 0, -1)) || !comps.NormalV.Equal(Vector(0, 0, -1)) {
		t.Fatal("failed")
	}
}
func TestPrepareComputationsOutside(t *testing.T) {
	r := NewRay(Point(0, 0, -5), Vector(0, 0, 1))
	shape := NewSphere()
	i := Intersection{4, shape}
	comps := PrepareComputations(i, r)
	if comps.Inside == true {
		t.Fatal("failed")
	}
}
func TestPrepareComputationsInside(t *testing.T) {
	r := NewRay(Point(0, 0, 0), Vector(0, 0, 1))
	shape := NewSphere()
	i := Intersection{1, shape}
	comps := PrepareComputations(i, r)
	if !comps.Point.Equal(Point(0, 0, 1)) || !comps.EyeV.Equal(Vector(0, 0, -1)) || !comps.Inside || !comps.NormalV.Equal(Vector(0, 0, -1)) {
		t.Fatal("failed")
	}
}

func TestWorld_ShadeHit(t *testing.T) {
	w := NewWorld()
	r := NewRay(Point(0, 0, -5), Vector(0, 0, 1))
	shape := w.Objects[0]
	i := Intersection{4, shape}
	comps := PrepareComputations(i, r)
	c := w.ShadeHit(comps)
	if !c.Equal(Color(0.38066, 0.47583, 0.2855)) {
		t.Fatal("failed color:", c)
	}
}
func TestWorld_ShadeHitInside(t *testing.T) {
	w := NewWorld()
	w.Lights[0] = NewPointLight(Point(0, 0.25, 0), Color(1, 1, 1))
	r := NewRay(Point(0, 0, 0), Vector(0, 0, 1))
	shape := w.Objects[1]
	i := Intersection{0.5, shape}
	comps := PrepareComputations(i, r)
	c := w.ShadeHit(comps)
	if !c.Equal(Color(0.90498, 0.90498, 0.90498)) {
		t.Fatal("failed color:", c)
	}
}

func TestWorld_ColorAt(t *testing.T) {
	w := NewWorld()
	r := NewRay(Point(0, 0, -5), Vector(0, 1, 0))
	c := w.ColorAt(r)
	if !c.Equal(Color(0, 0, 0)) {
		t.Fatal("failed")
	}
}

func TestWorld_ColorAt2(t *testing.T) {
	w := NewWorld()
	r := NewRay(Point(0, 0, -5), Vector(0, 0, 1))
	c := w.ColorAt(r)
	if !c.Equal(Color(0.38066, 0.47583, 0.2855)) {
		t.Fatal("failed color:", c)
	}
}

func TestWorld_ColorAt3(t *testing.T) {
	w := NewWorld()
	outer := w.Objects[0]
	outer.GetMaterial().Ambient = 1

	inner := w.Objects[1]
	inner.GetMaterial().Ambient = 1
	r := NewRay(Point(0, 0, 0.75), Vector(0, 0, -1))
	c := w.ColorAt(r)
	if !c.Equal(inner.GetMaterial().Color) {
		t.Fatal("failed color:", c)
	}
}

func TestWorld_IsShadowed1(t *testing.T) {
	w := NewWorld()
	p := Point(0, 10, 0)
	if w.IsShadowed(p) {
		t.Fatal("failed")
	}
}
func TestWorld_IsShadowed2(t *testing.T) {
	w := NewWorld()
	p := Point(10, -10, 10)
	if !w.IsShadowed(p) {
		t.Fatal("failed")
	}
}
func TestWorld_IsShadowed3(t *testing.T) {
	w := NewWorld()
	p := Point(-20, 20, -20)
	if w.IsShadowed(p) {
		t.Fatal("failed")
	}
}

func TestWorld_IsShadowed4(t *testing.T) {
	w := NewWorld()
	p := Point(-2, 2, -2)
	if w.IsShadowed(p) {
		t.Fatal("failed")
	}
}

func TestWorld_ShadeHit2(t *testing.T) {
	w := new(World)
	w.Lights = append(w.Lights, NewPointLight(Point(0, 0, -10), Color(1, 1, 1)))
	s1 := NewSphere()
	w.Objects = append(w.Objects, s1)
	s2 := NewSphere()
	s2.Transform = Translation(0, 0, 10)
	w.Objects = append(w.Objects, s2)
	r := NewRay(Point(0, 0, 5), Vector(0, 0, 1))
	i := Intersection{4, s2}
	comps := PrepareComputations(i, r)
	c := w.ShadeHit(comps)
	if !c.Equal(Color(0.1, 0.1, 0.1)) {
		t.Fatal("failed")
	}

}
