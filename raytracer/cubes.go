package raytracer

import (
	"math"
)

var CubeId = 0

// Cube 边长为2，各坐标绝对值为1
type Cube struct {
	Id        int
	Transform *Matrix
	Material  *Material
	// 仅用于测试
	//localRay    *Ray
	//localNormal *Tuple
}

func NewCube() *Cube {
	CubeId++
	return &Cube{
		Id:        CubeId,
		Transform: EyeMatrix(4),
		Material:  NewMaterial(),
	}
}
func (c *Cube) GetTransform() *Matrix {
	return c.Transform
}

func (c *Cube) SetTransform(matrix *Matrix) {
	c.Transform = matrix
}

func (c *Cube) GetMaterial() *Material {
	return c.Material
}

func (c *Cube) SetMaterial(material *Material) {
	c.Material = material
}

func (c *Cube) Intersect(r *Ray) (ret []Intersection) {
	inv, err := Inverse(c.Transform)
	if err != nil {
		panic(err)
	}
	// 世界坐标转换到物体坐标
	localRay := r.Transform(inv)
	return c.LocalIntersect(localRay)
}

func (c *Cube) NormalAt(worldPoint *Tuple) *Tuple {
	inv, err := Inverse(c.Transform)
	if err != nil {
		panic(err)
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

func checkAxis(origin, direction float64) (tmin, tmax float64) {
	tminNumerator := -1 - origin
	tmaxNumerator := 1 - origin
	if math.Abs(direction) >= Eps {
		tmin = tminNumerator / direction
		tmax = tmaxNumerator / direction
	} else {
		tmin = tminNumerator * math.Inf(1)
		tmax = tmaxNumerator * math.Inf(1)
	}
	if tmin > tmax {
		tmin, tmax = tmax, tmin
	}
	return tmin, tmax
}
func (c *Cube) LocalIntersect(ray *Ray) (xs []Intersection) {
	xtmin, xtmax := checkAxis(ray.Origin.X, ray.Direction.X)
	ytmin, ytmax := checkAxis(ray.Origin.Y, ray.Direction.Y)
	ztmin, ztmax := checkAxis(ray.Origin.Z, ray.Direction.Z)

	tmin := math.Max(math.Max(xtmin, ytmin), ztmin)
	tmax := math.Min(math.Min(xtmax, ytmax), ztmax)

	if tmin > tmax {
		return
	}

	xs = append(xs, Intersection{tmin, c})
	xs = append(xs, Intersection{tmax, c})
	return
}
func (c *Cube) localNormalAt(point *Tuple) *Tuple {
	maxC := math.Max(math.Max(math.Abs(point.X), math.Abs(point.Y)), math.Abs(point.Z))
	switch maxC {
	case math.Abs(point.X):
		return Vector(point.X, 0, 0)
	case math.Abs(point.Y):
		return Vector(0, point.Y, 0)
	case math.Abs(point.Z):
		return Vector(0, 0, point.Z)
	}
	return Vector(0, 0, 0)
}
