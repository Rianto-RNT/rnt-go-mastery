package main

import "fmt"

func main() {
	name := "Rianto"
	fmt.Println(&name)

	var x int = 2
	ptr := &x
	fmt.Printf(`ptr is type of %T with a value of %v\n`, ptr,ptr)
	fmt.Printf("Address of x is %p\n", &x)
	
	var ptr1 *float64
	_ = ptr1
	
	p := new(int)
	x =100
	p = &x	
	fmt.Printf("Type of p is: %T with value is %v\n", p, p)
	fmt.Printf("Address of x is %p\n", &x)

	*p = 90 // equivalent to x = 90
	fmt.Println(x, *p)
	fmt.Println("p==x:", *p==x)

	*p = 10 // x equivalent 10
	*p = *p / 2 // x / 2
	fmt.Println(x)

	// &value >>> pointer
	// *pointer >>> value
}