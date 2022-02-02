package main

import (
	"canvas"
	"intersections"
	"lights"
	"matrices"
	"rays"
	"spheres"
	"transformations"
	"tuples"
)

func main() {

	c := canvas.NewCanvas(400, 400)
	// 球在(200,200,0),球半径为 50
	// 观察者在(200,200,200)处
	// 屏幕在(0,0,0)处，左上角为(0,0,0)
	camPos := tuples.Point(200, 200, 200)

	translate := transformations.Translation(200, 200, 0)
	scal := transformations.Scaling(50, 50, 50)
	s := spheres.NewSphere()
	s.Material.Color = tuples.Color(1, 0.2, 1)
	lightPosition := tuples.Point(50, 50, 200)
	lightColor := tuples.Color(1, 1, 1)
	light := lights.NewPointLight(lightPosition, lightColor)

	s.SetTransform(matrices.Multiply(translate, scal))
	for i := 0; i < 400; i++ {
		for j := 0; j < 400; j++ {
			r := rays.NewRay(camPos, tuples.Subtract(tuples.Point(float64(i), float64(j), 0), camPos))
			xs := intersections.Intersect(s, r)
			hit, err := intersections.Hit(xs)
			if err == nil {
				point := rays.Position(r, hit.T)
				normal := hit.Obj.NormalAt(point)
				eye := tuples.Normalize(r.Direction.Negate())

				color := lights.Lighting(hit.Obj.GetMaterial(), light, point, eye, normal)
				canvas.WritePixel(c, i, j, color)
			}
		}
	}
	err := c.SaveFile("chapter06.ppm")
	if err != nil {
		return
	}
}
