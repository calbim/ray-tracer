package color

import (
	"testing"
)

func TestAddition(t *testing.T) {
	c1 := New(0.9, 0.6, 0.75)
	c2 := New(0.7, 0.1, 0.25)
	sum := c1.Add(c2)
	if !sum.Equals(New(1.6, 0.7, 1.0)) {
		t.Errorf("wanted c1+c2=%v, got %v", New(1.6, 0.7, 1.0), c1.Add(c2))
	}
}

func TestSubtraction(t *testing.T) {
	c1 := New(0.9, 0.6, 0.75)
	c2 := New(0.7, 0.1, 0.25)
	diff := c1.Subtract(c2)
	if !diff.Equals(New(0.2, 0.5, 0.5)) {
		t.Errorf("wanted c1-c2=%v, got %v", New(0.2, 0.5, 0.5), c1.Subtract(c2))
	}
}

func TestMultiplication(t *testing.T) {
	c := New(0.2, 0.3, 0.4)
	product := c.Multiply(2)
	if !product.Equals(New(0.4, 0.6, 0.8)) {
		t.Errorf("wanted c*2=%v, got %v", New(0.4, 0.6, 0.8), c.Multiply(2))
	}
}

func TestMultipicationWithColor(t *testing.T) {
	c1 := New(1, 0.2, 0.4)
	c2 := New(0.9, 1, 0.1)
	product := c1.MultiplyColor(c2)
	if !product.Equals(New(0.9, 0.2, 0.04)) {
		t.Errorf("wanted c1*c2=%v, got %v", New(0.9, 0.2, 0.04), product)
	}
}
