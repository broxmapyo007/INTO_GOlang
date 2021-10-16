// defer | recover | panic

package main

import "fmt"

func main() {

	/* use for function excution control
	   defer :
	   	*delay the function ,till other function done excution [LIFO]
	   	*clear connections,file open,etc
	   recover :
	   	*used after panic to normalise the function flow, it used with defer to recover
	   	 panic in goroutine
	   panic :
	   	*pani is similar to throwing exception ,
	   	*when it called other functions stops with normal flow and this excute first
	   	*deffered function are excuted normally

	*/
	fmt.Print("controll function flow \n\n")
	//normal flow
	f_run()
	s_run()

	//defer
	defer f_run()
	s_run()

	// pani,recover example in next code
}
func f_run() {
	fmt.Println("1st called")
}

func s_run() {
	fmt.Println("2nd called")
}
