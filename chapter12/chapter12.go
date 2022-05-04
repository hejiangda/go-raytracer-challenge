package main

import (
	"github.com/hejiangda/go-raytracer-challenge/raytracer"
	"math"
)

func main() {

	world := raytracer.NewWorld()
	floor := raytracer.NewPlane()
	floor.Transform = raytracer.Translation(0, -1, 0)
	floor.Material.Pattern = raytracer.NewCheckersPattern(raytracer.WHITE, raytracer.BLACK)
	floor.Material.Pattern.SetTransform(raytracer.Scaling(0.2, 0.2, 0.2))
	floor.Material.Reflective = 0.1

	cube := raytracer.NewCube()
	cube.Material.Color = raytracer.Color(1, 0, 0)
	//cube.Transform = raytracer.RotationX(math.SqrtPi).Multiply(raytracer.RotationY(math.SqrtPi)).Multiply(raytracer.RotationZ(math.SqrtPi)).Multiply(raytracer.Scaling(0.5, 0.5, 0.5))
	cube.Transform = raytracer.Scaling(0.5, 0.5, 0.5)
	cube.Material.Ambient = 0
	cube.Material.Reflective = 0
	cube.Material.Specular = 1
	cube.Material.RefractiveIndex = 1
	cube.Material.Transparency = 0
	cube.Material.Shininess = 300

	//ball1 := raytracer.GlassSphere()
	//ball1.Material.RefractiveIndex = 1
	//ball1.Material.Color = raytracer.Color(0, 0, 0)
	//ball1.Transform = raytracer.Scaling(0.5, 0.5, 0.5).Multiply(raytracer.Translation(0, 0.5, 0))

	world.Objects = append(world.Objects, floor, cube)

	pointlight := raytracer.NewPointLight(raytracer.Point(0, 10, 10), raytracer.Color(1, 1, 1))
	//pointlight1 := raytracer.NewPointLight(raytracer.Point(0, 10, -10), raytracer.Color(1, 1, 1))
	//pointlight2 := raytracer.NewPointLight(raytracer.Point(10, 10, 0), raytracer.Color(1, 1, 1))
	//pointlight3 := raytracer.NewPointLight(raytracer.Point(-10, 10, 0), raytracer.Color(1, 1, 1))
	world.Lights = append(world.Lights, pointlight)

	camera := raytracer.NewCamera(800, 400, math.Pi/3)
	camera.Transform = raytracer.ViewTransform(raytracer.Point(0, 5, 0), raytracer.Point(0, 0, 0), raytracer.Vector(1, 0, 0))

	canvas := camera.Render(world)
	err := canvas.SaveFile("chapter12.ppm")
	if err != nil {
		return
	}
}
