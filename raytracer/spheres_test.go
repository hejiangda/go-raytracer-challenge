package raytracer

import (
	"math"
	"testing"
)

func TestRayIntersectsSphereAtTwoPoints(t *testing.T) {
	r := NewRay(Point(0, 0, -5), Vector(0, 0, 1))
	s := NewSphere()
	xs := s.localIntersect(r)
	if len(xs) != 2 {
		t.Fatal("intersect error")
	}
	if !AlmostEqual(xs[0], 4.0, Eps) {
		t.Fatal("intersect error,xs[0]:", xs[0])
	}
	if !AlmostEqual(xs[1], 6.0, Eps) {
		t.Fatal("intersect error")
	}
}
func TestRayIntersectsSphereAtATangent(t *testing.T) {
	r := NewRay(Point(0, 1, -5), Vector(0, 0, 1))
	s := NewSphere()
	xs := s.localIntersect(r)
	if len(xs) != 2 {
		t.Fatal("intersect error")
	}
	if !AlmostEqual(xs[0], 5.0, Eps) {
		t.Fatal("intersect error")
	}
	if !AlmostEqual(xs[1], 5.0, Eps) {
		t.Fatal("intersect error")
	}
}
func TestRayMissesSphere(t *testing.T) {
	r := NewRay(Point(0, 2, -5), Vector(0, 0, 1))
	s := NewSphere()
	xs := s.Intersect(r)
	if len(xs) != 0 {
		t.Fatal("the ray is not miss the sphere")
	}
}
func TestRayOriginatesInsideTheSphere(t *testing.T) {
	r := NewRay(Point(0, 0, 0), Vector(0, 0, 1))
	s := NewSphere()
	xs := s.localIntersect(r)
	if len(xs) != 2 {
		t.Fatal("intersect error")
	}
	if !AlmostEqual(xs[0], -1.0, Eps) {
		t.Fatal("intersect error")
	}
	if !AlmostEqual(xs[1], 1.0, Eps) {
		t.Fatal("intersect error")
	}
}
func TestSphereDefaultTransformation(t *testing.T) {
	s := NewSphere()
	if !s.Transform.Equal(EyeMatrix(4)) {
		t.Fatal("s.transform != identity_matrix")
	}
}

func TestChangingSphereTransformation(t *testing.T) {
	s := NewSphere()
	tt := Translation(2, 3, 4)
	s.SetTransform(tt)
	if !s.Transform.Equal(tt) {
		t.Fatal("s.transform != t")
	}
}

func TestNormalAtSphere(t *testing.T) {
	s := NewSphere()
	n := s.NormalAt(Point(1, 0, 0))
	if !Vector(1, 0, 0).Equal(n) {
		t.Fatal("n != vector(1,0,0)")
	}
}
func TestNormalAtSphere1(t *testing.T) {
	s := NewSphere()
	n := s.NormalAt(Point(0, 1, 0))
	if !Vector(0, 1, 0).Equal(n) {
		t.Fatal("n != vector(0,1,0)")
	}
}
func TestNormalAtSphere2(t *testing.T) {
	s := NewSphere()
	n := s.NormalAt(Point(0, 0, 1))
	if !Vector(0, 0, 1).Equal(n) {
		t.Fatal("n != vector(0,0,1)")
	}
}
func TestNormalAtSphere3(t *testing.T) {
	s := NewSphere()
	n := s.NormalAt(Point(math.Sqrt(3)/3, math.Sqrt(3)/3, math.Sqrt(3)/3))
	if !Vector(math.Sqrt(3)/3, math.Sqrt(3)/3, math.Sqrt(3)/3).Equal(n) {
		t.Fatal("n != vector(0,0,1)")
	}
}
func TestNormalAtSphere4(t *testing.T) {
	s := NewSphere()
	n := s.NormalAt(Point(math.Sqrt(3)/3, math.Sqrt(3)/3, math.Sqrt(3)/3))

	if !Normalize(n).Equal(n) {
		t.Fatal("n != normalize(n)")
	}
}
func TestComputingTheNormalOnTranslatedSphere(t *testing.T) {
	s := NewSphere()
	s.SetTransform(Translation(0, 1, 0))
	n := s.NormalAt(Point(0, 1.70711, -0.70711))
	if !n.Equal(Vector(0, 0.70711, -0.70711)) {
		t.Fatal("n != vector(0, 0.70711, -0.70711)")
	}
}
func TestComputingTheNormalOnTranslatedSphere2(t *testing.T) {
	s := NewSphere()
	m := Scaling(1, 0.5, 1).Multiply(RotationZ(math.Pi / 5))
	s.SetTransform(m)
	n := s.NormalAt(Point(0, math.Sqrt2/2, math.Sqrt2/2))
	if n.Equal(Vector(0, 0.97014, -0.24254)) {
		t.Fatal("n != vector(0, 0.97014, -0.24254)")
	}
}

func TestSphereMaterial(t *testing.T) {
	s := NewSphere()
	m := NewMaterial()
	if !m.Equal(s.Material) {
		t.Fatal("m != material(),m:", m, "s.M", s.Material)
	}
}
func TestSphereAssignedAMaterial(t *testing.T) {
	s := NewSphere()
	m := NewMaterial()
	m.Ambient = 1
	s.Material = m
	if !s.Material.Equal(m) {
		t.Fatal("s.material != m")
	}
}

func TestGlassSphere(t *testing.T) {
	s := GlassSphere()
	if !s.Transform.Equal(EyeMatrix(4)) ||
		!AlmostEqual(s.Material.Transparency, 1.0, Eps) ||
		!AlmostEqual(s.Material.RefractiveIndex, 1.5, Eps) {
		t.Fatal()
	}
}
