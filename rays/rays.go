package rays

import (
	"matrices"
	"tuples"
)

type Ray struct {
	Origin    *tuples.Tuple
	Direction *tuples.Tuple
}

func NewRay(origin, direction *tuples.Tuple) *Ray {
	return &Ray{origin, direction}
}
func Position(r *Ray, t float64) *tuples.Tuple {
	return tuples.Add(r.Origin, r.Direction.Multiply(t))
}
func Transform(r *Ray, m *matrices.Matrix) *Ray {
	return NewRay(matrices.MultiplyTuple(m, r.Origin), matrices.MultiplyTuple(m, r.Direction))
}
