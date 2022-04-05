package main

import (
	"github.com/hejiangda/go-raytracer-challenge/raytracer"
)

func main() {

	c := raytracer.NewCanvas(400, 400)
	red := raytracer.Color(1, 0, 0)
	// 球在(200,200,0),球半径为 50
	// 观察者在(200,200,200)处
	// 屏幕在(0,0,0)处，左上角为(0,0,0)
	camPos := raytracer.Point(200, 200, 200)

	translate := raytracer.Translation(200, 200, 0)
	scal := raytracer.Scaling(50, 50, 50)
	s := raytracer.NewSphere()
	s.SetTransform(raytracer.Multiply(translate, scal))
	for i := 0; i < 400; i++ {
		for j := 0; j < 400; j++ {
			r := raytracer.NewRay(camPos, raytracer.Subtract(raytracer.Point(float64(i), float64(j), 0), camPos))
			xs := raytracer.Intersect(s, r)
			_, err := raytracer.Hit(xs)
			if err == nil {
				raytracer.WritePixel(c, i, j, red)
			}
		}
	}
	c.SaveFile("chapter05.ppm")
}
