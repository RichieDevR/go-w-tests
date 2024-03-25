package main

import (
	"math"
	"testing"
)

// Table Driven Tests for Perimeter method on shape interface.
// Tests that code for shape.Perimeter() functions properly in v2.
// similar to shape.Area() tests seen here and in v1. this was done
// as a challenge to myself and to drive home what was learned in the
// original lesson
func TestPerimeter(t *testing.T) {
	perimeterTests := []struct {
		shape Shape

		name string

		hasPerimeter float64
	}{
		{name: "Rectangle p", shape: Rectangle{12, 6}, hasPerimeter: 36.0},

		{name: "Circle p", shape: Circle{10}, hasPerimeter: 62.83},

		{name: "Triangle p", shape: Triangle{12, 6}, hasPerimeter: 31.42},
	}

	for _, tt := range perimeterTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Perimeter()

			if math.Abs(got-tt.hasPerimeter) > 0.01 {
				t.Errorf("%s %s got %.2f want %.2f", tt.name, tt.shape, got, tt.hasPerimeter)
			}
		})
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		shape Shape

		name string

		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{12, 6}, hasArea: 72.0},

		{name: "Circle", shape: Circle{10}, hasArea: 314.1592653589793},

		{name: "Triangle", shape: Triangle{12, 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()

			if got != tt.hasArea {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
			}
		})
	}
}
