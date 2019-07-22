package camera

import (
	"math"
	"testing"

	"github.com/calbim/ray-tracer/src/world"

	"github.com/calbim/ray-tracer/src/transformations"

	"github.com/calbim/ray-tracer/src/tuple"

	"github.com/calbim/ray-tracer/src/matrix"
)

func TestConstructCamera(t *testing.T) {
	hSize := 160.0
	vSize := 120.0
	fieldOfView := math.Pi / 2
	c := New(hSize, vSize, fieldOfView)
	if c.HSize != 160 {
		t.Errorf("Horizontal size of canvas should be %v, but it is %v", hSize, c.HSize)
	}
	if c.VSize != 120 {
		t.Errorf("Vertical size of canvas should be %v, but it is %v", vSize, c.VSize)
	}
	if c.FieldOfView != math.Pi/2 {
		t.Errorf("Field of view of camera should be %v, but it is %v", fieldOfView, c.FieldOfView)

	}
	if !matrix.Equals(c.Transform, matrix.NewIdentity(), 4, 4, 4, 4) {
		t.Errorf("Transform should be %v, but it is %v", matrix.NewIdentity(), c.Transform)
	}
}

func TestPixelSizeHorizontalCanvas(t *testing.T) {
	c := New(200, 125, math.Pi/2)
	if c.PixelSize != 0.01 {
		t.Errorf("Pixel size should be %v, but it is %v", 0.01, c.PixelSize)
	}
}

func TestPixelSizeVerticalCanvas(t *testing.T) {
	c := New(125, 200, math.Pi/2)
	if c.PixelSize != 0.01 {
		t.Errorf("Pixel size should be %v, but it is %v", 0.01, c.PixelSize)
	}
}

func TestConstructRayCenterOfCanvas(t *testing.T) {
	c := New(201, 101, math.Pi/2)
	r, err := RayForPixel(c, 100, 50)
	if err != nil {
		t.Errorf("Error %v while calculating ray for pixel", err)
	}
	if !tuple.Equals(r.Origin, tuple.Point(0, 0, 0)) {
		t.Errorf("Ray origin should be %v, but it is %v", tuple.Point(0, 0, 0), r.Origin)
	}
	if !tuple.Equals(r.Direction, tuple.Vector(0, 0, -1)) {
		t.Errorf("Ray vector should be %v, but it is %v", tuple.Vector(0, 0, -1), r.Direction)
	}
}

func TestConstructRayCornerOfCanvas(t *testing.T) {
	c := New(201, 101, math.Pi/2)
	r, err := RayForPixel(c, 0, 0)
	if err != nil {
		t.Errorf("Error %v while calculating ray for pixel", err)
	}
	if !tuple.Equals(r.Origin, tuple.Point(0, 0, 0)) {
		t.Errorf("Ray origin should be %v, but it is %v", tuple.Point(0, 0, 0), r.Origin)
	}
	if !tuple.Equals(r.Direction, tuple.Vector(0.66519, 0.33259, -0.66851)) {
		t.Errorf("Ray vector should be %v, but it is %v", tuple.Vector(0.66519, 0.33259, -0.66851), r.Direction)
	}
}

func TestConstructRayWhenCameraIsTransformed(t *testing.T) {
	c := New(201, 101, math.Pi/2)
	c.Transform = matrix.Multiply(transformations.RotationY(math.Pi/4), transformations.NewTranslation(0, -2, 5))
	r, err := RayForPixel(c, 100, 50)
	if err != nil {
		t.Errorf("Error %v while calculating ray for pixel", err)
	}
	if !tuple.Equals(r.Origin, tuple.Point(0, 2, -5)) {
		t.Errorf("Ray origin should be %v, but it is %v", tuple.Point(0, 2, -5), r.Origin)
	}
	if !tuple.Equals(r.Direction, tuple.Vector(math.Sqrt(2)/2, 0, -math.Sqrt(2)/2)) {
		t.Errorf("Ray vector should be %v, but it is %v", tuple.Vector(math.Sqrt(2)/2, 0, -math.Sqrt(2)/2), r.Direction)
	}
}

func TestRender(t *testing.T) {
	w, err := world.NewDefault()
	if err != nil {
		t.Errorf("Error creating world %v", err)
	}
	c := New(11, 11, math.Pi/2)
	from := tuple.Point(0, 0, -5)
	to := tuple.Point(0, 0, 0)
	up := tuple.Vector(0, 1, 0)
	c.Transform = transformations.ViewTransform(from, to, up)
	image, err := c.Render(*w)
	if err != nil {
		t.Errorf("Error rendering image due to error %v", err)
	}
	if !tuple.Equals(image.Pixels[5][5], tuple.Color(0.38066, 0.47583, 0.2855)) {
		t.Errorf("Colour of pixel at 5,5 should be %v, but it is %v", tuple.Color(0.38066, 0.47583, 0.2855), image.Pixels[5][5])
	}
}
