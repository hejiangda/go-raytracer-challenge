package raytracer

type Ray struct {
	Origin    *Tuple
	Direction *Tuple
}

func NewRay(origin, direction *Tuple) *Ray {
	return &Ray{origin, direction}
}
func Position(r *Ray, t float64) *Tuple {
	return Add(r.Origin, r.Direction.Multiply(t))
}
func Transform(r *Ray, m *Matrix) *Ray {
	return NewRay(MultiplyTuple(m, r.Origin), MultiplyTuple(m, r.Direction))
}
