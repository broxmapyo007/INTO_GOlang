package main

import "fmt"

func main() {
	rec1 := Rectangle{10, 5} //{height=10,width=5}
	fmt.Println(rec1.height)

	fmt.Println("area of rectangle :", rec1.area())
}

type Rectangle struct {
	height float64
	width  float64
}

func (rect *Rectangle) area() float64 {
	return rect.height * rect.width
}
