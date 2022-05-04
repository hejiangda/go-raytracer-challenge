package raytracer

import (
	"math"
	"testing"
)

func TestARayMissesACylinder(t *testing.T) {
	tests := map[string]struct {
		origin    *Tuple
		direction *Tuple
	}{
		"1": {Point(1, 0, 0), Vector(0, 1, 0)},
		"2": {Point(0, 0, 0), Vector(0, 1, 0)},
		"3": {Point(0, 0, -5), Vector(1, 1, 1)},
	}
	cyl := NewCylinder()

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			direction := Normalize(tc.direction)
			r := NewRay(tc.origin, direction)
			xs := cyl.localIntersect(r)
			if len(xs) != 0 {
				t.Fatal()
			}
		})
	}
}

func TestARayStrikesACylinder(t *testing.T) {
	tests := map[string]struct {
		origin    *Tuple
		direction *Tuple
		t0        float64
		t1        float64
	}{
		"1": {Point(1, 0, -5), Vector(0, 0, 1), 5, 5},
		"2": {Point(0, 0, -5), Vector(0, 0, 1), 4, 6},
		"3": {Point(0.5, 0, -5), Vector(0.1, 1, 1), 6.80798, 7.08872},
	}
	cyl := NewCylinder()
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			direction := Normalize(tc.direction)
			r := NewRay(tc.origin, direction)
			xs := cyl.localIntersect(r)
			if len(xs) != 2 {
				t.Fatal("len(xs)!=2")
			}
			if !AlmostEqual(xs[0].T, tc.t0, Eps) || !AlmostEqual(xs[1].T, tc.t1, Eps) {
				t.Fatal(xs[0].T, xs[1].T)
			}
		})
	}
}
func TestNormalVectorOnACylinder(t *testing.T) {
	tests := map[string]struct {
		point  *Tuple
		normal *Tuple
	}{
		"1": {Point(1, 0, 0), Vector(1, 0, 0)},
		"2": {Point(0, 5, -1), Vector(0, 0, -1)},
		"3": {Point(0, -2, 1), Vector(0, 0, 1)},
		"4": {Point(-1, 1, 0), Vector(-1, 0, 0)},
	}
	cyl := NewCylinder()
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			n := cyl.localNormalAt(tc.point)
			if !n.Equal(tc.normal) {
				t.Fatal()
			}
		})
	}
}
func TestDefaultMinimumAndMaximumForACylinder(t *testing.T) {
	cyl := NewCylinder()
	if !math.IsInf(cyl.Minimum, -1) ||
		!math.IsInf(cyl.Maximum, 1) {
		t.Fatal()
	}
}
func TestIntersectingAConstrainedCylinder(t *testing.T) {
	tests := map[string]struct {
		point     *Tuple
		direction *Tuple
		count     int
	}{
		"1": {Point(0, 1.5, 0), Vector(0.1, 1, 0), 0},
		"2": {Point(0, 3, -5), Vector(0, 0, 1), 0},
		"3": {Point(0, 0, -5), Vector(0, 0, 1), 0},
		"4": {Point(0, 2, -5), Vector(0, 0, 1), 0},
		"5": {Point(0, 1, -5), Vector(0, 0, 1), 0},
		"6": {Point(0, 1.5, -2), Vector(0, 0, 1), 2},
	}
	cyl := NewCylinder()
	cyl.Minimum = 1
	cyl.Maximum = 2

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			direction := Normalize(tc.direction)
			r := NewRay(tc.point, direction)
			xs := cyl.localIntersect(r)
			if len(xs) != tc.count {
				t.Fatal()
			}
		})
	}
}

func TestDefaultClosedValueForACylinder(t *testing.T) {
	cyl := NewCylinder()
	if cyl.Closed != false {
		t.Fatal()
	}
}

func TestNormalVectorOnACylindersEndCaps(t *testing.T) {
	tests := map[string]struct {
		point  *Tuple
		normal *Tuple
	}{
		"1": {Point(0, 1, 0), Vector(0, -1, 0)},
		"2": {Point(0.5, 1, 0), Vector(0, -1, 0)},
		"3": {Point(0, 1, 0.5), Vector(0, -1, 0)},
		"4": {Point(0, 2, 0), Vector(0, 1, 0)},
		"5": {Point(0.5, 2, 0), Vector(0, 1, 0)},
		"6": {Point(0, 2, 0.5), Vector(0, 1, 0)},
	}
	cyl := NewCylinder()
	cyl.Minimum = 1
	cyl.Maximum = 2
	cyl.Closed = true
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			n := cyl.localNormalAt(tc.point)
			if !n.Equal(tc.normal) {
				t.Fatal()
			}
		})
	}
}
