package transformations

import (
	"math"
	"matrices"
	"testing"
	"tuples"
)

func TestTranslation(t *testing.T) {
	trans := Translation(5, -3, 2)
	p := tuples.Point(-3, 4, 5)
	p1 := matrices.MultiplyTuple(trans, p)
	if !p1.Equal(tuples.Point(2, 1, 7)) {
		t.Fatal("translate failed")
	}
}

func TestInverseOfATranslation(t *testing.T) {
	trans := Translation(5, -3, 2)
	inv, err := matrices.Inverse(trans)
	if err != nil {
		t.Fatal(err)
	}
	p := tuples.Point(-3, 4, 5)
	if !matrices.MultiplyTuple(inv, p).Equal(tuples.Point(-8, 7, 3)) {
		t.Fatal("multiply by the inverse of a transformation matrix failed")
	}
}
func TestTranslationDoesNotAffectVector(t *testing.T) {
	trans := Translation(5, -3, 2)
	v := tuples.Vector(-3, 4, 5)
	if !v.Equal(matrices.MultiplyTuple(trans, v)) {
		t.Fatal("translation affect the vector")
	}
}
func TestScaling(t *testing.T) {
	trans := Scaling(2, 3, 4)
	p := tuples.Point(-4, 6, 8)
	if !tuples.Point(-8, 18, 32).Equal(matrices.MultiplyTuple(trans, p)) {
		t.Fatal("scaling failed")
	}
}

func TestScaling2(t *testing.T) {
	trans := Scaling(2, 3, 4)
	v := tuples.Vector(-4, 6, 8)
	if !tuples.Vector(-8, 18, 32).Equal(matrices.MultiplyTuple(trans, v)) {
		t.Fatal("scaling failed")
	}
}

func TestInverseOfScaling(t *testing.T) {
	trans := Scaling(2, 3, 4)
	inv, err := matrices.Inverse(trans)
	if err != nil {
		t.Fatal(err)
	}
	v := tuples.Vector(-4, 6, 8)
	res := matrices.MultiplyTuple(inv, v)
	if !tuples.Vector(-2, 2, 2).Equal(res) {
		t.Fatal("scaling failed,v:", res)
	}
}

func TestReflection(t *testing.T) {
	trans := Scaling(-1, 1, 1)
	p := tuples.Point(2, 3, 4)
	if !tuples.Point(-2, 3, 4).Equal(matrices.MultiplyTuple(trans, p)) {
		t.Fatal("reflect failed")
	}
}

func TestRotatingPointAroundXAxis(t *testing.T) {
	p := tuples.Point(0, 1, 0)
	half_quarter := RotationX(math.Pi / 4)
	full_quarter := RotationX(math.Pi / 2)
	p1 := matrices.MultiplyTuple(half_quarter, p)
	p2 := matrices.MultiplyTuple(full_quarter, p)
	if !tuples.Point(0, math.Sqrt2/2, math.Sqrt2/2).Equal(p1) {
		t.Fatal("rotate failed")
	}
	if !tuples.Point(0, 0, 1).Equal(p2) {
		t.Fatal("rotate failed")
	}

}
func TestInverseRotatingAroundXAxis(t *testing.T) {
	p := tuples.Point(0, 1, 0)
	half_quarter := RotationX(math.Pi / 4)
	inv, err := matrices.Inverse(half_quarter)
	if err != nil {
		t.Fatal(err)
	}
	if !tuples.Point(0, math.Sqrt2/2, -math.Sqrt2/2).Equal(matrices.MultiplyTuple(inv, p)) {
		t.Fatal("inverse rotate failed")
	}
}
func TestRotatingPointAroundYAxis(t *testing.T) {
	p := tuples.Point(0, 0, 1)
	half_quarter := RotationY(math.Pi / 4)
	full_quarter := RotationY(math.Pi / 2)
	p1 := matrices.MultiplyTuple(half_quarter, p)
	p2 := matrices.MultiplyTuple(full_quarter, p)
	if !tuples.Point(math.Sqrt2/2, 0, math.Sqrt2/2).Equal(p1) {
		t.Fatal("rotate failed")
	}
	if !tuples.Point(1, 0, 0).Equal(p2) {
		t.Fatal("rotate failed")
	}
}
func TestRotatingPointAroundZAxis(t *testing.T) {
	p := tuples.Point(0, 1, 0)
	half_quarter := RotationZ(math.Pi / 4)
	full_quarter := RotationZ(math.Pi / 2)
	p1 := matrices.MultiplyTuple(half_quarter, p)
	p2 := matrices.MultiplyTuple(full_quarter, p)
	if !tuples.Point(-math.Sqrt2/2, math.Sqrt2/2, 0).Equal(p1) {
		t.Fatal("rotate failed")
	}
	if !tuples.Point(-1, 0, 0).Equal(p2) {
		t.Fatal("rotate failed")
	}
}
func TestShearingXY(t *testing.T) {
	trans := Shearing(1, 0, 0, 0, 0, 0)
	p := tuples.Point(2, 3, 4)
	if !tuples.Point(5, 3, 4).Equal(matrices.MultiplyTuple(trans, p)) {
		t.Fatal("shearing xy failed")
	}
}

