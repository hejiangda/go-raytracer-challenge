package raytracer

import "testing"

func TestIntersectingARayParallelToTheTriangle(t *testing.T) {
	tr := NewTriangle(Point(0, 1, 0), Point(-1, 0, 0), Point(1, 0, 0))
	r := NewRay(Point(0, -1, -2), Vector(0, 1, 0))
	xs := tr.LocalIntersect(r)
	if len(xs) != 0 {
		t.Fatal("error")
	}
}

func TestARayMissesTheP1P3Edge(t *testing.T) {
	tr := NewTriangle(Point(0, 1, 0), Point(-1, 0, 0), Point(1, 0, 0))
	r := NewRay(Point(1, 1, -2), Vector(0, 0, 1))
	xs := tr.LocalIntersect(r)
	if len(xs) != 0 {
		t.Fatal("error")
	}
}

func TestARayMissesTheP1P2Edge(t *testing.T) {
	tr := NewTriangle(Point(0, 1, 0), Point(-1, 0, 0), Point(1, 0, 0))
	r := NewRay(Point(-1, 1, -2), Vector(0, 0, 1))
	xs := tr.LocalIntersect(r)
	if len(xs) != 0 {
		t.Fatal("error")
	}
}

func TestARayMissesTheP2P3Edge(t *testing.T) {
	tr := NewTriangle(Point(0, 1, 0), Point(-1, 0, 0), Point(1, 0, 0))
	r := NewRay(Point(0, -1, -2), Vector(0, 0, 1))
	xs := tr.LocalIntersect(r)
	if len(xs) != 0 {
		t.Fatal("error")
	}
}

func TestARayStrikesATriangle(t *testing.T) {
	tr := NewTriangle(Point(0, 1, 0), Point(-1, 0, 0), Point(1, 0, 0))
	r := NewRay(Point(0, 0.5, -2), Vector(0, 0, 1))
	xs := tr.LocalIntersect(r)
	if len(xs) != 1 || xs[0].T != 2 {
		t.Fatal("error")
	}
}
