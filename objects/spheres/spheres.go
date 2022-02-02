package spheres

import (
	"lights"
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
	Material  *lights.Material
}

func NewSphere() *Sphere {
	SphereId++
	return &Sphere{
		Id:        SphereId,
		Transform: matrices.EyeMatrix(4),
		Material:  lights.NewMaterial(),
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
	r = rays.Transform(r, inv)
	sphereToRay := tuples.Subtract(r.Origin, tuples.Point(0, 0, 0))
	a := tuples.Dot(r.Direction, r.Direction)
	b := 2 * tuples.Dot(r.Direction, sphereToRay)
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
func (s *Sphere) NormalAt(p *tuples.Tuple) *tuples.Tuple {
	return NormalAt(s, p)
}
func (s *Sphere) GetMaterial() *lights.Material {
	return s.Material
}
func NormalAt(s *Sphere, p *tuples.Tuple) *tuples.Tuple {
	inv, err := matrices.Inverse(s.Transform)
	if err != nil {
		log.Fatal(err)
	}
	// 世界坐标转换到物体坐标
	objectPoint := matrices.MultiplyTuple(inv, p)
	// 计算物体坐标系法向
	objectNormal := tuples.Subtract(objectPoint, tuples.Point(0, 0, 0))
	tinv := matrices.Transpose(inv)
	// 物体坐标系法向转换到世界坐标系下的法向
	worldNormal := matrices.MultiplyTuple(tinv, objectNormal)
	worldNormal.W = 0
	return tuples.Normalize(worldNormal)
}
