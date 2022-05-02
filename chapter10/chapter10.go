package main

import (
	"github.com/hejiangda/go-raytracer-challenge/raytracer"
	"math"
)

func main() {

	floor := raytracer.NewPlane()
	floor.Transform = raytracer.Scaling(10, 0.01, 10)
	floor.Material.Color = raytracer.Color(1, 0.9, 0.9)
	floor.Material.Specular = 0
	floor.Material.Pattern = raytracer.NewGradientPattern(raytracer.Color(1, 1, 0), raytracer.Color(0, 1, 1))
	floor.Material.Pattern.SetTransform(raytracer.Scaling(0.5, 0.5, 0.5))

	// large sphere in the middle is a unit sphere,
	// translated upward slightly and colored green
	middle := raytracer.NewSphere()
	middle.Transform = raytracer.Translation(-0.5, 1, 0.5)
	middle.Material.Color = raytracer.Color(0.1, 1, 0.5)
	middle.Material.Diffuse = 0.7
	middle.Material.Specular = 0.3
	middle.Material.Pattern = raytracer.NewCheckersPattern(raytracer.Color(1, 0, 0), raytracer.Color(0, 1, 0))
	middle.Material.Pattern.SetTransform(raytracer.Scaling(0.25, 0.25, 0.25))

	// smaller green sphere on the right is scaled in half
	right := raytracer.NewSphere()
	right.Transform = raytracer.Translation(1.5, 0.5, -0.5).Multiply(raytracer.Scaling(0.5, 0.5, 0.5))
	right.Material.Color = raytracer.Color(0.5, 1, 0.1)
	right.Material.Diffuse = 0.7
	right.Material.Specular = 0.3
	right.Material.Pattern = raytracer.NewStripePattern(raytracer.Color(1, 0, 0), raytracer.Color(0, 0, 1))
	right.Material.Pattern.SetTransform(raytracer.Scaling(0.5, 0.5, 0.5))

	// smallest sphere is scaled by a third
	left := raytracer.NewSphere()
	left.Transform = raytracer.Translation(-1.5, 0.33, -0.75).Multiply(raytracer.Scaling(0.33, 0.33, 0.33))
	left.Material.Color = raytracer.Color(1, 0.8, 0.1)
	left.Material.Diffuse = 0.7
	left.Material.Specular = 0.3
	left.Material.Pattern = raytracer.NewRingPattern(raytracer.BLACK, raytracer.WHITE)
	left.Material.Pattern.SetTransform(raytracer.Scaling(0.2, 0.2, 0.2))

	world := raytracer.NewWorld()
	world.Lights = append(world.Lights, raytracer.NewPointLight(raytracer.Point(-10, 10, -10), raytracer.Color(1, 1, 1)))
	world.Objects = append(world.Objects, middle, right, left, floor)

	camera := raytracer.NewCamera(800, 400, math.Pi/3)
	camera.Transform = raytracer.ViewTransform(raytracer.Point(0, 1.5, -5), raytracer.Point(0, 1, 0), raytracer.Vector(0, 1, 0))

	canvas := camera.Render(world)
	err := canvas.SaveFile("chapter10.ppm")
	if err != nil {
		return
	}
}
