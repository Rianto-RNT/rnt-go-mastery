package main

import "fmt"

func main() {
	var numbers [4]int

	fmt.Printf("%v\n", numbers)
	fmt.Printf("%#v\n", numbers)

	var a1 = [4]float64{}
	fmt.Printf("%#v\n", a1)

	var a2 = [3]int{-10, 1, 100}
	fmt.Printf("%#v\n", a2)

	a3 := [4]string{"Rian", "Ryan", "rianto", "RNT"}
	fmt.Printf("%#v\n", a3)

	a4 := [4]string{"x", "y"}
	fmt.Printf("%#v\n", a4)

	a5 := [...]int{1, 2, 5, 1, -10, 66}
	fmt.Printf(">> %#v\n", a5)
	fmt.Printf(">> the length is >> %d\n", len(a5))
}
