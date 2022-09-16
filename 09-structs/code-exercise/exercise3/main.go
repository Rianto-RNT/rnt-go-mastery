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

	for index, color := range me.colors {
		fmt.Printf("%d -> %q\n", index, color)
	}

	_ = you
}