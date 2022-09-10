package main

import "fmt"

func main() {
	s := "أهلا وسهلا RNT!"

	// converting string to rune slice
	r := []rune(s)

	// Printing out the rune slice
	fmt.Printf("%#v\n", r)

	// iterating over the rune slice and printing out each index and rune in the rune slice
	for i, v := range r {
		fmt.Printf("%d -> %c\n", i, v)
	}
}
