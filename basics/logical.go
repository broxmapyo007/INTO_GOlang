package main

import "fmt"

func main() {
	var a int = 5
	var b float32 = 10.5
	fmt.Print(a)
	fmt.Println("float", b)
	const pi float64 = 3.14151

	x, y := 20, 5

	// math operators

	fmt.Println("x+y : ", x+y)
	fmt.Println("x-y : ", x-y)
	fmt.Println("x*y : ", x*y)
	fmt.Println("x/y : ", x/y)

	// bools
	istrue := true
	isfale := false
	//logical operators

	fmt.Println("and ", istrue && isfale)
	fmt.Println("or ", istrue || isfale)
	fmt.Print("not ", !isfale)
}
