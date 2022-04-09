package raytracer

import (
	"math"
	"testing"
)

func TestTranslation(t *testing.T) {
	trans := Translation(5, -3, 2)
	p := Point(-3, 4, 5)
	p1 := MultiplyTuple(trans, p)
	if !p1.Equal(Point(2, 1, 7)) {
		t.Fatal("translate failed")
	}
}

func TestInverseOfATranslation(t *testing.T) {
	trans := Translation(5, -3, 2)
	inv, err := Inverse(trans)
	if err != nil {
		t.Fatal(err)
	}
	p := Point(-3, 4, 5)
	if !MultiplyTuple(inv, p).Equal(Point(-8, 7, 3)) {
		t.Fatal("multiply by the inverse of a transformation matrix failed")
	}
}
func TestTranslationDoesNotAffectVector(t *testing.T) {
	trans := Translation(5, -3, 2)
	v := Vector(-3, 4, 5)
	if !v.Equal(MultiplyTuple(trans, v)) {
		t.Fatal("translation affect the vector")
	}
}
func TestScaling(t *testing.T) {
	trans := Scaling(2, 3, 4)
	p := Point(-4, 6, 8)
	if !Point(-8, 18, 32).Equal(MultiplyTuple(trans, p)) {
		t.Fatal("scaling failed")
	}
}

func TestScaling2(t *testing.T) {
	trans := Scaling(2, 3, 4)
	v := Vector(-4, 6, 8)
	if !Vector(-8, 18, 32).Equal(MultiplyTuple(trans, v)) {
		t.Fatal("scaling failed")
	}
}

func TestInverseOfScaling(t *testing.T) {
	trans := Scaling(2, 3, 4)
	inv, err := Inverse(trans)
	if err != nil {
		t.Fatal(err)
	}
	v := Vector(-4, 6, 8)
	res := MultiplyTuple(inv, v)
	if !Vector(-2, 2, 2).Equal(res) {
		t.Fatal("scaling failed,v:", res)
	}
}

func TestReflection(t *testing.T) {
	trans := Scaling(-1, 1, 1)
	p := Point(2, 3, 4)
	if !Point(-2, 3, 4).Equal(MultiplyTuple(trans, p)) {
		t.Fatal("reflect failed")
	}
}

func TestRotatingPointAroundXAxis(t *testing.T) {
	p := Point(0, 1, 0)
	half_quarter := RotationX(math.Pi / 4)
	full_quarter := RotationX(math.Pi / 2)
	p1 := MultiplyTuple(half_quarter, p)
	p2 := MultiplyTuple(full_quarter, p)
	if !Point(0, math.Sqrt2/2, math.Sqrt2/2).Equal(p1) {
		t.Fatal("rotate failed")
	}
	if !Point(0, 0, 1).Equal(p2) {
		t.Fatal("rotate failed")
	}

}
func TestInverseRotatingAroundXAxis(t *testing.T) {
	p := Point(0, 1, 0)
	half_quarter := RotationX(math.Pi / 4)
	inv, err := Inverse(half_quarter)
	if err != nil {
		t.Fatal(err)
	}
	if !Point(0, math.Sqrt2/2, -math.Sqrt2/2).Equal(MultiplyTuple(inv, p)) {
		t.Fatal("inverse rotate failed")
	}
}
func TestRotatingPointAroundYAxis(t *testing.T) {
	p := Point(0, 0, 1)
	half_quarter := RotationY(math.Pi / 4)
	full_quarter := RotationY(math.Pi / 2)
	p1 := MultiplyTuple(half_quarter, p)
	p2 := MultiplyTuple(full_quarter, p)
	if !Point(math.Sqrt2/2, 0, math.Sqrt2/2).Equal(p1) {
		t.Fatal("rotate failed")
	}
	if !Point(1, 0, 0).Equal(p2) {
		t.Fatal("rotate failed")
	}
}
func TestRotatingPointAroundZAxis(t *testing.T) {
	p := Point(0, 1, 0)
	half_quarter := RotationZ(math.Pi / 4)
	full_quarter := RotationZ(math.Pi / 2)
	p1 := MultiplyTuple(half_quarter, p)
	p2 := MultiplyTuple(full_quarter, p)
	if !Point(-math.Sqrt2/2, math.Sqrt2/2, 0).Equal(p1) {
		t.Fatal("rotate failed")
	}
	if !Point(-1, 0, 0).Equal(p2) {
		t.Fatal("rotate failed")
	}
}
func TestShearingXY(t *testing.T) {
	trans := Shearing(1, 0, 0, 0, 0, 0)
	p := Point(2, 3, 4)
	if !Point(5, 3, 4).Equal(MultiplyTuple(trans, p)) {
		t.Fatal("shearing xy failed")
	}
}

