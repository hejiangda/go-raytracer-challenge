package raytracer

import (
	"testing"
)

func TestIntersection(t *testing.T) {
	s := NewSphere()
	i := Intersection{
		T:   3.5,
		Obj: s,
	}
	if i.T != 3.5 {
		t.Fatal("i.T!=3.5")
	}
	if i.Obj != s {
		t.Fatal("i.Obj!=s")
	}
}

func TestAggregatingIntersections(t *testing.T) {
	s := NewSphere()
	i1 := Intersection{1, s}
	i2 := Intersection{2, s}
	xs := Intersections(i1, i2)
	if len(xs) != 2 {
		t.Fatal("len(xs)!=2")
	}
	if xs[0].T != 1 {
		t.Fatal("xs[0].T!=1")
	}
	if xs[1].T != 2 {
		t.Fatal("xs[1].T!=2")
	}
}

func TestIntersectSetsTheObjectsOnTheIntersection(t *testing.T) {
	r := NewRay(Point(0, 0, -5), Vector(0, 0, 1))
	s := NewSphere()
	xs := Intersect(s, r)
	if len(xs) != 2 {
		t.Fatal("len(xs)!=2")
	}
	if xs[0].Obj != s {
		t.Fatal("xs[0].Obj!=s")
	}
	if xs[1].Obj != s {
		t.Fatal("xs[1].Obj!=s")
	}
}

func TestHit(t *testing.T) {
	s := NewSphere()
	i1 := Intersection{1, s}
	i2 := Intersection{2, s}
	xs := Intersections(i2, i1)
	i, err := Hit(xs)
	if err != nil {
		t.Fatal(err)
	}
	if i != i1 {
		t.Fatal("i!=i1")
	}
}

func TestHit2(t *testing.T) {
	s := NewSphere()
	i1 := Intersection{-1, s}
	i2 := Intersection{1, s}
	xs := Intersections(i2, i1)
	i, err := Hit(xs)
	if err != nil {
		t.Fatal(err)
	}
	if i != i2 {
		t.Fatal("i!=i2")
	}
}
func TestHit3(t *testing.T) {
	s := NewSphere()
	i1 := Intersection{-2, s}
	i2 := Intersection{-1, s}
	xs := Intersections(i2, i1)
	_, err := Hit(xs)
	if err == nil {
		t.Fatal("find hit")
	}
}

func TestHit4(t *testing.T) {
	s := NewSphere()
	i1 := Intersection{5, s}
	i2 := Intersection{7, s}
	i3 := Intersection{-3, s}
	i4 := Intersection{2, s}

	xs := Intersections(i2, i1, i3, i4)
	i, err := Hit(xs)
	if err != nil {
		t.Fatal(err)
	}
	if i != i4 {
		t.Fatal("i!=i4")
	}
}
func TestIntersectAScaledSphereWithARay(t *testing.T) {
	r := NewRay(Point(0, 0, -5), Vector(0, 0, 1))
	s := NewSphere()
	s.SetTransform(Scaling(2, 2, 2))
	xs := Intersect(s, r)
	if len(xs) != 2 {
		t.Fatal("xs.count != 2")
	}
	if xs[0].T != 3 {
		t.Fatal("xs[0].t != 3,val:", xs[0].T)
	}
	if xs[1].T != 7 {
		t.Fatal("xs[1].t != 7")
	}
}

func TestIntersectATranslatedSphereWithARay(t *testing.T) {
	r := NewRay(Point(0, 0, -5), Vector(0, 0, 1))
	s := NewSphere()
	s.SetTransform(Translation(5, 0, 0))
	xs := s.Intersect(r)
	if len(xs) != 0 {
		t.Fatal("xs.count != 0")
	}
}

func TestPrepareComputations2(t *testing.T) {
	r := NewRay(Point(0, 0, -5), Vector(0, 0, 1))
	shape := NewSphere()
	shape.Transform = Translation(0, 0, 1)
	i := Intersection{5, shape}
	comps := PrepareComputations(i, r)
	if !(comps.OverPoint.Z < -Eps/2) || !(comps.Point.Z > comps.OverPoint.Z) {
		t.Fatal("failed")
	}
}
