package main

import "fmt"

func main() {
	//panic and recover demostrate by divide 2 numbers

	fmt.Println(div(3, 0))
	fmt.Println(div(9, 3))
	dempanic()
}

func div(n1, n2 int) int {
	defer func() {
		fmt.Println(recover())
	}()

	solution := n1 / n2
	return solution

}

func dempanic() {
	defer func() {
		fmt.Println(recover())
	}()
	panic("Error of Panic")
}
