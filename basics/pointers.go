package main

import (
	"fmt"
)

func main() {
	fmt.Print("pointer are used to pass,store or reference memory space address \n\n")

	a := 5
	b := 10
	fmt.Println("a", a, "b", b)
	fmt.Println("> address a:", &a, "\n> address b:", &b)

	//it can use reference and change value
	changeval(&a, b)
	fmt.Println("a", a, &a)
	fmt.Println("b", b, &b)
}

func changeval(x *int, y int) {
	y = 26
	*x = 22

}
