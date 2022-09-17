package main

import "fmt"

// Create a function called cube() that takes a parameter of type float64 and returns the cube of that parameter (the parameter to the power of 3).
func cube(n float64) float64 {
	return n * n * n
}

func main() {
	x := cube(3)
	fmt.Println(x)
}
