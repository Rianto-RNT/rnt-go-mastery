package main

import "fmt"

func main() {
	var (
		a int
		b float64
		c bool
		d string
	)

	x, y, z := 20, 15.5, "RNT"

	fmt.Println(a, b, c, d)
	fmt.Println("x value is >>>", x, " y value is >>>", y, " z value is >>>", z)

	// _, _, _, _, _, _, _ = a, b, c, d, x, y, z //using blank identifier to mute the unused variable error
	// _ stays always on the left side of =
}
