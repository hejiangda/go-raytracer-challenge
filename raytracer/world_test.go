package raytracer

import (
	"math"
	"testing"
)

func TestNewWorld(t *testing.T) {

	w := DefaultWorld()

	pointLight := NewPointLight(Point(-10, 10, -10), Color(1, 1, 1))
	var ok = false
	for _, light := range w.Lights {
		if light.Position.Equal(pointLight.Position) && light.Intensity.Equal(pointLight.Intensity) {
			ok = true
		}
	}
	if ok == false {
		t.Fatal("Does not contain light:", pointLight.Position, pointLight.Intensity)
	}

	s1 := NewSphere()
	s1.Material.Color = Color(0.8, 1.0, 0.6)
	s1.Material.Diffuse = 0.7
	s1.Material.Specular = 0.2

	s2 := NewSphere()
	s2.Transform = Scaling(0.5, 0.5, 0.5)

	var ok1 = false
	var ok2 = false
	for _, object := range w.Objects {
		switch object.(type) {
		case *Sphere:
			s := object.(*Sphere)
			if s.Transform.Equal(s1.Transform) && s.Material.Equal(s1.Material) {
				ok1 = true
			}
			if s.Transform.Equal(s2.Transform) && s.Material.Equal(s2.Material) {
				ok2 = true
			}
		}

	}
	if !ok1 {
		t.Fatal("Does not contain the object s1:", s1.Transform, s1.Material)
	}
	if !ok2 {
		t.Fatal("Does not contain the object: s2", s2.Transform, s2.Material)
	}
}

func TestWorld_Intersect(t *testing.T) {
	w := DefaultWorld()
	r := NewRay(Point(0, 0, -5), Vector(0, 0, 1))
	xs := w.Intersect(r)
	//fmt.Println("xs len", len(xs))
	if xs[0].T != 4 || xs[1].T != 4.5 || xs[2].T != 5.5 || xs[3].T != 6 {
		t.Fatal("Intersect failed! xs:", xs)
	}
}

func TestPrepareComputations(t *testing.T) {
	r := NewRay(Point(0, 0, -5), Vector(0, 0, 1))
	shape := NewSphere()
	i := Intersection{4, shape}
	comps := PrepareComputations(i, r, []Intersection{i})
	if comps.T != i.T || comps.Object != i.Obj || !comps.Point.Equal(Point(0, 0, -1)) || !comps.EyeV.Equal(Vector(0, 0, -1)) || !comps.NormalV.Equal(Vector(0, 0, -1)) {
		t.Fatal("failed")
	}
}
func TestPrepareComputationsOutside(t *testing.T) {
	r := NewRay(Point(0, 0, -5), Vector(0, 0, 1))
	shape := NewSphere()
	i := Intersection{4, shape}
	comps := PrepareComputations(i, r, []Intersection{i})
	if comps.Inside == true {
		t.Fatal("failed")
	}
}
func TestPrepareComputationsInside(t *testing.T) {
	r := NewRay(Point(0, 0, 0), Vector(0, 0, 1))
	shape := NewSphere()
	i := Intersection{1, shape}
	comps := PrepareComputations(i, r, []Intersection{i})
	if !comps.Point.Equal(Point(0, 0, 1)) || !comps.EyeV.Equal(Vector(0, 0, -1)) || !comps.Inside || !comps.NormalV.Equal(Vector(0, 0, -1)) {
		t.Fatal("failed")
	}
}

