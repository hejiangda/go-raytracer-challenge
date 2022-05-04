package raytracer

import (
	"log"
	"math"
)

var PlaneId = 0

// Plane x-z plane
type Plane struct {
	Id        int
	Transform *Matrix
	Material  *Material
}

func NewPlane() *Plane {
	PlaneId++
	return &Plane{
		Id:        PlaneId,
		Transform: EyeMatrix(4),
		Material:  NewMaterial(),
	}
}
func (p *Plane) GetTransform() *Matrix {
	return p.Transform
}
func (p *Plane) SetTransform(m *Matrix) {
	p.Transform = m
}

func (p *Plane) GetMaterial() *Material {
	return p.Material
}

func (p *Plane) SetMaterial(m *Material) {
	p.Material = m
}

func (p *Plane) Intersect(r *Ray) (ret []Intersection) {
	inv, err := Inverse(p.Transform)
	if err != nil {
		panic(err)
	}
	// 世界坐标转换到物体坐标
	localRay := r.Transform(inv)
	t := p.localIntersect(localRay)
	ret = append(ret, Intersection{t, p})
	return
}
func (p *Plane) localIntersect(localRay *Ray) (t float64) {
	if math.Abs(localRay.Direction.Y) < Eps {
		return
	}
	t = -localRay.Origin.Y / localRay.Direction.Y
	return
}

func (p *Plane) NormalAt(worldPoint *Tuple) *Tuple {
	inv, err := Inverse(p.Transform)
	if err != nil {
		log.Fatal(err)
	}
	// 世界坐标转换到物体坐标
	objectPoint := MultiplyTuple(inv, worldPoint)

	// 获取物体坐标系下的法向
	localNormal := p.localNormalAt(objectPoint)

	// 物体坐标系法向转换到世界坐标系下的法向
	worldNormal := MultiplyTuple(Transpose(inv), localNormal)
	worldNormal.W = 0
	return Normalize(worldNormal)
}
func (p *Plane) localNormalAt(pos *Tuple) *Tuple {
	// 计算物体坐标系法向
	localNormal := Vector(0, 1, 0)
	return localNormal
}
