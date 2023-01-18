package raytracer

import (
	"errors"
	"math"
	"sort"
)

type Group struct {
	Id        int
	Transform *Matrix
	Material  *Material
	Children  []Shape

	Parent Shape
}

var GroupId = 0

func NewGroup() *Group {
	GroupId++
	return &Group{
		Id:        GroupId,
		Transform: EyeMatrix(4),
		Material:  NewMaterial(),
	}
}

func (g *Group) AddChild(s Shape) {
	g.Children = append(g.Children, s)
	s.SetParent(g)
}

func (g *Group) HasChild(s Shape) bool {
	for _, child := range g.Children {
		if child == s {
			return true
		}
	}
	return false
}

func (g *Group) GetParent() Shape {
	return g.Parent
}

func (g *Group) SetParent(s Shape) {
	g.Parent = s
}

func (g *Group) GetTransform() *Matrix {
	return g.Transform
}

func (g *Group) SetTransform(matrix *Matrix) {
	g.Transform = matrix
}

func (g *Group) GetMaterial() *Material {
	//TODO implement me
	panic("implement me")
}

func (g *Group) SetMaterial(material *Material) {
	//TODO implement me
	panic("implement me")
}

func (g *Group) Intersect(r *Ray) (ret []Intersection) {
	b := NewBoundsGroup(g)
	if b != nil {
		if !b.Intersect(r) {
			return ret
		}
	}

	inv, _ := Inverse(g.GetTransform())
	for _, child := range g.Children {
		child.SetTransform(g.GetTransform().Multiply(child.GetTransform()))
		ret = append(ret, child.Intersect(r)...)
		child.SetTransform(inv.Multiply(child.GetTransform()))
	}
	sort.Slice(ret, func(i, j int) bool {
		return ret[i].T < ret[j].T
	})
	return ret
}

func (g *Group) NormalAt(p *Tuple) *Tuple {
	//TODO implement me
	panic("implement me")
}

func (g *Group) localIntersect(localRay *Ray) (float64, float64, error) {
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
func (g *Group) World2Object(p *Tuple) *Tuple {
	if g.Parent != nil {
		p = g.Parent.World2Object(p)
	}
	inv, err := Inverse(g.Transform)
	if err != nil {
		panic(err)
	}
	localPoint := MultiplyTuple(inv, p)
	return localPoint
}
func (g *Group) Normal2World(t *Tuple) *Tuple {
	inv, err := Inverse(g.Transform)
	if err != nil {
		panic(err)
	}
	normal := MultiplyTuple(Transpose(inv), t)
	normal.W = 0
	normal = normal.Normalize()
	if g.Parent != nil {
		normal = g.Parent.Normal2World(normal)
	}
	return normal
}
