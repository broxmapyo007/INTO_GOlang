package main

import "fmt"

func main() {
	fmt.Println("loops : for loop only (it can use in multiple ways")

	//simple for loop- for <intiate>;<condition>;<increment>{}
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// for loop as while loop
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	//nested loops
	for i := 1; i < 10; i++ {
		for j := 1; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
