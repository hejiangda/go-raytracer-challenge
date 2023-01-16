package raytracer

import "math"

var (
	BLACK = &Tuple{0, 0, 0, 0}
	WHITE = &Tuple{1, 1, 1, 0}
)

type Pattern interface {
	AtShape(obj Shape, point *Tuple) (color *Tuple)
	SetTransform(m *Matrix)
}
type StripePattern struct {
	A         *Tuple
	B         *Tuple
	transform *Matrix
}

func NewStripePattern(a, b *Tuple) *StripePattern {
	return &StripePattern{a, b, EyeMatrix(4)}
}

func (p *StripePattern) At(point *Tuple) (color *Tuple) {
	if int(math.Floor(point.X))%2 == 0 {
		color = p.A
	} else {
		color = p.B
	}
	return color
}

func (p *StripePattern) AtShape(obj Shape, point *Tuple) (color *Tuple) {
	objPoint := obj.World2Object(point)
	patTransInv, _ := Inverse(p.transform)
	patPoint := MultiplyTuple(patTransInv, objPoint)
	return p.At(patPoint)
}

func (p *StripePattern) SetTransform(m *Matrix) {
	p.transform = m
}

type GradientPattern struct {
	A         *Tuple
	B         *Tuple
	transform *Matrix
}

func NewGradientPattern(a, b *Tuple) *GradientPattern {
	return &GradientPattern{a, b, EyeMatrix(4)}
}

func (p *GradientPattern) AtShape(obj Shape, point *Tuple) (color *Tuple) {
	objPoint := obj.World2Object(point)

	patTransInv, _ := Inverse(p.transform)
	patPoint := MultiplyTuple(patTransInv, objPoint)

	return p.At(patPoint)
}
func (p *GradientPattern) At(point *Tuple) (color *Tuple) {
	distance := p.B.Subtract(p.A)
	fraction := point.X - math.Floor(point.X)
	return p.A.Add(distance.Multiply(fraction))
}
func (p *GradientPattern) SetTransform(m *Matrix) {
	p.transform = m
}

type RingPattern struct {
	A         *Tuple
	B         *Tuple
	transform *Matrix
}

func NewRingPattern(a, b *Tuple) *RingPattern {
	return &RingPattern{a, b, EyeMatrix(4)}
}
func (p *RingPattern) At(point *Tuple) (color *Tuple) {
	if int(math.Floor(math.Sqrt(point.X*point.X+point.Z*point.Z)))%2 == 0 {
		return p.A
	} else {
		return p.B
	}
}
func (p *RingPattern) AtShape(obj Shape, point *Tuple) (color *Tuple) {
	objPoint := obj.World2Object(point)
	patTransInv, _ := Inverse(p.transform)
	patPoint := MultiplyTuple(patTransInv, objPoint)
	return p.At(patPoint)
}
func (p *RingPattern) SetTransform(m *Matrix) {
	p.transform = m
}

type CheckersPattern struct {
	A         *Tuple
	B         *Tuple
	transform *Matrix
}

func NewCheckersPattern(a, b *Tuple) *CheckersPattern {
	return &CheckersPattern{a, b, EyeMatrix(4)}
}

func (p *CheckersPattern) At(point *Tuple) (color *Tuple) {
	if int(math.Floor(point.X)+math.Floor(point.Y)+math.Floor(point.Z))%2 == 0 {
		return p.A
	} else {
		return p.B
	}
}
func (p *CheckersPattern) AtShape(obj Shape, point *Tuple) (color *Tuple) {
	objPoint := obj.World2Object(point)
	patTransInv, _ := Inverse(p.transform)
	patPoint := MultiplyTuple(patTransInv, objPoint)
	return p.At(patPoint)
}
func (p *CheckersPattern) SetTransform(m *Matrix) {
	p.transform = m
}
