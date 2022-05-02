package raytracer

import (
	"fmt"
	"testing"
)

func TestStripePattern(t *testing.T) {
	p := NewStripePattern(WHITE, BLACK)
	if !p.A.Equal(WHITE) || !p.B.Equal(BLACK) {
		t.Fatal("failed")
	}
}

func TestPattern_StripeAt(t *testing.T) {
	pattern := NewStripePattern(WHITE, BLACK)
	if !pattern.At(Point(0, 0, 0)).Equal(WHITE) ||
		!pattern.At(Point(0, 1, 0)).Equal(WHITE) ||
		!pattern.At(Point(0, 2, 0)).Equal(WHITE) {
		t.Fatal("failed")
	}
}
func TestPattern_StripeAt1(t *testing.T) {
	pattern := NewStripePattern(WHITE, BLACK)
	if !pattern.At(Point(0, 0, 0)).Equal(WHITE) ||
		!pattern.At(Point(0, 0, 1)).Equal(WHITE) ||
		!pattern.At(Point(0, 0, 2)).Equal(WHITE) {
		t.Fatal("failed")
	}
}
func TestPattern_StripeAt3(t *testing.T) {
	pattern := NewStripePattern(WHITE, BLACK)
	if !pattern.At(Point(0, 0, 0)).Equal(WHITE) ||
		!pattern.At(Point(0.9, 0, 0)).Equal(WHITE) ||
		!pattern.At(Point(1, 0, 0)).Equal(BLACK) ||
		!pattern.At(Point(-0.1, 0, 0)).Equal(BLACK) ||
		!pattern.At(Point(-1, 0, 0)).Equal(BLACK) ||
		!pattern.At(Point(-1.1, 0, 0)).Equal(WHITE) {

		t.Fatal("failed")
	}
}

func TestLightingWithAPatternApplied(t *testing.T) {
	material := NewMaterial()
	material.Pattern = NewStripePattern(WHITE, BLACK)
	material.Ambient = 1
	material.Diffuse = 0
	material.Specular = 0
	eyev := Vector(0, 0, -1)
	normalv := Vector(0, 0, -1)
	light := NewPointLight(Point(0, 0, -10), WHITE)
	obj := NewSphere()
	c1 := Lighting(material, obj, light, Point(0.9, 0, 0), eyev, normalv, false)
	c2 := Lighting(material, obj, light, Point(1.1, 0, 0), eyev, normalv, false)
	if !c1.Equal(WHITE) ||
		!c2.Equal(BLACK) {
		fmt.Println(c1)
		fmt.Println(c2)
		t.Fatal("failed")
	}
}
func TestPattern_StripeAtObject(t *testing.T) {
	obj := NewSphere()
	obj.SetTransform(Scaling(2, 2, 2))
	patt := NewStripePattern(WHITE, BLACK)
	c := patt.AtShape(obj, Point(1.5, 0, 0))
	if !c.Equal(WHITE) {
		t.Fatal("failed")
	}
}
func TestPattern_SetTransform(t *testing.T) {
	obj := NewSphere()
	patt := NewStripePattern(WHITE, BLACK)
	patt.SetTransform(Scaling(2, 2, 2))
	c := patt.AtShape(obj, Point(1.5, 0, 0))
	if !c.Equal(WHITE) {
		t.Fatal("failed")
	}
}

func TestPattern_StripeAtObject2(t *testing.T) {
	obj := NewSphere()
	obj.SetTransform(Scaling(2, 2, 2))
	patt := NewStripePattern(WHITE, BLACK)
	patt.SetTransform(Translation(0.5, 0, 0))
	c := patt.AtShape(obj, Point(2.5, 0, 0))
	if !c.Equal(WHITE) {
		t.Fatal("failed")
	}
}

func testPattern() Pattern {
	return NewStripePattern(WHITE, BLACK)
}

func TestAPatternWithAnObjectTransformation(t *testing.T) {
	shape := NewSphere()
	shape.SetTransform(Scaling(2, 2, 2))
	pattern := testPattern()
	c := pattern.AtShape(shape, Point(2, 3, 4))
	if c.Equal(Color(1, 1.5, 2)) {
		t.Fatal("failed")
	}
}

func TestAPatternWithAPatternTransformation(t *testing.T) {
	shape := NewSphere()
	pattern := testPattern()
	pattern.SetTransform(Scaling(2, 2, 2))
	c := pattern.AtShape(shape, Point(2, 3, 4))
	if c.Equal(Color(1, 1.5, 2)) {
		t.Fatal("failed")
	}
}

func TestAPatternWithBothAnObjectAndAPattern(t *testing.T) {
	shape := NewSphere()
	shape.SetTransform(Scaling(2, 2, 2))
	pattern := testPattern()
	pattern.SetTransform(Translation(0.5, 1, 1.5))
	c := pattern.AtShape(shape, Point(2.5, 3, 3.5))
	if c.Equal(Color(0.75, 0.5, 0.25)) {
		t.Fatal("failed")
	}
}
func TestGradientPattern_At(t *testing.T) {
	pattern := NewGradientPattern(WHITE, BLACK)
	if !pattern.At(Point(0, 0, 0)).Equal(WHITE) ||
		!pattern.At(Point(0.25, 0, 0)).Equal(Color(0.75, 0.75, 0.75)) ||
		!pattern.At(Point(0.5, 0, 0)).Equal(Color(0.5, 0.5, 0.5)) ||
		!pattern.At(Point(0.75, 0, 0)).Equal(Color(0.25, 0.25, 0.25)) {
		t.Fatal("failed")
	}
}

func TestRingPattern_At(t *testing.T) {
	pattern := NewRingPattern(WHITE, BLACK)
	if !pattern.At(Point(0, 0, 0)).Equal(WHITE) ||
		!pattern.At(Point(1, 0, 0)).Equal(BLACK) ||
		!pattern.At(Point(0, 0, 1)).Equal(BLACK) ||
		!pattern.At(Point(0.708, 0, 0.708)).Equal(BLACK) {
		t.Fatal("failed")
	}
}

func TestCheckersPattern_At(t *testing.T) {
	pattern := NewCheckersPattern(WHITE, BLACK)
	if !pattern.At(Point(0, 0, 0)).Equal(WHITE) ||
		!pattern.At(Point(0.99, 0, 0)).Equal(WHITE) ||
		!pattern.At(Point(1.01, 0, 0)).Equal(BLACK) {
		t.Fatal("failed")
	}
}

func TestCheckersPattern_At2(t *testing.T) {
	pattern := NewCheckersPattern(WHITE, BLACK)
	if !pattern.At(Point(0, 0, 0)).Equal(WHITE) ||
		!pattern.At(Point(0, 0.99, 0)).Equal(WHITE) ||
		!pattern.At(Point(0, 1.01, 0)).Equal(BLACK) {
		t.Fatal("failed")
	}
}

func TestCheckersPattern_At3(t *testing.T) {
	pattern := NewCheckersPattern(WHITE, BLACK)
	if !pattern.At(Point(0, 0, 0)).Equal(WHITE) ||
		!pattern.At(Point(0, 0, 0.99)).Equal(WHITE) ||
		!pattern.At(Point(0, 0, 1.01)).Equal(BLACK) {
		t.Fatal("failed")
	}
}
