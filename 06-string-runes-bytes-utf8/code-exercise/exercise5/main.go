package main

import "fmt"

func main() {
	s := "你好 Go!"

	// converting string to byte slice
	b := []byte(s)

	// printing out the byte slice
	fmt.Printf("%#v\n", b)

	// iterating over the byte slice and printing out each index and byte in the byte slice
	for i, v := range b {
		fmt.Printf("%d -> %d\n", i, v)
	}
}
