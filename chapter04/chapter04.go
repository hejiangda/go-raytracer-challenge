package main

import (
	"canvas"
	"math"
	"matrices"
	"transformations"
	"tuples"
)

func main() {

	c := canvas.NewCanvas(400, 400)
	white := tuples.Color(1, 1, 1)
	center := tuples.Point(0, 0, 0)
	translate := transformations.Translation(0, -100, 0)
	translate1 := transformations.Translation(200, 200, 0)
	for i := 0; i < 12; i++ {
		rotate := transformations.RotationZ(math.Pi / 6 * float64(i))
		pos := matrices.MultiplyTuple(translate1.Multiply(rotate).Multiply(translate), center)
		canvas.WritePixel(c, int(math.Round(pos.X)), int(math.Round(pos.Y)), white)
	}
	c.SaveFile("chapter04.ppm")
}