func TestShearingXZ(t *testing.T) {
	trans := Shearing(0, 1, 0, 0, 0, 0)
	p := Point(2, 3, 4)
	if !Point(6, 3, 4).Equal(MultiplyTuple(trans, p)) {
		t.Fatal("shearing xy failed")
	}
}

func TestShearingyx(t *testing.T) {
	trans := Shearing(0, 0, 1, 0, 0, 0)
	p := Point(2, 3, 4)
	if !Point(2, 5, 4).Equal(MultiplyTuple(trans, p)) {
		t.Fatal("shearing xy failed")
	}
}
func TestShearingyz(t *testing.T) {
	trans := Shearing(0, 0, 0, 1, 0, 0)
	p := Point(2, 3, 4)
	if !Point(2, 7, 4).Equal(MultiplyTuple(trans, p)) {
		t.Fatal("shearing xy failed")
	}
}
func TestShearingzx(t *testing.T) {
	trans := Shearing(0, 0, 0, 0, 1, 0)
	p := Point(2, 3, 4)
	if !Point(2, 3, 6).Equal(MultiplyTuple(trans, p)) {
		t.Fatal("shearing xy failed")
	}
}
func TestShearingzy(t *testing.T) {
	trans := Shearing(0, 0, 0, 0, 0, 1)
	p := Point(2, 3, 4)
	if !Point(2, 3, 7).Equal(MultiplyTuple(trans, p)) {
		t.Fatal("shearing xy failed")
	}
}

func TestIndividualTransformations(t *testing.T) {
	p := Point(1, 0, 1)
	a := RotationX(math.Pi / 2)
	b := Scaling(5, 5, 5)
	c := Translation(10, 5, 7)
	p2 := MultiplyTuple(a, p)
	if !Point(1, -1, 0).Equal(p2) {
		t.Fatal("rotate failed")
	}
	p3 := MultiplyTuple(b, p2)
	if !Point(5, -5, 0).Equal(p3) {
		t.Fatal("scale failed")
	}
	p4 := MultiplyTuple(c, p3)
	if !Point(15, 0, 7).Equal(p4) {
		t.Fatal("translate failed")
	}
}
func TestChainedTransformations(t *testing.T) {
	p := Point(1, 0, 1)
	a := RotationX(math.Pi / 2)
	b := Scaling(5, 5, 5)
	c := Translation(10, 5, 7)
	tt := c.Multiply(b).Multiply(a)
	if !Point(15, 0, 7).Equal(MultiplyTuple(tt, p)) {
		t.Fatal("chained transform failed")
	}
}
func TestFluentAPIs(t *testing.T) {
	p := Point(1, 0, 1)
	tt := EyeMatrix(4).Multiply(Translation(10, 5, 7)).Multiply(Scaling(5, 5, 5)).Multiply(RotationX(math.Pi / 2))
	if !Point(15, 0, 7).Equal(MultiplyTuple(tt, p)) {
		t.Fatal("chained transform failed")
	}
}

func TestViewTransform(t *testing.T) {
	from := Point(0, 0, 0)
	to := Point(0, 0, -1)
	up := Vector(0, 1, 0)
	tt := ViewTransform(from, to, up)
	if !IsSame(tt, EyeMatrix(4)) {
		t.Fatal("tt is not equal to I")
	}
}

func TestViewTransform2(t *testing.T) {
	from := Point(0, 0, 0)
	to := Point(0, 0, 1)
	up := Vector(0, 1, 0)
	tt := ViewTransform(from, to, up)
	if !IsSame(tt, Scaling(-1, 1, -1)) {
		t.Fatal("tt is not equal to Scaling(-1,1,-1)")
	}
}

func TestViewTransform3(t *testing.T) {
	from := Point(0, 0, 8)
	to := Point(0, 0, 0)
	up := Vector(0, 1, 0)
	tt := ViewTransform(from, to, up)
	if !IsSame(tt, Translation(0, 0, -8)) {
		t.Fatal("tt is not equal to Translation(0,0,-8),tt:", tt, "Translation(0, 0, -8):", Translation(0, 0, -8))
	}
}

func TestViewTransform4(t *testing.T) {
	from := Point(1, 3, 2)
	to := Point(4, -2, 8)
	up := Vector(1, 1, 0)
	tt := ViewTransform(from, to, up)
	m := NewMatrix(4,
		[][]float64{
			{-0.50709, 0.50709, 0.67612, -2.36643},
			{0.76772, 0.60609, 0.12122, -2.82843},
			{-0.35857, 0.59761, -0.71714, 0.00000},
			{0.00000, 0.00000, 0.00000, 1.00000},
		})
	if !IsSame(tt, m) {
		t.Fatal("tt is not equal to m")
	}
}
