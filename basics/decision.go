package main

import "fmt"

func print(str string) {
	// use only to print string
	fmt.Println(str)
}
func main() {
	//if else, switch

	age := 10
	if age > 18 {
		print("age is eligible to vote")
	} else {
		print("age not eligible to vote")
	}

	switch age {
	case 10:
		print("in school")
	case 16:
		print("college approved")
	case 18:
		print("voter id can issue")
	case 20:
		print("college complete")

	}
}
