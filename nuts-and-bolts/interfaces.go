package nutsandbolts

import (
	"math"
)

// source: https://gobyexample.com/interfaces

type geometry interface {
	area() float64
	perimeter() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

// To implement an interface in Go, we just need to implement all the methods in the interface.
// Here we implement geometry for the rect.

func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perimeter() float64 {
	return 2*r.width + 2*r.height
}

// ...and for the circle
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func Measure(g geometry) (float64, float64) {
	return g.area(), g.perimeter()
}

func DetectCircle(g geometry) bool {
	// g.(circle) return the casted object and bool indicating if the casting was successful
	_, ok := g.(circle)
	return ok
}