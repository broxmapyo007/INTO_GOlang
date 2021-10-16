package main

import "fmt"

func main() {
	var eventname [5]int

	eventname[0] = 1
	eventname[1] = 2
	eventname[2] = 3
	eventname[3] = 4
	eventname[4] = 5

	fmt.Println(eventname[1])

	//2nd array way iterate in array
	eventcol := [5]int{1, 2, 3, 4, 5}

	for i, val := range eventcol {
		fmt.Println(i, "array ", val)
	}

	//indice,slice

	slice_ls := eventcol[2:5] // [0:] , [:5]
	fmt.Println(slice_ls)

	slice2 := make([]int, 5, 10)

	copy(slice2, slice_ls)
	fmt.Println(slice2)

	slice3 := append(eventcol[0:3]) //append(slice2, 3, 0, 10) : add slice +3,0,10 in array
	fmt.Println(slice3)

	//list
	fmt.Println(addall(10, 20, 50, 100))
}

func addall(args ...int) int {
	sum := 0
	for _, val := range args {
		sum += val
	}
	return sum
}
