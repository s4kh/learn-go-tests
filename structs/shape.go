package structs

import "math"

type Shape interface {
	Area() float64
	GetName() string
}

type Rectangle struct {
	Height float64
	Width  float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

func (r Rectangle) GetName() string {
	return "Rectangle"
}

func (c Circle) GetName() string {
	return "Circle"
}

func (t Triangle) GetName() string {
	return "Triangle"
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) / 2
}

func (c Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}
