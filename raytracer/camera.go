package raytracer

import "math"

type Camera struct {
	HSize       int
	VSize       int
	FieldOfView float64
	Transform   *Matrix
	PixelSize   float64
	HalfWidth   float64
	HalfHeight  float64
}

func NewCamera(hsize, vsize int, fov float64) (cam *Camera) {
	cam = new(Camera)
	cam.HSize = hsize
	cam.VSize = vsize
	cam.FieldOfView = fov
	cam.Transform = EyeMatrix(4)
	halfView := math.Tan(fov / 2)
	aspect := float64(hsize) / float64(vsize)
	if aspect >= 1 {
		cam.HalfWidth = halfView
		cam.HalfHeight = halfView / aspect
	} else {
		cam.HalfWidth = halfView * aspect
		cam.HalfHeight = halfView
	}
	cam.PixelSize = (cam.HalfWidth * 2) / float64(cam.HSize)
	return
}

func (camera *Camera) RayForPixel(px, py int) (ray *Ray) {
	xOffset := (float64(px) + 0.5) * camera.PixelSize
	yOffset := (float64(py) + 0.5) * camera.PixelSize

	worldX := camera.HalfWidth - xOffset
	worldY := camera.HalfHeight - yOffset

	invTrans, _ := Inverse(camera.Transform)
	pixel := MultiplyTuple(invTrans, Point(worldX, worldY, -1))
	origin := MultiplyTuple(invTrans, Point(0, 0, 0))
	direction := Normalize(pixel.Subtract(origin))
	ray = NewRay(origin, direction)
	return
}

func (camera *Camera) Render(world *World) (image *Canvas) {
	image = NewCanvas(camera.HSize, camera.VSize)

	for y := 0; y < camera.VSize-1; y++ {
		for x := 0; x < camera.HSize-1; x++ {
			ray := camera.RayForPixel(x, y)
			color := world.ColorAt(ray)
			image.WritePixel(x, y, color)
		}
	}
	return
}
