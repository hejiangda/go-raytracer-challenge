package spheres

import (
	"log"
	"math"
	"matrices"
	"rays"
	"tuples"
)

var SphereId = 0

type Sphere struct {
	Id        int
	Transform *matrices.Matrix
}

func NewSphere() *Sphere {
	SphereId++
	return &Sphere{
		Id:        SphereId,
		Transform: matrices.EyeMatrix(4),
	}
}
func (s *Sphere) Intersect(r *rays.Ray) (ret []float64) {
	return Intersect(s, r)
}
func Intersect(s *Sphere, r *rays.Ray) (ret []float64) {
	inv, err := matrices.Inverse(s.Transform)
	if err != nil {
		log.Fatal(err)
	}
	r1 := rays.Transform(r, inv)
	sphereToRay := tuples.Subtract(r1.Origin, tuples.Point(0, 0, 0))
	a := tuples.Dot(r1.Direction, r1.Direction)
	b := 2 * tuples.Dot(r1.Direction, sphereToRay)
	c := tuples.Dot(sphereToRay, sphereToRay) - 1
	discriminant := b*b - 4*a*c
	if discriminant < 0 {
		return
	}
	t1 := (-b - math.Sqrt(discriminant)) / (2 * a)
	t2 := (-b + math.Sqrt(discriminant)) / (2 * a)
	if t1 > t2 {
		ret = append(ret, t2)
		ret = append(ret, t1)
	} else {
		ret = append(ret, t1)
		ret = append(ret, t2)
	}
	return
}
func (s *Sphere) SetTransform(t *matrices.Matrix) {
	s.Transform = t
}
func SetTransform(s *Sphere, t *matrices.Matrix) {
	s.Transform = t
}
