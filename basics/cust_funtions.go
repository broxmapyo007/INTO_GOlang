package main

import "fmt"

func main() {
	//math basic function

	x, y := 5, 6

	fmt.Println(add(x, y))

	//recursive function
	num := 5
	fmt.Println("factorial of 5:", factorial(num))
}

func add(num1, num2 int) int {
	return num1 + num2

}

func factorial(num int) int {
	if num == 0 {
		return 1
	}
	return num * factorial(num-1)

}
