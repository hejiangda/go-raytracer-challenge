package raytracer

import (
	"testing"
)

func TestRay(t *testing.T) {
	origin := Point(1, 2, 3)
	direction := Vector(4, 5, 6)
	r := NewRay(origin, direction)
	if !origin.Equal(r.Origin) {
		t.Fatal("wrong ray")
	}
	if !direction.Equal(r.Direction) {
		t.Fatal("wrong ray")
	}
}

func TestPosition(t *testing.T) {
	r := NewRay(Point(2, 3, 4), Vector(1, 0, 0))
	if !Point(2, 3, 4).Equal(r.Position(0)) {
		t.Fatal("wrong position")
	}
	if !Point(3, 3, 4).Equal(r.Position(1)) {
		t.Fatal("wrong position")
	}
	if !Point(1, 3, 4).Equal(r.Position(-1)) {
		t.Fatal("wrong position")
	}
	if !Point(4.5, 3, 4).Equal(r.Position(2.5)) {
		t.Fatal("wrong position")
	}
}

func TestTranslatingRay(t *testing.T) {
	r := NewRay(Point(1, 2, 3), Vector(0, 1, 0))
	m := Translation(3, 4, 5)
	r2 := r.Transform(m)
	if !r2.Origin.Equal(Point(4, 6, 8)) {
		t.Fatal("r2.origin != point(4,6,8)")
	}
	if !r2.Direction.Equal(Vector(0, 1, 0)) {
		t.Fatal("r2.direction != vector(0,1,0)")
	}
}

func TestScalingRay(t *testing.T) {
	r := NewRay(Point(1, 2, 3), Vector(0, 1, 0))
	m := Scaling(2, 3, 4)
	r2 := r.Transform(m)
	if !r2.Origin.Equal(Point(2, 6, 12)) {
		t.Fatal("r2.origin != point(2,6,12)")
	}
	if !r2.Direction.Equal(Vector(0, 3, 0)) {
		t.Fatal("r2.direction!=vector(0,3,0)")
	}
}
