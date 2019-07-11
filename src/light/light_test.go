package light

import (
	"testing"
	"github.com/calbim/ray-tracer/src/tuple"
)

func TestPointLight(t *testing.T){
	intensity := tuple.Color(1,1,1)
	position := tuple.Point(0,0,0)
	pointLight := PointLight{Intensity:intensity, Position:position}
	if pointLight.Intensity != intensity || pointLight.Position!=position {
		t.Errorf("Point light is initialized incorrectly")
	}
}