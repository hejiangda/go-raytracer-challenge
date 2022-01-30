package tuples

import "math"

type tuple struct {
	X float64
	Y float64
	Z float64
	W float64
}
type TupleType int
type TuplePtr *tuple

const Eps = 1e-5
const (
	VectorType TupleType = iota
	PointType
)

func (t *tuple) Type() TupleType {
	if AlmostEqual(t.W, 1.0, Eps) {
		return PointType
	} else {
		return VectorType
	}
}
func (t *tuple) Equal(a *tuple) bool {
	return AlmostEqual(a.X, t.X, Eps) &&
		AlmostEqual(a.Y, t.Y, Eps) &&
		AlmostEqual(a.Z, t.Z, Eps) &&
		AlmostEqual(a.W, t.W, Eps)
}
func AlmostEqual(a, b, eps float64) bool {
	return math.Abs(a-b) < eps
}

func Tuple(x, y, z, w float64) *tuple {
	t := new(tuple)
	t.X = x
	t.Y = y
	t.Z = z
	t.W = w
	return t
}
func Point(x, y, z float64) *tuple {
	t := new(tuple)
	t.X = x
	t.Y = y
	t.Z = z
	t.W = 1.0
	return t
}

func Vector(x, y, z float64) *tuple {
	t := new(tuple)
	t.X = x
	t.Y = y
	t.Z = z
	t.W = 0.0
	return t
}
func Add(a1, a2 *tuple) *tuple {
	t := new(tuple)
	t.X = a1.X + a2.X
	t.Y = a1.Y + a2.Y
	t.Z = a1.Z + a2.Z
	t.W = a1.W + a2.W
	return t
}

func Subtract(a1, a2 *tuple) *tuple {
	t := new(tuple)
	t.X = a1.X - a2.X
	t.Y = a1.Y - a2.Y
	t.Z = a1.Z - a2.Z
	t.W = a1.W - a2.W
	return t
}

func Negate(a *tuple) *tuple {
	t := new(tuple)
	t.X = -a.X
	t.Y = -a.Y
	t.Z = -a.Z
	t.W = -a.W
	return t
}
func (t *tuple) Multiply(val float64) *tuple {
	a := new(tuple)
	a.X = t.X * val
	a.Y = t.Y * val
	a.Z = t.Z * val
	a.W = t.W * val
	return a
}
func (t *tuple) Divide(val float64) *tuple {
	a := new(tuple)
	a.X = t.X / val
	a.Y = t.Y / val
	a.Z = t.Z / val
	a.W = t.W / val
	return a
}
func (t *tuple) Magnitude() float64 {
	return math.Sqrt(t.X*t.X + t.Y*t.Y + t.Z*t.Z + t.W*t.W)
}
func Normalize(t *tuple) *tuple {
	a := new(tuple)
	m := t.Magnitude()
	a.X = t.X / m
	a.Y = t.Y / m
	a.Z = t.Z / m
	a.W = t.W / m
	return a
}
func Dot(a, b *tuple) float64 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z + a.W*b.W
}
func Cross(a, b *tuple) *tuple {
	return Vector(a.Y*b.Z-a.Z*b.Y,
		a.Z*b.X-a.X*b.Z,
		a.X*b.Y-a.Y*b.X)
}
