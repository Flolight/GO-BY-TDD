package structs

import "math"

type Shape interface {
	Area() float64
}

// Rectangle shape model
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle shape model
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Perimeter computes the perimeter of a shape of <width> * <height>
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

// Area computes the area of a shape of <width> * <height>
func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}
