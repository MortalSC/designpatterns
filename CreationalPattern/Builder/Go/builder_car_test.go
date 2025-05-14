package builder

import (
	"testing"
)

func TestBuilderCar(t *testing.T) {
	builder := NewCarStudio()
	builder.Brand("sky").Speed(120).Engine("audi")
	car := builder.Build()
	if car.Brand() != "sky" {
		t.Errorf("Expected brand to be 'sky', got '%s'", car.Brand())
	}
	if car.Speed() != 120 {
		t.Errorf("Expected speed to be 120, got %d", car.Speed())
	}
}

func TestBuilderCarMore(t *testing.T) {
	builder := NewCarStudio()
	builder.Brand("land").Speed(110).Engine("bmw")
	builder.Engine("man made").Brand("panda").Wheel(15)
	car := builder.Build()

	if car.Brand() != "panda" {
		t.Errorf("Expected brand to be 'panda', got '%s'", car.Brand())
	}
	if car.Speed() != 110 {
		t.Errorf("Expected speed to be 110, got %d", car.Speed())
	}
	car.Brief()
}
