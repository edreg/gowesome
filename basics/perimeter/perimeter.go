package perimeter

import "math"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	width  float64
	height float64
}

type Circle struct {
	radius float64
}

type SameSidedTriangle struct {
	base   float64
	height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (r Rectangle) Perimeter() float64 {
	return 2*r.width + 2*r.height
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c Circle) Perimeter() float64 {
	return 2 * c.radius * math.Pi
}

func (t SameSidedTriangle) Area() float64 {
	return t.height * t.base * 0.5
}

func (t SameSidedTriangle) Perimeter() float64 {
	return 2 * t.height * t.base
}
