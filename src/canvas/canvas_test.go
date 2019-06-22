package canvas

import (
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
