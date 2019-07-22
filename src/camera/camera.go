package camera

import (
	"errors"
	"fmt"
	"math"

	"github.com/calbim/ray-tracer/src/canvas"

	"github.com/calbim/ray-tracer/src/world"

	"github.com/calbim/ray-tracer/src/ray"
	"github.com/calbim/ray-tracer/src/tuple"

	"github.com/calbim/ray-tracer/src/matrix"
)

// Camera represents a camera
type Camera struct {
	HSize       float64 //in pixels
	VSize       float64
	FieldOfView float64
	Transform   [][]float64
	PixelSize   float64
	HalfHeight  float64
	HalfWidth   float64
}

// New returns a new camera
func New(hSize float64, vSize float64, fieldOfView float64) Camera {
	transform := matrix.NewIdentity()
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
	c.PixelSize = c.HalfWidth * 2 / c.HSize
	return c
}

//RayForPixel returns the ray from the camera to point (x,y) on canvas
func RayForPixel(c Camera, x, y int) (*ray.Ray, error) {
	xOffset := (float64(x) + 0.5) * c.PixelSize
	yOffset := (float64(y) + 0.5) * c.PixelSize
	worldX := c.HalfWidth - xOffset
	worldY := c.HalfHeight - yOffset
	inverse, err := matrix.Inverse(c.Transform, 4)
	if err != nil {
		return nil, errors.New("could not compute inverse for matrix")
	}
	pixel := matrix.MultiplyWithTuple(inverse, tuple.Point(worldX, worldY, -1))
	origin := matrix.MultiplyWithTuple(inverse, tuple.Point(0, 0, 0))
	direction := tuple.Normalize(tuple.Subtract(pixel, origin))
	return &ray.Ray{
		Origin:    origin,
		Direction: direction,
	}, nil
}

//Render renders the world with a camera
func (c Camera) Render(w world.World) (*canvas.Canvas, error) {
	image := canvas.New(int(c.HSize), int(c.VSize))
	for y := 0; y < int(c.VSize); y++ {
		for x := 0; x < int(c.HSize); x++ {
			r, err := RayForPixel(c, x, y)
			if err != nil {
				return nil, fmt.Errorf("Could not find ray for pixel due to error %v", err)
			}
			color, err := world.ColorAt(w, *r)
			if err != nil {
				return nil, fmt.Errorf("Could not find color at ray due to error %v", err)
			}
			canvas.WritePixel(&image, x, y, *color)
		}
	}
	return &image, nil
}

func pixelSize(hSize int, vSize int, fieldOfView float64) float64 {
	halfView := math.Tan(fieldOfView / 2)
	aspect := float64(hSize / vSize)
	var halfWidth, halfHeight float64
	if aspect >= 1 {
		halfWidth = halfView
		halfHeight = halfView / aspect
	} else {
		halfWidth = halfView * aspect
		halfHeight = halfView
	}
	return halfWidth * 2 / halfHeight
}
