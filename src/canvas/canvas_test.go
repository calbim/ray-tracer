package canvas

import (
	"strings"
	"testing"

	"../tuple"
)

func TestCanvas(t *testing.T) {
	c := New(10, 20)
	if c.width != 10 || c.height != 20 {
		t.Errorf("Canvas width and height should be 10 and 20 respectively")
	}
	for i := 0; i < c.height; i++ {
		for j := 0; j < c.width; j++ {
			if !tuple.Equals(c.pixels[i][j], tuple.Color(0, 0, 0)) {
				t.Errorf("Every pixel of a new canvas should be black")
			}
		}
	}
}

func TestWriteToCanvas(t *testing.T) {
	c := New(10, 20)
	r := tuple.Color(1, 0, 0)
	WritePixel(&c, 2, 3, r)
	if !tuple.Equals(c.pixels[3][2], r) {
		t.Errorf("Pixel at width 2 amd height 3 should be red")
	}
}
func TestCanvastoPPM(t *testing.T) {
	c := New(5, 3)

	c1 := tuple.Color(1.5, 0, 0)
	c2 := tuple.Color(0, 0.5, 0)
	c3 := tuple.Color(-0.5, 0, 1)

	WritePixel(&c, 0, 0, c1)
	WritePixel(&c, 2, 1, c2)
	WritePixel(&c, 4, 2, c3)

	ppm := ToPPM(c)
	ppmSplit := strings.Split(ppm, "\n")
	if ppmSplit[0] != "P3" {
		t.Errorf("First line of ppm header should be P3")
	}
	if ppmSplit[1] != "5 3" {
		t.Errorf("Second line of ppm header should be 5 3")
	}
	if ppmSplit[2] != "255" {
		t.Errorf("Third line of ppm header should be 255")
	}
	if ppmSplit[3] != "255 0 0 0 0 0 0 0 0 0 0 0 0 0 0" {
		t.Errorf("Line 1: Incorect PPM conversion")
	}
	if ppmSplit[4] != "0 0 0 0 0 0 0 128 0 0 0 0 0 0 0" {
		t.Errorf("Line 2: Incorect PPM conversion")
	}
	if ppmSplit[5] != "0 0 0 0 0 0 0 0 0 0 0 0 0 0 255" {
		t.Errorf("Line 3: Incorect PPM conversion")
	}
}

// func TestSplitLongLinesinPPM(t *testing.T) {
// 	c := New(10, 2)
// 	color := tuple.Color(1, 0.8, 0.6)
// 	for i := 0; i < 10; i++ {
// 		for j := 0; j < 2; j++ {
// 			WritePixel(&c, i, j, color)
// 		}
// 	}
// 	ppm := ToPPM(c)
// 	ppmSplit := strings.Split(ppm, "\n")
// 	if ppmSplit[3] != "255 204 153 255 204 153 255 204 153 255 204 153 255 204" {
// 		t.Errorf("Incorrect splitting of long strings")
// 	}
// 	if ppmSplit[4] != "153 255 204 153 255 204 153 255 204 153 255 204 153" {
// 		t.Errorf("Incorrect splitting of long strings")
// 	}
// 	if ppmSplit[5] != "255 204 153 255 204 153 255 204 153 255 204 153 255 204" {
// 		t.Errorf("Incorrect splitting of long strings")
// 	}
// 	if ppmSplit[6] != "153 255 204 153 255 204 153 255 204 153 255 204 153" {
// 		t.Errorf("Incorrect splitting of long strings")
// 	}

// }
