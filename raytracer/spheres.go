package raytracer

import (
	"errors"
	"math"
)

var SphereId = 0

type Sphere struct {
	Id        int
	Transform *Matrix
	Material  *Material
	// 仅用于测试
	localRay *Ray
	Parent   Shape
}

func (s *Sphere) GetParent() Shape {
	return s.Parent
}

func (s *Sphere) SetParent(shape Shape) {
	s.Parent = shape
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
	inv, err := Inverse(s.Transform)
	if err != nil {
		panic(err)
	}
	// 世界坐标转换到物体坐标
	localRay := r.Transform(inv)
	s.localRay = localRay
	t1, t2, err := s.localIntersect(localRay)
	if err != nil {
		return
	}
	return Intersections(Intersection{t1, s}, Intersection{t2, s})
}
func (s *Sphere) localIntersect(localRay *Ray) (float64, float64, error) {
	sphereToRay := Subtract(localRay.Origin, Point(0, 0, 0))
	a := Dot(localRay.Direction, localRay.Direction)
	b := 2 * Dot(localRay.Direction, sphereToRay)
	c := Dot(sphereToRay, sphereToRay) - 1
	discriminant := b*b - 4*a*c
	if discriminant < 0 {
		return 0, 0, errors.New("b^2-4ac<0")
	}
	t1 := (-b - math.Sqrt(discriminant)) / (2 * a)
	t2 := (-b + math.Sqrt(discriminant)) / (2 * a)
	if t1 > t2 {
		return t2, t1, nil
	} else {
		return t1, t2, nil
	}
}

func (s *Sphere) GetTransform() *Matrix {
	return s.Transform
}
func (s *Sphere) SetTransform(t *Matrix) {
	s.Transform = t
}

func (s *Sphere) NormalAt(worldPoint *Tuple) *Tuple {

	// 世界坐标转换到物体坐标
	objectPoint := s.World2Object(worldPoint)

	// 获取物体坐标系下的法向
	localNormal := s.localNormalAt(objectPoint)

	// 物体坐标系法向转换到世界坐标系下的法向

	worldNormal := s.Normal2World(localNormal)
	worldNormal.W = 0
	return Normalize(worldNormal)
}

func (s *Sphere) GetMaterial() *Material {
	return s.Material
}
func (s *Sphere) SetMaterial(m *Material) {
	s.Material = m
}
func (s *Sphere) localNormalAt(objectPoint *Tuple) *Tuple {
	// 计算物体坐标系法向
	localNormal := Subtract(objectPoint, Point(0, 0, 0))
	return localNormal
}
func (s *Sphere) World2Object(p *Tuple) *Tuple {
	if s.Parent != nil {
		p = s.Parent.World2Object(p)
	}
	inv, err := Inverse(s.Transform)
	if err != nil {
		panic(err)
	}
	localPoint := MultiplyTuple(inv, p)
	return localPoint
}
func GlassSphere() (s *Sphere) {
	s = NewSphere()
	s.Material.Transparency = 1.0
	s.Material.RefractiveIndex = 1.5
	return s
}
func (s *Sphere) Normal2World(t *Tuple) *Tuple {
	inv, err := Inverse(s.Transform)
	if err != nil {
		panic(err)
	}
	normal := MultiplyTuple(Transpose(inv), t)
	normal.W = 0
	normal = normal.Normalize()
	if s.Parent != nil {
		normal = s.Parent.Normal2World(normal)
	}
	return normal
}
