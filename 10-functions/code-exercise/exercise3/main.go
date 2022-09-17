package main

import (
	"fmt"
	"os"
	"strconv"
)

func myFunc(n string) int {
	// Convert string to int
	i, err1 := strconv.Atoi(n)

	// Error handling
	if err1 != nil {
		fmt.Printf("Cannot convert %q to int.", n)
		os.Exit(1) // Exit the program

	}
	ii, _ := strconv.Atoi(n + n)
	iii, _ := strconv.Atoi(n + n + n)

	return i + ii + iii

}

func main() {
	fmt.Println(myFunc("5")) // >>> result 615
	fmt.Println(myFunc("3")) // >>> result 369
}
