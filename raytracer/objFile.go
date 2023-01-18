package raytracer

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

type ObjFile struct {
	Vertices     []*Tuple
	DefaultGroup *Group
}

func newObjFile() *ObjFile {
	var obj ObjFile
	// vertices 从1开始索引
	obj.Vertices = append(obj.Vertices, nil)
	obj.DefaultGroup = NewGroup()
	return &obj
}

func ParseObjFile(fp string) *ObjFile {
	obj := newObjFile()
	f, err := os.OpenFile(fp, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return nil
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		txt := scanner.Text()
		var cmd string
		fmt.Sscanf(txt, "%s", &cmd)
		switch cmd {
		case "v":
			var p1, p2, p3 float64
			fmt.Sscanf(txt, "%s%f%f%f", &cmd, &p1, &p2, &p3)
			obj.Vertices = append(obj.Vertices, Point(p1, p2, p3))
		case "f":
			var varr []int
			buf := bytes.NewBufferString(txt)
			fmt.Fscan(buf, &cmd)
			for {
				var idx int
				_, err := fmt.Fscan(buf, &idx)
				if err != nil {
					break
				}
				varr = append(varr, idx)
			}
			for i := 2; i < len(varr); i++ {
				obj.DefaultGroup.AddChild(NewTriangle(obj.Vertices[varr[0]], obj.Vertices[varr[i-1]], obj.Vertices[varr[i]]))
			}
		case "g":

		}
	}
	return obj
}
