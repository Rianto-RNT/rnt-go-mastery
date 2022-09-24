package main

import "fmt"

type cube struct {
	edge float64
}

func volume(c cube) float64 {
	return c.edge * c.edge * c.edge
}

func main() {
	var x interface{}
	x = cube{edge: 5}

	// Type Assertion
	v := volume(x.(cube))

	fmt.Printf("Cube volume: %v\n", v) // Cube volume: 125
}
