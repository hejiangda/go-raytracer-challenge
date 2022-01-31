package tuples

import "math"

type Tuple struct {
	X float64
	Y float64
	Z float64
	W float64
}
type TupleType int

const Eps = 1e-5
const (
	VectorType TupleType = iota
	PointType
)

func (t *Tuple) Type() TupleType {
	if AlmostEqual(t.W, 1.0, Eps) {
		return PointType
	} else {
		return VectorType
	}
}
func (t *Tuple) Equal(a *Tuple) bool {
	return AlmostEqual(a.X, t.X, Eps) &&
		AlmostEqual(a.Y, t.Y, Eps) &&
		AlmostEqual(a.Z, t.Z, Eps) &&
		AlmostEqual(a.W, t.W, Eps)
}
func AlmostEqual(a, b, eps float64) bool {
	return math.Abs(a-b) < eps
}

func Point(x, y, z float64) *Tuple {
	t := new(Tuple)
	t.X = x
	t.Y = y
	t.Z = z
	t.W = 1.0
	return t
}

func Vector(x, y, z float64) *Tuple {
	t := new(Tuple)
	t.X = x
	t.Y = y
	t.Z = z
	t.W = 0.0
	return t
}
func Add(a1, a2 *Tuple) *Tuple {
	t := new(Tuple)
	t.X = a1.X + a2.X
	t.Y = a1.Y + a2.Y
	t.Z = a1.Z + a2.Z
	t.W = a1.W + a2.W
	return t
}

func Subtract(a1, a2 *Tuple) *Tuple {
	t := new(Tuple)
	t.X = a1.X - a2.X
	t.Y = a1.Y - a2.Y
	t.Z = a1.Z - a2.Z
	t.W = a1.W - a2.W
	return t
}

func Negate(a *Tuple) *Tuple {
	t := new(Tuple)
	t.X = -a.X
	t.Y = -a.Y
	t.Z = -a.Z
	t.W = -a.W
	return t
}
func (t *Tuple) Multiply(val float64) *Tuple {
	a := new(Tuple)
	a.X = t.X * val
	a.Y = t.Y * val
	a.Z = t.Z * val
	a.W = t.W * val
	return a
}
func (t *Tuple) Divide(val float64) *Tuple {
	a := new(Tuple)
	a.X = t.X / val
	a.Y = t.Y / val
	a.Z = t.Z / val
	a.W = t.W / val
	return a
}
func (t *Tuple) Magnitude() float64 {
	return math.Sqrt(t.X*t.X + t.Y*t.Y + t.Z*t.Z + t.W*t.W)
}
func Normalize(t *Tuple) *Tuple {
	a := new(Tuple)
	m := t.Magnitude()
	a.X = t.X / m
	a.Y = t.Y / m
	a.Z = t.Z / m
	a.W = t.W / m
	return a
}
func Dot(a, b *Tuple) float64 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z + a.W*b.W
}
func Cross(a, b *Tuple) *Tuple {
	return Vector(a.Y*b.Z-a.Z*b.Y,
		a.Z*b.X-a.X*b.Z,
		a.X*b.Y-a.Y*b.X)
}
func Color(red, green, blue float64) *Tuple {
	c := new(Tuple)
	c.X = red
	c.Y = green
	c.Z = blue
	c.W = 0
	return c
}
func MultiplyColors(a, b *Tuple) *Tuple {
	c := new(Tuple)
	c.X = a.X * b.X
	c.Y = a.Y * b.Y
	c.Z = a.Z * b.Z
	c.W = a.W * b.W
	return c
}
