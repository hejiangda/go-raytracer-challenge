package main

import (
	"github.com/hejiangda/go-raytracer-challenge/raytracer"
	"math"
)

//func hexagonCorner() *raytracer.Sphere {
//	corner := raytracer.NewSphere()
//	corner.SetTransform(raytracer.Translation(0, 0, -1).Multiply(raytracer.Scaling(0.25, 0.25, 0.25)))
//	return corner
//}
//
//func hexagonEdge() *raytracer.Cylinder {
//	edge := raytracer.NewCylinder()
//	edge.Minimum = 0
//	edge.Maximum = 1
//	edge.SetTransform(raytracer.Translation(0, 0, -1).Multiply(raytracer.RotationY(-math.Pi / 6)).Multiply(raytracer.RotationZ(-math.Pi / 2)).Multiply(raytracer.Scaling(0.25, 1, 0.25)))
//	return edge
//}
//
//func hexagonSide() *raytracer.Group {
//	side := raytracer.NewGroup()
//	side.AddChild(hexagonCorner())
//	side.AddChild(hexagonEdge())
//	return side
//}
//func hexagon() *raytracer.Group {
//	hex := raytracer.NewGroup()
//	for i := 0; i < 6; i++ {
//		side := hexagonSide()
//		side.SetTransform(raytracer.RotationY(float64(i) * math.Pi / 3))
//		hex.AddChild(side)
//	}
//	return hex
//}
func main() {
	world := raytracer.NewWorld()
	floor := raytracer.NewPlane()
	floor.Transform = raytracer.Translation(0, -1, 0)
	floor.Material.Pattern = raytracer.NewCheckersPattern(raytracer.WHITE, raytracer.BLACK)
	floor.Material.Pattern.SetTransform(raytracer.Scaling(0.2, 0.2, 0.2))
	floor.Material.Reflective = 0.1

	//h := hexagon()

	t := raytracer.NewTriangle(raytracer.Point(1, 0, 0), raytracer.Point(-1, 0, 0), raytracer.Point(0, 0, 1))
	world.Objects = append(world.Objects, floor, t)

	pointlight := raytracer.NewPointLight(raytracer.Point(0, 1, 1), raytracer.Color(1, 1, 1))
	world.Lights = append(world.Lights, pointlight)

	camera := raytracer.NewCamera(800, 400, math.Pi/3)
	camera.Transform = raytracer.ViewTransform(raytracer.Point(0, 3, 5), raytracer.Point(0, 0, 0), raytracer.Vector(0, 1, 0))

	canvas := camera.Render(world)
	err := canvas.SaveFile("chapter15.ppm")
	if err != nil {
		return
	}
}
