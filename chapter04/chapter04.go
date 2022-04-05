package main

import (
	"github.com/hejiangda/go-raytracer-challenge/raytracer"
	"math"
)

func main() {

	c := raytracer.NewCanvas(400, 400)
	white := raytracer.Color(1, 1, 1)
	center := raytracer.Point(0, 0, 0)
	translate := raytracer.Translation(0, -100, 0)
	translate1 := raytracer.Translation(200, 200, 0)
	for i := 0; i < 12; i++ {
		rotate := raytracer.RotationZ(math.Pi / 6 * float64(i))
		pos := raytracer.MultiplyTuple(translate1.Multiply(rotate).Multiply(translate), center)
		raytracer.WritePixel(c, int(math.Round(pos.X)), int(math.Round(pos.Y)), white)
	}
	c.SaveFile("chapter04.ppm")
}