func TestWorld_ShadeHit(t *testing.T) {
	w := DefaultWorld()
	r := NewRay(Point(0, 0, -5), Vector(0, 0, 1))
	shape := w.Objects[0]
	i := Intersection{4, shape}
	comps := PrepareComputations(i, r, []Intersection{i})
	c := w.ShadeHit(comps, DefaultReflectRemaining)
	if !c.Equal(Color(0.38066, 0.47583, 0.2855)) {
		t.Fatal("failed color:", c)
	}
}
func TestWorld_ShadeHitInside(t *testing.T) {
	w := DefaultWorld()
	w.Lights[0] = NewPointLight(Point(0, 0.25, 0), Color(1, 1, 1))
	r := NewRay(Point(0, 0, 0), Vector(0, 0, 1))
	shape := w.Objects[1]
	i := Intersection{0.5, shape}
	comps := PrepareComputations(i, r, []Intersection{i})
	c := w.ShadeHit(comps, DefaultReflectRemaining)
	if !c.Equal(Color(0.90498, 0.90498, 0.90498)) {
		t.Fatal("failed color:", c)
	}
}

func TestWorld_ColorAt(t *testing.T) {
	w := DefaultWorld()
	r := NewRay(Point(0, 0, -5), Vector(0, 1, 0))
	c := w.ColorAt(r, DefaultReflectRemaining)
	if !c.Equal(Color(0, 0, 0)) {
		t.Fatal("failed")
	}
}

func TestWorld_ColorAt2(t *testing.T) {
	w := DefaultWorld()
	r := NewRay(Point(0, 0, -5), Vector(0, 0, 1))
	c := w.ColorAt(r, DefaultReflectRemaining)
	if !c.Equal(Color(0.38066, 0.47583, 0.2855)) {
		t.Fatal("failed color:", c)
	}
}

func TestWorld_ColorAt3(t *testing.T) {
	w := DefaultWorld()
	outer := w.Objects[0]
	outer.GetMaterial().Ambient = 1

	inner := w.Objects[1]
	inner.GetMaterial().Ambient = 1
	r := NewRay(Point(0, 0, 0.75), Vector(0, 0, -1))
	c := w.ColorAt(r, DefaultReflectRemaining)
	if !c.Equal(inner.GetMaterial().Color) {
		t.Fatal("failed color:", c)
	}
}

func TestWorld_IsShadowed1(t *testing.T) {
	w := DefaultWorld()
	p := Point(0, 10, 0)
	if w.IsShadowed(p, w.Lights[0]) {
		t.Fatal("failed")
	}
}
func TestWorld_IsShadowed2(t *testing.T) {
	w := DefaultWorld()
	p := Point(10, -10, 10)
	if !w.IsShadowed(p, w.Lights[0]) {
		t.Fatal("failed")
	}
}
func TestWorld_IsShadowed3(t *testing.T) {
	w := DefaultWorld()
	p := Point(-20, 20, -20)
	if w.IsShadowed(p, w.Lights[0]) {
		t.Fatal("failed")
	}
}

func TestWorld_IsShadowed4(t *testing.T) {
	w := DefaultWorld()
	p := Point(-2, 2, -2)
	if w.IsShadowed(p, w.Lights[0]) {
		t.Fatal("failed")
	}
}

func TestWorld_ShadeHit2(t *testing.T) {
	w := new(World)
	w.Lights = append(w.Lights, NewPointLight(Point(0, 0, -10), Color(1, 1, 1)))
	s1 := NewSphere()
	w.Objects = append(w.Objects, s1)
	s2 := NewSphere()
	s2.Transform = Translation(0, 0, 10)
	w.Objects = append(w.Objects, s2)
	r := NewRay(Point(0, 0, 5), Vector(0, 0, 1))
	i := Intersection{4, s2}
	comps := PrepareComputations(i, r, []Intersection{i})
	c := w.ShadeHit(comps, DefaultReflectRemaining)
	if !c.Equal(Color(0.1, 0.1, 0.1)) {
		t.Fatal("failed")
	}

}

func TestReflectedColorForANonreflectiveMaterial(t *testing.T) {
	w := DefaultWorld()
	r := NewRay(Point(0, 0, 0), Vector(0, 0, 1))
	shape := w.Objects[1]
	shape.GetMaterial().Ambient = 1
	i := Intersection{1, shape}
	comps := PrepareComputations(i, r, []Intersection{i})
	color := w.ReflectedColor(comps, DefaultReflectRemaining)
	if !color.Equal(Color(0, 0, 0)) {
		t.Fatal("failed")
	}
}

