package geometry

import "math"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	width  float64
	height float64
}
type Triangle struct {
	height float64
	base   float64
}

func (t Triangle) Area() (P float64) {
	return t.height * t.base * 0.5
}
func (t Triangle) Perimeter() (P float64) {
	return t.height + t.base + math.Sqrt(math.Pow(t.height, 2)+math.Pow(t.base, 2))
}
func (r Rectangle) Area() (P float64) {
	return r.width * r.height
}
func (r Rectangle) Perimeter() (P float64) {
	return (r.width + r.height) * 2
}

type Circle struct {
	radius float64
}

func (c Circle) Area() (P float64) {
	return math.Pi * c.radius * c.radius
}
func (c Circle) Perimeter() (P float64) {
	return c.radius * 2 * math.Pi
}
