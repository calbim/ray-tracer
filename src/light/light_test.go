package light

import (
	"testing"

	"github.com/calbim/ray-tracer/src/color"
	"github.com/calbim/ray-tracer/src/tuple"
)

func TestPointLight(t *testing.T) {
	intensity := color.New(1, 1, 1)
	position := tuple.Point(0, 0, 0)
	pointLight := PointLight(position, intensity)
	if pointLight.Intensity != intensity || pointLight.Position != position {
		t.Errorf("wanted point light to be %v, got %v", Light{Intensity: intensity, Position: position}, pointLight)
	}
}
