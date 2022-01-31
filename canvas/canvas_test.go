package canvas

import (
	"bufio"
	"strings"
	"testing"
	"tuples"
)

func TestCanvas(t *testing.T) {
	c := NewCanvas(10, 20)
	if c.Width != 10 {
		t.Fatal("c.Width!=10")
	}
	if c.Height != 20 {
		t.Fatal("c.Height!=20")
	}
	if c.Data == nil {
		t.Fatal("c.Data is not initialized")
	}
	for i := 0; i < c.Height*c.Width; i++ {
		if !c.Data[i].Equal(tuples.Color(0, 0, 0)) {
			t.Fatal("canvas initialized color is not equal to color(0,0,0)")
		}
	}
}

func TestWritePixelsToCanvas(t *testing.T) {
	c := NewCanvas(10, 20)
	red := tuples.Color(1, 0, 0)
	WritePixel(c, 2, 3, red)
	if !PixelAt(c, 2, 3).Equal(red) {
		t.Fatal("Pixel at c (2,3) is not red")
	}
}
func TestConstructPPMHeader(t *testing.T) {
	c := NewCanvas(5, 3)
	ppm := ToPPM(c)
	if ppm != "P3\n5 3\n255" {
		t.Fatal("construct ppm header failed! header:", ppm)
	}
}
func TestConstructPPmData(t *testing.T) {
	c := NewCanvas(5, 3)
	c1 := tuples.Color(1.5, 0, 0)
	c2 := tuples.Color(0, 0.5, 0)
	c3 := tuples.Color(-0.5, 0, 1)
	WritePixel(c, 0, 0, c1)
	WritePixel(c, 2, 1, c2)
	WritePixel(c, 4, 2, c3)
	ppm := ToPPM(c)
	scanner := bufio.NewScanner(strings.NewReader(ppm))
	var i int
	for scanner.Scan() {
		i++
		switch i {
		case 4:
			if scanner.Text() != "255 0 0 0 0 0 0 0 0 0 0 0 0 0 0" {
				t.Fatal("line 4 error line:", scanner.Text())
			}
		case 5:
			if scanner.Text() != "0 0 0 0 0 0 0 128 0 0 0 0 0 0 0" {
				t.Fatal("line 5 error line:", scanner.Text())
			}
		case 6:
			if scanner.Text() != "0 0 0 0 0 0 0 0 0 0 0 0 0 0 255" {
				t.Fatal("line 6 error")
			}
		}
	}
	if i < 6 {
		t.Fatal("Construct data failed")
	}
}
func TestConstructPPmDataSplitLongLines(t *testing.T) {
	c := NewCanvas(10, 2)
	pixel := tuples.Color(1, 0.8, 0.6)
	for i := 0; i < c.Width*c.Height; i++ {
		WritePixel(c, i%c.Width, i/c.Width, pixel)
	}
	ppm := ToPPM(c)
	scanner := bufio.NewScanner(strings.NewReader(ppm))
	var i int
	for scanner.Scan() {
		i++
		switch i {
		case 4:
			if scanner.Text() != "255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204" {
				t.Fatal("line 4 error line:", scanner.Text())
			}
		case 5:
			if scanner.Text() != "153 255 204 153 255 204 153 255 204 153 255 204 153" {
				t.Fatal("line 5 error line:", scanner.Text())
			}
		case 6:
			if scanner.Text() != "255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204" {
				t.Fatal("line 6 error")
			}
		case 7:
			if scanner.Text() != "153 255 204 153 255 204 153 255 204 153 255 204 153" {
				t.Fatal("line 5 error line:", scanner.Text())
			}
		}
	}
	if i < 7 {
		t.Fatal("Construct data failed")
	}
}
func TestPPMTerminatedByNewlineCharacter(t *testing.T) {
	c := NewCanvas(5, 3)
	ppm := ToPPM(c)
	lastChar := ppm[len(ppm)-1]
	if lastChar != '\n' {
		t.Fatal("PPM files are not terminated by a newline character")
	}

}
