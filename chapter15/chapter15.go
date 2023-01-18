package main

import (
	"github.com/hejiangda/go-raytracer-challenge/raytracer"
	"math"
)

func main() {
	world := raytracer.NewWorld()

	ellipsoidObj := raytracer.ParseObjFile("ellipsoid.obj")
	ellipsoidObj.DefaultGroup.SetTransform(raytracer.Translation(0, 0, 1))

	world.Objects = append(world.Objects, ellipsoidObj.DefaultGroup)

	pointlight := raytracer.NewPointLight(raytracer.Point(0, 5, 0), raytracer.Color(1, 1, 1))
	world.Lights = append(world.Lights, pointlight)

	camera := raytracer.NewCamera(80, 40, math.Pi/3)
	camera.Transform = raytracer.ViewTransform(raytracer.Point(5, 20, 5), raytracer.Point(0, 0, 0), raytracer.Vector(1, 0, 0))

	canvas := camera.Render(world)
	err := canvas.SaveFile("chapter15.ppm")
	if err != nil {
		return
	}
}
