package raytracer

import (
	"math"
)

type Triangle struct {
	P1     *Tuple
	P2     *Tuple
	P3     *Tuple
	E1     *Tuple
	E2     *Tuple
	Normal *Tuple

	Transform *Matrix
	Material  *Material
	Parent    Shape
}

func (t *Triangle) GetTransform() *Matrix {
	return t.Transform
}

func (t *Triangle) SetTransform(matrix *Matrix) {
	t.Transform = matrix
}

func (t *Triangle) GetMaterial() *Material {
	return t.Material
}

func (t *Triangle) SetMaterial(material *Material) {
	t.Material = material
}

func (t *Triangle) Intersect(r *Ray) (ret []Intersection) {
	inv, err := Inverse(t.Transform)
	if err != nil {
		panic(err)
	}
	// 世界坐标转换到物体坐标
	localRay := r.Transform(inv)
	return t.LocalIntersect(localRay)
}

func (t *Triangle) NormalAt(p *Tuple) *Tuple {

	// 世界坐标转换到物体坐标
	objectPoint := t.World2Object(p)

	// 获取物体坐标系下的法向
	localNormal := t.LocalNormalAt(objectPoint)

	// 物体坐标系法向转换到世界坐标系下的法向

	worldNormal := t.Normal2World(localNormal)
	worldNormal.W = 0
	return Normalize(worldNormal)
}

func (t *Triangle) GetParent() Shape {
	return t.Parent
}

func (t *Triangle) SetParent(s Shape) {
	t.Parent = s
}

func (t *Triangle) World2Object(p *Tuple) *Tuple {
	if t.Parent != nil {
		p = t.Parent.World2Object(p)
	}
	inv, err := Inverse(t.Transform)
	if err != nil {
		panic(err)
	}
	localPoint := MultiplyTuple(inv, p)
	return localPoint
}

func (t *Triangle) Normal2World(p *Tuple) *Tuple {
	inv, err := Inverse(t.Transform)
	if err != nil {
		panic(err)
	}
	normal := MultiplyTuple(Transpose(inv), p)
	normal.W = 0
	normal = normal.Normalize()
	if t.Parent != nil {
		normal = t.Parent.Normal2World(normal)
	}
	return normal
}

func NewTriangle(p1, p2, p3 *Tuple) (t *Triangle) {
	e1 := p2.Subtract(p1)
	e2 := p3.Subtract(p1)
	normal := Normalize(Cross(e2, e1))
	t = &Triangle{
		P1:        p1,
		P2:        p2,
		P3:        p3,
		E1:        e1,
		E2:        e2,
		Normal:    normal,
		Transform: EyeMatrix(4),
		Material:  NewMaterial(),
	}
	return
}

func (t *Triangle) LocalNormalAt(p *Tuple) *Tuple {
	return t.Normal
}

func (t *Triangle) LocalIntersect(ray *Ray) (ret []Intersection) {
	dirCrossE2 := Cross(ray.Direction, t.E2)
	det := Dot(t.E1, dirCrossE2)
	if math.Abs(det) < Eps {
		return
	}

	f := 1.0 / det
	p1ToOrigin := ray.Origin.Subtract(t.P1)
	u := f * Dot(p1ToOrigin, dirCrossE2)
	if u < 0 || u > 1 {
		return
	}

	originCrossE1 := Cross(p1ToOrigin, t.E1)
	v := f * Dot(ray.Direction, originCrossE1)
	if v < 0 || (v+u) > 1 {
		return
	}

	tt := f * Dot(t.E2, originCrossE1)

	ret = append(ret, Intersection{tt, t})

	return
}
