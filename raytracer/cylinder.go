package raytracer

import (
	"math"
)

var CylinderId = 0

// Cylinder radius=1
type Cylinder struct {
	Id        int
	Transform *Matrix
	Material  *Material
	// 仅用于测试
	localRay *Ray

	Minimum float64
	Maximum float64
	Closed  bool

	Parent Shape
}

func (c *Cylinder) GetParent() Shape {
	return c.Parent
}

func (c *Cylinder) SetParent(s Shape) {
	c.Parent = s
}

func NewCylinder() *Cylinder {
	CylinderId++
	return &Cylinder{
		Id:        CylinderId,
		Transform: EyeMatrix(4),
		Material:  NewMaterial(),
		Minimum:   math.Inf(-1),
		Maximum:   math.Inf(1),
	}
}

func (c *Cylinder) GetTransform() *Matrix {
	return c.Transform
}

func (c *Cylinder) SetTransform(matrix *Matrix) {
	c.Transform = matrix
}

func (c *Cylinder) GetMaterial() *Material {
	return c.Material
}

func (c *Cylinder) SetMaterial(material *Material) {
	c.Material = material
}

func (c *Cylinder) Intersect(r *Ray) (ret []Intersection) {
	inv, err := Inverse(c.Transform)
	if err != nil {
		panic(err)
	}
	// 世界坐标转换到物体坐标
	localRay := r.Transform(inv)
	return c.localIntersect(localRay)
}

func (c *Cylinder) NormalAt(worldPoint *Tuple) *Tuple {
	//// 世界坐标转换到物体坐标
	//objectPoint := MultiplyTuple(inv, worldPoint)
	objectPoint := c.World2Object(worldPoint)

	// 获取物体坐标系下的法向
	localNormal := c.localNormalAt(objectPoint)

	// 物体坐标系法向转换到世界坐标系下的法向
	worldNormal := c.Normal2World(localNormal)
	worldNormal.W = 0
	return Normalize(worldNormal)
}

func (c *Cylinder) localIntersect(localRay *Ray) (ret []Intersection) {
	a := localRay.Direction.X*localRay.Direction.X + localRay.Direction.Z*localRay.Direction.Z
	if AlmostEqual(a, 0.0, Eps) {
		ret = c.intersectCaps(localRay, ret)
		return
	}
	b := 2*localRay.Origin.X*localRay.Direction.X + 2*localRay.Origin.Z*localRay.Direction.Z
	cc := localRay.Origin.X*localRay.Origin.X + localRay.Origin.Z*localRay.Origin.Z - 1
	disc := b*b - 4*a*cc
	if disc < 0 {
		return
	}
	t0 := (-b - math.Sqrt(disc)) / (2 * a)
	t1 := (-b + math.Sqrt(disc)) / (2 * a)
	if t0 > t1 {
		t0, t1 = t1, t0
	}
	y0 := localRay.Origin.Y + t0*localRay.Direction.Y
	if c.Minimum < y0 && y0 < c.Maximum {
		ret = append(ret, Intersection{t0, c})
	}
	y1 := localRay.Origin.Y + t1*localRay.Direction.Y
	if c.Minimum < y1 && y1 < c.Maximum {
		ret = append(ret, Intersection{t1, c})
	}
	ret = c.intersectCaps(localRay, ret)
	return
}
func (c *Cylinder) localNormalAt(point *Tuple) *Tuple {
	dist := point.X*point.X + point.Z*point.Z
	if dist < 1 && point.Y >= c.Maximum-Eps {
		return Vector(0, 1, 0)
	} else if dist < 1 && point.Y <= c.Minimum+Eps {
		return Vector(0, -1, 0)
	}
	return Vector(point.X, 0, point.Z)
}

func (c *Cylinder) checkCap(ray *Ray, t float64) bool {
	x := ray.Origin.X + t*ray.Direction.X
	z := ray.Origin.Z + t*ray.Direction.Z
	return (x*x + z*z) <= 1
}

func (c *Cylinder) intersectCaps(ray *Ray, xs []Intersection) []Intersection {
	if !c.Closed || AlmostEqual(ray.Direction.Y, 0, Eps) {
		return xs
	}
	t := (c.Minimum - ray.Origin.Y) / ray.Direction.Y
	if c.checkCap(ray, t) {
		xs = append(xs, Intersection{t, c})
	}
	t = (c.Maximum - ray.Origin.Y) / ray.Direction.Y
	if c.checkCap(ray, t) {
		xs = append(xs, Intersection{t, c})
	}
	return xs
}
func (c *Cylinder) World2Object(p *Tuple) *Tuple {
	if c.Parent != nil {
		p = c.Parent.World2Object(p)
	}
	inv, err := Inverse(c.Transform)
	if err != nil {
		panic(err)
	}
	localPoint := MultiplyTuple(inv, p)
	return localPoint
}
func (c *Cylinder) Normal2World(t *Tuple) *Tuple {
	inv, err := Inverse(c.Transform)
	if err != nil {
		panic(err)
	}
	normal := MultiplyTuple(Transpose(inv), t)
	normal.W = 0
	normal = normal.Normalize()
	if c.Parent != nil {
		normal = c.Parent.Normal2World(normal)
	}
	return normal
}
