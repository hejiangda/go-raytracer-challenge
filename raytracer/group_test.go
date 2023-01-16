package raytracer

import (
	"testing"
)

func TestCreateGroup(t *testing.T) {
	g := NewGroup()
	trans := g.GetTransform()
	if !trans.Equal(EyeMatrix(4)) {
		t.Fatal("failed")
	}
}
func TestAddingAChildToAGroup(t *testing.T) {
	g := NewGroup()
	s := NewSphere()
	shape := interface{}(s).(Shape)
	g.AddChild(shape)
	shape.GetParent()
	if !g.HasChild(shape) || shape.GetParent() != interface{}(g).(Shape) {
		t.Fatal("failed")
	}
}

func TestIntersectingARayWithAnEmptyGroup(t *testing.T) {
	g := NewGroup()
	r := NewRay(Point(0, 0, 0), Vector(0, 0, 1))
	xs := g.Intersect(r)
	if len(xs) != 0 {
		t.Fatal("failed")
	}
}

func TestIntersectingARayWithANonemptyGroup(t *testing.T) {
	g := NewGroup()
	s1 := NewSphere()
	s2 := NewSphere()
	s2.SetTransform(Translation(0, 0, -3))
	s3 := NewSphere()
	s3.SetTransform(Translation(5, 0, 0))
	g.AddChild(s1)
	g.AddChild(s2)
	g.AddChild(s3)

	r := NewRay(Point(0, 0, -5), Vector(0, 0, 1))
	xs := g.Intersect(r)

	if len(xs) != 4 {
		t.Fatal("failed")
	}
	if xs[0].Obj != s2 || xs[1].Obj != s2 || xs[2].Obj != s1 || xs[3].Obj != s1 {
		t.Fatal("failed")
	}
}

func TestIntersectingATransformedGroup(t *testing.T) {
	g := NewGroup()
	g.SetTransform(Scaling(2, 2, 2))
	s := NewSphere()
	s.SetTransform(Translation(5, 0, 0))
	g.AddChild(s)
	r := NewRay(Point(10, 0, -10), Vector(0, 0, 1))
	xs := g.Intersect(r)

	if len(xs) != 2 {
		t.Fatal("failed")
	}
}
