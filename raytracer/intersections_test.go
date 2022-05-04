package raytracer

import (
	"math"
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
	xs := s.Intersect(r)
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
	xs := s.Intersect(r)
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
	comps := PrepareComputations(i, r, []Intersection{i})
	if !(comps.OverPoint.Z < -Eps/2) || !(comps.Point.Z > comps.OverPoint.Z) {
		t.Fatal("failed")
	}
}

func TestPrepareComputations3(t *testing.T) {
	shape := NewPlane()
	r := NewRay(Point(0, 1, -1), Vector(0, -math.Sqrt2/2, math.Sqrt2/2))
	i := Intersection{
		T:   math.Sqrt2,
		Obj: shape,
	}
	comps := PrepareComputations(i, r, []Intersection{i})
	if !comps.ReflectV.Equal(Vector(0, math.Sqrt2/2, math.Sqrt2/2)) {
		t.Fatal("failed")
	}
}
func TestFindingN1AndN2AtVariousIntersections(t *testing.T) {
	A := GlassSphere()
	A.Transform = Scaling(2, 2, 2)
	A.Material.RefractiveIndex = 1.5

	B := GlassSphere()
	B.Transform = Translation(0, 0, -0.25)
	B.Material.RefractiveIndex = 2.0

	C := GlassSphere()
	C.Transform = Translation(0, 0, 0.25)
	C.Material.RefractiveIndex = 2.5

	r := NewRay(Point(0, 0, -4), Vector(0, 0, 1))
	xs := Intersections(Intersection{2, A},
		Intersection{2.75, B},
		Intersection{3.25, C},
		Intersection{4.75, B},
		Intersection{5.25, C},
		Intersection{6, A})
	n1Tab := []float64{1.0, 1.5, 2.0, 2.5, 2.5, 1.5}
	n2Tab := []float64{1.5, 2.0, 2.5, 2.5, 1.5, 1.0}
	for i, x := range xs {
		comps := PrepareComputations(x, r, xs)
		if !(comps.N1 == n1Tab[i]) ||
			!(comps.N2 == n2Tab[i]) {
			t.Fatal()
		}
	}
}

func TestTheUnderPointIsOffsetBelowTheSurface(t *testing.T) {
	r := NewRay(Point(0, 0, -5), Vector(0, 0, 1))
	shape := GlassSphere()
	shape.Transform = Translation(0, 0, 1)
	i := Intersection{5, shape}
	xs := Intersections(i)
	comps := PrepareComputations(i, r, xs)
	if !(comps.UnderPoint.Z > Eps/2 && comps.Point.Z < comps.UnderPoint.Z) {
		t.Fatal()
	}
}

func TestSchlickApproximationUnderTotalInternalReflection(t *testing.T) {
	shape := GlassSphere()
	r := NewRay(Point(0, 0, math.Sqrt2/2), Vector(0, 1, 0))
	xs := Intersections(Intersection{-math.Sqrt2 / 2, shape}, Intersection{math.Sqrt2 / 2, shape})
	comps := PrepareComputations(xs[1], r, xs)
	reflectance := Schlick(comps)
	if reflectance != 1.0 {
		t.Fatal()
	}
}

func TestSchlickApproximationWithAPerpendicularViewingAngle(t *testing.T) {
	shape := GlassSphere()
	r := NewRay(Point(0, 0, 0), Vector(0, 1, 0))
	xs := Intersections(Intersection{-1, shape}, Intersection{1, shape})
	comps := PrepareComputations(xs[1], r, xs)
	reflectance := Schlick(comps)
	if !AlmostEqual(reflectance, 0.04, Eps) {
		t.Fatal()
	}
}

func TestSchlickApproximationWithSmallAngleAndN2LargerThanN1(t *testing.T) {
	shape := GlassSphere()
	r := NewRay(Point(0, 0.99, -2), Vector(0, 0, 1))
	xs := Intersections(Intersection{1.8589, shape})
	comps := PrepareComputations(xs[0], r, xs)
	reflectance := Schlick(comps)
	if !AlmostEqual(reflectance, 0.48873, Eps) {
		t.Fatal(reflectance)
	}
}
