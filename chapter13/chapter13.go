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

	cylinder := raytracer.NewCylinder()
	cylinder.Material.Color = raytracer.Color(1, 0, 0)
	//cylinder.Transform = raytracer.RotationX(math.SqrtPi).Multiply(raytracer.RotationY(math.SqrtPi)).Multiply(raytracer.RotationZ(math.SqrtPi)).Multiply(raytracer.Scaling(0.5, 0.5, 0.5))
	cylinder.Transform = raytracer.Scaling(0.5, 0.5, 0.5).Multiply(raytracer.Translation(0, 0, -4))
	cylinder.Material.Ambient = 0
	cylinder.Material.Reflective = 0
	cylinder.Material.Specular = 1
	cylinder.Material.RefractiveIndex = 1
	cylinder.Material.Transparency = 0
	cylinder.Material.Shininess = 300
	cylinder.Minimum = 0
	cylinder.Maximum = 1
	cylinder.Closed = true

	cone := raytracer.NewCone()
	cone.Material.Color = raytracer.Color(0, 0, 1)
	cone.Transform = raytracer.Scaling(0.5, 0.5, 0.5).Multiply(raytracer.Translation(1, 0.5, 4))
	cone.Material.Ambient = 0
	cone.Material.Reflective = 0
	cone.Material.Specular = 1
	cone.Material.RefractiveIndex = 1
	cone.Material.Transparency = 0
	cone.Material.Shininess = 300
	cone.Minimum = -1
	cone.Maximum = 0
	cone.Closed = true

	//ball1 := raytracer.GlassSphere()
	//ball1.Material.RefractiveIndex = 1
	//ball1.Material.Color = raytracer.Color(0, 0, 0)
	//ball1.Transform = raytracer.Scaling(0.5, 0.5, 0.5).Multiply(raytracer.Translation(0, 0.5, 0))

	world.Objects = append(world.Objects, floor, cylinder, cone)

	pointlight := raytracer.NewPointLight(raytracer.Point(0, 10, 10), raytracer.Color(1, 1, 1))
	//pointlight1 := raytracer.NewPointLight(raytracer.Point(0, 10, -10), raytracer.Color(1, 1, 1))
	//pointlight2 := raytracer.NewPointLight(raytracer.Point(10, 10, 0), raytracer.Color(1, 1, 1))
	//pointlight3 := raytracer.NewPointLight(raytracer.Point(-10, 10, 0), raytracer.Color(1, 1, 1))
	world.Lights = append(world.Lights, pointlight)

	camera := raytracer.NewCamera(800, 400, math.Pi/3)
	camera.Transform = raytracer.ViewTransform(raytracer.Point(0, 5, 0), raytracer.Point(0, 0, 0), raytracer.Vector(1, 0, 0))

	canvas := camera.Render(world)
	err := canvas.SaveFile("chapter13.ppm")
	if err != nil {
		return
	}
}
