package main

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

// Area returns the area of a rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Perimeter returns the perimeter of a rectangle
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

// Circle represents a circle...
type Circle struct {
	Radius float64
}

// Area returns the area of a rectangle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
