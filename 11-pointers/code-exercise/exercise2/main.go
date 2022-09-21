package main

import "fmt"

func main() {
	x, y := 10, 2
	ptrX, ptrY := &x, &y

	// Declare a new variable called z and initialize the variable by dividing x by y through the pointers.
	z := *ptrX / *ptrY
	fmt.Printf("z is %d\n", z)
}
