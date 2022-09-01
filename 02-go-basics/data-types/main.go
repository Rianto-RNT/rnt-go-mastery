package main

import "fmt"

func main() {
	// int type
	var i1 int8 = -128
	fmt.Printf("%T\n", i1)

	var i2 uint16 = 65535
	fmt.Printf("%T\n", i2)

	// Float type
	var f1, f2, f3 float64 = 1.1, -.2, 5.0
	fmt.Printf("%T %T %T\n", f1, f2, f3)

	// Rune Type
	var r rune = 'f'
	fmt.Printf("%T\n", r)
	fmt.Println(r)
	fmt.Printf("%x\n", r)

	// Bool Type
	var b bool = true
	fmt.Printf("%T\n", b)

	// String type
	var s string = "Hello RNT"
	fmt.Printf("%T\n", s)

}
