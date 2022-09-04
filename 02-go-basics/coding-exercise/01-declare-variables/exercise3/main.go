package main

import "fmt"

func main() {
	var a float64 = 7.1

	x, y := true, 3.7

	// Cannot assign more than 1 variable on left side of :=
	// a, x := 5.5, false
	a, x = 5.5, false

	_, _, _ = a, x, y

	fmt.Println("Result >>", a, x, y)
}
