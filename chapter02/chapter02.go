package main

import (
	"canvas"
	"math"
	"tuples"
)

func main() {
	p := projectile{
		Position: tuples.Point(0, 1, 0),
		Velocity: tuples.Normalize(tuples.Vector(1, 1.8, 0)).Multiply(11.25),
	}
	e := environment{
		Gravity: tuples.Vector(0, -0.1, 0),
		Wind:    tuples.Vector(-0.01, 0, 0),
	}
	c := canvas.NewCanvas(900, 550)
	red := tuples.Color(1, 0, 0)
	for i := 0; ; i++ {
		//fmt.Println("tick: ", i, " position: ", int(math.Round(p.Position.X)), int(math.Round(p.Position.Y)))
		if p.Position.Y <= 0 {
			break
		}
		canvas.WritePixel(c, int(math.Round(p.Position.X)), c.Height-int(math.Round(p.Position.Y)), red)
		p = tick(e, p)
	}
	c.SaveFile("chapter02.ppm")
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
