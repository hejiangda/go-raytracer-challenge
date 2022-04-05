package main

import (
	"github.com/hejiangda/go-raytracer-challenge/raytracer"
	"math"
)

func main() {
	p := projectile{
		Position: raytracer.Point(0, 1, 0),
		Velocity: raytracer.Normalize(raytracer.Vector(1, 1.8, 0)).Multiply(11.25),
	}
	e := environment{
		Gravity: raytracer.Vector(0, -0.1, 0),
		Wind:    raytracer.Vector(-0.01, 0, 0),
	}
	c := raytracer.NewCanvas(900, 550)
	red := raytracer.Color(1, 0, 0)
	for i := 0; ; i++ {
		//fmt.Println("tick: ", i, " position: ", int(math.Round(p.Position.X)), int(math.Round(p.Position.Y)))
		if p.Position.Y <= 0 {
			break
		}
		raytracer.WritePixel(c, int(math.Round(p.Position.X)), c.Height-int(math.Round(p.Position.Y)), red)
		p = tick(e, p)
	}
	c.SaveFile("chapter02.ppm")
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
