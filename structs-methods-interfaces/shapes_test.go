package main

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	rect := Rectangle{10.0, 10.0}

	got := Perimeter(rect)

	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

type Shape interface {
	Area() float64
}

func TestArea(t *testing.T) {
	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()

		got := shape.Area()

		if got != want {
			t.Errorf("go %g want %g", got, want)
		}
	}

	t.Run("area of rectangle", func(t *testing.T) {
		rectangle := Rectangle{12.0, 6.0}

		checkArea(t, rectangle, 72.0)
	})

	t.Run("area of a circle", func(t *testing.T) {
		circle := Circle{10.0}

		checkArea(t, circle, 314.1592653589793)
	})
}
