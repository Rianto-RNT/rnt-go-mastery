package main

import "fmt"

func main() {

	// iota is number generator for constants which starts from zero
	// and is incremented by 1 automatically.

	const (
		c1 = iota
		c2 = iota
		c3 = iota
	)

	fmt.Println(c1, c2, c3) // 0 1 2

	const (
		c11 = iota
		c22
		c33
	)
	fmt.Println(c11, c22, c33)

	const (
		North = iota // default 0
		East
		South
		West
	)
	fmt.Println(West)

	const (
		a = (iota * 2) + 1 // 0
		b                  // 2
		c                  // 4

	)
	fmt.Println(a, b, c)

	// x = -1, y = -4, z = -5

	const (
		x = -(iota + 2) // -2
		_               // -3
		y               // -4
		z               // -5
	)

	fmt.Println(x, y, z)
}
