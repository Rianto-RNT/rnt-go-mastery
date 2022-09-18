package main

import "fmt"

func f1(a int) {
	fmt.Println(a)
}

func main() {
	x := f1
	fmt.Printf("%T\n", x) // >>> func(int)
	x(8)
}
