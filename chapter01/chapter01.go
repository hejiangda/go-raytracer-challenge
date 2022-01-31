package main

import (
	"fmt"
	"tuples"
)

func main() {
	p := projectile{
		Position: tuples.Point(0, 1, 0),
		Velocity: tuples.Normalize(tuples.Vector(1, 1, 0)),
	}
	e := environment{
		Gravity: tuples.Vector(0, -0.1, 0),
		Wind:    tuples.Vector(-0.01, 0, 0),
	}
	for i := 0; ; i++ {
		fmt.Println("tick: ", i, " position: ", p.Position)
		if p.Position.Y <= 0 {
			break
		}
		p = tick(e, p)
	}
}

type environment struct {
	Gravity *tuples.Tuple
	Wind    *tuples.Tuple
}

type projectile struct {
	Position *tuples.Tuple
	Velocity *tuples.Tuple
}

func tick(e environment, p projectile) (res projectile) {
	res.Position = tuples.Add(p.Position, p.Velocity)
	res.Velocity = tuples.Add(tuples.Add(p.Velocity, e.Gravity), e.Wind)
	return
}
