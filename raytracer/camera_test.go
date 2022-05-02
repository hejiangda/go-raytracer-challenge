package raytracer

import (
	"math"
	"testing"
)

func TestNewCamera(t *testing.T) {
	hsize := int(160)
	vsize := int(120)
	fieldOfView := math.Pi / 2
	identityMatrix := EyeMatrix(4)
	c := NewCamera(hsize, vsize, fieldOfView)
	if c.HSize != hsize || c.VSize != vsize || c.FieldOfView != fieldOfView || !c.Transform.Equal(identityMatrix) {
		t.Fatal("failed! c:", c)
	}
}

func TestNewCamera2(t *testing.T) {
	c := NewCamera(125, 200, math.Pi/2)
	if !AlmostEqual(c.PixelSize, 0.01, Eps) {
		t.Fatal("failed c:", c)
	}
}
func TestNewCamera3(t *testing.T) {
	c := NewCamera(200, 125, math.Pi/2)
	if !AlmostEqual(c.PixelSize, 0.01, Eps) {
		t.Fatal("failed c:", c)
	}
}

func TestRayForPixel(t *testing.T) {
	c := NewCamera(201, 101, math.Pi/2)
	r := c.RayForPixel(100, 50)
	if !r.Origin.Equal(Point(0, 0, 0)) || !r.Direction.Equal(Vector(0, 0, -1)) {
		t.Fatal("failed! r:", r)
	}
}

func TestRayForPixel2(t *testing.T) {
	c := NewCamera(201, 101, math.Pi/2)
	r := c.RayForPixel(0, 0)
	if !r.Origin.Equal(Point(0, 0, 0)) || !r.Direction.Equal(Vector(0.66519, 0.33259, -0.66851)) {
		t.Fatal("failed! r:", r)
	}
}

func TestRayForPixel3(t *testing.T) {
	c := NewCamera(201, 101, math.Pi/2)
	c.Transform = RotationY(math.Pi / 4).Multiply(Translation(0, -2, 5))
	r := c.RayForPixel(100, 50)
	if !r.Origin.Equal(Point(0, 2, -5)) || !r.Direction.Equal(Vector(math.Sqrt2/2, 0, -math.Sqrt2/2)) {
		t.Fatal("failed! r:", r)
	}
}

func TestCamera_Render(t *testing.T) {
	w := DefaultWorld()
	c := NewCamera(11, 11, math.Pi/2)
	from := Point(0, 0, -5)
	to := Point(0, 0, 0)
	up := Vector(0, 1, 0)
	c.Transform = ViewTransform(from, to, up)
	image := c.Render(w)
	if !PixelAt(image, 5, 5).Equal(Color(0.38066, 0.47583, 0.2855)) {
		t.Fatal("failed! image", image)
	}
}
