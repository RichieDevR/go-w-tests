package main

import "math"

// Shape is implemented by anything that can tell us its Area or Perimeter.
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Rectangle has the dimensions of a rectangle and represents a rectangle.
type Rectangle struct {
	Width float64

	Height float64
}

// Area returns the area of the rectangle.
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Perimeter returns the perimeter of a rectangle.
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Height + r.Width)
}

// Circle represents a circle...
type Circle struct {
	Radius float64
}

// Area returns the area of the circle.
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Perimeter returns the perimeter of a circle.
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Triangle represents a triangle...
type Triangle struct {
	Base float64

	Height float64
}

// Area returns the area of the triangle.
func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

// Perimeter returns the perimeter of the triangle.
func (t Triangle) Perimeter() float64 {
	getHypotenuse := func(base, height float64) float64 {
		return math.Sqrt(base*base + height*height)
	}

	hypotenuse := getHypotenuse(t.Base, t.Height)

	perimeter := t.Base + t.Height + hypotenuse

	return perimeter
}
