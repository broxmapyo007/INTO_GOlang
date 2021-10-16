package main

import "fmt"

func main() {
	fmt.Println("dictionary same use")

	//maps : dictionary is key : value pairs
	studentage := make(map[string]int)

	studentage["a1"] = 10
	studentage["a2"] = 10
	studentage["a3"] = 10

	fmt.Println(studentage)
	fmt.Println(studentage["a2"])
	fmt.Println(len(studentage))

	//nested map
	superhero :=
		map[string]map[string]string{
			"superman": map[string]string{
				"realname": "clark",
			},

			"batman": {
				"power": "rich man",
			},
		}

	if temp, hero := superhero["batman"]; hero {
		fmt.Println(temp, temp["power"])
	}
	fmt.Println(superhero, superhero["superman"]["realname"])

}
