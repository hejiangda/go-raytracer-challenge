package raytracer

import (
	"errors"
	"math"
	"sort"
)

type Intersection struct {
	T   float64
	Obj Shape
}

//type PhysicalObject interface {
//	Intersect(r *Ray) (ret []float64)
//	NormalAt(p *Tuple) *Tuple
//	GetMaterial() *Material
//}

func Intersections(args ...Intersection) (ret []Intersection) {
	for _, arg := range args {
		ret = append(ret, arg)
	}
	return
}

func Hit(xs []Intersection) (ret Intersection, err error) {
	sort.Slice(xs, func(i, j int) bool {
		return xs[i].T < xs[j].T
	})
	for _, x := range xs {
		if x.T >= 0 {
			ret = x
			return ret, nil
		}
	}
	return ret, errors.New("no hit")
}

func Schlick(comps PreComputations) float64 {
	cos := Dot(comps.EyeV, comps.NormalV)
	if comps.N1 > comps.N2 {
		n := comps.N1 / comps.N2
		sin2T := n * n * (1.0 - cos*cos)
		if sin2T >= 1.0 {
			return 1.0
		}
		cosT := math.Sqrt(1.0 - sin2T)
		cos = cosT
	}

	r0 := ((comps.N1 - comps.N2) / (comps.N1 + comps.N2)) * ((comps.N1 - comps.N2) / (comps.N1 + comps.N2))

	return r0 + (1-r0)*math.Pow(1-cos, 5)
}
