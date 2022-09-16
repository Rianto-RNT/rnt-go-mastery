package main

import "fmt"

// 1. Create your own struct type called person that stores the following data: name, age and a list with favorite colors.
type person struct {
	name   string
	age    int
	colors []string
}

func main() {
	// 2. Declare and initialize two values of type person, one called me and another called you.
	me := person{name: "Rian", age: 29, colors: []string{"black", "white"}}
	you := person{name: "Jet", age: 23, colors: []string{"blue", "white"}}

	// 3. Print out the struct values.
	fmt.Printf("%+v\n", me)
	fmt.Printf("%+v\n", you)
}