func TestShearingXZ(t *testing.T) {
	trans := Shearing(0, 1, 0, 0, 0, 0)
	p := tuples.Point(2, 3, 4)
	if !tuples.Point(6, 3, 4).Equal(matrices.MultiplyTuple(trans, p)) {
		t.Fatal("shearing xy failed")
	}
}

func TestShearingyx(t *testing.T) {
	trans := Shearing(0, 0, 1, 0, 0, 0)
	p := tuples.Point(2, 3, 4)
	if !tuples.Point(2, 5, 4).Equal(matrices.MultiplyTuple(trans, p)) {
		t.Fatal("shearing xy failed")
	}
}
func TestShearingyz(t *testing.T) {
	trans := Shearing(0, 0, 0, 1, 0, 0)
	p := tuples.Point(2, 3, 4)
	if !tuples.Point(2, 7, 4).Equal(matrices.MultiplyTuple(trans, p)) {
		t.Fatal("shearing xy failed")
	}
}
func TestShearingzx(t *testing.T) {
	trans := Shearing(0, 0, 0, 0, 1, 0)
	p := tuples.Point(2, 3, 4)
	if !tuples.Point(2, 3, 6).Equal(matrices.MultiplyTuple(trans, p)) {
		t.Fatal("shearing xy failed")
	}
}
func TestShearingzy(t *testing.T) {
	trans := Shearing(0, 0, 0, 0, 0, 1)
	p := tuples.Point(2, 3, 4)
	if !tuples.Point(2, 3, 7).Equal(matrices.MultiplyTuple(trans, p)) {
		t.Fatal("shearing xy failed")
	}
}

func TestIndividualTransformations(t *testing.T) {
	p := tuples.Point(1, 0, 1)
	a := RotationX(math.Pi / 2)
	b := Scaling(5, 5, 5)
	c := Translation(10, 5, 7)
	p2 := matrices.MultiplyTuple(a, p)
	if !tuples.Point(1, -1, 0).Equal(p2) {
		t.Fatal("rotate failed")
	}
	p3 := matrices.MultiplyTuple(b, p2)
	if !tuples.Point(5, -5, 0).Equal(p3) {
		t.Fatal("scale failed")
	}
	p4 := matrices.MultiplyTuple(c, p3)
	if !tuples.Point(15, 0, 7).Equal(p4) {
		t.Fatal("translate failed")
	}
}
func TestChainedTransformations(t *testing.T) {
	p := tuples.Point(1, 0, 1)
	a := RotationX(math.Pi / 2)
	b := Scaling(5, 5, 5)
	c := Translation(10, 5, 7)
	tt := c.Multiply(b).Multiply(a)
	if !tuples.Point(15, 0, 7).Equal(matrices.MultiplyTuple(tt, p)) {
		t.Fatal("chained transform failed")
	}
}
func TestFluentAPIs(t *testing.T) {
	p := tuples.Point(1, 0, 1)
	tt := matrices.EyeMatrix(4).Multiply(Translation(10, 5, 7)).Multiply(Scaling(5, 5, 5)).Multiply(RotationX(math.Pi / 2))
	if !tuples.Point(15, 0, 7).Equal(matrices.MultiplyTuple(tt, p)) {
		t.Fatal("chained transform failed")
	}
}
