package main

import "fmt"

func main() {
	x, y, z := 10, 15.5, "Rianto"
	score := []int{10, 20, 30}

	// 1.
	fmt.Printf("x is %d, y is %f, z is %s\n", x, y, z)
	fmt.Printf("score is %#v\n", score)

	// 2.
	fmt.Printf("z is %q\n", z)

	// 3.
	fmt.Printf("x is %v, y is %v, z is %v, score is %v\n", x, y, z, score)

	// 4.
	fmt.Printf("y type: %T, score type: %T\n", y, score)
}
