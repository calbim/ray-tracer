package intersections
 
import (
	"github.com/calbim/ray-tracer/src/sphere"
	"testing"
)
func TestIntersectionObject(t *testing.T){
	s:= sphere.New()
	i := Intersection{3.5,s}
	if i.Value != 3.5 {
		t.Errorf("The intersection point should be %f",3.5)
	}
	if i.Object != s{
		t.Errorf("The intersected object should be %v", s)
	}
}