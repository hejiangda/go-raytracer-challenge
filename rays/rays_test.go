package rays

import (
	"testing"
	"transformations"
	"tuples"
)

func TestRay(t *testing.T) {
	origin := tuples.Point(1, 2, 3)
	direction := tuples.Vector(4, 5, 6)
	r := NewRay(origin, direction)
	if !origin.Equal(r.Origin) {
		t.Fatal("wrong ray")
	}
	if !direction.Equal(r.Direction) {
		t.Fatal("wrong ray")
	}
}

func TestPosition(t *testing.T) {
	r := NewRay(tuples.Point(2, 3, 4), tuples.Vector(1, 0, 0))
	if !tuples.Point(2, 3, 4).Equal(Position(r, 0)) {
		t.Fatal("wrong position")
	}
	if !tuples.Point(3, 3, 4).Equal(Position(r, 1)) {
		t.Fatal("wrong position")
	}
	if !tuples.Point(1, 3, 4).Equal(Position(r, -1)) {
		t.Fatal("wrong position")
	}
	if !tuples.Point(4.5, 3, 4).Equal(Position(r, 2.5)) {
		t.Fatal("wrong position")
	}
}

func TestTranslatingRay(t *testing.T) {
	r := NewRay(tuples.Point(1, 2, 3), tuples.Vector(0, 1, 0))
	m := transformations.Translation(3, 4, 5)
	r2 := Transform(r, m)
	if !r2.Origin.Equal(tuples.Point(4, 6, 8)) {
		t.Fatal("r2.origin != point(4,6,8)")
	}
	if !r2.Direction.Equal(tuples.Vector(0, 1, 0)) {
		t.Fatal("r2.direction != vector(0,1,0)")
	}
}

func TestScalingRay(t *testing.T) {
	r := NewRay(tuples.Point(1, 2, 3), tuples.Vector(0, 1, 0))
	m := transformations.Scaling(2, 3, 4)
	r2 := Transform(r, m)
	if !r2.Origin.Equal(tuples.Point(2, 6, 12)) {
		t.Fatal("r2.origin != point(2,6,12)")
	}
	if !r2.Direction.Equal(tuples.Vector(0, 3, 0)) {
		t.Fatal("r2.direction!=vector(0,3,0)")
	}
}
