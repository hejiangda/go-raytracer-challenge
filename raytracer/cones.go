package raytracer

import (
	"log"
	"math"
)

var ConeId = 0

type Cone struct {
	Id        int
	Transform *Matrix
	Material  *Material
	// 仅用于测试
	localRay *Ray

	Minimum float64
	Maximum float64
	Closed  bool
}

func (c *Cone) GetTransform() *Matrix {
	return c.Transform
}

func (c *Cone) SetTransform(matrix *Matrix) {
	c.Transform = matrix
}

func (c *Cone) GetMaterial() *Material {
	return c.Material
}

func (c *Cone) SetMaterial(material *Material) {
	c.Material = material
}

func (c *Cone) Intersect(r *Ray) (ret []Intersection) {
	inv, err := Inverse(c.Transform)
	if err != nil {
		panic(err)
	}
	// 世界坐标转换到物体坐标
	localRay := r.Transform(inv)
	return c.localIntersect(localRay)
}

func (c *Cone) NormalAt(worldPoint *Tuple) *Tuple {
	inv, err := Inverse(c.Transform)
	if err != nil {
		log.Fatal(err)
	}
	// 世界坐标转换到物体坐标
	objectPoint := MultiplyTuple(inv, worldPoint)

	// 获取物体坐标系下的法向
	localNormal := c.localNormalAt(objectPoint)

	// 物体坐标系法向转换到世界坐标系下的法向
	worldNormal := MultiplyTuple(Transpose(inv), localNormal)
	worldNormal.W = 0
	return Normalize(worldNormal)
}

func NewCone() *Cone {
	ConeId++
	return &Cone{
		Id:        ConeId,
		Transform: EyeMatrix(4),
		Material:  NewMaterial(),
		Minimum:   math.Inf(-1),
		Maximum:   math.Inf(1),
	}
}

func (c *Cone) localIntersect(localRay *Ray) (ret []Intersection) {
	a := localRay.Direction.X*localRay.Direction.X - localRay.Direction.Y*localRay.Direction.Y + localRay.Direction.Z*localRay.Direction.Z
	b := 2*localRay.Origin.X*localRay.Direction.X - 2*localRay.Origin.Y*localRay.Direction.Y + 2*localRay.Origin.Z*localRay.Direction.Z

	cc := localRay.Origin.X*localRay.Origin.X - localRay.Origin.Y*localRay.Origin.Y + localRay.Origin.Z*localRay.Origin.Z
	if AlmostEqual(a, 0.0, Eps) && !AlmostEqual(b, 0.0, Eps) {
		t := -cc / (2 * b)
		ret = append(ret, Intersection{t, c})
		return c.intersectCaps(localRay, ret)
	}
	if AlmostEqual(a, 0.0, Eps) && AlmostEqual(b, 0.0, Eps) {
		ret = c.intersectCaps(localRay, ret)
		return
	}
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

func (c *Cone) intersectCaps(ray *Ray, xs []Intersection) []Intersection {
	if !c.Closed || AlmostEqual(ray.Direction.Y, 0, Eps) {
		return xs
	}
	t := (c.Minimum - ray.Origin.Y) / ray.Direction.Y
	if c.checkCap(ray, t, c.Minimum) {
		xs = append(xs, Intersection{t, c})
	}
	t = (c.Maximum - ray.Origin.Y) / ray.Direction.Y
	if c.checkCap(ray, t, c.Maximum) {
		xs = append(xs, Intersection{t, c})
	}
	return xs
}

func (c *Cone) checkCap(ray *Ray, t, y float64) bool {
	x := ray.Origin.X + t*ray.Direction.X
	z := ray.Origin.Z + t*ray.Direction.Z

	return (x*x + z*z) <= y*y
}

func (c *Cone) localNormalAt(point *Tuple) *Tuple {
	dist := point.X*point.X + point.Z*point.Z
	if dist < 1 && point.Y >= c.Maximum-Eps {
		return Vector(0, 1, 0)
	} else if dist < 1 && point.Y <= c.Minimum+Eps {
		return Vector(0, -1, 0)
	}
	y := math.Sqrt(dist)
	if point.Y > 0 {
		y = -y
	}
	return Vector(point.X, y, point.Z)
}
