package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	//file write
	file, err := os.Create("sample.txt")

	if err != nil {
		log.Fatal(err)
	}

	file.WriteString("hii file is writing by golang")
	file.Close()

	//file read
	stream, err := ioutil.ReadFile("sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	read1 := string(stream)

	fmt.Println(read1)
}
