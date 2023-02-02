package shapes

import "math"

type Shape interface {
	Area() float64
}

// Rectangle shape struct
type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

func (r Rectangle) Area() float64 {
	return r.height * r.width
}

// Circle shape struct
type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return (c.radius * c.radius) * math.Pi
}
