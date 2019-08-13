package canvas

import (
	"fmt"
	"strings"
	"testing"

	"github.com/calbim/ray-tracer/src/color"
)

func TestCanvas(t *testing.T) {
	c := New(10, 20)
	if c.width != 10 || c.height != 20 {
		t.Errorf("Canvas width and height should be 10 and 20 respectively")
	}
	for i := 0; i < c.height; i++ {
		for j := 0; j < c.width; j++ {
			pixelColor := c.Pixels[i][j]
			if !pixelColor.Equals(color.Black) {
				t.Errorf("wanted pixel color to be %v, got %v", color.Black, pixelColor)
			}
		}
	}
}

func TestWriteToCanvas(t *testing.T) {
	c := New(10, 20)
	r := color.New(1, 0, 0)
	c.WritePixel(2, 3, r)
	col := c.Pixels[3][2]
	if !col.Equals(r) {
		t.Errorf("wanted color at (3,2)= %v, got %v", r, col)
	}
}

func TestCanvastoPPM(t *testing.T) {
	c := New(5, 3)

	c1 := color.New(1.5, 0, 0)
	c2 := color.New(0, 0.5, 0)
	c3 := color.New(-0.5, 0, 1)

	c.WritePixel(0, 0, c1)
	c.WritePixel(2, 1, c2)
	c.WritePixel(4, 2, c3)

	ppm := c.ToPPM()
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
	fmt.Println(ppmSplit[3])

	if strings.Trim(ppmSplit[3], " ") != "255 0 0 0 0 0 0 0 0 0 0 0 0 0 0" {
		t.Errorf("Line 1: Incorect PPM conversion")
	}
	if strings.Trim(ppmSplit[4], " ") != "0 0 0 0 0 0 0 128 0 0 0 0 0 0 0" {
		t.Errorf("Line 2: Incorect PPM conversion")
	}
	if strings.Trim(ppmSplit[5], " ") != "0 0 0 0 0 0 0 0 0 0 0 0 0 0 255" {
		t.Errorf("Line 3: Incorect PPM conversion")
	}
}
