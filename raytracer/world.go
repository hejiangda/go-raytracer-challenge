package raytracer

import (
	"math"
	"sort"
)

var WorldId = 0

const DefaultReflectRemaining = 5 //5
const DefaultRefractRemaining = 5

type World struct {
	// 目前只实现了点光源和球
	// 就先这样吧
	Id      int
	Lights  []*PointLight
	Objects []Shape
}

type PreComputations struct {
	T          float64
	Object     Shape
	Point      *Tuple
	OverPoint  *Tuple
	EyeV       *Tuple
	NormalV    *Tuple
	Inside     bool
	ReflectV   *Tuple
	N1         float64
	N2         float64
	UnderPoint *Tuple
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

func PrepareComputations(i Intersection, ray *Ray, xs []Intersection) (comps PreComputations) {
	comps.T = i.T
	comps.Object = i.Obj

	comps.Point = ray.Position(comps.T)
	comps.EyeV = ray.Direction.Multiply(-1).Normalize()
	comps.NormalV = comps.Object.NormalAt(comps.Point)
	if Dot(comps.NormalV, comps.EyeV) < 0 {
		comps.Inside = true
		comps.NormalV = comps.NormalV.Multiply(-1)
	}

	comps.OverPoint = comps.Point.Add(comps.NormalV.Multiply(Eps))
	comps.UnderPoint = comps.Point.Add(comps.NormalV.Multiply(-Eps))
	comps.ReflectV = Reflect(ray.Direction, comps.NormalV)

	containers := make([]Shape, 0)
	for _, x := range xs {
		if x == i {
			if len(containers) == 0 {
				comps.N1 = 1
			} else {
				comps.N1 = (containers[len(containers)-1]).GetMaterial().RefractiveIndex
			}
		}
		checkContain := func() (idx int, ok bool) {
			for j, container := range containers {
				if container == x.Obj {
					return j, true
				}
			}
			return -1, false
		}
		removeContain := func(idx int) {
			containers = append(containers[:idx], containers[idx+1:]...)
		}
		if idx, ok := checkContain(); ok == true {
			removeContain(idx)
		} else {
			containers = append(containers, x.Obj)
		}
		if x == i {
			if len(containers) == 0 {
				comps.N2 = 1.0
			} else {
				comps.N2 = containers[len(containers)-1].GetMaterial().RefractiveIndex
			}
		}
	}
	return
}

func (w *World) ShadeHit(comps PreComputations, remaining int) (color *Tuple) {

	color = Color(0, 0, 0)
	for _, light := range w.Lights {
		surface := Lighting(comps.Object.GetMaterial(), comps.Object, light, comps.OverPoint, comps.EyeV, comps.NormalV, w.IsShadowed(comps.OverPoint, light))
		color = color.Add(surface)
		reflected := w.ReflectedColor(comps, remaining)
		refracted := w.RefractedColor(comps, remaining)
		material := comps.Object.GetMaterial()
		if material.Reflective > 0 && material.Transparency > 0 {
			reflectance := Schlick(comps)
			color = color.Add(reflected.Multiply(reflectance)).Add(refracted.Multiply(1 - reflectance))
		} else {
			color = color.Add(reflected).Add(refracted)
		}
	}
	return
}
func (w *World) ColorAt(ray *Ray, remaining int) (color *Tuple) {
	color = Color(0, 0, 0)
	inters := w.Intersect(ray)

	for _, inter := range inters {
		if inter.T > 0 {
			comps := PrepareComputations(inter, ray, inters)
			return w.ShadeHit(comps, remaining)
		}
	}
	return
}
func (w *World) IsShadowed(point *Tuple, light *PointLight) bool {
	v := light.Position.Subtract(point)
	distance := v.Magnitude()
	direction := v.Normalize()

	r := NewRay(point, direction)
	intersections := w.Intersect(r)

	h, err := Hit(intersections)
	if err == nil && h.T < distance {
		return true
	}
	return false
}
func (w *World) ReflectedColor(comps PreComputations, remaining int) (color *Tuple) {
	if remaining < 1 {
		return Color(0, 0, 0)
	}

	if AlmostEqual(comps.Object.GetMaterial().Reflective, 0, Eps) {
		color = Color(0, 0, 0)
	} else {
		reflectRay := NewRay(comps.OverPoint, comps.ReflectV)
		color = w.ColorAt(reflectRay, remaining-1).Multiply(comps.Object.GetMaterial().Reflective)
	}

	return color
}
func (w *World) RefractedColor(comps PreComputations, remaining int) (color *Tuple) {
	if remaining < 1 {
		return Color(0, 0, 0)
	}
	if comps.Object.GetMaterial().Transparency == 0 {
		return Color(0, 0, 0)
	} else {
		nRatio := comps.N1 / comps.N2
		cosI := Dot(comps.EyeV, comps.NormalV)
		sin2T := nRatio * nRatio * (1 - cosI*cosI)
		if sin2T >= 1 {
			return Color(0, 0, 0)
		}

		cosT := math.Sqrt(1.0 - sin2T)
		direction := comps.NormalV.Multiply(nRatio*cosI - cosT).Subtract(comps.EyeV.Multiply(nRatio))
		//fmt.Println(direction)
		refractRay := NewRay(comps.UnderPoint, direction)

		color = w.ColorAt(refractRay, remaining-1).Multiply(comps.Object.GetMaterial().Transparency)
		return color
	}
}