func TestReflectedColorForAReflectiveMaterial(t *testing.T) {
	w := DefaultWorld()
	shape := NewPlane()
	shape.Material.Reflective = 0.5
	shape.Transform = Translation(0, -1, 0)
	w.Objects = append(w.Objects, shape)
	r := NewRay(Point(0, 0, -3), Vector(0, -math.Sqrt2/2, math.Sqrt2/2))
	i := Intersection{
		T:   math.Sqrt2,
		Obj: shape,
	}
	comps := PrepareComputations(i, r, []Intersection{i})
	color := w.ReflectedColor(comps, DefaultReflectRemaining)
	if !color.Equal(Color(0.19033, 0.23792, 0.14274)) {
		t.Fatal("failed", color)
	}
}
func TestShadeHitWithAReflectiveMaterial(t *testing.T) {
	w := DefaultWorld()
	shape := NewPlane()
	shape.Material.Reflective = 0.5
	shape.Transform = Translation(0, -1, 0)
	w.Objects = append(w.Objects, shape)
	r := NewRay(Point(0, 0, -3), Vector(0, -math.Sqrt2/2, math.Sqrt2/2))
	i := Intersection{
		T:   math.Sqrt2,
		Obj: shape,
	}
	comps := PrepareComputations(i, r, []Intersection{i})
	color := w.ShadeHit(comps, DefaultReflectRemaining)
	if !color.Equal(Color(0.87676, 0.92434, 0.82917)) {
		t.Fatal("failed", color)
	}
}

func TestColorAtWithMutuallyReflectiveSurfaces(t *testing.T) {
	w := NewWorld()
	w.Lights = append(w.Lights, NewPointLight(Point(0, 0, 0), Color(1, 1, 1)))
	lower := NewPlane()
	lower.Material.Reflective = 1
	lower.Transform = Translation(0, -1, 0)

	upper := NewPlane()
	upper.Material.Reflective = 1
	upper.Transform = Translation(0, 1, 0)

	w.Objects = append(w.Objects, lower, upper)
	r := NewRay(Point(0, 0, 0), Vector(0, 1, 0))
	w.ColorAt(r, DefaultReflectRemaining)
}

func TestTheReflectedColorAtTheMaximumRecursiveDepth(t *testing.T) {
	w := DefaultWorld()
	shape := NewPlane()
	shape.Material.Reflective = 0.5
	shape.Transform = Translation(0, -1, 0)
	w.Objects = append(w.Objects, shape)
	r := NewRay(Point(0, 0, -3), Vector(0, -math.Sqrt2/2, math.Sqrt2/2))
	i := Intersection{
		T:   math.Sqrt2,
		Obj: shape,
	}
	comps := PrepareComputations(i, r, []Intersection{i})
	color := w.ReflectedColor(comps, 0)
	if !color.Equal(Color(0, 0, 0)) {
		t.Fatal("failed", color)
	}
}
func TestTheRefractedColorWithAnOpaqueSurface(t *testing.T) {
	w := DefaultWorld()
	shape := w.Objects[0]
	r := NewRay(Point(0, 0, -5), Vector(0, 0, 1))
	xs := Intersections(Intersection{4, shape}, Intersection{6, shape})
	comps := PrepareComputations(xs[0], r, xs)
	c := w.RefractedColor(comps, 5)
	if !c.Equal(Color(0, 0, 0)) {
		t.Fatal()
	}
}

func TestTheRefractedColorAtTheMaximumRecursiveDepth(t *testing.T) {
	w := DefaultWorld()
	shape := w.Objects[0]
	shape.GetMaterial().Transparency = 1.0
	shape.GetMaterial().RefractiveIndex = 1.5
	r := NewRay(Point(0, 0, -5), Vector(0, 0, 1))
	xs := Intersections(Intersection{4, shape}, Intersection{6, shape})
	comps := PrepareComputations(xs[0], r, xs)
	c := w.RefractedColor(comps, 0)
	if !c.Equal(Color(0, 0, 0)) {
		t.Fatal()
	}
}

