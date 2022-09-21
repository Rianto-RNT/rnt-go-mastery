package main

import "fmt"

func swapValue(a, b *float64) {
	*a, *b = *b, *a
}

func main() {
	x, y := 5.5, 8.8
	swapValue(&x, &y)

	fmt.Printf("x is %v and y is %v\n", x, y)
}
