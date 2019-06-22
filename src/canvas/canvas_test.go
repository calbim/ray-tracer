package canvas

import (
	"fmt"
	"strings"
	"testing"

	"../tuple"
)

func TestCanvas(t *testing.T) {
	c := New(10, 20)
	if c.width != 10 || c.height != 20 {
		t.Errorf("Canvas width and height should be 10 and 20 respectively")
	}
	for i := 0; i < c.width; i++ {
		for j := 0; j < c.height; j++ {
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
	if !tuple.Equals(c.pixels[2][3], r) {
		t.Errorf("Pixel at (2,3) should be red")
	}
}

func TestCanvastoPPM(t *testing.T) {
	c := New(5, 3)
	ppm := ToPPM(c)
	ppmSplit := strings.Split(ppm, "\n")
	fmt.Println(ppmSplit)
	if ppmSplit[0] != "P3" {
		t.Errorf("First line of ppm header should be P3")
	}
	if ppmSplit[1] != "5 3" {
		t.Errorf("Second line of ppm header should be 5 3")
	}
	if ppmSplit[2] != "255" {
		t.Errorf("Third line of ppm header should be 255")
	}
}
