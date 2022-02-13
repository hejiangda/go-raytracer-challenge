package main

import (
	"fmt"
	"raytracer"
)

func main() {
	p := projectile{
		Position: raytracer.Point(0, 1, 0),
		Velocity: raytracer.Normalize(raytracer.Vector(1, 1, 0)),
	}
	e := environment{
		Gravity: raytracer.Vector(0, -0.1, 0),
		Wind:    raytracer.Vector(-0.01, 0, 0),
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
	Gravity *raytracer.Tuple
	Wind    *raytracer.Tuple
}

type projectile struct {
	Position *raytracer.Tuple
	Velocity *raytracer.Tuple
}

func tick(e environment, p projectile) (res projectile) {
	res.Position = raytracer.Add(p.Position, p.Velocity)
	res.Velocity = raytracer.Add(raytracer.Add(p.Velocity, e.Gravity), e.Wind)
	return
}
