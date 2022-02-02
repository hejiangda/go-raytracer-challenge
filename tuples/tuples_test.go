package tuples

import (
	"math"
	"testing"
)

func TestATupleIsAPoint(t *testing.T) {
	a := &Tuple{4.3, -4.2, 3.1, 1.0}
	if !AlmostEqual(a.X, 4.3, Eps) {
		t.Fatal("a.X != 4.3")
	}
	if !AlmostEqual(a.Y, -4.2, Eps) {
		t.Fatal("a.Y != -4.2")
	}
	if !AlmostEqual(a.Z, 3.1, Eps) {
		t.Fatal("a.Z != 3.1")
	}
	if !AlmostEqual(a.W, 1.0, Eps) {
		t.Fatal("a.W != 1.0")
	}
	if a.Type() != PointType {
		t.Fatal("a is not a point")
	}
	if a.Type() == VectorType {
		t.Fatal("a is a vector")
	}
	if !a.Equal(&Tuple{4.3, -4.2, 3.1, 1.0}) {
		t.Fatal("a is not equal to itself")
	}
}

func TestATupleIsAVector(t *testing.T) {
	a := &Tuple{4.3, -4.2, 3.1, 0.0}
	if !AlmostEqual(a.X, 4.3, Eps) {
		t.Fatal("a.X != 4.3")
	}
	if !AlmostEqual(a.Y, -4.2, Eps) {
		t.Fatal("a.Y != -4.2")
	}
	if !AlmostEqual(a.Z, 3.1, Eps) {
		t.Fatal("a.Z != 3.1")
	}
	if !AlmostEqual(a.W, 0.0, Eps) {
		t.Fatal("a.W != 0.0")
	}
	if a.Type() == PointType {
		t.Fatal("a is a point")
	}
	if a.Type() != VectorType {
		t.Fatal("a is not a vector")
	}
	if !a.Equal(&Tuple{4.3, -4.2, 3.1, 0.0}) {
		t.Fatal("a is not equal to itself")
	}
}

func TestCreatePoint(t *testing.T) {
	p := Point(4, -4, 3)
	if !p.Equal(&Tuple{4, -4, 3, 1}) {
		t.Fatal("p is not equal to equivalent tuple")
	}
}

func TestCreateVector(t *testing.T) {
	v := Vector(4, -4, 3)
	if !v.Equal(&Tuple{4, -4, 3, 0}) {
		t.Fatal("v is not equal to equivalent tuple")
	}
}

func TestAdd(t *testing.T) {
	a1 := &Tuple{3, -2, 5, 1}
	a2 := &Tuple{-2, 3, 1, 0}
	if !Add(a1, a2).Equal(&Tuple{1, 1, 6, 1}) {
		t.Fatal("a1 add a2 is not equal to tuple(1,1,6,1)")
	}
}

func TestSubtract(t *testing.T) {
	p1 := Point(3, 2, 1)
	p2 := Point(5, 6, 7)
	if !Subtract(p1, p2).Equal(Vector(-2, -4, -6)) {
		t.Fatal("p1 subtract p2 is not equal to vector(-2,-4,-6)")
	}
}
func TestSubtractAVectorFromAPoint(t *testing.T) {
	p := Point(3, 2, 1)
	v := Vector(5, 6, 7)
	if !Subtract(p, v).Equal(Point(-2, -4, -6)) {
		t.Fatal("p subtract v is not equal to point(-2,-4,-6)")
	}
}

func TestSubtractTwoVectors(t *testing.T) {
	v1 := Vector(3, 2, 1)
	v2 := Vector(5, 6, 7)
	if !Subtract(v1, v2).Equal(Vector(-2, -4, -6)) {
		t.Fatal("v1 subtract v2 is not equal to vector(-2,-4,-6)")
	}
}

func TestSubtractAVectorFromZero(t *testing.T) {
	zero := Vector(0, 0, 0)
	v := Vector(1, -2, 3)
	if !Subtract(zero, v).Equal(Vector(-1, 2, -3)) {
		t.Fatal("zero subtract v is not equal to vector(-1,2,-3)")
	}
}

func TestNegate(t *testing.T) {
	a := &Tuple{1, -2, 3, -4}
	na := Negate(a)
	if !na.Equal(&Tuple{-1, 2, -3, 4}) {
		t.Fatal("negated a is not equal to tuple(-1,2,-3,4)")
	}
}

func TestMultiplyATupleByAScalar(t *testing.T) {
	a := &Tuple{1, -2, 3, -4}
	b := a.Multiply(3.5)
	if !b.Equal(&Tuple{3.5, -7, 10.5, -14}) {
		t.Fatal("a multiply by 3.5 is not equal to tuple(3.5,-7,10.5,-14)")
	}
}

func TestMultiplyATupleByAFraction(t *testing.T) {
	a := &Tuple{1, -2, 3, -4}
	b := a.Multiply(0.5)
	if !b.Equal(&Tuple{0.5, -1, 1.5, -2}) {
		t.Fatal("a multiply by 0.5 is not equal to tuple(0.5,-1,1.5,-2)")
	}
}

func TestDivideATupleByAScalar(t *testing.T) {
	a := &Tuple{1, -2, 3, -4}
	b := a.Divide(2)
	if !b.Equal(&Tuple{0.5, -1, 1.5, -2}) {
		t.Fatal("a divide by 2 is not equal to tuple(0.5,-1,1.5,-2)")
	}
}

