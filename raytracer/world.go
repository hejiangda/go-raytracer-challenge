package raytracer

import (
	"sort"
)

var WorldId = 0

type World struct {
	// 目前只实现了点光源和球
	// 就先这样吧
	Id      int
	Lights  []*PointLight
	Objects []Shape
}

type PreComputations struct {
	T         float64
	Object    Shape
	Point     *Tuple
	OverPoint *Tuple
	EyeV      *Tuple
	NormalV   *Tuple
	Inside    bool
}

func NewWorld() (w *World) {
	WorldId++
	return &World{
		Id:      WorldId,
		Lights:  nil,
		Objects: nil,
	}
}
func DefaultWorld() (w *World) {

	w = NewWorld()
	w.Lights = append(w.Lights, NewPointLight(Point(-10, 10, -10), Color(1, 1, 1)))

	s1 := NewSphere()
	s1.Material.Color = Color(0.8, 1.0, 0.6)
	s1.Material.Diffuse = 0.7
	s1.Material.Specular = 0.2

	s2 := NewSphere()
	s2.Transform = Scaling(0.5, 0.5, 0.5)
	w.Objects = append(w.Objects, s1, s2)
	return
}

func (w *World) Intersect(ray *Ray) (ret []Intersection) {
	for _, object := range w.Objects {
		ret = append(ret, object.Intersect(ray)...)
	}
	sort.Slice(ret, func(i, j int) bool {
		return ret[i].T < ret[j].T
	})
	return
}

func PrepareComputations(i Intersection, ray *Ray) (comps PreComputations) {
	comps.T = i.T
	comps.Object = i.Obj

	comps.Point = ray.Position(comps.T)
	comps.EyeV = ray.Direction.Multiply(-1)
	comps.NormalV = comps.Object.NormalAt(comps.Point)
	if Dot(comps.NormalV, comps.EyeV) < 0 {
		comps.Inside = true
		comps.NormalV = comps.NormalV.Multiply(-1)
	}

	comps.OverPoint = comps.Point.Add(comps.NormalV.Multiply(Eps))
	return
}

func (w *World) ShadeHit(comps PreComputations) (color *Tuple) {
	color = Color(0, 0, 0)
	for _, light := range w.Lights {
		color = color.Add(Lighting(comps.Object.GetMaterial(), comps.Object, light, comps.Point, comps.EyeV, comps.NormalV, w.IsShadowed(comps.OverPoint)))
	}
	return
}
func (w *World) ColorAt(ray *Ray) (color *Tuple) {
	color = Color(0, 0, 0)
	inters := w.Intersect(ray)
	for _, inter := range inters {
		if inter.T > 0 {
			comps := PrepareComputations(inter, ray)
			return w.ShadeHit(comps)
		}
	}
	return
}
func (w *World) IsShadowed(point *Tuple) bool {
	for _, light := range w.Lights {
		v := light.Position.Subtract(point)
		distance := v.Magnitude()
		direction := v.Normalize()

		r := NewRay(point, direction)
		intersections := w.Intersect(r)

		h, err := Hit(intersections)
		if err == nil && h.T < distance {
			return true
		}
	}
	return false
}
