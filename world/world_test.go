package world

import (
	"testing"
)

func TestDefaultWorldHasNothing(t *testing.T){
	w := World{}
	if w.Objects != nil || w.Light != nil {
		t.Errorf("A default world object should not have any objects in it or a light source.")
	}
}