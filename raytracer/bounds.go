package raytracer

import "math"

type Bounds struct {
	Min *Tuple
	Max *Tuple
}

func (b *Bounds) Trans(t *Matrix) {
	b.Min = MultiplyTuple(t, b.Min)
	b.Max = MultiplyTuple(t, b.Max)
}
func (b *Bounds) Update(nb *Bounds) {
	b.Min.X = math.Min(nb.Min.X, b.Min.X)
	b.Min.Y = math.Min(nb.Min.Y, b.Min.Y)
	b.Min.Z = math.Min(nb.Min.Z, b.Min.Z)

	b.Max.X = math.Max(nb.Max.X, b.Max.X)
	b.Max.Y = math.Max(nb.Max.Y, b.Max.Y)
	b.Max.Z = math.Max(nb.Max.Z, b.Max.Z)
}
func NewBoundsShape(shape Shape) *Bounds {
	switch t := shape.(type) {
	case *Sphere:
		return &Bounds{Point(-1, -1, -1), Point(1, 1, 1)}
	case *Cube:
		return &Bounds{Point(-1, -1, -1), Point(1, 1, 1)}
	case *Cylinder:
		return &Bounds{Point(-1, t.Minimum, -1), Point(1, t.Maximum, 1)}
	case *Cone:
		return &Bounds{Point(-1, t.Minimum, -1), Point(1, t.Maximum, 1)}
	case *Plane:
		return &Bounds{Point(math.Inf(-1), 0, math.Inf(-1)), Point(math.Inf(1), 0, math.Inf(1))}
	}
	return nil
}

func NewBoundsGroup(group *Group) *Bounds {
	var groupBounds *Bounds
	for i, child := range group.Children {
		var objBounds *Bounds
		switch t := child.(type) {
		case *Group:
			objBounds = NewBoundsGroup(t)
			objBounds.Trans(child.GetTransform())

		default:
			objBounds = NewBoundsShape(t)
			objBounds.Trans(child.GetTransform())
		}
		if i == 0 {
			groupBounds = objBounds
		} else {
			groupBounds.Update(objBounds)
		}
	}
	groupBounds.Trans(group.GetTransform())
	return groupBounds
}
func (b *Bounds) checkAxis(axis int, origin, direction float64) (tmin, tmax float64) {
	var tminNumerator float64
	var tmaxNumerator float64
	switch axis {
	case 0:
		tminNumerator = b.Min.X - origin
		tmaxNumerator = b.Max.X - origin
	case 1:
		tminNumerator = b.Min.Y - origin
		tmaxNumerator = b.Max.Y - origin
	case 2:
		tminNumerator = b.Min.Z - origin
		tmaxNumerator = b.Max.Z - origin
	}

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
func (b *Bounds) LocalIntersect(ray *Ray) bool {
	xtmin, xtmax := checkAxis(ray.Origin.X, ray.Direction.X)
	ytmin, ytmax := checkAxis(ray.Origin.Y, ray.Direction.Y)
	ztmin, ztmax := checkAxis(ray.Origin.Z, ray.Direction.Z)

	tmin := math.Max(math.Max(xtmin, ytmin), ztmin)
	tmax := math.Min(math.Min(xtmax, ytmax), ztmax)

	if tmin > tmax {
		return false
	}
	return true
}
func (b *Bounds) Intersect(r *Ray) bool {
	return b.LocalIntersect(r)
}
