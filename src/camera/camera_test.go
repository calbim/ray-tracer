package camera

import (
	"github.com/calbim/ray-tracer/src/color"
	"github.com/calbim/ray-tracer/src/world"
	"fmt"
	"math"
	"testing"

	"github.com/calbim/ray-tracer/src/matrix"
	"github.com/calbim/ray-tracer/src/transforms"
	"github.com/calbim/ray-tracer/src/tuple"
)

func TestConstructCamera(t *testing.T) {
	hSize := 160.0
	vSize := 120.0
	fieldOfView := math.Pi / 2
	c := New(hSize, vSize, fieldOfView)
	if c.HSize != 160 {
		t.Errorf("wanted horizontal size=%v, got %v", hSize, c.HSize)
	}
	if c.VSize != 120 {
		t.Errorf("wanted vertical size=%v, got %v", vSize, c.VSize)
	}
	if c.FieldOfView != math.Pi/2 {
		t.Errorf("wanted field of view=%v, got %v", fieldOfView, c.FieldOfView)
	}
	if !c.Transform.Equals(matrix.Identity) {
		t.Errorf("wanted transform=%v, got %v", matrix.Identity, c.Transform)
	}
}

func TestPixelSizeHorizontalCanvas(t *testing.T) {
	c := New(200, 125, math.Pi/2)
	if c.PixelSize != 0.01 {
		t.Errorf("wanted pixel size=%v, got %v", 0.01, c.PixelSize)
	}
}

func TestPixelSizeVerticalCanvas(t *testing.T) {
	c := New(125, 200, math.Pi/2)
	if c.PixelSize != 0.01 {
		t.Errorf("wanted pixel size=%v, got %v", 0.01, c.PixelSize)
	}
}

func TestConstructRayCenterOfCanvas(t *testing.T) {
	c := New(201, 101, math.Pi/2)
	r := c.RayForPixel(100, 50)
	if !r.Origin.Equals(tuple.Point(0, 0, 0)) {
		t.Errorf("wanted origin=%v, got %v", tuple.Point(0, 0, 0), r.Origin)
	}
	if !r.Direction.Equals(tuple.Vector(0, 0, -1)) {
		t.Errorf("wanted vector=%v, got %v", tuple.Vector(0, 0, -1), r.Direction)
	}
}

func TestConstructRayCornerOfCanvas(t *testing.T) {
	c := New(201, 101, math.Pi/2)
	r := c.RayForPixel(0, 0)
	if !r.Origin.Equals(tuple.Point(0, 0, 0)) {
		t.Errorf("ray origin=%v, got %v", tuple.Point(0, 0, 0), r.Origin)
	}
	if !r.Direction.Equals(tuple.Vector(0.66519, 0.33259, -0.66851)) {
		t.Errorf("wanted ray vector=%v, got %v", tuple.Vector(0.66519, 0.33259, -0.66851), r.Direction)
	}
}

func TestConstructRayWhenCameraIsTransformed(t *testing.T) {
	c := New(201, 101, math.Pi/2)
	m := transforms.RotationY(math.Pi / 4)
	n := transforms.Translation(0, -2, 5)
	p := m.Multiply(n)
	c.Transform = p
	cp := &c
	fmt.Println(c.Transform)
	r := cp.RayForPixel(100, 50)
	if !r.Origin.Equals(tuple.Point(0, 2, -5)) {
		t.Errorf("wanted ray origin=%v, got %v", tuple.Point(0, 2, -5), r.Origin)
	}
	if !r.Direction.Equals(tuple.Vector(math.Sqrt(2)/2, 0, -math.Sqrt(2)/2)) {
		t.Errorf("wanted ray direction=%v, got %v", tuple.Vector(math.Sqrt(2)/2, 0, -math.Sqrt(2)/2), r.Direction)
	}
}

func TestRender(t *testing.T) {
	w := world.Default()
	c := New(11, 11, math.Pi/2)
	from := tuple.Point(0, 0, -5)
	to := tuple.Point(0, 0, 0)
	up := tuple.Vector(0, 1, 0)
	c.Transform = transforms.ViewTransform(from, to, up)
	image := c.Render(w)
	if !(image.Pixels[5][5].Equals(color.New(0.38066, 0.47583, 0.2855))) {
		t.Errorf("wanted pixel at 5,5 to =%v, but it is %v", color.New(0.38066, 0.47583, 0.2855), image.Pixels[5][5])
	}
}
