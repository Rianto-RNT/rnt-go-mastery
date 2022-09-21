package main

import "fmt"

func main() {
	x := 10.10
	fmt.Println(x)

	ptr := &x
	fmt.Printf("ptr type %T, Value of ptr: %v\n", ptr, ptr)

	fmt.Printf("The address of ptr: %p\n", ptr)
	fmt.Printf("The value of x through the pointer: %f\n", *ptr)
}
