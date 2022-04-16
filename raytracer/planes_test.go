package raytracer

import "testing"

func TestNormalOfAPlaneIsConstantEverywhere(t *testing.T) {
	p := NewPlane()
	n1 := p.localNormalAt(Point(0, 0, 0))
	n2 := p.localNormalAt(Point(10, 0, -10))
	n3 := p.localNormalAt(Point(-5, 0, 150))
	if !n1.Equal(Vector(0, 1, 0)) ||
		!n2.Equal(Vector(0, 1, 0)) ||
		!n3.Equal(Vector(0, 1, 0)) {
		t.Fatal("failed")
	}
}

func TestIntersectWithARayParallelToThePlane(t *testing.T) {
	p := NewPlane()
	r := NewRay(Point(0, 10, 0), Vector(0, 0, 1))
	xs := p.Intersect(r)
	if len(xs) != 0 {
		t.Fatal("failed")
	}
}

func TestIntersectWithACoplanarRay(t *testing.T) {
	p := NewPlane()
	r := NewRay(Point(0, 0, 0), Vector(0, 0, 1))
	xs := p.Intersect(r)
	if len(xs) != 0 {
		t.Fatal("failed")
	}
}

func TestARayIntersectingAPlaneFromAbove(t *testing.T) {
	p := NewPlane()
	r := NewRay(Point(0, 1, 0), Vector(0, -1, 0))
	xs := p.Intersect(r)
	if len(xs) != 1 {
		t.Fatal("failed")
	}
	if xs[0].T != 1 || xs[0].Obj != p {
		t.Fatal("failed")
	}
}