func TestMagnitudeOfVector1(t *testing.T) {
	v := Vector(1, 0, 0)
	m := v.Magnitude()
	if !AlmostEqual(m, 1, Eps) {
		t.Fatal("magnitude of vector(1,0,0) is not equal to 1")
	}
}
func TestMagnitudeOfVector2(t *testing.T) {
	v := Vector(0, 1, 0)
	m := v.Magnitude()
	if !AlmostEqual(m, 1, Eps) {
		t.Fatal("magnitude of vector(1,0,0) is not equal to 1")
	}
}
func TestMagnitudeOfVector3(t *testing.T) {
	v := Vector(0, 0, 1)
	m := v.Magnitude()
	if !AlmostEqual(m, 1, Eps) {
		t.Fatal("magnitude of vector(1,0,0) is not equal to 1")
	}
}

func TestMagnitudeOfVector4(t *testing.T) {
	v := Vector(1, 2, 3)
	m := v.Magnitude()
	if !AlmostEqual(m, math.Sqrt(14), Eps) {
		t.Fatal("magnitude of vector(1,2,3) is not equal to sqrt(14)")
	}
}

func TestNormalizeVector1(t *testing.T) {
	v := Vector(4, 0, 0)
	v1 := Normalize(v)
	if !v1.Equal(Vector(1, 0, 0)) {
		t.Fatal("normalize(v) is not equal to vector(1,0,0)")
	}
}
func TestNormalizeVector2(t *testing.T) {
	v := Vector(1, 2, 3)
	v1 := Normalize(v)
	if !v1.Equal(Vector(0.26726, 0.53452, 0.80178)) {
		t.Fatal("normalize(v) is not equal to vector(0.26726,0.53452,0.80178)")
	}
}

func TestMagnitudeOfANormalizedVector(t *testing.T) {
	v := Vector(1, 2, 3)
	norm := Normalize(v)
	m := norm.Magnitude()
	if !AlmostEqual(m, 1, Eps) {
		t.Fatal("The magnitude of a normalized vector is not equal to 1")
	}
}

func TestDotProduct(t *testing.T) {
	a := Vector(1, 2, 3)
	b := Vector(2, 3, 4)
	res := Dot(a, b)
	if !AlmostEqual(res, 20, Eps) {
		t.Fatal("The dot product of two tuples doesn't equal to 20")
	}
}

func TestCrossProduct(t *testing.T) {
	a := Vector(1, 2, 3)
	b := Vector(2, 3, 4)
	c1 := Cross(a, b)
	c2 := Cross(b, a)
	if !c1.Equal(Vector(-1, 2, -1)) {
		t.Fatal("The cross product of two vectors is not equal vector(-1,2,-1)")
	}
	if !c2.Equal(Vector(1, -2, 1)) {
		t.Fatal("The cross product of two vectors is not equal vector(1,-2,1)")
	}
}

func TestColorsTuples(t *testing.T) {
	c := Color(-0.5, 0.4, 1.7)
	if !c.Equal(&Tuple{-0.5, 0.4, 1.7, 0}) {
		t.Fatal("The color is not equal to color(-0.5, 0.4, 1.7)")
	}
}

func TestAddColors(t *testing.T) {
	c1 := Color(0.9, 0.6, 0.75)
	c2 := Color(0.7, 0.1, 0.25)
	c3 := Add(c1, c2)
	if !c3.Equal(Color(1.6, 0.7, 1.0)) {
		t.Fatal("c1 + c2 is not equal to color(1.6,0.7,1.0)")
	}
}
func TestSubtractColor(t *testing.T) {
	c1 := Color(0.9, 0.6, 0.75)
	c2 := Color(0.7, 0.1, 0.25)
	c3 := Subtract(c1, c2)
	if !c3.Equal(Color(0.2, 0.5, 0.5)) {
		t.Fatal("c1 - c2 is not equal to color(0.2,0.5,0.5")
	}
}

func TestMultiplyColorByScalar(t *testing.T) {
	c := Color(0.2, 0.3, 0.4)
	c2 := c.Multiply(2)
	if !c2.Equal(Color(0.4, 0.6, 0.8)) {
		t.Fatal("c * 2 is not equal to color(0.4,0.6,0.8)")
	}
}

func TestMultiplyColors(t *testing.T) {
	c1 := Color(1, 0.2, 0.4)
	c2 := Color(0.9, 1, 0.1)
	c3 := MultiplyColors(c1, c2)
	if !c3.Equal(Color(0.9, 0.2, 0.04)) {
		t.Fatal("c1 * c2 is not equal to color(0.9,0.2,0.04)")
	}

}

func TestReflectAVector45(t *testing.T) {
	v := Vector(1, -1, 0)
	n := Vector(0, 1, 0)
	r := Reflect(v, n)
	if !Vector(1, 1, 0).Equal(r) {
		t.Fatal("r != vector(1, 1, 0)")
	}
}

func TestReflectAVectorSlantedSurface(t *testing.T) {
	v := Vector(0, -1, 0)
	n := Vector(math.Sqrt2/2, math.Sqrt2/2, 0)
	r := Reflect(v, n)
	if !Vector(1, 0, 0).Equal(r) {
		t.Fatal("r != vector(1, 0, 0)")
	}
}
