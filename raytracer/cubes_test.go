package raytracer

import "testing"

func TestCube_LocalIntersect(t *testing.T) {
	tests := map[string]struct {
		origin    *Tuple
		direction *Tuple
		t1        float64
		t2        float64
	}{
		"plusX":  {Point(5, 0.5, 0), Vector(-1, 0, 0), 4, 6},
		"minusX": {Point(-5, 0.5, 0), Vector(1, 0, 0), 4, 6},
		"plusY":  {Point(0.5, 5, 0), Vector(0, -1, 0), 4, 6},
		"minusY": {Point(0.5, -5, 0), Vector(0, 1, 0), 4, 6},
		"plusZ":  {Point(0.5, 0, 5), Vector(0, 0, -1), 4, 6},
		"minusZ": {Point(0.5, 0, -5), Vector(0, 0, 1), 4, 6},
		"inside": {Point(0, 0.5, 0), Vector(0, 0, 1), -1, 1},
	}

	c := NewCube()
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			r := NewRay(tc.origin, tc.direction)
			xs := c.LocalIntersect(r)
			if len(xs) != 2 {
				t.Fatal("len(xs)!=2")
			}
			if !AlmostEqual(xs[0].T, tc.t1, Eps) || !AlmostEqual(xs[1].T, tc.t2, Eps) {
				t.Fatal(xs)
			}
		})
	}
}

func TestARayMissesACube(t *testing.T) {
	c := NewCube()
	tests := map[string]struct {
		origin    *Tuple
		direction *Tuple
	}{
		"1": {Point(-2, 0, 0), Vector(0.2673, 0.5345, 0.8018)},
		"2": {Point(0, -2, 0), Vector(0.8018, 0.2673, 0.5345)},
		"3": {Point(0, 0, -2), Vector(0.5345, 0.8018, 0.2673)},
		"4": {Point(2, 0, 2), Vector(0, 0, -1)},
		"5": {Point(0, 2, 2), Vector(0, -1, 0)},
		"6": {Point(2, 2, 0), Vector(-1, 0, 0)},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			r := NewRay(tc.origin, tc.direction)
			xs := c.LocalIntersect(r)
			if len(xs) != 0 {
				t.Fatal()
			}
		})
	}
}

func TestNormalOnTheSurfaceOfACube(t *testing.T) {
	tests := map[string]struct {
		point  *Tuple
		normal *Tuple
	}{
		"1": {Point(1, 0.5, -0.8), Vector(1, 0, 0)},
		"2": {Point(-1, -0.2, 0.9), Vector(-1, 0, 0)},
		"3": {Point(-0.4, 1, -0.1), Vector(0, 1, 0)},
		"4": {Point(-0.4, 1, -0.1), Vector(0, 1, 0)},
		"5": {Point(-0.6, 0.3, 1), Vector(0, 0, 1)},
		"6": {Point(0.4, 0.4, -1), Vector(0, 0, -1)},
		"7": {Point(1, 1, 1), Vector(1, 0, 0)},
		"8": {Point(-1, -1, -1), Vector(-1, 0, 0)},
	}
	c := NewCube()
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			normal := c.localNormalAt(tc.point)
			if !normal.Equal(tc.normal) {
				t.Fatal(normal)
			}
		})
	}
}
