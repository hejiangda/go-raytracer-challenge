package spheres

import (
	"lights"
	"math"
	"matrices"
	"rays"
	"testing"
	"transformations"
	"tuples"
)

func TestRayIntersectsSphereAtTwoPoints(t *testing.T) {
	r := rays.NewRay(tuples.Point(0, 0, -5), tuples.Vector(0, 0, 1))
	s := NewSphere()
	xs := Intersect(s, r)
	if len(xs) != 2 {
		t.Fatal("intersect error")
	}
	if !tuples.AlmostEqual(xs[0], 4.0, tuples.Eps) {
		t.Fatal("intersect error,xs[0]:", xs[0])
	}
	if !tuples.AlmostEqual(xs[1], 6.0, tuples.Eps) {
		t.Fatal("intersect error")
	}
}
func TestRayIntersectsSphereAtATangent(t *testing.T) {
	r := rays.NewRay(tuples.Point(0, 1, -5), tuples.Vector(0, 0, 1))
	s := NewSphere()
	xs := Intersect(s, r)
	if len(xs) != 2 {
		t.Fatal("intersect error")
	}
	if !tuples.AlmostEqual(xs[0], 5.0, tuples.Eps) {
		t.Fatal("intersect error")
	}
	if !tuples.AlmostEqual(xs[1], 5.0, tuples.Eps) {
		t.Fatal("intersect error")
	}
}
func TestRayMissesSphere(t *testing.T) {
	r := rays.NewRay(tuples.Point(0, 2, -5), tuples.Vector(0, 0, 1))
	s := NewSphere()
	xs := Intersect(s, r)
	if len(xs) != 0 {
		t.Fatal("the ray is not miss the sphere")
	}
}
func TestRayOriginatesInsideTheSphere(t *testing.T) {
	r := rays.NewRay(tuples.Point(0, 0, 0), tuples.Vector(0, 0, 1))
	s := NewSphere()
	xs := Intersect(s, r)
	if len(xs) != 2 {
		t.Fatal("intersect error")
	}
	if !tuples.AlmostEqual(xs[0], -1.0, tuples.Eps) {
		t.Fatal("intersect error")
	}
	if !tuples.AlmostEqual(xs[1], 1.0, tuples.Eps) {
		t.Fatal("intersect error")
	}
}
func TestSphereDefaultTransformation(t *testing.T) {
	s := NewSphere()
	if !matrices.IsSame(s.Transform, matrices.EyeMatrix(4)) {
		t.Fatal("s.transform != identity_matrix")
	}
}

func TestChangingSphereTransformation(t *testing.T) {
	s := NewSphere()
	tt := transformations.Translation(2, 3, 4)
	SetTransform(s, tt)
	if !matrices.IsSame(s.Transform, tt) {
		t.Fatal("s.transform != t")
	}
}

func TestNormalAtSphere(t *testing.T) {
	s := NewSphere()
	n := NormalAt(s, tuples.Point(1, 0, 0))
	if !tuples.Vector(1, 0, 0).Equal(n) {
		t.Fatal("n != vector(1,0,0)")
	}
}
func TestNormalAtSphere1(t *testing.T) {
	s := NewSphere()
	n := NormalAt(s, tuples.Point(0, 1, 0))
	if !tuples.Vector(0, 1, 0).Equal(n) {
		t.Fatal("n != vector(0,1,0)")
	}
}
func TestNormalAtSphere2(t *testing.T) {
	s := NewSphere()
	n := NormalAt(s, tuples.Point(0, 0, 1))
	if !tuples.Vector(0, 0, 1).Equal(n) {
		t.Fatal("n != vector(0,0,1)")
	}
}
func TestNormalAtSphere3(t *testing.T) {
	s := NewSphere()
	n := NormalAt(s, tuples.Point(math.Sqrt(3)/3, math.Sqrt(3)/3, math.Sqrt(3)/3))
	if !tuples.Vector(math.Sqrt(3)/3, math.Sqrt(3)/3, math.Sqrt(3)/3).Equal(n) {
		t.Fatal("n != vector(0,0,1)")
	}
}
func TestNormalAtSphere4(t *testing.T) {
	s := NewSphere()
	n := NormalAt(s, tuples.Point(math.Sqrt(3)/3, math.Sqrt(3)/3, math.Sqrt(3)/3))

	if !tuples.Normalize(n).Equal(n) {
		t.Fatal("n != normalize(n)")
	}
}
func TestComputingTheNormalOnTranslatedSphere(t *testing.T) {
	s := NewSphere()
	s.SetTransform(transformations.Translation(0, 1, 0))
	n := NormalAt(s, tuples.Point(0, 1.70711, -0.70711))
	if !n.Equal(tuples.Vector(0, 0.70711, -0.70711)) {
		t.Fatal("n != vector(0, 0.70711, -0.70711)")
	}
}
func TestComputingTheNormalOnTranslatedSphere2(t *testing.T) {
	s := NewSphere()
	m := transformations.Scaling(1, 0.5, 1).Multiply(transformations.RotationZ(math.Pi / 5))
	s.SetTransform(m)
	n := NormalAt(s, tuples.Point(0, math.Sqrt2/2, math.Sqrt2/2))
	if n.Equal(tuples.Vector(0, 0.97014, -0.24254)) {
		t.Fatal("n != vector(0, 0.97014, -0.24254)")
	}
}

func TestSphereMaterial(t *testing.T) {
	s := NewSphere()
	m := lights.NewMaterial()
	if !m.Equal(s.Material) {
		t.Fatal("m != material(),m:", m, "s.M", s.Material)
	}
}
func TestSphereAssignedAMaterial(t *testing.T) {
	s := NewSphere()
	m := lights.NewMaterial()
	m.Ambient = 1
	s.Material = m
	if !s.Material.Equal(m) {
		t.Fatal("s.material != m")
	}
}
