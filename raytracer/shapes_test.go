package raytracer

import (
	"math"
	"testing"
)

func TestGetDefaultShapeTransformation(t *testing.T) {
	s := NewShape("sphere")
	trans := s.GetTransform()
	if !trans.Equal(EyeMatrix(4)) {
		t.Fatal("failed")
	}
}

func TestSetShapeTransformation(t *testing.T) {
	s := NewShape("sphere")
	s.SetTransform(Translation(2, 3, 4))
	if !s.GetTransform().Equal(Translation(2, 3, 4)) {
		t.Fatal("failed")
	}
}

func TestGetShapeMaterial(t *testing.T) {
	s := NewShape("sphere")
	m := s.GetMaterial()
	if !m.Equal(NewMaterial()) {
		t.Fatal("failed")
	}
}

func TestSetShapeMaterial(t *testing.T) {
	s := NewShape("sphere")
	m := NewMaterial()
	m.Ambient = 1
	s.SetMaterial(m)
	if !s.GetMaterial().Equal(m) {
		t.Fatal("failed")
	}
}

func TestIntersectingAScaledShapeWithARay(t *testing.T) {
	r := NewRay(Point(0, 0, -5), Vector(0, 0, 1))
	s := NewShape("sphere")
	s.SetTransform(Scaling(2, 2, 2))
	s.Intersect(r)
	if !s.(*Sphere).localRay.Origin.Equal(Point(0, 0, -2.5)) {
		t.Fatal("failed")
	}
	if !s.(*Sphere).localRay.Direction.Equal(Vector(0, 0, 0.5)) {
		t.Fatal("failed")
	}
}
func TestIntersectingATranslatedShapeWithARay(t *testing.T) {
	r := NewRay(Point(0, 0, -5), Vector(0, 0, 1))
	s := NewShape("sphere")
	s.SetTransform(Translation(5, 0, 0))
	s.Intersect(r)
	if !s.(*Sphere).localRay.Origin.Equal(Point(-5, 0, -5)) {
		t.Fatal("failed")
	}
	if !s.(*Sphere).localRay.Direction.Equal(Vector(0, 0, 1)) {
		t.Fatal("failed")
	}
}

func TestComputingTheNormalOnATranslatedShape(t *testing.T) {
	s := NewShape("sphere")
	s.SetTransform(Translation(0, 1, 0))
	n := s.NormalAt(Point(0, 1.70711, -0.70711))
	if !n.Equal(Vector(0, 0.70711, -0.70711)) {
		t.Fatal("failed")
	}
}

func TestComputingTheNormalOnATransformedShape(t *testing.T) {
	s := NewShape("sphere")
	m := Scaling(1, 0.5, 1).Multiply(RotationZ(math.Pi / 5))
	s.SetTransform(m)
	n := s.NormalAt(Point(0, math.Sqrt2/2, -math.Sqrt2/2))
	if !n.Equal(Vector(0, 0.97014, -0.24254)) {
		t.Fatal("failed n:", n)
	}
}

func TestShapeHasParent(t *testing.T) {
	s := NewShape("sphere")
	if s.GetParent() != nil {
		t.Fatal("default shape should have no parent!")
	}
}

func TestConvertingAPointFromWorldToObjectSpace(t *testing.T) {
	g1 := NewGroup()
	g1.SetTransform(RotationY(math.Pi / 2))
	g2 := NewGroup()
	g2.SetTransform(Scaling(2, 2, 2))
	g1.AddChild(g2)
	s := NewSphere()
	s.SetTransform(Translation(5, 0, 0))
	g2.AddChild(s)

	p := s.World2Object(Point(-2, 0, -10))
	if Point(0, 0, -1).Equal(p) == false {
		t.Fatal("failed")
	}
}

func TestConvertingANormalFromObjectToWorldSpace(t *testing.T) {
	g1 := NewGroup()
	g1.SetTransform(RotationY(math.Pi / 2))
	g2 := NewGroup()
	g2.SetTransform(Scaling(1, 2, 3))
	g1.AddChild(g2)
	s := NewSphere()
	s.SetTransform(Translation(5, 0, 0))
	g2.AddChild(s)
	n := s.Normal2World(Vector(math.Sqrt(3)/3, math.Sqrt(3)/3, math.Sqrt(3)/3))
	//fmt.Println(n)
	if !n.Equal(Vector(0.28571, 0.42857, -0.85714)) {
		t.Fatal("failed")
	}
}

func TestFindingTheNormalOnAChildObject(t *testing.T) {
	g1 := NewGroup()
	g1.SetTransform(RotationY(math.Pi / 2))
	g2 := NewGroup()
	g2.SetTransform(Scaling(1, 2, 3))
	g1.AddChild(g2)
	s := NewSphere()
	s.SetTransform(Translation(5, 0, 0))
	g2.AddChild(s)
	n := s.NormalAt(Point(1.7321, 1.1547, -5.5774))
	//fmt.Println(n)
	if !n.Equal(Vector(0.2857, 0.42854, -0.85716)) {
		t.Fatal("failed")
	}
}
