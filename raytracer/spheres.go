package raytracer

import (
	"log"
	"math"
)

var SphereId = 0

type Sphere struct {
	Id          int
	Transform   *Matrix
	Material    *Material
	localRay    *Ray
	localNormal *Tuple
}

func NewSphere() *Sphere {
	SphereId++
	return &Sphere{
		Id:        SphereId,
		Transform: EyeMatrix(4),
		Material:  NewMaterial(),
	}
}

func (s *Sphere) Intersect(r *Ray) (ret []Intersection) {
	arr := s.localIntersect(r)
	for _, t := range arr {
		ret = append(ret, Intersection{t, s})
	}
	return
}
func (s *Sphere) localIntersect(r *Ray) (ret []float64) {
	inv, err := Inverse(s.Transform)
	if err != nil {
		log.Fatal(err)
	}
	localRay := r.Transform(inv)
	s.localRay = localRay
	sphereToRay := Subtract(localRay.Origin, Point(0, 0, 0))
	a := Dot(localRay.Direction, localRay.Direction)
	b := 2 * Dot(localRay.Direction, sphereToRay)
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

func (s *Sphere) GetTransform() *Matrix {
	return s.Transform
}
func (s *Sphere) SetTransform(t *Matrix) {
	s.Transform = t
}

func (s *Sphere) NormalAt(p *Tuple) *Tuple {
	return s.localNormalAt(p)
}

func (s *Sphere) GetMaterial() *Material {
	return s.Material
}
func (s *Sphere) SetMaterial(m *Material) {
	s.Material = m
}
func (s *Sphere) localNormalAt(p *Tuple) *Tuple {
	inv, err := Inverse(s.Transform)
	if err != nil {
		log.Fatal(err)
	}
	// 世界坐标转换到物体坐标
	// localPoint
	objectPoint := MultiplyTuple(inv, p)
	// 计算物体坐标系法向
	localNormal := Subtract(objectPoint, Point(0, 0, 0))
	s.localNormal = localNormal
	tinv := Transpose(inv)
	// 物体坐标系法向转换到世界坐标系下的法向
	worldNormal := MultiplyTuple(tinv, localNormal)
	worldNormal.W = 0
	return Normalize(worldNormal)
}

func GlassSphere() (s *Sphere) {
	s = NewSphere()
	s.Material.Transparency = 1.0
	s.Material.RefractiveIndex = 1.5
	return s
}
