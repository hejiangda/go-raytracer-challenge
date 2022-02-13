package raytracer

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Canvas struct {
	Width  int
	Height int
	Data   []Tuple
}

func NewCanvas(width, height int) *Canvas {
	c := new(Canvas)
	c.Width = width
	c.Height = height
	c.Data = make([]Tuple, width*height)
	//for i := 0; i < width*height; i++ {
	//	c.Data[i] = Color(0, 0, 0)
	//}
	return c
}

func WritePixel(c *Canvas, x, y int, pixel *Tuple) {
	c.Data[y*c.Width+x] = *pixel
}
func PixelAt(c *Canvas, x, y int) *Tuple {
	return &c.Data[y*c.Width+x]
}

func ToPPM(c *Canvas) string {
	header := fmt.Sprintf("P3\n%v %v\n255", c.Width, c.Height)
	var builder strings.Builder
	builder.WriteString(header)
	limitAndClamp := func(val float64) int {
		val *= 255
		if val > 255 {
			val = 255
		}
		if val < 0 {
			val = 0
		}
		return int(math.Round(val))
	}
	var lineLen int

	for i := 0; i < c.Width*c.Height; i++ {
		color := c.Data[i]
		k := 0
		getNextValStr := func() string {
			var ret string
			switch k {
			case 0:
				ret = strconv.Itoa(limitAndClamp(color.X))
			case 1:
				ret = strconv.Itoa(limitAndClamp(color.Y))
			case 2:
				ret = strconv.Itoa(limitAndClamp(color.Z))
			}
			k++
			return ret
		}
		if i%c.Width == 0 {
			builder.WriteString("\n")
			lineLen = 0
			str := getNextValStr()
			lineLen += len(str)
			builder.WriteString(str)
		}
		for k <= 2 {
			str := getNextValStr()
			if lineLen+len(str)+1 > 70 {
				builder.WriteString("\n")
				lineLen = 0
				lineLen += len(str)
				builder.WriteString(str)
			} else {
				builder.WriteString(" ")
				builder.WriteString(str)
				lineLen += len(str) + 1
			}
		}
	}
	builder.WriteString("\n")
	return builder.String()
}
func (c *Canvas) SaveFile(fname string) error {
	f, err := os.Create(fname)
	if err != nil {
		return err
	}
	defer f.Close()
	s := ToPPM(c)
	_, err = f.WriteString(s)
	if err != nil {
		return err
	}
	return nil
}
