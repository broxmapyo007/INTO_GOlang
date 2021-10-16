package main

import (
	"fmt"
	"math"
)

func main() {
	// interface implemented by structure
	fmt.Println("interface ....")

	rec := Rectangle{10, 20}
	cir := Circle{4}
	fmt.Println("area of rectangle :", getArea(rec))
	fmt.Println("area of circle :", getArea(cir))
}

type Shape interface {
	area() float64
}

type Rectangle struct {
	height float64
	width  float64
}

type Circle struct {
	radius float64
}

func (r1 Rectangle) area() float64 {
	return r1.height * r1.width
}

func (r Circle) area() float64 {
	return math.Pi * math.Pow(r.radius, 2)
}

func getArea(s Shape) float64 {
	return s.area()
}
