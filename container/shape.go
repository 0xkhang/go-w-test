package container

import "math"

type Rectangle struct {
	width  float64
	height float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}

func (r Rectangle) Perimiter() float64 {
	return 2 * (r.width + r.height)
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}
