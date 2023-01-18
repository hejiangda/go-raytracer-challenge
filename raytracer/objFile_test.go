package raytracer

import (
	"fmt"
	"testing"
)

// Ignoring unrecognized lines
func TestParseObjFile1(t *testing.T) {
	obj := ParseObjFile("test/objFileTest1.obj")
	if len(obj.Vertices) != 1 || len(obj.DefaultGroup.Children) != 0 {
		fmt.Println(len(obj.Vertices), len(obj.DefaultGroup.Children))
		t.Fatal("error")
	}
}

func TestParseObjFile2(t *testing.T) {
	obj := ParseObjFile("test/objFileTest2.obj")
	if len(obj.Vertices) != 5 || len(obj.DefaultGroup.Children) != 0 {
		fmt.Println(len(obj.Vertices), len(obj.DefaultGroup.Children))
		t.Fatal("error")
	}
	if !obj.Vertices[1].Equal(Point(-1, 1, 0)) ||
		!obj.Vertices[2].Equal(Point(-1, 0.5, 0)) ||
		!obj.Vertices[3].Equal(Point(1, 0, 0)) ||
		!obj.Vertices[4].Equal(Point(1, 1, 0)) {
		t.Fatal("error")
	}
}

func TestParseObjFile3(t *testing.T) {
	obj := ParseObjFile("test/objFileTest3.obj")
	if len(obj.Vertices) != 5 || len(obj.DefaultGroup.Children) != 2 {
		fmt.Println(len(obj.Vertices), len(obj.DefaultGroup.Children))
		t.Fatal("error")
	}
	t1 := obj.DefaultGroup.Children[0].(*Triangle)
	t2 := obj.DefaultGroup.Children[1].(*Triangle)

	if !t1.P1.Equal(obj.Vertices[1]) ||
		!t1.P2.Equal(obj.Vertices[2]) ||
		!t1.P3.Equal(obj.Vertices[3]) ||
		!t2.P1.Equal(obj.Vertices[1]) ||
		!t2.P2.Equal(obj.Vertices[3]) ||
		!t2.P3.Equal(obj.Vertices[4]) {
		t.Fatal("error")
	}

}