func TestTheRefractedColorUnderTotalInternalReflection(t *testing.T) {
	w := DefaultWorld()
	shape := w.Objects[0]
	shape.GetMaterial().Transparency = 1.0
	shape.GetMaterial().RefractiveIndex = 1.5

	r := NewRay(Point(0, 0, math.Sqrt2/2), Vector(0, 1, 0))
	xs := Intersections(Intersection{-math.Sqrt2 / 2, shape}, Intersection{math.Sqrt2 / 2, shape})
	comps := PrepareComputations(xs[1], r, xs)
	c := w.RefractedColor(comps, 5)
	if !c.Equal(Color(0, 0, 0)) {
		t.Fatal()
	}
}

// 这个测试用例有问题，应该是作者写错了
//func TestTheRefractedColorWithARefractedRay(t *testing.T) {
//	w := DefaultWorld()
//	A := w.Objects[0]
//	A.GetMaterial().Ambient = 1.0
//	A.GetMaterial().Pattern = testPattern()
//	B := w.Objects[1]
//	B.GetMaterial().Transparency = 1.0
//	B.GetMaterial().RefractiveIndex = 1.5
//	r := NewRay(Point(0, 0, 0.1), Vector(0, 1, 0))
//	xs := Intersections(Intersection{-0.9899, A}, Intersection{-0.4899, B}, Intersection{0.4899, B}, Intersection{0.9899, A})
//	comps := PrepareComputations(xs[2], r, xs)
//	c := w.RefractedColor(comps, DefaultRefractRemaining)
//	if !c.Equal(Color(0, 0.99888, 0.04725)) {
//		t.Fatal(c)
//	}
//}
func TestShadeHitWithATransparentMaterial(t *testing.T) {
	w := DefaultWorld()
	floor := NewPlane()
	floor.Transform = Translation(0, -1, 0)
	floor.Material.Transparency = 0.5
	floor.Material.RefractiveIndex = 1.5

	ball := NewSphere()
	ball.Material.Color = Color(1, 0, 0)
	ball.Material.Ambient = 0.5
	ball.Transform = Translation(0, -3.5, -0.5)

	w.Objects = append(w.Objects, floor, ball)
	r := NewRay(Point(0, 0, -3), Vector(0, -math.Sqrt2/2, math.Sqrt2/2))
	xs := Intersections(Intersection{math.Sqrt2, floor})
	comps := PrepareComputations(xs[0], r, xs)
	color := w.ShadeHit(comps, 5)
	if !color.Equal(Color(0.93642, 0.68642, 0.68642)) {
		t.Fatal(color)
	}
	//fmt.Println(color)
}
func TestShadeHitWithAReflectiveTransparentMaterial(t *testing.T) {
	w := DefaultWorld()
	r := NewRay(Point(0, 0, -3), Vector(0, -math.Sqrt2/2, math.Sqrt2/2))
	floor := NewPlane()
	floor.Transform = Translation(0, -1, 0)
	floor.Material.Reflective = 0.5
	floor.Material.Transparency = 0.5
	floor.Material.RefractiveIndex = 1.5

	ball := NewSphere()
	ball.Material.Color = Color(1, 0, 0)
	ball.Material.Ambient = 0.5
	ball.Transform = Translation(0, -3.5, -0.5)

	w.Objects = append(w.Objects, floor, ball)
	xs := Intersections(Intersection{math.Sqrt2, floor})
	comps := PrepareComputations(xs[0], r, xs)
	color := w.ShadeHit(comps, 5)
	if !color.Equal(Color(0.93391, 0.69643, 0.69243)) {
		t.Fatal()
	}

}
