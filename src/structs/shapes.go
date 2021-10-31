package structs

import "math"

type Rectangle struct {
	width  float64
	height float64
}

type Circle struct {
	Redius float64
}

type Shape interface {
	Area() float64
}

type Triangle struct {
	Base   float64
	Height float64
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.width + rectangle.height)
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (c Circle) Area() float64 {
	return math.Pi * c.Redius * c.Redius
}

func (t Triangle) Area() float64 {
	return 0
}
