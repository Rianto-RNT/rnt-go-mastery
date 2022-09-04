package main

import "fmt"

// ERROR -> := is not allowed in package scope (only local scope)
// version := "3.1"
var version = "3.1"

func main() {
	// ERROR -> A string is initialized only using double quotes ("")
	// name = 'Golang'

	name := "rnt"
	fmt.Println(name)
}
