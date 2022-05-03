package raytracer

import (
	"log"
	"math"
)

var PlaneId = 0

// Plane x-z plane
type Plane struct {
	Id          int
	Transform   *Matrix
	Material    *Material
	localRay    *Ray
	localNormal *Tuple
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
	arr := p.localIntersect(r)
	for _, t := range arr {
		ret = append(ret, Intersection{t, p})
	}
	return
}
func (p *Plane) localIntersect(r *Ray) (ret []float64) {
	inv, err := Inverse(p.Transform)
	if err != nil {
		log.Fatal(err)
	}
	localRay := r.Transform(inv)
	if math.Abs(localRay.Direction.Y) < Eps {
		return
	}
	t := -localRay.Origin.Y / localRay.Direction.Y
	ret = append(ret, t)
	return
}

func (p *Plane) NormalAt(pos *Tuple) *Tuple {
	return p.localNormalAt(pos)
}
func (p *Plane) localNormalAt(pos *Tuple) *Tuple {
	inv, err := Inverse(p.Transform)
	if err != nil {
		log.Fatal(err)
	}
	// 世界坐标转换到物体坐标
	// 计算物体坐标系法向
	localNormal := Vector(0, 1, 0)

	tinv := Transpose(inv)
	// 物体坐标系法向转换到世界坐标系下的法向
	worldNormal := MultiplyTuple(tinv, localNormal)
	worldNormal.W = 0
	return Normalize(worldNormal)
}
