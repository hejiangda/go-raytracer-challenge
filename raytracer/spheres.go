package raytracer

import (
	"log"
	"math"
)

var SphereId = 0

type Sphere struct {
	Id        int
	Transform *Matrix
	Material  *Material
}

func NewSphere() *Sphere {
	SphereId++
	return &Sphere{
		Id:        SphereId,
		Transform: EyeMatrix(4),
		Material:  NewMaterial(),
	}
}
func (s *Sphere) Intersect(r *Ray) (ret []float64) {
	//return Intersect(s, r)
	inv, err := Inverse(s.Transform)
	if err != nil {
		log.Fatal(err)
	}
	r = Transform(r, inv)
	sphereToRay := Subtract(r.Origin, Point(0, 0, 0))
	a := Dot(r.Direction, r.Direction)
	b := 2 * Dot(r.Direction, sphereToRay)
	c := Dot(sphereToRay, sphereToRay) - 1
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

//func Intersect(s *Sphere, r *Ray) (ret []float64) {
//
//}
func (s *Sphere) SetTransform(t *Matrix) {
	s.Transform = t
}
func SetTransform(s *Sphere, t *Matrix) {
	s.Transform = t
}
func (s *Sphere) NormalAt(p *Tuple) *Tuple {
	return NormalAt(s, p)
}
func (s *Sphere) GetMaterial() *Material {
	return s.Material
}
func NormalAt(s *Sphere, p *Tuple) *Tuple {
	inv, err := Inverse(s.Transform)
	if err != nil {
		log.Fatal(err)
	}
	// 世界坐标转换到物体坐标
	objectPoint := MultiplyTuple(inv, p)
	// 计算物体坐标系法向
	objectNormal := Subtract(objectPoint, Point(0, 0, 0))
	tinv := Transpose(inv)
	// 物体坐标系法向转换到世界坐标系下的法向
	worldNormal := MultiplyTuple(tinv, objectNormal)
	worldNormal.W = 0
	return Normalize(worldNormal)
}
