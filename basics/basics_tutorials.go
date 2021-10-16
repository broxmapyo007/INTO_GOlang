package main

import (
	"fmt"
)

func main() {
	//single cmt
	/* multi
	lne comment */

	//print statements ,print simply
	fmt.Print(2, "-integer   | ")

	//print in new line, auto add \n at end of print
	fmt.Println(" name : king")                              //will continue after integer
	fmt.Println(" name :" + "kong")                          // will be on new line
	fmt.Println("len of string:king kong", len("king kong")) //len of sting

	//print formatter ( f,T,t,d,v,s)
	/*
	   %f : float
	   %t : bool values
	   %T : Type
	   %b : binary
	   %v : value
	   %s : string
	   %x : hexcode
	   %d : integer
	   %c : key code eg 33 is ?
	*/
	fmt.Printf("%s  : string format\n", "kingo") //string
	fmt.Printf("%d  : integer format\n", 234)    //integers
	fmt.Printf("%f  : float format\n", 3.124)    //float
	fmt.Printf("%t  : bool format\n", true)      //BOOL
	fmt.Printf("%T  : Type type\n", 3.124)       //type
	fmt.Printf("%b  : float format\n", 3)        //binary
	fmt.Printf("%c  : key format\n", 33)         //key code
	fmt.Printf("%v  : value format\n", 2312)     //value
	fmt.Printf("%x  : hexcode format\n", 3)      //hex

	// varible : integer , float, string, bool
	// var|const <var_name> <type> = <value>
	// <var_name> := <value>

	x := 10
	fmt.Print(x)

	y := "string"
	var z string = "string2"
	println(y, z)

	isbool := true
	isfalse := false
	println(isbool, isfalse)

}
