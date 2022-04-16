package raytracer

type Ray struct {
	Origin    *Tuple
	Direction *Tuple
}

func NewRay(origin, direction *Tuple) *Ray {
	return &Ray{origin, direction}
}
func (r *Ray) Position(t float64) *Tuple {
	return Add(r.Origin, r.Direction.Multiply(t))
}
func (r *Ray) Transform(m *Matrix) *Ray {
	return NewRay(MultiplyTuple(m, r.Origin), MultiplyTuple(m, r.Direction))
}
