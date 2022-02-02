package spheres

import (
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
