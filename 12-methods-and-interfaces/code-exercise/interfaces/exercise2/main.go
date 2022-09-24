package main

import "fmt"

func main() {
	// 1. Declare a variable called empty of type empty interface. Print out its type.
	var empty interface{}
	fmt.Printf("%T\n", empty) // <nill>

	// 2. Assign an int value to the variable called empty. Print out its type.
	empty = 5
	fmt.Printf("%T\n", empty) // int

	// 3. Assign a float64 value to empty. Print out its type.
	empty = 9.99
	fmt.Printf("%T\n", empty) // float

	// 4. Assign an int slice value to empty. Print out its type.
	empty = []int{1, 3, 8}
	fmt.Printf("%T\n", empty) // []int

	// 5. Add a new int value to the slice (empty variable).
	empty = append(empty.([]int), 10) // Assertion Type

	// 6. Print out the slice (empty variable).
	fmt.Printf("%v\n", empty)
}
