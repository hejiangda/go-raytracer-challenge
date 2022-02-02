package main

import (
	"canvas"
	"intersections"
	"matrices"
	"rays"
	"spheres"
	"transformations"
	"tuples"
)

func main() {

	c := canvas.NewCanvas(400, 400)
	red := tuples.Color(1, 0, 0)
	// 球在(200,200,0),球半径为 50
	// 观察者在(200,200,200)处
	// 屏幕在(0,0,0)处，左上角为(0,0,0)
	camPos := tuples.Point(200, 200, 200)

	translate := transformations.Translation(200, 200, 0)
	scal := transformations.Scaling(50, 50, 50)
	s := spheres.NewSphere()
	s.SetTransform(matrices.Multiply(translate, scal))
	for i := 0; i < 400; i++ {
		for j := 0; j < 400; j++ {
			r := rays.NewRay(camPos, tuples.Subtract(tuples.Point(float64(i), float64(j), 0), camPos))
			xs := intersections.Intersect(s, r)
			_, err := intersections.Hit(xs)
			if err == nil {
				canvas.WritePixel(c, i, j, red)
			}
		}
	}
	c.SaveFile("chapter05.ppm")
}
