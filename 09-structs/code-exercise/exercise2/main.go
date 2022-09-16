package main

import "fmt"

type person struct {
	name   string
	age    int
	colors []string
}

func main() {
	me := person{name: "Rian", age: 29, colors: []string{"black", "white"}}
	you := person{name: "Jet", age: 23, colors: []string{"blue", "white"}}

	// 1. Change the name or the struct value called me to "Morrison"
	me.name = "Morrison"

	// 2. Take in a new variable the favorites colors of struct value called you. Print out the type and the value of the new variable.
	var colors []string = you.colors
	fmt.Printf("Type: %T, value: %v\n", colors, colors)

	// 3. Add a new favorite color to the second struct value called you.
	colors = append(colors, "Grey")
	you.colors = colors

	// 4. Print out the struct values.
	fmt.Printf("%+v\n", me)
	fmt.Printf("%+v\n", you)
}