package main

import (
	structs "learn-go-with-test/struct_method_interface/struct"
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := structs.Rectangle{Width: 10, Height: 10.0}
	got := Perimeter(rectangle)
	want := 40.0
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	areaTest := []struct {
		name  string
		shape structs.Shape
		want  float64
	}{
		{
			name:  "test area rectangle",
			shape: structs.Rectangle{Width: 10, Height: 10.0},
			want:  100.0,
		},
		{
			name:  "test area circle",
			shape: structs.Circle{Radius: 10},
			want:  314.1592653589793,
		},
		{
			name:  "test area triangle",
			shape: structs.Triangle{Base: 10.0, Height: 10.0},
			want:  50.0,
		},
	}

	for _, test := range areaTest {
		got := test.shape.Area()
		if got != test.want {
			t.Errorf("%#v got %g want %g", test.shape, got, test.want)
		}
	}
}
