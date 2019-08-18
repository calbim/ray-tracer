package camera

import (
	"math"

	"github.com/calbim/ray-tracer/src/canvas"
	"github.com/calbim/ray-tracer/src/matrix"
	"github.com/calbim/ray-tracer/src/world"

	"github.com/calbim/ray-tracer/src/ray"
	"github.com/calbim/ray-tracer/src/tuple"
)

//Camera represents a camera
type Camera struct {
	HSize       float64 //in pixels
	VSize       float64
	FieldOfView float64
	Transform   *matrix.Matrix
	PixelSize   float64
	HalfHeight  float64
	HalfWidth   float64
}

// New returns a new camera
func New(hSize float64, vSize float64, fieldOfView float64) Camera {
	transform := matrix.Identity
	c := Camera{
		HSize:       hSize,
		VSize:       vSize,
		FieldOfView: fieldOfView,
		Transform:   transform,
	}
	halfView := math.Tan(fieldOfView / 2)
	aspect := float64(hSize / vSize)
	if aspect >= 1 {
		c.HalfWidth = halfView
		c.HalfHeight = halfView / aspect
	} else {
		c.HalfWidth = halfView * aspect
		c.HalfHeight = halfView
	}
	c.PixelSize = (c.HalfWidth * 2) / c.HSize
	return c
}

//RayForPixel returns the ray from the camera to point (x,y) on canvas
func (c *Camera) RayForPixel(x, y int) *ray.Ray {
	xOffset := (float64(x) + 0.5) * c.PixelSize
	yOffset := (float64(y) + 0.5) * c.PixelSize
	worldX := c.HalfWidth - xOffset
	worldY := c.HalfHeight - yOffset
	inverse, err := c.Transform.Inverse()
	if err != nil {
		return nil
	}
	pixel := inverse.MultiplyTuple(tuple.Point(worldX, worldY, -1))
	origin := inverse.MultiplyTuple(tuple.Point(0, 0, 0))
	direction := pixel.Subtract(origin)
	direction = direction.Normalize()
	r := ray.New(origin, direction)
	return &r
}

//Render renders the world with a camera
func (c Camera) Render(w world.World) *canvas.Canvas {
	image := canvas.New(int(c.HSize), int(c.VSize))
	for y := 0; y < int(c.VSize); y++ {
		for x := 0; x < int(c.HSize); x++ {
			r := c.RayForPixel(x, y)
			col := w.ColorAt(*r)
			image.WritePixel(x, y, col)
		}
	}
	return &image
}
