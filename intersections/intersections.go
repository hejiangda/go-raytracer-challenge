package intersections

import (
	"errors"
	"rays"
	"sort"
)

type Intersection struct {
	T   float64
	Obj PhysicalObject
}

type PhysicalObject interface {
	Intersect(r *rays.Ray) (ret []float64)
	//SetTransform(t *matrices.Matrix)
}

func Intersections(args ...Intersection) (ret []Intersection) {
	for _, arg := range args {
		ret = append(ret, arg)
	}
	return
}

func Intersect(obj PhysicalObject, r *rays.Ray) (ret []Intersection) {
	arr := obj.Intersect(r)
	for _, t := range arr {
		ret = append(ret, Intersection{t, obj})
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
