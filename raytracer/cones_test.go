package raytracer

import (
	"math"
	"testing"
)

func TestIntersectingAConeWithARay(t *testing.T) {
	tests := map[string]struct {
		origin    *Tuple
		direction *Tuple
		t0        float64
		t1        float64
	}{
		"1": {Point(0, 0, -5), Vector(0, 0, 1), 5, 5},
		"2": {Point(0, 0, -5), Vector(1, 1, 1), 8.66025, 8.66025},
		"3": {Point(1, 1, -5), Vector(-0.5, -1, 1), 4.55006, 49.44994},
	}
	shape := NewCone()
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			direction := Normalize(tc.direction)
			r := NewRay(tc.origin, direction)
			xs := shape.localIntersect(r)
			if len(xs) != 2 {
				t.Fatal()
			}
			if !AlmostEqual(xs[0].T, tc.t0, Eps) ||
				!AlmostEqual(xs[1].T, tc.t1, Eps) {
				t.Fatal("xs[0]", xs[0].T, tc.t0, "xs[1]", xs[1].T, tc.t1)
			}
		})
	}
}

func TestIntersectingAConesEndCaps(t *testing.T) {
	tests := map[string]struct {
		origin    *Tuple
		direction *Tuple
		count     int
	}{
		"1": {Point(0, 0, -5), Vector(0, 1, 0), 0},
		"2": {Point(0, 0, -0.25), Vector(0, 1, 1), 2},
		"3": {Point(0, 0, -0.25), Vector(0, 1, 0), 4},
	}
	shape := NewCone()
	shape.Minimum = -0.5
	shape.Maximum = 0.5
	shape.Closed = true
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			direction := Normalize(tc.direction)
			r := NewRay(tc.origin, direction)
			xs := shape.localIntersect(r)
			if len(xs) != tc.count {
				t.Fatal(len(xs), tc.count)
			}
		})
	}
}
func TestComputingTheNormalVectorOnACone(t *testing.T) {
	tests := map[string]struct {
		point  *Tuple
		normal *Tuple
	}{
		"1": {Point(0, 0, 0), Vector(0, 0, 0)},
		"2": {Point(1, 1, 1), Vector(1, -math.Sqrt2, 1)},
		"3": {Point(-1, -1, 0), Vector(-1, 1, 0)},
	}
	shape := NewCone()
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			n := shape.localNormalAt(tc.point)
			if !n.Equal(tc.normal) {
				t.Fatal()
			}
		})
	}
}
